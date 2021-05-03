package main

import (
	"fmt"
	models "github.com/pzentenoe/poc-chain-of-responsability/domain/usecase"
	"gopkg.in/yaml.v3"
	"io"
	"net/http"
)

func main() {
	// Read and unmarshall the pipeline configuration from the YAML file
	pipelineConfigFile := "https://gist.githubusercontent.com/thomaspoignant/2499a88c939f654c7e15295194445fd7/raw/0c1b1f5c3ba0a0c73c121f8f002317ae87d04b7d/pipeline.yaml"
	resp, _ := http.Get(pipelineConfigFile)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var pipelineConfig models.PipelineConfig
	_ = yaml.Unmarshal(body, &pipelineConfig)

	// Create the pipeline from the config
	pipeline, _ := models.NewPipeline(pipelineConfig)

	// You can use a context object that can be used by every step
	context := make([]string, 0)
	pipeline.Execute(&context)

	// Check what all steps have done
	fmt.Println(context)
}
