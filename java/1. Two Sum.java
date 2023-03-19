class Solution {
    public int[] twoSum(int[] nums, int target) {
        int[] result = new int[2];
        HashMap<Integer, Integer> hashmap = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            if(hashmap.containsKey(nums[i]) && hashmap.get(nums[i]) != i) {
                result[0] = hashmap.get(nums[i]);
                result[1] = i;
                break;
            } else {
                hashmap.put(target - nums[i], i);
                // test
            }
        }
        return result;
    }
}