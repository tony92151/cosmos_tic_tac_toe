package keeper

import (
	"github.com/cosmos_tic_tac_toe/cttt/x/cttt/types"
)

var _ types.QueryServer = Keeper{}
