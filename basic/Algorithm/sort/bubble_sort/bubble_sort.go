package bubble_sort


// 冒泡排序
// O(n^2)
// 每走一趟，产生一个最大值，右侧有序数目不断增多
// 优化：swap 记录当次扫描是否有交换，如果没有交换表示前面的已经全部有序
func Bubble_Sort(arr []int){
	swap:=false
	for cnt:=len(arr)-1;cnt>0;cnt--{
		swap=false
		for i:=1;i<=cnt;i++{
			if arr[i-1]>arr[i]{
				arr[i-1],arr[i] = arr[i],arr[i-1]
				swap = true
			}
		}
		if swap == false{
			break
		}
	}
}