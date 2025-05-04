package state

import (
	"design-patterns/state/version_1/document"
	"fmt"
)

type DraftState struct {
	BaseState // Embed BaseState for common methods
}

func NewDraftState() document.State {
	return &DraftState{}
}

func (s *DraftState) GetName() string { return "Draft" }

func (s *DraftState) Save() error {
	fmt.Println("DraftState: Saving document...")
	fmt.Println("DraftState: Document saved.")
	return nil
}

func (s *DraftState) SubmitForReview() error {
	fmt.Println("DraftState: Submitting document for review...")
	s.document.SetState(NewModerationState())
	fmt.Println("DraftState: Document submitted for review.")
	return nil
}
