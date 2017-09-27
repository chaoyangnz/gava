import ds_util.TreeNode;
import ds_util.TreeNode.Pair;
import ds_util.TreeUtil;

import java.util.ArrayList;
import java.util.List;
import java.util.Stack;

/**
 * https://leetcode.com/problems/binary-tree-inorder-traversal
 */
public class TreeInOrderTraverse {
    public List<Integer> inorderTraversal(TreeNode root) {
        List<Integer> list = new ArrayList<>();

        Stack<Pair> stack = new Stack<>();
        stack.push(new Pair(root, false));
        while (!stack.empty()) {
            Pair current = stack.pop();
            if(current.node == null) continue;
            if(current.visited) {
                list.add(current.node.val);
            } else {
                // other traversal: change this order
                stack.push(new Pair(current.node.right, false));
                stack.push(new Pair(current.node, true));
                stack.push(new Pair(current.node.left, false));
            }
        }
        return list;
    }

    private static TreeInOrderTraverse solution = new TreeInOrderTraverse();
    public static void main(String[] args) {
        solution.test1();
        solution.test2();
        solution.test3();
    }

    //	   			3
//			   / \
//			   9  20
//			     /  \
//			    15   7
    public void test1() {
        Integer[] arr = {3,9,20,null,null,15,7};
        System.out.println(solution.inorderTraversal(TreeUtil.buildTree(arr)));
    }

    //	   			1
//			   / \
//			  2   3
//			 /     \
//			4       5
    public void test2() {
        Integer[] arr = {1,2,3,4,null,null,5};
        System.out.println(solution.inorderTraversal(TreeUtil.buildTree(arr)));
    }

    public void test3() {
        Integer[] arr = {1,2,3,null,4,5,6,null,null,7,8,null,null,9};
        System.out.println(solution.inorderTraversal(TreeUtil.buildTree(arr)));
    }
}
