import Catalog from "./pages/Catalog";
import Cart from "./components/Cart";

export default function App() {
  return (
    <div className="container">
      <h1>Mini Catalog</h1>

      <Catalog />
      <Cart />
    </div>
  );
}