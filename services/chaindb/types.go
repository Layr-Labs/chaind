// Copyright © 2020, 2021 Weald Technology Trading.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package chaindb

import (
	"time"

	"github.com/attestantio/go-eth2-client/spec/phase0"
)

// Block holds information about a block.
type Block struct {
	Slot             phase0.Slot
	ProposerIndex    phase0.ValidatorIndex
	Root             phase0.Root
	Graffiti         []byte
	RANDAOReveal     phase0.BLSSignature
	BodyRoot         phase0.Root
	ParentRoot       phase0.Root
	StateRoot        phase0.Root
	Canonical        *bool
	ETH1BlockHash    []byte
	ETH1DepositCount uint64
	ETH1DepositRoot  phase0.Root
}

// Validator holds information about a validator.
type Validator struct {
	PublicKey                  phase0.BLSPubKey
	Index                      phase0.ValidatorIndex
	EffectiveBalance           phase0.Gwei
	Slashed                    bool
	ActivationEligibilityEpoch phase0.Epoch
	ActivationEpoch            phase0.Epoch
	ExitEpoch                  phase0.Epoch
	WithdrawableEpoch          phase0.Epoch
}

// ValidatorBalance holds information about a validator's balance at a given epoch.
type ValidatorBalance struct {
	Index            phase0.ValidatorIndex
	Epoch            phase0.Epoch
	Balance          phase0.Gwei
	EffectiveBalance phase0.Gwei
}

// AggregateValidatorBalance holds aggreated information about validators' balances at a given epoch.
type AggregateValidatorBalance struct {
	Epoch            phase0.Epoch
	Balance          phase0.Gwei
	EffectiveBalance phase0.Gwei
}

// BeaconCommittee holds information for beacon committees.
type BeaconCommittee struct {
	Slot      phase0.Slot
	Index     phase0.CommitteeIndex
	Committee []phase0.ValidatorIndex
}

// ProposerDuty holds information for proposer duties.
type ProposerDuty struct {
	Slot           phase0.Slot
	ValidatorIndex phase0.ValidatorIndex
}

// AttesterDuty holds information for attester duties.
type AttesterDuty struct {
	Slot           phase0.Slot
	Committee      phase0.CommitteeIndex
	ValidatorIndex phase0.ValidatorIndex
	// CommitteeIndex is the index of the validator in the committee.
	CommitteeIndex uint64
}

// Attestation holds information about an attestation included by a block.
type Attestation struct {
	InclusionSlot      phase0.Slot
	InclusionBlockRoot phase0.Root
	InclusionIndex     uint64
	Slot               phase0.Slot
	CommitteeIndex     phase0.CommitteeIndex
	AggregationBits    []byte
	AggregationIndices []phase0.ValidatorIndex
	BeaconBlockRoot    phase0.Root
	SourceEpoch        phase0.Epoch
	SourceRoot         phase0.Root
	TargetEpoch        phase0.Epoch
	TargetRoot         phase0.Root
	Canonical          *bool
	TargetCorrect      *bool
	HeadCorrect        *bool
}

// SyncAggregate holds information about a sync aggregate included in a block.
type SyncAggregate struct {
	InclusionSlot      phase0.Slot
	InclusionBlockRoot phase0.Root
	Bits               []byte
	Indices            []phase0.ValidatorIndex
}

// Deposit holds information about an Ethereum 2 deposit included by a block.
type Deposit struct {
	InclusionSlot         phase0.Slot
	InclusionBlockRoot    phase0.Root
	InclusionIndex        uint64
	ValidatorPubKey       phase0.BLSPubKey
	WithdrawalCredentials []byte
	Amount                phase0.Gwei
}

// ETH1Deposit holds information about an Ethereum 2 deposit made on the Ethereum 1 chain.
type ETH1Deposit struct {
	ETH1BlockNumber       uint64
	ETH1BlockHash         []byte
	ETH1BlockTimestamp    time.Time
	ETH1TxHash            []byte
	ETH1LogIndex          uint64
	ETH1Sender            []byte
	ETH1Recipient         []byte
	ETH1GasUsed           uint64
	ETH1GasPrice          uint64
	DepositIndex          uint64
	ValidatorPubKey       phase0.BLSPubKey
	WithdrawalCredentials []byte
	Signature             phase0.BLSSignature
	Amount                phase0.Gwei
}

