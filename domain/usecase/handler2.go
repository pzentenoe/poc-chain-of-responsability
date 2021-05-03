package usecase

type HandlerImpl2 struct {
	next Handler
}

func (e *HandlerImpl2) Init(name string, step *PipelineStep, availableHandlers map[string]Handler) error {
	if step.Next != "" {
		e.next = availableHandlers[step.Next]
	}
	return nil
}

func (e *HandlerImpl2) Execute(context *[]string) error {
	*context = append(*context, "HandlerImpl2 called")
	if e.next != nil {
		return e.next.Execute(context)
	}
	return nil
}
