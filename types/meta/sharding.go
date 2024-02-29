package meta

import (
	"fmt"
)

func MakeShardFullKey(label string, ringSize int) string {
	return fmt.Sprintf("%sMod%d", label, ringSize)
}

func (m *Meta) GetShardValue(label string, ringSize int) int64 {
	key := fmt.Sprintf("%sMod%d", label, ringSize)
	return m.GetShards()[key]
}

func (m *Meta) SetShardValueByLabel(label string, ringSize int, hash int64) {
	key := fmt.Sprintf("%sMod%d", label, ringSize)
	if m.Shards == nil {
		m.Shards = make(map[string]int64)
	}
	m.Shards[key] = hash % int64(ringSize)
}

func (m *Meta) SetShardValueByKey(key string, value int64) {
	if m.Shards == nil {
		m.Shards = make(map[string]int64)
	}
	m.Shards[key] = value
}
