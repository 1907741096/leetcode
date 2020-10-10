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
    ListNode* detectCycle(ListNode* head) {
        if (head == NULL || head->next == NULL) {
            return NULL;
        }
        ListNode* quick = head->next->next;
        ListNode* slow = head->next;
        while (quick != NULL) {
            if (quick == slow) {
                ListNode* ptr = head;
                while (ptr != slow) {
                    ptr = ptr->next;
                    slow = slow->next;
                }
                return ptr;
            }
            if (quick->next == NULL) {
                return NULL;
            }
            quick = quick->next->next;
            slow = slow->next;
        }
        return NULL;
    }
};