package jsonpatch

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

var arrayBase = `{
  "persons": [{"name":"Ed"},{}]
}`

var arrayUpdated = `{
  "persons": [{"name":"Ed"},{},{}]
}`

func TestArrayAddMultipleEmptyObjects(t *testing.T) {
	patch, e := CreatePatch([]byte(arrayBase), []byte(arrayUpdated))
	assert.NoError(t, e)
	t.Log("Patch:", patch)
	assert.Equal(t, 1, len(patch), "they should be equal")
	sort.Sort(ByPath(patch))

	change := patch[0]
	assert.Equal(t, "add", change.Operation, "they should be equal")
	assert.Equal(t, "/persons/2", change.Path, "they should be equal")
	assert.Equal(t, map[string]interface{}{}, change.Value, "they should be equal")
	// change = patch[1]
	// assert.Equal(t, "add", change.Operation, "they should be equal")
	// assert.Equal(t, "/goods/2/batters/batter/2", change.Path, "they should be equal")
	// assert.Equal(t, map[string]interface{}{"id": "1003", "type": "Vanilla"}, change.Value, "they should be equal")
	// change = patch[2]
	// assert.Equal(t, change.Operation, "remove", "they should be equal")
	// assert.Equal(t, change.Path, "/goods/2/topping/2", "they should be equal")
	// assert.Equal(t, nil, change.Value, "they should be equal")
}
