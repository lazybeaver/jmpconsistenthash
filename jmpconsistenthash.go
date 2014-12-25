package jmpconsistenthash

// Hash computes the Jump Consistent Hash
func Hash(key uint64, shards uint64) uint64 {
	s := uint64(0)
	for i := uint64(0); i < uint64(shards); {
		s = i
		key = key*uint64(2862933555777941757) + 1
		i = uint64(float64(s+1) * float64(1<<31) / float64((key>>33)+1))
	}
	return s
}
