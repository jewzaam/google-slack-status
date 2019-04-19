package controller

import (
	"github.com/jewzaam/google-slack-status/pkg/controller/googleslackstatus"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, googleslackstatus.Add)
}
