package insert_sort
// 挪出一个坑来，把这个arr[i] 目标值填进去
// 可以想象成摸牌的过程：右手拿到新牌，放到左手的有序牌组中
// 时间：O(N^2) 稳定
// 适用场景：数组基本有序且数据量较小
func InsertSort(arr []int){
	for i := 1; i < len(arr); i++ {
		value := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > value {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = value
	}
}

// 默写：第一次
func InsertSort_P(arr []int){
	for i:=1;i<len(arr);i++{
		v:=arr[i]
		j:=i-1
		for j>=0&&arr[j]>v{
			arr[j] = arr[j+1]
			j--
		}
		arr[j+1] = v
	}
}