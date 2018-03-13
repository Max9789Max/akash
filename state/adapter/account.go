package adapter

import (
	"github.com/gogo/protobuf/proto"
	"github.com/ovrclk/photon/state"
	"github.com/ovrclk/photon/types"
)

type account struct {
	db state.DB
}

func (a account) Get(key AccountKey) (*types.Account, error) {
}

func (a account) List(key AccountKey) (*types.Accounts, error) {
}

func (a account) Save(key AccountKey, account *types.Account) {
}

func (a account) read(key Key, obj *proto.Message) error {
	buf, err := a.db.Get(a.pathFor(key))
	if err != nil {
		*obj = nil
		return err
	}
}
