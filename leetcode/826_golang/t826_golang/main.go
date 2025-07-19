package main

import "container/heap"

type WorkHeap struct{}

func (w WorkHeap) Len() int {
	//TODO implement me
	panic("implement me")
}

func (w WorkHeap) Less(i, j int) bool {
	//TODO implement me
	panic("implement me")
}

func (w WorkHeap) Swap(i, j int) {
	//TODO implement me
	panic("implement me")
}

func (w WorkHeap) Push(x any) {
	//TODO implement me
	panic("implement me")
}

func (w WorkHeap) Pop() any {
	//TODO implement me
	panic("implement me")
}

func createHeap() heap.Interface {
	return &WorkHeap{}
}

func main() {

}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {

}
