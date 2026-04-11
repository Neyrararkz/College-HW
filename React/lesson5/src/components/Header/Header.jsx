import Logo from "./Logo";
import Nav from "./Nav";
import ToggleThemeButton from "./ToggleThemeButton";

export default function Header({ theme, setTheme }) {
  return (
    <header className="header">
      <Logo />

      <Nav />

      <div className="header-right">
        <Logo small />
        <ToggleThemeButton theme={theme} setTheme={setTheme} />
      </div>
    </header>
  );
}