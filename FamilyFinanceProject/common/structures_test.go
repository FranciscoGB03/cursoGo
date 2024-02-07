package common

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestUserStruct(t *testing.T) {
	expectedFields := map[string]interface{}{
		"ID":        reflect.Int,
		"FirstName": reflect.String,
		"LastName":  reflect.String,
		"Relation":  reflect.String,
		"Income":    reflect.Float64,
		"Expenses":  reflect.Float64,
		"Balance":   reflect.Float64,
		"Budget":    reflect.Float64,
		"Bill":      reflect.Float64,
	}

	testStructFields(t, "User", User{}, expectedFields)
}

func TestIncomingStruct(t *testing.T) {
	expectedFields := map[string]interface{}{
		"Income_ID":  reflect.Int,
		"Date":       reflect.String,
		"Account_ID": reflect.Int,
		"Vendor":     reflect.String,
		"Category":   reflect.String,
		"Deposit":    reflect.Float64,
		"Balance":    reflect.Float64,
		"Notes":      reflect.String,
	}

	testStructFields(t, "Incoming", Incoming{}, expectedFields)
}

func TestExpensesStruct(t *testing.T) {
	expectedFields := map[string]interface{}{
		"Expenses_ID": reflect.Int,
		"Date":        reflect.String,
		"Account_ID":  reflect.Int,
		"Category":    reflect.String,
		"Vendor":      reflect.String,
		"Payment":     reflect.Float64,
		"Balance":     reflect.Float64,
		"Notes":       reflect.String,
	}

	testStructFields(t, "Expenses", Expenses{}, expectedFields)
}

// Helper function to test the fields of a struct

func testStructFields(t *testing.T, structName string, s interface{}, expectedFields interface{}) {

	expectedData := s

	// Marshal to JSON
	expectedJSONData, err := json.Marshal(expectedData)
	if err != nil {
		t.Fatalf("Failed to marshal %s to JSON: %v", structName, err)
	}

	// Marshal again after unmarshaling
	actualJSONData, err := json.Marshal(s)
	if err != nil {
		t.Fatalf("Failed to marshal %s to JSON after unmarshaling: %v", structName, err)
	}

	// Compare JSON
	if string(actualJSONData) != string(expectedJSONData) {
		t.Errorf("JSON marshaling and unmarshaling for %s failed. Expected: %s, Got: %s", structName, string(expectedJSONData), string(actualJSONData))
	}
}
