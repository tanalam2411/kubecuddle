package controller

import (
	"github.com/tanalam2411/kubecuddle/pkg/controller/tgik"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, tgik.Add)
}
