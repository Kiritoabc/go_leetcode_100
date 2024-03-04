package demo

import (
	"sort"
	"time"
)

// 1. 给定一个整数 n ，返回 n! 结果中尾随零的数量。

func trailingZeros(n int) int {
	var ans int
	for n > 0 {
		ans += n / 5
		n /= 5
	}
	return ans
}

// 2. 给你两个链表 list1 和 list2 ，它们包含的元素分别为 n 个和 m 个
//。  请你将 list1 中下标从 a 到 b 的全部节点都删除，并将list2 接在被删除节点的位置。

// 3.实现支持下列接口的「快照数组」- SnapshotArray：

type SnapshotArray struct {
	id   int
	snap map[int]map[int]int
}

func Constructor(length int) SnapshotArray {
	return SnapshotArray{0, make(map[int]map[int]int, 0)}
}

func (ssArray *SnapshotArray) Set(index int, val int) {
	if ssArray.snap[index] == nil {
		ssArray.snap[index] = make(map[int]int, 0)
	}
	//记录下一次快照前，改变过的index值
	ssArray.snap[index][ssArray.id] = val
}

func (ssArray *SnapshotArray) Snap() int {
	ssArray.id++
	//返回上一次的snap_id
	return ssArray.id - 1
}

/*如果index当前snap_id不存在，则说明snap_id这次快照跟snap_id-1这次快照之间index的值没变化过
二分找到snpa_id前第一个存在的快照id
但是go没有unorderedmap，还要枚举allkeys排序再二分,还不如直接遍历O(N)
*/

func (ssArray *SnapshotArray) Get(index int, snap_id int) int {
	if value, ok := ssArray.snap[index][snap_id]; ok {
		return value
	}
	//枚举所有key
	allkeys := make([]int, len(ssArray.snap[index]))
	i := 0
	for key, _ := range ssArray.snap[index] {
		allkeys[i] = key
		i++
	}
	//排序已备二分
	sort.Ints(allkeys)

	low := 0
	high := len(allkeys) - 1
	for low <= high {
		mid := (low + high) / 2
		if allkeys[mid] < snap_id {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	//如果low==0，说明不存在这样的快照，返回初始值0
	if low == 0 {
		return 0
	} else {
		//找到第一个比目标值大的下标，需要-1
		return ssArray.snap[index][allkeys[low-1]]
	}
}

func (ssArray *SnapshotArray) Get2(index int, snap_id int) int {
	for snap_id >= 0 {
		if v, ok := ssArray.snap[index][snap_id]; ok {
			return v
		} else {
			snap_id--
		}
	}
	return 0
}

func demo() {
	select {
	case <-time.After(10 * time.Second):
	}
}

// 排序算法

func bubbleSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func selectSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}

func insertSort(arr []int) []int {
	for i, v := range arr {
		preIndex := i - 1
		for preIndex > 0 && arr[preIndex] > v {
			arr[preIndex+1] = arr[preIndex]
			preIndex--
		}
		arr[preIndex+1] = v
	}
	return arr
}

func shellSort(arr []int) []int {
	length := len(arr)
	gap := 1
	for gap < length/3 {
		gap = gap*3 + 1
	}
	for gap > 0 {
		for i := gap; i < length; i++ {
			temp := arr[i]
			j := i - gap
			for j >= 0 && arr[j] > temp {
				arr[j+gap] = arr[j]
				j -= gap
			}
			arr[j+gap] = temp
		}
		gap /= 3
	}
	return arr
}

func mergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	mid := length / 2
	left := arr[0:mid]
	right := arr[mid:]
	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) []int {
	var result []int
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	if len(left) > 0 {
		result = append(result, left...)
	}

	if len(right) > 0 {
		result = append(result, right...)
	}

	return result
}

func quickSort(arr []int) []int {
	return _quickSort(arr, 0, len(arr)-1)
}

func _quickSort(arr []int, left, right int) []int {
	if left < right {
		partitionIndex := partition(arr, left, right)
		_quickSort(arr, left, partitionIndex-1)
		_quickSort(arr, partitionIndex+1, right)
	}
	return arr
}

func partition(arr []int, left, right int) int {
	pivot := left
	index := pivot + 1

	for i := index; i <= right; i++ {
		if arr[i] < arr[pivot] {
			arr[i], arr[index] = arr[index], arr[i]
			index++
		}
	}
	arr[pivot], arr[index-1] = arr[index-1], arr[pivot]
	return index - 1
}

func countingSort(arr []int, maxValue int) []int {
	bucketLen := maxValue + 1
	bucket := make([]int, bucketLen)
	sortedIndex := 0
	// 放入
	for _, v := range arr {
		bucket[v]++
	}
	// 拿出
	for i, v := range bucket {
		for v > 0 {
			arr[sortedIndex] = i
			sortedIndex++
			v--
		}
	}
	return arr
}

func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	path := make([]int, k)
	var back func(start, sum int)
	back = func(start, sum int) {
		// result
		if len(path) == k && sum == n {
			res = append(res, append([]int{}, path...))
			return
		}

		// login loop
		for i := start; i <= 9; i++ {
			// condition

			path = append(path, i)
			back(i+1, sum+i)
			path = path[:len(path)-1]
		}
	}
	back(1, 0)
	return res
}

var (
	res  [][]int
	path []int
)

func combinationSum(candidates []int, target int) [][]int {
	res, path = make([][]int, 0), make([]int, 0)
	sort.Ints(candidates)
	dfs(candidates, target, 0, 0)
	return res
}

func dfs(candidates []int, target, start, sum int) {
	if sum >= target {
		if sum == target {
			res = append(res, append([]int{}, path...))
		}
		return
	}

	for i := start; i < len(candidates); i++ {
		if sum+candidates[i] > target {
			return
		}
		path = append(path, candidates[i])
		dfs(candidates, target, i, sum+candidates[i])
		path = path[:len(path)-1]
	}
}
