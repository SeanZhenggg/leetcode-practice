// wrong O(log(m * n)) solution, cuz for loop will take O(m * n) times, but still pass.
class Solution {
  public static boolean searchMatrix(int[][] matrix, int target) {
    int[] flatten = new int[matrix.length * matrix[0].length]; // m * n

    for (int i = 0; i < matrix.length; i++) {
      for (int j = 0; j < matrix[i].length; j++) {
        flatten[i * matrix[0].length + j] = matrix[i][j];
      }
    }

    int left = 0, right = flatten.length - 1;

    while (left <= right) {
      int mid = (left + right) / 2;

      if (target > flatten[mid]) {
        left = mid + 1;
      } else if (target < flatten[mid]) {
        right = mid - 1;
      } else {
        return true;
      }
    }
    return false;
  }
}

// correct O(log(m * n)) solution.
class Solution {
    public static boolean searchMatrix(int[][] matrix, int target) {
        int left = 0, right = matrix.length * matrix[0].length - 1;

        while (left <= right) {
            int mid = (left + right) / 2;
            int i = mid / matrix[0].length;
            int j = mid % matrix[0].length;
            
            if (target > matrix[i][j]) left = mid + 1;
            else if (target < matrix[i][j]) right = mid - 1;
            else return true;
        }
        return false;
    }
}