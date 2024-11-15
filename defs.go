package main

type BupRegular struct {
	Hash [32]byte
}

type BupSymlink struct {
	Target string
}
