package deployment_test

import (
	"fmt"
	"testing"

	"github.com/ovrclk/akash/app/deployment"
	"github.com/ovrclk/akash/app/fulfillment"
	"github.com/ovrclk/akash/app/lease"
	"github.com/ovrclk/akash/app/order"
	"github.com/ovrclk/akash/app/provider"
	apptypes "github.com/ovrclk/akash/app/types"
	"github.com/ovrclk/akash/keys"
	"github.com/ovrclk/akash/query"
	pstate "github.com/ovrclk/akash/state"
	"github.com/ovrclk/akash/testutil"
	"github.com/ovrclk/akash/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	tmtypes "github.com/tendermint/abci/types"
)

func TestAcceptQuery(t *testing.T) {
	state := testutil.NewState(t, nil)

	address := testutil.DeploymentAddress(t)

	app, err := deployment.NewApp(state, testutil.Logger())
	require.NoError(t, err)

	{
		path := query.DeploymentPath(address)
		assert.True(t, app.AcceptQuery(tmtypes.RequestQuery{Path: path}))
	}

	{
		path := fmt.Sprintf("%v%x", "/foo/", address)
		assert.False(t, app.AcceptQuery(tmtypes.RequestQuery{Path: path}))
	}
}

func TestCreateTx(t *testing.T) {
	const groupseq = 1
	state := testutil.NewState(t, nil)
	app, err := deployment.NewApp(state, testutil.Logger())
	require.NoError(t, err)
	account, key := testutil.CreateAccount(t, state)
	nonce := uint64(1)

	depl, groups := testutil.CreateDeployment(t, app, account, key, nonce)

	{
		path := query.DeploymentPath(depl.Address)
		resp := app.Query(tmtypes.RequestQuery{Path: path})
		assert.Empty(t, resp.Log)
		require.True(t, resp.IsOK())

		dep := new(types.Deployment)
		require.NoError(t, dep.Unmarshal(resp.Value))

		assert.Equal(t, depl.Tenant, dep.Tenant)
		assert.Equal(t, depl.Address, dep.Address)
	}

	{
		path := query.DeploymentGroupPath(groups.Items[0].DeploymentGroupID)
		resp := app.Query(tmtypes.RequestQuery{Path: path})
		assert.Empty(t, resp.Log)
		require.True(t, resp.IsOK())

		grps := new(types.DeploymentGroup)
		require.NoError(t, grps.Unmarshal(resp.Value))

		assert.Equal(t, grps.Requirements, groups.GetItems()[0].Requirements)
		assert.Equal(t, grps.Resources, groups.GetItems()[0].Resources)
	}

	{
		path := pstate.DeploymentPath
		resp := app.Query(tmtypes.RequestQuery{Path: path})
		assert.Empty(t, resp.Log)
		require.True(t, resp.IsOK())
	}

	badgroup := types.DeploymentGroupID{
		Deployment: depl.Address,
		Seq:        2,
	}

	goodgroup := groups.GetItems()[0].DeploymentGroupID

	{
		path := fmt.Sprintf("%v%v",
			pstate.DeploymentPath,
			keys.DeploymentGroupID(badgroup).Path())
		resp := app.Query(tmtypes.RequestQuery{Path: path})
		assert.NotEmpty(t, resp.Log)
		require.False(t, resp.IsOK())
	}

	{
		path := query.DeploymentGroupPath(goodgroup)
		resp := app.Query(tmtypes.RequestQuery{Path: path})
		assert.Empty(t, resp.Log)
		require.True(t, resp.IsOK())
	}

	{
		path := query.DeploymentGroupPath(badgroup)
		resp := app.Query(tmtypes.RequestQuery{Path: path})
		assert.NotEmpty(t, resp.Log)
		require.False(t, resp.IsOK())
	}

	{
		grps, err := state.DeploymentGroup().ForDeployment(depl.Address)
		require.NoError(t, err)
		require.Len(t, grps, 1)

		assert.Equal(t, grps[0].Requirements, groups.GetItems()[0].Requirements)
		assert.Equal(t, grps[0].Resources, groups.GetItems()[0].Resources)
	}
}

