package graph

import (
	"math/rand"
	"testing"
	"time"
)

const (
	maxGraphSize = 100
)

func TestUngraphHasVertices(t *testing.T) {
	for i := 0; i < maxGraphSize; i = (i + 1) * 10 {
		g := NewGraph(i)
		if g.V() != i {
			t.Errorf("Graph doesn't have as many vertices as given for "+
				"constructor, expected %d but was %d", i, g.V())
		}
	}
}

func TestUngraphHasEdges(t *testing.T) {

	for i := 10; i < maxGraphSize; i = (i + 1) * 10 {
		g := NewGraph(i)

		expect := (i - 1) / 2

		for j := 0; j < i-2; j += 2 {
			g.AddEdge(j, j+1)
		}

		actual := g.E()
		if expect != actual {
			t.Errorf("Expected %d edges but was %d", expect, actual)
		}
	}
}

func TestUngraphCanStringify(t *testing.T) {
	for i := 2; i < maxGraphSize; i = (i + 1) * 10 {
		g := NewGraph(i)
		g.AddEdge(0, i-1)
		g.GoString()
	}
}

func TestUngraphCanAddEdges(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	rInt := func(from, to int) int {
		if from <= 0 {
			from = 0
		}
		if to <= 0 {
			return 0
		}
		return from + r.Intn(to-from)
	}

	for i := 1; i < maxGraphSize; i = (i + 1) * 10 {
		g := NewGraph(i)

		for v := 0; v < g.V(); v++ {
			adj := g.Adj(v)
			if len(adj) != 0 {
				t.Fatalf("Graph should have no edges at this point, "+
					"had at least %d", len(adj))
			}
		}
		if i < 2 {
			continue
		}

		from := rInt(0, i/2)
		to := rInt(i/2, i-1)

		g.AddEdge(from, to)
		adj := g.Adj(from)
		if len(adj) != 1 {
			t.Fatalf("Vertex %d should have only 1 edge", from)
		}

		if adj[0] != to {
			t.Fatalf("Vertex %d should have an edge with %d", from, to)
		}
	}
}
