#Задание 1

# prices = [1200, 3500, 800, 2100, 5000, 1700]

# print("Первые три: ", prices[:3])
# print("Последние две: ", prices[-2:])
# print("Список в обратном порядке: ", prices[::-1])
# print(f"Мин: {min(prices)} | Макс: {max(prices)} | Длина: {len(prices)}")

#Задание 2

# students = [
#   {"name": "Анна", "grade": 5},
#   {"name": "Иван", "grade": 4},
#   {"name": "Олег", "grade": 3}
# ]

# students.append({"name": "Елена", "grade": 5})
# for student in students:
#     print(f"Имя: {student["name"]}, оценка: {student["grade"]}")
# students[2]["grade"] = 4

#Задание 3

# text = input("Введите текст заметки: ")
# with open("3YEAR/Python_Advanced/3-4/notes.txt", "a", encoding="utf-8") as file:
#     file.write(text)
# with open("3YEAR/Python_Advanced/3-4/notes.txt", "r", encoding="utf-8") as file:
#     print(file.read())
    
#Задание 4

# import json 

# books = [
#     {"title": "Гарри Поттер", "year": 1997},
#     {"title": "Маленький принц", "year": 1943}
# ]

# with open("3YEAR/Python_Advanced/3-4/books.json", "w", encoding="utf-8") as file:
#     json.dump(books, file, ensure_ascii=False, indent=4)
# with open("3YEAR/Python_Advanced/3-4/books.json", "r", encoding="utf-8") as file:
#     books = json.load(file)
# for book in books:
#     print(f"Название: {book["title"]}, Год: {book["year"]}")

#Задание 5

class Film:
    def __init__(self, title, rating):
        self.title = title
        self.__rating = rating

    def get_rating(self):
        return self.__rating
    def set_rating(self, rating):
        if 0 < rating <= 10:
            self.__rating = rating

    def show_info(self):
        print(f"Название: {self.title}, Рейтинг: {self.get_rating()}") 

new_film = Film("Avengers", 8)
new_film.show_info()
new_film.set_rating(9)
new_film.show_info()
print(new_film.get_rating())