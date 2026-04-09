import styles from "./ProductCard.module.css"

export default function ProductCard({ product, onRemove, onInc, onDec}) {
    return (
        <div className={styles.card}>
            <h3 className={styles.name}>{product.name}</h3>
            <p className={styles.price}>Price: {product.price}</p>
            <div className={styles.qtyControls}>               
                <button onClick={() => onDec(product.id)}>-</button>
                <span>{product.qty}</span>
                <button onClick={() => onInc(product.id)}>+</button>
            </div>
            <button className={styles.deleteBtn} onClick={() => onRemove(product.id)}>🗑</button>
        </div>
    )
}