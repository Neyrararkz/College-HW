import { useState } from "react"
import "./ShoppingList.css"

export default function ShoppingList() {
  const [title, setTitle] = useState("")
  const [qty, setQty] = useState("")

  const [items, setItems] = useState([])

  const [editId, setEditId] = useState(null)
  const [editTitle, setEditTitle] = useState("")
  const [editQty, setEditQty] = useState("")

  const [filter, setFilter] = useState("all")

  const addItem = () => {
    if (title.trim() === "") return

    const newItem = {
      id: Date.now(),
      title: title, 
      qty: qty ? Number(qty) : 1,
      bought: false
    }

    setItems([...items, newItem])

    setTitle("")
    setQty("")
  }

  const deleteItem = (id) => {
    setItems(items.filter(item => item.id !== id))
  }

  const toggleBought = (id) => {
    const updated = items.map(item => item.id === id ? {...item, bought: !item.bought} : item)
  
  setItems(updated)
  }

  const startEdit = (item) => {
    setEditId(item.id)
    setEditTitle(item.title)
    setEditQty(item.qty)
  }

  const saveEdit = () => {
    if(editTitle.trim() === "") return

    const updated = items.map(item => item.id === editId ? {...item, title: editTitle, qty: Number(editQty)} : item)
  
    setItems(updated)
    setEditId(null)
  }

  const filteredItems = items.filter(item => {
    if (filter === "bought") return item.bought
    if (filter === "notBought") return !item.bought

    return true
  })

  const total = items.length
  const bought = items.filter(i => i.bought).length
  const left = total - bought

  return (
    <div className="shopping">

      <h2>Shopping List</h2>
      <div className="shopping-top">
        <p>Всего: {total}</p>
        <p>Куплено: {bought}</p>
        <p>Осталось: {left}</p>
      </div>

      <div className="shopping-inputs">
        <input 
          className="title"
          placeholder="Товар"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          />
        <input 
          className="qty"
          type="number" 
          placeholder="Кол-во"
          value={qty}
          onChange={(e) => setQty(e.target.value)}
        />
        <button onClick={addItem}>Добавить</button>
      </div>

      <div className="shopping-filters">
        <button 
          className={filter === "all" ? "active" : ""}
          onClick={() => setFilter("all")}
        >Все</button>
        <button 
          className={filter === "bought" ? "active" : ""}
          onClick={() => setFilter("bought")}
        >Куплено</button>
        <button
          className={filter === "left" ? "active" : ""}
          onClick={() => setFilter("notBought")}
        >Не куплено</button>
      </div>

      <ul className="shopping-list">
        {filteredItems.map(item => (
          <li 
            key={item.id}
            className={`shopping-item ${item.bought ? "done" : ""}`}
          >
            {editId === item.id ? (
              <>
                <input
                  value={editTitle}
                  onChange={(e) => setEditTitle(e.target.value)}
                />
                <input 
                  type="number" 
                  value={editQty}
                  onChange={(e) => setEditQty(e.target.value)}
                />
                <button onClick={saveEdit}>Сохранить</button>
              </>
            ) : (
              <> 
                <div className="shopping-left">
                  <input 
                    type="checkbox" 
                    checked={item.bought}
                    onChange={() => toggleBought(item.id)}
                  />
                  <span className="shopping-title">
                    {item.title}
                  </span>
                  <span className="shopping-qty">
                    {item.qty}
                  </span>
                </div>
                <div className="shopping-actions">
                  <button onClick={() => startEdit(item)}>✏</button>                  
                  <button onClick={() => deleteItem(item.id)}>🗑</button>
                </div> 
              </>
            )}
          </li>
        ))}
      </ul>

    </div>
  )
}