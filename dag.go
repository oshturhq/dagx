package dagx

type DAG struct {
	adjacency map[string][]string
	reverse   map[string][]string
}

func NewDAG() *DAG {
	return &DAG{
		adjacency: make(map[string][]string),
		reverse:   make(map[string][]string),
	}
}

func (d *DAG) AddNode(id string) {
	if id == "" {
		return
	}
	if _, ok := d.adjacency[id]; !ok {
		d.adjacency[id] = []string{}
		d.reverse[id] = []string{}
	}
}

func (d *DAG) AddEdge(from, to string) {
	d.AddNode(from)
	d.AddNode(to)
	d.adjacency[from] = append(d.adjacency[from], to)
	d.reverse[to] = append(d.reverse[to], from)
}

func (d *DAG) Nodes() []string {
	ids := make([]string, 0, len(d.adjacency))
	for id := range d.adjacency {
		ids = append(ids, id)
	}
	return ids
}

func (d *DAG) Edges() []Edge {
	var edges []Edge
	for from, tos := range d.adjacency {
		for _, to := range tos {
			edges = append(edges, Edge{From: from, To: to})
		}
	}
	return edges
}

func (d *DAG) Parents(id string) []string {
	return d.reverse[id]
}

func (d *DAG) Children(id string) []string {
	return d.adjacency[id]
}

func (d *DAG) InDegree(id string) int {
	return len(d.reverse[id])
}

func (d *DAG) TopologicalOrder() ([]string, error) {
	if len(d.adjacency) == 0 {
		return nil, ErrEmptyGraph
	}

	inDegree := make(map[string]int)
	for id := range d.adjacency {
		inDegree[id] = 0
	}
	for _, tos := range d.adjacency {
		for _, to := range tos {
			inDegree[to]++
		}
	}

	queue := make([]string, 0)
	for id, deg := range inDegree {
		if deg == 0 {
			queue = append(queue, id)
		}
	}

	order := make([]string, 0, len(d.adjacency))
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		order = append(order, n)

		for _, child := range d.adjacency[n] {
			inDegree[child]--
			if inDegree[child] == 0 {
				queue = append(queue, child)
			}
		}
	}

	if len(order) != len(d.adjacency) {
		return nil, ErrCycleDetected
	}

	return order, nil
}

func (d *DAG) HasCycle() bool {
	_, err := d.TopologicalOrder()
	return err == ErrCycleDetected
}
