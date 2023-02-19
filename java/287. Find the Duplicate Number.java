// fast slow pointer solution
class Solution {
    public int findDuplicate(int[] nums) {
        int slow = 0;
        int fast = 0;
        do {
            slow = nums[slow];
            fast = nums[nums[fast]];
        } while (slow != fast);

        fast = 0;

        do {
            slow = nums[slow];
            fast = nums[fast];
        } while (slow != fast);

        return slow;
    }
}

// binary search solution
class Solution {
    public int findDuplicate(int[] nums) {
        int left = 1;
        int right = nums.length - 1;
        int duplicate = -1;

        while (left <= right) {
            int curr = (left + right) / 2;

            int count = 0;
            for (int i = 0; i < nums.length; i++) {
                if (nums[i] <= curr) count ++;
            }

            if (count > curr) {
                duplicate = curr;
                right = curr - 1;
            } else {
                left = curr + 1;
            }
        }

        return duplicate;
    }
}