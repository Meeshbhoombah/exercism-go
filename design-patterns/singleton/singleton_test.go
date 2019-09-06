package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	counterOne := GetInstance()

	if counterOne == nil {
		t.Error("Expected pointer to Singleton after calling GetInstance(), not nil")
	}

	expectedCounter := counterOne
}

