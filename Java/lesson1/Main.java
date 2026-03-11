// import java.util.Scanner;

// public class Main {
//     public static void main(String[] args) {
//         Scanner sc = new Scanner(System.in);
//         int choice = -1;
//         int operationCount = 0;
//         double lastResult = 0;

//         while (choice != 0) {
//             System.out.println("===== CALC MENU =====");
//             System.out.println("1 — Сложение (+)");
//             System.out.println("2 — Вычитание (-)");
//             System.out.println("3 — Умножение (*)");
//             System.out.println("4 — Деление (/)");
//             System.out.println("5 — Остаток от деления (%)");
//             System.out.println("6 — Степень (a^b)");
//             System.out.println("7 — Среднее из N чисел");
//             System.out.println("8 — Показать статистику");
//             System.out.println("0 — Выход");
//             System.out.println("=====================");
            
//             if (!sc.hasNextInt()) {
//                 System.out.println("Ошибка: нужно ввести номер пункта.");
//                 sc.next();
//                 continue;
//             }
//             choice = sc.nextInt();    

//             double a;
//             double b;
//             double result;

//             switch (choice) {
//                 case 1:
//                     System.out.println("Первое число: ");
//                     a = sc.nextDouble();
//                     System.out.println("Второе число: ");
//                     b = sc.nextDouble();

//                     result = a + b;
//                     System.out.println("Результат: " + result);
//                     operationCount++;
//                     lastResult = result;
//                     break;
//                 case 2:
//                     System.out.println("Первое число: ");
//                     a = sc.nextDouble();
//                     System.out.println("Второе число: ");
//                     b = sc.nextDouble();

//                     result = a - b;
//                     System.out.println("Результат: " + result);
//                     operationCount++;
//                     lastResult = result;
//                     break;
//                 case 3:
//                     System.out.println("Первое число: ");
//                     a = sc.nextDouble();
//                     System.out.println("Второе число: ");
//                     b = sc.nextDouble();

//                     result = a * b;
//                     System.out.println("Результат: " + result);
//                     operationCount++;
//                     lastResult = result;
//                     break;
//                 case 4:
//                     System.out.println("Первое число: ");
//                     a = sc.nextDouble();
//                     System.out.println("Второе число: ");
//                     b = sc.nextDouble();

//                     if (b == 0) {
//                         System.out.println("Ошибка: деление на 0 запрещено.");
//                     } else {
//                         result = a / b;
//                         System.out.println("Результат: " + result);
//                         operationCount++;
//                         lastResult = result;
//                     }                    
//                     break;
//                 case 5:
//                     System.out.println("Первое число: ");
//                     a = sc.nextDouble();
//                     System.out.println("Второе число: ");
//                     b = sc.nextDouble();

//                     if (b == 0) {
//                         System.out.println("Ошибка: деление на 0 запрещено.");
//                     } else {
//                         result = a % b;
//                         System.out.println("Результат: " + result);
//                         operationCount++;
//                         lastResult = result;
//                     }
//                     break;
//                 case 6:
//                     System.out.println("Первое число: ");
//                     a = sc.nextDouble();
//                     System.out.println("Второе число: ");
//                     b = sc.nextDouble();

//                     result = Math.pow(a, b);
//                     System.out.println("Результат: " + result);
//                     operationCount++;
//                     lastResult = result;
//                     break;
//                 case 7:
//                     System.out.println("Сколько чисел?");
//                     int n = sc.nextInt();

//                     if (n <= 0) {
//                         System.out.println("Количество чисел должно быть больше 0.");
//                         break;
//                     }

//                     double sum = 0;
//                     for (int i = 0; i < n; i++) {
//                         System.out.println("Введите число:");
//                         double num = sc.nextDouble();
//                         sum += num;
//                     }
//                     result = sum / n;
//                     System.out.println("Среднее: " + result);
//                     operationCount++;
//                     lastResult = result;
//                     break;
//                 case 8:
//                     System.out.println("Количество операций: " + operationCount);
//                     System.out.println("Последний результат: " + lastResult);
//                     break;
//                 case 0:
//                     System.out.println("Выход.");
//                     break;
//                 default:
//                     System.out.println("Нет такого варианта.");
//             }
//         }
//         sc.close();
//     }
// }