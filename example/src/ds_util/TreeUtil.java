package ds_util;

import java.util.*;

public class TreeUtil {
  /**
   * This way to build a tree is compatible with the leetcode Tree Node Visualizer
   * <p>Print int the Leetcode way </p>
   * @param arr
   * @return
   */
  public static TreeNode buildTree(Integer... arr) {
    final int N = arr.length;
    if (N == 0) return null;

    TreeNode[] nodes = new TreeNode[N];
    Queue<TreeNode> kids = new LinkedList();
    for (int i = 0; i < N; ++i) {
      nodes[i] = arr[i] != null ? new TreeNode(arr[i]) : null;
      kids.offer(nodes[i]);
    }
    TreeNode root = kids.poll();
    for(TreeNode node : nodes) {
      if(node != null) {
        if(!kids.isEmpty()) node.left  = kids.poll();
        if(!kids.isEmpty()) node.right  = kids.poll();
      }
    }

    print(root);
    return root;
  }

  /**
   * This way to build a tree is based on the array representation of a complete tree.<br>
   * The above <code>buildTree</code> is more sparse than this one.
   *
   * @param arr
   * @return
   */
  public static TreeNode buildTree1(Integer[] arr) {
    final int N = arr.length;
    if (N == 0) return null;

    TreeNode[] nodes = new TreeNode[N];
    nodes[0] = new TreeNode(arr[0]);
    for (int i = 0; i < N; ++i) {
      TreeNode node = nodes[i];
      if (2 * i + 1 < N && node != null && arr[2 * i + 1] != null) {
        node.left = new TreeNode(arr[2 * i + 1]);
        nodes[2 * i + 1] = node.left;
      }
      if (2 * i + 2 < N && node != null && arr[2 * i + 2] != null) {
        node.right = new TreeNode(arr[2 * i + 2]);
        nodes[2 * i + 2] = node.right;
      }
    }

    TreeNode root = nodes[0];
    print(root);
    return root;
  }

  public static void print(TreeNode root) {

//    BTreePrinter.printNode(root);
    new TreePrinter().printPretty(root, 1, 0, System.err);
  }
}


class BTreePrinter {

  public static <T extends Comparable<?>> void printNode(TreeNode root) {
    int maxLevel = BTreePrinter.maxLevel(root);

    printNodeInternal(Collections.singletonList(root), 1, maxLevel);
  }

  private static <T extends Comparable<?>> void printNodeInternal(List<TreeNode> nodes, int level, int maxLevel) {
    if (nodes.isEmpty() || BTreePrinter.isAllElementsNull(nodes))
      return;

    int floor = maxLevel - level;
    int endgeLines = (int) Math.pow(2, (Math.max(floor - 1, 0)));
    int firstSpaces = (int) Math.pow(2, (floor)) - 1;
    int betweenSpaces = (int) Math.pow(2, (floor + 1)) - 1;

    BTreePrinter.printWhitespaces(firstSpaces);

    List<TreeNode> newNodes = new ArrayList<TreeNode>();
    for (TreeNode node : nodes) {
      if (node != null) {
        System.out.print(node.val);
        newNodes.add(node.left);
        newNodes.add(node.right);
      } else {
        newNodes.add(null);
        newNodes.add(null);
        System.out.print(" ");
      }

      BTreePrinter.printWhitespaces(betweenSpaces);
    }
    System.out.println("");

    for (int i = 1; i <= endgeLines; i++) {
      for (int j = 0; j < nodes.size(); j++) {
        BTreePrinter.printWhitespaces(firstSpaces - i);
        if (nodes.get(j) == null) {
          BTreePrinter.printWhitespaces(endgeLines + endgeLines + i + 1);
          continue;
        }

        if (nodes.get(j).left != null)
          System.out.print("/");
        else
          BTreePrinter.printWhitespaces(1);

        BTreePrinter.printWhitespaces(i + i - 1);

        if (nodes.get(j).right != null)
          System.out.print("\\");
        else
          BTreePrinter.printWhitespaces(1);

        BTreePrinter.printWhitespaces(endgeLines + endgeLines - i);
      }

      System.out.println("");
    }

    printNodeInternal(newNodes, level + 1, maxLevel);
  }

  private static void printWhitespaces(int count) {
    for (int i = 0; i < count; i++)
      System.out.print(" ");
  }

  private static <T extends Comparable<?>> int maxLevel(TreeNode node) {
    if (node == null)
      return 0;

    return Math.max(BTreePrinter.maxLevel(node.left), BTreePrinter.maxLevel(node.right)) + 1;
  }

  private static <T> boolean isAllElementsNull(List<T> list) {
    for (Object object : list) {
      if (object != null)
        return false;
    }

    return true;
  }

}
