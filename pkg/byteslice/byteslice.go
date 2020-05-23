package byteslice

type ByteSlice []byte

func (p *ByteSlice) Append(data []byte) {
	slice := *p

	ln := len(slice)
	if ln + len(data) > cap(slice) { // reallocate
		newSlice := make([]byte, (ln + len(data)) * 2)
		copy(newSlice, slice)
		slice = newSlice
	}
	copy(slice[ln:], data)

	*p = slice
}

func (p *ByteSlice) Write(data []byte) (n int, err error) {
	slice := *p
	slice.Append(data)
	*p = slice

	return len(data), nil
}