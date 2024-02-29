package std_decorators

import (
	"context"
	"hash/fnv"
	"strings"

	"github.com/cloudwan/goten-sdk/runtime/resource"
	"github.com/cloudwan/goten-sdk/runtime/resource/sharding"
)

const (
	ShardingByServiceIdLabel = "byServiceId"

	ServiceNamePatternPrefix = "services/{service}"
)

type shardingByServiceIdAlg struct{}

func (c *shardingByServiceIdAlg) Label() string {
	return ShardingByServiceIdLabel
}

type ServiceIdExtractorShardingOverride interface {
	resource.Resource
	GetServiceIdForSharding() string
}

func (c *shardingByServiceIdAlg) Shardable(res resource.Resource) bool {
	if strings.HasPrefix(string(res.GetRawName().GetPattern()), ServiceNamePatternPrefix) {
		return true
	}
	if _, ok := res.(ServiceIdExtractorShardingOverride); ok {
		return true
	}
	return false
}

func (c *shardingByServiceIdAlg) ComputeHash(_ context.Context, res resource.Resource) (uint64, bool) {
	if override, ok := res.(ServiceIdExtractorShardingOverride); ok {
		if serviceId := override.GetServiceIdForSharding(); serviceId != "" {
			hasher := fnv.New64a()
			_, _ = hasher.Write([]byte(serviceId))
			return hasher.Sum64(), true
		}
	}
	if serviceId := res.GetRawName().GetIdParts()["serviceId"]; serviceId != "" {
		hasher := fnv.New64a()
		_, _ = hasher.Write([]byte(serviceId))
		return hasher.Sum64(), true
	}
	return 0, false
}

// NewShardingByServiceId gives decorator which assigns shards based solely on service ID.
// All ring configs must contain one FieldPath object with Label = "byServiceId".
func NewShardingByServiceId(rings ...*sharding.RingConfig) *sharding.Decorator {
	return sharding.NewDecorator([]sharding.Algorithm{&shardingByServiceIdAlg{}}, rings)
}
