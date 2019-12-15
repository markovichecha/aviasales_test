package parsers

import (
	"os"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXMLParse(t *testing.T) {
	entry, _ := os.Open("test_suites/test.xml")
	expectedResult.Rating = 0

	hc := make(chan Hotel)
	go XMLParse(entry, hc)
	result := <-hc
	close(hc)
	sort.Strings(result.Images)
	assert.Equal(t, expectedResult, result)
}
