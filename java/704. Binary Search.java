class Solution {
  public static int search(int[] nums, int target) {
    int left = 0, right = nums.length - 1;
    int ret = -1;
    while (left <= right) {
      int mid = (left + right) / 2;

      if (nums[mid] < target) {
        left = mid + 1;
      } else if (nums[mid] > target) {
        right = mid - 1;
      } else {
        ret = mid;
        break;
      }
    }

    return ret;
  }
}
