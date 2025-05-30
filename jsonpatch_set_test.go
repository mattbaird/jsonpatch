package jsonpatch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var simpleObjEmtpyPrmitiveSet = `{"a":100, "b":[]}`
var simpleObjPrimitiveSetWithOneItem = `{"a":100, "b":[1]}`
var simpleObjPrimitiveSetWithMultipleItems = `{"a":100, "b":[1,2]}`
var simpleObjAddSingleItemToPrimitiveSet = `{"b":[3]}`
var simpleObjAddMultipleItemsToPrimitiveSet = `{"b":[3,4]}`
var simpleObjAddDuplicateItemToPrimitiveSet = `{"b":[2]}`
var simpleObjSingletonObjectSet = `{"a":100, "b":[{"c":1}]}`
var simpleObjAddObjectSetItem = `{"b":[{"c":2}]}`
var simpleObjAddDuplicateObjectSetItem = `{"b":[{"c":1}]}`

var nestedObj = `{"a":100, "b":{"c":200}}`
var nestedObjModifyProp = `{"b":{"c":250}}`
var nestedObjAddProp = `{"b":{"d":"hello"}}`
var nestedObjPrimitiveSet = `{"a":100, "b":{"c":[200]}}`
var nestedObjAddPrimitiveSetItem = `{"b":{"c":[250]}}`

var setTestCollections = Collections{
	EntitySets: EntitySets{},
	Arrays:     []string{}, // No arrays in this test, only sets
}

func TestCreatePatch_AddItemToEmptyPrimitiveSetInEnsureExistsMode_GeneratesAddOperation(t *testing.T) {
	patch, err := CreatePatch([]byte(simpleObjEmtpyPrmitiveSet), []byte(simpleObjAddSingleItemToPrimitiveSet), setTestCollections, PatchStrategyEnsureExists)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(patch), "they should be equal")
	change := patch[0]
	assert.Equal(t, "add", change.Operation, "they should be equal")
	assert.Equal(t, "/b/0", change.Path, "they should be equal")
	var expected float64 = 3
	assert.Equal(t, expected, change.Value, "they should be equal")
}

func TestCreatePatch_AddItemToEmptyPrimitiveSetInEnsureExactMatchMode_GeneratesAddOperation(t *testing.T) {
	patch, err := CreatePatch([]byte(simpleObjEmtpyPrmitiveSet), []byte(simpleObjAddSingleItemToPrimitiveSet), setTestCollections, PatchStrategyExactMatch)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(patch), "they should be equal")
	change := patch[0]
	assert.Equal(t, "add", change.Operation, "they should be equal")
	assert.Equal(t, "/b/0", change.Path, "they should be equal")
	var expected float64 = 3
	assert.Equal(t, expected, change.Value, "they should be equal")
}

func TestCreatePatch_AddItemToPrimitiveSetWithOneItemInEnsureExistsMode_GeneratesAddOperation(t *testing.T) {
	patch, err := CreatePatch([]byte(simpleObjPrimitiveSetWithOneItem), []byte(simpleObjAddSingleItemToPrimitiveSet), setTestCollections, PatchStrategyEnsureExists)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(patch), "they should be equal")
	change := patch[0]
	assert.Equal(t, "add", change.Operation, "they should be equal")
	assert.Equal(t, "/b/1", change.Path, "they should be equal")
	var expected float64 = 3
	assert.Equal(t, expected, change.Value, "they should be equal")
}

func TestCreatePatch_AddItemToPrimitiveSetWithOneItemInExactMatchMode_GeneratesARemoveAndAnAddOperation(t *testing.T) {
	patch, err := CreatePatch([]byte(simpleObjPrimitiveSetWithOneItem), []byte(simpleObjAddSingleItemToPrimitiveSet), setTestCollections, PatchStrategyExactMatch)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(patch), "they should be equal")
	change := patch[0]
	assert.Equal(t, "remove", change.Operation, "they should be equal")
	assert.Equal(t, "/b/0", change.Path, "they should be equal")
	change = patch[1]
	assert.Equal(t, "add", change.Operation, "they should be equal")
	assert.Equal(t, "/b/0", change.Path, "they should be equal")
	var expected float64 = 3
	assert.Equal(t, expected, change.Value, "they should be equal")
}

