# Extensions for go-yaml/yaml

Introduction
------------

The yamlx package extends [go-yaml/yaml.v3](gopkg.in/yaml.v3).
It is initially to add custom tag support for importing files and referencing imported anchors

Installation and usage
----------------------

The import path for the package is *gopkg.in/DirtyCajunRice/go-utility/yamlx.v0*.

To install it, run:

    go get gopkg.in/DirtyCajunRice/go-utility/yamlx.v0

API documentation
-----------------

If opened in a browser, the import path itself leads to the API documentation:

- [https://gopkg.in/DirtyCajunRice/go-utility/yamlx.v0](https://gopkg.in/DirtyCajunRice/go-utility/yamlx.v0)

Example
-------

```go
package main

import (
	"io/ioutil"

	"gopkg.in/DirtyCajunRice/go-utility/yamlx.v0"
	"gopkg.in/yaml.v3"
)

type Example struct {
	FieldA string `yaml:"field_a"`
	FieldB string `yaml:"field_b"`
}

var data1 = `
a: !ref a
b: &b
  c: 2
  d: [3, 4]
`

var data2 = `
field_a: &a Easy!
.some_key: !include data1.yaml
field_b: 
  <<: !ref b 
  e: 5
`

func main() {
	var example Example

	err := ioutil.WriteFile("data1.yaml", []byte(data1), 0644)
	if err != nil {
		panic(err)
	}
	
	processor := yamlx.NewProcessor(example)
	err = yaml.Unmarshal([]byte(data2), processor)
	if err != nil {
		panic(err)
	}
}
```
