import { useDispatch } from "react-redux";
import { addToCart } from "../redux/cartSlice";

export default function ProductCard({ product }) {
  const dispatch = useDispatch();

  return (
    <div className="card">
      <img src={product.thumbnail} alt={product.title} />
      <h3>{product.title}</h3>
      <p>{product.price} ₸</p>

      <button
        onClick={() => dispatch(addToCart(product))}
      >
        В корзину
      </button>
    </div>
  );
}