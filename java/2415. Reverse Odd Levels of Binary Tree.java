// bfs solution
class Solution {

  public TreeNode reverseOddLevels(TreeNode root) {
    Deque<TreeNode> queue = new LinkedList<>();
    queue.offer(root);
    int level = 0;
    while (!queue.isEmpty()) {
      int size = queue.size();
      while (size > 0) {
        TreeNode node = queue.poll();
        if (node.left != null) {
          queue.offer(node.left);
        }
        if (node.right != null) {
          queue.offer(node.right);
        }
        size--;
      }
      level++;
      if (level % 2 == 1 && !queue.isEmpty()) {
        int[] nums = new int[queue.size()];
        int i = 0;
        for (TreeNode q : queue) {
          nums[i++] = q.val;
        }
        i = 0;
        for (TreeNode q : queue) {
          q.val = nums[nums.length - 1 - (i++)];
        }
      }
    }
    return root;
  }
}

// dfs solution
class Solution {
    public TreeNode reverseOddLevels(TreeNode root) {
        recursion(root.left, root.right, 1);
        return root;
    }
    
    public void recursion(TreeNode nodeLeft, TreeNode nodeRight, int level) {
        if (nodeLeft == null || nodeRight == null) return;
        
        if(level % 2 == 1) {
            int val = nodeLeft.val;
            nodeLeft.val = nodeRight.val;
            nodeRight.val = val;
        }
        
        recursion(nodeLeft.left, nodeRight.right, level + 1);
        recursion(nodeLeft.right, nodeRight.left, level + 1);
    }
}