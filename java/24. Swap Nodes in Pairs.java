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

// first solution
class Solution {
    public ListNode swapPairs(ListNode head) {
        int index = 0;
        ListNode dummy = new ListNode(-1);
        ListNode dummyHead = dummy;
        ListNode prev = dummy;

        while (head != null) {
            ListNode next = head.next;
            if(index % 2 == 1) {
                dummyHead.next = new ListNode(head.val, new ListNode(prev.val));
                dummyHead = dummyHead.next.next;
            } else {
                if(head.next == null) {
                    dummyHead.next = new ListNode(head.val);
                }
            }
            prev = head;
            head = head.next;
            index ++;
        }

        return dummy.next;
    }
}

// second solution
class Solution {
    public ListNode swapPairs(ListNode head) {
        if(head == null || head.next == null) return head;
        ListNode dummy = new ListNode(-1);
        ListNode prev = dummy;
        
        while (head != null && head.next != null) {
            prev.next = head.next;
            head.next = prev.next.next;
            prev.next.next = head;
            
            prev = head;
            head = head.next;
        } 

        return dummy.next;
    }
}