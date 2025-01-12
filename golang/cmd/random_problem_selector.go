package main

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"strings"
	"time"
)

const (
	Filename           = "./problems"
	LeetcodeUrlPattern = "https://leetcode.com/problems/%s"
)

func main() {
	choosed := make(map[int]bool)
	// 讀取 keyboard std input

	for {
		fmt.Printf("Enter command(n: choose a random question, q: exit): \n")
		var command string
		_, err := fmt.Scan(&command)
		if err != nil {
			fmt.Printf("fmt.Scan err occurred: %s\n", err)
			return
		}

		// 讀入 problems 底下的所有資料夾下的檔案名稱, 存入 list
		dirs, err := os.ReadDir(Filename)
		if err != nil {
			fmt.Printf("os.ReadDir(FILENAME) err occurred: %s\n", err)
			return
		}

		problemList := make([]os.DirEntry, 0)
		for _, dir := range dirs {
			if dir.IsDir() {
				dirName := path.Join(Filename, dir.Name())
				files, err := os.ReadDir(dirName)
				if err != nil {
					fmt.Printf("os.ReadDir(dirName) err occurred: %s\n", err)
					return
				}

				for _, file := range files {
					if !file.IsDir() {
						problemList = append(problemList, file)
					}
				}
			}
		}
		problemList = shuffle(problemList)

		switch command {
		case "n":
			// n -> choose question

			// random 選擇一個 list 中的 item, 不得與之前重複
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			idx := r.Intn(len(problemList) - 1)
			for choosed[idx] {
				r = rand.New(rand.NewSource(time.Now().UnixNano()))
				idx = r.Intn(len(problemList) - 1)
			}
			choosed[idx] = true
			filename := problemList[idx].Name()

			// output
			splitted := strings.SplitN(filename, ".", 3)
			no, filename, _ := splitted[0], splitted[1], splitted[2]
			filename = strings.TrimLeft(filename, " ")
			url := fmt.Sprintf(LeetcodeUrlPattern, strings.ToLower(strings.ReplaceAll(strings.TrimLeft(filename, " "), " ", "-")))
			fmt.Printf("Problem %s: %s\n", no, filename)
			fmt.Println(url)
		case "q":
			// q -> exit
			fmt.Println("Bye")
			return
		default:
			// 已經選擇過的不重複選擇
			fmt.Printf("invalid command: %s, program exit\n", command)
			return
		}

		fmt.Println()
	}
}

func shuffle[T any](input []T) []T {
	for i := 0; i < len(input); i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))

		j := r.Intn(i + 1)
		input[i], input[j] = input[j], input[i]
	}
	return input
}
