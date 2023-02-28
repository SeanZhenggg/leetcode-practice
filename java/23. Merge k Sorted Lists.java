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
    private ListNode dummy = new ListNode(-1);
    public ListNode mergeKLists(ListNode[] lists) {
         for (ListNode list : lists) {
             this.merge(this.dummy.next, list);
         }
         return this.dummy.next;
    }

    public void merge(ListNode l1, ListNode l2) {
        ListNode tempHead = this.dummy;
        while (l1 != null && l2 != null) {
            if (l1.val < l2.val) {
                tempHead.next = l1;
                l1 = l1.next;
            } else {
                tempHead.next = l2;
                l2 = l2.next;
            }
            tempHead = tempHead.next;
        }

        while (l1 != null) {
            tempHead.next = l1;
            l1 = l1.next;
            tempHead = tempHead.next;
        }

        while (l2 != null) {
            tempHead.next = l2;
            l2 = l2.next;
            tempHead = tempHead.next;
        }
    }
}
