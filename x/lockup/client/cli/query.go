package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/osmosis-labs/osmosis/v13/osmoutils/osmocli"
	"github.com/osmosis-labs/osmosis/v13/x/lockup/types"
)

// GetQueryCmd returns the cli query commands for this module.
func GetQueryCmd() *cobra.Command {
	cmd := osmocli.QueryIndexCmd(types.ModuleName)

	cmd.AddCommand(
		GetCmdModuleBalance(),
		GetCmdModuleLockedAmount(),
		GetCmdAccountUnlockableCoins(),
		GetCmdAccountUnlockingCoins(),
		GetCmdAccountLockedCoins(),
		GetCmdAccountLockedPastTime(),
		GetCmdAccountLockedPastTimeNotUnlockingOnly(),
		GetCmdAccountUnlockedBeforeTime(),
		GetCmdAccountLockedPastTimeDenom(),
		GetCmdLockedByID(),
		GetCmdAccountLockedLongerDuration(),
		GetCmdAccountLockedLongerDurationNotUnlockingOnly(),
		GetCmdAccountLockedLongerDurationDenom(),
		GetCmdTotalLockedByDenom(),
		GetCmdOutputLocksJson(),
		GetCmdSyntheticLockupsByLockupID(),
		GetCmdAccountLockedDuration(),
		osmocli.GetParams[*types.QueryParamsRequest](
			types.ModuleName, types.NewQueryClient),
	)

	return cmd
}

// GetCmdModuleBalance returns full balance of the lockup module.
// Lockup module is where coins of locks are held.
// This includes locked balance and unlocked balance of the module.
func GetCmdModuleBalance() *cobra.Command {
	return osmocli.SimpleQueryCmd[*types.ModuleBalanceRequest](
		"module-balance",
		"Query module balance",
		`{{.Short}}`, types.ModuleName, types.NewQueryClient)
}

// GetCmdModuleLockedAmount returns locked balance of the module,
// which are all the tokens not unlocking + tokens that are not finished unlocking.
func GetCmdModuleLockedAmount() *cobra.Command {
	return osmocli.SimpleQueryCmd[*types.ModuleLockedAmountRequest](
		"module-locked-amount",
		"Query locked amount",
		`{{.Short}}`, types.ModuleName, types.NewQueryClient)
}

// GetCmdAccountUnlockableCoins returns unlockable coins which has finsihed unlocking.
// TODO: DELETE THIS + Actual query in subsequent PR
func GetCmdAccountUnlockableCoins() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "account-unlockable-coins <address>",
		Short: "Query account's unlockable coins",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query account's unlockable coins.

Example:
$ %s query lockup account-unlockable-coins <address>
`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.AccountUnlockableCoins(cmd.Context(), &types.AccountUnlockableCoinsRequest{Owner: args[0]})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdAccountUnlockingCoins returns unlocking coins of a specific account.
func GetCmdAccountUnlockingCoins() *cobra.Command {
	return osmocli.SimpleQueryCmd[*types.AccountUnlockingCoinsRequest](
		"account-unlocking-coins <address>",
		"Query account's unlocking coins",
		`{{.Short}}{{.ExampleHeader}}
{{.CommandPrefix}} account-unlocking-coins <address>
`, types.ModuleName, types.NewQueryClient)
}

// GetCmdAccountLockedCoins returns locked coins that that are still in a locked state from the specified account.
func GetCmdAccountLockedCoins() *cobra.Command {
	return osmocli.SimpleQueryCmd[*types.AccountLockedCoinsRequest](
		"account-locked-coins <address>",
		"Query account's locked coins",
		`{{.Short}}{{.ExampleHeader}}
{{.CommandPrefix}} account-locked-coins <address>
`, types.ModuleName, types.NewQueryClient)
}

// GetCmdAccountLockedPastTime returns locks of an account with unlock time beyond timestamp.
func GetCmdAccountLockedPastTime() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "account-locked-pastime <address> <timestamp>",
		Short: "Query locked records of an account with unlock time beyond timestamp",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query locked records of an account with unlock time beyond timestamp.

Example:
$ %s query lockup account-locked-pastime <address> <timestamp>
`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			i, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				panic(err)
			}
			timestamp := time.Unix(i, 0)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.AccountLockedPastTime(cmd.Context(), &types.AccountLockedPastTimeRequest{Owner: args[0], Timestamp: timestamp})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdAccountLockedPastTimeNotUnlockingOnly returns locks of an account with unlock time beyond provided timestamp
