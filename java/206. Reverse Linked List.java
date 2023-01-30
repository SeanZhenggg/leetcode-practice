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
    public ListNode reverseList(ListNode head) {
        ArrayList<ListNode> arrays = new ArrayList<>();
        ListNode result = new ListNode(-50000);
        while (head != null) {
            arrays.add(head);
            head = head.next;
        }

        for(int i = 0; i < arrays.size(); i++) {
            ListNode head2 = arrays.get(i);
            if(i == 0) {
                head2.next = null;
            }
            if(i > 0) {
                ListNode prev = arrays.get(i - 1);
                head2.next = prev;
            }
            if (i == arrays.size() - 1) {
                result.next = head2;
            }
        }

        return result.next;
    }
}

// second solution
class Solution {
    public ListNode reverseList(ListNode head) {
        ListNode prevNode = null;
        ListNode nextNode = null;
        while (head != null) {
            nextNode = head.next;
            head.next = prevNode;
            prevNode = head;
            head = nextNode;
        }
        return prevNode;
    }
}