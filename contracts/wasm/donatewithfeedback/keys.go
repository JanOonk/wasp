// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package donatewithfeedback

import "github.com/iotaledger/wasp/packages/vm/wasmlib"

const (
	IdxParamAmount         = 0
	IdxParamFeedback       = 1
	IdxParamNr             = 2
	IdxResultAmount        = 3
	IdxResultCount         = 4
	IdxResultDonator       = 5
	IdxResultError         = 6
	IdxResultFeedback      = 7
	IdxResultMaxDonation   = 8
	IdxResultTimestamp     = 9
	IdxResultTotalDonation = 10
	IdxStateLog            = 11
	IdxStateMaxDonation    = 12
	IdxStateTotalDonation  = 13
)

const keyMapLen = 14

var keyMap = [keyMapLen]wasmlib.Key{
	ParamAmount,
	ParamFeedback,
	ParamNr,
	ResultAmount,
	ResultCount,
	ResultDonator,
	ResultError,
	ResultFeedback,
	ResultMaxDonation,
	ResultTimestamp,
	ResultTotalDonation,
	StateLog,
	StateMaxDonation,
	StateTotalDonation,
}

var idxMap [keyMapLen]wasmlib.Key32