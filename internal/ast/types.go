package ast

import "google.golang.org/protobuf/reflect/protoreflect"

type FullName = protoreflect.FullName

type Node interface {
	PrinterTo
}
