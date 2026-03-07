package main

import "fmt"

//BLOCK 1
//1
func inRange(x int, min int, max int) bool {
	if x >= min && x <= max {
		return true
	} else {
		return false
	}
}

//2
func ageCategory(age int) string {
	switch {
	case age < 12:
		return "Child"
	case age < 18:
		return "Teen"
	case age < 65:
		return "Adult"
	default:
		return "Senior"
	}
}

//3
func compareThree(a int, b int, c int) string {
	if a == b && a == c {
		return "all equal"
	} else if a == b || b == c || c == a {
		return "two equal"
	} else {
		return "all different"
	}
}

//BLOCK 2
//4
func countDivisors(n int) int {
	var count int
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}
	}
	return count
}

//5
func firstEven() int {
	for {
		var num int
		fmt.Println("5. Введите четное число: ")
		fmt.Scan(&num)
		if num%2 == 0 {
			return num
		}
	}
}

//6
func averageUntilZero() float64 {
	var sum float64
	var count float64
	for {
		var num float64
		fmt.Println("6. Введите число: ")
		fmt.Scan(&num)
		if num == 0 {
			if count == 0 {
				return 0
			}
			break
		}
		sum += num
		count++
	}
	return sum / count
}

//BLOCK 3
//7
func readPositive() int {
	for {
		var num int
		fmt.Println("7. Введите положительное число: ")
		fmt.Scan(&num)
		if num > 0 {
			return num
		}
	}
}

//8
func safeCompare(a1 int, b1 int) (string, bool) {
	switch {
	case a1 == b1:
		return "equal", true
	case a1 > b1:
		return "a greater", true
	default:
		return "b greater", true
	}
}

//BLOCK 4
//9
func menu() {
	for{
		var cmd int
		fmt.Print("Что вы хотите сделать?\n1 - Проверка диапазона\n2 - Категория возраста\n3 - Количество делителей\n4 - Среднее до нуля\n0 - Выход\n")
		fmt.Scan(&cmd)

		switch cmd{
		case 1:
			var x, min, max int
			fmt.Println("1. Введите число, минимальное и максимльное значение: ")
			fmt.Scan(&x, &min, &max)
			fmt.Println(inRange(x, min, max))
		case 2:
			var age int
			fmt.Println("2. Введите свой возраст: ")
			fmt.Scan(&age)
			fmt.Println(ageCategory(age))
		case 3:
			var n int
			fmt.Println("4. Введите число: ")
			fmt.Scan(&n)
			fmt.Println(countDivisors(n))
		case 4:
			fmt.Println(averageUntilZero())
		case 0:
			fmt.Println("Вы вышли.")
				return
		default:
			fmt.Println("Нельзя так")
		}
	}
}

//10
func gcd(o int, k int) int {
	if (o > k){
		var e int
		for{
			e = o%k
			if (e == 0){
				return k
			}
			o = k
			k = e
		}
	} else {
		var e int
		for{
			e = k%o
			if (e == 0){
				return k
			}
			k = o
			o = e
		}
	}
}

func main() {
	//1
	var x, min, max int
	fmt.Println("1. Введите число, минимальное и максимльное значение: ")
	fmt.Scan(&x, &min, &max)
	fmt.Println(inRange(x, min, max))

	//2
	var age int
	fmt.Println("2. Введите свой возраст: ")
	fmt.Scan(&age)
	fmt.Println(ageCategory(age))

	//3
	var a, b, c int
	fmt.Println("3. Введите три числа: ")
	fmt.Scan(&a, &b, &c)
	fmt.Println(compareThree(a, b, c))

	// BLOCK 2
	//4
	var n int
	fmt.Println("4. Введите число: ")
	fmt.Scan(&n)
	fmt.Println(countDivisors(n))

	//5
	fmt.Println(firstEven())

	//6
	fmt.Println(averageUntilZero())

	// BLOCK 3
	//7
	fmt.Println(readPositive())

	//8
	var a1, b1 int
	fmt.Println("8. Введите два числа: ")
	fmt.Scan(&a1, &b1)
	fmt.Println(safeCompare(a1, b1))

	//BLOK 4 
	//9
	menu()

	//10
	var o, k int
	fmt.Println("10. Введите два числа: ")
	fmt.Scan(&o, &k)
	fmt.Println(gcd(o, k))
}
