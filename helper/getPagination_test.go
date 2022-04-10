package helper

import (
	"fmt"
	"testing"
)

func TestGetCategory(t *testing.T) {

	categories, count, err := GetCategory(1, 10)

	fmt.Println(categories)

	if err != nil {
		t.Error("Expected nil, got error")
	}

	if len(categories) != 10 {
		t.Error("Expected 10, got ", len(categories))
	}

	if count != 10 {
		t.Error("Expected 10, got ", count)
	}

}
