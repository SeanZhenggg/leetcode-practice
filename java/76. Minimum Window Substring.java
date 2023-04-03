class Solution {

  public static String minWindow(String s, String t) {
    if (s.length() < t.length()) return "";

    if (s.length() == 0 || t.length() == 0) return "";

    Map<Character, Integer> tMap = new HashMap<>();
    Map<Character, Integer> sMap = new HashMap<>();
    int[] res = new int[] { -1, -1 };
    int resLength = Integer.MAX_VALUE;

    for (int i = 0; i < t.length(); i++) {
      tMap.put(t.charAt(i), tMap.getOrDefault(t.charAt(i), 0) + 1);
    }

    int need = tMap.size(), have = 0;
    int l = 0;
    for (int r = 0; r < s.length(); r++) {
      sMap.put(s.charAt(r), sMap.getOrDefault(s.charAt(r), 0) + 1);
      if (
        tMap.containsKey(s.charAt(r)) &&
        sMap.get(s.charAt(r)).equals(tMap.get(s.charAt(r)))
      ) {
        have += 1;
      }
      while (have == need) {
        if ((r - l + 1) < resLength) {
          resLength = (r - l + 1);
          res[0] = l;
          res[1] = r;
        }
        sMap.put(s.charAt(l), sMap.get(s.charAt(l)) - 1);
        if (
          tMap.containsKey(s.charAt(l)) &&
          (int) sMap.get(s.charAt(l)) < (int) tMap.get(s.charAt(l))
        ) {
          have -= 1;
        }
        l++;
      }
    }

    if (res[0] == -1 && res[1] == -1) return "";
    return s.substring(res[0], res[1] + 1);
  }
}

class Solution {

  public static String minWindow(String s, String t) {
    if (s.length() < t.length()) return "";

    if (s.length() == 0 || t.length() == 0) return "";

    Map<Character, Integer> tMap = new HashMap<>();

    int[] res = new int[] { -1, -1 };
    int resLength = Integer.MAX_VALUE;

    for (int i = 0; i < t.length(); i++) {
      tMap.put(t.charAt(i), tMap.getOrDefault(t.charAt(i), 0) + 1);
    }

    int need = tMap.size(), have = 0;
    int l = 0;
    for (int r = 0; r < s.length(); r++) {
      if (tMap.containsKey(s.charAt(r))) {
        tMap.put(s.charAt(r), tMap.getOrDefault(s.charAt(r), 0) - 1);
        if (tMap.get(s.charAt(r)) == 0) have += 1;
      }

      while (have == need) {
        if ((r - l + 1) < resLength) {
          resLength = (r - l + 1);
          res[0] = l;
          res[1] = r;
        }
        if (tMap.containsKey(s.charAt(l))) {
          if (tMap.get(s.charAt(l)) == 0) have -= 1;
          tMap.put(s.charAt(l), tMap.get(s.charAt(l)) + 1);
        }
        l++;
      }
    }

    if (res[0] == -1 && res[1] == -1) return "";
    return s.substring(res[0], res[1] + 1);
  }
}
