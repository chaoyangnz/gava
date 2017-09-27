import ds_util.TreeNode;
import ds_util.TreeNode.Pair;
import ds_util.TreeUtil;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;

/**
 * https://leetcode.com/problems/binary-tree-level-order-traversal
 */
public class TreeLevelOrderTraverse {

    // improved version: use one queue but record the level.
    public List<List<Integer>> levelOrder(TreeNode root) {
        List<List<Integer>> lists = new ArrayList<>();
        Queue<Pair> thisLevel = new LinkedList<>(); // all nodes are not null

        thisLevel.offer(new Pair(root, 1));
        int level = 1;
        List<Integer> list = new ArrayList<>();
        while (!thisLevel.isEmpty()) {
            Pair current = thisLevel.poll();

            if(level < current.level) { // first on in a new level
                lists.add(list);
                list = new ArrayList<>();
                level = current.level;
            }

            list.add(current.node.val);

            assert current.node != null;
            if(current.node.left != null)  thisLevel.offer(new Pair(current.node.left, current.level+1));
            if(current.node.right != null) thisLevel.offer(new Pair(current.node.right, current.level+1));
        }
        lists.add(list); // the last level

        return lists;
    }

    public List<List<Integer>> levelOrder1(TreeNode root) {
        List<List<Integer>> lists = new ArrayList<>();

        List<Integer> levelVals = new ArrayList<>(); // for one level only
        Queue<TreeNode> thisLevel = new LinkedList<>(); // all nodes are not null
        Queue<TreeNode> nextLevel = new LinkedList<>(); // all nodes are not null

        thisLevel.offer(root);
        while (!thisLevel.isEmpty()) {
            TreeNode node = thisLevel.poll();
            assert node != null;
            if(node.left != null)  nextLevel.offer(node.left);
            if(node.right != null) nextLevel.offer(node.right);
            levelVals.add(node.val);

            // transfer next level to current level
            if(thisLevel.isEmpty()) {
                lists.add(levelVals);
                levelVals = new ArrayList<>();

                while (!nextLevel.isEmpty()) {
                    thisLevel.offer(nextLevel.poll());
                }
            }
        }

        return lists;
    }

    private static TreeLevelOrderTraverse solution = new TreeLevelOrderTraverse();
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
        System.out.println(solution.levelOrder(TreeUtil.buildTree(arr)));
    }

    //	   			1
//			   / \
//			  2   3
//			 /     \
//			4       5
    public void test2() {
        Integer[] arr = {1,2,3,4,null,null,5};
        System.out.println(solution.levelOrder(TreeUtil.buildTree(arr)));
    }

    public void test3() {
        Integer[] arr = {1,2,3,null,4,5,6,null,null,7,8,null,null,9};
        System.out.println(solution.levelOrder(TreeUtil.buildTree(arr)));
    }
}
