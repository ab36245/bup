package main

import (
	"fmt"
	"time"
)

type Info struct {
	When time.Time
	What What

	Created  time.Time
	Gid      uint32
	Kind     Kind
	Modified time.Time
	Perm     uint16
	Uid      uint32
	Regular  BupRegular
	Symlink  BupSymlink
}

func (i Info) String() string {
	s := "BupInfo {\n"

	field := func(name string, mesg string, args ...any) {
		if len(args) > 0 {
			mesg = fmt.Sprintf(mesg, args...)
		}
		s += fmt.Sprintf("  %-10s %s\n", name, mesg)

	}

	field("When", i.When.Format(time.DateTime))
	// field("What", "%s", i.What)

	if i.What&iwCreated == iwCreated {
		field("Created", i.Created.Format(time.DateTime))
	}

	if i.What&iwGid == iwGid {
		field("Gid", "%d", i.Gid)
	}

	if i.What&iwKind == iwKind {
		field("Kind", "%s", i.Kind)
	}

	if i.What&iwModified == iwModified {
		field("Modified", i.Modified.Format(time.DateTime))
	}

	if i.What&iwPerm == iwPerm {
		field("Perm", "0%03o", i.Perm)
	}

	if i.What&iwUid == iwUid {
		field("Uid", "%d", i.Uid)
	}

	if i.What&iwRegular == iwRegular {
		field("Hash", fmt.Sprintf("%x", i.Regular.Hash))
	}

	if i.What&iwSymlink == iwSymlink {
		field("Target", "%q", i.Symlink.Target)
	}

	s += "}"
	return s
}
