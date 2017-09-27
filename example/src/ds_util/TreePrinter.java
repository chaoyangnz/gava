package ds_util;

import java.io.PrintStream;
import java.util.Deque;
import java.util.Iterator;
import java.util.LinkedList;

public class TreePrinter {
  // Find the maximum height of the binary tree
  private int maxHeight(TreeNode p) {
    if (p == null) return 0;
    int leftHeight = maxHeight(p.left);
    int rightHeight = maxHeight(p.right);
    return (leftHeight > rightHeight) ? leftHeight + 1: rightHeight + 1;
  }

  // Convert an integer value to string
  private String intToString(int val) {
//    ostringstream ss;
//    ss << val;
//    return ss.str();
    return String.valueOf(val);
  }

  private String fill = " ";
  private String setw(int length, String str) {
    if(str.length() > length) return str;

    String spaces = "";
    int blanks = length - str.length();

    for (int i = 0; i < blanks; i++) {
      spaces += fill;
    }

    spaces += str;
    return spaces;
  }

  // Print the arm branches (eg, /    \ ) on a line
  private void printBranches(int branchLen, int nodeSpaceLen, int startLen, int nodesInThisLevel, Deque<TreeNode> nodesQueue, PrintStream out) {
    Iterator<TreeNode> iter = nodesQueue.iterator();
    for (int i = 0; i < nodesInThisLevel / 2; i++) {

      int w = (i == 0) ? startLen-1 : nodeSpaceLen-2;
      out.print(setw(w, ""));

      TreeNode node = iter.next();
      out.print(node != null ? "/" : " ");

      w = 2*branchLen+2;
      out.print(setw(w, ""));

      node = iter.next();
      out.print(node != null ? "\\" : " ");
//      out << ((i == 0) ? setw(startLen-1) : setw(nodeSpaceLen-2)) << "" << ((*iter++) ? "/" : " ");
//      out << setw(2*branchLen+2) << "" << ((*iter++) ? "\\" : " ");
    }
    out.println();
  }

  // Print the branches and node (eg, ___10___ )
  private void printNodes(int branchLen, int nodeSpaceLen, int startLen, int nodesInThisLevel, Deque<TreeNode> nodesQueue, PrintStream out) {
    Iterator<TreeNode> iter = nodesQueue.iterator();
    for (int i = 0; i < nodesInThisLevel; i++) {
      TreeNode node = iter.next();
      int w = (i == 0) ? startLen : nodeSpaceLen;
      out.print(setw(w, ""));

      fill = (node != null && node.left != null) ? "_" : " ";
      w = branchLen+2;
      out.print(setw(w, node != null ? intToString(node.val) : ""));

      fill = (node != null && node.right != null) ? "_" : " ";
      w = branchLen;
      out.print(setw(w, ""));

      fill = " ";

//      out << ((i == 0) ? setw(startLen) : setw(nodeSpaceLen)) << "" << ((*iter && (*iter)->left) ? setfill('_') : setfill(' '));
//      out << setw(branchLen+2) << ((*iter) ? intToString((*iter)->data) : "");
//      out << ((*iter && (*iter)->right) ? setfill('_') : setfill(' ')) << setw(branchLen) << "" << setfill(' ');
    }
//    out << endl;
    out.println();
  }

  // Print the leaves only (just for the bottom row)
  private void printLeaves(int indentSpace, int level, int nodesInThisLevel, Deque<TreeNode> nodesQueue, PrintStream out) {
    Iterator<TreeNode> iter = nodesQueue.iterator();
    for (int i = 0; i < nodesInThisLevel; i++) {
      int w = (i == 0) ? indentSpace+2 : 2*level+2;
      TreeNode node = iter.next();
      out.print(setw(w, node != null ? intToString(node.val) : ""));
//      out << ((i == 0) ? setw(indentSpace+2) : setw(2*level+2)) << ((*iter) ? intToString((*iter)->data) : "");
    }
//    out << endl;
    out.println();
  }

  // Pretty formatting of a binary tree to the output stream
// @ param
// level  Control how wide you want the tree to sparse (eg, level 1 has the minimum space between nodes, while level 2 has a larger space between nodes)
// indentSpace  Change this to add some indent space to the left (eg, indentSpace of 0 means the lowest level of the left node will stick to the left margin)
  public void printPretty(TreeNode root, int level, int indentSpace, PrintStream out) {
    int h = maxHeight(root);
    int nodesInThisLevel = 1;

    int branchLen = 2*((int)Math.pow(2.0,h)-1) - (3-level)*(int)Math.pow(2.0,h-1);  // eq of the length of branch for each node of each level
    int nodeSpaceLen = 2 + (level+1)*(int)Math.pow(2.0,h);  // distance between left neighbor node's right arm and right neighbor node's left arm
    int startLen = branchLen + (3-level) + indentSpace;  // starting space to the first node to print of each level (for the left most node of each level only)

    Deque<TreeNode> nodesQueue = new LinkedList<>();
    nodesQueue.offerLast(root);
    for (int r = 1; r < h; r++) {
      printBranches(branchLen, nodeSpaceLen, startLen, nodesInThisLevel, nodesQueue, out);
      branchLen = branchLen/2 - 1;
      nodeSpaceLen = nodeSpaceLen/2 + 1;
      startLen = branchLen + (3-level) + indentSpace;
      printNodes(branchLen, nodeSpaceLen, startLen, nodesInThisLevel, nodesQueue, out);

      for (int i = 0; i < nodesInThisLevel; i++) {
        TreeNode currNode = nodesQueue.getFirst();
        nodesQueue.pollFirst();
        if (currNode != null) {
          nodesQueue.offerLast(currNode.left);
          nodesQueue.offerLast(currNode.right);
        } else {
          nodesQueue.offerLast(null);
          nodesQueue.offerLast(null);
        }
      }
      nodesInThisLevel *= 2;
    }
    printBranches(branchLen, nodeSpaceLen, startLen, nodesInThisLevel, nodesQueue, out);
    printLeaves(indentSpace, level, nodesInThisLevel, nodesQueue, out);
  }

  public static void main(String[] args) {
    TreeNode root = new TreeNode(30);
    root.left = new TreeNode(20);
    root.right = new TreeNode(40);
    root.left.left = new TreeNode(10);
    root.left.right = new TreeNode(25);
//    root.right.left = new TreeNode(35);
    root.right.right = new TreeNode(50);
    root.left.left.left = new TreeNode(5);
    root.left.left.right = new TreeNode(15);
    root.left.right.right = new TreeNode(28);
    root.right.right.left = new TreeNode(41);

//    System.out.println("Tree pretty print with level=1 and indentSpace=0\n\n");
    // Output to console
    new TreePrinter().printPretty(root, 1, 0, System.err);

//    System.out.println("\n\nTree pretty print with level=5 and indentSpace=3,\noutput to file \"tree_pretty.txt\".\n\n");
    // Create a file and output to that file
//    ofstream fout("tree_pretty.txt");
    // Now print a tree that's more spread out to the file
//    printPretty(root, 5, 0, fout);
  }
}
