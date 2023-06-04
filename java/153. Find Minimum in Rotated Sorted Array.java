class Solution {
    public static int findMin(int[] nums) {
        int left = 0, right = nums.length - 1;
        int res = Integer.MAX_VALUE;

        while (left <= right) {
            if(nums[left] <= nums[right]) {
                res = Math.min(res, nums[left]);
                break;
            }

            int mid = (left + right) / 2;

            res = Math.min(res, nums[mid]);
            
            if(nums[mid] >= nums[left]) {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }
        return res;
    }
}

class Solution {
    public static int findMin(int[] nums) {
        int left = 0, right = nums.length - 1;

        if(nums[left] < nums[right]) return nums[left];

        while (left < right) {
            int mid = (left + right) / 2;
            
            if(nums[mid] >= nums[right]) {
                left = mid + 1;
            } else {
                right = mid;
            }
        }

        return nums[right];
    }
}