func TestCreatePatch_AddItemToPrimitiveSetWithMultipleItems_InEnsureExistsMode_GeneratesAddOperation(t *testing.T) {
	patch, err := CreatePatch([]byte(simpleObjPrimitiveSetWithMultipleItems), []byte(simpleObjAddSingleItemToPrimitiveSet), setTestCollections, PatchStrategyEnsureExists)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(patch), "they should be equal")
	change := patch[0]
	assert.Equal(t, "add", change.Operation, "they should be equal")
	assert.Equal(t, "/b/2", change.Path, "they should be equal")
	var expected float64 = 3
	assert.Equal(t, expected, change.Value, "they should be equal")
}

func TestCreatePatch_AddItemToPrimitiveSetWithMultipleItems_InExactMatchMode_GeneratesTwoRemovesAndOneAddOperation(t *testing.T) {
	patch, err := CreatePatch([]byte(simpleObjPrimitiveSetWithMultipleItems), []byte(simpleObjAddSingleItemToPrimitiveSet), setTestCollections, PatchStrategyExactMatch)
	assert.NoError(t, err)
	assert.Equal(t, 3, len(patch), "they should be equal")
	change := patch[0]
	assert.Equal(t, "remove", change.Operation, "they should be equal")
	assert.Equal(t, "/b/1", change.Path, "they should be equal")
	change = patch[1]
	assert.Equal(t, "remove", change.Operation, "they should be equal")
	assert.Equal(t, "/b/0", change.Path, "they should be equal")
	change = patch[2]
	assert.Equal(t, "add", change.Operation, "they should be equal")
	assert.Equal(t, "/b/0", change.Path, "they should be equal")
	var expected float64 = 3
	assert.Equal(t, expected, change.Value, "they should be equal")
}

func TestCreatePatch_AddMultipleItemsToPrimitiveSetWithMultipleItems_InEnsureExistsMode_GeneratesTwoAddOperations(t *testing.T) {
	patch, err := CreatePatch([]byte(simpleObjPrimitiveSetWithMultipleItems), []byte(simpleObjAddMultipleItemsToPrimitiveSet), setTestCollections, PatchStrategyEnsureExists)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(patch), "they should be equal")
	change := patch[0]
	assert.Equal(t, "add", change.Operation, "they should be equal")
	assert.Equal(t, "/b/2", change.Path, "they should be equal")
	var expected float64 = 3
	assert.Equal(t, expected, change.Value, "they should be equal")
	change = patch[1]
	assert.Equal(t, "add", change.Operation, "they should be equal")
	assert.Equal(t, "/b/3", change.Path, "they should be equal")
	expected = 4
	assert.Equal(t, expected, change.Value, "they should be equal")
}

func TestCreatePatch_AddMultipleItemsToPrimitiveSetWithMultipleItems_InExactMatchMode_GeneratesTwoRemovesAndTwoAddOperations(t *testing.T) {
	patch, err := CreatePatch([]byte(simpleObjPrimitiveSetWithMultipleItems), []byte(simpleObjAddMultipleItemsToPrimitiveSet), setTestCollections, PatchStrategyExactMatch)
	assert.NoError(t, err)
	assert.Equal(t, 4, len(patch), "they should be equal")
	change := patch[0]
	assert.Equal(t, "remove", change.Operation, "they should be equal")
	assert.Equal(t, "/b/1", change.Path, "they should be equal")
	change = patch[1]
	assert.Equal(t, "remove", change.Operation, "they should be equal")
	assert.Equal(t, "/b/0", change.Path, "they should be equal")
	change = patch[2]
	assert.Equal(t, "add", change.Operation, "they should be equal")
	assert.Equal(t, "/b/0", change.Path, "they should be equal")
	var expected float64 = 3
	assert.Equal(t, expected, change.Value, "they should be equal")
	change = patch[3]
	assert.Equal(t, "add", change.Operation, "they should be equal")
	assert.Equal(t, "/b/1", change.Path, "they should be equal")
	expected = 4
	assert.Equal(t, expected, change.Value, "they should be equal")
}

func TestCreatePatch_AddDuplicateItemToPrimitiveSetWithOneMultipleItems_InEnsureExistsMode_GeneratesNoOperations(t *testing.T) {
	patch, err := CreatePatch([]byte(simpleObjPrimitiveSetWithMultipleItems), []byte(simpleObjAddDuplicateItemToPrimitiveSet), setTestCollections, PatchStrategyEnsureExists)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(patch), "they should be equal")
}

