class Solution {
    public static int[] dailyTemperatures(int[] temperatures) {
        Stack<Integer> stack = new Stack<>();
        int[] results = new int[temperatures.length];

        for(int i = 0; i < temperatures.length; i++) {
            if (stack.empty()) {
                stack.push(i);
                continue;
            }
            if(temperatures[stack.peek()] < temperatures[i]) {
                while(!stack.empty() && temperatures[stack.peek()] < temperatures[i]) {
                    int lastTop = stack.pop();
                    results[lastTop] = i - lastTop;
                }
            }
            stack.push(i);
        }

        return results;
    }
}