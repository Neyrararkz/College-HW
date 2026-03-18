package main

import "fmt"

// 1
type User struct{
	name string
	age int
}
func NewUser(name string, age int) User{
	return User{
		name: name,
		age: age,
	}
}
func (u User) GetName() string{
	return u.name
}
func (u User) GetAge() int{
	return u.age
}
func (u User) Introduce(){
	fmt.Println("Name:", u.name)
	fmt.Println("Age:", u.age)
}

// 2
type Course struct{
	Title string
	MaxScore int
	StudentsCount int
}
func (c Course) GetInfo(){
	fmt.Println("Title:", c.Title)
	fmt.Println("MaxScore:", c.MaxScore)
	fmt.Println("StudentsCount:", c.StudentsCount)
}
func (c *Course) AddStudent(){
	c.StudentsCount++
}

// 3
type Student struct{
	User
	Score int
	Course
}
type Teacher struct{
	User
	Subject string
}

// 4
func (s *Student) Study(){
	s.Score++
	if (s.Score > s.MaxScore){
		s.Score = s.MaxScore
	}
}
func (s Student) GetInfo(){
	fmt.Println("Name:", s.GetName())
	fmt.Println("Age:", s.GetAge())
	fmt.Println("Score:", s.Score)
	fmt.Println("Course: ", s.Title)
}
func (t Teacher) Teach(){
	println(t.GetName(),"преподает", t.Subject)
}
func (t Teacher) GetInfo(){
	fmt.Println("Name:", t.GetName())
	fmt.Println("Age:", t.GetAge())
	fmt.Println("Subject:", t.Subject)
}

//5
type Participant interface{
	GetInfo()
	Act()
}
func MakeGetInfo(p Participant){
	p.GetInfo()
}
func MakeAct(p Participant){
	p.Act()
}

//6
func (s *Student) Act(){
	s.Study()
	fmt.Println(s.GetName(), "изучает материалы курса", s.Title)
}
func (t Teacher) Act(){
	t.Teach()
}

func main() {
	c := Course{
		Title: "Go",
		MaxScore: 100,
		StudentsCount: 20,
	}
	s1 := Student{
		User: NewUser("Daria", 17),
		Score: 80,
		Course: c,
	}
	s2 := Student{
		User: NewUser("Arman", 18),
		Score: 70,
		Course: c,
	}	
	t := Teacher{
		User: NewUser("Madina", 29),
		Subject: "Golang",
	}
	participants := []Participant{&s1, &s2, t}

	for _, p := range participants{
		MakeGetInfo(p)
		MakeAct(p)
	}
	for _, p := range participants{
		MakeGetInfo(p)
	}


}
