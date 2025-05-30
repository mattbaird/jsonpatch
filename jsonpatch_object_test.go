package jsonpatch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var simpleObj = `{"a":100, "b":20}`
var simpleObjModifyProp = `{"b":250}`
var simpleObjAddProp = `{"c":"hello"}`

func TestCreatePatch_ModifyProperty_GeneratesReplaceOperation(t *testing.T) {
	patch, err := CreatePatch([]byte(simpleObj), []byte(simpleObjModifyProp), setTestCollections, PatchStrategyEnsureExists)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(patch), "they should be equal")
	change := patch[0]
	assert.Equal(t, "replace", change.Operation, "they should be equal")
	assert.Equal(t, "/b", change.Path, "they should be equal")
	var expected float64 = 250
	assert.Equal(t, expected, change.Value, "they should be equal")
}

func TestCreatePatch_AddProperty_GeneratesAddOperation(t *testing.T) {
	patch, err := CreatePatch([]byte(simpleObj), []byte(simpleObjAddProp), setTestCollections, PatchStrategyEnsureExists)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(patch), "they should be equal")
	change := patch[0]
	assert.Equal(t, "add", change.Operation, "they should be equal")
	assert.Equal(t, "/c", change.Path, "they should be equal")
	assert.Equal(t, "hello", change.Value, "they should be equal")
}

func TestCreatePatch_NestedObject_ModifyProperty_GeneratesReplaceOperation(t *testing.T) {
	patch, err := CreatePatch([]byte(nestedObj), []byte(nestedObjModifyProp), setTestCollections, PatchStrategyEnsureExists)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(patch), "they should be equal")
	change := patch[0]
	assert.Equal(t, "replace", change.Operation, "they should be equal")
	assert.Equal(t, "/b/c", change.Path, "they should be equal")
	var expected float64 = 250
	assert.Equal(t, expected, change.Value, "they should be equal")
}

func TestCreatePatch_NestedObject_AddProperty_GeneratesAddOperation(t *testing.T) {
	patch, err := CreatePatch([]byte(nestedObj), []byte(nestedObjAddProp), setTestCollections, PatchStrategyEnsureExists)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(patch), "they should be equal")
	change := patch[0]
	assert.Equal(t, "add", change.Operation, "they should be equal")
	assert.Equal(t, "/b/d", change.Path, "they should be equal")
	assert.Equal(t, "hello", change.Value, "they should be equal")
}
