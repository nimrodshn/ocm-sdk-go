/*
Copyright (c) 2019 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1

// ResourceQuotaBuilder contains the data and logic needed to build 'resource_quota' objects.
//
//
type ResourceQuotaBuilder struct {
	id                   *string
	href                 *string
	link                 bool
	organizationID       *string
	sku                  *string
	resourceName         *string
	resourceType         *string
	byoc                 *bool
	availabilityZoneType *string
	allowed              *int
	reserved             *int
}

// NewResourceQuota creates a new builder of 'resource_quota' objects.
func NewResourceQuota() *ResourceQuotaBuilder {
	return new(ResourceQuotaBuilder)
}

// ID sets the identifier of the object.
func (b *ResourceQuotaBuilder) ID(value string) *ResourceQuotaBuilder {
	b.id = &value
	return b
}

// HREF sets the link to the object.
func (b *ResourceQuotaBuilder) HREF(value string) *ResourceQuotaBuilder {
	b.href = &value
	return b
}

// Link sets the flag that indicates if this is a link.
func (b *ResourceQuotaBuilder) Link(value bool) *ResourceQuotaBuilder {
	b.link = value
	return b
}

// OrganizationID sets the value of the 'organization_ID' attribute
// to the given value.
//
//
func (b *ResourceQuotaBuilder) OrganizationID(value string) *ResourceQuotaBuilder {
	b.organizationID = &value
	return b
}

// SKU sets the value of the 'SKU' attribute
// to the given value.
//
//
func (b *ResourceQuotaBuilder) SKU(value string) *ResourceQuotaBuilder {
	b.sku = &value
	return b
}

// ResourceName sets the value of the 'resource_name' attribute
// to the given value.
//
//
func (b *ResourceQuotaBuilder) ResourceName(value string) *ResourceQuotaBuilder {
	b.resourceName = &value
	return b
}

// ResourceType sets the value of the 'resource_type' attribute
// to the given value.
//
//
func (b *ResourceQuotaBuilder) ResourceType(value string) *ResourceQuotaBuilder {
	b.resourceType = &value
	return b
}

// BYOC sets the value of the 'BYOC' attribute
// to the given value.
//
//
func (b *ResourceQuotaBuilder) BYOC(value bool) *ResourceQuotaBuilder {
	b.byoc = &value
	return b
}

// AvailabilityZoneType sets the value of the 'availability_zone_type' attribute
// to the given value.
//
//
func (b *ResourceQuotaBuilder) AvailabilityZoneType(value string) *ResourceQuotaBuilder {
	b.availabilityZoneType = &value
	return b
}

// Allowed sets the value of the 'allowed' attribute
// to the given value.
//
//
func (b *ResourceQuotaBuilder) Allowed(value int) *ResourceQuotaBuilder {
	b.allowed = &value
	return b
}

// Reserved sets the value of the 'reserved' attribute
// to the given value.
//
//
func (b *ResourceQuotaBuilder) Reserved(value int) *ResourceQuotaBuilder {
	b.reserved = &value
	return b
}

// Build creates a 'resource_quota' object using the configuration stored in the builder.
func (b *ResourceQuotaBuilder) Build() (object *ResourceQuota, err error) {
	object = new(ResourceQuota)
	object.id = b.id
	object.href = b.href
	object.link = b.link
	if b.organizationID != nil {
		object.organizationID = b.organizationID
	}
	if b.sku != nil {
		object.sku = b.sku
	}
	if b.resourceName != nil {
		object.resourceName = b.resourceName
	}
	if b.resourceType != nil {
		object.resourceType = b.resourceType
	}
	if b.byoc != nil {
		object.byoc = b.byoc
	}
	if b.availabilityZoneType != nil {
		object.availabilityZoneType = b.availabilityZoneType
	}
	if b.allowed != nil {
		object.allowed = b.allowed
	}
	if b.reserved != nil {
		object.reserved = b.reserved
	}
	return
}
