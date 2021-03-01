package heap_sort

// 堆：完全二叉树(除了最后一层节点都铺满，且最后一层是从左到右)+ 子节点大于或者小于父节点
// 完全二叉树，可以使用数组进行存储
// 大顶堆：根节点最大
// 小顶堆：根节点最小

// heapify

// 堆排序其实就是将无序数据，首先进行堆化，然后不断取父节点的过程
// 过程：首先找到最后一个节点的父节点，然后进行heapify操作，直到根节点
// 原理：从下向上构建堆，最后整个堆都符合性质。
func heap_sort(arr []int) []int {
	//	 获取目标节点
	tar := (len(arr) - 2) / 2
	for i := tar; i >= 0; i-- {
		heapify(arr, i)
	}
	//	此时大根堆已经形成，只需要不断的取根节点，然后heapify就行
	tmp:=make([]int,len(arr))
	copy(tmp,arr)
	for i:=0;i<len(arr);i++{
		tmp[len(tmp)],tmp[0]= tmp[0],tmp[len(tmp)]
		
	}


}

//大根堆化
// 算法过程：找出父子几点最大的，换到父节点，然后如果发生了交换需要将变化传递下去
func heapify(arr []int, root int) {
	// 终止条件
	if root >= len(arr) {
		return
	}
	l := root*2 + 1
	r := root*2 + 2
	max := root
	if arr[l] > arr[max] {
		max = l
	}
	if arr[r] > arr[max] {
		max = r
	}
	if max != root {
		arr[root], arr[max] = arr[max], arr[root]
		heapify(arr, max)
	}
}
