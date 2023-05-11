// using stack
class Solution {
    public static int carFleet(int target, int[] position, int[] speed) {
        int [][] cars = new int[position.length][2];
        Stack<Double> stack = new Stack<>();
        for(int i = 0; i < position.length; i++) {
            cars[i][0] = position[i];
            cars[i][1] = speed[i];
        }

        Arrays.sort(cars, (a, b) -> b[0] - a[0]);

        for (int[] car: cars) {
            double desTime = (double) (target - car[0]) / car[1];
            stack.push(desTime);
            if(stack.size() >= 2 && stack.get(stack.size() - 1) <= stack.get(stack.size() - 2)) {
                stack.pop();
            }
        }

        return stack.size();
    }
}

// using vars
class Solution {
    public static int carFleet(int target, int[] position, int[] speed) {
        int [][] cars = new int[position.length][2];
        double curTime = 0.0;
        int fleets = 0;
        for(int i = 0; i < position.length; i++) {
            cars[i][0] = position[i];
            cars[i][1] = speed[i];
        }

        Arrays.sort(cars, (a, b) -> b[0] - a[0]);

        for (int[] car: cars) {
            double desTime = (double) (target - car[0]) / car[1];
            
            if(curTime < desTime) {
                fleets += 1;
                curTime = desTime;
            }
        }

        return fleets;
    }
}

// better performance
class Solution {
    public static int carFleet(int target, int[] position, int[] speed) {
        double [] desTimes = new double[target + 1];
        double curTime = 0.0;
        int fleets = 0;

        for(int i = 0; i < position.length; i++) {
            desTimes[position[i]] = (double) (target - position[i]) / speed[i];
        }

        for(int i = target; i >= 0; i--) {
            if(curTime < desTimes[i]) {
                fleets += 1;
                curTime = desTimes[i];
            }
        }

        return fleets;
    }
}