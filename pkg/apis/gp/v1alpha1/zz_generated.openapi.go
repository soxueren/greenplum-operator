// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/soxueren/greenplum-operator/pkg/apis/gp/v1alpha1.GPDBCluster":       schema_pkg_apis_gp_v1alpha1_GPDBCluster(ref),
		"github.com/soxueren/greenplum-operator/pkg/apis/gp/v1alpha1.GPDBClusterSpec":   schema_pkg_apis_gp_v1alpha1_GPDBClusterSpec(ref),
		"github.com/soxueren/greenplum-operator/pkg/apis/gp/v1alpha1.GPDBClusterStatus": schema_pkg_apis_gp_v1alpha1_GPDBClusterStatus(ref),
	}
}

func schema_pkg_apis_gp_v1alpha1_GPDBCluster(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GPDBCluster is the Schema for the gpdbclusters API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/soxueren/greenplum-operator/pkg/apis/gp/v1alpha1.GPDBClusterSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/soxueren/greenplum-operator/pkg/apis/gp/v1alpha1.GPDBClusterStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/soxueren/greenplum-operator/pkg/apis/gp/v1alpha1.GPDBClusterSpec", "github.com/soxueren/greenplum-operator/pkg/apis/gp/v1alpha1.GPDBClusterStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_gp_v1alpha1_GPDBClusterSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GPDBClusterSpec defines the desired state of GPDBCluster",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_gp_v1alpha1_GPDBClusterStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GPDBClusterStatus defines the observed state of GPDBCluster",
				Type:        []string{"object"},
			},
		},
	}
}
