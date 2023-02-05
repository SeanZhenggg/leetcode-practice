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

// two pointer solution
class Solution {
    public ListNode removeNthFromEnd(ListNode head, int n) {
        ListNode dummy = head;
        ListNode fast = dummy;
        ListNode slow = dummy;
        
        for(int i = 1; i <= n; i++) {
            fast = fast.next;
        }

        if(fast == null) return slow.next;
        
        while (fast != null && fast.next != null) {
            slow = slow.next;
            fast = fast.next;
        }

        slow.next = slow.next.next;

        return dummy;
    }
}

// cool solution
class Solution {
    public ListNode removeNthFromEnd(ListNode head, int n) {
        ListNode dummy = new ListNode(-1);
        ListNode slow = dummy;
        dummy.next = head;
        
        for(int i = 1; head != null; head = head.next, i++) {
            if(i > n) slow = slow.next;
        }

        slow.next = slow.next.next;

        return dummy.next;
    }
}