// stack solution
class Solution {
    public static double[] getCollisionTimes(int[][] cars) {
        double[] ans = new double[cars.length];

        for (int i = 0; i < ans.length; i++) ans[i] = -1.0;

        Stack<Integer> stack = new Stack<>();

        for(int i = cars.length - 1; i >= 0; i--) {

            while (!stack.empty() && (cars[i][1] <= cars[stack.peek()][1])) stack.pop();

            while (!stack.empty()) {
                double collisionTime = (double)(cars[stack.peek()][0] - cars[i][0]) / (cars[i][1] - cars[stack.peek()][1]);
                if(collisionTime >= ans[stack.peek()] && ans[stack.peek()] != -1.0) {
                    stack.pop();
                } else {
                    ans[i] = collisionTime;
                    break;
                }
            }

            stack.push(i);
        }

        return ans;
    }
}