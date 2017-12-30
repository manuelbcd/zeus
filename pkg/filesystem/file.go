package filesystem

import (
    "os"
)

// Write opens a existing file or creates if it does not exist, then writes
// the given text into it.
func Write(absolutePath string, content string) error {
    osFile, err := os.OpenFile(absolutePath, os.O_RDWR | os.O_CREATE | os.O_TRUNC, os.ModeDevice)
    if err == nil {
        _, err = osFile.WriteString(content)
    }

    return err
}
