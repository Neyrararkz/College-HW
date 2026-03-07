package main

import "fmt"

func main(){
	registered := false
	authorized := false
	for{
		var cmd int
		fmt.Print("Что вы хотите сделать?\n1 – Регистрация\n2 – Вход\n3 – Проверка статуса\n4 – Выход\n")
		fmt.Scan(&cmd)

		switch cmd{
		case 1:
			if(authorized == true){
				fmt.Println("Вы уже в системе.")
			} else{
				if(registered == true){
					fmt.Println("Вы уже зарегестрированы.")
				} else {
					fmt.Println("Вы зарегестрировались!")
					registered = true
				}
			}
		case 2:
			if(authorized == true){
				fmt.Println("Вы уже в системе.")
			} else{
				if(registered == false){
					fmt.Println("Сначала вы должны зарегестрироваться.")
				} else {
					fmt.Println("Вы в системе!")
					authorized = true
				}
			}
		case 3:
			if(registered == true){
				fmt.Println("Вы зарегестрированы.")
			} else {
				fmt.Println("Вы не зарегестрированы.")
			}
			if(authorized == true){
				fmt.Println("Вы в системе.")
			} else{
				fmt.Println("Вы не в системе.")
			}
		case 4:
			fmt.Println("Вы вышли.")
				return
		default:
			fmt.Println("Нельзя так")
		}
	}
	
}
