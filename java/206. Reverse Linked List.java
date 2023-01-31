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

// iterative solution
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

// recursive solution
class Solution {
    public ListNode firstHead = new ListNode(-1);
    public ListNode recursion(ListNode head) {
        if (head == null || head.next == null) {
            firstHead.next = head;
            return head;
        }
        ListNode nextNode = recursion(head.next);
        nextNode.next = head;
        head.next = null;
        return nextNode.next;
    }
    public ListNode reverseList(ListNode head) {
        recursion(head);

        return firstHead.next;
    }
}

// recursive solution from others
class Solution {
    public ListNode reverseList(ListNode head) {
        // Special case...
        if (head == null || head.next == null) return head;
        // Create a new node to call the function recursively and we get the reverse linked list...
        ListNode res = reverseList(head.next);
        // Set head node as head.next.next...
        head.next.next = head;
        //set head's next to be null...
        head.next = null;
        return res;     // Return the reverse linked list...
    }
}