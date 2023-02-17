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