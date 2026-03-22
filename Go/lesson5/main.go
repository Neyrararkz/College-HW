package main

import "fmt"

type Course struct {
    Id    int
    Title string
    Price int
}
type Student struct {
    Name    string
    Courses []Course
}

var catalog = make(map[int]Course)
var students = make(map[string]Student)

func AddCourse(c Course) {
	catalog[c.Id] = c
}

func AddStudent(name string) {
	students[name] = Student{
		Name:    name,
		Courses: []Course{},
	}
}

func Enroll(name string, courseId int) {
	student := students[name]
	course := catalog[courseId]

	student.Courses = append(student.Courses, course)
	students[name] = student 
}

func TotalCost(name string) int {
	student := students[name]

	sum := 0
	for _, c := range student.Courses {
		sum += c.Price
	}
	return sum
}

func PrintStudent(name string) {
	student := students[name]
	fmt.Println("Student:", student.Name)
	fmt.Println("Courses:")

	for _, c := range student.Courses {
		fmt.Println(c.Title, "—", c.Price)
	}

	fmt.Println("Total:", TotalCost(name))
	fmt.Println()
}

func main() {
	AddCourse(Course{Id: 1, Title: "Go Basics", Price: 20000})
	AddCourse(Course{Id: 2, Title: "SQL", Price: 15000})
	AddCourse(Course{Id: 3, Title: "Java", Price: 18000})

	AddStudent("Ali")
	AddStudent("Amina")

	Enroll("Ali", 1)
	Enroll("Ali", 2)

	Enroll("Amina", 2)
	Enroll("Amina", 3)

	PrintStudent("Ali")
	PrintStudent("Amina")
}