package test

import (
	"time"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"github.com/tendermint/tendermint/types"
)

func MakeVote(
	val types.PrivValidator,
	chainID string,
	valIndex int32,
	height int64,
	round int32,
	step int,
	blockID types.BlockID,
	time time.Time,
) (*types.Vote, error) {
	pubKey, err := val.GetPubKey()
	if err != nil {
		return nil, err
	}

	v := &types.Vote{
		ValidatorAddress: pubKey.Address(),
		ValidatorIndex:   valIndex,
		Height:           height,
		Round:            round,
		Type:             tmproto.SignedMsgType(step),
		BlockID:          blockID,
		Timestamp:        time,
	}

	vpb := v.ToProto()
	if err := val.SignVote(chainID, vpb); err != nil {
		return nil, err
	}

	v.Signature = vpb.Signature
	if tmproto.SignedMsgType(step) == tmproto.PrecommitType {
		v.ExtensionSignature = vpb.ExtensionSignature
	}
	return v, nil
}

func MakePrecommit(
	val types.PrivValidator,
	valIndex int32,
	height int64,
	round int32,
	blockID types.BlockID,
	time time.Time,
) (*types.Vote, error) {
	return MakeVote(val, DefaultTestChainID, valIndex, height, round, int(tmproto.PrecommitType), blockID, time)
}
