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
		vector<int> res;
		if (root == nullptr) {
			return res;
		}

		stack<TreeNode*> stk;
		TreeNode* prev = nullptr;
		while (root != nullptr || !stk.empty()) {
			while (root != nullptr) {
				stk.emplace(root);
				root = root->left;
			}
			root = stk.top();
			stk.pop();
			if (root->right == nullptr || root->right == prev) {
				res.emplace_back(root->val);
				prev = root;
				root = nullptr;
			}
			else {
				stk.emplace(root);
				root = root->right;
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