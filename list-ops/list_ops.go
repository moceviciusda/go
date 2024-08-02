package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	if s.Length() < 1 {
		return initial
	}

	initial = fn(initial, s[0])
	s = s[1:]

	return s.Foldl(fn, initial)
}

// weird exercise/test case. Why would you make 'initial' the second argument?
// Think correct solution should be run the predicate on every element but in
// a reversed list and keep the initial argument as an actual initial value
func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	for i := s.Length() - 1; i >= 0; i-- {
		initial = fn(s[i], initial)
	}
	return initial
}

func (s IntList) Filter(fn func(int) bool) IntList {
	newList := make(IntList, 0)
	for _, num := range s {
		if fn(num) {
			newList = append(newList, num)
		}
	}
	return newList
}

func (s IntList) Length() (sLen int) {
	for range s {
		sLen++
	}
	return
}

func (s IntList) Map(fn func(int) int) IntList {
	for i, num := range s {
		s[i] = fn(num)
	}

	return s
}

func (s IntList) Append(lst IntList) IntList {
	newList := make(IntList, s.Length()+lst.Length())
	for i, num := range s {
		newList[i] = num
	}
	for i, num := range lst {
		newList[s.Length()+i] = num
	}

	return newList
}

func (s IntList) Reverse() IntList {
	newList := make(IntList, 0)
	for i := s.Length() - 1; i >= 0; i-- {
		newList = append(newList, s[i])
	}
	return newList
}

func (s IntList) Concat(lists []IntList) IntList {
	for _, list := range lists {
		s = s.Append(list)
	}

	return s
}
