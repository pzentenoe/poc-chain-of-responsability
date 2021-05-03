package usecase

import (
	"fmt"
)



type Pipeline struct {
	root     string
	steps    map[string]*PipelineStep
	handlers map[string]Handler
}

// NewPipeline will create a new Pipeline ready to be executed
func NewPipeline(pipelineConfig PipelineConfig) (Pipeline, error) {
	p := Pipeline{
		steps: pipelineConfig.Steps,
		root:  pipelineConfig.Root,
	}

	// Get handlers from  config
	p.handlers = make(map[string]Handler, len(p.steps))
	for name, step := range p.steps {
		handler, _ := p.getHandlerFromType(step.StepType)
		p.handlers[name] = handler
	}

	// Init all handlers
	for name, step := range p.steps {
		err := p.handlers[name].Init(name, step, p.handlers)
		if err != nil {
			return Pipeline{}, fmt.Errorf("impossible to init the step named '%s': %v", name, err)
		}
	}
	// Check that root step exists
	if _, ok := p.handlers[p.root]; !ok {
		return Pipeline{}, fmt.Errorf("impossible to start with step \"%s\" because it does not exists", p.root)
	}
	return p, nil
}

// getHandlerFromType is mapping handler type name in your configuration to proper handlers.
func (p *Pipeline) getHandlerFromType(s string) (Handler, error) {
	// mapping list for the handlers
	handlers := map[string]Handler{
		"handlerImpl1": &HandlerImpl1{},
		"handlerImpl2": &HandlerImpl2{},
	}

	stepHandler, handlerExists := handlers[s]
	if !handlerExists {
		return nil, fmt.Errorf("impossible to find a matching step handler for %s", s)
	}
	return stepHandler, nil
}

// Execute the search Pipeline by taking the 1st Step and execute it.
func (p *Pipeline) Execute(context *[]string) error {
	return p.handlers[p.root].Execute(context)
}


// PipelineConfig is the representation of a pipeline in the configuration.
type PipelineConfig struct {
	// Steps is the list of step in your pipeline.
	Steps map[string]*PipelineStep `yaml:"steps"`

	// Root is the name of the first step in your pipeline, we will start by calling it, and it will call the next steps after it.
	Root string `yaml:"root"`
}


// PipelineStep is a step representation in the configuration.
type PipelineStep struct {
	// StepType is the type of the Handler to map for this step configuration, the list of
	// the available types is in the method getHandlerFromType
	StepType string `yaml:"type"`

	// Next is the next step we should call after this one. This param is not mandatory.
	Next string `yaml:"next"`
}