func TestCreatePatch_AddDuplicateItemToPrimitiveSetWithOneMultipleItems_InExactMatchMode_GeneratesARemoveOperation(t *testing.T) {
	patch, err := CreatePatch([]byte(simpleObjPrimitiveSetWithMultipleItems), []byte(simpleObjAddDuplicateItemToPrimitiveSet), setTestCollections, PatchStrategyExactMatch)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(patch), "they should be equal")
	change := patch[0]
	assert.Equal(t, "remove", change.Operation, "they should be equal")
	assert.Equal(t, "/b/0", change.Path, "they should be equal")
}

func TestCreatePatch_AddItemToPrimitiveSetInNestedObject_InEnsureExistsMode_GeneratesAddOperation(t *testing.T) {
	patch, err := CreatePatch([]byte(nestedObjPrimitiveSet), []byte(nestedObjAddPrimitiveSetItem), setTestCollections, PatchStrategyEnsureExists)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(patch), "they should be equal")
	change := patch[0]
	assert.Equal(t, "add", change.Operation, "they should be equal")
	assert.Equal(t, "/b/c/1", change.Path, "they should be equal")
	var expected float64 = 250
	assert.Equal(t, expected, change.Value, "they should be equal")
}

func TestCreatePatch_AddItemToPrimitiveSetInNestedObject_InExactMatchMode_GeneratesAddOperation(t *testing.T) {
	patch, err := CreatePatch([]byte(nestedObjPrimitiveSet), []byte(nestedObjAddPrimitiveSetItem), setTestCollections, PatchStrategyExactMatch)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(patch), "they should be equal")
	change := patch[0]
	assert.Equal(t, "remove", change.Operation, "they should be equal")
	assert.Equal(t, "/b/c/0", change.Path, "they should be equal")
	change = patch[1]
	assert.Equal(t, "add", change.Operation, "they should be equal")
	assert.Equal(t, "/b/c/0", change.Path, "they should be equal")
	var expected float64 = 250
	assert.Equal(t, expected, change.Value, "they should be equal")
}

func TestCreatePatch_AddItemToObjectSetWithOneItem_InEnsureExistsMode_GeneratesAddOperation(t *testing.T) {
	patch, err := CreatePatch([]byte(simpleObjSingletonObjectSet), []byte(simpleObjAddObjectSetItem), setTestCollections, PatchStrategyEnsureExists)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(patch), "they should be equal")
	change := patch[0]
	assert.Equal(t, "add", change.Operation, "they should be equal")
	assert.Equal(t, "/b/1", change.Path, "they should be equal")
	var expected = map[string]any{"c": float64(2)}
	assert.Equal(t, expected, change.Value, "they should be equal")
}

func TestCreatePatch_AddItemToObjectSetWithOneItem_InExactMatchMode_GeneratesAddOperation(t *testing.T) {
	patch, err := CreatePatch([]byte(simpleObjSingletonObjectSet), []byte(simpleObjAddObjectSetItem), setTestCollections, PatchStrategyExactMatch)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(patch), "they should be equal")
	change := patch[0]
	assert.Equal(t, "remove", change.Operation, "they should be equal")
	assert.Equal(t, "/b/0", change.Path, "they should be equal")
	change = patch[1]
	assert.Equal(t, "add", change.Operation, "they should be equal")
	assert.Equal(t, "/b/0", change.Path, "they should be equal")
	var expected = map[string]any{"c": float64(2)}
	assert.Equal(t, expected, change.Value, "they should be equal")
}

func TestCreatePatch_AddDuplicateItemToObjectSetWithOneItem_InEnsureExistsMode_GeneratesNoOperations(t *testing.T) {
	patch, err := CreatePatch([]byte(simpleObjSingletonObjectSet), []byte(simpleObjAddDuplicateObjectSetItem), setTestCollections, PatchStrategyEnsureExists)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(patch), "they should be equal")
}

func TestCreatePatch_AddDuplicateItemToObjectSetWithOneItem_InExactMatchMode_GeneratesNoOperations(t *testing.T) {
	patch, err := CreatePatch([]byte(simpleObjSingletonObjectSet), []byte(simpleObjAddDuplicateObjectSetItem), setTestCollections, PatchStrategyExactMatch)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(patch), "they should be equal")
}
