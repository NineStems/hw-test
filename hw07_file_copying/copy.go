package main

import (
	"errors"
	"github.com/cheggaaa/pb/v3"
	"io"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	ff, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer ff.Close()

	stat, err := ff.Stat()
	if err != nil {
		return err
	}

	if stat.Size() < offset {
		return ErrOffsetExceedsFileSize
	}

	if limit == 0 || limit > stat.Size() {
		limit = stat.Size()
	}

	bar := pb.Full.Start64(limit)
	sectionRead := io.NewSectionReader(ff, offset, limit)

	ft, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer ft.Close()

	_, err = io.Copy(ft, bar.NewProxyReader(sectionRead))
	if err != nil {
		return err
	}
	bar.Finish()

	return nil
}
