// O(n * k) failed solution
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

// O(n) monotonic deque solution
class Solution {

  public static int[] maxSlidingWindow(int[] nums, int k) {
    int l = 0, idx = 0;
    int[] res = new int[nums.length - k + 1];
    ArrayDeque<Integer> queue = new ArrayDeque<>();
    for (int r = 0; r < nums.length; r++) {
      while (
        !queue.isEmpty() && nums[queue.getLast()] < nums[r]
      ) queue.pollLast();
      queue.offerLast(r);

      if (l > queue.peekFirst()) {
        queue.pollFirst();
      }

      if (r + 1 >= k) {
        res[idx++] = nums[queue.peekFirst()];
        l++;
      }
    }

    return res;
  }
}

// O(n) dp solution
class Solution {

  public static int[] maxSlidingWindow(int[] nums, int k) {
    int[] max_left = new int[nums.length];
    int[] max_right = new int[nums.length];

    max_left[0] = nums[0];
    max_right[nums.length - 1] = nums[nums.length - 1];

    for (int i = 1, j = nums.length - i - 1; i < nums.length; i++, j--) {
      max_left[i] = (i % k == 0) ? nums[i] : Math.max(max_left[i - 1], nums[i]);
      max_right[j] =
        (j % k == (k - 1)) ? nums[j] : Math.max(max_right[j + 1], nums[j]);
    }

    int[] sliding_max = new int[nums.length - k + 1];
    for (int i = 0, j = i + k - 1; j < nums.length; i++, j++) {
      sliding_max[i] = Math.max(max_right[i], max_left[j]);
    }

    return sliding_max;
  }
}
