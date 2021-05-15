package mempool_old

import (
	"sort"
	"sync"
	"time"

	"github.com/iotaledger/hive.go/logger"
	"github.com/iotaledger/wasp/packages/chain"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/coretypes/request"
	"github.com/iotaledger/wasp/packages/state"
	"github.com/iotaledger/wasp/packages/vm/core/blocklog"
)

type mempool struct {
	mutex       sync.RWMutex
	incounter   int
	outcounter  int
	stateReader state.StateReader
	requestRefs map[coretypes.RequestID]*requestRef
	chStop      chan bool
	blobCache   coretypes.BlobCache
	log         *logger.Logger
}

type requestRef struct {
	req             coretypes.Request
	whenMsgReceived time.Time
	seen            map[uint16]bool
}

const constSolidificationLoopDelay = 200 * time.Millisecond

var _ chain.MempoolOld = &mempool{}

func New(stateReader state.StateReader, blobCache coretypes.BlobCache, log *logger.Logger) chain.MempoolOld {
	ret := &mempool{
		stateReader: stateReader,
		requestRefs: make(map[coretypes.RequestID]*requestRef),
		chStop:      make(chan bool),
		blobCache:   blobCache,
		log:         log.Named("m"),
	}
	go ret.solidificationLoop()
	return ret
}

func (m *mempool) ReceiveRequest(req coretypes.Request) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	// only allow off-ledger requests with valid signature
	if offLedgerReq, ok := req.(*request.RequestOffLedger); ok {
		if !offLedgerReq.VerifySignature() {
			m.log.Errorf("ReceiveRequest.VerifySignature: invalid signature")
			return
		}
	}
	id := req.ID()
	if blocklog.IsRequestProcessed(m.stateReader, &id) {
		return
	}
	if _, ok := m.requestRefs[req.ID()]; ok {
		return
	}

	// attempt solidification for those requests that do not require blobs
	// instead of having to wait for the solidification goroutine to kick in
	// also weeds out requests with solidification errors
	_, err := req.SolidifyArgs(m.blobCache)
	if err != nil {
		m.log.Errorf("ReceiveRequest.SolidifyArgs: %s", err)
		return
	}

	m.incounter++
	tl := req.TimeLock()
	if tl.IsZero() {
		m.log.Debugf("IN MEMPOOL %s (+%d / -%d)", req.ID(), m.incounter, m.outcounter)
	} else {
		m.log.Debugf("IN MEMPOOL %s (+%d / -%d) timelocked for %v", req.ID(), m.incounter, m.outcounter, tl.Sub(time.Now()))
	}
	m.requestRefs[req.ID()] = &requestRef{
		req:             req,
		whenMsgReceived: time.Now(),
		seen:            make(map[uint16]bool),
	}
}

func (m *mempool) MarkSeenByCommitteePeer(reqid *coretypes.RequestID, peerIndex uint16) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if _, ok := m.requestRefs[*reqid]; !ok {
		m.requestRefs[*reqid] = &requestRef{
			seen: make(map[uint16]bool),
		}
	}
	m.requestRefs[*reqid].seen[peerIndex] = true
}

func (m *mempool) ClearSeenMarks() {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	for _, rec := range m.requestRefs {
		rec.seen = make(map[uint16]bool)
	}
}

func (m *mempool) RemoveRequests(reqs ...coretypes.RequestID) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	for _, rid := range reqs {
		if _, ok := m.requestRefs[rid]; ok {
			m.outcounter++
		}
		delete(m.requestRefs, rid)
		m.log.Debugf("OUT MEMPOOL %s (+%d / -%d)", rid, m.incounter, m.outcounter)
	}
}

const timeAheadTolerance = 1000 * time.Nanosecond

func isRequestReady(ref *requestRef, seenThreshold uint16, nowis time.Time) bool {
	if ref.req == nil {
		return false
	}
	if len(ref.seen) < int(seenThreshold) {
		return false
	}
	if _, paramsReady := ref.req.Params(); !paramsReady {
		return false
	}
	if !ref.req.TimeLock().IsZero() {
		timeBaseline := nowis.Add(timeAheadTolerance)
		if ref.req.TimeLock().After(timeBaseline) {
			return false
		}
	}
	return true
}

