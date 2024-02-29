package meta

import (
	"sort"
)

func (m *Meta) EnsureServices() *ServicesInfo {
	if m.Services == nil {
		m.Services = &ServicesInfo{}
	}
	return m.Services
}

func (si *ServicesInfo) EnsureHasAllowed(svcId string) {
	found := false
	for _, currentlyAllowed := range si.GetAllowedServices() {
		if currentlyAllowed == svcId {
			found = true
			break
		}
	}
	if !found {
		si.AllowedServices = append(si.AllowedServices, svcId)
		sort.Strings(si.AllowedServices)
	}
}

func (si *ServicesInfo) EnsureRemovedFromAllowed(svcId string) {
	currentlyAllowed := si.GetAllowedServices()
	si.AllowedServices = make([]string, 0, len(currentlyAllowed))
	for _, allowed := range currentlyAllowed {
		if allowed != svcId {
			si.AllowedServices = append(si.AllowedServices, allowed)
		}
	}
}
