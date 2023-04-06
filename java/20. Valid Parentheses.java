class Solution {

  public boolean isValid(String s) {
    Stack<Character> stack = new Stack<Character>();
    for (int i = 0; i < s.length(); i++) {
      char c = s.charAt(i);
      if (c == '[' || c == '(' || c == '{') {
        stack.push(c);
        continue;
      } else {
        if (stack.empty()) return false;
      }

      if (c == ']') {
        char last = stack.pop();
        if (last != '[') return false;
      } else if (c == ')') {
        char last = stack.pop();
        if (last != '(') return false;
      } else if (c == '}') {
        char last = stack.pop();
        if (last != '{') return false;
      }
    }
    return stack.empty();
  }
}

class Solution {

  public boolean isValid(String s) {
    if (s.length() <= 1) return false;
    ArrayList<Character> stack = new ArrayList<>();

    for (int i = 0; i < s.length(); i++) {
      switch (s.charAt(i)) {
        case ')':
          if (stack.isEmpty()) return false;
          if (stack.get(stack.size() - 1) != '(') return false;
          stack.remove(stack.size() - 1);
          break;
        case '}':
          if (stack.isEmpty()) return false;
          if (stack.get(stack.size() - 1) != '{') return false;
          stack.remove(stack.size() - 1);
          break;
        case ']':
          if (stack.isEmpty()) return false;
          if (stack.get(stack.size() - 1) != '[') return false;
          stack.remove(stack.size() - 1);
          break;
        case '(':
        case '{':
        case '[':
          stack.add(s.charAt(i));
          break;
      }
    }
    return stack.isEmpty();
  }
}
