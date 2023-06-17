class Solution {
    public static int minSubArrayLen(int target, int[] nums) {
        int min = Integer.MAX_VALUE;
        int l = 0, sum = 0;

        for(int r = 0; r < nums.length; r++) {
            sum += nums[r];

            while(sum >= target && l <= r) {
                min = Math.min(min, r - l + 1);
                sum -= nums[l++];
            }
        }


        return min == Integer.MAX_VALUE ? 0 : min;
    }
}