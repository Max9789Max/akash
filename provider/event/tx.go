package event

import (
	"context"

	"github.com/ovrclk/akash/marketplace"
	"github.com/ovrclk/akash/types"
	tmtmtypes "github.com/tendermint/tendermint/types"
	"github.com/tendermint/tmlibs/log"
)

type TxCreateOrder struct {
	tx *types.TxCreateOrder
}

func (ev TxCreateOrder) Tx() *types.TxCreateOrder {
	return ev.tx
}

func (ev TxCreateOrder) OrderID() types.OrderID {
	return ev.tx.OrderID
}

type TxCreateFulfillment struct {
	tx *types.TxCreateFulfillment
}

func (ev TxCreateFulfillment) Tx() *types.TxCreateFulfillment {
	return ev.tx
}

func (ev TxCreateFulfillment) FulfillmentID() types.FulfillmentID {
	return ev.tx.FulfillmentID
}

type TxCreateLease struct {
	tx *types.TxCreateLease
}

func (ev TxCreateLease) Tx() *types.TxCreateLease {
	return ev.tx
}

func (ev TxCreateLease) LeaseID() types.LeaseID {
	return ev.tx.LeaseID
}

type TxCloseDeployment struct {
	tx *types.TxCloseDeployment
}

func (ev TxCloseDeployment) Tx() *types.TxCloseDeployment {
	return ev.tx
}

func (ev TxCloseDeployment) DeploymentID() []byte {
	return ev.tx.Deployment
}

type TxCloseFulfillment struct {
	tx *types.TxCloseFulfillment
}

func (ev TxCloseFulfillment) Tx() *types.TxCloseFulfillment {
	return ev.tx
}

func (ev TxCloseFulfillment) FulfillmentID() types.FulfillmentID {
	return ev.tx.FulfillmentID
}

func MarketplaceTxPublisher(ctx context.Context, log log.Logger, tmbus tmtmtypes.EventBusSubscriber, bus Bus) (marketplace.Monitor, error) {
	handler := MarketplaceTxHandler(bus)
	return marketplace.NewMonitor(ctx, log, tmbus, "tx-publisher", handler, marketplace.TxQuery())
}

func MarketplaceTxHandler(bus Bus) marketplace.Handler {
	return marketplace.NewBuilder().
		OnTxCreateOrder(func(tx *types.TxCreateOrder) {
			bus.Publish(TxCreateOrder{tx})
		}).
		OnTxCreateFulfillment(func(tx *types.TxCreateFulfillment) {
			bus.Publish(TxCreateFulfillment{tx})
		}).
		OnTxCreateLease(func(tx *types.TxCreateLease) {
			bus.Publish(TxCreateLease{tx})
		}).
		OnTxCloseDeployment(func(tx *types.TxCloseDeployment) {
			bus.Publish(TxCloseDeployment{tx})
		}).
		OnTxCloseFulfillment(func(tx *types.TxCloseFulfillment) {
			bus.Publish(TxCloseFulfillment{tx})
		}).
		Create()
}
