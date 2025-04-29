// document/state/states.go
package state

import (
	"design-patterns/state/document"
	"fmt"
)

// BaseState provides common functionality for all states (optional but helpful).
// It holds the reference to the context (document.Document).
// Note: Embedding BaseState in concrete states is common, but you must
// ensure the SetContext call on the concrete state sets the *embedded*
// BaseState's document field.
type BaseState struct {
	doc *document.Document // Reference to the context (from parent package)
}

// SetContext implementation for BaseState.
// Concrete states embedding BaseState will implicitly have this method.
func (b *BaseState) SetContext(d *document.Document) {
	b.doc = d
}

// Default implementations for operations that are not valid in a state.
// These methods are part of the document.DocumentState interface.
func (b *BaseState) Save() error {
	return fmt.Errorf("operation Save() not allowed in state %s", b.GetName())
}
func (b *BaseState) SubmitForReview() error {
	return fmt.Errorf("operation SubmitForReview() not allowed in state %s", b.GetName())
}
func (b *BaseState) Approve() error {
	return fmt.Errorf("operation Approve() not allowed in state %s", b.GetName())
}
func (b *BaseState) Reject() error {
	return fmt.Errorf("operation Reject() not allowed in state %s", b.GetName())
}
func (b *BaseState) Archive() error {
	return fmt.Errorf("operation Archive() not allowed in state %s", b.GetName())
}
func (b *BaseState) GetName() error {
	return nil
}

func (b *BaseState) SetContext(d *Document) {
	return nil
}

// --- Concrete State Implementations ---

// DraftState: Represents the Draft state.
type DraftState struct {
	BaseState // Embed BaseState for common methods and context reference
}

// GetName returns the name of this state.
func (s *DraftState) GetName() string { return "Draft" }

// Override allowed operations from BaseState.
func (s *DraftState) Save() error {
	fmt.Println("DraftState: Saving document...")
	// Simulate saving logic
	fmt.Println("DraftState: Document saved.")
	return nil // Saving is always allowed in Draft state
}

func (s *DraftState) SubmitForReview() error {
	fmt.Println("DraftState: Submitting document for review...")
	// Simulate submission process
	// Transition to Moderation state using the context reference
	// We need to instantiate ModerationState, which is in the same package (state)
	s.doc.SetState(&ModerationState{}) // Change state in Context (Document)
	fmt.Println("DraftState: Document submitted for review.")
	return nil
}

// ModerationState: Represents the Moderation state.
type ModerationState struct {
	BaseState
}

func (s *ModerationState) GetName() string { return "Moderation" }

// Override allowed operations.
func (s *ModerationState) Approve() error {
	fmt.Println("ModerationState: Approving document...")
	// Simulate approval process
	// Transition to Published state
	s.doc.SetState(&PublishedState{}) // Change state in Context
	fmt.Println("ModerationState: Document approved.")
	return nil
}

func (s *ModerationState) Reject() error {
	fmt.Println("ModerationState: Rejecting document...")
	// Simulate rejection process
	// Transition back to Draft state
	s.doc.SetState(&DraftState{}) // Change state in Context
	fmt.Println("ModerationState: Document rejected, moved back to Draft.")
	return nil
}

// PublishedState: Represents the Published state.
type PublishedState struct {
	BaseState
}

func (s *PublishedState) GetName() string { return "Published" }

// Override allowed operations.
func (s *PublishedState) Archive() error {
	fmt.Println("PublishedState: Archiving document...")
	// Simulate archiving process
	// Transition to Archived state
	s.doc.SetState(&ArchivedState{}) // Change state in Context
	fmt.Println("PublishedState: Document archived.")
	return nil
}

// ArchivedState: Represents the Archived state.
type ArchivedState struct {
	BaseState
}

func (s *ArchivedState) GetName() string { return "Archived" }

// No operations overridden here. They will use BaseState's default error messages.

// --- State Factory Functions (Optional but can be helpful) ---
// These can provide a cleaner way to create state instances if constructors are complex
// or you want to control state creation via interfaces or dependency injection.
// For this example, direct struct instantiation (&State{}) is sufficient.

/*
func NewDraftState(doc *document.Document) document.DocumentState {
	s := &DraftState{}
	s.SetContext(doc) // Set the context immediately upon creation
	return s
}
// Define similar factories for other states
*/
