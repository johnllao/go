package zipfs

import "archive/zip"
import "errors"
import "io/ioutil"
import "net/http"
import "bytes"

// ZipFileSystem implements a http FileSystem based on ZIP
type ZipFileSystem struct {
	filename string
}

// NewZipFileSystem creates new instance of ZipFileSystem
func NewZipFileSystem(f string) ZipFileSystem {
	return ZipFileSystem {
		filename: f,
	}
}

// Open opens a file
func (fs ZipFileSystem) Open(name string) (http.File, error) {

	r, err := zip.OpenReader(name)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	for _, f := range r.File {
		if !f.FileInfo().IsDir() && name == "/"+f.Name {
			rc, err := f.Open()
			if err != nil {
				return nil, err
			}
			contents, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, err
			}
			return newZipFile(f.Name, f.FileInfo(), bytes.NewReader(contents)), nil
		}
	}

	return nil, errors.New("file not found "+name)
}