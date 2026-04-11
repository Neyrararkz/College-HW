export default function ProductCard({ product }) {
  const desc = product.description.slice(0, 80);

  return (
    <div className="product-card">

      <img src={product.thumbnail} alt={product.title} />

      <h3>{product.title}</h3>

      <p>{desc}...</p>

      <p>
        Rating: {product.rating ? product.rating : "No rating"}
      </p>

      <p className="price">
        ${product.price}
      </p>

      <button onClick={() => alert(product.title)}>
        MORE
      </button>

    </div>
  );
}