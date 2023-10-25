package utils

import "testing"

func TestCreateQuery(t *testing.T) {
	filter := SqlFilter{
		Data:     map[string]string{"age": "12", "name": "Alex", "surname": "Rud"},
		FindKeys: []string{"age", "name"},
	}
	if filter.CreateQuery() != "age = 12 AND name = Alex" {
		t.Error("CreateQuery() != age = 12 AND name = Alex")
	}
}
