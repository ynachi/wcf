// This file implements the heap interface as a max heap
package wcf

type MinHeap []Word

func (mh MinHeap) Len() int {
	return len(mh)
}

func (mh MinHeap) Less(i, j int) bool {
	return mh[i].Count < mh[j].Count
}

func (mh MinHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh *MinHeap) Push(x any) {
	*mh = append(*mh, x.(Word))
}

func (mh *MinHeap) Pop() any {
	old := *mh
	n := len(old)
	x := old[n-1]
	*mh = old[0 : n-1]
	return x
}
