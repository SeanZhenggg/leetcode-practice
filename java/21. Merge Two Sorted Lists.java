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
    public ListNode mergeTwoLists(ListNode list1, ListNode list2) {
        ListNode result = null;
        ListNode answer = null;
        while (list1 != null || list2 != null) {
            ListNode nextNode = null;
            if (list1 == null) {
                nextNode = new ListNode(list2.val);
                list2 = list2.next;
            } else if (list2 == null) {
                nextNode = new ListNode(list1.val);
                list1 = list1.next;
            } else {
                if (list1.val > list2.val) {
                    nextNode = new ListNode(list2.val);
                    list2 = list2.next;
                } else {
                    nextNode = new ListNode(list1.val);
                    list1 = list1.next;
                }
            }

            if(result == null) {
                result = nextNode;
                answer = result;
            }
            else {
                result.next = nextNode;
                result = result.next;
            }
        }

        return answer;
    }
}