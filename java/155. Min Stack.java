// first solution using db-linked-list

class ListNode {
    public long val;
    public ListNode next = null;
    public ListNode prev = null;
    public ListNode(long val) {
        this.val = val;
    }
}

class MinStack {
    private ArrayList<ListNode> stack;
    private ListNode head;
    private ListNode tail;

    public MinStack() {
        this.head = new ListNode(Long.MIN_VALUE);
        this.tail = new ListNode(Long.MAX_VALUE);
        this.stack = new ArrayList<>();
    }
    
    public void push(int val) {
        ListNode newNode = new ListNode(val);
        this.stack.add(newNode);
        if (this.head.next == null && this.tail.prev == null) {
            this.head.next = newNode;
            newNode.next = this.tail;
            this.tail.prev = newNode;
            newNode.prev = this.head;
        } else {
            ListNode headNextNode = this.head.next;
            if(headNextNode.val > val) {
                this.head.next = newNode;
                newNode.next = headNextNode;
                headNextNode.prev = newNode;
                newNode.prev = this.head;
            } else {
                ListNode headNextNextNode = headNextNode.next;
                headNextNode.next = newNode;
                newNode.next = headNextNextNode;
                headNextNextNode.prev = newNode;
                newNode.prev = headNextNode;
            }
        }
    }
    
    public void pop() {
        ListNode topNode = this.stack.get(this.stack.size() - 1);
        topNode.prev.next = topNode.next;
        topNode.next.prev = topNode.prev;
        this.stack.remove(this.stack.size() - 1);
    }
    
    public int top() {
        return (int) this.stack.get(this.stack.size() - 1).val;
    }
    
    public int getMin() {
        return (int) this.head.next.val;
    }
}