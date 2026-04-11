// import java.lang.reflect.Array;
// import java.util.Arrays;
// import java.util.Random;
// import java.util.Scanner;

// public class Main {
//     public static void main(String[] args) {
//         Scanner sc = new Scanner(System.in);
//         // Часть 1
//         // 1
//         for (int i=1; i<=10; i++){
//             for (int j=1; j<=10; j++){
//                 System.out.println(i + " * " + j + " = " + i*j);
//             }
//             System.out.println("");
//         }
//         // 2
//         System.out.print("n: ");
//         int n = sc.nextInt();
//         System.out.print("m: ");
//         int m = sc.nextInt();
//         for (int i=1; i<=n; i++){
//             for (int j=1; j<=m; j++){
//                 System.out.print("*");
//             }
//             System.out.println();
//         }
//         // 3
//         for (int i=1; i<=5; i++){
//             for (int j=1; j<=i; j++){
//                 System.out.print("*");
//             }
//             System.out.println();
//         }
//         // 4
//         int count = 0;
//         for (int i=1; i <=100; i+=10) {
//             for (int j=1; j<=10; j++) {
//                 int number = i + j;
//                 if (number % 2 == 0) {
//                     count++;
//                 }
//             }
//         }
//         System.out.println("Even count: " + count);
//         // 5
//         int sum = 0;
//         for (int i=1; i <=100; i++) {
//             sum += i;
//         }
//         System.out.println("Sum: " + sum);

//         // Часть 2
//         int[][] matrix = {
//             {1, 2, 3},
//             {4, 5, 6},
//             {7, 8, 9}
//         };
//         // 6        
//         for (int i = 0; i < matrix.length; i++) {
//             for (int j = 0; j < matrix[i].length; j++) {
//                 System.out.print(matrix[i][j] + "\t");
//             }
//             System.out.println();
//         }
//         // 7
//         int totalSum = 0;
//         for (int[] row : matrix) {
//             for (int val : row) {
//                 totalSum += val;
//             }
//         }
//         System.out.println("Sum: " + totalSum);
//         // 8
//         int max = matrix[0][0];
//         for (int[] row : matrix) {
//             for (int val : row) {
//                 if (val > max) max = val;
//             }
//         }
//         System.out.println("Max: " + max);
//         // 9
//         int evenCount = 0;
//         for (int[] row : matrix) {
//             for (int val : row) {
//                 if (val % 2 == 0) evenCount++;
//             }
//         }
//         System.out.println("Even Count: " + evenCount);
//         // 10
//         for (int i = 0; i < matrix.length; i++) {
//             int rowSum = 0;
//             for (int j = 0; j < matrix[i].length; j++) {
//                 rowSum += matrix[i][j];
//             }
//             System.out.print("String " + i + ": " + rowSum + "; ");
//         }
//         System.out.println();
//         // 11
//         for (int j = 0; j < matrix[0].length; j++) {
//             int colSum = 0;
//             for (int i = 0; i < matrix.length; i++) {
//                 colSum += matrix[i][j];
//             }
//             System.out.print("Column " + j + ": " + colSum + "; ");
//         }
//         System.out.println();
//         // 12
//         for (int i = 0; i < matrix.length; i++) {
//             System.out.print(matrix[i][i] + " ");
//         }
//         System.out.println();
//         // 13
//         int[][] transposed = new int[3][3];
//         for (int i = 0; i < matrix.length; i++) {
//             for (int j = 0; j < matrix[i].length; j++) {
//                 transposed[j][i] = matrix[i][j];
//             }
//         }
//         for (int[] row : transposed) {
//             for (int val : row) {
//                 System.out.print(val + "\t");
//             }
//             System.out.println();
//         }

//         // Часть 3
//         int n = 5;
//         int[][] matrix = new int[n][n];
//         Random rand = new Random();

//         // 14
//         for (int i = 0; i < n; i++) {
//             for (int j = 0; j < n; j++) {
//                 matrix[i][j] = rand.nextInt(50) + 1;
//                 System.out.print(matrix[i][j] + "\t");
//             }
//             System.out.println();
//         }

//         // 15
//         int min = matrix[0][0];
//         int minRow = 0, minCol = 0;
//         double sum = 0;

//         for (int i = 0; i < n; i++) {
//             for (int j = 0; j < n; j++) {
//                 sum += matrix[i][j];
//                 if (matrix[i][j] < min) {
//                     min = matrix[i][j];
//                     minRow = i;
//                     minCol = j;
//                 }
//             }
//         }
//         System.out.println("\nMin: " + min + " in position [" + minRow + "][" + minCol + "]");

//         // 16
//         boolean isSymmetric = true;
//         for (int i = 0; i < n; i++) {
//             for (int j = 0; j < i; j++) { 
//                 if (matrix[i][j] != matrix[j][i]) {
//                     isSymmetric = false;
//                     break;
//                 }
//             }
//         }
//         System.out.println("Is Symmetric: " + (isSymmetric ? "Yes" : "No"));

//         // 17
//         int[] temp = matrix[0];
//         matrix[0] = matrix[n - 1];
//         matrix[n - 1] = temp;

//         // 18
//         double average = sum / (n * n);
//         System.out.println("Average: " + String.format("%.2f", average));
//         for (int i = 0; i < n; i++) {
//             for (int j = 0; j < n; j++) {
//                 if (matrix[i][j] < average) {
//                     matrix[i][j] = 0;
//                 }
//                 System.out.print(matrix[i][j] + "\t");
//             }
//             System.out.println();
//         }
        
//         sc.close();
//     }
// }