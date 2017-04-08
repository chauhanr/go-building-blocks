package interfaces

type Element interface{}

type Miner interface {
	Less(i int, j int) bool
	Len() int
	Get(index int) Element
}

func Min(data Miner) Element {
	if data.Len() != 0 {
		min := 0
		for i := 1; i < data.Len(); i++ {
			if data.Less(i, min) {
				min = i
				//log.Printf("index choosen %v \n", min)
			}
		}
		return data.Get(min)
	}
	return nil
}

type IntElement []int

func (p IntElement) Len() int {
	return len(p)
}

func (p IntElement) Less(i int, j int) bool {
	return p[i] < p[j]
}

func (p IntElement) Get(index int) Element {
	length := len(p)
	if index > length {
		panic("Try to access an element not there")
	}
	return p[index]
}

type FloatElement []float32

func (p FloatElement) Len() int {
	return len(p)
}

func (p FloatElement) Less(i int, j int) bool {
	return p[i] < p[j]
}

func (p FloatElement) Get(index int) Element {
	length := len(p)
	if index > length {
		panic("Try to access an element not there")
	}
	return p[index]
}
