import { useState } from "react"
import styles from "./ProductForm.module.css"


export default function ProductForm({ setProducts }) {
    const [form, setForm] = useState({
        name: "",
        price: "",
        qty: 1
    })

    const handleChange = (e) => {
        setForm({
            ...form,
            [e.target.name]: e.target.value 
        })
    }

    const addProduct = () => {
        if (form.name.trim() === "") return

        const newProduct = {
            id: Date.now(),
            name: form.name,
            price: Number(form.price),
            qty: Number(form.qty)
        }

        setProducts(prev => [newProduct, ...prev])

        setForm({
            name: "",
            price: "",
            qty: 1
        })
    }

    return (
        <div className={styles.form}>
            <input  className={styles.input}
                name="name"
                type="text" 
                placeholder="Product name"
                value={form.name}
                onChange={handleChange}
            />
            <input  className={styles.input}
                name="price"
                type="number" 
                placeholder="Price"
                value={form.price}
                onChange={handleChange}
            />
            <input  className={styles.input}
                name="qty"
                type="number" 
                placeholder="Qty"
                value={form.qty}
                onChange={handleChange}
            />
            <button className={styles.addBtn} onClick={addProduct}>Add</button>
        </div>
    )
}