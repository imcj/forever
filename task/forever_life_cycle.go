package task

type ForeverLifeCycle int

const (
	ForeverLifeCycleStart ForeverLifeCycle = iota
	ForeverLifeCycleRunning
	ForeverLifeCycleStop
)
