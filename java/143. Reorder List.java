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

// stack solution
class Solution {
    public void reorderList(ListNode head) {
        Stack<ListNode> st = new Stack<>();

        ListNode curr = head;
        while (curr != null) {
            st.push(curr);
            curr = curr.next;
        }

        curr = head;
        int size = st.size() / 2;

        for (int i = 0; i < size; i++) {
            ListNode next = head.next;
            ListNode popped = st.pop();
            head.next = popped;
            popped.next = next;
            head = next;
        }
        head.next = null;
    }
}

// two-pointer with merge op solution

// recursive solution 1
class Solution {
    public void reorderList(ListNode head) {
        if (head.next == null || head.next.next == null) return;

        ListNode tempHead = head;
        while (tempHead.next.next != null) {
            tempHead = tempHead.next;
        }


        ListNode penultimate = tempHead;
        ListNode next = head.next;
        head.next = penultimate.next;
        penultimate.next.next = next;
        penultimate.next = null;

        reorderList(next);
    }
}

// recursive solution 2
