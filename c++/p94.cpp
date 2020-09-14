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

private:
    vector<int> res;

public:
    vector<int> inorderTraversal(TreeNode* root) {
         res = new vector<int>;
         midOrder(root);
         return res;
    }

    void midOrder(TreeNode* node) {
        if (node == NULL) {
            return;
        }
        midOrder(node.left);
        res.push_back(node.val);
        midOrder(node.right);
    }


};