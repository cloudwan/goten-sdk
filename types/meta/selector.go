package meta

// LabelSelector test to match meta of selected object
// object meta may be null. In case no labels are selector labels are specified
// it will match everything
func (ls *LabelSelector) Matches(meta *Meta) bool {
	lbls := meta.GetLabels()
	if lbls == nil {
		lbls = make(map[string]string)
	}
	for mk, mv := range ls.GetMatchLabels() {
		if lbls[mk] != mv {
			return false
		}
	}
	return true
}
