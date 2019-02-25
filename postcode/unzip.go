package postcode

import (
    "archive/zip"
    "io"
    "os"
    "path/filepath"
)

func Unzip(src string, dest string) ([]string, error) {
    var filenames []string

    r, err := zip.OpenReader(src)
    if err != nil {
        return filenames, err
    }
    defer r.Close()

    for _, f := range r.File {
	rc, err := f.Open()
	if err != nil {
            return filenames, err
        }
        defer rc.Close()

        fpath := filepath.Join(dest, f.Name)
        filenames = append(filenames, fpath)
        if f.FileInfo().IsDir() {

            os.MkdirAll(fpath, os.ModePerm)

        }else{
            if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
                return filenames, err
            }
            outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
            if err != nil {
                return filenames, err
            }

            _, err = io.Copy(outFile, rc)
            outFile.Close()
            if err != nil {
                return filenames, err
            }

        }
    }
    return filenames, nil
}