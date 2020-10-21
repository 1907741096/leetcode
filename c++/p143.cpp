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
	void reorderList(ListNode* head) {
		if (head == nullptr) {
			return;
		}
		ListNode* midNode = getMidNode(head);
		ListNode* l2 = midNode->next;
		midNode->next = nullptr;
		l2 = reverseList(l2);
		head = mergeList(head, l2);
	}

	ListNode* getMidNode(ListNode* head) {
		ListNode* slow = head;
		ListNode* fast = head;
		while (fast->next != nullptr && fast->next->next != nullptr) {
			slow = slow->next;
			fast = fast->next->next;
		}
		return slow;
	}

	ListNode* reverseList(ListNode* node) {
		ListNode* lastNode = nullptr;
		ListNode* nextNode;
		while (node != nullptr) {
			nextNode = node->next;
			node->next = lastNode;
			lastNode = node;
			node = nextNode;
		}
		return lastNode;
	}

	ListNode* mergeList(ListNode* l1, ListNode* l2) {
		ListNode* head = l1;
		ListNode* nextNode;
		while (l1 != nullptr && l2 != nullptr) {
			nextNode = l1->next;
			l1->next = l2;
			l2 = l2->next;
			l1->next->next = nextNode;
			l1 = l1->next->next;
		}
		return head;
	}
};