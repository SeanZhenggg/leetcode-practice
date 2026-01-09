package arraysAndHashing

import (
	"github.com/stretchr/testify/assert"
	"log"
	"sort"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	assertion := assert.New(t)

	ans1 := topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
	log.Println("ans1: ", ans1)
	sort.Ints(ans1)
	assertion.Equal(ans1, []int{1, 2})

	ans2 := topKFrequent([]int{1}, 1)
	log.Println("ans2: ", ans2)
	sort.Ints(ans2)
	assertion.Equal(ans2, []int{1})
}

func TestTopKFrequent2(t *testing.T) {
	assertion := assert.New(t)

	ans1 := topKFrequent2([]int{1, 1, 1, 2, 2, 3}, 2)
	log.Println("ans1: ", ans1)
	sort.Ints(ans1)
	assertion.Equal(ans1, []int{1, 2})

	ans2 := topKFrequent2([]int{1}, 1)
	log.Println("ans2: ", ans2)
	sort.Ints(ans2)
	assertion.Equal(ans2, []int{1})

	ans3 := topKFrequent2([]int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}, 2)
	log.Println("ans3: ", ans3)
	sort.Ints(ans3)
	assertion.Equal(ans3, []int{1, 2})

	ans4 := topKFrequent2([]int{-5, -2, -1, 2, 5, -5, 2, -1}, 3)
	log.Println("ans4: ", ans4)
	sort.Ints(ans4)
	assertion.Equal(ans4, []int{-5, -1, 2})
}

func TestTopKFrequent3(t *testing.T) {
	assertion := assert.New(t)

	ans1 := topKFrequent3([]int{1, 1, 1, 2, 2, 3}, 2)
	log.Println("ans1: ", ans1)
	sort.Ints(ans1)
	assertion.Equal(ans1, []int{1, 2})

	ans2 := topKFrequent3([]int{1}, 1)
	log.Println("ans2: ", ans2)
	sort.Ints(ans2)
	assertion.Equal(ans2, []int{1})

	ans3 := topKFrequent3([]int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}, 2)
	log.Println("ans3: ", ans3)
	sort.Ints(ans3)
	assertion.Equal(ans3, []int{1, 2})

	ans4 := topKFrequent3([]int{-5, -2, -1, 2, 5, -5, 2, -1}, 3)
	log.Println("ans4: ", ans4)
	sort.Ints(ans4)
	assertion.Equal(ans4, []int{-5, -1, 2})
}

func TestTopKFrequent4(t *testing.T) {
	assertion := assert.New(t)

	ans1 := topKFrequent4([]int{1, 1, 1, 2, 2, 3}, 2)
	log.Println("ans1: ", ans1)
	sort.Ints(ans1)
	assertion.Equal(ans1, []int{1, 2})

	ans2 := topKFrequent4([]int{1}, 1)
	log.Println("ans2: ", ans2)
	sort.Ints(ans2)
	assertion.Equal(ans2, []int{1})

	ans3 := topKFrequent4([]int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}, 2)
	log.Println("ans3: ", ans3)
	sort.Ints(ans3)
	assertion.Equal(ans3, []int{1, 2})

	ans4 := topKFrequent4([]int{-5, -2, -1, 2, 5, -5, 2, -1}, 3)
	log.Println("ans4: ", ans4)
	sort.Ints(ans4)
	assertion.Equal(ans4, []int{-5, -1, 2})
}
