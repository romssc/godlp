package godlp

import (
	"path/filepath"
)

type Config struct {
	PkgDir           string
	FileOutputFormat string
}

func New(c Config) *App {
	return &App{
		config: config{
			py: py{
				dir:    c.PkgDir,
				python: filepath.Join(c.PkgDir, ".py/venv/bin/python"),
				worker: filepath.Join(c.PkgDir, ".py/worker.py"),
				script: filepath.Join(c.PkgDir, ".py/setup.sh"),
			},
			version:          "v0.1.0",
			author:           "https://github.com/romssc",
			fileOutputFormat: c.FileOutputFormat,
		},
	}
}
