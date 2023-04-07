class Solution {
    public static List<String> generateParenthesis(int n) {
        List<String> result = new ArrayList<>();
        backtracking(n, "", result);

        return result;
    }

    public static void backtracking(int n, String curr, List<String> result) {
        if(curr.length() >= n * 2) {
            List<String> stack = new ArrayList<>();
            for(int i = 0; i < curr.length(); i++) {
                if(curr.charAt(i) == '(') stack.add("(");
                else if (stack.size() == 0) return;
                else {
                    String popped = stack.get(stack.size() - 1);
                    if(popped != "(") return;
                    stack.remove(stack.size() - 1);
                }
            }
            if (!stack.isEmpty()) return;
            result.add(curr);
            return;
        }

        for(int i = 0; i < 2; i++) {
            backtracking(n, i == 0 ? curr + "(" : curr + ")", result);
        }
    }
}