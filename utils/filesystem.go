// TODO
package utils

import (
    "os"
)

// Exists checks if a file or directory exists using absolutePath.
func Exists(absolutePath string) bool {
    _, err := os.Stat(absolutePath)
    return !os.IsNotExist(err)
}

// TODO
func Write(absolutePath string, content string) error {

    osFile, err := os.OpenFile(absolutePath, os.O_RDWR | os.O_CREATE | os.O_TRUNC , os.ModeDevice)
    if err == nil {
        _, err = osFile.WriteString(content)
    }

    return err
}
