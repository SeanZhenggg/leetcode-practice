class Solution {
    public List<List<Integer>> threeSum(int[] nums) {
        List<List<Integer>> result = new ArrayList<>();
        Arrays.sort(nums);

        for(int i = 0; i <= nums.length - 3; i++) {
            if (i > 0 && nums[i] == nums[i - 1]) continue;
            if(nums[i] > 0) break;

            int l = i + 1, r = nums.length - 1;
            while (l < r) {
                if ((l > i + 1 && nums[l] == nums[l - 1]) && (r < nums.length - 1 && nums[r] == nums[r + 1])) {
                    l++;
                    r--;
                    continue;
                }

                if (nums[i] + nums[l] + nums[r] == 0) {
                    result.add(new ArrayList<Integer>(Arrays.asList(nums[i], nums[l], nums[r])));
                    l++;
                    r--;
                }
                else if (nums[i] + nums[l] + nums[r] > 0) r--;
                else l++;
            }
        }
        return result;
    }
}

// better solution
class Solution {
    public List<List<Integer>> threeSum(int[] nums) {
        List<List<Integer>> result = new ArrayList<>();
        Arrays.sort(nums);

        for(int i = 0; i <= nums.length - 3; i++) {
            if (i > 0 && nums[i] == nums[i - 1]) continue;
            if(nums[i] > 0) break;

            int l = i + 1, r = nums.length - 1;
            while (l < r) {
                if (nums[i] + nums[l] + nums[r] == 0) {
                    result.add(new ArrayList<Integer>(Arrays.asList(nums[i], nums[l], nums[r])));
                    do {
                        l++;
                    }
                    while (l < r && nums[l] == nums[l - 1]);
                }

                else if (nums[i] + nums[l] + nums[r] > 0) r--;
                else l++;
            }
        }
        return result;
    }
}