package parsers

import (
	"os"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONParse(t *testing.T) {
	entry, _ := os.Open("test_suites/test.json")

	hc := make(chan Hotel)
	go JSONParse(entry, hc)
	result := <-hc
	close(hc)
	sort.Strings(result.Images)
	assert.Equal(t, expectedResult, result)
}
