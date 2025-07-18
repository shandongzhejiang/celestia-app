package app

import (
	"github.com/celestiaorg/celestia-app/v4/pkg/appconsts"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MaxEffectiveSquareSize returns the max effective square size.
func (app *App) MaxEffectiveSquareSize(ctx sdk.Context) int {
	govMax := app.BlobKeeper.GetParams(ctx).GovMaxSquareSize
	hardMax := appconsts.SquareSizeUpperBound
	return min(int(govMax), hardMax)
}
