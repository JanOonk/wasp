package commiteeimpl

import (
	"fmt"
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address"
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/iotaledger/wasp/packages/committee"
	"github.com/iotaledger/wasp/packages/util"
	"github.com/iotaledger/wasp/plugins/peering"
	"github.com/iotaledger/wasp/plugins/publisher"
	"time"
)

func init() {
	committee.ConstructorNew = newCommitteeObj
}

func (c *committeeObj) IsOpenQueue() bool {
	if c.isOpenQueue.Load() {
		return true
	}
	c.mutexIsReady.Lock()
	defer c.mutexIsReady.Unlock()

	return c.checkReady()
}

// implements Committee interface

func (c *committeeObj) Params() *committee.Parameters {
	return c.params
}

func (c *committeeObj) SetReadyStateManager() {
	c.mutexIsReady.Lock()
	defer c.mutexIsReady.Unlock()

	c.isReadyStateManager = true
	c.log.Debugf("State Manager object was created")
	c.checkReady()
}

func (c *committeeObj) SetReadyConsensus() {
	c.mutexIsReady.Lock()
	defer c.mutexIsReady.Unlock()

	c.isReadyConsensus = true
	c.log.Debugf("Consensus object was created")
	c.checkReady()
}

func (c *committeeObj) SetInitConnectPeriodOver() {
	c.mutexIsReady.Lock()
	defer c.mutexIsReady.Unlock()

	c.isInitConnectPeriodOver = true
	c.log.Debugf("init connect period is over")
	c.checkReady()
}

func (c *committeeObj) SetQuorumOfConnectionsReached() {
	c.mutexIsReady.Lock()
	defer c.mutexIsReady.Unlock()

	c.isQuorumOfConnectionsReached = true
	c.log.Debugf("quorum of connections has been reached")
	c.checkReady()
}

func (c *committeeObj) checkReady() bool {
	if c.isReadyConsensus && c.isReadyStateManager && c.isInitConnectPeriodOver && c.isQuorumOfConnectionsReached {
		c.isOpenQueue.Store(true)
		c.startTimer()
		c.onActivation()

		c.log.Infof("committee now is fully initialized")
		publisher.Publish("active_committee", c.address.String())
	}
	return c.isReadyConsensus && c.isReadyStateManager && c.isInitConnectPeriodOver && c.isQuorumOfConnectionsReached
}

func (c *committeeObj) startTimer() {
	go func() {
		tick := 0
		for c.isOpenQueue.Load() {
			time.Sleep(c.params.TimerTickPeriod)
			c.ReceiveMessage(committee.TimerTick(tick))
			tick++
		}
	}()
}

func (c *committeeObj) Dismiss() {
	c.log.Infof("Dismiss committee for %s", c.address.String())

	c.dismissOnce.Do(func() {
		c.isOpenQueue.Store(false)
		c.dismissed.Store(true)

		close(c.chMsg)

		for _, pa := range c.peers {
			if pa != nil {
				peering.StopUsingPeer(pa.PeeringId())
			}
		}
	})

	publisher.Publish("dismissed_committee", c.address.String())
}

func (c *committeeObj) IsDismissed() bool {
	return c.dismissed.Load()
}

func (c *committeeObj) Address() *address.Address {
	return &c.address
}

func (c *committeeObj) OwnerAddress() *address.Address {
	return &c.ownerAddress
}

func (c *committeeObj) Color() *balance.Color {
	return &c.color
}

func (c *committeeObj) Size() uint16 {
	return c.size
}

func (c *committeeObj) ReceiveMessage(msg interface{}) {
	if c.isOpenQueue.Load() {
		select {
		case c.chMsg <- msg:
		case <-time.After(c.params.ReceiveMsgChannelTimeout):
			c.log.Warnf("timeout on ReceiveMessage type '%T'. Will be repeated", msg)
			go c.ReceiveMessage(msg)
		}
	}
}

// sends message to peer by index. It can be both committee peer or access peer
func (c *committeeObj) SendMsg(targetPeerIndex uint16, msgType byte, msgData []byte) error {
	if int(targetPeerIndex) >= len(c.peers) {
		return fmt.Errorf("SendMsg: wrong peer index")
	}
	peer := c.peers[targetPeerIndex]
	if peer == nil {
		return fmt.Errorf("SendMsg: wrong peer")
	}
	msg := &peering.PeerMessage{
		Address:     c.address,
		SenderIndex: c.ownIndex,
		MsgType:     msgType,
		MsgData:     msgData,
	}
	return peer.SendMsg(msg)
}

func (c *committeeObj) SendMsgToCommitteePeers(msgType byte, msgData []byte) (uint16, int64) {
	msg := &peering.PeerMessage{
		Address:     c.address,
		SenderIndex: c.ownIndex,
		MsgType:     msgType,
		MsgData:     msgData,
	}
	return peering.SendMsgToPeers(msg, c.committeePeers()...)
}

// sends message to the peer seq[seqIndex]. If receives error, seqIndex = (seqIndex+1) % size and repeats
// if is not able to send message after size attempts, returns an error
// seqIndex is start seqIndex
// returned index is seqIndex of the successful send
func (c *committeeObj) SendMsgInSequence(msgType byte, msgData []byte, seqIndex uint16, seq []uint16) (uint16, error) {
	if len(seq) != int(c.Size()) || seqIndex >= c.Size() || !util.ValidPermutation(seq) {
		return 0, fmt.Errorf("SendMsgInSequence: wrong params")
	}
	numAttempts := uint16(0)
	for ; numAttempts < c.Size(); seqIndex = (seqIndex + 1) % c.Size() {
		if seq[seqIndex] >= c.Size() {
			return 0, fmt.Errorf("SendMsgInSequence: wrong params")
		}
		if err := c.SendMsg(seq[seqIndex], msgType, msgData); err == nil {
			return seqIndex, nil
		}
		numAttempts++
	}
	return 0, fmt.Errorf("failed to send")
}

// returns true if peer is alive. Used by the operator to determine current leader
func (c *committeeObj) IsAlivePeer(peerIndex uint16) bool {
	if int(peerIndex) >= len(c.peers) {
		return false
	}
	if peerIndex == c.ownIndex {
		return true
	}
	return c.peers[peerIndex].IsAlive()
}

func (c *committeeObj) OwnPeerIndex() uint16 {
	return c.ownIndex
}

func (c *committeeObj) NumPeers() uint16 {
	return uint16(len(c.peers))
}

// first N peers are committee peers, the rest are access peers in any
func (c *committeeObj) committeePeers() []*peering.Peer {
	return c.peers[:c.size]
}

func (c *committeeObj) HasQuorum() bool {
	count := uint16(0)
	for _, peer := range c.committeePeers() {
		if peer == nil {
			count++
		} else {
			if peer.IsAlive() {
				count++
			}
		}
		if count >= c.quorum {
			return true
		}
	}
	return false
}

func (c *committeeObj) ConnectedPeers() []uint16 {
	ret := make([]uint16, 0)
	for i, peer := range c.committeePeers() {
		if peer == nil {
			ret = append(ret, uint16(i))
		} else {
			if peer.IsAlive() {
				ret = append(ret, uint16(i))
			}
		}
	}
	return ret
}
