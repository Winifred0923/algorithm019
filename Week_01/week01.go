package Week_01

import "fmt"

//1.两个数组交集
//中级优解：用map找出重复值,利用map的key值去重复
func intersection(nums1 []int, nums2 []int) []int {
	removeRepeat := make(map[int]int)

	for _, num := range nums1 {
		removeRepeat[num] = 1
	}

	key := 0
	for _, num := range nums2 {
		if val, ok := removeRepeat[num]; ok && val > 0 {
			removeRepeat[num] -= 1
			nums2[key] = num
			key++
		}
	}

	return nums2[0:key]
}
//最优解：双指针


///2.两个数组的交集
func intersect(nums1 []int, nums2 []int) []int {
	hash := map[int]int{}
	res := []int{}
	for _, num1 := range nums1 {
		if hash[num1] > 0 {
			hash[num1]++
		} else {
			hash[num1] = 1
		}
	}
	for _, num2 := range nums2 {
		if hash[num2] > 0 {
			fmt.Println(num2)
			res = append(res, num2)
			hash[num2]--
		}
	}
	return res
}

//给定目标值，返回数组中相加等于目标值的下标
//自己做的
func twoSum(nums []int, target int) []int {
	length := len(nums)
	for i, v := range nums{

		for j:=i+1; j<length; j++ {
			if nums[j] == target - v{
				return []int{i, j}
			}
		}
	}
	return []int{}
}

//最优解， 利用map，少一次循环
func twoSum2(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		if k, ok := m[target-v]; ok {
			return []int{k, i}
		}
		m[v] = i
	}
	return nil
}

//3.盛水最大面积,典型的夹逼方法，移动最小值，
func maxArea(height []int) int {
	var l, r = 0, len(height) - 1
	var maxArea, tmp = 0, 0

	for {
		if l == r {
			break
		}
		lh, rh := height[l], height[r]
		if lh < rh {
			tmp = lh * (r - l)
			l++
		} else {
			tmp = rh * (r - l)
			r--
		}

		if tmp > maxArea {
			maxArea = tmp
		}
	}
	return maxArea
}

//4.爬楼梯,两种解决方法，一种暴力求解，递归，而另一种是动态规划，这种事要培养的思想，找到规律
func fibonacci(n int)(res int) {
	if n <= 1 {
		return 1
	}
	res = fibonacci(n-1) + fibonacci(n - 2)
	return
}

//动态规划:
func climbStairs(n int)(res int){
	switch n {
	case 1:
		return 1
	case 2:
		return 2
	default:
		start , end := 1, 2
		for i := 2; i < n; i++ {
			start, end = end, start + end
		}
		return end
	}
}

//5.加一：[1，2，0]返回[1,2,1],[1,3,9]返回[1,4,0]
//自己解决方法存在的问题，循环遍历，耗时多，最优解的思想其实只考虑最后一位是否为九，是的话前一位加一，不是的话本位加一
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 { // 当前位置不用进位，+1，然后直接返回
			digits[i]++
			return digits
		} else { // 要进位，当前位置置0
			digits[i] = 0
		}
	}
	return append([]int{1}, digits...)
}