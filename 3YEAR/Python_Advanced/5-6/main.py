from user import Librarian, Reader, PremiumReader

users = []

stand = Reader("st", "1234")
prem = PremiumReader("pr", "1234")
lib = Librarian("li", "1234")
users.append(stand)
users.append(prem)
users.append(lib)

books_catalog = [
    {
        "title": "The Hunger Games",
        "author": "Suzanne Collins",
        "year": 2008
    },    
    {
        "title": "The Maze Runner",
        "author": "James Dashner",
        "year": 2009
    },       
    {
        "title": "Dune",
        "author": "Frank Herbert",
        "year": 1965
    },    
    {
        "title": "Harry Potter",
        "author": "J.K. Rowling",
        "year": 1997
    }
]
books_onhand = []

# Регистрация & Авторизация: функции

def validate_user(username, password):
    if username.strip() == "":
        print("\nИмя пользователя не может быть пустым.")
        return
    for user in users:
        if user.username == username:
            print("\nЭто имя пользователя уже занято.")
            return
    if len(password) < 4:
        print("\nПароль должен содержать минимум 4 символа.")
        return
    return True


def register(type, username, password):
    if type == "librarian":
        user = Librarian(username, password)
    elif type == "reader":
        user = Reader(username, password, books_list=[], status="не читает")
    elif type == "premium": 
        user = PremiumReader(username, password, books_list=[], status="не читает", story=[])
    else: 
        print("Некорректный тип аккаунта.")
        return

    users.append(user)
    return True

def login(username, password):
    for user in users:
        if user.username == username:
            if user.check_password(password):
                return user

# Меню: функции
def show_books_catalog():
    if books_catalog:
        print("\nКаталог:")
        count = 1
        for book in books_catalog:
            print(f"{count}. «{book["title"]}»: {book["author"]}, {book["year"]}")
            count += 1
    else:
        print("\nКаталог пуст.")

def take_book(account, index):
    if index < 0 or index >= len(books_catalog):
        print("\nКнига не найдена.")
        return
    book_title = books_catalog[index]["title"]
    print(f"\nВы успешно взяли книгу: «{book_title}»")
    account.books_list.append(books_catalog[index])
    books_onhand.append(books_catalog[index])
    books_catalog.pop(index)
    account.status = "читает"
    return book_title

def return_book(account, index):
    if index < 0 or index >= len(account.books_list):
        print("\nКнига не найдена.")
        return
    book_title = account.books_list[index]["title"]
    print(f"\nВы успешно вернули книгу: «{book_title}»")
    books_catalog.append(account.books_list[index])
    books_onhand.remove(account.books_list[index])
    account.books_list.pop(index)
    if len(account.books_list) == 0:
        account.status = "не читает"
    return book_title

# Меню: функции библиотекаря

def show_books_onhand():
    if books_onhand:
        print("\nВыданные книги:")
        count = 1
        for book in books_onhand:
            print(f"{count}. «{book["title"]}»: {book["author"]}, {book["year"]}")
            count += 1
    else:
        print("\nНет выданных книг.")

def show_users():
    hasReader = False
    for user in users:
        if isinstance(user, Reader):
            hasReader = True
    if hasReader:
        print("\nСписок пользователей:")
        count = 1
        for user in users:
            if isinstance(user, Reader):
                print(f"{count}. {user.username}:")
                if user.books_list:
                    c = 1
                    for book in user.books_list:
                        print(f"\t{c}) «{book["title"]}»: {book["author"]}, {book["year"]}")
                        c += 1
                else:
                    print("\tНет книг на руках.")
                count += 1
    else:
        print("\nНет зарегистрированных пользователей")

# Меню: цикл
def main(account):
    while True:
        if isinstance(account, Reader):
            print("\nМеню:\n1. Каталог книг\n2. Взять книгу\n3. Вернуть книгу\n4. Моя информация\n5. Моя история\n6. Очистить историю\n0. Выход")
            choice = input("→ ")

            match choice:
                case "1":
                    show_books_catalog()
                case "2":
                    index = int(input("Номер книги, которую вы хотите взять: ")) - 1              
                    book_title = take_book(account, index)
                    if isinstance(account, PremiumReader):
                        account.story.append({"action": "Взяли книгу", "book_title": book_title})
                case "3":                     
                    index = int(input("Номер книги, которую вы хотите вернуть: ")) - 1                    
                    book_title = return_book(account, index)
                    if isinstance(account, PremiumReader):
                        account.story.append({"action": "Вернули книгу", "book_title": book_title})
                case "4":
                    account.show_info()
                case "5":
                    if not isinstance(account, PremiumReader):
                        print("\nУ вас нет доступа к этой функции. Необходимо приобрести Premium подписку.")
                        continue
                    account.show_story()
                case "6":
                    if not isinstance(account, PremiumReader):
                        print("\nУ вас нет доступа к этой функции. Необходимо приобрести Premium подписку.")
                        continue
                    print("\nВы уверены, что хотите очистить историю?(да/нет)")
                    choice = input("→ ")
                    if choice.lower().strip() == "да":
                        account.delete_story()
                        print("История успешно очищена.")
                    continue
                case "0":
                    print("До встречи!")
                    break
                case _:
                    print("Неккоректный ввод. Повторите попытку.")

        else:
            print("\nМеню:\n1. Каталог книг\n2. Книги на руках\n3. Пользователи\n4. Моя информация\n\n0. Выход")
            choice = input("→ ")

            match choice:
                case "1":
                    show_books_catalog()
                case "2":
                    show_books_onhand()
                case "3":
                    show_users()
                # case "4":
                case "5":
                    account.show_info()
                case "0":
                    print("До встречи!")
                    break
                case _:
                    print("Неккоректный ввод. Повторите попытку.")


# Регистрация & Авторизация: цикл

while True:
    print("\n1. Зарегестрироваться\n2. Войти\n0. Выйти")
    choice = input("→ ")

    match choice:
        case "1":
            print("\nВведите тип аккаунта:\n\nlibrarian — Библиотекарь\nreader — Читатель\npremium — Премиум читатель")
            type = input("→ ").lower().strip()
            username = input("Имя пользователя: ")
            password = input("Придумайте пароль: ")
            if not validate_user(username, password):
                continue
            if register(type, username, password):
                print(f"\nРегистрация прошла успешно! Добро пожаловать, {username}!")
        case "2":
            username = input("Имя пользователя: ")
            password = input("Пароль: ")
            account = login(username, password)
            if not account:
                print("\nНеверное имя пользователя или пароль")
                continue
            print(f"\nВход выполнен успешно. С возвращением, {username}!")
            main(account)
        case "0":
            print("До встречи!")
            break
        case _:
            print("Некорректный ввод. Повторите попытку.")

            
            

    


    