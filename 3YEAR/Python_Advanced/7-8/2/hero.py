from abc import ABC, abstractmethod

class Hero(ABC):
    def __init__(self, name, damage):
        self.name = name
        self.damage = damage

    @abstractmethod
    def attack(self):
        pass

class Warrior(Hero):
    def attack(self):
            print(f"{self.name} рубит мечом на {self.damage} урона!")

class Mage(Hero):
    def attack(self):
            print(f"{self.name} пускает огненный шар на {self.damage} урона!")
        