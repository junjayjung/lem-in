package main

type Matrix struct {
	StartRoom string
	EndRoom   string
	Edges     map[string][]string
	Vertices  []string
}

type Dot struct {
	NumAnts int
}

type PathData struct {
	AllPaths        [][]string
	NonIntercepting [][][]string
	SortedComb      map[int][][]string
	BestComb        [][]string
}
