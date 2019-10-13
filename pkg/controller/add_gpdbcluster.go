package controller

import (
	"github.com/soxueren/greenplum-operator/pkg/controller/gpdbmaster"
	"github.com/soxueren/greenplum-operator/pkg/controller/gpdbsegment"
	"github.com/soxueren/greenplum-operator/pkg/controller/gpdbmirror"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, gpdbmaster.Add)
	AddToManagerFuncs = append(AddToManagerFuncs, gpdbsegment.Add)
	AddToManagerFuncs = append(AddToManagerFuncs, gpdbmirror.Add)	
}