// amongst the locks that are in the unlocking queue.
func GetCmdAccountLockedPastTimeNotUnlockingOnly() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "account-locked-pastime-not-unlocking <address> <timestamp>",
		Short: "Query locked records of an account with unlock time beyond timestamp within not unlocking queue",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query locked records of an account with unlock time beyond timestamp within not unlocking queue.

Example:
$ %s query lockup account-locked-pastime-not-unlocking <address> <timestamp>
`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			i, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				panic(err)
			}
			timestamp := time.Unix(i, 0)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.AccountLockedPastTimeNotUnlockingOnly(cmd.Context(), &types.AccountLockedPastTimeNotUnlockingOnlyRequest{Owner: args[0], Timestamp: timestamp})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdAccountUnlockedBeforeTime returns locks with unlock time before the provided timestamp.
func GetCmdAccountUnlockedBeforeTime() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "account-locked-beforetime <address> <timestamp>",
		Short: "Query account's unlocked records before specific time",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query account's the total unlocked records with unlock time before timestamp.

Example:
$ %s query lockup account-locked-pastime <address> <timestamp>
`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			i, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				panic(err)
			}
			timestamp := time.Unix(i, 0)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.AccountUnlockedBeforeTime(cmd.Context(), &types.AccountUnlockedBeforeTimeRequest{Owner: args[0], Timestamp: timestamp})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdAccountLockedPastTimeDenom returns locks of an account whose unlock time is
// beyond given timestamp, and locks with the specified denom.
func GetCmdAccountLockedPastTimeDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "account-locked-pastime-denom <address> <timestamp> <denom>",
		Short: "Query account's lock records by address, timestamp, denom",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query account's lock records by address, timestamp, denom.

