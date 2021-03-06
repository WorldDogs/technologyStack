package quick_sort

func sort(arr []int) []int {
	var recurse func(left int, right int)
	var partition func(left int, right int, pivot int) int
	// 分区函数
	partition = func(left int, right int, pivot int) int {
		v := arr[pivot]                                 // 获取基准点
		right--                                         // right 指向arr 最后一个元素
		arr[pivot], arr[right] = arr[right], arr[pivot] //将基准点的值和最右端的值交换

		// i,left 都指向最左端
		// i一直向右移动，left只有在指向的元素比v小才会移动，所以遇到大于v的会停下，
		// 再有比v小的就会进行交换，
		for i := left; i < right; i++ {
			if arr[i] <= v {
				arr[i], arr[left] = arr[left], arr[i]
				left++
			}
		}
		// 最后left和right交换
		arr[left], arr[right] = arr[right], arr[left]
		return left
	}

	recurse = func(left int, right int) {
		if left < right {
			pivot := (right + left) / 2
			pivot = partition(left, right, pivot)
			recurse(left, pivot)
			recurse(pivot+1, right)
		}
	}

	recurse(0, len(arr))
	return arr
}
func QuickSort_p1(arr []int) []int {
	var recurse func(left, right int)
	// 分区函数，将arr分成两部分，左边比arr[pivot]小，右边比arr[pivot]大
	// 最后返回 分区位置，为下次递归用
	partition := func(left, right, pivot int) int {
		v := arr[pivot]
		// 将基准点换到最后面
		arr[pivot], arr[right] = arr[right], arr[pivot]

		for i := left; i < left; i++ {
			if arr[i] <= v {
				arr[i], arr[left] = arr[left], arr[i]
				left++ // 如果arr[i]<=v i,v 同步运动。否则，i走在前面,left指向了比v大的值，
				//	再次遇到比v小的，就会发生交换
			}
		}
		arr[left], arr[right] = arr[right], arr[left]
		return left
	}
	recurse = func(left, right int) {
		if left < right {
			pivot := (right + left) / 2
			pivot = partition(left, right, pivot)
			recurse(left, pivot)
			recurse(left+1, right)
		}
	}
	recurse(0, len(arr)-1)
	return arr
}

// 算法过程：选定基点，使得数据=》 比基点小的在基点左边，比基点大的在基点右边，
//          然后根据基点，将数据一份为二分而治之
// 			分为两个函数：partation、recurce
func QuickSort_p2(arr []int) []int {
	var recurse func(left, right int)
	partation := func(left, right, pivot int) int {
		v := arr[pivot] // 获取基准值
		//	将基准值
		// 将基准点换到最后
		arr[pivot], arr[right] = arr[right], arr[pivot]
		for i := left; i < left; i++ {
			if arr[i] <= v {
				if i != left {
					arr[i], arr[left] = arr[left], arr[i]
					left++
				}
			}
		}
		arr[left], arr[right] = arr[right], arr[left]
		return left
	}
	recurse = func(left, right int) {
		// 确定基准
		if left < right {
			pivot := (left + right) / 2
			pivot = partation(left, right, pivot)
			recurse(left, pivot)
			recurse(pivot+1, right)
		}
	}
	recurse(0, len(arr)-1)
	return arr
}
// 算法流程：根据基准点，小于基准点的放左边，大于的放右边，
// 根据基准点一份为二，分而治之。最后所有的局部是有序的，整体也变为了有序的
//func QuickSort_p3(arr []int)[]int{
//	partation:=func()
//}