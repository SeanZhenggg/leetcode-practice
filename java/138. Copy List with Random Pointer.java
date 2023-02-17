/*
// Definition for a Node.
class Node {
    int val;
    Node next;
    Node random;

    public Node(int val) {
        this.val = val;
        this.next = null;
        this.random = null;
    }
}
*/

class Solution {
    public Node copyRandomList(Node head) {
        Node curr = new Node(-1);
        Node dummy = curr;
        HashMap<Node, ArrayList<Integer>> reverseMap = new HashMap<>();
        HashMap<Integer, Integer> map = new HashMap<>();
        HashMap<Integer, Node> newMap = new HashMap<>();
        int index = 0;
        Node currHead = head;
        while (currHead != null) {
            curr.next = new Node(currHead.val);
            ArrayList<Integer> valueArr = reverseMap.get(currHead.random);
            if(valueArr == null) {
                valueArr = new ArrayList<>();
            }
            valueArr.add(index);
            reverseMap.put(currHead.random, valueArr);
            newMap.put(index, curr.next);
            currHead = currHead.next;
            curr = curr.next;
            index++;
        }

        currHead = head;
        index = 0;
        while (currHead != null) {
            ArrayList<Integer> mapIndexs = reverseMap.get(currHead);
            if(mapIndexs != null) {
                for (Integer j : mapIndexs) {
                    if(map.get(j) == null) {
                        map.put(j, index);
                    }
                }
            }
            currHead = currHead.next;
            index++;
        }
        ArrayList<Integer> mapIndexs = reverseMap.get(null);
        if(mapIndexs != null) {
            for (Integer j : mapIndexs) {
                map.put(j, -1);
            }
        }

        index = 0;
        curr = dummy.next;
        while (curr != null) {
            int mapNodeIndex = map.get(index);
            Node mapNode = mapNodeIndex == -1 ? null : newMap.get(mapNodeIndex);
            curr.random = mapNode;
            curr = curr.next;
            index ++;
        }

        return dummy.next;
    }
}