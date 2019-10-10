package controller

import (
	"github.com/soxueren/greenplum-operator/pkg/controller/dbnodeservice"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, dbnodeservice.Add)
}
