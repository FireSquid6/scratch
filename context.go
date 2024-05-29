package main

import (
	"github.com/goccy/go-yaml"
	"github.com/otiai10/copy"
	"os"
	"path"
	"strings"
)

type CommandContext struct {
	fs   Filesystem
	conf Config
}

type Filesystem interface {
	Read(path string) (File, error)
	Write(path string, text string) error
	Exists(path string) bool
	Delete(path string) error
	GetFilepathsInDirectory(path string) ([]string, error)
	Copy(src string, dest string) error
}

type File struct {
	text string
	path string
}

type Config struct {
	ScratchDirectory   string `yaml:"scratch_directory"`
	TemplatesDirectory string `yaml:"templates_directory"`
	ProjectsDirectory  string `yaml:"projects_directory"`
}

func getDeafultConfig() Config {
	return Config{
		ScratchDirectory:   "~/scratch",
		TemplatesDirectory: "~/templates",
		ProjectsDirectory:  "~/source",
	}
}

func formatPath(p string) string {
	home := os.Getenv("HOME")

	if strings.HasPrefix(p, "~/") {
		return path.Join(home, p[2:])
	}

	return p
}

func fakeFormatPath(p string) string {
	if strings.HasPrefix(p, "~/") {
		return path.Join("/home/user", p[2:])
	}

	return p
}

type RealFilesystem struct{}

func (fs *RealFilesystem) Read(path string) (File, error) {
	p := formatPath(path)
	f, err := os.ReadFile(p)
	if err != nil {
		return File{}, err
	}

	return File{text: string(f), path: p}, nil
}

func (fs *RealFilesystem) Write(path string, text string) error {
	p := formatPath(path)
	return os.WriteFile(p, []byte(text), 0644)
}

func (fs *RealFilesystem) Exists(path string) bool {
	p := formatPath(path)
	_, err := os.Stat(p)
	return err == nil
}

func (fs *RealFilesystem) Delete(path string) error {
	p := formatPath(path)
	return os.Remove(p)
}

func (fs *RealFilesystem) Copy(src string, dest string) error {
	src = formatPath(src)
	dest = formatPath(dest)
	return copy.Copy(src, dest)
}


func (fs *RealFilesystem) GetFilepathsInDirectory(path string) ([]string, error) {
	p := formatPath(path)
	files, err := os.ReadDir(p)

	filepaths := []string{}
	for _, file := range files {
		filepaths = append(filepaths, file.Name())
	}

	return filepaths, err
}

func GetContext() CommandContext {
	fs := &RealFilesystem{}
	conf := getDeafultConfig()

	return CommandContext{fs: fs, conf: conf}
}

func readConfig(fs Filesystem) Config {
	file, err := fs.Read("~/.config/scratch.yaml")
	if err != nil {
		return getDeafultConfig()
	}

	config := getDeafultConfig()
	yaml.Unmarshal([]byte(file.text), &config)

	return config
}
