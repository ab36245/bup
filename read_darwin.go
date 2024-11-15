package main

import (
	"fmt"
	"io/fs"
	"syscall"
	"time"
)

func readCreated(fi fs.FileInfo) time.Time {
	stat := darwin_stat(fi)
	return time.Unix(stat.Birthtimespec.Sec, stat.Birthtimespec.Nsec)
}

func readGid(fi fs.FileInfo) uint32 {
	stat := darwin_stat(fi)
	return stat.Gid
}

func readUid(fi fs.FileInfo) uint32 {
	stat := darwin_stat(fi)
	return stat.Uid
}

func darwin_stat(fi fs.FileInfo) *syscall.Stat_t {
	stat, ok := fi.Sys().(*syscall.Stat_t)
	if !ok {
		panic(fmt.Errorf("type cast failed???"))
	}
	return stat
}
