package ds_util;

// Definition for a binary tree node.
public class TreeNode {
  public int val;
  public TreeNode left;
  public TreeNode right;

  public TreeNode(int x) {
    val = x;
  }

  public static class Pair {
    public TreeNode node;
    public boolean visited;


    public Pair(TreeNode node, boolean visited) {
      this.node = node;
      this.visited = visited;
    }

    public int level;
    public Pair(TreeNode node, int level) {
      this.node = node;
      this.level = level;
    }
  }
}

