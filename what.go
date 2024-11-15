package main

import "strings"

type What uint16

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

func (w What) String() string {
	var set []string
	add := func(f What, s string) {
		if w&f == f {
			set = append(set, s)
		}
	}
	add(iwCreated, "iwCreated")
	add(iwGid, "iwGid")
	add(iwKind, "iwKind")
	add(iwModified, "iwModified")
	add(iwPerm, "iwPerm")
	add(iwUid, "iwUid")
	add(iwRegular, "iwRegular")
	add(iwSymlink, "iwCreated")
	return strings.Join(set, ", ")
}
