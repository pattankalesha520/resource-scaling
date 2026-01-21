package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	id   int
	load int
}

type Cluster struct {
	nodes []Node
}

func (c *Cluster) avgLoad() int {
	sum := 0
	for _, n := range c.nodes {
		sum += n.load
	}
	return sum / len(c.nodes)
}

func (c *Cluster) scaleUp() {
	id := len(c.nodes) + 1
	c.nodes = append(c.nodes, Node{id: id, load: 0})
}

func (c *Cluster) scaleDown() {
	if len(c.nodes) > 1 {
		c.nodes = c.nodes[:len(c.nodes)-1]
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	c := Cluster{nodes: []Node{{1, 0}, {2, 0}, {3, 0}}}
	for i := 0; i < 15; i++ {
		for j := range c.nodes {
			c.nodes[j].load = rand.Intn(100)
		}
		avg := c.avgLoad()
		if avg > 70 {
			c.scaleUp()
		}
		if avg < 30 {
			c.scaleDown()
		}
		fmt.Println("Nodes:", len(c.nodes), "AvgLoad:", avg)
		time.Sleep(500 * time.Millisecond)
	}
}
