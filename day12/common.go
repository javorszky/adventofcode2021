package day12

import (
	"io/ioutil"
	"strings"

	"github.com/javorszky/adventofcode2021/util"
)

const (
	startName = "start"
	endName   = "end"
)

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs(fn string) []string {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.TrimRight(string(data), util.NewLine), util.NewLine)
}

type node struct {
	name string
	smol bool
	lynx map[string]*node
}

func (n *node) Link(otherNode *node) {
	if _, ok := n.lynx[otherNode.name]; !ok {
		n.lynx[otherNode.name] = otherNode
	}
}

func parseIntoNodes(input []string) map[string]*node {
	m := make(map[string]*node)

	for _, line := range input {
		parts := strings.Split(line, "-")
		if _, ok := m[parts[0]]; !ok {
			m[parts[0]] = newNode(parts[0])
		}

		if _, ok := m[parts[1]]; !ok {
			m[parts[1]] = newNode(parts[1])
		}

		m[parts[0]].Link(m[parts[1]])
		m[parts[1]].Link(m[parts[0]])
	}

	return m
}

func newNode(name string) *node {
	l := make(map[string]*node)

	return &node{
		name: name,
		smol: strings.ToLower(name) == name,
		lynx: l,
	}
}
