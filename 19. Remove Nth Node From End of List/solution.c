// Time: 100%
// Memory: 96.39%

/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
struct ListNode* removeNthFromEnd(struct ListNode* head, int n) {
    struct ListNode *p = head, *np = head;
    size_t i = 0;
    // Going with p for n nodes
    for (; i < n && p != NULL; i++, p = p->next);
    
    // There are no n nodes in the list
    if (i < n)
        return head;
    // n-th from end - head of list
    if (p == NULL)
        return np->next;
    // when p reaches end, np will be n-th from end
    for (; p->next != NULL; p = p->next, np = np->next);
    // n > 1, so no need to check if np->next == NULL
    // delete n-th from end
    np->next = np->next->next;
    return head;   
}