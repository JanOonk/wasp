package jsonable

import (
	"encoding/json"
	"github.com/iotaledger/wasp/packages/coretypes"

	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address"
)

type Address struct {
	address address.Address
}

type ChainID struct {
	chainID coretypes.ChainID
}

func NewAddress(address *address.Address) *Address {
	return &Address{address: *address}
}

func (a Address) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.Address().String())
}

func (a *Address) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	addr, err := address.FromBase58(s)
	a.address = addr
	return err
}

func (a Address) Address() *address.Address {
	return &a.address
}

func NewChainID(chainID *coretypes.ChainID) *ChainID {
	return &ChainID{chainID: *chainID}
}

func (a ChainID) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.ChainID().String())
}

func (a *ChainID) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	chainID, err := coretypes.NewChainIDFromBase58(s)
	a.chainID = chainID
	return err
}

func (a ChainID) ChainID() *coretypes.ChainID {
	return &a.chainID
}