// first solution
class Solution {

  public static int search(int[] nums, int target) {
    int left = 0, right = nums.length - 1;

    while (left <= right) {
      int mid = (left + right) / 2;

      if (nums[left] < nums[right]) {
        if (nums[mid] > target) right = mid - 1;
        else if (nums[mid] < target) left = mid + 1;
        else return mid;
      }
      else {
        if (nums[mid] > target) {
          if (target >= nums[left]) right = mid - 1;
          else {
            if (nums[mid] >= nums[left]) left = mid + 1;
            else right = mid - 1;
          }
        }
        else if (nums[mid] < target) {
          if (target >= nums[left]) {
            if (nums[mid] >= nums[left]) left = mid + 1;
            else right = mid - 1;
          } else left = mid + 1;
        }
        else {
          return mid;
        }
      }
    }

    return -1;
  }
}

// correct solution
class Solution {

  public static int search(int[] nums, int target) {
    int left = 0, right = nums.length - 1;

    while (left <= right) {
      int mid = (left + right) / 2;

      if (nums[mid] == target) return mid;

      if (nums[left] <= nums[mid]) {
        if (target < nums[left] || target > nums[mid]) left = mid + 1;
        else right = mid - 1;
      }
      else {
        if (target < nums[mid] || target > nums[right]) right = mid - 1;
        else left = mid + 1;
      }
    }

    return -1;
  }
}

class Solution {

  public static int search(int[] nums, int target) {
    int left = 0, right = nums.length - 1;

    while (left <= right) {
      int mid = (left + right) / 2;

      if (target < nums[mid]) {
        if (target >= nums[left]) right = mid - 1;
        else {
          if (nums[mid] >= nums[left]) left = mid + 1;
          else right = mid - 1;
        }
      }
      else if (target > nums[mid]) {
        if (target <= nums[right]) left = mid + 1;
        else {
          if (nums[mid] <= nums[right]) right = mid - 1;
          else left = mid + 1;
        }
      } else {
        return mid;
      }
    }

    return -1;
  }
}
