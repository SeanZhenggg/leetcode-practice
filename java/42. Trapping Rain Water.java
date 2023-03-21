// two pointer with one record solution
class Solution {
    public int trap(int[] height) {
        int l = 0, r = height.length - 1;
        int trap = 0;
        int lastMinHeight = 0;

        while (l < r) {
            int minHeight = height[l] > height[r] ? height[r] : height[l];
            if(minHeight > lastMinHeight) lastMinHeight = minHeight;

            if(l != 0 && lastMinHeight - height[l] > 0) trap += (lastMinHeight - height[l]);
            if(r != height.length - 1 && lastMinHeight - height[r] > 0) trap += (lastMinHeight - height[r]);

            if (height[l] > height[r]) {
                r--;
            } else {
                l++;
            }
        }

        return trap;
    }
}

// two pointer with two max record solution
class Solution {
    public int trap(int[] height) {
        int l = 0, r = height.length - 1;
        int trap = 0, lMax = height[l], rMax = height[r];

        while (l < r) {
            if(lMax < rMax) {
                l++;
                lMax = Math.max(lMax, height[l]);
                trap += lMax - height[l];
            } else {
                r--;
                rMax = Math.max(rMax, height[r]);
                trap += rMax - height[r];
            }
        }

        return trap;
    }
}