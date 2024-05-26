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

func fakeFormatPath(p string) (string) {
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


func GetContext() CommandContext {
  fs := &RealFilesystem{}
  conf := getDeafultConfig()

  return CommandContext{fs: fs, conf: conf}
}
