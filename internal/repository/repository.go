package repository

import (
	"errors"
	"github.com/marco2704/zeus/internal/config"
	"github.com/marco2704/zeus/pkg/filesystem"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// messageSeparator is used to separate data in the commit message.
const messageSeparator = "#"

// Repository represents the git repository that each users has.
type Repository struct {
	*repository
}

// repository is a unexported struct that have all Repository fields.
type repository struct {
	absolutePath string       `json:"-"`
	Directories  *[]Directory `json:"directories,omitempty"`
	Files        *[]file      `json:"files,omitempty"`
}

// File represents a document.
type file struct {
	Name string
}

// Directory represents a folder which may store File and another Directory.
type Directory struct {
	Name        string
	directories *[]Directory
	files       *[]file
}

// NewRepository creates a instance of Repository using the userId as folder name.
func NewRepository(userId string) *Repository {

	return &Repository{
		&repository{
			filepath.Join(config.Config.RepoRootPath(), userId),
			nil,
			nil,
		},
	}
}

// Init creates a git repository using the absolute path.
func (repository *repository) Init() error {

	err := repository.createDirectory("")
	if err != nil {
		return err
	}

	cmd := exec.Command("git", "init")
	cmd.Dir = repository.absolutePath
	return cmd.Run()
}

// createDirectory creates a directory using string directory parameter into repository directory.
// If directory parameter is empty the repository directory is created (called by Init method).
func (repository *repository) createDirectory(directory string) error {

	path := repository.absolutePath
	if directory != "" {
		path = filepath.Join(path, directory)
	}

	if filesystem.Exists(path) {
		return errors.New("error - the directory already exists")
	} else {
		err := os.Mkdir(path, os.ModeDevice)
		if err != nil {
			return err
		}
	}
	return nil
}

// SaveFile open or creates if it does not exist a file with the given name in
// the repository, then fil the content given into it. A no nil error is returned
// if something went wrong.
func (repository *repository) SaveFile(file string, content string) error {
	return filesystem.Write(filepath.Join(repository.absolutePath, file), content)
}

// CreateVersion creates a git commit which only includes the given file (if it exists).
func (repository *repository) CreateVersion(file string, userId string) error {

	path := filepath.Join(repository.absolutePath, file)
	if !filesystem.Exists(path) {
		return errors.New("error - the file does not exist")
	}

	addCmd := exec.Command("git", "add", file)
	addCmd.Dir = repository.absolutePath
	if err := addCmd.Run(); err != nil {
		return err
	}

	message := strings.Join([]string{
		string(time.Now().Nanosecond()),
		userId,
	}, messageSeparator)

	commitCmd := exec.Command("git", "commit", "-m", message)
	commitCmd.Dir = repository.absolutePath
	return commitCmd.Run()
}
