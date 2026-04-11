export default function ToggleThemeButton({ theme, setTheme }) {
  const toggle = () => {
    setTheme(theme === "light" ? "dark" : "light");
  };

  return (
    <button onClick={toggle} className="toggle-button">
      {theme === "light" ? "Dark" : "Light"}
    </button>
  );
}