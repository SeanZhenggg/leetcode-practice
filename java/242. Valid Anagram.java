// sort solution
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

// array solution
class Solution {
    public boolean isAnagram(String s, String t) {
        if(s.length() != t.length()) return false;

        int[] counts = new int[26];
        
        char base = 'a';

        for(int i = 0; i < s.length(); i++) {
            counts[s.charAt(i) - base] ++;
            counts[t.charAt(i) - base] --;
        }

        for (int i : counts) if (i != 0) return false;

        return true;
    }
}

// faster array solution
class Solution {
    public boolean isAnagram(String s, String t) {
        if(s.length() != t.length()) return false;

        int[] records=new int[26];
        for(char c:s.toCharArray()){
            records[c-'a']++;
        }
        for(char c:t.toCharArray()){
            if(records[c-'a']==0) return false;
            records[c-'a']--;
        }
        return true;
    }
}