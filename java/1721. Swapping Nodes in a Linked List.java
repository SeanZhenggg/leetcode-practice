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

// dummy first solution🥲
class Solution {
    public ListNode swapNodes(ListNode head, int k) {
        ListNode dummy = new ListNode(-1);
        ListNode prev1 = dummy;
        ListNode prev2 = dummy;
        dummy.next = head;
        int length = 0;
        int count = 1;

        while (head != null) {
            head = head.next;
            length++;
        }
        head = dummy.next;

        while (count < k) {
            prev1 = prev1.next;
            head = head.next;
            count ++;
        }

        head = dummy.next;
        count = 1;

        while (count < (length - k + 1)) {
            prev2 = prev2.next;
            head = head.next;
            count ++;
        }

        int tempVal = prev1.next.val;
        prev1.next.val = prev2.next.val;
        prev2.next.val = tempVal;

        return dummy.next;
    }
}

// two-pointer better solution
class Solution {
    public ListNode swapNodes(ListNode head, int k) {
        ListNode dummy = new ListNode(-1);
        ListNode fast = head;
        ListNode slow = head;
        dummy.next = head;
        for (int i = 1; i < k; i++) fast = fast.next;
        ListNode tempFast = fast;

        while (fast.next != null) {
            fast = fast.next;
            slow = slow.next;
        }

        int tempFastVal = tempFast.val;
        tempFast.val = slow.val;
        slow.val = tempFastVal;
        
        return dummy.next;
    }
}