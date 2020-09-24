#include <iostream>
#include <algorithm>

using namespace std;
class Solution {
public:
	bool isBalanced(TreeNode* root) {
		if (root == NULL) {
			return true;
		}
		int leftDepth = isBalancedDepth(root->left);
		int rightDepth = isBalancedDepth(root->right);
		return leftDepth >= 0 && rightDepth >= 0 && abs(double(leftDepth - rightDepth)) <= 1;
	}

	int isBalancedDepth(TreeNode* root) {
		if (root == NULL) {
			return 0;
		}
		int leftDepth = isBalancedDepth(root->left);
		int rightDepth = isBalancedDepth(root->right);
		if (abs(double(leftDepth - rightDepth)) <= 1) {
			return max(leftDepth, rightDepth) + 1;
		}
		else {
			return -1;
		}
	}
};