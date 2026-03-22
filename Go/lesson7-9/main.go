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
func (u *User) SetName(name string) {
	if name == "" {
		fmt.Println("Имя не должно быть пустым")
		return
	}
	u.name = name
}

func (u User) GetAge() int{
	return u.age
}
func (u *User) SetAge(age int) {
	if age < 0 {
		fmt.Println("Возраст не может быть отрицательным")
		return
	}
	u.age = age
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
	MaxStudents int
}
func (c Course) GetInfo(){
	fmt.Println("Title:", c.Title)
	fmt.Println("MaxScore:", c.MaxScore)
	fmt.Println("StudentsCount:", c.StudentsCount)
}

func (c *Course) AddStudent(){
	if c.StudentsCount >= c.MaxStudents{
		fmt.Println("Курс переполнен")
		return
	}
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
	fmt.Println(t.GetName(), "преподает", t.Subject)
}
func (t Teacher) GetInfo(){
	fmt.Println("Name:", t.GetName())
	fmt.Println("Age:", t.GetAge())
	fmt.Println("Subject:", t.Subject)
}
func (t Teacher) GradeStudent(student *Student, points int){
	student.Score += points
	if student.Score > student.MaxScore{
		student.Score = student.MaxScore
	}
	fmt.Println(t.GetName(), "оценил", student.GetName())
}

func AverageScore(students []*Student) float64{
	if len(students) == 0{
		return 0
	}
	sum := 0
	for _, s := range students{
		sum += s.Score
	}
	return float64(sum) / float64(len(students))
}
func BestStudent(students []*Student) *Student{
	if len(students) == 0{
		return nil
	}
	best := students[0]
	for _, s := range students{
		if s.Score > best.Score{
			best = s
		}
	}
	return best
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
		StudentsCount: 0,
		MaxStudents: 3,
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

	s3 := Student{
		User: NewUser("Ali", 19),
		Score: 60,
		Course: c,
	}

	t := Teacher{
		User: NewUser("Madina", 29),
		Subject: "Golang",
	}

	c.AddStudent()
	c.AddStudent()
	c.AddStudent()

	students := []*Student{&s1, &s2, &s3}

	participants := []Participant{&s1, &s2, &s3, t}

	for _, p := range participants{
		MakeGetInfo(p)
		MakeAct(p)
	}

	t.GradeStudent(&s1, 10)
	t.GradeStudent(&s2, 15)
	t.GradeStudent(&s3, 20)

	fmt.Println("Average score:", AverageScore(students))

	best := BestStudent(students)

	if best != nil{
		fmt.Println("Best student:", best.GetName(), "Score:", best.Score)
	}
}