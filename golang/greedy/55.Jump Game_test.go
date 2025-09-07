package greedy

import (
	"log"
	"testing"
)

func Test_canJump(t *testing.T) {
	case1 := []int{2, 3, 1, 1, 4}
	ans1 := canJump(case1)
	log.Printf("ans1: %v", ans1)
	log.Println()
	case2 := []int{3, 2, 1, 0, 4}
	ans2 := canJump(case2)
	log.Printf("ans2: %v", ans2)
	log.Println()
	case3 := []int{2, 1, 0, 1, 4}
	ans3 := canJump(case3)
	log.Printf("ans3: %v", ans3)

	case4 := []int{3, 2, 4, 0, 5, 0, 2, 1, 0, 3, 10}
	ans4 := canJump(case4)
	log.Printf("ans4: %v", ans4)
}

func Test_canJump2(t *testing.T) {
	case1 := []int{2, 3, 1, 1, 4}
	ans1 := canJump2(case1)
	log.Printf("ans1: %v", ans1)
	log.Println()
	case2 := []int{3, 2, 1, 0, 4}
	ans2 := canJump2(case2)
	log.Printf("ans2: %v", ans2)
	log.Println()
	case3 := []int{2, 1, 0, 1, 4}
	ans3 := canJump2(case3)
	log.Printf("ans3: %v", ans3)

	case4 := []int{3, 2, 4, 0, 5, 0, 2, 1, 0, 3, 10}
	ans4 := canJump2(case4)
	log.Printf("ans4: %v", ans4)
}

func Test_canJump3(t *testing.T) {
	case1 := []int{2, 3, 1, 1, 4}
	ans1 := canJump3(case1)
	log.Printf("ans1: %v", ans1)
	log.Println()
	case2 := []int{3, 2, 1, 0, 4}
	ans2 := canJump3(case2)
	log.Printf("ans2: %v", ans2)
	log.Println()
	case3 := []int{2, 1, 0, 1, 4}
	ans3 := canJump3(case3)
	log.Printf("ans3: %v", ans3)

	case4 := []int{3, 2, 4, 0, 5, 0, 2, 1, 0, 3, 10}
	ans4 := canJump3(case4)
	log.Printf("ans4: %v", ans4)
}

func Test_canJump4(t *testing.T) {
	case1 := []int{2, 3, 1, 1, 4}
	ans1 := canJump4(case1)
	log.Printf("ans1: %v", ans1)
	log.Println()
	case2 := []int{3, 2, 1, 0, 4}
	ans2 := canJump4(case2)
	log.Printf("ans2: %v", ans2)
	log.Println()
	case3 := []int{2, 1, 0, 1, 4}
	ans3 := canJump4(case3)
	log.Printf("ans3: %v", ans3)

	case4 := []int{3, 2, 4, 0, 5, 0, 2, 1, 0, 3, 10}
	ans4 := canJump4(case4)
	log.Printf("ans4: %v", ans4)
}
