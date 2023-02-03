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
    public ListNode removeNthFromEnd(ListNode head, int n) {
        ListNode dummy = head;
        ListNode prev = null;
        int length = 0;
        int count = 1;
        while (head != null) {
            length++;
            head = head.next;
        }
        
        if(length == 1 && n == 1) return null;

        head = dummy;
        while (head != null) {
            if (count == length - n + 1) {
                if(count == 1) return head.next;
                prev.next = head.next;
                break;
            }
            prev = head;
            head = head.next;
            count ++;
        }

        return dummy;
    }
}