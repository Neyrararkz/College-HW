package main

import "fmt"

//1
type Book struct {
    Title string
    Author string
    pages int
	isAvailable bool
}
func newBook(title string, author string, pages int) Book {
	return Book{
		Title: title,
		Author: author,
		pages: pages,
		isAvailable: true,
	}
}

func (b Book) Info() {
	fmt.Println("Title:", b.Title)
	fmt.Println("Author:", b.Author)
	fmt.Println("Pages:", b.pages)
	fmt.Println("Available:", b.isAvailable)
}
func (b *Book) Borrow() {
	if b.isAvailable {
		b.isAvailable = false
		fmt.Println("You borrowed the book:", b.Title)
	} else {
		fmt.Println("Book is not available")
	}
}
func (b *Book) ReturnBook() {
	b.isAvailable = true
	fmt.Println("Book returned:", b.Title)
}
func (b Book) GetPages() int {
	return b.pages
}
func (b *Book) SetPages(p int) {
	if p <= 0 {
		fmt.Println("Invalid number of pages")
		return
	}
	b.pages = p
}

//2
type Worker interface {
	Work() string
	GetName() string
}

type Programmer struct {
	Name string
	Language string
}
type Designer struct {
	Name string
	Tool string
}

func (p Programmer) Work() string {
	return "Программист " + p.Name + " пишет код на " + p.Language
}
func (p Programmer) GetName() string {
	return p.Name
}

func (d Designer) Work() string {
	return "Дизайнер " + d.Name + " делает макет в " + d.Tool
}

func (d Designer) GetName() string {
	return d.Name
}
func ShowWork(w Worker) {
	fmt.Println("Name:", w.GetName())
	fmt.Println(w.Work())
	fmt.Println()
}

//3
type Product struct {
	Name string
	price float64
	Quantity int
}
func newProduct(name string, price float64, quantity int) Product {
	if price < 0 {
		price = 0
	}
	if quantity < 0 {
		quantity = 0
	}
	return Product{
		Name:     name,
		price:    price,
		Quantity: quantity,
	}
}
func (p Product) GetPrice() float64 {
	return p.price
}
func (p *Product) SetPrice(newPrice float64) {
	if newPrice < 0 {
		fmt.Println("Invalid price")
		return
	}
	p.price = newPrice
}
func (p *Product) Buy(amount int) {
	if amount <= 0 {
		fmt.Println("Invalid amount")
		return
	}
	if amount > p.Quantity {
		fmt.Println("Not enough products in stock")
		return
	}
	p.Quantity -= amount
	fmt.Println("Purchased:", amount, p.Name)
}
func (p *Product) Restock(amount int) {
	if amount <= 0 {
		fmt.Println("Invalid restock amount")
		return
	}
	p.Quantity += amount
}
func (p Product) Info() {
	fmt.Println("Product:", p.Name)
	fmt.Println("Price:", p.price)
	fmt.Println("Stock:", p.Quantity)
}


func main() {
	// //1
	// book := newBook("Sherlock Holmes", "Arthur Conan Doyle", 350)

	// book.Info()

	// book.Borrow()
	// book.Borrow()

	// book.ReturnBook()

	// book.SetPages(400)

	// book.Info()

	// //2
	// p1 := Programmer{"Slava", "Go"}
	// p2 := Programmer{"Dana", "Java"}

	// d1 := Designer{"Lucy", "Figma"}
	// d2 := Designer{"John", "Photoshop"}

	// ShowWork(p1)
	// ShowWork(p2)
	// ShowWork(d1)
	// ShowWork(d2)

	// //3
	// product := newProduct("Laptop", 1500, 10)

	// product.Info()

	// product.Buy(3)
	// product.Buy(20)

	// product.SetPrice(1400)

	// product.Restock(5)

	// product.Info()
}