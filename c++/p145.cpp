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
class Solution {
public:
    vector<int> res;

	vector<int> postorderTraversal(TreeNode* root) {
		std::stack<TreeNode*> stk;
		TreeNode* pre;
		while (root != nullptr || !stk.empty()) {
			if (root != nullptr) {
				stk.push(root);
				root = root->left;
			}
			else {
				root = stk.top();
				if (root->right != nullptr && root->right != pre) {
					root = root->right;
					stk.push(root);
					root = root->left;
				}
				else {
					stk.pop();
					res.push_back(root->val);
					pre = root;
					root = nullptr;
				}
			}
		}
		return res;
	}

    void backTraversal(TreeNode* root) {
        if (root == nullptr) {
            return;
        }
        backTraversal(root->left);
        backTraversal(root->right);
        res.push_back(root->val);
    }
};