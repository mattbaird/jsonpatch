package jsonpatch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var simpleObj = `{"a":100, "b":20}`
var simpleObjModifyProp = `{"b":250}`
var simpleObjAddProp = `{"c":"hello"}`
var simpleObjEmtpyPrmitiveArray = `{"a":100, "b":[]}`
var simpleObjSingletonPrimitiveArray = `{"a":100, "b":[1]}`
var simpleObjMultipleItemPrimitiveArray = `{"a":100, "b":[1,2]}`
var simpleObjAddPrimitiveArrayItem = `{"b":[3]}`
var simpleObjAddDuplicateArrayItem = `{"b":[2]}`
var simpleObjSingletonObjectArray = `{"a":100, "b":[{"c":1}]}`
var simpleObjAddObjectArrayItem = `{"b":[{"c":2}]}`
var simpleObjAddDuplicateObjectArrayItem = `{"b":[{"c":1}]}`
var simpleObjKeyValueArray = `{"a":100, "t":[{"k":1, "v":1},{"k":2, "v":2}]}`
var simpleObjAddKeyValueArrayItem = `{"t":[{"k":3, "v":3}]}`
var simpleObjModifyKeyValueArrayItem = `{"t":[{"k":2, "v":3}]}`
var simpleObjAddDuplicateKeyValueArrayItem = `{"t":[{"k":2, "v":2}]}`
var complexNextedKeyValueArray = `{
    "a":100,
    "t":[
    {"k":1,
    "v":[
    {"nk":11, "c":"x", "d":[1,2]},
    {"nk":22, "c":"y", "d":[3,4]}
    ]
    },
    {"k":2,
    "v":[
    {"nk":33, "c":"z", "d":[5,6]}
    ]
    }
    ]}`
var complexNextedKeyValueArrayModifyItem = `{
    "t":[
    {"k":2,
    "v":[
    {"nk":33, "c":"zz", "d":[7,8]}
    ]
    }
    ]}`

var nestedObj = `{"a":100, "b":{"c":200}}`
var nestedObjModifyProp = `{"b":{"c":250}}`
var nestedObjAddProp = `{"b":{"d":"hello"}}`
var nestedObjPrimitiveArray = `{"a":100, "b":{"c":[200]}}`
var nestedObjAddPrimitiveArrayItem = `{"b":{"c":[250]}}`

var ensureExistsStrategyTestCollections = Collections{
	entitySets: EntitySets{
		Path("$.t"):      Key("k"),
		Path("$.t[*].v"): Key("nk"),
	},
	arrays: []string{}, // No arrays in this test, only sets
}

func TestCreatePatch_ModifyProperty_GeneratesReplaceOperation(t *testing.T) {
	patch, err := CreatePatch([]byte(simpleObj), []byte(simpleObjModifyProp), ensureExistsStrategyTestCollections, PatchStrategyEnsureExists)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(patch), "they should be equal")
	change := patch[0]
	assert.Equal(t, "replace", change.Operation, "they should be equal")
	assert.Equal(t, "/b", change.Path, "they should be equal")
	var expected float64 = 250
	assert.Equal(t, expected, change.Value, "they should be equal")
}

func TestCreatePatch_AddProperty_GeneratesAddOperation(t *testing.T) {
	patch, err := CreatePatch([]byte(simpleObj), []byte(simpleObjAddProp), ensureExistsStrategyTestCollections, PatchStrategyEnsureExists)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(patch), "they should be equal")
	change := patch[0]
	assert.Equal(t, "add", change.Operation, "they should be equal")
	assert.Equal(t, "/c", change.Path, "they should be equal")
	assert.Equal(t, "hello", change.Value, "they should be equal")
}

