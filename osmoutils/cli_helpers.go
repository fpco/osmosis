package osmoutils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func DefaultFeeString(cfg network.Config) string {
	feeCoins := sdk.NewCoins(sdk.NewCoin(cfg.BondDenom, sdk.NewInt(10)))
	return fmt.Sprintf("--%s=%s", flags.FlagFees, feeCoins.String())
}

const (
	base   = 10
	bitlen = 64
)

func ParseUint64SliceFromString(s string, separator string) ([]uint64, error) {
	var parsedInts []uint64
	for _, s := range strings.Split(s, separator) {
		s = strings.TrimSpace(s)

		parsed, err := strconv.ParseUint(s, base, bitlen)
		if err != nil {
			return []uint64{}, err
		}
		parsedInts = append(parsedInts, parsed)
	}
	return parsedInts, nil
}

func ParseSdkIntFromString(s string, separator string) ([]sdk.Int, error) {
	var parsedInts []sdk.Int
	for _, weightStr := range strings.Split(s, separator) {
		weightStr = strings.TrimSpace(weightStr)

		parsed, err := strconv.ParseUint(weightStr, base, bitlen)
		if err != nil {
			return parsedInts, err
		}
		parsedInts = append(parsedInts, sdk.NewIntFromUint64(parsed))
	}
	return parsedInts, nil
}

func ParseSdkDecFromString(s string, separator string) ([]sdk.Dec, error) {
	var parsedDec []sdk.Dec
	for _, weightStr := range strings.Split(s, separator) {
		weightStr = strings.TrimSpace(weightStr)

		parsed, err := sdk.NewDecFromStr(weightStr)
		if err != nil {
			return parsedDec, err
		}

		parsedDec = append(parsedDec, parsed)
	}
	return parsedDec, nil
}

func ParseSdkValAddressFromString(s string, separator string) []sdk.ValAddress {
	var parsedAddr []sdk.ValAddress
	for _, addr := range strings.Split(s, separator) {
		valAddr := sdk.ValAddress([]byte(addr))
		parsedAddr = append(parsedAddr, valAddr)
	}

	return parsedAddr
}
