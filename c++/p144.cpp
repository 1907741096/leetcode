/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
#include <stack>
class Solution {
public:
    vector<int> res;
    vector<int> preorderTraversal(TreeNode* root) {
        std::stack<TreeNode*> stk;
        stk.push(root);
        while (!stk.empty()) {
            TreeNode* node = stk.top();
            stk.pop();
            if (node == nullptr) {
                continue;
            }
            res.push_back(node->val);
            stk.push(node->right);
		    stk.push(node->left);
        }
        return res;
    }

    void dfs(TreeNode* node) {
        if (node == nullptr) {
            return;
        }
        res.push_back(node->val);
        dfs(node->left);
        dfs(node->right);
    }
};