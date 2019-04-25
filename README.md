# jsonpatch

[![Build Status][build-image]][build-url]


[build-url]: https://travis-ci.com/benitogf/jsonpatch
[build-image]: https://api.travis-ci.com/benitogf/jsonpatch.svg?branch=master&style=flat-square

As per http://jsonpatch.com/ JSON Patch is specified in RFC 6902 from the IETF.

an attempt to join the efforts made by [mattbaird](https://github.com/mattbaird/jsonpatch) and [evanphx](https://github.com/evanphx/json-patch)

using [jsondiff](https://github.com/nsf/jsondiff) on the tests


## Create Patch

```go
package main

import (
	"fmt"
	"github.com/benitogf/jsonpatch"
)

var simpleA = `{"a":100, "b":200, "c":"hello"}`
var simpleB = `{"a":100, "b":200, "c":"goodbye"}`

func main() {
	patch, e := jsonpatch.CreatePatch([]byte(simpleA), []byte(simpleA))
	if e != nil {
		fmt.Printf("Error creating JSON patch:%v", e)
		return
	}
	for _, operation := range patch {
		fmt.Printf("%s\n", operation.Json())
	}
}
```


## Apply Patch

```go
package main

import (
	"fmt"
	"github.com/benitogf/jsonpatch"
)

func main() {
	original := []byte(`{"name": "John", "age": 24, "height": 3.21}`)
	patchJSON := []byte(`[
		{"op": "replace", "path": "/name", "value": "Jane"},
		{"op": "remove", "path": "/height"}
	]`)

	patch, err := jsonpatch.DecodePatch(patchJSON)
	if err != nil {
		panic(err)
	}

	modified, err := patch.Apply(original)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Original document: %s\n", original)
	fmt.Printf("Modified document: %s\n", modified)
}
```