package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func Solve(r *Scanner, w *Writer) {
	N, M := r.Int(), r.Int()
	results := make([]int, M)
	for i := 0; i < N; i++ {
		K := r.Int()
		for k := 0; k < K; k++ {
			results[r.Int()-1]++
		}
	}
	ans := 0
	for _, v := range results {
		if v == N {
			ans++
		}
	}
	w.Println(ans)
}

func main() {
	r, w := NewScanner(os.Stdin, MaxBufferSize), NewWriter(os.Stdout)
	defer w.Flush()
	Solve(r, w)
}

const MaxBufferSize = 1 * 1024 * 1024

type Scanner struct{ sc *bufio.Scanner }

func NewScanner(r io.Reader, size int) *Scanner {
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, size), size)
	sc.Split(bufio.ScanWords)
	return &Scanner{sc}
}
func (s *Scanner) scan() bool       { return s.sc.Scan() }
func (s *Scanner) text() string     { return s.sc.Text() }
func (s *Scanner) String() string   { s.scan(); return s.text() }
func (s *Scanner) Int() int         { return atoi(s.String()) }
func (s *Scanner) Float64() float64 { return atof(s.String()) }

type Writer struct{ bf *bufio.Writer }

func NewWriter(w io.Writer) *Writer        { return &Writer{bufio.NewWriter(w)} }
func (w *Writer) Print(a ...interface{})   { fmt.Fprint(w.bf, a...) }
func (w *Writer) Println(a ...interface{}) { fmt.Fprintln(w.bf, a...) }
func (w *Writer) Flush()                   { w.bf.Flush() }

// unwrap returns the value of v if err is nil and panics otherwise.
func unwrap[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

// atoi returns an integer converted from s.
func atoi(s string) int { return unwrap(strconv.Atoi(s)) }

// atof returns a float converted from s.
func atof(s string) float64 { return unwrap(strconv.ParseFloat(s, 64)) }

// UnionFind is a disjoint-set data structure.
type UnionFind struct {
	parent []int // parent[i] = parent of i
	size   []int // size[i] = number of elements in subtree rooted at i
}

// NewUnionFind creates a new union-find data structure with n elements.
func NewUnionFind(n int) *UnionFind {
	u := &UnionFind{
		parent: make([]int, n),
		size:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		u.parent[i] = -1
		u.size[i] = 1
	}
	return u
}

// Root returns the root of the component that element x belongs to.
func (u *UnionFind) Root(x int) int {
	// x is the root of the tree
	if u.parent[x] == -1 {
		return x
	}

	// Use path compression heuristic.
	u.parent[x] = u.Root(u.parent[x])
	return u.parent[x]
}

// IsSameSet returns true if elements x and y belong to the same component.
func (u *UnionFind) IsSameSet(x, y int) bool {
	return u.Root(x) == u.Root(y)
}

// Union merges the components that elements x and y belong to.
func (u *UnionFind) Union(x, y int) bool {
	xRoot, yRoot := u.Root(x), u.Root(y)
	if xRoot == yRoot {
		return false
	}

	// Use union by size heuristic.
	// Merge smaller component into the larger one.
	if u.size[xRoot] < u.size[yRoot] {
		xRoot, yRoot = yRoot, xRoot
	}
	u.parent[yRoot] = xRoot
	u.size[xRoot] += u.size[yRoot]
	return true
}

// Size returns the size of the component that element x belongs to.
func (u *UnionFind) Size(x int) int {
	return u.size[u.Root(x)]
}
