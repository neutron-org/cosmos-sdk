package types

import (
	"sort"

	"cosmossdk.io/math"

	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/gogoproto/proto"

	"cosmossdk.io/errors"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

// NewHistoricalInfo will create a historical information struct from header and valset
// it will first sort valset before inclusion into historical info
func NewHistoricalInfo(header cmtproto.Header, valSet Validators, powerReduction math.Int) HistoricalInfo {
	// Must sort in the same way that CometBFT does
	sort.SliceStable(valSet, func(i, j int) bool {
		return ValidatorsByVotingPower(valSet).Less(i, j, powerReduction)
	})

	return HistoricalInfo{
		Header: header,
		Valset: valSet,
	}
}

// MustUnmarshalHistoricalInfo wll unmarshal historical info and panic on error
func MustUnmarshalHistoricalInfo(cdc codec.BinaryCodec, value []byte) HistoricalInfo {
	hi, err := UnmarshalHistoricalInfo(cdc, value)
	if err != nil {
		panic(err)
	}

	return hi
}

// UnmarshalHistoricalInfo will unmarshal historical info and return any error
func UnmarshalHistoricalInfo(cdc codec.BinaryCodec, value []byte) (hi HistoricalInfo, err error) {
	err = cdc.Unmarshal(value, &hi)
	return hi, err
}

// ValidateBasic will ensure HistoricalInfo is not nil and sorted
func ValidateBasic(hi HistoricalInfo) error {
	if len(hi.Valset) == 0 {
		return errors.Wrap(ErrInvalidHistoricalInfo, "validator set is empty")
	}

	if !sort.IsSorted(Validators(hi.Valset)) {
		return errors.Wrap(ErrInvalidHistoricalInfo, "validator set is not sorted by address")
	}

	return nil
}

// Equal checks if receiver is equal to the parameter
func (hi *HistoricalInfo) Equal(hi2 *HistoricalInfo) bool {
	if !proto.Equal(&hi.Header, &hi2.Header) {
		return false
	}
	if len(hi.Valset) != len(hi2.Valset) {
		return false
	}
	for i := range hi.Valset {
		if !hi.Valset[i].Equal(&hi2.Valset[i]) {
			return false
		}
	}
	return true
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (hi HistoricalInfo) UnpackInterfaces(c codectypes.AnyUnpacker) error {
	for i := range hi.Valset {
		if err := hi.Valset[i].UnpackInterfaces(c); err != nil {
			return err
		}
	}
	return nil
}

// func MustUnmarshalConsPubKeyRotationHistory(cdc codec.BinaryCodec, value []byte) ConsPubKeyRotationHistory {
// 	hi, err := UnmarshalConsPubKeyRotationHistory(cdc, value)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return hi
// }

// func UnmarshalConsPubKeyRotationHistory(cdc codec.BinaryCodec, value []byte) (hi ConsPubKeyRotationHistory, err error) {
// 	err = cdc.Unmarshal(value, &hi)
// 	return hi, err
// }

// func MustMarshalConsPubKeyRotationHistory(cdc codec.BinaryCodec, history ConsPubKeyRotationHistory) []byte {
// 	hi, err := MarshalConsPubKeyRotationHistory(cdc, history)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return hi
// }

// func MarshalConsPubKeyRotationHistory(cdc codec.BinaryCodec, hi ConsPubKeyRotationHistory) ([]byte, error) {
// 	return cdc.Marshal(&hi)
// }

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (hi ConsPubKeyRotationHistory) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var oldPubKey cryptotypes.PubKey
	err := unpacker.UnpackAny(hi.OldConsPubkey, &oldPubKey)
	if err != nil {
		return err
	}
	var newPubKey cryptotypes.PubKey
	return unpacker.UnpackAny(hi.NewConsPubkey, &newPubKey)
}
