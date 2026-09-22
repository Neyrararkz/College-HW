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
        print("\nДолжность: Библиотекарь")

class Reader(User):
    def __init__(self, username, password, books_list=[], status="не читает",):
        super().__init__(username, password)        
        self.books_list = books_list        
        self.status = status

    def show_books_list(self):  
        if self.books_list:      
            print("\n\nВзятые книги: ")
            count = 1        
            for book in self.books_list:
                print(f"{count}. «{book["title"]}»: {book["author"]}, {book["year"]}")
                count += 1
            print("")
        else:
            print("\n\nНет взятых книг.")


    def show_info(self):
        super().show_info()
        self.show_books_list()
        print(f"Статус: {self.status}")

class PremiumReader(Reader):
    def __init__(self, username, password, books_list=[], status="не читает", story=[]):
        super().__init__(username, password, books_list, status)
        self.story = story

    def show_story(self):        
        if self.story:
            print("\nИстория: ")
            count = 1
            for str in self.story:
                print(f"{count}. {str["action"]}: «{str["book_title"]}»")
                count += 1
            print("")
        else:
            print("\nИстория пуста.")

    def delete_story(self):
        self.story = []


