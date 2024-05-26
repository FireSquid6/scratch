package main

import (
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

func formatPath(p string) (string) {
  home := os.Getenv("HOME")

  if strings.HasPrefix(p, "~/") {
    return path.Join(home, p[2:])
  }

  return p
}

type RealFilesystem struct{}

func (fs *RealFilesystem) Open(path string) (File, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return File{}, err
	}

	return File{text: string(f), path: path}, nil
}

func (fs *RealFilesystem) Write(path string, text string) error {
	return os.WriteFile(path, []byte(text), 0644)
}

func (fs *RealFilesystem) Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func (fs *RealFilesystem) Delete(path string) error {
	return os.Remove(path)
}

type FakeFilesystem struct {
	files map[string]string
}

func (fs *FakeFilesystem) Read(path string) (File, error) {
	return File{text: fs.files[path], path: path}, nil
}

func (fs *FakeFilesystem) Write(path string, text string) error {
	fs.files[path] = text
	return nil
}

func (fs *FakeFilesystem) Exists(path string) bool {
	_, ok := fs.files[path]
	return ok
}
