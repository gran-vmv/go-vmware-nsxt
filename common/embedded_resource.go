/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package common

type EmbeddedResource struct {

	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`

	Schema string `json:"_schema,omitempty"`

	Self *SelfResourceLink `json:"_self,omitempty"`

	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected.
	Revision int64 `json:"_revision"`

	Owner *OwnerResourceLink `json:"_owner,omitempty"`

	// Description of this resource
	Description string `json:"description,omitempty"`

	// Defaults to ID if not set
	DisplayName string `json:"display_name,omitempty"`

	// Identifier of the resource
	Id string `json:"id,omitempty"`

	// The type of this resource.
	ResourceType string `json:"resource_type,omitempty"`
}
