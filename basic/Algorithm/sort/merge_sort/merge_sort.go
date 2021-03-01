package main

import "fmt"

// 此实现修改了原数据
// sort 注释 是第一遍注释
func sort(arr []int) {
	var s = make([]int, len(arr)/2+1)
	if len(arr) < 2 { // 数据<2 直接返回 （不需要再归并下去了）
		return
	}

	mid := len(arr) / 2 //取中点

	sort(arr[:mid]) // 左侧递归 [left,mid)
	sort(arr[mid:]) // 右侧递归 [mid,right]

	if arr[mid-1] <= arr[mid] { // 如果左右两个区间已经有序，无需数据重组，直接返回
		return
	}

	copy(s, arr[:mid]) // 拷贝 s=[left,mid)

	l, r := 0, mid

	for i := 0; ; i++ {
		if s[l] <= arr[r] {
			arr[i] = s[l]
			l++

			if l == mid { // l=mid，说明右区间最小值大于左区间最大值，直接返回
				break
			}
		} else {
			arr[i] = arr[r]
			r++
			if r == len(arr) { // r=len(arr) 说明右区间最大值小于左区间最小值，所以直接将左区间追加
				copy(arr[i+1:], s[l:mid])
				break
			}
		}
	}
	return
}

// 此注释是第二次注释，更精简
func MergeSort_p2(arr []int) {
	if len(arr) < 2 { // 数据<2 直接返回
		return
	}
	//	取中点
	mid := len(arr) / 2

	// 归并左右区间
	sort(arr[:mid])
	sort(arr[mid:])

	//	优化
	if arr[mid-1] <= arr[mid] { // 如果左右区间已经有序，直接返回
		return
	}

	// 错误代码实例
	// s:=[]int{} // len cap 都为0
	// copy(s,arr[:mid])
	// copy 拷贝长度为两个切片长度最小值
	// 所以s 还是为空

	s:=make([]int,len(arr)/2)
	copy(s,arr[:mid])
	l,r:=0,mid
	for i:=0;;i++{
		fmt.Println(i,l,r)
		if s[l] <=arr[r]{
			arr[i] = s[l]
			l++
			if l==mid{  // 左区间先走完了，右区间数据不用动
				break
			}
		}else{
			arr[i] = arr[r]
			r++
			if r==len(arr){ // 右区间走完了，需要把左区间剩余的数据拷贝过来
				copy(arr[i+1:],s[l:mid])
				break
			}
		}
	}
}
