package filter

type condOptionsCfg struct {
	negateCond bool
}

type FilterCondOptions interface {
	IsNot() bool
}

func (c *condOptionsCfg) IsNot() bool {
	return c.negateCond
}

func MakeFilterCondOptions(opts []FilterConditionOption) FilterCondOptions {
	condOpts := &condOptionsCfg{
		negateCond: false,
	}
	for _, opt := range opts {
		opt(condOpts)
	}
	return condOpts
}

type FilterConditionOption func(*condOptionsCfg)

func Not() FilterConditionOption {
	return func(c *condOptionsCfg) {
		c.negateCond = true
	}
}
