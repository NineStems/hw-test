package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")

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

	ft, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer ft.Close()

	return copyReader(ff, ft, offset, limit, stat.Size())
}

func copyReader(fr io.ReaderAt, fw io.Writer, offset, limit, size int64) error {
	if size < offset {
		return ErrOffsetExceedsFileSize
	}

	if limit == 0 || limit > size {
		limit = size
	}

	sectionRead := io.NewSectionReader(fr, offset, limit)

	switch {
	case offset > 0 && limit == size:
		limit -= offset
	case offset > 0 && size-(limit+offset) < 0:
		limit = size - offset
	}

	bar := pb.Full.Start64(limit)
	_, err := io.Copy(fw, bar.NewProxyReader(sectionRead))
	if err != nil {
		return err
	}
	bar.Finish()
	return nil
}
