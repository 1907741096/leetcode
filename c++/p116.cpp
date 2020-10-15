/*
// Definition for a Node.
class Node {
public:
    int val;
    Node* left;
    Node* right;
    Node* next;

    Node() : val(0), left(NULL), right(NULL), next(NULL) {}

    Node(int _val) : val(_val), left(NULL), right(NULL), next(NULL) {}

    Node(int _val, Node* _left, Node* _right, Node* _next)
        : val(_val), left(_left), right(_right), next(_next) {}
};
*/
#include <queue>

class Solution {
public:
    Node* connect(Node* root) {
        if (root == nullptr) {
            return root;
        }
        int size;
        std::queue<Node*> q;
        q.push(root);
        Node* nowNode;
        while (!q.empty()) {
            size = q.size();
            for (int i = 0; i < size; i++) {
				nowNode = q.front();
				q.pop();
                if (nowNode->left != nullptr) q.push(nowNode->left);
                if (nowNode->right != nullptr) q.push(nowNode->right);
				if (i < size - 1) {
					nowNode->next = q.front();
				}
            }
        }
        return root;
    }
};