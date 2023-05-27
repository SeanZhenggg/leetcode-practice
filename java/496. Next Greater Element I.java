class Solution {
    public static int[] nextGreaterElement(int[] nums1, int[] nums2) {
        if(nums2.length == 1 && nums1.length == 1) return new int[] {-1};
        HashMap<Integer, Integer> map = new HashMap<>();
        Stack<Integer> stack = new Stack<>();
        int[] ans = new int[nums1.length];
        for(int i = 0; i < nums2.length; i++) {
            while (!stack.empty() && nums2[stack.peek()] < nums2[i]) {
                int top = stack.pop();
                map.put(nums2[top], nums2[i]);
            }
            stack.push(i);
        }

        for(int j = 0; j < nums1.length; j++) {
            ans[j] = map.getOrDefault(nums1[j], -1);
        }
        return ans;
    }
}