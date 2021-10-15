// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package fairroulette

import "github.com/iotaledger/wasplib/packages/vm/wasmlib"

type PayWinnersCall struct {
	Func *wasmlib.ScFunc
}

type PlaceBetCall struct {
	Func   *wasmlib.ScFunc
	Params MutablePlaceBetParams
}

type PlayPeriodCall struct {
	Func   *wasmlib.ScFunc
	Params MutablePlayPeriodParams
}

type LastWinningNumberCall struct {
	Func    *wasmlib.ScView
	Results ImmutableLastWinningNumberResults
}

type RoundNumberCall struct {
	Func    *wasmlib.ScView
	Results ImmutableRoundNumberResults
}

type RoundStartedAtCall struct {
	Func    *wasmlib.ScView
	Results ImmutableRoundStartedAtResults
}

type RoundStatusCall struct {
	Func    *wasmlib.ScView
	Results ImmutableRoundStatusResults
}

type Funcs struct{}

var ScFuncs Funcs

func (sc Funcs) PayWinners(ctx wasmlib.ScFuncCallContext) *PayWinnersCall {
	return &PayWinnersCall{Func: wasmlib.NewScFunc(HScName, HFuncPayWinners)}
}

func (sc Funcs) PlaceBet(ctx wasmlib.ScFuncCallContext) *PlaceBetCall {
	f := &PlaceBetCall{Func: wasmlib.NewScFunc(HScName, HFuncPlaceBet)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

func (sc Funcs) PlayPeriod(ctx wasmlib.ScFuncCallContext) *PlayPeriodCall {
	f := &PlayPeriodCall{Func: wasmlib.NewScFunc(HScName, HFuncPlayPeriod)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

func (sc Funcs) LastWinningNumber(ctx wasmlib.ScViewCallContext) *LastWinningNumberCall {
	f := &LastWinningNumberCall{Func: wasmlib.NewScView(HScName, HViewLastWinningNumber)}
	f.Func.SetPtrs(nil, &f.Results.id)
	return f
}

func (sc Funcs) RoundNumber(ctx wasmlib.ScViewCallContext) *RoundNumberCall {
	f := &RoundNumberCall{Func: wasmlib.NewScView(HScName, HViewRoundNumber)}
	f.Func.SetPtrs(nil, &f.Results.id)
	return f
}

func (sc Funcs) RoundStartedAt(ctx wasmlib.ScViewCallContext) *RoundStartedAtCall {
	f := &RoundStartedAtCall{Func: wasmlib.NewScView(HScName, HViewRoundStartedAt)}
	f.Func.SetPtrs(nil, &f.Results.id)
	return f
}

func (sc Funcs) RoundStatus(ctx wasmlib.ScViewCallContext) *RoundStatusCall {
	f := &RoundStatusCall{Func: wasmlib.NewScView(HScName, HViewRoundStatus)}
	f.Func.SetPtrs(nil, &f.Results.id)
	return f
}
