package jsonpatch

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	jsonpatch "github.com/evanphx/json-patch"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJSONPatchCreate(t *testing.T) {
	cases := map[string]struct {
		a string
		b string
	}{
		"object": {
			`{"asdf":"qwerty"}`,
			`{"asdf":"zzz"}`,
		},
		"object with array": {
			`{"items":[{"asdf":"qwerty"}]}`,
			`{"items":[{"asdf":"bla"},{"asdf":"zzz"}]}`,
		},
		"array": {
			`[{"asdf":"qwerty"}]`,
			`[{"asdf":"bla"},{"asdf":"zzz"}]`,
		},
		"from empty array": {
			`[]`,
			`[{"asdf":"bla"},{"asdf":"zzz"}]`,
		},
		"to empty array": {
			`[{"asdf":"bla"},{"asdf":"zzz"}]`,
			`[]`,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			patch, err := CreatePatch([]byte(tc.a), []byte(tc.b))
			assert.NoError(t, err)

			patchBytes, err := json.Marshal(patch)
			assert.NoError(t, err)

			fmt.Printf("%s\n", string(patchBytes))

			p, err := jsonpatch.DecodePatch(patchBytes)
			assert.NoError(t, err)

			res, err := p.Apply([]byte(tc.a))
			assert.NoError(t, err)
			spew.Dump(res)

			assert.Equal(t, tc.b, string(res))
		})
	}
}