from abc import ABC, abstractmethod

class PaymentMethod(ABC):
    @abstractmethod 
    def pay(self, amount):
        pass

class CreditCard(PaymentMethod):
    def __init__(self, balance):
        self.balance = balance

    def pay(self, amount):
        if amount <= 0:
            print("Сумма должна быть больше 0")
            return
        if amount > self.balance:
            print("На карте недостаточно средств")
            return
        self.balance -= amount
        print(f"Списано {amount} с карты. Остаток: {self.balance}")

class PayPal(PaymentMethod):
    def pay(self, amount):
        print(f"перевод {amount} через PayPal успешно выполнен")

