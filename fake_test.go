package main

import "testing"

func TestFake(t *testing.T) {
	fdb := NewFake(map[string][]string{
		"table1": []string{"field1", "field2"},
	})

	fdb.InsertRow("table1", map[string]Value{
		"field1": Value(1),
		"field2": Value("hello"),
		"field3": Value("val1"),
		"field4": Value("val2"),
	})

	row := fdb.GetRow("table1", map[string]Value{
		"field1": Value(1),
		"field2": Value("hello"),
	})

	if row["field3"] != "val1" {
		t.Fatalf("feild3 did not equal val1: %s", row["field3"])
	}
	if row["field4"] != "val2" {
		t.Fatalf("feild3 did not equal val2: %s", row["field4"])
	}
}
