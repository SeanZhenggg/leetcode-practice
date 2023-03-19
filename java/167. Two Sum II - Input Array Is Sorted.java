// two pointer
class Solution {
    public int[] twoSum(int[] numbers, int target) {
        int left = 0, right = numbers.length - 1;

        while (target != numbers[left] + numbers[right]) {
            if(target > numbers[left] + numbers[right]) left ++;
            else right --;
        }

        return new int[] {left + 1, right + 1};
    }
}

// binary search
class Solution {
    public int[] twoSum(int[] numbers, int target) {
        for(int i = 0; i < numbers.length; i++) {
            int left = i + 1, right = numbers.length - 1;

            while(left <= right) {
                int mid = (left + right) / 2;

                if(numbers[i] + numbers[mid] == target) {
                    return new int[] {i + 1, mid + 1};
                }
                else if (numbers[i] + numbers[mid] > target) {
                    right = mid - 1;
                }
                else {
                    left = mid + 1;
                }
            }
        }

        return result;
    }
}

// faster two pointer
class Solution {
    public int[] twoSum(int[] numbers, int target) {
        int i = 0, j = numbers.length - 1;
        while(i <= j){
            int mid = (i + j)/2;
            if(numbers[i] + numbers[j] < target){
                if(numbers[mid] + numbers[j] < target){
                    i = mid + 1;
                } else {
                    i++;
                }
            } else if(numbers[i] + numbers[j] > target){
                if(numbers[i] + numbers[mid] > target){
                    j = mid - 1;
                } else {
                    j--;
                }
            } else {
                return new int[] {i + 1, j + 1};
            }
        }
        return null;
    }
}