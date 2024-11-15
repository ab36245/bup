package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"io/fs"
	"os"
	"time"
)

func Read(root string) []*Item {
	r := &Reader{
		root: root,
		when: time.Now(),
	}
	r.Read()
	return r.list
}

func readKind(fi fs.FileInfo) Kind {
	switch fi.Mode() & fs.ModeType {
	case 0:
		return ikRegular
	case fs.ModeDir:
		return ikDir
	case fs.ModeSymlink:
		return ikSymlink
	case fs.ModeNamedPipe:
		return ikNamedPipe
	case fs.ModeSocket:
		return ikSocket
	case fs.ModeDevice:
		return ikBlockDevice
	case fs.ModeCharDevice:
		return ikCharDevice
	default:
		panic(fmt.Errorf("unknown mode 0%o\n", fi.Mode()&fs.ModeType))
	}
}

func readModified(fi fs.FileInfo) time.Time {
	return fi.ModTime()
}

func readPerm(fi fs.FileInfo) uint16 {
	return uint16(fi.Mode() & fs.ModePerm)
}

func readRegular(path string) BupRegular {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		panic(err)
	}

	var hash [32]byte
	copy(hash[:], h.Sum(nil))

	return BupRegular{
		Hash: hash,
	}
}

func readSymlink(path string) BupSymlink {
	target, err := os.Readlink(path)
	if err != nil {
		panic(err)
	}
	return BupSymlink{
		Target: target,
	}
}
