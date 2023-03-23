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

// priority queue solution
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

// bucket sort solution
class Solution {
    public static int[] topKFrequent(int[] nums, int k) {
        if (k == 1 && nums.length == 1) return nums;

        Map<Integer, Integer> map = new HashMap<>();
        List<Integer> buckets[] = new ArrayList[nums.length + 1];

        for (int n : nums) {
            map.merge(n, 1, Integer::sum);
        }

        for (Map.Entry<Integer, Integer> entry : map.entrySet()) {
            int freq = entry.getValue();
            if(buckets[freq] == null) {
                buckets[freq] = new ArrayList<>();
            }
            buckets[freq].add(entry.getKey());
        }

        int j = 0;
        int[] result = new int[k];
        for(int i = nums.length; i >= 0; i--) {
            if(buckets[i] != null) {
                for(int n : buckets[i]) {
                    result[j++] = n;
                    if(j == k) break;
                }
                if(j == k) break;
            }
        }
        return result;
    }
}