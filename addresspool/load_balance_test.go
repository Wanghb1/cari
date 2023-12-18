package addresspool_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-chassis/cari/addresspool"
)

func BenchmarkRoundRobin(b *testing.B) {
	eps := []string{"172.0.0.1", "172.0.0.2", "172.0.0.3", "172.0.0.4",
		"172.0.0.5", "172.0.0.6", "172.0.0.7", "172.0.0.8", "172.0.0.9",
		"172.0.0.10", "172.0.0.11", "172.0.0.12"}
	next := addresspool.RoundRobin(eps)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = next()
	}
}
func TestLoadbalanceEmpty(t *testing.T) {
	t.Log("Testing Round robin with empty endpoint arrays")
	var sArrEmpty []string

	next := addresspool.RoundRobin(sArrEmpty)
	_, err := next()
	assert.Error(t, err)

}
func TestLoadbalance(t *testing.T) {
	t.Log("Testing Round robin function")
	var sArr []string

	sArr = append(sArr, "s1")
	sArr = append(sArr, "s2")

	next := addresspool.RoundRobin(sArr)
	_, err := next()
	assert.NoError(t, err)
}
