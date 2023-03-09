class Solution {
    public int[] productExceptSelf(int[] nums) {
        int[] leftProduct = new int[nums.length];
        int[] rightProduct = new int[nums.length];

        for(int i = 0; i < nums.length; i++) {
            leftProduct[i] = i != 0 ? leftProduct[i - 1] * nums[i] : nums[i];
        }

        for(int i = nums.length - 1; i >= 0; i--) {
            rightProduct[i] = i != nums.length - 1 ? rightProduct[i + 1] * nums[i] : nums[i];
        }

        for(int i = 0; i < leftProduct.length; i++) {
            int leftPd = (i - 1 < 0 ? 1 : leftProduct[i - 1]);
            int rightPd = (i + 1 > leftProduct.length - 1 ? 1 : rightProduct[i + 1]);
            nums[i] = leftPd * rightPd;
        }
        
        return nums;
    }
}