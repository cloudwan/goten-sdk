package sharding

func (m *RingConfig) GetDefaultFieldPath() string {
	if len(m.GetFieldPaths()) == 0 {
		return ""
	}
	return m.GetFieldPaths()[0].GetPath()
}

func (m *RingConfig) GetFieldPathForLabel(label string) string {
	for _, fp := range m.GetFieldPaths() {
		if fp.GetLabel() == label {
			return fp.Path
		}
	}
	return ""
}
