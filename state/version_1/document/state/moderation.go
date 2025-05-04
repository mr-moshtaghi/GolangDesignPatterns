package state

import (
	"design-patterns/state/version_1/document"
	"fmt"
)

type ModerationState struct {
	BaseState // Embed BaseState for common methods
}

func NewModerationState() document.State {
	return &ModerationState{}
}

func (s *ModerationState) GetName() string { return "Moderation" }

func (s *ModerationState) Approve() error {
	fmt.Println("ModerationState: Approving document...")
	s.document.SetState(NewPublishedState())
	fmt.Println("ModerationState: Document approved.")
	return nil
}

func (s *ModerationState) Reject() error {
	fmt.Println("ModerationState: Rejecting document...")
	s.document.SetState(NewDraftState())
	fmt.Println("ModerationState: Document rejected, moved back to Draft.")
	return nil
}
