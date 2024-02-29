package service

func ReferencesContains(refs []*Reference, desired *Reference) bool {
	for _, svcRef := range refs {
		if svcRef.GotenEqual(desired) {
			return true
		}
	}
	return false
}

func (s *Service) GetVersionOfImportedService(importingSvcVersion string, imported *Name) string {
	for _, importInfo := range s.GetImportedVersions() {
		if importInfo.GetTargetService().ServiceId == imported.ServiceId &&
			importInfo.GetCurrentServiceVersion() == importingSvcVersion {
			return importInfo.GetTargetServiceVersion()
		}
	}
	return ""
}
