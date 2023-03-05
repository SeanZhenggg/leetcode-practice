class Solution {
    public boolean isAnagram(String s, String t) {
        String sorted1 = Arrays.stream(s.split(""))
                .parallel()
                .sorted()
                .collect(Collectors.joining());
        String sorted2 = Arrays.stream(t.split(""))
                .parallel()
                .sorted()
                .collect(Collectors.joining());

        return sorted1.equals(sorted2);
    }
}

class Solution {
    public boolean isAnagram(String s, String t) {
        if(s.length() != t.length()) return false;

        int[] counts = new int[26];
        
        int base = "a".codePointAt(0);

        for(int i = 0; i < s.length(); i++) {
            // counts[s - 'a'] ++;
            // counts[t - 'a'] --;
            counts[s.codePointAt(i) - base] ++;
            counts[t.codePointAt(i) - base] --;
        }

        for (int i : counts) if (i != 0) return false;

        return true;
    }
}