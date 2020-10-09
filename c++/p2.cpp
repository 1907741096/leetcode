/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        ListNode* l3 = new ListNode;
        addTwoNumbersToThree(l1, l2, l3, 0);
        return l3;
    }

    void addTwoNumbersToThree(ListNode* l1, ListNode* l2, ListNode* l3, int n) {
        int n1, n2 = 0;
        if (l1 != nullptr) {
            n1 = l1->val;
            l1 = l1->next;
        }
        if (l2 != nullptr) {
            n2 = l2->val;
            l2 = l2->next;
        }
        l3->val = (n1 + n2 + n) % 10;
        if (l1 == nullptr && l2 == nullptr && (n1 + n2 + n) / 10 == 0) {
            return;
        }
        l3->next = new ListNode;
        addTwoNumbersToThree(l1, l2, l3->next, (n1 + n2 + n) / 10);
    }
};