import { useReducer } from "react";

const initialState = {
  items: [],
  form: { title: "", price: "" },
  errors: {},
  mode: "add",
  editingId: null,
};

function makeId() {
  return `${Date.now().toString(36)}-${Math.floor(Math.random() * 1e6)
    .toString(36)
    .padStart(4, "0")}`;
}

function validate(form) {
  const errors = {};

  if (!String(form.title ?? "").trim()) errors.title = "Название обязательно";

  const priceStr = String(form.price ?? "").trim();
  const priceNumber = Number(priceStr);

  if (!priceStr) errors.price = "Цена обязательна";
  else if (Number.isNaN(priceNumber)) errors.price = "Цена должна быть числом";
  else if (priceNumber <= 0) errors.price = "Цена должна быть > 0";

  return errors;
}

function reducer(state, action) {
  switch (action.type) {
    case "SET_FIELD": {
      const { field, value } = action;
      return {
        ...state,
        form: {
          ...state.form,
          [field]: value,
        },
      };
    }

    case "SET_ERRORS": {
      return { ...state, errors: action.errors };
    }

    case "RESET_FORM": {
      return { ...state, form: { title: "", price: "" }, errors: {} };
    }

    case "ADD_ITEM": {
      return {
        ...state,
        items: [action.item, ...state.items],
      };
    }

    case "DELETE_ITEM": {
      const id = action.id;
      const nextItems = state.items.filter((it) => it.id !== id);

      const wasEditing = state.editingId === id;
      return {
        ...state,
        items: nextItems,
        ...(wasEditing
          ? {
              mode: "add",
              editingId: null,
              form: { title: "", price: "" },
              errors: {},
            }
          : {}),
      };
    }

    case "START_EDIT": {
      const item = action.item;
      return {
        ...state,
        mode: "edit",
        editingId: item.id,
        form: { title: item.title, price: String(item.price) },
        errors: {},
      };
    }

    case "CANCEL_EDIT": {
      return {
        ...state,
        mode: "add",
        editingId: null,
        form: { title: "", price: "" },
        errors: {},
      };
    }

    case "SAVE_EDIT": {
      const { id, nextData } = action;
      const nextItems = state.items.map((it) =>
        it.id === id ? { ...it, ...nextData } : it
      );

      return {
        ...state,
        items: nextItems,
        mode: "add",
        editingId: null,
        form: { title: "", price: "" },
        errors: {},
      };
    }

    default:
      return state;
  }
}

export default function App() {
  const [state, dispatch] = useReducer(reducer, initialState);

  const isEdit = state.mode === "edit";
  const totalPrice = state.items.reduce((sum, it) => sum + it.price, 0);

  const submit = (e) => {
    e.preventDefault();

    const errors = validate(state.form);
    dispatch({ type: "SET_ERRORS", errors });

    if (Object.keys(errors).length > 0) return;

    const title = state.form.title.trim();
    const price = Number(state.form.price);

    if (!isEdit) {
      dispatch({
        type: "ADD_ITEM",
        item: { id: makeId(), title, price },
      });
      dispatch({ type: "RESET_FORM" });
      return;
    }

    dispatch({
      type: "SAVE_EDIT",
      id: state.editingId,
      nextData: { title, price },
    });
  };
  return (
    <div className="app">
      <form onSubmit={submit} className="form">

        <h2>
          {isEdit ? "Редактировать товар" : "Добавить товар"}
        </h2>

        <input
          type="text"
          placeholder="Название"
          value={state.form.title}
          onChange={(e) =>
            dispatch({
              type: "SET_FIELD",
              field: "title",
              value: e.target.value,
            })
          }
        />
        {state.errors.title && (
          <p className="error">{state.errors.title}</p>
        )}

        <input
          type="number"
          placeholder="Цена"
          value={state.form.price}
          onChange={(e) =>
            dispatch({
              type: "SET_FIELD",
              field: "price",
              value: e.target.value,
            })
          }
        />
        {state.errors.price && (
          <p className="error">{state.errors.price}</p>
        )}

        <button type="submit">
          {isEdit ? "Сохранить" : "Добавить"}
        </button>

        {isEdit && (
          <button
            type="button"
            onClick={() => dispatch({ type: "CANCEL_EDIT" })}
          >
            Отмена
          </button>
        )}
      </form>

      <div className="list">

        {state.items.length === 0 ? (
          <p>Список пуст</p>
        ) : (
          state.items.map((it) => (
            <div
              key={it.id}
              className={`item ${
                state.editingId === it.id ? "editing" : ""
              }`}
            >
              <span>
                {it.title} —{" "}
                {it.price.toLocaleString()} ₸
              </span>

              <div className="actions">
                <button
                  onClick={() =>
                    dispatch({ type: "START_EDIT", item: it })
                  }
                >
                  ✏
                </button>

                <button
                  onClick={() =>
                    dispatch({ type: "DELETE_ITEM", id: it.id })
                  }
                >
                  🗑
                </button>
              </div>
            </div>
          ))
        )}

      </div>

      <h3 className="total">
        Total: {totalPrice.toLocaleString()} ₸
      </h3>

    </div>
  );
}