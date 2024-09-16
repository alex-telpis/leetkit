package leetkit

// IntMinHeap is a Min-Heap implementation for integers.
// It allows you to quickly test heap-based solutions.
// However, after testing, you'll need to copy this code and include it in your submission.
type IntMinHeap []int

func (h IntMinHeap) Len() int               { return len(h) }
func (h IntMinHeap) Less(i int, j int) bool { return h[i] < h[j] }
func (h IntMinHeap) Swap(i int, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntMinHeap) Push(x any)            { *h = append(*h, x.(int)) }
func (h *IntMinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// IntMaxHeap is a Max-Heap implementation for integers.
// It allows you to quickly test heap-based solutions.
// However, after testing you'll need to copy this code and include it in your submission.
type IntMaxHeap []int

func (h IntMaxHeap) Len() int               { return len(h) }
func (h IntMaxHeap) Less(i int, j int) bool { return h[i] > h[j] }
func (h IntMaxHeap) Swap(i int, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntMaxHeap) Push(x any)            { *h = append(*h, x.(int)) }
func (h *IntMaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
