// union find solution
class UnionFind {
    private int[] sets;

    public int[] getSets() {
        return sets;
    }

    public UnionFind(int[] sets) {
        this.sets = sets;
    }

    public void union(int x, int y) {
        int rootX = this.findRoot(x);
        int rootY = this.findRoot(y);
        if (this.sets[rootX] <= this.sets[rootY]) {
            this.sets[rootX] += this.sets[rootY];
            this.sets[rootY] = rootX;
        } else {
            this.sets[rootY] += this.sets[rootX];
            this.sets[rootX] = rootY;
        }
    }

    public int findRoot(int node) {
        int root;
        for (root = node; this.sets[root] >= 0; root = this.sets[root]) ;

        while (root != node) {
            int parent = this.sets[node];
            this.sets[node] = root;
            node = parent;
        }
        return root;
    }

    public boolean isSameRoot(int x, int y) {
        if (this.sets[x] < 0 || this.sets[y] < 0) return false;
        return this.findRoot(x) == this.findRoot(y);
    }
}

class Solution {
    public int longestConsecutive(int[] nums) {
        if(nums.length <= 1) return nums.length;

        int[] sets = new int[nums.length];
        Arrays.fill(sets, -1);

        UnionFind unionFind = new UnionFind(sets);
        doUnion(nums, unionFind);

        int max = 1;
        int[] resultSets = unionFind.getSets();
        for (int item : resultSets) {
            if (item > 0) continue;

            int absCount = Math.abs(item);
            if (absCount >= max) {
                max = absCount;
            }
        }
        return max;
    }

    public void doUnion(int[] nums, UnionFind unionFind) {
        Map<Integer, Integer> map = new HashMap<>();

        for (int i = 0; i < nums.length; i++) {
            if (map.containsKey(nums[i])) continue;

            if (map.containsKey(nums[i] - 1)) {
                int x = map.get(nums[i] - 1);
                int y = i;
                if (!unionFind.isSameRoot(x, y)) unionFind.union(x, y);
            }

            if (map.containsKey(nums[i] + 1)) {
                int x = map.get(nums[i] + 1);
                int y = i;
                if (!unionFind.isSameRoot(x, y)) unionFind.union(x, y);
            }

            map.put(nums[i], i);
        }
    }
}

// set solution
class Solution {
    public int longestConsecutive(int[] nums) {
        if (nums.length == 0) return 0;
        HashSet<Integer> set = new HashSet<>();
        int ans = 1;
        for (int num : nums) set.add(num);
        for (int num : nums) {
            if (!set.contains(num + 1)) {
                int count = 1;
                while (set.contains(num - 1)) {
                    num--;
                    count++;
                }
                ans = Math.max(count, ans);
            }
        }
        return ans;
    }
}