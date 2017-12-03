//oauth2_github
package git

import (
    "path/filepath"
    "os"
    "errors"
    "os/exec"
    "time"
    "strings"
)

//
const messageSeparator  = "#"

//
type Repository struct {
    *repository
}

//
type repository struct {
    absolutePath string
}

//
func NewRepository(absolutePath string) *Repository {
    return &Repository{&repository{absolutePath}}
}

//
func exists(path string) bool {
    _, err := os.Stat(path)
    return !os.IsNotExist(err)
}

//
func (repository *repository) Init() error {

    err := repository.CreateDirectory("")
    if err != nil {
        return err
    }

    cmd := exec.Command("git","init")
    cmd.Dir = repository.absolutePath
    return cmd.Run()
}

//
func (repository *repository) CreateDirectory(directory string) error {

    path := repository.absolutePath
    if directory != "" {
        path = filepath.Join(path,directory)
    }

    if exists(path){
        return errors.New("[ERROR] trying create a existing directory")
    }else {
        err := os.Mkdir(path, os.ModeDevice)
        if err != nil {
            return err
        }
    }
    return nil
}

//
func (repository *repository) SaveFile(file string, content string) error {

    filePath := filepath.Join(repository.absolutePath, file)
    osFile, err := os.OpenFile(filePath, os.O_RDWR | os.O_CREATE | os.O_TRUNC , os.ModeDevice)

    if err == nil {
        _, err = osFile.WriteString(content)
    }

    return err
}

//
func (repository *repository) CreateVersion(file string, userId string) error {

    path := filepath.Join(repository.absolutePath,file)
    if !exists(path) {
        return errors.New("[ERROR] the file does not exist")
    }

    addCmd := exec.Command("git","add",file)
    addCmd.Dir = repository.absolutePath
    if err := addCmd.Run(); err != nil  {
        return err
    }

    message := strings.Join([]string{
        string(time.Now().Nanosecond()),
        userId,
        }, messageSeparator)

    commitCmd := exec.Command("git","commit","-m",message)
    commitCmd.Dir = repository.absolutePath
    return commitCmd.Run()
}