Example:
$ %s query lockup account-locked-pastime-denom <address> <timestamp> <denom>
`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			i, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				panic(err)
			}
			timestamp := time.Unix(i, 0)

			denom := args[2]

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.AccountLockedPastTimeDenom(cmd.Context(), &types.AccountLockedPastTimeDenomRequest{Owner: args[0], Timestamp: timestamp, Denom: denom})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdLockedByID returns lock by id.
func GetCmdLockedByID() *cobra.Command {
	q := osmocli.QueryDescriptor{
		Use:   "lock-by-id <id>",
		Short: "Query account's lock record by id",
		Long: `{{.Short}}{{.ExampleHeader}}
{{.CommandPrefix}} lock-by-id 1`,
		QueryFnName: "LockedByID",
	}
	q.Long = osmocli.FormatLongDesc(q.Long, osmocli.NewLongMetadata(types.ModuleName).WithShort(q.Short))
	return osmocli.SimpleQueryFromDescriptor[*types.LockedRequest](q, types.NewQueryClient)
}

// GetCmdSyntheticLockupsByLockupID returns synthetic lockups by lockup id.
func GetCmdSyntheticLockupsByLockupID() *cobra.Command {
	return osmocli.SimpleQueryCmd[*types.SyntheticLockupsByLockupIDRequest](
		"synthetic-lockups-by-lock-id <id>",
		"Query synthetic lockups by lockup id",
		`{{.Short}}`, types.ModuleName, types.NewQueryClient)
}

// GetCmdAccountLockedLongerDuration returns account locked records with longer duration.
func GetCmdAccountLockedLongerDuration() *cobra.Command {
	return osmocli.SimpleQueryCmd[*types.AccountLockedLongerDurationRequest](
		"account-locked-longer-duration <address> <duration>",
		"Query account locked records with longer duration",
		`{{.Short}}`, types.ModuleName, types.NewQueryClient)
}

// GetCmdAccountLockedLongerDuration returns account locked records with longer duration.
func GetCmdAccountLockedDuration() *cobra.Command {
	return osmocli.SimpleQueryCmd[*types.AccountLockedDurationRequest](
		"account-locked-duration <address> <duration>",
		"Query account locked records with a specific duration",
		`{{.Short}}{{.ExampleHeader}}
{{.CommandPrefix}} account-locked-duration osmo1yl6hdjhmkf37639730gffanpzndzdpmhxy9ep3 604800s`, types.ModuleName, types.NewQueryClient)
}

// GetCmdAccountLockedLongerDurationNotUnlockingOnly returns account locked records with longer duration from unlocking only queue.
func GetCmdAccountLockedLongerDurationNotUnlockingOnly() *cobra.Command {
	return osmocli.SimpleQueryCmd[*types.AccountLockedLongerDurationNotUnlockingOnlyRequest](
		"account-locked-longer-duration-not-unlocking <address> <duration>",
		"Query account locked records with longer duration from unlocking only queue",
		`{{.Short}}`, types.ModuleName, types.NewQueryClient)
}

// GetCmdAccountLockedLongerDurationDenom returns account's locks for a specific denom
// with longer duration than the given duration.
func GetCmdAccountLockedLongerDurationDenom() *cobra.Command {
	return osmocli.SimpleQueryCmd[*types.AccountLockedLongerDurationDenomRequest](
		"account-locked-longer-duration-denom <address> <duration> <denom>",
		"Query locked records for a denom with longer duration",
		`{{.Short}}`, types.ModuleName, types.NewQueryClient)
}

func GetCmdTotalLockedByDenom() *cobra.Command {
	cmd := osmocli.SimpleQueryFromDescriptor[*types.LockedDenomRequest](osmocli.QueryDescriptor{
		Use:   "total-locked-of-denom <denom>",
		Short: "Query locked amount for a specific denom bigger then duration provided",
		Long: osmocli.FormatLongDescDirect(`{{.Short}}{{.ExampleHeader}}
{{.CommandPrefix}} total-locked-of-denom uosmo --min-duration=0s`, types.ModuleName),
		CustomFlagOverrides: map[string]string{
			"duration": FlagMinDuration,
		},
		QueryFnName: "LockedDenom",
	}, types.NewQueryClient)

	cmd.Flags().AddFlagSet(FlagSetMinDuration())
	return cmd
}

// GetCmdOutputLocksJson outputs all locks into a file called lock_export.json.
func GetCmdOutputLocksJson() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "output-all-locks <max lock ID>",
		Short: "output all locks into a json file",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Output all locks into a json file.
Example:
$ %s query lockup output-all-locks <max lock ID>
`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			maxLockID, err := strconv.ParseInt(args[0], 10, 32)
			if err != nil {
				return err
			}

			// status
			const (
				doesnt_exist_status = iota
				unbonding_status
				bonded_status
			)

			type LockResult struct {
				Id            int
				Status        int // one of {doesnt_exist, }
				Denom         string
				Amount        sdk.Int
				Address       string
				UnbondEndTime time.Time
			}
			queryClient := types.NewQueryClient(clientCtx)

			results := []LockResult{}
			for i := 0; i <= int(maxLockID); i++ {
				curLockResult := LockResult{Id: i}
				res, err := queryClient.LockedByID(cmd.Context(), &types.LockedRequest{LockId: uint64(i)})
				if err != nil {
					curLockResult.Status = doesnt_exist_status
					results = append(results, curLockResult)
					continue
				}
				// 1527019420 is hardcoded time well before launch, but well after year 1
				if res.Lock.EndTime.Before(time.Unix(1527019420, 0)) {
					curLockResult.Status = bonded_status
				} else {
					curLockResult.Status = unbonding_status
					curLockResult.UnbondEndTime = res.Lock.EndTime
					curLockResult.Denom = res.Lock.Coins[0].Denom
					curLockResult.Amount = res.Lock.Coins[0].Amount
					curLockResult.Address = res.Lock.Owner
				}
				results = append(results, curLockResult)
			}

			bz, err := json.Marshal(results)
			if err != nil {
				return err
			}
			err = os.WriteFile("lock_export.json", bz, 0o777)
			if err != nil {
				return err
			}

			fmt.Println("Writing to lock_export.json")
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
