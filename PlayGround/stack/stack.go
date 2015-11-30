package stack

type Stack []interface{}

func (k *Stack) Push(s interface{}) {
	*k = append(*k, s)
}

func (k *Stack) Pop() (s interface{}, ok bool) {
	if k.Empty() {
		return
	}

	last := len(*k) - 1
	s = (*k)[last]
	*k = (*k)[:last]

	return s, true
}

func (k *Stack) Peek() (s interface{}, ok bool) {
	if k.Empty() {
		return
	}

	last := len(*k) - 1
	s = (*k)[last]

	return s, true
}

func (k *Stack) Empty() bool {
	return len(*k) == 0
}
