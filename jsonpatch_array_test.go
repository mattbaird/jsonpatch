package jsonpatch

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

var arraySrc = `
{
  "spec": {
    "loadBalancerSourceRanges": [
      "192.101.0.0/16",
      "192.0.0.0/24"
    ]
  }
}
`

var arrayDst = `
{
  "spec": {
    "loadBalancerSourceRanges": [
      "192.101.0.0/24"
    ]
  }
}
`

func TestArraySame(t *testing.T) {
	patch, e := CreatePatch([]byte(arraySrc), []byte(arraySrc))
	assert.NoError(t, e)
	assert.Equal(t, len(patch), 0, "they should be equal")
}

func TestArrayBoolReplace(t *testing.T) {
	patch, e := CreatePatch([]byte(arraySrc), []byte(arrayDst))
	assert.NoError(t, e)
	assert.Equal(t, 2, len(patch), "they should be equal")
	sort.Sort(ByPath(patch))

	change := patch[0]
	assert.Equal(t, "replace", change.Operation, "they should be equal")
	assert.Equal(t, "/spec/loadBalancerSourceRanges/0", change.Path, "they should be equal")
	assert.Equal(t, "192.101.0.0/24", change.Value, "they should be equal")
	change = patch[1]
	assert.Equal(t, change.Operation, "remove", "they should be equal")
	assert.Equal(t, change.Path, "/spec/loadBalancerSourceRanges/1", "they should be equal")
	assert.Equal(t, nil, change.Value, "they should be equal")
}

func TestArrayAlmostSame(t *testing.T) {
	src := `{"Lines":[1,2,3,4,5,6,7,8,9,10]}`
	to := `{"Lines":[2,3,4,5,6,7,8,9,10,11]}`
	patch, e := CreatePatch([]byte(src), []byte(to))
	assert.NoError(t, e)
	assert.Equal(t, 2, len(patch), "they should be equal")
	sort.Sort(ByPath(patch))

	change := patch[0]
	assert.Equal(t, "remove", change.Operation, "they should be equal")
	assert.Equal(t, "/Lines/0", change.Path, "they should be equal")
	assert.Equal(t, nil, change.Value, "they should be equal")
	change = patch[1]
	assert.Equal(t, change.Operation, "add", "they should be equal")
	assert.Equal(t, change.Path, "/Lines/10", "they should be equal")
	assert.Equal(t, float64(11), change.Value, "they should be equal")
}
