package alliance

import (
	"github.com/terra-money/alliance/x/alliance/keeper"

	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// EndBlocker
func EndBlocker(ctx sdk.Context, k keeper.Keeper) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}
