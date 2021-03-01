package heap_sort

import "fmt"

// 堆：完全二叉树(除了最后一层节点都铺满，且最后一层是从左到右)+ 子节点大于或者小于父节点
// 完全二叉树，可以使用数组进行存储
// 大顶堆：根节点最大
// 小顶堆：根节点最小

// heapify

// 堆排序其实就是将无序数据，首先进行堆化，然后不断取父节点的过程
// 过程：首先找到最后一个节点的父节点，然后进行heapify操作，直到根节点
// 原理：从下向上构建堆，最后整个堆都符合性质。
// 应用：1000w数据 最小的5个，此时可以用最大堆。过程：首先拿五个元素构建堆，新数据如果比堆定小，删除堆顶，然后新元素放入堆顶，heapify。
// todo @zsl 将应用进行实现
func heap_sort(arr []int) []int {
	//	 获取目标节点
	tar := (len(arr) - 2) / 2
	for i := tar; i >= 0; i-- {
		heapify(arr, i)
	}
	fmt.Println("build heap", arr)
	//	此时大根堆已经形成，只需要不断的取根节点，然后heapify就行
	tmp := make([]int, len(arr))
	ans := make([]int, 0, len(arr))
	copy(tmp, arr)
	for i := 0; i < len(arr); i++ {
		ans = append(ans, tmp[0])
		tmp[len(tmp)-1], tmp[0] = tmp[0], tmp[len(tmp)-1]
		tmp = tmp[:len(tmp)-1]
		heapify(tmp, 0)
		fmt.Println("heapify", tmp)
	}
	fmt.Println("ans", ans)
	return ans
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
	if l < len(arr) && arr[l] > arr[max] {
		max = l
	}
	if r < len(arr) && arr[r] > arr[max] {
		max = r
	}
	if max != root {
		arr[root], arr[max] = arr[max], arr[root]
		heapify(arr, max)
	}
}



//func main() {
//	target := []int{1, 1, 2}
//	heapify(target, 0)
//	fmt.Println(target)
//	heap_sort([]int{2, 1, 2, 3, 5, 10})
//}
