package dynamicprogramming

import "log"

func rob(nums []int) int {
	visited := make([]int, len(nums))
	for i := range visited {
		visited[i] = -1
	}

	var dfs func(n int) int
	dfs = func(n int) (num int) {
		if n < 0 {
			return 0
		}

		if visited[n] != -1 {
			return visited[n]
		}

		defer func() {
			visited[n] = num
		}()
		return max(dfs(n-1), dfs(n-2)+nums[n])
	}

	return dfs(len(nums) - 1)
}

func rob2(nums []int) int {
	visited := make([]int, len(nums)+2)

	for i := 0; i < len(nums); i++ {
		visited[i+2] = max(visited[i+1], visited[i]+nums[i])
	}

	return visited[len(nums)+1]
}

func Test_rob() {
	nums1 := []int{1, 2, 3, 1}
	ans1 := rob(nums1)
	log.Printf("ans1: %v", ans1)

	nums2 := []int{2, 7, 9, 3, 1}
	ans2 := rob(nums2)
	log.Printf("ans2: %v", ans2)

}

func Test_rob2() {
	nums1 := []int{1, 2, 3, 1}
	ans1 := rob2(nums1)
	log.Printf("ans1: %v", ans1)

	nums2 := []int{2, 7, 9, 3, 1}
	ans2 := rob2(nums2)
	log.Printf("ans2: %v", ans2)

}
