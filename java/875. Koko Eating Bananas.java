class Solution {

  public int minEatingSpeed(int[] piles, int h) {
    int left = 1, right = Integer.MIN_VALUE;

    for (int p : piles) {
      right = Math.max(right, p);
    }

    int minHours = right;

    while (left <= right) {
      int mid = (left + right) / 2;
      int totalHours = 0;

      for (int pile : piles) {
        totalHours += Math.ceil((double) pile / mid);
      }

      if (totalHours > h) {
        left = mid + 1;
      } else {
        minHours = Math.min(mid, minHours);
        right = mid - 1;
      }
    }

    return minHours;
  }
}
