package day14

import (
	"strings"
)

type polymerElement struct {
	point uint
	next  *polymerElement
}

func (p *polymerElement) Point() uint {
	return p.point
}

func (p *polymerElement) SetNext(el *polymerElement) {
	p.next = el
}

func (p *polymerElement) Next() *polymerElement {
	return p.next
}

func (p *polymerElement) String() string {
	return codePointToLetter[p.point]
}

func newPolymerElement(s uint) *polymerElement {
	return &polymerElement{
		point: s,
		next:  nil,
	}
}

func (p *polymerElement) InsertLink(newElement *polymerElement) *polymerElement {
	newElement.SetNext(p.next)
	p.SetNext(newElement)

	return newElement
}

func task1LinkedList(template string, rules []string) int {
	betterRules := parseBetterRules(rules)
	polymer := parseTemplateLinkedList(template)

	for j := 0; j < 10; j++ {
		workLinkedList(polymer, betterRules)
	}

	counts := map[uint]int{}
	walker := polymer
	most := 0
	least := 0

	for {
		counts[walker.Point()]++
		walker = walker.Next()

		if walker == nil {
			break
		}

		least++
	}

	for _, v := range counts {
		if v > most {
			most = v
		}

		if v < least {
			least = v
		}
	}

	return most - least
}

func parseTemplateLinkedList(template string) *polymerElement {
	first := newPolymerElement(uint(template[0]))
	prev := first

	for i := 1; i < len(template); i++ {
		el := newPolymerElement(uint(template[i]))
		prev.SetNext(el)
		prev = el
	}

	return first
}

func workLinkedList(start *polymerElement, betterRules map[uint]uint) *polymerElement {
	walker := start

	for {
		el := newPolymerElement(betterRules[walker.Point()<<8|walker.Next().Point()])
		el.SetNext(walker.Next())
		walker.SetNext(el)
		walker = el.Next()

		if walker.Next() == nil {
			break
		}
	}

	return start
}

func drain(start *polymerElement) string {
	var sb strings.Builder

	for {
		sb.WriteString(start.String())
		start = start.Next()

		if start == nil {
			break
		}
	}

	return sb.String()
}
