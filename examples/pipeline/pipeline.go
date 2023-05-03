package main

import (
	"github.com/myntra/pipeline"
	slogpipeline "github.com/phxvlabs.dev/sdk-go/v1/slog-pipeline"
)

type work struct {
	pipeline.StepContext
	id int
}

func (w *work) Exec(request *pipeline.Request) *pipeline.Result {
	w.Status
}

func main() {
	logger := slogpipeline.NewPipelineChannel("app")
	p, err := pipeline.New(logger)

	if err != nil {
		return err
	}
}
