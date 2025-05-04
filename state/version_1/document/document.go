package document

import (
	"fmt"
)

type State interface {
	Save() error
	SubmitForReview() error
	Approve() error
	Reject() error
	Archive() error
	GetName() string
	SetContext(d *Document)
}

type Document struct {
	State   State
	Title   string
	Content string
}

func NewDocument(title, content string) *Document {
	doc := &Document{
		Title:   title,
		Content: content,
	}
	return doc
}

func (d *Document) SetState(state State) {
	d.State = state
	state.SetContext(d)
	fmt.Printf("--- Document state changed to: %s ---\n", d.State.GetName())
}

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

func (d *Document) GetCurrentStateName() string {
	if d.State == nil {
		return "Unknown"
	}
	return d.State.GetName()
}
