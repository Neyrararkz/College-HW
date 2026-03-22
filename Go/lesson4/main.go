package main

import "fmt"

func main() {

	// ========= Массивы =========
	arr := [10]int{1,2,3,4,5,6,7,8,9,10}

	//1
	for _, i := range arr {
		fmt.Println(i)
	}

	//2
	sum := 0
	for _, i := range arr {
		sum += i
	}
	fmt.Println("Sum: ", sum)

	//3
	max := arr[0]
	min := arr[0]
	for i := range arr{
		if(arr[i] > max){
			max = arr[i]
		}
		if(arr[i] < min){
			min = arr[i]
		}
	}
	fmt.Println("Max: ", max, "\nMin: ", min)

	//4
	count := 0
	for _, i := range arr{
		if(i%2 == 0){
			count++
		}
	}
	fmt.Println("Count: ", count)

	//5	
	for i := range arr{
		arr[i] *= 2 
	}
	for _, i := range arr{
		fmt.Println(i)
	}


	// ========= Слайсы =========
	s := []int{1, 2, 3, 4, 5}

	//6
	s = append(s, 6, 7, 8)
	fmt.Println(s)		

	// 7
	s = append(s[:2], s[3:]...)
	fmt.Println(s)

	//8
	sum2 := 0
	for _, i := range s {
		sum += i
	}
	fmt.Println(float64(sum2) / float64(len(s)))

	//9
	s1 := []string{"Lydia", "Theo", "Dylan"}
	short := s1[0]
	for _, i := range s1{
		if (len(i) < len(short)){
			short = i
		}
	}
	fmt.Println(short)

	//10
	s2 := []int{1,10,2,20,3,30}
	s3 := []int{}
	for _, i := range s2{
		if(i>=10){
			s3 = append(s3, i)
		}
	}
	fmt.Println(s3)
	
}