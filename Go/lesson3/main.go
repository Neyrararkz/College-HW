package main

import "fmt"

func task9(s []int) {
	s[0] = 9
}

func task10(s []int) {
	s = append(s, 10)
	fmt.Println(s, len(s), cap(s))
}

func main() {
	// BLOCK 1
	// //1
	// arr := [5]int{1,2,3,4,5}

	// fmt.Println(arr)
	// fmt.Println(len(arr))

	// //2
	// arr2 := [4]int{1,2,3,4}
	// fmt.Println(arr2)
	// arr2[1] = 5
	// arr2[3] = 6
	// fmt.Println(arr2)

	// //3
	// a := []int{1,2,3}
	// b := a[:]
	// a[2] = 4
	// fmt.Println(a, b)

	// //BLOCK 2
	// //4
	// var s []int
	// fmt.Println(s)
	// fmt.Println(len(s))
	// fmt.Println(cap(s))
	// fmt.Println(s == nil)

	// //5
	// s5 := []int{1,2,3}
	// s5 = append(s5, 4)
	// fmt.Println(s5)

	// //6
	// s6 := make([]int, 3, 5)
	// s6[0] = 1
	// s6[1] = 2
	// s6[2] = 3
	// fmt.Println(s6)
	// fmt.Println(cap(s6))
	// s6 = append(s6, 4)
	// s6 = append(s6, 5)
	// s6 = append(s6, 6)
	// fmt.Println(s6)
	// fmt.Println(cap(s6))

	// //BLOCK 3
	// //7
	// arr7 := [6]int{1,2,3,4,5,6}
	// fmt.Println(arr7)
	// s7 := arr7[2:4]
	// s7[1] = 7
	// fmt.Println(arr7)
	// fmt.Println(s7)

	// //8
	// s8 := s7[:]
	// s8[0] = 8
	// fmt.Println(arr7)
	// fmt.Println(s7)
	// fmt.Println(s8)
	// //так он изменился
	
	// //BLOCK 4
	// //9
	// s9 := []int{1,2,3,4}
	// fmt.Println(s9)
	// task9(s9)
	// fmt.Println(s9)

	//10
	s10 := make([]int, 4, 4)
	task10(s10)
	fmt.Println(s10, len(s10), cap(s10))
	s10[0], s10[1], s10[2], s10[3] = 1,2,3,4 
	
}