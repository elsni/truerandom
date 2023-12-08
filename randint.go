package truerandom

type IntGenerator struct {
	numberpool []int
	num        int
	poolptr    int
	Min        int
	Max        int
	Buffersize int
}

func NewIntGenerator(min, max, num int) (IntGenerator, error) {
	result := IntGenerator{}
	result.Min = min
	result.Max = max
	result.num = num
	err := result.Refresh()
	return result, err
}

func (i *IntGenerator) Refresh() error {
	list, err := fetch(i.Min, i.Max, i.num)
	i.numberpool = list
	i.Buffersize = len(i.numberpool)
	i.poolptr = 0
	return err
}

func (i *IntGenerator) Rand() int {
	if i.Buffersize == 0 {
		return -1
	}
	i.poolptr += 1
	if i.poolptr == i.Buffersize {
		i.Refresh()
	}
	return i.numberpool[i.poolptr]
}
