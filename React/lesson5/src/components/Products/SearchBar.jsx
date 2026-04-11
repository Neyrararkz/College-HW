import { useState } from "react";

export default function SearchBar({ onSearch }) {
  const [text, setText] = useState("");

  const handleSearch = () => {
    onSearch(text);
  };

  return (
    <div className="search">
      <input
        value={text}
        onChange={(e) => setText(e.target.value)}
        placeholder="Search products..."
      />

      <button onClick={handleSearch}>
        Search
      </button>
    </div>
  );
}