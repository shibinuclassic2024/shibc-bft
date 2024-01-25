package core

import (
	"github.com/shibinuclassic2024/shibc-bft/messages"
	"github.com/shibinuclassic2024/shibc-bft/messages/proto"
)


type MessageConstructor interface {
	// BuildPrePrepareMessage builds a PREPREPARE message based on the passed in view and proposal
	BuildPrePrepareMessage(
		rawProposal []byte,
		certificate *proto.RoundChangeCertificate,
		view *proto.View,
	) *proto.Message

	
	BuildPrepareMessage(proposalHash []byte, view *proto.View) *proto.Message

	
	BuildCommitMessage(proposalHash []byte, view *proto.View) *proto.Message


	BuildRoundChangeMessage(
		proposal *proto.Proposal,
		certificate *proto.PreparedCertificate,
		view *proto.View,
	) *proto.Message
}

// Verifier defines the verifier interface
type Verifier interface {
	// IsValidProposal checks if the proposal is valid
	IsValidProposal(rawProposal []byte) bool

	IsValidValidator(msg *proto.Message) bool

	// IsProposer checks if the passed in ID is the Proposer for current view (sequence, round)
	IsProposer(id []byte, height, round uint64) bool

	// IsValidProposalHash checks if the hash matches the proposal
	IsValidProposalHash(proposal *proto.Proposal, hash []byte) bool

	// IsValidCommittedSeal checks
	// if signature for proposal hash in committed seal is signed by a validator
	IsValidCommittedSeal(proposalHash []byte, committedSeal *messages.CommittedSeal) bool
}

// Backend defines an interface all backend implementations
// need to implement
type Backend interface {
	MessageConstructor
	Verifier
	ValidatorBackend

	// BuildProposal builds a new proposal for the given view (height and round)
	BuildProposal(view *proto.View) []byte

	
	InsertProposal(proposal *proto.Proposal, committedSeals []*messages.CommittedSeal)

	// ID returns the validator's ID
	ID() []byte
}
