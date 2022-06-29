package algorithms

import (
	dse "dsa/general"
	"fmt"
)

var stat bool
var msg string

func initvar() {
	stat = false
	msg = ""
}

type Search interface {
	LSearch() (string, error)
}

type Ilist struct {
	List    []int
	Element int
}

type Slist struct {
	List    []string
	Element string
}

func (il Ilist) LSearch() (string, error) {
	initvar()
	var errmsg string = "%d does not found in given list."
	for i, e := range il.List {
		if e == il.Element {
			stat = true
			msg = fmt.Sprintf("%d found at location %d", il.Element, i)
			break
		}
	}
	if !stat {
		return msg, &dse.DsaError{Err: fmt.Sprintf(errmsg, il.Element)}
	}
	return msg, nil
}

func (sl Slist) LSearch() (string, error) {
	initvar()
	var errmsg string = "%s does not found in given list."
	for i, e := range sl.List {
		if e == sl.Element {
			stat = true
			msg = fmt.Sprintf("%s found at location %d", sl.Element, i)
			break
		}
	}
	if !stat {
		return msg, &dse.DsaError{Err: fmt.Sprintf(errmsg, sl.Element)}
	}
	return msg, nil
}

func LinearSearch(s Search) (string, error) {
	return s.LSearch()
}
