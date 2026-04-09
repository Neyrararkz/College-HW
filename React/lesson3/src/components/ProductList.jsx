import ProductCard from "./ProductCard"
import styles from "./ProductList.module.css"

export default function ProductList({ products, setProducts }) {
    const removeProduct = (id) => {
        setProducts(prev => prev.filter(p => p.id !== id))
    }

    const incQty = (id) => {
        setProducts(prev => prev.map(p => p.id == id ? 
            {...p, qty: p.qty + 1 } : p))
    }

    const decQty = (id) => {
        setProducts(prev => prev.map(p => p.id == id && p.qty > 1 ? 
            {...p, qty: p.qty - 1 } : p))
    }

    const total = products.reduce(
        (sum, p) => sum + p.price * p.qty, 0
    )

    return (
        <div className={styles.list}>
            {products.map(product => (
                <ProductCard 
                    key={product.id}
                    product={product}
                    onRemove={removeProduct}
                    onInc={incQty}
                    onDec={decQty}
                />
            ))}
            <h3 className={styles.total}>
                Total: {total}
            </h3>
        </div>
        
    )
}