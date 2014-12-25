package jmpconsistenthash

import "testing"

func TestShardBounds(t *testing.T) {
	for shards := uint64(0); shards < 100; shards++ {
		for key := uint64(0); key < 10*shards; key++ {
			hash := Hash(key, shards)
			if hash < 0 || hash >= shards {
				t.Errorf("Hash(%d, %d) = %d should be in the range[0,%d)", key, shards, hash, shards)
			}
		}
	}
}

func delta(a uint64, b uint64) uint64 {
	if a > b {
		return a - b
	}
	return b - a
}

func checkDistribution(distribution map[uint64]uint64, ideal uint64) bool {
	sumdelta := uint64(0)
	for _, count := range distribution {
		// Consider any shard with zero keys as unacceptable
		if count == 0 && ideal > 1 {
			return false
		}
		sumdelta += delta(count, ideal)
	}
	// Normalize delta-from-ideal per shard
	normalized := float64(sumdelta) / float64(len(distribution))
	// If normalized delta is more than 5% of ideal, reject the distribution.
	if normalized > float64(0.05)*float64(ideal) {
		return false
	}
	return true
}

// TODO: Write a statistical test of significance instead of this hack.
func TestShardSpreads(t *testing.T) {
	for shards := uint64(0); shards < 100; shards++ {
		distribution := make(map[uint64]uint64)
		for i := uint64(0); i < shards; i++ {
			distribution[i] = 0
		}

		for key := uint64(0); key < 1000*shards; key++ {
			hash := Hash(key, shards)
			distribution[hash]++
		}

		if !checkDistribution(distribution, 1000) {
			t.Errorf("Shards (%d) has unacceptable distribution: %v", shards, distribution)
		}
	}
}
