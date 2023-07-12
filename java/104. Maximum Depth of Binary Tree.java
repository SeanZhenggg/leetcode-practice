// dfs solution
class Solution {

  private int depth = 0;

  public int maxDepth(TreeNode root) {
    if (root == null) return 0;

    recursion(root, 1);

    return depth;
  }

  public void recursion(TreeNode root, int level) {
    if (root == null) return;

    depth = Math.max(depth, level);
    recursion(root.left, level + 1);
    recursion(root.right, level + 1);
  }
}

// bfs solution
class Solution {

  public int maxDepth(TreeNode root) {
    if (root == null) return 0;
    Deque<TreeNode> queue = new LinkedList<>();
    int level = 0;
    queue.offer(root);

    while (!queue.isEmpty()) {
      int size = queue.size();
      while (size-- > 0) {
        TreeNode node = queue.poll();
        if (node.left != null) queue.offer(node.left);
        if (node.right != null) queue.offer(node.right);
      }
      level++;
    }

    return level;
  }
}

// dfs another solution
class Solution {

  private int depth = 0;

  public int maxDepth(TreeNode root) {
    if (root == null) return 0;

    int leftDepth = maxDepth(root.left);
    int rightDepth = maxDepth(root.right);
    return Math.max(leftDepth, rightDepth) + 1; // the +1 is for current root in the recursion
  }
}
