package parsers

import (
	"encoding/csv"
	"os"
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCSVParse(t *testing.T) {
	entry, _ := os.Open("test_suites/test.csv")
	hc := make(chan Hotel)
	go CSVParse(entry, hc)
	result := <-hc
	close(hc)
	sort.Strings(result.Images)
	assert.Equal(t, expectedResult, result)
}

func TestFieldToIndex(t *testing.T) {
	entry, _ := os.Open("test_suites/test.csv")
	expectedResult := map[int]fieldInfo{
		0:  fieldInfo{"Name", reflect.String},
		1:  fieldInfo{"Description", reflect.String},
		2:  fieldInfo{"Address", reflect.String},
		3:  fieldInfo{"City", reflect.String},
		4:  fieldInfo{"Country", reflect.String},
		5:  fieldInfo{"CountryCode", reflect.String},
		6:  fieldInfo{"StarRating", reflect.Uint8},
		7:  fieldInfo{"Longitude", reflect.Float64},
		8:  fieldInfo{"Latitude", reflect.Float64},
		9:  fieldInfo{"Images", reflect.Slice},
		10: fieldInfo{"Images", reflect.Slice},
		11: fieldInfo{"Images", reflect.Slice},
		12: fieldInfo{"Images", reflect.Slice},
		13: fieldInfo{"Images", reflect.Slice},
		14: fieldInfo{"Rating", reflect.Float64},
	}

	r := csv.NewReader(entry)
	fields, _ := r.Read()
	result := fieldToIndex(fields)
	assert.Equal(t, expectedResult, result)
}
