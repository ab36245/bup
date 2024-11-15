package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"time"
)

type Reader struct {
	list []*Item
	root string
	when time.Time
}

func (r *Reader) Read() {
	r.readDir("")
}

func (r *Reader) readDir(dir string) {
	entries, err := os.ReadDir(r.pathFull(dir))
	if err != nil {
		panic(err)
	}

	// Might be too slow??
	sort.Sort(byName(entries))

	for _, entry := range entries {
		r.readItem(r.pathJoin(dir, entry.Name()))
	}
}

func (r *Reader) readItem(path string) {
	full := r.pathFull(path)
	fi, err := os.Lstat(full)
	if err != nil {
		panic(err)
	}

	info := &Info{}

	info.When = r.when

	info.Created = readCreated(fi)
	info.What |= iwCreated

	info.Gid = readGid(fi)
	info.What |= iwGid

	info.Kind = readKind(fi)
	info.What |= iwKind

	info.Modified = readModified(fi)
	info.What |= iwModified

	info.Perm = readPerm(fi)
	info.What |= iwPerm

	info.Uid = readUid(fi)
	info.What |= iwUid

	if info.Kind == ikRegular {
		info.Regular = readRegular(full)
		info.What |= iwRegular
	} else if info.Kind == ikSymlink {
		info.Symlink = readSymlink(full)
		info.What |= iwSymlink
	}

	item := &Item{
		Path: path,
		Info: []*Info{info},
	}

	r.list = append(r.list, item)

	if info.Kind == ikDir {
		r.readDir(item.Path)
	}
}

func (r Reader) pathFull(path string) string {
	return r.pathJoin(r.root, path)
}

func (r Reader) pathJoin(a, b string) string {
	if a == "" {
		return b
	}
	if b == "" {
		return a
	}
	return filepath.Join(a, b)
}

type byName []fs.DirEntry

func (e byName) Len() int {
	return len(e)
}

func (e byName) Less(i, j int) bool {
	return e[i].Name() < e[j].Name()
}

func (e byName) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}
