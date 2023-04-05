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

// O(n) monotonic deque solution
class Solution {
    public static int[] maxSlidingWindow(int[] nums, int k) {
        int l = 0, idx = 0;
        int[] res = new int[nums.length - k + 1];
        ArrayDeque<Integer> queue = new ArrayDeque<>();
        for (int r = 0; r < nums.length; r++) {
            while (!queue.isEmpty() && nums[queue.getLast()] < nums[r]) queue.pollLast();
            queue.offerLast(r);

            if (l > queue.peekFirst()) {
                queue.pollFirst();
            }

            if (r + 1 >= k) {
                res[idx++] = nums[queue.peekFirst()];
                l ++;
            }
        }

        return res;
    }
}