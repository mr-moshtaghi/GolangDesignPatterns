package state

import (
	"design-patterns/state/version_1/document"
	"fmt"
)

type PublishedState struct {
	BaseState
}

func NewPublishedState() document.State {
	return &PublishedState{}
}

func (s *PublishedState) GetName() string { return "Published" }

func (s *PublishedState) Archive() error {
	fmt.Println("PublishedState: Archiving document...")
	s.document.SetState(NewArchivedState())
	fmt.Println("PublishedState: Document archived.")
	return nil
}
