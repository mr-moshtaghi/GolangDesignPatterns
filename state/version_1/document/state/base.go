package state

import (
	"design-patterns/state/version_1/document"
	"fmt"
)

type BaseState struct {
	document *document.Document // Reference to the context
}

func (b *BaseState) SetContext(d *document.Document) {
	b.document = d
}

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

func (b *BaseState) GetName() string {
	return ""
}
