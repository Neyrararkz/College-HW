import java.util.Random;
import java.util.Scanner;

public class Main {

    static Random random = new Random();

    static int readInt(Scanner sc) {
        while (!sc.hasNextInt()) {
            System.out.println("Ошибка ввода. Введите число:");
            sc.next();
        }
        return sc.nextInt();
    }

    static int readInt(Scanner sc, String message) {
        System.out.print(message);
        return readInt(sc);
    }

    static void showMenu() {
        System.out.println("\n=== Mini RPG Arena ===");
        System.out.println("1) Показать героя");
        System.out.println("2) Тренировка");
        System.out.println("3) Бой");
        System.out.println("4) Отдых");
        System.out.println("5) Магазин");
        System.out.println("6) Выход");
        System.out.print("Выбор: ");
    }

    static void showHero(int hp, int energy, int power, int coins) {
        System.out.println("\n=== Герой ===");
        System.out.println("HP: " + hp);
        System.out.println("Energy: " + energy);
        System.out.println("Power: " + power);
        System.out.println("Coins: " + coins);
    }

    static int[] workout(int energy, int power) {

        if (energy < 10) {
            System.out.println("Недостаточно энергии для тренировки.");
            return new int[]{energy, power};
        }

        energy = Math.max(0, energy - 10);
        power += 5;

        System.out.println("Тренировка завершена!");
        System.out.println("Power увеличена.");

        return new int[]{energy, power};
    }

    static int[] relax(int hp, int energy) {

        hp = Math.min(100, hp + 10);
        energy = Math.min(100, energy + 15);

        System.out.println("Вы отдохнули.");
        System.out.println("+HP и +Energy");

        return new int[]{hp, energy};
    }

    static int[] shop(Scanner sc, int coins, int hp) {

        System.out.println("\n=== Магазин ===");
        System.out.println("1) Зелье лечения (+20 HP) — 50 coins");
        System.out.println("0) Назад");

        int choice = readInt(sc);

        switch (choice) {

            case 1:
                if (coins >= 50) {
                    coins -= 50;
                    hp = Math.min(100, hp + 20);
                    System.out.println("Вы выпили зелье!");
                } else {
                    System.out.println("Недостаточно монет.");
                }
                break;

            case 0:
                break;

            default:
                System.out.println("Нет такого товара.");
        }

        return new int[]{coins, hp};
    }

    static int[] fight(Scanner sc, int hp, int energy, int power, int coins) {

        int monsterHP = 80;

        System.out.println("\nМонстр появился! HP монстра: " + monsterHP);

        while (hp > 0 && monsterHP > 0) {

            System.out.println("\nВаш HP: " + hp + " | Energy: " + energy);
            int hit = readInt(sc, "С какой силой ударить? ");

            if (hit <= 0) {
                System.out.println("Сила удара должна быть больше 0.");
                continue;
            }

            if (hit > power) {
                System.out.println("Вы не можете ударить сильнее своей силы.");
                continue;
            }

            if (energy < hit) {
                System.out.println("Недостаточно энергии.");
                continue;
            }

            monsterHP = Math.max(0, monsterHP - hit);
            energy = Math.max(0, energy - hit / 2);

            System.out.println("Вы ударили на " + hit);
            System.out.println("HP монстра: " + monsterHP);

            if (monsterHP <= 0) break;

            int damage = random.nextInt(15) + 5;
            hp = Math.max(0, hp - damage);

            System.out.println("Монстр ударил на " + damage);
        }

        if (hp > 0) {
            coins += 50;
            System.out.println("Вы победили! +50 coins");
        } else {
            System.out.println("Вы проиграли...");
        }

        return new int[]{hp, energy, coins};
    }

    public static void main(String[] args) {

        Scanner sc = new Scanner(System.in);

        int hp = 100;
        int energy = 100;
        int power = 20;
        int coins = 100;

        boolean running = true;

        while (running) {

            showMenu();
            int choice = readInt(sc);

            switch (choice) {

                case 1:
                    showHero(hp, energy, power, coins);
                    break;

                case 2:
                    int[] workoutResult = workout(energy, power);
                    energy = workoutResult[0];
                    power = workoutResult[1];
                    break;

                case 3:
                    int[] fightResult = fight(sc, hp, energy, power, coins);
                    hp = fightResult[0];
                    energy = fightResult[1];
                    coins = fightResult[2];
                    break;

                case 4:
                    int[] relaxResult = relax(hp, energy);
                    hp = relaxResult[0];
                    energy = relaxResult[1];
                    break;

                case 5:
                    int[] shopResult = shop(sc, coins, hp);
                    coins = shopResult[0];
                    hp = shopResult[1];
                    break;

                case 6:
                    running = false;
                    System.out.println("Игра завершена.");
                    break;

                default:
                    System.out.println("Нет такого пункта.");
            }
        }

        sc.close();
    }
}