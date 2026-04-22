import { useSelector, useDispatch } from "react-redux";
import { removeFromCart, clearCart } from "../redux/cartSlice";

export default function Cart() {
  const items = useSelector((state) => state.cart.items);
  const dispatch = useDispatch();

  return (
    <div className="cart">
      <h2>Корзина ({items.length})</h2>

      {items.map((item) => (
        <div key={item.id}>
          {item.title}
          <button
            onClick={() =>
              dispatch(removeFromCart(item.id))
            }
          >
            ❌
          </button>
        </div>
      ))}

      <button onClick={() => dispatch(clearCart())}>
        Очистить корзину
      </button>
    </div>
  );
}