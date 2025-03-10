package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"

	incentivestypes "github.com/osmosis-labs/osmosis/v13/x/incentives/types"
	lockuptypes "github.com/osmosis-labs/osmosis/v13/x/lockup/types"
)

// AccountKeeper interface contains functions for getting accounts and the module address
type AccountKeeper interface {
	GetModuleAddress(name string) sdk.AccAddress
	GetModuleAccount(ctx sdk.Context, name string) authtypes.ModuleAccountI
}

// BankKeeper sends tokens across modules and is able to get account balances.
type BankKeeper interface {
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
}

// GAMMKeeper gets the pool interface from poolID.
type GAMMKeeper interface {
	GetNextPoolId(ctx sdk.Context) uint64
}

// IncentivesKeeper creates and gets gauges, and also allows additions to gauge rewards.
type IncentivesKeeper interface {
	CreateGauge(ctx sdk.Context, isPerpetual bool, owner sdk.AccAddress, coins sdk.Coins, distrTo lockuptypes.QueryCondition, startTime time.Time, numEpochsPaidOver uint64) (uint64, error)
	GetGaugeByID(ctx sdk.Context, gaugeID uint64) (*incentivestypes.Gauge, error)
	GetGauges(ctx sdk.Context) []incentivestypes.Gauge

	AddToGaugeRewards(ctx sdk.Context, owner sdk.AccAddress, coins sdk.Coins, gaugeID uint64) error
}

// DistrKeeper handles pool-fees functionality - setting / getting fees and funding the community pool.
type DistrKeeper interface {
	SetFeePool(ctx sdk.Context, feePool distrtypes.FeePool)
	FundCommunityPool(ctx sdk.Context, amount sdk.Coins, sender sdk.AccAddress) error
}
