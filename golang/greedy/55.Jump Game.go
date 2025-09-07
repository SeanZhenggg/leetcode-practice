package greedy

func canJump(nums []int) bool {
	pathMap := map[int]bool{}

	for i := 0; i < len(nums); i++ {
		pathMap[i] = true
	}

	var dfs func(i int) bool
	dfs = func(i int) bool {
		// log.Printf("i now: %d", i)
		if i == len(nums)-1 {
			return true
		}

		if nums[i] == 0 {
			// log.Printf("i: %d failed, record false", i)
			pathMap[i] = false
			return false
		}

		for j := nums[i]; j >= 1; j-- {
			// log.Printf("try i+j: %d", i+j)
			if !pathMap[i+j] {
				continue
			}

			if dfs(i + j) {
				return true
			}
		}
		// log.Printf("i: %d failed, record false", i)
		pathMap[i] = false
		return false
	}

	return dfs(0)
}

func canJump2(nums []int) bool {
	memo := make([]bool, len(nums))

	memo[len(nums)-1] = true

	for i := len(nums) - 2; i >= 0; i-- {
		next := min(i+nums[i], len(nums)-1)

		for j := i + 1; j <= next; j++ {
			if memo[j] {
				memo[i] = true
				break
			}
		}
	}
	return memo[0] == true
}

func canJump3(nums []int) bool {
	leftMost := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i]+i >= leftMost {
			leftMost = i
		}
	}

	return leftMost == 0
}

func canJump4(nums []int) bool {
	maxReach := 0                    // 目前能到達的最遠位置
	for i := 0; i < len(nums); i++ { // 每一個位置都檢查, 包含終點
		if i > maxReach { // 看看之前記錄的最遠位置是否能夠到達當前位置, 不行的話代表我們根本走不到這裡, 也就意味著走不到終點
			return false
		}
		maxReach = max(maxReach, i+nums[i]) // 更新在當前位置上能夠到達的最遠位置
		if maxReach >= len(nums)-1 {        // 如果能到達的最遠位置已經是終點或是比終點更遠，代表我們能夠抵達終點
			return true
		}
	}

	return false // 如果迴圈的最後在終點位置 maxReach 沒有辦法到達的話，代表我們走不到終點
}
