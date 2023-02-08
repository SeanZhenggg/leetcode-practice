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