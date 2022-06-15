/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
var addTwoNumbers = function(l1, l2) {
    let answer = new ListNode(); // 答案
    let tempLinkedList = answer; // 用來暫存每一個答案的linked list
    let jingWei = 0; // 進位紀錄
    // 條件為: 只要l1, l2有一個還有連結下個linked list或是有進位
    while(l1 !== null || l2 !== null || jingWei != 0) {
        const l1Val = l1 ? l1.val : 0;
        const l2Val = l2 ? l2.val : 0;
        let sum = (l1Val + l2Val + jingWei);
        let residual = 0; // 加總扣除進位後的餘數, 代表每一個linked list的value
        if(sum >= 10) {
            jingWei = 1;
            residual = sum - 10;
        } 
        else {
            jingWei = 0;
            residual = sum;
        }
        
        tempLinkedList.val = residual;
        // 指向l1, l2的下一個 linked list
        l1 = l1 ? l1.next : null; 
        l2 = l2 ? l2.next : null;
        // 避免創建下一個空的linked list的方法，目前沒想到更好的做法
        if(l1 === null && l2 === null && jingWei === 0) break;
        // 創建並指向下一個答案的linked list
        tempLinkedList.next = new ListNode();
        tempLinkedList = tempLinkedList.next;
    }
    return answer;
};