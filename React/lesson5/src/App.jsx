import { useState, useEffect } from "react";
import Header from "./components/Header/Header";
import ProductsPage from "./components/Products/ProductsPage";
import "./styles/theme.css";

export default function App() {
  const [theme, setTheme] = useState("light");

  useEffect(() => {
    const saved = localStorage.getItem("theme");
    if (saved) setTheme(saved);
  }, []);

  useEffect(() => {
    document.body.className = theme;
    localStorage.setItem("theme", theme);
  }, [theme]);

  return (
    <div>
      <Header theme={theme} setTheme={setTheme} />
      <ProductsPage />
    </div>
  );
}