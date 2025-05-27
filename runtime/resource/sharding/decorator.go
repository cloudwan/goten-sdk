package sharding

import (
	"context"
	"fmt"

	"github.com/cloudwan/goten-sdk/runtime/resource"
)

type Algorithm interface {
	Label() string
	Shardable(resource.Resource) bool
	ComputeHash(context.Context, resource.Resource) (uint64, bool)
}

type Decorator struct {
	algRingBindings []algorithmRingBindings
}

type algorithmRingBindings struct {
	rings []ringOnPath
	alg   Algorithm
}

type ringOnPath struct {
	size uint64
	path string
}

func NewDecorator(algorithms []Algorithm, rings []*RingConfig) *Decorator {
	assigner := &Decorator{}
	for _, algorithm := range algorithms {
		algRingsBinding := algorithmRingBindings{alg: algorithm}
		for _, ringCfg := range rings {
			var selectedFp *RingConfig_FieldPath
			for _, fp := range ringCfg.FieldPaths {
				if fp.Label == algorithm.Label() {
					selectedFp = fp
					break
				}
			}
			if selectedFp == nil {
				panic(fmt.Sprintf(
					"Decorator configuration error: Could not find field path with label %s",
					algorithm.Label()))
			}
			algRingsBinding.rings = append(algRingsBinding.rings, ringOnPath{size: ringCfg.Size, path: selectedFp.Path})
		}
		assigner.algRingBindings = append(assigner.algRingBindings, algRingsBinding)
	}
	return assigner
}

func (a *Decorator) AssignShards(ctx context.Context, resource resource.Resource) {
	for _, algWithRings := range a.algRingBindings {
		if !algWithRings.alg.Shardable(resource) {
			continue
		}

		hash, set := algWithRings.alg.ComputeHash(ctx, resource)
		for _, ring := range algWithRings.rings {
			fp, err := resource.GetResourceDescriptor().ParseFieldPath(ring.path)
			if err == nil {
				if set {
					fp.WithRawIValue(int64(hash % ring.size)).SetToRaw(resource)
				} else {
					fp.ClearValueRaw(resource)
				}
			} else {
				panic(fmt.Errorf("cannot parse field path %s for resource of type %s",
					ring.path, resource.GetResourceDescriptor().GetResourceTypeName().FullyQualifiedTypeName()))
			}
		}
	}
}
