import java.util.Scanner;

public class Main {
    public static void main(String[] args) {

        Scanner sc = new Scanner(System.in);

        System.out.print("Размер массива: ");
        int n = sc.nextInt();
        int[] arr = new int[n];
        System.out.println("Введите элементы:");
        for (int i = 0; i < n; i++) {
            arr[i] = sc.nextInt();
        }

        //1
        System.out.println("1. Массив в обратном порядке:");
        for (int i = arr.length - 1; i >= 0; i--) {
            System.out.print(arr[i] + " ");
        }

        //2
        int count = 0;
        for (int i = 0; i < arr.length; i++) {
            if (arr[i] == 0){
                count++;
            }
        }
        System.out.println("\n2. Количество нулей: " + count);

        //3
        int sum = 0;
        for (int i = 0; i < arr.length; i++) {
            sum += arr[i];
        }
        System.out.println("3. Сумма элементов: " + sum);

        //4
        int max = arr[0];
        for (int i = 0; i < arr.length; i++) {
            if (arr[i] > max) {
                max = arr[i];
            }
        }
        System.out.println("4. Максимальный элемент: " + max);

        //5
        boolean if10 = false;
        for (int i = 0; i < arr.length; i++) {
            if (arr[i] == 10) {
                if10 = true;
                break;
            }       
        }
        if (if10) {
            System.out.println("6. YES");
        } else {
            System.out.println("6. NO");
        }

        sc.close();
    }
}