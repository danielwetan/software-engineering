package setbenchmark

type OnDemandArraySet struct {
	data []int
}

func (s *OnDemandArraySet) Add(value int) {
	s.data = append(s.data, value)
}

func (s *OnDemandArraySet) Contains(value int) (exists bool) {
	for _, v := range s.data {
		if v == value {
			return true
		}
	}
	return false
}

func (s *OnDemandArraySet) Length() int {
	return len(s.data)
}
func (s *OnDemandArraySet) RemoveDuplicates() {
	length := len(s.data) - 1
	for i := 0; i < length; i++ {
		for j := i + 1; j <= length; j++ {
			if s.data[i] == s.data[j] {
				s.data[j] = s.data[length]
				s.data = s.data[0:length]
				length--
				j--
			}
		}
	}
}

func NewOnDemandArraySet() Set {
	return &OnDemandArraySet{make([]int, 0, 100)}
}
