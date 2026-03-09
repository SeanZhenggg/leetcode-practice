package stack

func isValid(s string) bool {
	st := make([]string, 0, len(s))
	parenthesesMap := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}

	for _, ch := range s {
		sc := string(ch)
		if len(st) > 0 && parenthesesMap[sc] == st[len(st)-1] {
			st = st[:len(st)-1]
		} else {
			st = append(st, sc)
		}
	}

	return len(st) == 0
}

func isValid2(s string) bool {
	st := make([]rune, 0, len(s))
	parenthesesMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, ch := range s {
		if open, found := parenthesesMap[ch]; found {
			if len(st) > 0 && open == st[len(st)-1] {
				st = st[:len(st)-1]
			} else {
				return false
			}
		} else {
			st = append(st, ch)
		}
	}

	return len(st) == 0
}

// 不能用 two pointer 解，因為可以有這種 use case: "()[]{}"
func isValidTwoPointerFailed(s string) bool {
	l, r := 0, len(s)-1

	m := map[uint8]uint8{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for l < r {
		if matched, ok := m[s[l]]; !ok || matched != s[r] {
			return false
		}

		l++
		r--
	}

	return true
}

type Stack[T any] struct {
	st []T
}

func (s *Stack[T]) Top() T {
	var rt T
	if len(s.st) == 0 {
		return rt
	}
	return s.st[len(s.st)-1]
}

func (s *Stack[T]) Push(value T) {
	s.st = append(s.st, value)
}

func (s *Stack[T]) Pop() T {
	var rt T
	if len(s.st) == 0 {
		return rt
	}
	rt = s.st[len(s.st)-1]
	s.st = s.st[:len(s.st)-1]
	return rt
}

func (s *Stack[T]) Len() int {
	return len(s.st)
}

func isValidStack(s string) bool {
	st := new(Stack[uint8])
	m := map[uint8]uint8{
		')': '(',
		']': '[',
		'}': '{',
	}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			st.Push(s[i])
		} else {
			t := st.Pop()
			if m[s[i]] != t {
				return false
			}
		}
	}

	return st.Len() == 0
}
