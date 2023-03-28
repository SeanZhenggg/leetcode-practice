// first solution
class Solution {
    public boolean checkInclusion(String s1, String s2) {
        HashMap<Character, Integer> map = new HashMap<>();
        
        for(int i = 0; i < s1.length(); i++) {
            map.merge(s1.charAt(i), 1, Integer::sum);
        }

        int l = 0;
        for(int r = 0; r < s2.length(); r++) {
            if(map.containsKey(s2.charAt(r))) {
                if(r == 0 || !map.containsKey(s2.charAt(r - 1))) {
                    l = r;
                }

                if (r - l + 1 == s1.length()) {
                    int tempL = l;
                    HashMap<Character, Integer> cloned = (HashMap<Character, Integer>) map.clone();

                    for(int i = l; i <= r; i++) {
                        cloned.put(s2.charAt(i), cloned.getOrDefault(s2.charAt(i), 0) - 1);
                    }

                    boolean isValid = true;
                    int negCount = 0;
                    for(Map.Entry<Character, Integer> entry : cloned.entrySet()) {
                        if(entry.getValue() != 0) {
                            isValid = false;
                        }

                        if(entry.getValue() < 0) {
                            negCount += entry.getValue();
                        }
                    }

                    if(isValid) return true;
                    else if (negCount != 0) l -= negCount;
                }
            }
        }

        return false;
    }
}

// using matches solution
class Solution {
    public static boolean checkInclusion(String s1, String s2) {
        if(s1.length() > s2.length()) return false;
        Map<Character, Integer> s1Map = new HashMap<>();
        Map<Character, Integer> s2Map = new HashMap<>();

        for (int i = 0; i < s1.length(); i++) {
            s1Map.merge(s1.charAt(i), 1, Integer::sum);
            s2Map.merge(s2.charAt(i), 1, Integer::sum);
        }

        int matches = 0;

        for (char i = 'a'; i <= 'z'; i++) {
            char key = (char) i;
            if (!s1Map.containsKey(key)) s1Map.put(key, 0);
            if (!s2Map.containsKey(key)) s2Map.put(key, 0);
            if (s1Map.get(key).equals(s2Map.get(key))) {
                matches += 1;
            }
        }
        if (matches == 26) return true;

        for (int i = s1.length(); i < s2.length(); i++) {

            char nextKey = s2.charAt(i);
            s2Map.put(nextKey, s2Map.get(nextKey) + 1);
            if (s1Map.get(nextKey).equals(s2Map.get(nextKey))) matches += 1;
            else if (s2Map.get(nextKey).equals(s1Map.get(nextKey) + 1)) matches -= 1;

            char prevKey = s2.charAt(i - s1.length());
            s2Map.put(prevKey, s2Map.get(prevKey) - 1); // need to place here instead of below put nextKey!!!
            if (s1Map.get(prevKey).equals(s2Map.get(prevKey))) matches += 1;
            else if (s2Map.get(prevKey).equals(s1Map.get(prevKey) - 1)) matches -= 1;

            if (matches == 26) return true;
        }
        return matches == 26;
    }
}

// using array map solution
class Solution {
    public boolean checkInclusion(String s1, String s2) {
        if (s1.length() > s2.length()) return false;

        int[] map = new int[26];
        char[] s1Chars = s1.toCharArray();
        char[] s2Chars = s2.toCharArray();

        for(char s1c: s1Chars) {
            map[s1c - 'a']--;
        }
        
        for(int i = 0; i < s1Chars.length; i++) {
            map[s2Chars[i] - 'a']++;
        }

        if(isPermutation(map)) return true;

        for(int i = s1Chars.length; i < s2Chars.length; i++) {
            int prev = i - s1Chars.length;
            int next = i;
            
            map[s2Chars[next] - 'a'] ++;
            map[s2Chars[prev] - 'a'] --;

            if(isPermutation(map)) return true;
        }
        
        return false;
    }

    public boolean isPermutation(int[] map) {
        boolean isPermutation = true;
        for(int count : map) {
            if(count != 0) isPermutation = false;
        }
        return isPermutation;
    }
}