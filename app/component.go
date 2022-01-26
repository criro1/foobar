package app

type (
	Component struct {
		RunFn   Run
		CloseFn Close
		ReadyFn Ready
	}
)

func (c *Component) Run() error {
	return c.RunFn()
}

func (c *Component) Close() error {
	return c.CloseFn()
}

func (c *Component) Ready() <-chan struct{} {
	if c.ReadyFn == nil {
		return nil
	}

	return c.ReadyFn()
}
