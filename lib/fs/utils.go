package fs

import (
	"context"
	"os"
	"path/filepath"

	"golang.org/x/sync/errgroup"
)

func WalkDir(path string, cb func(path string) error) error {
	group, _ := errgroup.WithContext(context.Background())

	err := filepath.Walk(path, func(Path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			group.Go(func() error {
				err := cb(Path)
				return err
			})
		}
		return nil
	})

	if err != nil {
		return err
	}

	if err := group.Wait(); err != nil {
		return err
	}

	return nil
}
