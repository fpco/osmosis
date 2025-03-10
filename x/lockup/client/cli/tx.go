package cli

import (
	"github.com/spf13/cobra"

	"github.com/osmosis-labs/osmosis/v13/osmoutils/osmocli"
	"github.com/osmosis-labs/osmosis/v13/x/lockup/types"
)

// GetTxCmd returns the transaction commands for this module.
func GetTxCmd() *cobra.Command {
	cmd := osmocli.TxIndexCmd(types.ModuleName)
	cmd.AddCommand(
		NewLockTokensCmd(),
		NewBeginUnlockingAllCmd(),
		NewBeginUnlockByIDCmd(),
		NewForceUnlockByIdCmd(),
	)

	return cmd
}

func NewLockTokensCmd() *cobra.Command {
	cmd := osmocli.BuildTxCli[*types.MsgLockTokens](&osmocli.TxCliDesc{
		Use:   "lock-tokens [tokens]",
		Short: "lock tokens into lockup pool from user account",
		CustomFlagOverrides: map[string]string{
			"duration": FlagDuration,
		},
	})

	cmd.Flags().AddFlagSet(FlagSetLockTokens())
	err := cmd.MarkFlagRequired(FlagDuration)
	if err != nil {
		panic(err)
	}
	return cmd
}

// TODO: We should change the Use string to be unlock-all
func NewBeginUnlockingAllCmd() *cobra.Command {
	return osmocli.BuildTxCli[*types.MsgBeginUnlockingAll](&osmocli.TxCliDesc{
		Use:   "begin-unlock-tokens",
		Short: "begin unlock not unlocking tokens from lockup pool for sender",
	})
}

// NewBeginUnlockByIDCmd unlocks individual period lock by ID.
func NewBeginUnlockByIDCmd() *cobra.Command {
	cmd := osmocli.BuildTxCli[*types.MsgBeginUnlocking](&osmocli.TxCliDesc{
		Use:   "begin-unlock-by-id [id]",
		Short: "begin unlock individual period lock by ID",
		CustomFlagOverrides: map[string]string{
			"coins": FlagAmount,
		},
	})

	cmd.Flags().AddFlagSet(FlagSetUnlockTokens())
	return cmd
}

// NewForceUnlockByIdCmd force unlocks individual period lock by ID if proper permissions exist.
func NewForceUnlockByIdCmd() *cobra.Command {
	cmd := osmocli.BuildTxCli[*types.MsgBeginUnlocking](&osmocli.TxCliDesc{
		Use:   "force-unlock-by-id [id]",
		Short: "force unlocks individual period lock by ID",
		Long:  "force unlocks individual period lock by ID. if no amount provided, entire lock is unlocked",
		CustomFlagOverrides: map[string]string{
			"coins": FlagAmount,
		},
	})

	cmd.Flags().AddFlagSet(FlagSetUnlockTokens())
	return cmd
}
