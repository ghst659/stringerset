// Package stringerset implements a set of Stringers.
package stringerset

import (
	"fmt"
)

type StringerSet map[string]struct{}

// toString returns the string representation of an interface.
func toString(x interface{}) string {
	switch v := x.(type) {
	case string:
		return v
	case int:
		return fmt.Sprintf("%d", v)
	case float64:
		return fmt.Sprintf("%f", v)
	default:
		return fmt.Sprintf("%+v", v)
	}
}

// New creates a new set initialised with zero or more members.
func New(members ...interface{}) StringerSet {
	ss := make(StringerSet)
	ss.Add(members...)
	return ss
}

// Add adds the given Stringers to the set.
func (ss StringerSet) Add(members ...interface{}) {
	for _, m := range members {
		ss[toString(m)] = struct{}{}
	}
}

// Remove removes the given members from the set, whether or not it is present.
func (ss StringerSet) Remove(members ...interface{}) {
	for _, m := range members {
		delete(ss, toString(m))
	}
}

// HasAll returns true if all of the given args are members of the set.
func (ss StringerSet) HasAll(query ...interface{}) bool {
	for _, q := range query {
		if _, present := ss[toString(q)]; !present {
			return false
		}
	}
	return true
}

// HasAny returns true if all of the given args are members of the set.
func (ss StringerSet) HasAny(query ...interface{}) bool {
	for _, q := range query {
		if _, present := ss[toString(q)]; present {
			return true
		}
	}
	return false
}

// Union returns a StringerSet that has the members of all the args.
func Union(ssArgs ...StringerSet) StringerSet {
	u := New()
	for _, ss := range ssArgs {
		for m, _ := range ss {
			u[m] = struct{}{}
		}
	}
	return u
}

// Intersect returns a StringerSet that has the intersection of all args.
func Intersection(ssArgs ...StringerSet) StringerSet {
	i := New()
	for order, ss := range ssArgs {
		if order == 0 {
			for m, _ := range ss {
				i[m] = struct{}{}
			}
		} else {
			for m, _ := range i {
				if _, present := ss[m]; !present {
					delete(i, m)
				}
			}
		}
	}
	return i
}
