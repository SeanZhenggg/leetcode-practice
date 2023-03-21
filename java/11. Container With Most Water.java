// brute force solution
class Solution {
    public int maxArea(int[] height) {
        if(height.length == 2) return calArea(height, 0, 1);
        int max = 0;

        for (int i = 0; i < height.length; i++) {
            for(int j = i + 1; j < height.length; j++) {
                max = Math.max(max, calArea(height, i, j));
            }
        }

        return max;
    }

    public int calArea(int[] height, int l, int r) {
        int h = Math.min(height[l], height[r]);
        int w = r - l;
        return w * h;
    }
}

// two pointer solution
class Solution {
    public int maxArea(int[] height) {
        if(height.length == 2) return calArea(height, 0, 1);
        int max = 0, l = 0, r = height.length - 1;

        while (l < r) {
            max = Math.max(max, calArea(height, l, r));
            if(height[l] > height[r]) r--;
            else l++;
        }

        return max;
    }

    public int calArea(int[] height, int l, int r) {
        int h = Math.min(height[l], height[r]);
        int w = r - l;
        return w * h;
    }
}

// faster two pointer solution
class Solution {
    public int maxArea(int[] height) {
        int max = 0, l = 0, r = height.length - 1;

        while (l < r) {
            int minHeight = height[l] > height[r] ? height[r] : height[l];
            max = Math.max(max, minHeight * (r - l));
            
            while(height[l] <= minHeight && l < r) l++;
            while(height[r] <= minHeight && l < r) r--;
        }

        return max;
    }
}