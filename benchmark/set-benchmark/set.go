package setbenchmark 

type Set interface {
	Add(value int)
	Contains(value int) bool 
	Length() int 
	RemoveDuplicates()
}

