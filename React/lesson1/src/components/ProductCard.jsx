export default function ProductCard({ title, price, inStock }) {
    return (
        <div>
            <h3>{title}</h3>
            <p>Цена: {price} T</p>
            <p>
                {inStock ? "В наличии" : "Нет в наличии"}
            </p>
        </div>
    )
}