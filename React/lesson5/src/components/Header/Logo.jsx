export default function Logo({ small }) {
  return (
    <div className={small ? "logo logo-small" : "logo"}>
      LOGO
    </div>
  );
}