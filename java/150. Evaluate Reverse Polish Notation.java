class Solution {
    public int evalRPN(String[] tokens) {
        Stack<Integer> stack = new Stack<>();
        for(int i = 0; i < tokens.length; i++) {
            try {
                int number = Integer.parseInt(tokens[i]);
                stack.push(number);
            } catch (NumberFormatException e) {
                int[] nums = new int[2];
                int j = 0;
                while (j < 2 && !stack.empty()) {
                    nums[j] = stack.pop();
                    j ++;
                }
                switch(tokens[i]) {
                    case "+":
                        stack.push(nums[1] + nums[0]);
                        break;
                    case "-":
                        stack.push(nums[1] - nums[0]);
                        break;
                    case "*":
                        stack.push(nums[1] * nums[0]);
                        break;
                    case "/":
                        stack.push(nums[1] / nums[0]);
                        break;
                    default:
                        break;
                }
            }
        }

        return stack.peek();
    }
}