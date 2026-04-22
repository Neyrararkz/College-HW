import { useEffect, useState } from "react";
import ProductCard from "../components/ProductCard";

export default function Catalog() {
  const [products, setProducts] = useState([]);
  const [query, setQuery] = useState("");
  const [filter, setFilter] = useState("all");

  useEffect(() => {
    fetch("https://dummyjson.com/products")
      .then(res => res.json())
      .then(data => setProducts(data.products));
  }, []);

  const searched = products.filter(p =>
    p.title.toLowerCase().includes(query.toLowerCase())
  );

  const filtered = searched.filter(p => {
    if (filter === "cheap") return p.price < 100;
    if (filter === "expensive") return p.price >= 100;
    return true;
  });

  return (
    <div className="catalog">

      <div className="search">
        <input
            placeholder="Поиск..."
            value={query}
            onChange={(e) => setQuery(e.target.value)}
        />
      </div>

      <div className="filters">
        <button onClick={() => setFilter("all")}>Все</button>
        <button onClick={() => setFilter("cheap")}>Дешевые</button>
        <button onClick={() => setFilter("expensive")}>Дорогие</button>
      </div>

      <div className="grid">
        {filtered.map(product => (
          <ProductCard key={product.id} product={product} />
        ))}
      </div>

    </div>
  );
}