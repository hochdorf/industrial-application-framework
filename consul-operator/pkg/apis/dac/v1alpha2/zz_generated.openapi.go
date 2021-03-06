// Copyright 2020 Nokia
// Licensed under the BSD 3-Clause License.
// SPDX-License-Identifier: BSD-3-Clause

// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha2

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./pkg/apis/dac/v1alpha2.Consul":       schema_pkg_apis_dac_v1alpha2_Consul(ref),
		"./pkg/apis/dac/v1alpha2.ConsulSpec":   schema_pkg_apis_dac_v1alpha2_ConsulSpec(ref),
		"./pkg/apis/dac/v1alpha2.ConsulStatus": schema_pkg_apis_dac_v1alpha2_ConsulStatus(ref),
	}
}

func schema_pkg_apis_dac_v1alpha2_Consul(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Consul is the Schema for the consuls API",
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
							Ref: ref("./pkg/apis/dac/v1alpha2.ConsulSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/dac/v1alpha2.ConsulStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/dac/v1alpha2.ConsulSpec", "./pkg/apis/dac/v1alpha2.ConsulStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_dac_v1alpha2_ConsulSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ConsulSpec defines the desired state of Consul",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"replicaCount": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: RunCrTemplater \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"ports": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/dac/v1alpha2.Ports"),
						},
					},
				},
				Required: []string{"replicaCount", "ports"},
			},
		},
		Dependencies: []string{
			"./pkg/apis/dac/v1alpha2.Ports"},
	}
}

func schema_pkg_apis_dac_v1alpha2_ConsulStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ConsulStatus defines the observed state of Consul",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"prevSpec": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: RunCrTemplater \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
							Ref:         ref("./pkg/apis/dac/v1alpha2.ConsulSpec"),
						},
					},
					"appStatus": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"appReportedData": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/dac/v1alpha2.AppReporteData"),
						},
					},
					"appliedResources": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/nokia/industrial-application-framework/consul-operator/pkg/k8sdynamic.ResourceDescriptor"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/dac/v1alpha2.AppReporteData", "./pkg/apis/dac/v1alpha2.ConsulSpec", "github.com/nokia/industrial-application-framework/consul-operator/pkg/k8sdynamic.ResourceDescriptor"},
	}
}
