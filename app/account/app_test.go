package account_test

import (
	"testing"

	"github.com/ovrclk/akash/app/account"
	apptypes "github.com/ovrclk/akash/app/types"
	"github.com/ovrclk/akash/query"
	"github.com/ovrclk/akash/testutil"
	"github.com/ovrclk/akash/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	tmtypes "github.com/tendermint/abci/types"
)

func TestAccountApp(t *testing.T) {

	const (
		balance uint64 = 150
		amount  uint64 = 100
	)

	keyfrom := testutil.PrivateKey(t)
	addrfrom := keyfrom.PubKey().Address().Bytes()
	keyto := testutil.PrivateKey(t)
	addrto := keyto.PubKey().Address().Bytes()

	send := &types.TxPayload_TxSend{
		TxSend: &types.TxSend{
			From:   addrfrom,
			To:     addrto,
			Amount: amount,
		},
	}

	state := testutil.NewState(t, &types.Genesis{
		Accounts: []types.Account{
			types.Account{Address: addrfrom, Balance: balance},
		},
	})

	ctx := apptypes.NewContext(&types.Tx{
		Key: keyfrom.PubKey().Bytes(),
		Payload: types.TxPayload{
			Payload: send,
		},
	})

	app, err := account.NewApp(state, testutil.Logger())
	require.NoError(t, err)

	assert.True(t, app.AcceptQuery(tmtypes.RequestQuery{Path: query.AccountPath(addrfrom)}))

	assert.True(t, app.AcceptTx(ctx, send))

	{
		resp := app.CheckTx(ctx, send)
		assert.True(t, resp.IsOK(), resp.Log)
	}

	{
		resp := app.DeliverTx(ctx, send)
		assert.True(t, resp.IsOK(), resp.Log)
	}

	{
		resp := app.Query(tmtypes.RequestQuery{Path: query.AccountPath(addrfrom)})
		assert.Empty(t, resp.Log)
		require.True(t, resp.IsOK())

		addr := new(types.Account)
		require.NoError(t, addr.Unmarshal(resp.Value))

		assert.Equal(t, send.TxSend.From, addr.Address)
		assert.Equal(t, balance-amount, addr.Balance)
	}

	{
		resp := app.Query(tmtypes.RequestQuery{Path: query.AccountPath(addrto)})
		assert.Empty(t, resp.Log)
		require.True(t, resp.IsOK())

		addr := new(types.Account)
		require.NoError(t, addr.Unmarshal(resp.Value))

		assert.Equal(t, send.TxSend.To, addr.Address)
		assert.Equal(t, amount, addr.Balance)
	}

}

func TestTx_BadTxType(t *testing.T) {
	state_ := testutil.NewState(t, nil)
	app, err := account.NewApp(state_, testutil.Logger())
	require.NoError(t, err)
	account, key := testutil.CreateAccount(t, state_)
	tx := testutil.ProviderTx(account, key, 10)
	ctx := apptypes.NewContext(tx)
	assert.False(t, app.AcceptTx(ctx, tx.Payload.Payload))
	cresp := app.CheckTx(ctx, tx.Payload.Payload)
	assert.False(t, cresp.IsOK())
	dresp := app.DeliverTx(ctx, tx.Payload.Payload)
	assert.False(t, dresp.IsOK())
}
