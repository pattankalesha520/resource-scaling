package main
import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
	rand.Seed(time.Now().UnixNano())
	minReplicas, maxReplicas := 1, 20
	nodeCapacity, minNodes := 4, 1
	replicas, nodes := 2, 1
	highCPU, lowCPU := 70.0, 30.0
	for i := 0; i < 20; i++ {
		cpu := 20 + rand.Float64()*90
		mem := 20 + rand.Float64()*90
		scaleUp := cpu > highCPU || mem > highCPU
		scaleDown := cpu < lowCPU && mem < lowCPU
		if scaleUp && replicas < maxReplicas {
			replicas++
		} else if scaleDown && replicas > minReplicas {
			replicas--
		}
		neededNodes := (replicas + nodeCapacity - 1) / nodeCapacity
		if neededNodes > nodes {
			nodes = neededNodes
		} else if replicas <= (nodes-1)*nodeCapacity && nodes > minNodes {
			nodes--
		}
		fmt.Printf("tick=%02d cpu=%.1f mem=%.1f pods=%d nodes=%d\n", i+1, cpu, mem, replicas, nodes)
		time.Sleep(150 * time.Millisecond)
	}
}
