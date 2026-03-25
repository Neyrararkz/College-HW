export default function UserCard(props) {
    return (
        <div>
            <h3>{props.name}</h3>
            <p>Возраст: {props.age}</p>
            <p>Город: {props.city}</p>
            <p>
                {props.isStudent ? "Стдент" : "Не студент"}
            </p>
        </div>
    )
}