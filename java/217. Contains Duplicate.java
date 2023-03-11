// HashMap solution
class Solution {
    public boolean containsDuplicate(int[] nums) {
        Map<Integer, Integer> map = new HashMap<>();
        for(int i = 0; i < nums.length; i++) {
            if (map.containsKey(nums[i])) return true;
            map.put(nums[i], 1);
        }

        return false;
    }
}

// sort solution
class Solution {
    public boolean containsDuplicate(int[] nums) {
        Arrays.sort(nums);
        for(int i = 1; i < nums.length; i++) {
            if(nums[i] == nums[i - 1]) return true;
        }

        return false;
    }
}

// HashSet solution
class Solution {
    public boolean containsDuplicate(int[] nums) {
        Set<Integer> set = new HashSet<>();
        for(int i = 0; i < nums.length; i++)
            if(!set.add(nums[i])) return true;

        return false;
    }
}