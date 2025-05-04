package main

import (
	"fmt"
	"strings"
	"time"
)

type UserData struct {
	ID                 string
	Name               string
	Email              string
	Phone              string
	City               string
	ProcessedTimestamp string
}

type DataProcessingStrategy interface {
	Process([]UserData) ([]UserData, error)
}

type NormalizationStrategy struct {
	FieldToNormalize []string
}

func (ns NormalizationStrategy) Process(data []UserData) ([]UserData, error) {
	fmt.Println("Applying Normalization Strategy...")
	processedData := make([]UserData, len(data))
	copy(processedData, data)

	for i := range processedData {
		user := &processedData[i]
		for _, fieldName := range ns.FieldToNormalize {
			switch fieldName {
			case "Name":
				user.Name = strings.TrimSpace(user.Name)
				user.Name = strings.ToLower(user.Name)
			case "Email":
				user.Email = strings.TrimSpace(user.Email)
				user.Email = strings.ToLower(user.Email)
			case "City":
				user.City = strings.TrimSpace(user.City)
				user.City = strings.ToLower(user.City)
			}
		}
	}
	return processedData, nil
}

type RedactionStrategy struct {
	FieldsToRedact []string
}

func (rs RedactionStrategy) Process(data []UserData) ([]UserData, error) {
	fmt.Println("Applying Redaction Strategy...")
	processedData := make([]UserData, len(data))
	copy(processedData, data)
	placeholder := "***REDACTED***"
	for i := range processedData {
		user := &processedData[i] // Get a pointer

		for _, fieldName := range rs.FieldsToRedact {
			switch fieldName {
			case "Email":
				user.Email = placeholder
			case "Phone":
				user.Phone = placeholder
				// Add more fields as needed
			}
		}
	}
	return processedData, nil
}

type TimestampStrategy struct{}

func (s *TimestampStrategy) Process(data []UserData) ([]UserData, error) {
	fmt.Println("Applying Timestamp Strategy...")
	processedData := make([]UserData, len(data))
	copy(processedData, data) // Work on a copy

	currentTime := time.Now().Format(time.RFC3339)

	for i := range processedData {
		user := &processedData[i] // Get a pointer
		user.ProcessedTimestamp = currentTime
	}
	return processedData, nil
}

type DataProcessor struct {
	strategy DataProcessingStrategy
}

func (dp *DataProcessor) SetStrategy(strategy DataProcessingStrategy) {
	dp.strategy = strategy
}

func (dp *DataProcessor) ProcessData(data []UserData) ([]UserData, error) {
	if dp.strategy == nil {
		return nil, fmt.Errorf("no processing strategy set")
	}
	fmt.Println("Processing data using the configured strategy...")
	return dp.strategy.Process(data)
}

func main() {
	sampleData := []UserData{
		{ID: "1", Name: "  ALICE  ", Email: "Alice.Wonderland@example.com", Phone: "123-456-7890", City: "  NEW YORK  "},
		{ID: "2", Name: "BOB", Email: "Bob.Builder@example.com", Phone: "987-654-3210", City: "los angeles"},
		{ID: "3", Name: "Charlie", Email: "charlie.chaplin@example.com", Phone: "555-1212", City: "london"},
	}
	processor := &DataProcessor{}

	fmt.Println("--- Scenario 1: Normalization & Redaction ---")

	normalizationStrategy := &NormalizationStrategy{
		FieldToNormalize: []string{"Name", "Email", "City"},
	}
	processor.SetStrategy(normalizationStrategy)
	normalizedData, err := processor.ProcessData(sampleData)
	if err != nil {
		fmt.Println("Error during normalization:", err)
		return
	}
	printUserData(normalizedData)

	redactionStrategy := &RedactionStrategy{
		FieldsToRedact: []string{"Email", "Phone"},
	}
	processor.SetStrategy(redactionStrategy)
	redactedData, err := processor.ProcessData(normalizedData)
	if err != nil {
		fmt.Println("Error during redaction:", err)
		return
	}
	fmt.Println("\nRedacted Data:")
	printUserData(redactedData)

	fmt.Println("\n--- Scenario 2: Timestamping ---")

	timestampStrategy := &TimestampStrategy{}

	processor.SetStrategy(timestampStrategy)
	timestampedData, err := processor.ProcessData(sampleData)
	if err != nil {
		fmt.Println("Error during timestamping:", err)
		return
	}
	fmt.Println("Timestamped Data:")
	printUserData(timestampedData)

	fmt.Println("\n--- Scenario 3: No Strategy ---")

	processor.SetStrategy(nil)
	_, err = processor.ProcessData(sampleData)
	if err != nil {
		fmt.Println("Error (Expected):", err)
	}
}

func printUserData(data []UserData) {
	for _, user := range data {
		fmt.Printf("  ID: %s, Name: %s, Email: %s, Phone: %s, City: %s, Timestamp: %s\n",
			user.ID, user.Name, user.Email, user.Phone, user.City, user.ProcessedTimestamp)
	}
}
