package setbenchmark

type RealtimeArraySet struct {
	data []int
}

func (s *RealtimeArraySet) Add(value int) {
	if !s.Contains(value) {
		s.data = append(s.data, value)
	}
}

func (s *RealtimeArraySet) Contains(value int) (exists bool) {
	for _, v := range s.data {
		if v == value {
			return true
		}
	}
	return false
}

func (s *RealtimeArraySet) Length() int {
	return len(s.data)
}
func (s *RealtimeArraySet) RemoveDuplicates() {}

func NewRealtimeArraySet() Set {
	return &RealtimeArraySet{make([]int, 0, 100)}
}
