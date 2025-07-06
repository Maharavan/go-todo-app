package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fields := strings.Fields(input)

	fmt.Println(fields)

	for _, out := range fields {
		num, _ := strconv.Atoi(out)
		if num%2 == 0 {
			fmt.Printf("Given number is even %d\n", num)
		} else {
			fmt.Printf("Given number is odd %d\n", num)
		}
	}

}
