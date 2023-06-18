// sliding window solution
class Solution {

  public static int minSubArrayLen(int target, int[] nums) {
    int min = Integer.MAX_VALUE;
    int l = 0, sum = 0;

    for (int r = 0; r < nums.length; r++) {
      sum += nums[r];

      while (sum >= target && l <= r) {
        min = Math.min(min, r - l + 1);
        sum -= nums[l++];
      }
    }

    return min == Integer.MAX_VALUE ? 0 : min;
  }
}

// binary-search solution
class Solution {

  public int minSubArrayLen(int target, int[] nums) {
    int left = 1, right = nums.length;
    int ans = 0;
    while (left <= right) {
      int mid = (left + right) / 2;

      if (windowExist(nums, target, mid)) {
        ans = mid;
        right = mid - 1;
      } else {
        left = mid + 1;
      }
    }

    return ans;
  }

  private boolean windowExist(int[] nums, int target, int mid) {
    int sum = 0;
    for (int i = 0; i < mid; i++) sum += nums[i];

    if (sum >= target) return true;

    int l = 0;

    for (int r = mid; r < nums.length; r++) {
      sum -= nums[l++];
      sum += nums[r];

      if (sum >= target) return true;
    }

    return false;
  }
}
