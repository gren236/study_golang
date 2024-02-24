package hasht

func HashDJB2(s []byte) uint64 {
	var hash uint64 = 5381

	for _, c := range s {
		hash = ((hash << 5) + hash) + uint64(c)
	}

	return hash
}
