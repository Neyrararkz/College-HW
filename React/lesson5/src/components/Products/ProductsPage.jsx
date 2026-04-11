import { useState, useEffect } from "react";
import SearchBar from "./SearchBar";
import ProductGrid from "./ProductGrid";

export default function ProductsPage() {
  const [products, setProducts] = useState([]);
  const [filtered, setFiltered] = useState([]);

  useEffect(() => {
    fetch("https://dummyjson.com/products")
      .then(res => res.json())
      .then(data => {
        setProducts(data.products);
        setFiltered(data.products.slice(0, 8));
      })
      .catch(() => console.log("Error loading products"));
  }, []);

  const search = (query) => {
    const result = products.filter(p =>
      p.title.toLowerCase().includes(query.toLowerCase())
    );

    setFiltered(result.slice(0, 8));
  };

  return (
    <div className="products-page">
      <h1>Products</h1>

      <SearchBar onSearch={search} />

      <ProductGrid products={filtered} />
    </div>
  );
}