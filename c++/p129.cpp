/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    int sumNumbers(TreeNode* root) {
        return sumNumber(root, 0);
    }

    int sumNumber(TreeNode* root, int n) {
        if (root == NULL) {
            return 0;
        }
        if (root->left == NULL && root->right == NULL) {
            return n * 10 + root->val;
        }
        return sumNumber(root->left, n * 10 + root->val) + sumNumber(root->right, n * 10 + root->val);
    }
};