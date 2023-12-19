package wd

import (
	"context"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/k0sproject/rig/pkg/rigfs"
	"github.com/octohelm/unifs/pkg/filesystem"
	"golang.org/x/net/webdav"
)

func WrapRigFS(fsys rigfs.Fsys) filesystem.FileSystem {
	return &rigfsWrapper{fsys: fsys}
}

type rigfsWrapper struct {
	fsys rigfs.Fsys
}

func (r *rigfsWrapper) Mkdir(ctx context.Context, name string, perm os.FileMode) error {
	return r.fsys.MkDirAll(name, perm)
}

func (r *rigfsWrapper) OpenFile(ctx context.Context, name string, flag int, perm os.FileMode) (webdav.File, error) {
	if perm.IsDir() {
		info, err := r.fsys.Stat(name)
		if err != nil {
			return nil, err
		}

		return &wrapper{
			fsys: r.fsys,
			path: name,
			info: info,
		}, err
	}

	f, err := r.fsys.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return &wrapper{
		fsys: r.fsys,
		path: name,
		file: f,
	}, err
}

var _ fs.ReadDirFile = &wrapper{}

type wrapper struct {
	fsys rigfs.Fsys
	path string
	info fs.FileInfo
	file fs.File
}

func (f *wrapper) Close() error {
	if f.file != nil {
		return f.file.Close()
	}
	return nil
}

func (f *wrapper) Stat() (fs.FileInfo, error) {
	if f.info != nil {
		return f.info, nil
	}
	return f.file.Stat()
}

func (f *wrapper) Read(bytes []byte) (int, error) {
	if seeker, ok := f.file.(io.Reader); ok {
		return seeker.Read(bytes)
	}
	return -1, &fs.PathError{
		Op:   "read",
		Path: f.path,
		Err:  fs.ErrInvalid,
	}
}

func (f *wrapper) Seek(offset int64, whence int) (int64, error) {
	if seeker, ok := f.file.(io.Seeker); ok {
		return seeker.Seek(offset, whence)
	}
	return -1, &fs.PathError{
		Op:   "seek",
		Path: f.path,
		Err:  fs.ErrInvalid,
	}
}

func (f *wrapper) Write(p []byte) (n int, err error) {
	if writer, ok := f.file.(io.Writer); ok {
		return writer.Write(p)
	}
	return -1, &fs.PathError{
		Op:   "write",
		Path: f.path,
		Err:  fs.ErrPermission,
	}
}

func (f *wrapper) ReadDir(n int) ([]fs.DirEntry, error) {
	if ff, ok := f.file.(fs.ReadDirFile); ok {
		if n < 0 {
			n = 0
		}
		return ff.ReadDir(n)
	}
	return fs.ReadDir(f.fsys, f.path)
}

func (f *wrapper) Readdir(count int) ([]fs.FileInfo, error) {
	list, err := f.ReadDir(count)
	if err != nil {
		return nil, err
	}

	infos := make([]fs.FileInfo, len(list))
	for i := range list {
		info, err := list[i].Info()
		if err != nil {
			return nil, err
		}
		infos[i] = info
	}

	return infos, nil
}

func (r *rigfsWrapper) RemoveAll(ctx context.Context, name string) error {
	return r.fsys.RemoveAll(name)
}

func (r *rigfsWrapper) Stat(ctx context.Context, name string) (os.FileInfo, error) {
	return r.fsys.Stat(name)
}

func (r *rigfsWrapper) Rename(ctx context.Context, oldName, newName string) (err error) {
	oldFile, err := r.OpenFile(ctx, oldName, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer func() {
		_ = oldFile.Close()
		if err != nil {
			_ = r.RemoveAll(ctx, oldName)
		}
	}()

	if err := r.Mkdir(ctx, filepath.Dir(newName), os.ModeDir); err != nil {
		return err
	}

	newFile, err := r.OpenFile(ctx, oldName, os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}

	defer newFile.Close()
	if _, err := io.Copy(newFile, oldFile); err != nil {
		return err
	}
	return nil
}
