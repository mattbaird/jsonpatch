package jsonpatch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalNullableValue(t *testing.T) {
	p1 := Operation{
		Operation: "replace",
		Path:      "/a1",
		Value:     nil,
	}
	assert.JSONEq(t, `{"op":"replace", "path":"/a1","value":null}`, p1.JSON())

	p2 := Operation{
		Operation: "replace",
		Path:      "/a2",
		Value:     "v2",
	}
	assert.JSONEq(t, `{"op":"replace", "path":"/a2", "value":"v2"}`, p2.JSON())
}

func TestMarshalNonNullableValue(t *testing.T) {
	p1 := Operation{
		Operation: "remove",
		Path:      "/a1",
	}
	assert.JSONEq(t, `{"op":"remove", "path":"/a1"}`, p1.JSON())

}