func TestTx_BadTxType(t *testing.T) {
	state_ := testutil.NewState(t, nil)
	app, err := deployment.NewApp(state_, testutil.Logger())
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

func TestCloseTx_1(t *testing.T) {
	state := testutil.NewState(t, nil)
	app, err := deployment.NewApp(state, testutil.Logger())
	require.NoError(t, err)
	account, key := testutil.CreateAccount(t, state)
	nonce := uint64(1)

	depl, groups := testutil.CreateDeployment(t, app, account, key, nonce)

	group := groups.Items[0]

	check := func(
		dstate types.Deployment_DeploymentState,
		gstate types.DeploymentGroup_DeploymentGroupState) {
		assertDeploymentState(t, app, depl.Address, dstate)
		assertDeploymentGroupState(t, app, group.DeploymentGroupID, gstate)
	}

	check(types.Deployment_ACTIVE, types.DeploymentGroup_OPEN)

	testutil.CloseDeployment(t, app, &depl.Address, key)

	check(types.Deployment_CLOSED, types.DeploymentGroup_CLOSED)
}

func TestCloseTx_2(t *testing.T) {

	const (
		oseq = 3
	)

	state := testutil.NewState(t, nil)
	app, err := deployment.NewApp(state, testutil.Logger())
	require.NoError(t, err)
	account, key := testutil.CreateAccount(t, state)
	nonce := uint64(1)

	depl, groups := testutil.CreateDeployment(t, app, account, key, nonce)
	group := groups.Items[0]

	oapp, err := order.NewApp(state, testutil.Logger())
	require.NoError(t, err)

	order := testutil.CreateOrder(t, oapp, account, key, depl.Address, group.Seq, oseq)

	check := func(
		dstate types.Deployment_DeploymentState,
		gstate types.DeploymentGroup_DeploymentGroupState,
		ostate types.Order_OrderState) {
		assertDeploymentState(t, app, depl.Address, dstate)
		assertDeploymentGroupState(t, app, order.GroupID(), gstate)
		assertOrderState(t, oapp, order.OrderID, ostate)
	}

	check(types.Deployment_ACTIVE, types.DeploymentGroup_OPEN, types.Order_OPEN)

	testutil.CloseDeployment(t, app, &depl.Address, key)

	check(types.Deployment_CLOSED, types.DeploymentGroup_CLOSED, types.Order_CLOSED)
}

func TestCloseTx_3(t *testing.T) {

	const (
		oseq  = 3
		price = 0
	)

	state := testutil.NewState(t, nil)
	app, err := deployment.NewApp(state, testutil.Logger())
	require.NoError(t, err)
	account, key := testutil.CreateAccount(t, state)
	nonce := uint64(1)
	depl, groups := testutil.CreateDeployment(t, app, account, key, nonce)
	group := groups.Items[0]

	orderapp, err := order.NewApp(state, testutil.Logger())
	require.NoError(t, err)
	order := testutil.CreateOrder(t, orderapp, account, key, depl.Address, group.Seq, oseq)

	providerapp, err := provider.NewApp(state, testutil.Logger())
	prov := testutil.CreateProvider(t, providerapp, account, key, nonce)

	fulfillmentapp, err := fulfillment.NewApp(state, testutil.Logger())
	fulfillment := testutil.CreateFulfillment(t, fulfillmentapp, prov.Address, key, depl.Address, group.Seq, order.Seq, price)

	check := func(
		dstate types.Deployment_DeploymentState,
		gstate types.DeploymentGroup_DeploymentGroupState,
		ostate types.Order_OrderState,
		fstate types.Fulfillment_FulfillmentState) {
		assertDeploymentState(t, app, depl.Address, dstate)
		assertDeploymentGroupState(t, app, group.DeploymentGroupID, gstate)
		assertOrderState(t, orderapp, order.OrderID, ostate)
		assertFulfillmentState(t, fulfillmentapp, fulfillment.FulfillmentID, fstate)
	}

	check(types.Deployment_ACTIVE, types.DeploymentGroup_OPEN, types.Order_OPEN, types.Fulfillment_OPEN)

	testutil.CloseDeployment(t, app, &depl.Address, key)

	check(types.Deployment_CLOSED, types.DeploymentGroup_CLOSED, types.Order_CLOSED, types.Fulfillment_CLOSED)
}

func TestCloseTx_4(t *testing.T) {

	const (
		gseq  = 1
		oseq  = 3
		price = 0
	)

	state := testutil.NewState(t, nil)
	app, err := deployment.NewApp(state, testutil.Logger())
	require.NoError(t, err)
	account, key := testutil.CreateAccount(t, state)
	nonce := uint64(1)
	depl, _ := testutil.CreateDeployment(t, app, account, key, nonce)

	orderapp, err := order.NewApp(state, testutil.Logger())
	require.NoError(t, err)
	order := testutil.CreateOrder(t, orderapp, account, key, depl.Address, gseq, oseq)

	providerapp, err := provider.NewApp(state, testutil.Logger())
	prov := testutil.CreateProvider(t, providerapp, account, key, nonce)

	fulfillmentapp, err := fulfillment.NewApp(state, testutil.Logger())
	fulfillment := testutil.CreateFulfillment(t, fulfillmentapp, prov.Address, key, depl.Address, gseq, oseq, price)

	leaseapp, err := lease.NewApp(state, testutil.Logger())
	lease := testutil.CreateLease(t, leaseapp, prov.Address, key, depl.Address, gseq, oseq, price)

	check := func(
		dstate types.Deployment_DeploymentState,
		gstate types.DeploymentGroup_DeploymentGroupState,
		ostate types.Order_OrderState,
		fstate types.Fulfillment_FulfillmentState,
		lstate types.Lease_LeaseState) {
		assertDeploymentState(t, app, depl.Address, dstate)
		assertDeploymentGroupState(t, app, order.GroupID(), gstate)
		assertOrderState(t, orderapp, order.OrderID, ostate)
		assertFulfillmentState(t, fulfillmentapp, fulfillment.FulfillmentID, fstate)
		assertLeaseState(t, leaseapp, lease.LeaseID, lstate)
	}

	check(types.Deployment_ACTIVE, types.DeploymentGroup_OPEN, types.Order_MATCHED, types.Fulfillment_OPEN, types.Lease_ACTIVE)

	testutil.CloseDeployment(t, app, &depl.Address, key)

	check(types.Deployment_CLOSED, types.DeploymentGroup_CLOSED, types.Order_CLOSED, types.Fulfillment_CLOSED, types.Lease_CLOSED)
}

// check deployment and group query & status
func assertDeploymentState(
	t *testing.T,
	app apptypes.Application,
	daddr []byte,
	dstate types.Deployment_DeploymentState) {

	path := query.DeploymentPath(daddr)
	resp := app.Query(tmtypes.RequestQuery{Path: path})
	assert.Empty(t, resp.Log)
	require.True(t, resp.IsOK())

	dep := new(types.Deployment)
	require.NoError(t, dep.Unmarshal(resp.Value))

	assert.Equal(t, dstate, dep.State)
}

// check deployment and group query & status
func assertDeploymentGroupState(
	t *testing.T,
	app apptypes.Application,
	id types.DeploymentGroupID,
	gstate types.DeploymentGroup_DeploymentGroupState) {

	path := query.DeploymentGroupPath(id)
	resp := app.Query(tmtypes.RequestQuery{Path: path})
	assert.Empty(t, resp.Log)
	require.True(t, resp.IsOK())

	group := new(types.DeploymentGroup)
	require.NoError(t, group.Unmarshal(resp.Value))

	assert.Equal(t, gstate, group.State)
}

func assertOrderState(
	t *testing.T,
	app apptypes.Application,
	id types.OrderID,
	ostate types.Order_OrderState) {

	path := query.OrderPath(id)
	resp := app.Query(tmtypes.RequestQuery{Path: path})
	assert.Empty(t, resp.Log)
	require.True(t, resp.IsOK())

	order := new(types.Order)
	require.NoError(t, order.Unmarshal(resp.Value))
	assert.Equal(t, ostate, order.State)
}

func assertFulfillmentState(
	t *testing.T,
	app apptypes.Application,
	id types.FulfillmentID,
	state types.Fulfillment_FulfillmentState) {

	path := query.FulfillmentPath(id)
	resp := app.Query(tmtypes.RequestQuery{Path: path})
	assert.Empty(t, resp.Log)
	require.True(t, resp.IsOK())

	obj := new(types.Fulfillment)
	require.NoError(t, obj.Unmarshal(resp.Value))
	assert.Equal(t, state, obj.State)
}

func assertLeaseState(
	t *testing.T,
	app apptypes.Application,
	id types.LeaseID,
	state types.Lease_LeaseState) {

	// check fulfillment state
	path := query.LeasePath(id)
	resp := app.Query(tmtypes.RequestQuery{Path: path})
	assert.Empty(t, resp.Log)
	require.True(t, resp.IsOK())

	obj := new(types.Lease)
	require.NoError(t, obj.Unmarshal(resp.Value))
	assert.Equal(t, state, obj.State)
}
