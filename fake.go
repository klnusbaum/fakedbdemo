package main

import (
	"fmt"
	"strings"
)

type Value interface{}

type FakeDB struct {
	TableKeys map[string][]string
	Data      map[string]map[string]map[string]Value
}

func NewFake(tableKeys map[string][]string) *FakeDB {
	return &FakeDB{
		TableKeys: tableKeys,
		Data:      make(map[string]map[string]map[string]Value),
	}
}

func (f *FakeDB) InsertRow(table string, row map[string]Value) {
	pk := buildPartitionKey(f.TableKeys[table], row)
	if _, ok := f.Data[table]; !ok {
		f.Data[table] = make(map[string]map[string]Value)
	}

	f.Data[table][pk] = row
}

func (f *FakeDB) GetRow(table string, row map[string]Value) map[string]Value {
	pk := buildPartitionKey(f.TableKeys[table], row)
	return f.Data[table][pk]
}

func buildPartitionKey(keyFields []string, row map[string]Value) string {
	pkValues := make([]string, 0)
	for _, keyField := range keyFields {
		pkValues = append(pkValues, valString(row[keyField]))
	}

	return strings.Join(pkValues, "-")
}

func valString(value Value) string {
	switch t := value.(type) {
	case string:
		return t
	case int:
		return fmt.Sprintf("%d", t)
	default:
		return ""
	}
}
