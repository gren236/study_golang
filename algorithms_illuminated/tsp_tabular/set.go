package main

type binSet struct {
	length uint // Max number of different elements
}

func (s binSet) encode(elements []int) uint {
	var res uint

	for _, e := range elements {
		res |= 1 << e
	}

	return res
}

func (s binSet) decode(b uint) []int {
	var res []int

	for i := uint(0); i < s.length; i++ {
		if (b>>i)&1 == 1 {
			res = append(res, int(i))
		}
	}

	return res
}

func (s binSet) remove(b uint, v int) uint {
	mask := ^(1 << v)
	b &= uint(mask)

	return b
}
