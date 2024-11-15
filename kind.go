package main

type Kind uint16

const (
	ikBlockDevice = iota
	ikCharDevice
	ikDir
	ikNamedPipe
	ikRegular
	ikSocket
	ikSymlink
)

func (k Kind) String() string {
	switch k {
	case ikBlockDevice:
		return "ikBlockDevice"
	case ikCharDevice:
		return "ikCharDevice"
	case ikDir:
		return "ikDir"
	case ikNamedPipe:
		return "ikNamedPipe"
	case ikRegular:
		return "ikRegular"
	case ikSocket:
		return "ikSocket"
	case ikSymlink:
		return "ikSymlink"
	default:
		return "???"
	}
}
