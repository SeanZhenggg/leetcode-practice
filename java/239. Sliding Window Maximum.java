// O(k * (n-k)) failed solution
class Solution {

  public static int[] maxSlidingWindow(int[] nums, int k) {
    int l = 0;
    int[] res = new int[nums.length - k + 1];
    int idx = 0, curMax = Integer.MIN_VALUE;
    for (int i = k - 1; i < nums.length; i++) {
      for (int j = l; j <= i; j++) {
        if (curMax < nums[j]) curMax = nums[j];
      }
      res[idx++] = curMax;
      curMax = Integer.MIN_VALUE;
      l++;
    }
    return res;
  }
}
