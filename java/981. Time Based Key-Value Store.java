class Data {
    String val;
    int timestamp;

    public Data(String v, int t) {
        this.val = v;
        this.timestamp = t;
    }
}

class TimeMap {
    private Map<String, ArrayList<Data>> map;

    public TimeMap() {
        this.map = new HashMap<>();
    }

    public void set(String key, String value, int timestamp) {
        if (!this.map.containsKey(key)) {
            ArrayList<Data> temp = new ArrayList<>();
            this.map.put(key, temp);
        }

        ArrayList<Data> list = this.map.get(key);
        Data d = new Data(value, timestamp);
        list.add(d);

        this.map.put(key, list);
    }

    public String get(String key, int timestamp) {
        if (!this.map.containsKey(key)) return "";

        ArrayList<Data> list = this.map.get(key);
        int left = 0, right = list.size() - 1;

        while (left <= right) {
            int mid = (left + right) / 2;

            if (timestamp < list.get(mid).timestamp) {
                right = mid - 1;
            } else if (timestamp > list.get(mid).timestamp) {
                if(mid + 1 >= list.size()) return list.get(mid).val;
                if (list.get(mid + 1).timestamp > timestamp) return list.get(mid).val;
                left = mid + 1;
            } else {
                return list.get(mid).val;
            }
        }

        return "";
    }
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * TimeMap obj = new TimeMap();
 * obj.set(key,value,timestamp);
 * String param_2 = obj.get(key,timestamp);
 */