// VoluntaryExit holds information about a voluntary exit included in a block.
type VoluntaryExit struct {
	InclusionSlot      phase0.Slot
	InclusionBlockRoot phase0.Root
	InclusionIndex     uint64
	ValidatorIndex     phase0.ValidatorIndex
	Epoch              phase0.Epoch
}

// AttesterSlashing holds information about an attester slashing included by a block.
type AttesterSlashing struct {
	InclusionSlot               phase0.Slot
	InclusionBlockRoot          phase0.Root
	InclusionIndex              uint64
	Attestation1Indices         []phase0.ValidatorIndex
	Attestation1Slot            phase0.Slot
	Attestation1CommitteeIndex  phase0.CommitteeIndex
	Attestation1BeaconBlockRoot phase0.Root
	Attestation1SourceEpoch     phase0.Epoch
	Attestation1SourceRoot      phase0.Root
	Attestation1TargetEpoch     phase0.Epoch
	Attestation1TargetRoot      phase0.Root
	Attestation1Signature       phase0.BLSSignature
	Attestation2Indices         []phase0.ValidatorIndex
	Attestation2Slot            phase0.Slot
	Attestation2CommitteeIndex  phase0.CommitteeIndex
	Attestation2BeaconBlockRoot phase0.Root
	Attestation2SourceEpoch     phase0.Epoch
	Attestation2SourceRoot      phase0.Root
	Attestation2TargetEpoch     phase0.Epoch
	Attestation2TargetRoot      phase0.Root
	Attestation2Signature       phase0.BLSSignature
}

// ProposerSlashing holds information about a proposer slashing included by a block.
type ProposerSlashing struct {
	InclusionSlot        phase0.Slot
	InclusionBlockRoot   phase0.Root
	InclusionIndex       uint64
	Block1Root           phase0.Root
	Header1Slot          phase0.Slot
	Header1ProposerIndex phase0.ValidatorIndex
	Header1ParentRoot    phase0.Root
	Header1StateRoot     phase0.Root
	Header1BodyRoot      phase0.Root
	Header1Signature     phase0.BLSSignature
	Block2Root           phase0.Root
	Header2Slot          phase0.Slot
	Header2ProposerIndex phase0.ValidatorIndex
	Header2ParentRoot    phase0.Root
	Header2StateRoot     phase0.Root
	Header2BodyRoot      phase0.Root
	Header2Signature     phase0.BLSSignature
}

// ValidatorEpochSummary provides a summary of a validator's operations for an epoch.
type ValidatorEpochSummary struct {
	Index                     phase0.ValidatorIndex
	Epoch                     phase0.Epoch
	ProposerDuties            int
	ProposalsIncluded         int
	AttestationIncluded       bool
	AttestationTargetCorrect  *bool
	AttestationHeadCorrect    *bool
	AttestationInclusionDelay *int
}

// BlockSummary provides a summary of an epoch.
type BlockSummary struct {
	Slot                          phase0.Slot
	AttestationsForBlock          int
	DuplicateAttestationsForBlock int
	VotesForBlock                 int
}

// EpochSummary provides a summary of an epoch.
type EpochSummary struct {
	Epoch                         phase0.Epoch
	ActivationQueueLength         int
	ActivatingValidators          int
	ActiveValidators              int
	ActiveRealBalance             phase0.Gwei
	ActiveBalance                 phase0.Gwei
	AttestingValidators           int
	AttestingBalance              phase0.Gwei
	TargetCorrectValidators       int
	TargetCorrectBalance          phase0.Gwei
	HeadCorrectValidators         int
	HeadCorrectBalance            phase0.Gwei
	AttestationsForEpoch          int
	AttestationsInEpoch           int
	DuplicateAttestationsForEpoch int
	ProposerSlashings             int
	AttesterSlashings             int
	Deposits                      int
	ExitingValidators             int
	CanonicalBlocks               int
}

// SyncCommittee holds information for sync committees.
type SyncCommittee struct {
	Period    uint64
	Committee []phase0.ValidatorIndex
}
