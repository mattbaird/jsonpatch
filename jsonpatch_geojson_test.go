package jsonpatch

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

var point = `{"type":"Point", "coordinates":[0.0, 1.0]}`
var lineString = `{"type":"LineString", "coordinates":[[0.0, 1.0], [2.0, 3.0]]}`

func TestPointLineStringReplace(t *testing.T) {
	patch, e := CreatePatch([]byte(point), []byte(lineString))
	assert.NoError(t, e)
	for i, v := range patch {
		fmt.Printf("patch[%d]: %#v\n", i, v)
	}
}

func TestLineStringPointReplace(t *testing.T) {
	patch, e := CreatePatch([]byte(lineString), []byte(point))
	assert.NoError(t, e)
	for i, v := range patch {
		fmt.Printf("patch[%d]: %#v\n", i, v)
	}
}
