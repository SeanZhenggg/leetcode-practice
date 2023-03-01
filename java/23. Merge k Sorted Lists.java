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
    private ListNode dummy = new ListNode(-1);
    public ListNode mergeKLists(ListNode[] lists) {
         for (ListNode list : lists) {
             this.merge(this.dummy.next, list);
         }
         return this.dummy.next;
    }

    public void merge(ListNode l1, ListNode l2) {
        ListNode tempHead = this.dummy;
        while (l1 != null && l2 != null) {
            if (l1.val < l2.val) {
                tempHead.next = l1;
                l1 = l1.next;
            } else {
                tempHead.next = l2;
                l2 = l2.next;
            }
            tempHead = tempHead.next;
        }

        if (l1 != null) tempHead.next = l1;
        if (l2 != null) tempHead.next = l2;
    }
}

// merge sort way
class Solution {
    public ListNode mergeKLists(ListNode[] lists) {
        if (lists.length == 0) return null;
        return mergeSort(lists, 0, lists.length - 1);
    }

    public ListNode mergeSort(ListNode[] lists, int left, int right) {
        if (left == right) return lists[left];
        if (left > right) return null;

        int mid = (left + right) / 2;
        ListNode l = mergeSort(lists, left, mid);
        ListNode r = mergeSort(lists, mid + 1, right);
        return merge(l, r);
    }

    public ListNode merge(ListNode l1, ListNode l2) {
        ListNode dummy = new ListNode(-1);
        ListNode tempHead = dummy;
        while (l1 != null && l2 != null) {
            if (l1.val < l2.val) {
                tempHead.next = l1;
                l1 = l1.next;
            } else {
                tempHead.next = l2;
                l2 = l2.next;
            }
            tempHead = tempHead.next;
        }

        if (l1 != null) tempHead.next = l1;

        if (l2 != null) tempHead.next = l2;

        return dummy.next;
    }
}
