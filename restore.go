// This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain Dedication
// license. Its contents can be found at:
// http://creativecommons.org/publicdomain/zero/1.0/

package bindata

import (
	"fmt"
	"io"
)

func writeRestore(w io.Writer) error {
	_, err := fmt.Fprintf(w, `
// RestoreAsset restores an asset to the given file path
func RestoreAsset(writeOutPath, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(filepath.Dir(writeOutPath), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(writeOutPath, data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(writeOutPath, info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(_filePath(dir,child), filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

`)
	return err
}
