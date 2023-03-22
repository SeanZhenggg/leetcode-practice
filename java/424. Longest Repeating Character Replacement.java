// sliding window w/ recalculating maxF solution
class Solution {
    public int characterReplacement(String s, int k) {
        int[] charMap = new int[26];
        int l = 0;
        int longest = 0;
        int maxF = 0;
        for(int r = 0; r < s.length(); r++) {
            charMap[s.charAt(r) - 'A']++;
            for(int i = 0; i < 26; i++) maxF = Math.max(maxF, charMap[i]);

            while((r - l + 1) - maxF > k) {
                charMap[s.charAt(l) - 'A']--;
                l++;
                for(int i = 0; i < 26; i++) maxF = Math.max(maxF, charMap[i]);
            }

            longest = Math.max(longest, r - l + 1);
        }

        return longest;
    }
}

// sliding window w/o recalculating maxF solution
class Solution {
    public int characterReplacement(String s, int k) {
        int[] charMap = new int[26];
        int l = 0;
        int longest = 0;
        int maxF = 0;
        for(int r = 0; r < s.length(); r++) {
            charMap[s.charAt(r) - 'A']++;
            
            // only to find maxF through entire array
            // no need to be re-calculating cuz it won't affect the answer
            // be-cuz k is constant,
            // if we need to get longest substring(window length),
            // we need to get the maxF longer and longer.
            maxF = Math.max(maxF, charMap[s.charAt(r) - 'A']);

            while((r - l + 1) - maxF > k) {
                charMap[s.charAt(l) - 'A']--;
                l++;
            }

            longest = Math.max(longest, r - l + 1);
        }

        return longest;
    }
}