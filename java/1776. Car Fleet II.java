// stack solution
class Solution {
    public static double[] getCollisionTimes(int[][] cars) {
        double[] ans = new double[cars.length];

        for (int i = 0; i < ans.length; i++) ans[i] = -1.0;

        Stack<Integer> stack = new Stack<>();

        for(int i = cars.length - 1; i >= 0; i--) {

            while (!stack.empty() && (cars[i][1] <= cars[stack.peek()][1])) stack.pop();

            while (!stack.empty()) {
                double collisionTime = (double)(cars[stack.peek()][0] - cars[i][0]) / (cars[i][1] - cars[stack.peek()][1]);
                if(collisionTime >= ans[stack.peek()] && ans[stack.peek()] != -1.0) {
                    stack.pop();
                } else {
                    ans[i] = collisionTime;
                    break;
                }
            }

            stack.push(i);
        }

        return ans;
    }
}

// using priority queue solution
public static double[] getCollisionTimes(int[][] cars) {
    double[] ans = new double[cars.length];

    for (int i = 0; i < ans.length; i++) ans[i] = -1.0;

    PriorityQueue<Node> pq = new PriorityQueue<>((a, b) -> Double.compare(a.time, b.time));

    TreeSet<Integer> ts = new TreeSet<>();
    for (int i = 0; i < cars.length; i++) ts.add(i);

    for(int i = 0; i < cars.length - 1; i++) {
        if(cars[i][1] > cars[i + 1][1]) {
            pq.offer(new Node(i, (double)(cars[i + 1][0] - cars[i][0]) / (cars[i][1] - cars[i+1][1])));
        }
    }

    while(!pq.isEmpty()) {
        Node top = pq.poll();

        if(ans[top.node] != -1.0) continue;

        ans[top.node] = top.time;

        Object prev = ts.lower(top.node);

        if(prev == null) continue;

        int next = ts.higher(top.node);

        if(cars[(int)prev][1] > cars[next][1]) {
            pq.offer(new Node((int) prev, (double)(cars[next][0] - cars[(int)prev][0]) / (cars[(int)prev][1] - cars[next][1])));
        }

        ts.remove(top.node);
    }

    return ans;
}