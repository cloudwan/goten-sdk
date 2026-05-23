package resource

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

// forbiddenInheritActions is the set of permission action segments that must
// never appear in SystemCatalogSpec.InheritPermissionIds. These are mutating
// verbs whose project/org-scope evaluation must NOT stand in for a real
// system-scope grant — see SystemCatalogSpec docs.
var forbiddenInheritActions = map[string]struct{}{
	"create":             {},
	"update":             {},
	"delete":             {},
	"setIamPolicy":       {},
	"testIamPermissions": {},
}

// ValidateInheritPermissionIds returns a non-nil error if any entry in perms
// names a mutating verb that must never be inherited from project/org scope
// to system scope. The check matches the action segment exactly (text after
// the last '.'), so reference-form permissions like "connectors!attach"
// pass unaffected, and lookalikes like "things.somedotcreate" do not
// false-trip.
func ValidateInheritPermissionIds(perms []string) error {
	for _, p := range perms {
		dot := strings.LastIndex(p, ".")
		if dot < 0 {
			continue
		}
		action := p[dot+1:]
		if _, bad := forbiddenInheritActions[action]; bad {
			return fmt.Errorf(
				"inheritPermissionIds must not include mutating verb: %q "+
					"(create/update/delete/setIamPolicy/testIamPermissions are "+
					"never inherited from project/org scope)", p)
		}
	}
	return nil
}

// SystemCatalogSpec describes how a resource kind participates in the
// "shared system catalog" pattern: system-scoped instances of the kind are
// intended to be visible to every project/org, and references to them from
// project/org-scoped resources are authorized against the caller's
// project/org-scoped permissions rather than requiring a system-scope
// RoleBinding.
//
// Typical use: a platform admin publishes a small, curated set of shared
// resources (e.g. chat models, connectors, capability templates) directly
// under the system parent, and wants ordinary project users to discover and
// reference them without any extra IAM configuration.
//
// The IAM layer consults SystemCatalogSpec via GetSystemCatalog to decide
// whether to rewrite an auth check against a system-scoped object so that it
// is evaluated under the caller's request scope instead.
type SystemCatalogSpec struct {
	// InheritPermissionIds is the whitelist of permission IDs that a
	// project/org-scope RoleBinding may satisfy on behalf of system scope.
	// Entries are raw IAM permission ids WITHOUT the service domain prefix,
	// for example: "chatModels.get", "chatModels.list", "chatModels!attach",
	// "connectors.list", "connectors!attach".
	//
	// Mutating verbs (.create, .update, .delete, .setIamPolicy,
	// .testIamPermissions) MUST NOT be included. Writes to system-scoped
	// instances always require a real system-scope RoleBinding.
	InheritPermissionIds []string
}

// Inherits reports whether the given permission ID (e.g. "chatModels.list"
// or "connectors!attach") is inherited from project/org scope to system
// scope for this kind.
func (s *SystemCatalogSpec) Inherits(permissionId string) bool {
	if s == nil {
		return false
	}
	for _, p := range s.InheritPermissionIds {
		if p == permissionId {
			return true
		}
	}
	return false
}

var (
	systemCatalogMu       sync.RWMutex
	systemCatalogRegistry = map[Descriptor]*SystemCatalogSpec{}
)

// RegisterSystemCatalog marks desc as a shared system catalog with the given
// spec. It is idempotent: registering the same descriptor twice with the same
// spec is allowed; registering with a different spec panics to surface
// inconsistent configuration at startup.
//
// Intended usage is from package init() in the package that owns the
// resource descriptor.
func RegisterSystemCatalog(desc Descriptor, spec *SystemCatalogSpec) {
	if desc == nil || spec == nil {
		return
	}
	if err := ValidateInheritPermissionIds(spec.InheritPermissionIds); err != nil {
		panic(fmt.Sprintf(
			"resource: invalid SystemCatalogSpec for %s: %s",
			desc.GetResourceTypeName().FullyQualifiedTypeName(), err))
	}
	systemCatalogMu.Lock()
	defer systemCatalogMu.Unlock()
	if existing, ok := systemCatalogRegistry[desc]; ok {
		if !sameSpec(existing, spec) {
			panic(fmt.Sprintf(
				"resource: conflicting RegisterSystemCatalog for %s: existing %v vs new %v",
				desc.GetResourceTypeName().FullyQualifiedTypeName(),
				existing.InheritPermissionIds,
				spec.InheritPermissionIds,
			))
		}
		return
	}
	systemCatalogRegistry[desc] = spec
}

// GetSystemCatalog returns the SystemCatalogSpec for desc if one has been
// registered, otherwise nil. The returned spec is the same pointer that was
// passed to RegisterSystemCatalog and MUST NOT be mutated by the caller;
// specs are expected to be immutable for the lifetime of the process.
func GetSystemCatalog(desc Descriptor) *SystemCatalogSpec {
	if desc == nil {
		return nil
	}
	systemCatalogMu.RLock()
	defer systemCatalogMu.RUnlock()
	return systemCatalogRegistry[desc]
}

// IsSystemCatalog is a convenience predicate for callers that only care
// whether the kind participates in the pattern at all.
func IsSystemCatalog(desc Descriptor) bool {
	return GetSystemCatalog(desc) != nil
}

// sameSpec reports whether a and b carry the same inherit-permission
// multiset, independent of order. It is used by RegisterSystemCatalog to
// allow idempotent re-registration when two call sites build the same
// logical spec but list permissions in a different order. Sorted-slice
// equality keeps the comparison correct even if a list happens to contain
// duplicates.
func sameSpec(a, b *SystemCatalogSpec) bool {
	if len(a.InheritPermissionIds) != len(b.InheritPermissionIds) {
		return false
	}
	as := append([]string(nil), a.InheritPermissionIds...)
	bs := append([]string(nil), b.InheritPermissionIds...)
	sort.Strings(as)
	sort.Strings(bs)
	for i := range as {
		if as[i] != bs[i] {
			return false
		}
	}
	return true
}
