class Solution {
    public List<List<String>> groupAnagrams(String[] strs) {
        List<List<String>> result = new ArrayList<>();
        if(strs.length == 1) {
            List<String> newCell = new ArrayList<>();
            newCell.add(strs[0]);
            result.add(newCell);
            return result;
        }

        Map<String, Integer> map = new HashMap<>();

        for(int i = 0; i < strs.length; i++) {
            char[] arr = strs[i].toCharArray();
            Arrays.sort(arr);
            String sorted = new String(arr);
            Object index = map.get(sorted);
            if(index != null) {
                result.get((int) index).add(strs[i]);
            } else {
                List<String> newCell = new ArrayList<>();
                newCell.add(strs[i]);
                result.add(newCell);
                map.put(sorted, result.size() - 1);
            }
        }
        return result;
    }
}