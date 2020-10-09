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
    bool hasCycle(ListNode* head) {
        if (head == NULL) {
            return false;
        }
        ListNode* quick = head->next;
        ListNode* slow = head;
        while (slow != NULL && quick != NULL) {
            if (quick == slow) {
                return true;
            }
            if (quick->next == NULL) {
                return false;
            }
            quick = quick->next->next;
            slow = slow->next;
        }
        return false;
    }
};