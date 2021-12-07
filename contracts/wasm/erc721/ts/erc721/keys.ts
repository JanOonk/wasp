// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "wasmlib";
import * as sc from "./index";

export const IdxParamApproval = 0;
export const IdxParamApproved = 1;
export const IdxParamData     = 2;
export const IdxParamFrom     = 3;
export const IdxParamName     = 4;
export const IdxParamOperator = 5;
export const IdxParamOwner    = 6;
export const IdxParamSymbol   = 7;
export const IdxParamTo       = 8;
export const IdxParamTokenID  = 9;

export const IdxResultAmount   = 10;
export const IdxResultApproval = 11;
export const IdxResultApproved = 12;
export const IdxResultName     = 13;
export const IdxResultOwner    = 14;
export const IdxResultSymbol   = 15;
export const IdxResultTokenURI = 16;

export const IdxStateApprovedAccounts  = 17;
export const IdxStateApprovedOperators = 18;
export const IdxStateBalances          = 19;
export const IdxStateName              = 20;
export const IdxStateOwners            = 21;
export const IdxStateSymbol            = 22;

export let keyMap: string[] = [
	sc.ParamApproval,
	sc.ParamApproved,
	sc.ParamData,
	sc.ParamFrom,
	sc.ParamName,
	sc.ParamOperator,
	sc.ParamOwner,
	sc.ParamSymbol,
	sc.ParamTo,
	sc.ParamTokenID,
	sc.ResultAmount,
	sc.ResultApproval,
	sc.ResultApproved,
	sc.ResultName,
	sc.ResultOwner,
	sc.ResultSymbol,
	sc.ResultTokenURI,
	sc.StateApprovedAccounts,
	sc.StateApprovedOperators,
	sc.StateBalances,
	sc.StateName,
	sc.StateOwners,
	sc.StateSymbol,
];

export let idxMap: wasmlib.Key32[] = new Array(keyMap.length);