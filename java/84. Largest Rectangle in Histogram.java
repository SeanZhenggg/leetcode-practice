class Solution {
    public static int largestRectangleArea(int[] heights) {
        Stack<int[]> stack = new Stack<>(); // pair (index, height)
        int start;
        int maxArea = 0;
        for(int i = 0; i < heights.length; i++) {
            start = i;
            while(!stack.empty() && heights[i] <= stack.peek()[1]) {
                int[] top = stack.pop();
                maxArea = Math.max(maxArea, (i - top[0]) * top[1]);
                start = top[0];
            }

            stack.push(new int[]{ start, heights[i] });
        }

        if(!stack.empty()) {
            Iterator<int[]> iterator = stack.iterator();
            while(iterator.hasNext()) {
                int[] element = iterator.next();
                maxArea = Math.max(maxArea, (heights.length - element[0]) * element[1]);
            }
        }

        return maxArea;
    }
}