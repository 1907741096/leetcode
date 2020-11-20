/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */

class Solution {
public:
    ListNode* insertionSortList(ListNode* head) {
        if (head == nullptr) {
            return head;
        }
        ListNode *newHead = new ListNode;
        newHead->next = head;
        ListNode *node = head->next;
        head->next = nullptr;
        ListNode* nowNode;
        ListNode* temp;
        bool flag = false;
        while (node != nullptr) {
            flag = false;
			nowNode = newHead;
            while (nowNode->next != nullptr) {
                if (nowNode->next->val <= node->val) {
                    nowNode = nowNode->next;
                }
                else {
                    flag = true;
                    temp = node;
                    node = node->next;
					temp->next = nowNode->next;
                    nowNode->next = temp;
                    break;
                }
            }
            if (!flag) {
                nowNode->next = node;
                nowNode->next->next = nullptr;
                node = node->next;
            }
        }
        return newHead->next;
    }
};