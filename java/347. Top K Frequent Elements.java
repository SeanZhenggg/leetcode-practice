class Node implements Comparable<Node> {
    public int getKey() {
        return key;
    }

    public void setKey(int key) {
        this.key = key;
    }

    public int getTimes() {
        return times;
    }

    public void setTimes(int times) {
        this.times = times;
    }

    private int key;
    private int times;

    public Node(int key, int times) {
        this.key = key;
        this.times = times;
    }

    @Override
    public int compareTo(Node o) {
        return o.times > this.times ? 1 : -1;
    }
}

class Solution {
    public int[] topKFrequent(int[] nums, int k) {
        if(k == 1 && nums.length == 1) return nums;

        PriorityQueue<Node> pq = new PriorityQueue<>();
        Map<Integer, Integer> map = new HashMap<>();
        int[] result = new int[k];
        for(int num : nums) {
            Object val = map.get(num);
            map.put(num, val == null ? 1 : (int) val + 1);
        }

        map.forEach((key, v) -> {
            pq.add(new Node(key, v));
        });
        int i = 0;
        while (k-- > 0) {
            Node polled = pq.poll();
            result[i] = polled.getKey();
            i++;
        }
        
        return result;
    }
}