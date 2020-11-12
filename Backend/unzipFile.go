package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Unzip(src string, destination string) ([]string, error) {

	var filenames []string
	r, err := zip.OpenReader(src)

	if err != nil {
		return filenames, err
	}

	defer r.Close()

	// Executes until files in the source directory keep storing filenames and extracts into destination until an error
	for _, f := range r.File {

		// Store "path/filename" for returning and using later on
		fpath := filepath.Join(destination, f.Name)

		// Checking for any invalid file paths
		if !strings.HasPrefix(fpath, filepath.Clean(destination)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s is an illegal filepath", fpath)
		}

		//append the accessed filenames
		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {

			// Creating a new Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Creating the files in the target directory
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		// The created file will be stored in
		// outFile with permissions to write &/or truncate
		outFile, err := os.OpenFile(fpath,
			os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
			f.Mode())

		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()

		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		//close the file without defer
		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}
	}

	return filenames, nil
}
