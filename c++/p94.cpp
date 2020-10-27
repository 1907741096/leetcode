using namespace std;

struct TreeNode {
    int val;
    TreeNode* left;
    TreeNode* right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

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
    vector<int> res;
	vector<int> inorderTraversal(TreeNode* root) {
		std::stack<TreeNode*> stk;
		while (root != nullptr || !stk.empty()) {
			if (root != nullptr) {
				stk.push(root);
				root = root->left;
			}
			else {
				root = stk.top();
				stk.pop();
				res.push_back(root->val);
				root = root->right;
			}
		}
		return res;
	}

    void midOrder(TreeNode* node) {
        if (node == NULL) {
            return;
        }
        midOrder(node->left);
        res.push_back(node->val);
        midOrder(node->right);
    }


};