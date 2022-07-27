package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

const (
	initialBufSize = 10000
)

func main() {
	buf := make([]byte, initialBufSize)
	sc.Buffer(buf, 999999999999999999)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	var array []int
	var result []int
	for i := 0; i < n; i++ {
		sc.Scan()
		tmp, _ := strconv.Atoi(sc.Text())
		array = append(array, tmp)
		result = append(result, tmp)
	}
	for i := 0; i < n-1; i++ {
		for result[i] != result[i+1] {
			if result[i] < result[i+1] {
				result[i] = result[i] + array[i]
			} else {
				result[i+1] = result[i+1] + array[i+1]
			}
		}
		array[i+1] = result[i+1]
	}
	fmt.Println(result[n-1])
}
