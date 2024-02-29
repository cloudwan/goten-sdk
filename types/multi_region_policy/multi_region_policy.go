package multi_region_policy

import (
	"sort"
)

func (p *MultiRegionPolicy) RegionEnabled(name string) bool {
	for _, availableRegion := range p.GetEnabledRegions() {
		if availableRegion == name {
			return true
		}
	}
	return false
}

func (p *MultiRegionPolicy) RegionsEnabled(regions []string) bool {
	for _, name := range regions {
		if !p.RegionEnabled(name) {
			return false
		}
	}
	return true
}

func (p *MultiRegionPolicy) CanReadFromRegion(readingRegion, dataRegion string, resFqn string) bool {
	return readingRegion == dataRegion || (!p.forbiddingRuleExists(readingRegion, dataRegion, resFqn) && p.RegionEnabled(readingRegion))
}

func (p *MultiRegionPolicy) CanCreateInRegion(regionId string) bool {
	return p.RegionEnabled(regionId)
}

func (p *MultiRegionPolicy) CanAccessDataFromRegion(readingRegion string, dataRegions []string, resFqn string) bool {
	for _, dataRegion := range dataRegions {
		if !p.CanReadFromRegion(readingRegion, dataRegion, resFqn) {
			return false
		}
	}
	return true
}

func (p *MultiRegionPolicy) ComputeSyncingRegions(resFqn string, owningRegion string) []string {
	regions := p.GetEnabledRegions()
	syncingRegions := make([]string, 0)
	for _, region := range regions {
		if region == owningRegion || !p.forbiddingRuleExists(region, owningRegion, resFqn) {
			syncingRegions = append(syncingRegions, region)
		}
	}
	sort.Strings(syncingRegions)
	return syncingRegions
}

func (p *MultiRegionPolicy) forbiddingRuleExists(readingRegion, dataRegion string, resFqn string) bool {
	for _, criteria := range p.GetCriteriaForDisabledSync() {
		if (criteria.SourceRegion == "" || criteria.SourceRegion == dataRegion) &&
			(criteria.DestRegion == "" || criteria.DestRegion == readingRegion) &&
			(criteria.ResourceTypeName == "" || criteria.ResourceTypeName == resFqn) {
			return true
		}
	}
	return false
}
