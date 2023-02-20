// hashtable with double linked list solution
class ListNode {
    int key;
    int val;
    ListNode prev;
    ListNode next;
    public ListNode (int key, int val, ListNode prev, ListNode next) {
        this.key = key;
        this.val = val;
        this.prev = prev;
        this.next = next;
    }
}

class LRUCache {
    private HashMap<Integer, ListNode> map = new HashMap<>();
    private int capacity;
    private ListNode head;
    private ListNode tail;

    public LRUCache(int capacity) {
        this.capacity = capacity;

        this.head = new ListNode(-1, -1, null, null);
        this.tail = new ListNode(-1, -1, null, null);

        this.head.next = this.tail; // head -> tail
        this.tail.prev = this.head; // head <- tail
    }
    
    public int get(int key) {
        ListNode node = this.map.get(key); // get node by key
        if (node == null) return -1; // not found

        // update node position in double linked list
        this.remove(node);
        this.add(node);
        return node.val;
    }
    
    public void put(int key, int value) {
        ListNode node;
        if (this.map.containsKey(key)) {
            // node has already existed in linked list
            node = this.map.get(key);
            this.remove(node);
            node.val = value;
        } else {
            // create new node
            node = new ListNode(key, value, null, null);
        }

        // update node position in double linked list
        this.add(node);
        this.map.put(key, node);

        // if map size exceed capacity, remove the least recently used node (right of the head)
        if(this.map.size() > this.capacity) {
            ListNode lruNode = this.head.next;
            this.remove(lruNode);
            this.map.remove(lruNode.key);
        }
    }

    public void add (ListNode node) {
        ListNode tailPrev = this.tail.prev;
        tailPrev.next = node; // tailPrev -> node
        node.next = this.tail; // node -> tail
        this.tail.prev = node; // node <- tail
        node.prev = tailPrev; // tailPrev <- node
    }

    public void remove (ListNode node) {
        ListNode nodePrev = node.prev;
        ListNode nodeNext = node.next;
        nodePrev.next = nodeNext;
        nodeNext.prev = nodePrev;
    }
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * LRUCache obj = new LRUCache(capacity);
 * int param_1 = obj.get(key);
 * obj.put(key,value);
 */