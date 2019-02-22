package table

import (
	"errors"
	"fmt"
	"sort"
)

// TableData is useful for working on CSV data
type TableData struct {
	Columns []string
	Records [][]string
}

// LoadMergedRows is used to load data from `[][]string` sources
// like csv.ReadAll()
func (t *TableData) LoadMergedRows(data [][]string) {
	if len(data) == 0 {
		return
	}
	t.Columns = data[0]
	if len(data) > 1 {
		t.Records = data[1:]
	}
}

func (t *TableData) ColumnIndex(wantCol string) int {
	for i, col := range t.Columns {
		if col == wantCol {
			return i
		}
	}
	return -1
}

func (t *TableData) ColumnValuesDistinct(wantCol string) (map[string]int, error) {
	data := map[string]int{}
	idx := t.ColumnIndex(wantCol)
	if idx < 0 {
		return data, fmt.Errorf("Column Not Found [%v]", wantCol)
	}

	for _, rec := range t.Records {
		if len(rec) > idx {
			val := rec[idx]
			_, ok := data[val]
			if !ok {
				data[val] = 0
			}
			data[val] += 1
		}
	}
	return data, nil
}

func (t *TableData) ColumnValuesMinMax(wantCol string) (string, string, error) {
	vals, err := t.ColumnValuesDistinct(wantCol)
	if err != nil {
		return "", "", err
	}
	if len(vals) == 0 {
		return "", "", errors.New("No Values Found")
	}

	arr := []string{}
	for val := range vals {
		arr = append(arr, val)
	}

	sort.Strings(arr)
	return arr[0], arr[len(arr)-1], nil
}

func (t *TableData) RecordValue(wantCol string, record []string) (string, error) {
	idx := t.ColumnIndex(wantCol)
	if idx < 0 {
		return "", fmt.Errorf("Column Not Found [%v]", wantCol)
	}
	if idx >= len(record) {
		return "", fmt.Errorf("Record does not have enough columns [%v]", idx+1)
	}
	return record[idx], nil
}