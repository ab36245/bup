package main

type BupItem struct {
	Path string
	Info []*Info
}

type BupRegular struct {
	Hash [32]byte
}

type BupSymlink struct {
	Target string
}
