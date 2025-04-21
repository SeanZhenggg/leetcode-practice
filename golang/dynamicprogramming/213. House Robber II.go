package dynamicprogramming

import "log"

func robII(nums []int) int {
	visited := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		visited[i] = -1
	}
	var dfs func(n int, minN int) (num int)
	dfs = func(n int, minN int) (num int) {
		if n < minN {
			return 0
		}

		if visited[n] != -1 {
			return visited[n]
		}

		defer func() {
			visited[n] = num
		}()
		return max(dfs(n-1, minN), dfs(n-2, minN)+nums[n])
	}

	n1 := dfs(len(nums)-1, 1)
	for i := 0; i < len(nums); i++ {
		visited[i] = -1
	}
	n2 := dfs(len(nums)-2, 0)
	return max(n1, n2)
}

func robII2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums)+2)

	for i := 0; i < len(nums)-1; i++ {
		dp[i+2] = max(dp[i+1], dp[i]+nums[i])
	}

	v1 := dp[len(nums)]

	dp = make([]int, len(nums)+2)
	for i := 1; i < len(nums); i++ {
		dp[i+2] = max(dp[i+1], dp[i]+nums[i])
	}
	v2 := dp[len(nums)+1]

	return max(v1, v2)
}

func Test_robII() {
	nums1 := []int{2, 3, 2}
	ans1 := robII(nums1)
	log.Printf("ans1: %v", ans1)

	nums2 := []int{1, 2, 3, 1}
	ans2 := robII(nums2)
	log.Printf("ans2: %v", ans2)

	nums3 := []int{1, 2, 3}
	ans3 := robII(nums3)
	log.Printf("ans3: %v", ans3)

}

func Test_robII2() {
	nums1 := []int{2, 3, 2}
	ans1 := robII2(nums1)
	log.Printf("ans1: %v", ans1)

	nums2 := []int{1, 2, 3, 1}
	ans2 := robII2(nums2)
	log.Printf("ans2: %v", ans2)

	nums3 := []int{1, 2, 3}
	ans3 := robII2(nums3)
	log.Printf("ans3: %v", ans3)

}
