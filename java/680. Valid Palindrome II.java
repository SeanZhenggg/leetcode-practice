class Solution {
    public boolean deleted = false;
    public boolean validPalindrome(String s) {
        return _isPalindrome(s, 0, s.length() - 1);
    }

    public boolean _isPalindrome(String s, int l, int r) {
        while(l <= r) {
            if(s.charAt(l) != s.charAt(r)) {
                if(deleted) return false;

                deleted = true;
                return _isPalindrome(s, l + 1, r) || _isPalindrome(s, l, r - 1);
            }
            else {
                l++;
                r--;   
            }
        }

        return true;
    }
}

class Solution {
    public boolean validPalindrome(String s) {
        return _isPalindrome(s, 0, s.length() - 1, 1);
    }

    public boolean _isPalindrome(String s, int l, int r, int delCnt) {
        while(l <= r) {
            if(s.charAt(l) != s.charAt(r)) {
                if(delCnt == 0) return false;

                return _isPalindrome(s, l + 1, r, delCnt - 1) || _isPalindrome(s, l, r - 1, delCnt - 1);
            }
            else {
                l++;
                r--;   
            }
        }

        return true;
    }
}