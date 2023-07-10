// first solution
class Solution {
    public TreeNode invertTree(TreeNode root) {
        if(root == null) return null;

        Deque<TreeNode> queue = new LinkedList<>();
        queue.offer(root);

        while (!queue.isEmpty()) {
            TreeNode node = queue.poll();
            TreeNode tempLeft = node.left;
            node.left = node.right;
            node.right = tempLeft;

            if(node.left != null) {
                queue.offer(node.left);
            }

            if(node.right != null) {
                queue.offer(node.right);
            }
        }

        return root;
    }
}

// dfs solution
class Solution {
    public TreeNode invertTree(TreeNode root) {
        if(root == null) return null;

        TreeNode tempLeft = root.left;
        root.left = invertTree(root.right);
        root.right = invertTree(tempLeft);

        return root;
    }
}

// bfs with linked list solution
class Solution {
    public TreeNode invertTree(TreeNode root) {
        recursion(root);

        return root;
    }

    public void recursion(TreeNode root) {
        if(root == null || root.left == null && root.right == null) return;

        TreeNode temp = root.left;
        root.left = root.right;
        root.right = temp;

        recursion(root.left);
        recursion(root.right);
    }
}