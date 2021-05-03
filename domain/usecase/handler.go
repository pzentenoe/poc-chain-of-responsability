package usecase

// Handler is defining how a step looks like.
type Handler interface {
	// Init configure the step from the configuration file
	Init(name string, step *PipelineStep, availableHandlers map[string]Handler) error

	// Execute apply the action of the step and move to the next step
	Execute(context *[]string) error
}
