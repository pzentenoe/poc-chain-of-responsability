package usecase

type HandlerImpl1 struct {
	next Handler
}

func (e *HandlerImpl1) Init(name string, step *PipelineStep, availableHandlers map[string]Handler) error {
	// This is a simplified version of the init method, you can check that next step is not it-self
	// and that the handler is available.
	if step.Next != "" {
		e.next = availableHandlers[step.Next]
	}
	return nil
}

func (e *HandlerImpl1) Execute(context *[]string) error {
	// You can add logic before and after the next step is called.
	*context = append(*context, "HandlerImpl1: before the call")
	if e.next != nil {
		_ = e.next.Execute(context)
	}
	*context = append(*context, "HandlerImpl1: after the call")
	return nil
}
