package godlp

import (
	"context"
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/romssc/godlp/internal/utils"
)

type App struct {
	config config
}

func (a *App) Start() error {
	_, err := utils.Exec(
		context.TODO(),
		map[string]string{
			"PKG_DIR": a.config.py.dir,
		},
		"bash",
		[]string{
			a.config.py.script,
		},
	)
	if err != nil {
		return fmt.Errorf("%w. %v", ErrStarting, err)
	}
	utils.StartupMessage(a.config.version, a.config.author, a.config.fileOutputFormat)
	return nil
}

func (a *App) Close() error {
	return nil
}

type config struct {
	py py

	version string
	author  string

	fileOutputFormat string
}

type py struct {
	dir    string
	python string
	worker string
	script string
}

/*

FETCHING METHODS

*/

type Metadata struct {
	Title      string      `json:"title"`
	Duration   int         `json:"duration"`
	Uploader   string      `json:"uploader"`
	UploadDate string      `json:"upload_date"`
	Thumbnails []Thumbnail `json:"thumbnails"`
}

type Thumbnail struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	URL    string `json:"url"`
}

type FetchMetadataOptions struct {
	URL string
}

func (a *App) FetchMetadata(ctx context.Context, opts FetchMetadataOptions) (Metadata, error) {
	payload, err := json.Marshal(map[string]string{"url": opts.URL})
	if err != nil {
		return Metadata{}, fmt.Errorf("%w. %v", ErrPreparing, err)
	}
	output, err := utils.Exec(
		ctx,
		nil,
		a.config.py.python,
		[]string{
			a.config.py.worker,
			"metadata",
			string(payload),
		},
	)
	if err != nil {
		return Metadata{}, fmt.Errorf("%w. %v", ErrFetching, err)
	}
	var m Metadata
	if err := json.Unmarshal(output, &m); err != nil {
		return Metadata{}, fmt.Errorf("%w. %v", ErrResponding, err)
	}
	return m, nil
}

type File struct {
	Path string `json:"path"`
	Type string `json:"type"`
}

type QualityOptions string

var (
	LowQuality    QualityOptions = "bestvideo[height<=480]+bestaudio/best[height<=480]"
	MediumQuality QualityOptions = "bestvideo[height<=720]+bestaudio/best[height<=720]"
	HighQuality   QualityOptions = "bestvideo[height<=1080]+bestaudio/best[height<=1080]"
	BestQuality   QualityOptions = "bestvideo+bestaudio/best"
)

type FetchFileOptions struct {
	Quality QualityOptions
	URL     string
}

func (a *App) FetchFile(ctx context.Context, store string, opts FetchFileOptions) (File, error) {
	payload, err := json.Marshal(map[string]string{
		"url":    opts.URL,
		"output": filepath.Join(store, a.config.fileOutputFormat),
		"format": string(opts.Quality),
	})
	if err != nil {
		return File{}, fmt.Errorf("%w. %v", ErrPreparing, err)
	}
	output, err := utils.Exec(
		ctx,
		nil,
		a.config.py.python,
		[]string{
			a.config.py.worker,
			"download",
			string(payload),
		})
	if err != nil {
		return File{}, fmt.Errorf("%w. %v", ErrFetching, err)
	}
	var f File
	if err := json.Unmarshal(output, &f); err != nil {
		return File{}, fmt.Errorf("%w. %v", ErrResponding, err)
	}
	return f, nil
}
