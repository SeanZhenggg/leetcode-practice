/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {
    public int carry = 0;
    public ListNode dummy = null;
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        dummy = l1;
        while (l1 != null || l2 != null) {
            l1.val = calculation(l1, l2);
            if (l1.next == null && l2.next == null) break;
            else if (l1.next == null && l2.next != null) l1.next = new ListNode(0);
            else if (l1.next != null && l2.next == null) l2.next = new ListNode(0);
            l1 = l1.next;
            l2 = l2.next;
        }

        if (carry > 0) l1.next = new ListNode(carry);

        return dummy;
    }

    public int calculation(ListNode l1, ListNode l2) {
        boolean isCarry = (l1.val + l2.val + carry) >= 10;
        int val = isCarry ? (l1.val + l2.val + carry - 10) : (l1.val + l2.val + carry);
        carry = isCarry ? 1 : 0;
        return val;
    }
}