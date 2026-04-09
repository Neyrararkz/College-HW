import { useState } from "react"
import ProductForm from "./components/ProductForm"
import ProductList from "./components/ProductList"
import styles from "./App.module.css"

export default function App() {
  const [products, setProducts] = useState([])

  return (
    <div className={styles.app}>
      <h1 className={styles.title}>Mini Shop Cart (HW)</h1>

      <ProductForm setProducts={setProducts} />

      <ProductList
        products={products}
        setProducts={setProducts}
      />
      
    </div>
  )
}