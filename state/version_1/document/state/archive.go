package state

import (
	"design-patterns/state/version_1/document"
)

type ArchivedState struct {
	BaseState // Embed BaseState for common methods
}

func NewArchivedState() document.State {
	return &ArchivedState{}
}

func (s *ArchivedState) GetName() string { return "Archived" }
