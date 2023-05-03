module github.com/phxvlabs.dev/sdk-go/examples/pipeline

go 1.20

require (
	github.com/myntra/pipeline v0.0.0-20180618182531-2babf4864ce8
	github.com/phxvlabs.dev/sdk-go/v1/slog-pipeline v0.0.0-unpublished
)

require (
	github.com/fatih/color v1.15.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	golang.org/x/exp v0.0.0-20230418202329-0354be287a23 // indirect
	golang.org/x/sys v0.6.0 // indirect
)

replace github.com/phxvlabs.dev/sdk-go/v1/slog-pipeline => ../../v1/slog-pipeline
