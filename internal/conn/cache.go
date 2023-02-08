package conn

type Cache []Connection

// implement sort.Interface
func (cc Cache) Len() int {
	return len(cc)
}

func (cc Cache) Less(i, j int) bool {
	iel := cc[i]
	jel := cc[j]

	return iel.Delay < jel.Delay
}

func (cc Cache) Swap(i, j int) {
	cc[i], cc[j] = cc[j], cc[i]
}

// implement container.heap.Interface
func (cc *Cache) Push(c any) {
	conn := c.(*Connection)
	*cc = append(*cc, *conn)
}

func (cc *Cache) Pop() any {
	old := *cc
	n := len(old)

	el := old[n-1]
	*cc = old[:n-1]

	return el
}
