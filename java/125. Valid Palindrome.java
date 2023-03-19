class Solution {
    public boolean isPalindrome(String s) {
        if (s.length() == 1) return true;

        String[] splitted = s.split("");
        String onlyAlpha = "";

        for(String c : splitted) {
            String lowerC = c.toLowerCase();
            int codePointC = lowerC.codePointAt(0);
            if((codePointC >= 48 && codePointC <= 57) || (codePointC >= 97 && codePointC <= 122)) {
                onlyAlpha += lowerC;
            }
        }

        if (onlyAlpha.length() <= 1) return true;

        int left = 0, right = onlyAlpha.length() - 1;

        while(left <= right) {
            if(onlyAlpha.charAt(left) != onlyAlpha.charAt(right)) return false;
            left++;
            right--;
        }

        return true;
    }
}