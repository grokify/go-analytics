package table

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	//"github.com/grokify/gotilla/encoding/csvutil"
)

const StyleSimple = "border:1px solid #000;border-collapse:collapse"

var DebugReadCSV = false // should not need to use this.

// Table is useful for working on CSV data
type Table struct {
	Name       string
	Columns    []string
	Records    [][]string
	FormatMap  map[int]string
	FormatFunc func(val string, colIdx uint) (interface{}, error)
	ID         string
	Class      string
	Style      string
}

func NewTable() Table {
	return Table{
		Columns:   []string{},
		Records:   [][]string{},
		FormatMap: map[int]string{}}
}

// LoadMergedRows is used to load data from `[][]string` sources
// like csv.ReadAll()
func (t *Table) LoadMergedRows(data [][]string) {
	if len(data) == 0 {
		return
	}
	t.Columns = data[0]
	if len(data) > 1 {
		t.Records = data[1:]
	}
}

func (t *Table) ColumnIndex(wantCol string) int {
	for i, col := range t.Columns {
		if col == wantCol {
			return i
		}
	}
	return -1
}

func (t *Table) ColumnsValuesDistinct(wantCols []string, stripSpace bool) (map[string]int, error) {
	data := map[string]int{}
	wantIdxs := []int{}
	maxIdx := -1
	for _, wantCol := range wantCols {
		wantIdx := t.ColumnIndex(wantCol)
		if wantIdx < 0 {
			return data, fmt.Errorf("Column Not Found [%v]", wantCol)
		}
		wantIdxs = append(wantIdxs, wantIdx)
		if wantIdx > maxIdx {
			maxIdx = wantIdx
		}
	}
	for _, rec := range t.Records {
		if len(rec) > maxIdx {
			vals := []string{}
			for _, wantIdx := range wantIdxs {
				val := rec[wantIdx]
				if stripSpace {
					val = strings.TrimSpace(val)
				}
				vals = append(vals, val)
			}
			valsStr := strings.Join(vals, " ")
			_, ok := data[valsStr]
			if !ok {
				data[valsStr] = 0
			}
			data[valsStr]++
		}
	}
	return data, nil
}

func (t *Table) ColumnValuesDistinct(wantCol string) (map[string]int, error) {
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
			data[val]++
		}
	}
	return data, nil
}

func (t *Table) ColumnValuesMinMax(wantCol string) (string, string, error) {
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

func (t *Table) RecordValue(wantCol string, record []string) (string, error) {
	idx := t.ColumnIndex(wantCol)
	if idx < 0 {
		return "", fmt.Errorf("Column Not Found [%v]", wantCol)
	}
	if idx >= len(record) {
		return "", fmt.Errorf("Record does not have enough columns [%v]", idx+1)
	}
	return record[idx], nil
}

func (t *Table) RecordValueOrEmpty(wantCol string, record []string) string {
	val, err := t.RecordValue(wantCol, record)
	if err != nil {
		return ""
	}
	return val
}

func (t *Table) NewTableFiltered(wantColNameValues map[string]string) (Table, error) {
	t2 := Table{Columns: t.Columns}
	records, err := t.FilterRecords(wantColNameValues)
	if err != nil {
		return t2, err
	}
	t2.Records = records
	return t2, nil
}

func (t *Table) FilterRecords(wantColNameValues map[string]string) ([][]string, error) {
	data := [][]string{}
	wantColIndexes := map[string]int{}
	maxIdx := -1
	for wantColName := range wantColNameValues {
		wantColIdx := t.ColumnIndex(wantColName)
		if wantColIdx < 0 {
			return data, fmt.Errorf("Column Not Found [%v]", wantColName)
		}
		if wantColIdx > maxIdx {
			maxIdx = wantColIdx
		}
		wantColIndexes[wantColName] = wantColIdx
	}
RECORDS:
	for _, rec := range t.Records {
		if len(rec) > maxIdx {
			for wantColName, wantColIdx := range wantColIndexes {
				colValue := rec[wantColIdx]
				wantColValue, ok := wantColNameValues[wantColName]
				if !ok {
					return data, fmt.Errorf("Column Name [%v] has no desired value", wantColName)
				}
				if colValue != wantColValue {
					continue RECORDS
				}
			}
			data = append(data, rec)
		}
	}
	return data, nil
}

func (t *Table) ColIndex(colName string) int {
	for i, tryColName := range t.Columns {
		if tryColName == colName {
			return i
		}
	}
	return -1
}

func (t *Table) ColValuesByColName(colName string) ([]string, error) {
	colIdx := t.ColIndex(colName)
	if colIdx < 0 {
		return []string{}, fmt.Errorf("E_NO_COL_FOR_NAME [%s]", colName)
	}
	vals := []string{}
	for _, row := range t.Records {
		if colIdx < len(row) {
			vals = append(vals, row[colIdx])
		} else {
			return vals, fmt.Errorf("E_COL_IDX [%d] ROW_LEN [%d]", colIdx, len(row))
		}
	}
	return vals, nil
}

func (t *Table) WriteXLSX(path, sheetname string) error {
	t.Name = sheetname
	return WriteXLSX(path, t)
}

func (t *Table) WriteCSV(path string) error {
	return WriteCSV(path, t)
}

func (t *Table) RecordToMSS(record []string) map[string]string {
	mss := map[string]string{}
	for i, key := range t.Columns {
		if i < len(t.Columns) {
			mss[key] = record[i]
		}
	}
	return mss
}

func (t *Table) ToSliceMSS() []map[string]string {
	slice := []map[string]string{}
	for _, rec := range t.Records {
		slice = append(slice, t.RecordToMSS(rec))
	}
	return slice
}
