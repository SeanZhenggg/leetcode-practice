class Solution {
    public int lengthOfLongestSubstring(String s) {
        int l = 0, r = 0;
        int longest = 0;
        for(int i = 0; i < s.length(); i++) {
            r = i;
            int tempL = l;
            while(tempL < r && s.charAt(tempL) != s.charAt(r)) tempL++;
            if(tempL != r) l = tempL + 1;
            longest = Math.max(longest, r - l + 1);
        }
        return longest;
    }
}