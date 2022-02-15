// Copyright 2020 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

// Code generated by ./scripts/makeOpenApiInfoDotGo.sh; DO NOT EDIT.

package kubernetesapi

import (
	"sigs.k8s.io/kustomize/kyaml/openapi/kubernetesapi/v1212"
)

const Info = "{title:Kubernetes,version:v1.21.2}"

var OpenAPIMustAsset = map[string]func(string) []byte{
	"v1212": v1212.MustAsset,
}

const DefaultOpenAPI = "v1212"
