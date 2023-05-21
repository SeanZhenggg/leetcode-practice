// using monotonic stack with (index, height) pair
class Solution {

  public static int largestRectangleArea(int[] heights) {
    Stack<int[]> stack = new Stack<>(); // pair (index, height)
    int start;
    int maxArea = 0;
    for (int i = 0; i < heights.length; i++) {
      start = i;
      while (!stack.empty() && heights[i] <= stack.peek()[1]) {
        int[] top = stack.pop();
        maxArea = Math.max(maxArea, (i - top[0]) * top[1]);
        start = top[0];
      }

      stack.push(new int[] { start, heights[i] });
    }

    if (!stack.empty()) {
      Iterator<int[]> iterator = stack.iterator();
      while (iterator.hasNext()) {
        int[] element = iterator.next();
        maxArea = Math.max(maxArea, (heights.length - element[0]) * element[1]);
      }
    }

    return maxArea;
  }
}

// using 2 monotonic stacks (forward min/backward min)
class Solution {
  public static int largestRectangleArea(int[] heights) {
    Stack<Integer> stack = new Stack<>();
    int[] fwdMin = new int[heights.length + 2];
    int[] bwdMin = new int[heights.length + 2];
    int[] newHeights = new int[heights.length + 2];
    int maxHeight = 0;

    for (int i = 0; i < newHeights.length; i++) {
      if (i == 0 || i == newHeights.length - 1) newHeights[i] =
        -1; else newHeights[i] = heights[i - 1];
    }

    for (int i = 1; i < newHeights.length; i++) {
      while (!stack.empty() && newHeights[stack.peek()] > newHeights[i]) {
        int top = stack.pop();
        bwdMin[top] = i;
      }
      stack.push(i);
    }

    stack.clear();

    for (int i = newHeights.length - 1; i >= 0; i--) {
      while (!stack.empty() && newHeights[stack.peek()] > newHeights[i]) {
        int top = stack.pop();
        fwdMin[top] = i;
      }
      stack.push(i);
    }

    stack.clear();

    for (int i = 1; i < newHeights.length - 1; i++) {
      maxHeight =
        Math.max(maxHeight, newHeights[i] * (bwdMin[i] - fwdMin[i] - 1));
    }

    return maxHeight;
  }
}
