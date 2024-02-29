package std_decorators

import (
	"context"
	"hash/fnv"

	"github.com/cloudwan/goten-sdk/runtime/resource"
	"github.com/cloudwan/goten-sdk/runtime/resource/sharding"
)

const (
	ShardingByNameLabel = "byName"
)

type shardingByNameAlg struct{}

func (c *shardingByNameAlg) Label() string {
	return ShardingByNameLabel
}

func (c *shardingByNameAlg) Shardable(res resource.Resource) bool {
	return res.GetResourceDescriptor().SupportsMetadata()
}

func (c *shardingByNameAlg) ComputeHash(_ context.Context, res resource.Resource) (uint64, bool) {
	hasher := fnv.New64a()
	_, _ = hasher.Write([]byte(res.GetRawName().String()))
	return hasher.Sum64(), true
}

// NewShardingByName gives decorator which assigns shards based solely resource name.
// All ring configs must contain one FieldPath object with Label = "byName".
func NewShardingByName(rings ...*sharding.RingConfig) *sharding.Decorator {
	return sharding.NewDecorator([]sharding.Algorithm{&shardingByNameAlg{}}, rings)
}

func NewShardingByNameAlg() sharding.Algorithm {
	return &shardingByNameAlg{}
}
