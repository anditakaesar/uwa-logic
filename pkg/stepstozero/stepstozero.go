package stepstozero

import "fmt"

func Run() {
	fmt.Println("[main] stepstozero")
	nums := []int{1, 3, 14, 8}
	for _, num := range nums {
		fmt.Printf("num: %d, steps: %d\n", num, CountStepsToZero(num))
	}

	for _, num := range nums {
		fmt.Printf("num: %d, steps: %d\n", num, CountStepsToZero2(num))
	}
}

func CountStepsToZero(num int) int {
	var count int
	for num > 0 {
		count += 1
		if num == 1 {
			return count
		}
		if num%2 == 0 {
			num /= 2
			continue
		}
		num = num - 1
	}
	return count
}

func CountStepsToZero2(num int) int {
	//fmt.Printf("%d >> %d = %d\n", 8, 1, 8>>1)
	//fmt.Printf("%d & %d = %d\n", 7, 1, 7&1)
	//fmt.Printf("%d & %d = %d\n", 8, 1, 8&1)
	var count int
	for num > 0 {
		count += 1
		if num == 1 {
			return count
		}
		if num&1 == 0 { // this compare the last bit representation of num
			num = num >> 1
			continue
		}
		num -= 1
	}

	return 0
}
