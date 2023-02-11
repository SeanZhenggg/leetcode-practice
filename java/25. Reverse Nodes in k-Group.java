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
    public ListNode reverseKGroup(ListNode head, int k) {
        if (k == 1) return head;

        ListNode dummy = new ListNode(-1);
        ArrayList<ListNode> restPrevs = new ArrayList<>();
        ListNode prev = dummy;
        ListNode tempCurr = head;
        ListNode currPrev = null;
        ListNode tempNextPrev = null;

        dummy.next = head;

        while (head != null) {
            tempCurr = head;

            for(int count = 1; count < k; count++) {
                if(tempCurr.next == null) return dummy.next;
                restPrevs.add(tempCurr);
                tempCurr = tempCurr.next;
            }

            prev.next = tempCurr;

            for (int j = 0; j < restPrevs.size(); j++) {
                currPrev = restPrevs.get(j);
                if(j == 0) {
                    currPrev.next = prev.next.next;
                    tempNextPrev = currPrev;
                } else {
                    ListNode prevPrev = restPrevs.get(j - 1);
                    currPrev.next = prevPrev;
                }
            }

            prev.next.next = currPrev;

            restPrevs.clear();

            prev = tempNextPrev;
            head = head.next;
        }

        return dummy.next;
    }
}

// better iterative solution
class Solution {
    public ListNode reverseKGroup(ListNode head, int k) {
        ListNode dummy = new ListNode(-1);
        ListNode begin = dummy;
        dummy.next = head;
        for(int i = 1; head != null; i++) {
            if(i % k == 0) {
                begin = reverse(begin, head.next);
                head = begin.next;
            } else {
                head = head.next;
            }
        }

        return dummy.next;
    }

    public ListNode reverse(ListNode begin, ListNode end) {
        ListNode curr = begin.next;
        ListNode first = curr;
        ListNode prev = begin;

        while (curr != end) {
            ListNode next = curr.next;
            curr.next = prev;
            prev = curr;
            curr = next;
        }

        begin.next = prev;
        first.next = end;

        return first;
    }
}

// recursive solution 1
class Solution {
    public ListNode reverseKGroup(ListNode head, int k) {
        ListNode curr = head;
        int i = 0;
        while(i != k && curr != null) {
            curr = curr.next;
            i++;
        }

        if(i != k) return head;

        curr = reverseKGroup(curr, k);
        while (i-- > 0) {
            ListNode next = head.next;
            head.next = curr;
            curr = head;
            head = next;
        }
        head = curr;

        return head;
    }
}

// recursive solution 2
class Solution {
    public ListNode reverseKGroup(ListNode head, int k) {
        ListNode pt = head;
        for(int i = 0; i < k; i++) {
            if(pt == null) return head;
            pt = pt.next;
        }

        ListNode prev = null;
        ListNode curr = head;
        for(int i = 0; i < k; i++) {
            ListNode next = curr.next;
            curr.next = prev;
            prev = curr;
            curr = next;
        }
        head.next = reverseKGroup(curr, k);
        return prev;
    }
}