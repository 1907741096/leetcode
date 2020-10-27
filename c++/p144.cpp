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
		while (root != nullptr || !stk.empty()) {
			if (root != nullptr) {
				stk.push(root);
				res.push_back(root->val);
				root = root->left;
			}
			else {
				root = stk.top();
				stk.pop();
				root = root->right;
			}
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