func TestCreatePatch_NestedObject_ModifyProperty_GeneratesReplaceOperation(t *testing.T) {
	patch, err := CreatePatch([]byte(nestedObj), []byte(nestedObjModifyProp), ensureExistsStrategyTestCollections, PatchStrategyEnsureExists)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(patch), "they should be equal")
	change := patch[0]
	assert.Equal(t, "replace", change.Operation, "they should be equal")
	assert.Equal(t, "/b/c", change.Path, "they should be equal")
	var expected float64 = 250
	assert.Equal(t, expected, change.Value, "they should be equal")
}

func TestCreatePatch_NestedObject_AddProperty_GeneratesAddOperation(t *testing.T) {
	patch, err := CreatePatch([]byte(nestedObj), []byte(nestedObjAddProp), ensureExistsStrategyTestCollections, PatchStrategyEnsureExists)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(patch), "they should be equal")
	change := patch[0]
	assert.Equal(t, "add", change.Operation, "they should be equal")
	assert.Equal(t, "/b/d", change.Path, "they should be equal")
	assert.Equal(t, "hello", change.Value, "they should be equal")
}

func TestCreatePatch_EmptyPrimitiveArray_AddItem_GeneratesAddOperation(t *testing.T) {
	patch, err := CreatePatch([]byte(simpleObjEmtpyPrmitiveArray), []byte(simpleObjAddPrimitiveArrayItem), ensureExistsStrategyTestCollections, PatchStrategyEnsureExists)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(patch), "they should be equal")
	change := patch[0]
	assert.Equal(t, "add", change.Operation, "they should be equal")
	assert.Equal(t, "/b/0", change.Path, "they should be equal")
	var expected float64 = 3
	assert.Equal(t, expected, change.Value, "they should be equal")
}

func TestCreatePatch_SingletonPrimitiveArray_AddItem_GeneratesAddOperation(t *testing.T) {
	patch, err := CreatePatch([]byte(simpleObjSingletonPrimitiveArray), []byte(simpleObjAddPrimitiveArrayItem), ensureExistsStrategyTestCollections, PatchStrategyEnsureExists)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(patch), "they should be equal")
	change := patch[0]
	assert.Equal(t, "add", change.Operation, "they should be equal")
	assert.Equal(t, "/b/1", change.Path, "they should be equal")
	var expected float64 = 3
	assert.Equal(t, expected, change.Value, "they should be equal")
}

func TestCreatePatch_MultipleItemPrimitiveArray_AddItem_GeneratesAddOperation(t *testing.T) {
	patch, err := CreatePatch([]byte(simpleObjMultipleItemPrimitiveArray), []byte(simpleObjAddPrimitiveArrayItem), ensureExistsStrategyTestCollections, PatchStrategyEnsureExists)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(patch), "they should be equal")
	change := patch[0]
	assert.Equal(t, "add", change.Operation, "they should be equal")
	assert.Equal(t, "/b/2", change.Path, "they should be equal")
	var expected float64 = 3
	assert.Equal(t, expected, change.Value, "they should be equal")
}

func TestCreatePatch_SingletonPrimitiveArray_AddDuplicateItem_GeneratesNoOperations(t *testing.T) {
	patch, err := CreatePatch([]byte(simpleObjMultipleItemPrimitiveArray), []byte(simpleObjAddDuplicateArrayItem), ensureExistsStrategyTestCollections, PatchStrategyEnsureExists)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(patch), "they should be equal")
}

func TestCreatePatch_NestedObject_PrimitiveArray_AddItem_GeneratesAddOperation(t *testing.T) {
	patch, err := CreatePatch([]byte(nestedObjPrimitiveArray), []byte(nestedObjAddPrimitiveArrayItem), ensureExistsStrategyTestCollections, PatchStrategyEnsureExists)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(patch), "they should be equal")
	change := patch[0]
	assert.Equal(t, "add", change.Operation, "they should be equal")
	assert.Equal(t, "/b/c/1", change.Path, "they should be equal")
	var expected float64 = 250
	assert.Equal(t, expected, change.Value, "they should be equal")
}

