package document

import (
	"fmt"
)

type DocumentState interface {
	Save() error
	SubmitForReview() error
	Approve() error
	Reject() error
	Archive() error
	GetName() string
	SetContext(d *Document)
}

// Document is the Context that holds the current state.
type Document struct {
	State   DocumentState // Reference to the current state
	Title   string
	Content string
}

// NewDocument creates a new document with an initial state.
// The initial state (DraftState) is created within this package's logic.
func NewDocument(title, content string) *Document {
	doc := &Document{
		Title:   title,
		Content: content,
	}
	doc.SetState(NewDraftState(doc))

	return doc
}

// SetState is the method used by states to change the document's current state.
// This method is public as states (which are in another package) need to call it.
func (d *Document) SetState(state DocumentState) {
	d.State = state
	state.SetContext(d) // Let the new state know about its context
	fmt.Printf("--- Document state changed to: %s ---\n", d.State.GetName())
}

// Delegate methods from Context to the current State.
// These methods are the public API clients use to interact with the Document.
func (d *Document) Save() error {
	fmt.Printf("Document: Calling Save() in state %s\n", d.State.GetName())
	return d.State.Save()
}

func (d *Document) SubmitForReview() error {
	fmt.Printf("Document: Calling SubmitForReview() in state %s\n", d.State.GetName())
	return d.State.SubmitForReview()
}

func (d *Document) Approve() error {
	fmt.Printf("Document: Calling Approve() in state %s\n", d.State.GetName())
	return d.State.Approve()
}

func (d *Document) Reject() error {
	fmt.Printf("Document: Calling Reject() in state %s\n", d.State.GetName())
	return d.State.Reject()
}

func (d *Document) Archive() error {
	fmt.Printf("Document: Calling Archive() in state %s\n", d.State.GetName())
	return d.State.Archive()
}

// GetCurrentStateName provides the current state's name.
func (d *Document) GetCurrentStateName() string {
	if d.State == nil {
		return "Unknown"
	}
	return d.State.GetName()
}

// --- State Factory Functions (Can be in document/state or here for simplicity) ---
// For cleaner separation, let's put these in document/state and import them.
// This requires document to import its subpackage state, which is acceptable.
// import "your_module_name/document/state" // Assuming your module name is your_module_name

// Placeholder for state factories - actual implementations in document/state/states.go
// We define these here *just* so NewDocument can call one.
// The actual factory functions will be in the state subpackage and imported.
// You need to uncomment the import statement for the state package once you create it.

/*
func NewDraftState(doc *Document) DocumentState {
    // Actual implementation will be in document/state
    // This is just a placeholder signature matching what NewDocument expects
    return nil
}

func NewModerationState(doc *Document) DocumentState { return nil }
func NewPublishedState(doc *Document) DocumentState { return nil }
func NewArchivedState(doc *Document) DocumentState { return nil }
*/

// --- Let's adjust NewDocument slightly to call concrete constructors directly for this example's simplicity ---
// In a real project, factories or dependency injection for states might be preferred.
// We will keep the state structs public in the subpackage and use them directly here.

// Corrected NewDocument implementation:
func NewDocumentFixed(title, content string) *Document {
	doc := &Document{
		Title:   title,
		Content: content,
	}
	// Use the concrete state struct from the subpackage directly
	// Requires importing the state subpackage
	// import "your_module_name/document/state" // Replace your_module_name
	// For this local example, we'll just use a placeholder for DraftState initially
	// The actual instantiation will be in main or within the state package.
	// Let's make the state constructors public in the state package and use them from main or via factories.
	// A common State pattern approach is for the Context to create the initial state itself.
	// So, we'll put the state struct definitions in document/state and make them public.
	// The NewDocument will then import and use the initial state's struct.

	// Let's revert to the original simple instantiation but update the SetState call.
	// The states themselves will set the context.
	// We need the state structs to be defined in the subpackage.
	// The initial state needs to be instantiated here or passed in.
	// The cleanest way is for NewDocument to instantiate the initial state from the subpackage.
	// So, NewDocument will import the state subpackage.

	// Final plan for NewDocument:
	// 1. Create the Document struct.
	// 2. Instantiate the initial state (DraftState) using its public constructor/struct from document/state.
	// 3. Call doc.SetState() with the new state. SetState will then call state.SetContext().
	// This requires document package to import the state subpackage.
	// import "your_module_name/document/state" // Uncomment and replace your_module_name

	// Example assuming state.DraftState is a public struct in document/state
	// initialState := &state.DraftState{}
	// doc.SetState(initialState)
	// The above needs actual state struct definition in the subpackage.

	// Let's define the initial state creation *after* we define the state package.
	// For now, keep the basic structure.

	// Re-implementing NewDocument to create the initial state:
	// This requires importing the state subpackage.
	// import "your_module_name/document/state" // Uncomment and replace your_module_name

	// Example:
	// doc := &Document{Title: title, Content: content}
	// draftState := &state.DraftState{} // Assuming state.DraftState is public
	// doc.SetState(draftState)
	// return doc

	// Let's finalize the NewDocument after the state package is defined.
	// The code below is the final version of document/document.go
}

// Corrected and Final document/document.go
