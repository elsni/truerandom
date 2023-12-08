package truerandom

import "errors"

type FloatGenerator struct {
	numberpool []int
	num        int
	poolptr    int
	Min        float64
	Max        float64
	Buffersize int
}

func NewFloatGenerator(min float64, max float64, num int) (FloatGenerator, error) {
	result := FloatGenerator{}
	result.Min = min
	result.Max = max
	result.num = num
	if max <= min {
		return result, errors.New("max must be greater than min")
	}
	err := result.Refresh()
	return result, err
}

func (i *FloatGenerator) Refresh() error {
	list, err := fetch(0, 1e9, i.num)
	i.numberpool = list
	i.Buffersize = len(i.numberpool)
	i.poolptr = 0
	return err
}

func (i *FloatGenerator) Rand() float64 {
	if i.Buffersize == 0 {
		return -1
	}
	i.poolptr += 1
	if i.poolptr == i.Buffersize {
		i.Refresh()
	}
	val := float64(i.numberpool[i.poolptr])
	distance := i.Max - i.Min

	return val/float64(1e9)*distance + i.Min
}
