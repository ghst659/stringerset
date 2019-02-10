package stringerset

import (
	"testing"
)

func TestNewHas(t *testing.T) {
	elements := []string{"a", "b", "c"}
	ss := New("a", "b", "c")
	if !ss.HasAll("a", "b", "c") {
		t.Errorf("HasAll failed: %v", elements)
	}
	for _, e := range elements {
		if !ss.HasAny(e, "not_there") {
			t.Errorf("HasAny missing expected member: %s", e)
		}
	}
}

func TestMembers(t *testing.T) {
	ss := New("b", "a")
	got := ss.Members()
	want := []string{"a", "b"}
	if len(got) != len(want) {
		t.Errorf("Members() incorrect element count: got %d want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("Members() mismatched element: got %q want %q", got[i], want[i])
		}
	}
}

func TestSize(t *testing.T) {
	ss := New("p", "q", "r")
	if ss.Size() != 3 {
		t.Errorf("incorrect size: got %d want 3", ss.Size())
	}
}

func TestRemove(t *testing.T) {
	ss := New("a", "b", "c")
	ss.Remove("b")
	if ss.HasAny("b") {
		t.Errorf("failed Remove method")
	}
}

func TestUnion(t *testing.T) {
	x := New("a", "b")
	y := New("c")
	z := New("d")
	u := Union(x, y, z)
	if !u.HasAll("a", "b", "c", "d") {
		t.Errorf("Union failed")
	}
}

func TestIntersection(t *testing.T) {
	x := New("a", "b")
	y := New("c", "b")
	i := Intersection(x, y)
	if !i.HasAll("b") {
		t.Errorf("Intersection failed inclusion")
	}
	if i.HasAny("a", "c", "d") {
		t.Errorf("Intersection failed exclusion")
	}
}
