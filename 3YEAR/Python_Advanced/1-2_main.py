# 1
# name = input("Введите ваше имя: ")
# age = int(input("Введите ваш возраст: "))
# print(f"Привет, {name}! Через год тебе будет {age+1} лет")

#2
# num = int(input("Введите целое число: "))

# if num%2==0:
#     parity = "чётное"
# else:
#     parity = "нечётное"

# if num>0:
#     print("Положительное,", parity)
# elif num<0:
#     print("Отрицательное,", parity)
# else:
#     print("Ноль,", parity)

#3
# qty = int(input("Сколько товаров вы хотите добавить?: "))

# products = []
# total = 0

# for x in range(qty):
#     name = input(f"\nВведите название {x+1} товара: ")  
#     price = int(input(f"Введите цену {x+1} товара: "))

#     products.append({"name": name, "price": price})
#     total += price

# print("\nВаш список товаров: ")
# for x in products:
#     print(f"{x['name']} — {x['price']}, ", end="")

# print("Общая стоимость: ", total)

#4
# name = input("Введите ваше имя: ")
# scores = [int(x) for x in input("Введите ваши оценки от 1 до 5 (через пробел): ").split()]
# avg = sum(scores)/len(scores)
# print(f"Средняя: {avg}, Максимум: {max(scores)}, Минимум: {min(scores)}")
# if avg>=4:
#     print("Хороший результат")
# else:
#     print("Нужно повторить тему")

#5
users = {}
current_user = None

def register():
    login = input("Логин: ")
    if login in users:
        print("Пользователь с таким логином уже существует")
        return
    password = input("Придумайте пароль: ")
    users[login] = {
        "password": password,
        "todos": []
        }
    print("Регистрация успешна")

def login():
    login = input("Логин: ")
    password = input("Пароль: ")
    if login in users and users[login]["password"] == password:
        print("Авторизация успешна")
        return login
    else:
        print("Неверный логин или пароль")
        return None

def add_task(user):
    task = input("Задача: ")
    users[user]["todos"].append(task)
    print("Задача добавлена")

def show_tasks(user):
    print("Ваши задачи:")
    for i, task in enumerate(users[user]["todos"]):
        print(f"{i}. {task}")

def delete_task(user):
    show_tasks(user)
    index = int(input("Индекс задачи: "))
    if 0 <= index < len(users[user]["todos"]):
        users[user]["todos"].pop(index)
        print("Задача удалена")
    else:
        print("Неверный индекс")

while True:
    if not current_user:
        choice = int(input("\n1 зарегестрироваться\n2 — войти в аккаунт\n"))

        if choice == 1:
            register()
        elif choice == 2:
            current_user = login()
        else:
            print("Некорректный выбор")

    else:
        choice = int(input("\nМеню:\n1 — добавить задачу\n2 — показать задачи\n3 — удалить задачу\n0 — выйти из аккаунта\n"))

        if choice == 1:
            add_task(current_user)
        elif choice == 2:
            show_tasks(current_user)
        elif choice == 3:
            delete_task(current_user)
        elif choice == 0:
            current_user = None
            print("Вы вышли из аккаунта")
        else:
            print("Некорректный выбор")



