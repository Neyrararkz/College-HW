import Header from "./components/Header"
import Main from "./components/Main"
import Footer from "./components/Footer"
import UserCard from "./components/UserCard"
import Box from "./components/Box"
import ProductCard from "./components/ProductCard"

export default function App() {
  const name = "John"
  const year = new Date().getFullYear
  const isStudent = true

  return (
    <>
      <h1>Мой первый React-проект</h1>

      <p>Имя: {name}</p>
      <p>Год: {year}</p>

      <p>
        Статус: {isStudent ? "Студент" : "Не студент"}
      </p>

      <Header />
      <Main />

      <UserCard 
        name = "Aliya"
        age = {20}
        city = "Almaty"
        isStudent = {true}
      />
      <UserCard 
        name = "Boris"
        age = {25}
        city = "Astana"
        isStudent = {false}
      />
      <UserCard 
        name = "Alexandra"
        age = {19}
        city = "Moscow"
        isStudent = {true}
      />

      <Box>
        <p>Это текст внутри Box</p>
      </Box ><Box>
        <p>Другой текст</p>
      </Box >

      <ProductCard
        title="Latte"
        price="900"
        inStock={true}
      />
      <ProductCard
        title="Donat"
        price="700"
        inStock={false}
      />
      <ProductCard
        title="Sandwich"
        price="1500"
        inStock={true}
      />

      <Footer />
    </>
  )
}