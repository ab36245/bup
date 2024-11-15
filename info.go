package main

import (
	"fmt"
	"time"
)

type Info struct {
	When time.Time
	What InfoWhat

	Created  time.Time
	Gid      uint32
	Kind     InfoKind
	Modified time.Time
	Perm     uint16
	Uid      uint32
	Regular  BupRegular
	Symlink  BupSymlink
}

type InfoKind uint16

const (
	ikBlockDevice = iota
	ikCharDevice
	ikDir
	ikNamedPipe
	ikRegular
	ikSocket
	ikSymlink
)

type InfoWhat uint16

const (
	iwCreated = 1 << iota
	iwGid
	iwKind
	iwModified
	iwPerm
	iwUid
	iwRegular
	iwSymlink
)

func (i Info) String() string {
	f := "  %-10s %s,\n"
	s := "BupInfo {\n"

	s += fmt.Sprintf(f, "When", i.When.Format(time.DateTime))

	// var w []string
	// addWhat := func(f InfoWhat, s string) {
	// 	if i.What&f == f {
	// 		w = append(w, s)
	// 	}
	// }
	// addWhat(iwCreated, "iwCreated")
	// addWhat(iwGid, "iwGid")
	// addWhat(iwKind, "iwKind")
	// addWhat(iwModified, "iwModified")
	// addWhat(iwPerm, "iwPerm")
	// addWhat(iwUid, "iwUid")
	// addWhat(iwRegular, "iwRegular")
	// addWhat(iwSymlink, "iwCreated")
	// s += fmt.Sprintf(f, "What", strings.Join(w, ", "))

	if i.What&iwCreated == iwCreated {
		s += fmt.Sprintf(f, "Created", i.Created.Format(time.DateTime))
	}

	if i.What&iwGid == iwGid {
		s += fmt.Sprintf(f, "Gid", fmt.Sprintf("%d", i.Gid))
	}

	if i.What&iwKind == iwKind {
		var k string
		switch i.Kind {
		case ikBlockDevice:
			k = "ikBlockDevice"
		case ikCharDevice:
			k = "ikCharDevice"
		case ikDir:
			k = "ikDir"
		case ikNamedPipe:
			k = "ikNamedPipe"
		case ikRegular:
			k = "ikRegular"
		case ikSocket:
			k = "ikSocket"
		case ikSymlink:
			k = "ikSymlink"
		default:
			k = "???"
		}
		s += fmt.Sprintf(f, "Kind", k)
	}

	if i.What&iwModified == iwModified {
		s += fmt.Sprintf(f, "Modified", i.Modified.Format(time.DateTime))
	}

	if i.What&iwPerm == iwPerm {
		s += fmt.Sprintf(f, "Perm", fmt.Sprintf("0%03o", i.Perm))
	}

	if i.What&iwUid == iwUid {
		s += fmt.Sprintf(f, "Uid", fmt.Sprintf("%d", i.Uid))
	}

	if i.What&iwRegular == iwRegular {
		s += fmt.Sprintf(f, "Hash", fmt.Sprintf("%x", i.Regular.Hash))
	}

	if i.What&iwSymlink == iwSymlink {
		s += fmt.Sprintf(f, "Target", fmt.Sprintf("%q", i.Symlink.Target))
	}

	s += "}"
	return s
}