func (m *mempool) GetReadyList(seenThreshold ...uint16) []coretypes.Request {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	var thr uint16
	if len(seenThreshold) > 0 {
		thr = seenThreshold[0]
	}
	ret := make([]coretypes.Request, 0, len(m.requestRefs))
	nowis := time.Now()
	for _, ref := range m.requestRefs {
		if isRequestReady(ref, thr, nowis) {
			ret = append(ret, ref.req)
		}
	}
	sort.Slice(ret, func(i, j int) bool {
		return ret[i].Order() < ret[j].Order()
	})
	return ret
}

// GetRequestsByIDs returns slice with requests. Corresponding nil if not ready
func (m *mempool) GetRequestsByIDs(nowis time.Time, reqids ...coretypes.RequestID) []coretypes.Request {
	ret := make([]coretypes.Request, len(reqids))
	for i := range ret {
		reqref, ok := m.requestRefs[reqids[i]]
		if !ok {
			continue
		}
		if !isRequestReady(reqref, 0, nowis) {
			continue
		}
		ret[i] = reqref.req
	}
	return ret
}

func (m *mempool) GetReadyListFull(seenThreshold ...uint16) []*chain.ReadyListRecord {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	var thr uint16
	if len(seenThreshold) > 0 {
		thr = seenThreshold[0]
	}
	ret := make([]*chain.ReadyListRecord, 0, len(m.requestRefs))
	nowis := time.Now()
	for _, ref := range m.requestRefs {
		if isRequestReady(ref, thr, nowis) {
			rec := &chain.ReadyListRecord{
				Request: ref.req,
				Seen:    make(map[uint16]bool),
			}
			for p := range ref.seen {
				rec.Seen[p] = true
			}
			ret = append(ret, rec)
		}
	}
	sort.Slice(ret, func(i, j int) bool {
		return ret[i].Request.Order() < ret[j].Request.Order()
	})
	return ret
}

func (m *mempool) TakeAllReady(nowis time.Time, reqids ...coretypes.RequestID) ([]coretypes.Request, bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	ret := make([]coretypes.Request, len(reqids))
	for i := range reqids {
		ref, ok := m.requestRefs[reqids[i]]
		if !ok {
			return nil, false
		}
		if !isRequestReady(ref, 0, nowis) {
			return nil, false
		}
		ret[i] = ref.req
	}
	return ret, true
}

func (m *mempool) HasRequest(id coretypes.RequestID) bool {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	rec, ok := m.requestRefs[id]
	return ok && rec.req != nil
}

// Stats return total number, number with messages, number solid
func (m *mempool) StatsOld() (int, int, int) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	total := len(m.requestRefs)
	withMsg, solid := 0, 0
	for _, ref := range m.requestRefs {
		if ref.req != nil {
			withMsg++
			if isSolid, _ := ref.req.SolidifyArgs(m.blobCache); isSolid {
				solid++
			}
		}
	}
	return total, withMsg, solid
}

func (m *mempool) Stats() chain.MempoolStatsOld {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	total := len(m.requestRefs)
	withMsg, solid := 0, 0
	for _, ref := range m.requestRefs {
		if ref.req != nil {
			withMsg++
			if isSolid, _ := ref.req.SolidifyArgs(m.blobCache); isSolid {
				solid++
			}
		}
	}
	return chain.MempoolStatsOld{
		Total:        total,
		WithMessages: withMsg,
		Solid:        solid,
		InCounter:    m.incounter,
		OutCounter:   m.outcounter,
	}
}

func (m *mempool) Close() {
	close(m.chStop)
}

func (m *mempool) solidificationLoop() {
	for {
		select {
		case <-m.chStop:
			return
		default:
			m.doSolidifyRequests()
		}
		time.Sleep(constSolidificationLoopDelay)
	}
}

func (m *mempool) doSolidifyRequests() {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	for _, ref := range m.requestRefs {
		if ref.req != nil {
			_, _ = ref.req.SolidifyArgs(m.blobCache)
		}
	}
}
