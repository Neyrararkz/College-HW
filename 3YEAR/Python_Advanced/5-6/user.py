class User:
    def __init__(self, username, password):
        self.username = username
        self.__password = password

    def get_password(self):
        return self.__password

    def set_password(self, password):
        if len(password) >= 4:
            self.__password = password

    def check_password(self, password):
        if self.__password == password:
            return True

    def show_info(self):
        print(f"\nИмя: {self.username}", end="")

class Librarian(User):
    def __init__(self, username, password):
        super().__init__(username, password)

    def show_info(self):
        super().show_info()
        print("Должность: Библиотекарь")

class Reader(User):
    def __init__(self, username, password, books_list=[], status="не читает",):
        super().__init__(username, password)        
        self.books_list = books_list        
        self.status = status

    def show_books_list(self):        
        print("\n\nВзятые книги: ")
        count = 1
        if self.books_list:
            for book in self.books_list:
                print(f"{count}. «{book["title"]}»: {book["author"]}, {book["year"]}")
                count += 1
            print("")
        else:
            print("Список пуст.\n")


    def show_info(self):
        super().show_info()
        self.show_books_list()
        print("Статус: ", self.status)

class PremiumReader(Reader):
    def __init__(self, username, password, books_list=[], status="не читает", story=[]):
        super().__init__(username, password, books_list, status)
        self.story = story

    def show_story(self):
        print("\nИстория: ")
        count = 1
        if self.story:
            for str in self.story:
                print(f"{count}. {str["action"]}: «{str["book_title"]}»")
                count += 1
            print("")
        else:
            print("История пуста.\n")

    def delete_story(self):
        self.story = []


