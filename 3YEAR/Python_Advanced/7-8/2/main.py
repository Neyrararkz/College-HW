from hero import Warrior, Mage

heroes = []
heroes.append(Warrior("Арагорн", 50))
heroes.append(Mage("Стрендж", 70))

for hero in heroes:
    hero.attack()