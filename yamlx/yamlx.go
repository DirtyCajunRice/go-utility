package yamlx

import (
	"errors"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

var anchors = make(map[string]*yaml.Node)

// Document is an intermediary struct for holding included files
type Document struct {
	content *yaml.Node
}

// UnmarshalYAML saves the entire included yaml file as its parent yaml.node
func (d *Document) UnmarshalYAML(value *yaml.Node) error {
	var err error
	d.content, err = resolve(value)
	return err
}

// Processor is a struct that holds the intended struct for unmarshalling
type Processor struct {
	target interface{}
}

func (i *Processor) setTarget(target interface{}) {
	i.target = target
}

// UnmarshalYAML saves the entire yaml file as its parent yaml.node
func (i *Processor) UnmarshalYAML(value *yaml.Node) error {
	r, err := resolve(value)
	if err != nil {
		return err
	}
	return r.Decode(i.target)
}

// NewProcessor accepts a target struct for unmarshal, and returns a Processor
func NewProcessor(target interface{}) *Processor {
	return &Processor{target: target}
}

func resolve(node *yaml.Node) (*yaml.Node, error) {
	if node.Anchor != "" {
		anchors[node.Anchor] = node
	}
	if node.Tag == "!include" {
		if node.Kind != yaml.ScalarNode {
			return nil, errors.New("!include on a non-scalar node")
		}
		file, err := ioutil.ReadFile(node.Value)
		if err != nil {
			return nil, err
		}
		var f Document
		err = yaml.Unmarshal(file, &f)
		return f.content, err
	}
	if node.Tag == "!ref" {
		node.Kind = yaml.AliasNode
		node.Alias = anchors[node.Value]
		if anchors[node.Value] == nil {
			return nil, errors.New("could not find anchor")
		}
		return node, nil
	}
	if node.Kind == yaml.SequenceNode || node.Kind == yaml.MappingNode {
		var err error
		for i := range node.Content {
			node.Content[i], err = resolve(node.Content[i])
			if err != nil {
				return nil, err
			}
		}
	}
	return node, nil
}

func Unmarshal(in []byte, out interface{}) (err error) {
	return yaml.Unmarshal(in, out)
}