func TestCreatePatch_SingletonObjectArray_AddItem_GeneratesAddOperation(t *testing.T) {
	patch, err := CreatePatch([]byte(simpleObjSingletonObjectArray), []byte(simpleObjAddObjectArrayItem), ensureExistsStrategyTestCollections, PatchStrategyEnsureExists)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(patch), "they should be equal")
	change := patch[0]
	assert.Equal(t, "add", change.Operation, "they should be equal")
	assert.Equal(t, "/b/1", change.Path, "they should be equal")
	var expected = map[string]any{"c": float64(2)}
	assert.Equal(t, expected, change.Value, "they should be equal")
}

func TestCreatePatch_KeyValueArray_AddItem_GeneratesAddOperation(t *testing.T) {
	patch, err := CreatePatch([]byte(simpleObjKeyValueArray), []byte(simpleObjAddKeyValueArrayItem), ensureExistsStrategyTestCollections, PatchStrategyEnsureExists)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(patch), "they should be equal")
	change := patch[0]
	assert.Equal(t, "add", change.Operation, "they should be equal")
	assert.Equal(t, "/t/2", change.Path, "they should be equal")
	var expected = map[string]any{"k": float64(3), "v": float64(3)}
	assert.Equal(t, expected, change.Value, "they should be equal")
}

func TestCreatePatch_SingletonObjectArray_AddDuplicateItem_GeneratesNoOperations(t *testing.T) {
	patch, err := CreatePatch([]byte(simpleObjSingletonObjectArray), []byte(simpleObjAddDuplicateObjectArrayItem), ensureExistsStrategyTestCollections, PatchStrategyEnsureExists)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(patch), "they should be equal")
}

func TestCreatePatch_KeyValueArray_ModifyItem_GeneratesReplaceOperation(t *testing.T) {
	patch, err := CreatePatch([]byte(simpleObjKeyValueArray), []byte(simpleObjModifyKeyValueArrayItem), ensureExistsStrategyTestCollections, PatchStrategyEnsureExists)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(patch), "they should be equal")
	change := patch[0]
	assert.Equal(t, "replace", change.Operation, "they should be equal")
	assert.Equal(t, "/t/1/v", change.Path, "they should be equal")
	var expected float64 = 3
	assert.Equal(t, expected, change.Value, "they should be equal")
}

func TestCreatePatch_KeyValueArray_AddDuplicateItem_GeneratesNoOperations(t *testing.T) {
	patch, err := CreatePatch([]byte(simpleObjKeyValueArray), []byte(simpleObjAddDuplicateKeyValueArrayItem), ensureExistsStrategyTestCollections, PatchStrategyEnsureExists)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(patch), "they should be equal")
}

func TestCreatePatch_ComplexNestedKeyValueArray_ModifyItem_GeneratesReplaceOperation(t *testing.T) {
	patch, err := CreatePatch([]byte(complexNextedKeyValueArray), []byte(complexNextedKeyValueArrayModifyItem), ensureExistsStrategyTestCollections, PatchStrategyEnsureExists)
	assert.NoError(t, err)
	assert.Equal(t, 3, len(patch), "they should be equal")
	change := patch[0]
	assert.Equal(t, "replace", change.Operation, "they should be equal")
	assert.Equal(t, "/t/1/v/0/c", change.Path, "they should be equal")
	assert.Equal(t, "zz", change.Value, "they should be equal")
	change = patch[1]
	assert.Equal(t, "add", change.Operation, "they should be equal")
	assert.Equal(t, "/t/1/v/0/d/2", change.Path, "they should be equal")
	assert.Equal(t, float64(7), change.Value, "they should be equal")
	change = patch[2]
	assert.Equal(t, "add", change.Operation, "they should be equal")
	assert.Equal(t, "add", change.Operation, "they should be equal")
	assert.Equal(t, "/t/1/v/0/d/3", change.Path, "they should be equal")
	assert.Equal(t, float64(8), change.Value, "they should be equal")
}
