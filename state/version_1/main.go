// main.go
package main

import (
	"design-patterns/state/version_1/document"
	"design-patterns/state/version_1/document/state"
	"fmt"
)

func main() {
	fmt.Println("Creating a new doc...")
	doc := document.NewDocument("My Awesome Article", "This is the content of the article.")

	initialState := state.NewDraftState()
	doc.SetState(initialState)

	fmt.Printf("Current state: %s\n", doc.GetCurrentStateName())
	fmt.Println("--------------------")

	err := doc.Save()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("Current state: %s\n", doc.GetCurrentStateName())
	fmt.Println("--------------------")

	err = doc.Approve()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("Current state: %s\n", doc.GetCurrentStateName())
	fmt.Println("--------------------")

	err = doc.SubmitForReview()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("Current state: %s\n", doc.GetCurrentStateName())
	fmt.Println("--------------------")

	err = doc.SubmitForReview()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("Current state: %s\n", doc.GetCurrentStateName())
	fmt.Println("--------------------")

	err = doc.Reject()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("Current state: %s\n", doc.GetCurrentStateName())
	fmt.Println("--------------------")

	err = doc.SubmitForReview()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("Current state: %s\n", doc.GetCurrentStateName())
	fmt.Println("--------------------")

	err = doc.Approve()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("Current state: %s\n", doc.GetCurrentStateName())
	fmt.Println("--------------------")

	err = doc.Save()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("Current state: %s\n", doc.GetCurrentStateName())
	fmt.Println("--------------------")

	err = doc.Archive()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("Current state: %s\n", doc.GetCurrentStateName())
	fmt.Println("--------------------")

	err = doc.Save()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("Current state: %s\n", doc.GetCurrentStateName())
	fmt.Println("--------------------")

	err = doc.Approve()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("Current state: %s\n", doc.GetCurrentStateName())
	fmt.Println("--------------------")
}
