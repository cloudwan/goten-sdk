package bootstrap

func (m *Service_Proto_Package) FQN() string {
	if m.GetCurrentVersion() != "" {
		return m.GetName() + "." + m.GetCurrentVersion()
	} else {
		return m.GetName()
	}
}
