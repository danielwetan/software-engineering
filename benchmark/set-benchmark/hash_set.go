package setbenchmark

type HashSet struct {
	data map[int]bool
}

func (h *HashSet) Add(value int) {
	h.data[value] = true
}

func (h *HashSet) Contains(value int) bool {
	_, exists := h.data[value]
	return exists
}

func (h *HashSet) Length() int {
	return len(h.data)
}

func (h *HashSet) RemoveDuplicates() {}

func NewSet() Set {
	return &HashSet{make(map[int]bool)}
}
