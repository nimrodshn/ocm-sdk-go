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

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

// IdentityProviderBuilder contains the data and logic needed to build 'identity_provider' objects.
//
// Representation of an identity provider.
type IdentityProviderBuilder struct {
	id            *string
	href          *string
	link          bool
	type_         *IdentityProviderType
	name          *string
	challenge     *bool
	login         *bool
	mappingMethod *IdentityProviderMappingMethod
	github        *GithubIdentityProviderBuilder
	gitlab        *GitlabIdentityProviderBuilder
	google        *GoogleIdentityProviderBuilder
	ldap          *LdapidentityProviderBuilder
	openID        *OpenIdidentityProviderBuilder
}

// NewIdentityProvider creates a new builder of 'identity_provider' objects.
func NewIdentityProvider() *IdentityProviderBuilder {
	return new(IdentityProviderBuilder)
}

// ID sets the identifier of the object.
func (b *IdentityProviderBuilder) ID(value string) *IdentityProviderBuilder {
	b.id = &value
	return b
}

// HREF sets the link to the object.
func (b *IdentityProviderBuilder) HREF(value string) *IdentityProviderBuilder {
	b.href = &value
	return b
}

// Link sets the flag that indicates if this is a link.
func (b *IdentityProviderBuilder) Link(value bool) *IdentityProviderBuilder {
	b.link = value
	return b
}

// Type sets the value of the 'type' attribute
// to the given value.
//
// Type of identity provider.
func (b *IdentityProviderBuilder) Type(value IdentityProviderType) *IdentityProviderBuilder {
	b.type_ = &value
	return b
}

// Name sets the value of the 'name' attribute
// to the given value.
//
//
func (b *IdentityProviderBuilder) Name(value string) *IdentityProviderBuilder {
	b.name = &value
	return b
}

// Challenge sets the value of the 'challenge' attribute
// to the given value.
//
//
func (b *IdentityProviderBuilder) Challenge(value bool) *IdentityProviderBuilder {
	b.challenge = &value
	return b
}

// Login sets the value of the 'login' attribute
// to the given value.
//
//
func (b *IdentityProviderBuilder) Login(value bool) *IdentityProviderBuilder {
	b.login = &value
	return b
}

// MappingMethod sets the value of the 'mapping_method' attribute
// to the given value.
//
// Controls how mappings are established between provider identities and user objects.
func (b *IdentityProviderBuilder) MappingMethod(value IdentityProviderMappingMethod) *IdentityProviderBuilder {
	b.mappingMethod = &value
	return b
}

// Github sets the value of the 'github' attribute
// to the given value.
//
// Details for `github` identity providers.
func (b *IdentityProviderBuilder) Github(value *GithubIdentityProviderBuilder) *IdentityProviderBuilder {
	b.github = value
	return b
}

// Gitlab sets the value of the 'gitlab' attribute
// to the given value.
//
// Details for `gitlab` identity providers.
func (b *IdentityProviderBuilder) Gitlab(value *GitlabIdentityProviderBuilder) *IdentityProviderBuilder {
	b.gitlab = value
	return b
}

// Google sets the value of the 'google' attribute
// to the given value.
//
// Details for `google` identity providers.
func (b *IdentityProviderBuilder) Google(value *GoogleIdentityProviderBuilder) *IdentityProviderBuilder {
	b.google = value
	return b
}

// LDAP sets the value of the 'LDAP' attribute
// to the given value.
//
// Details for `ldap` identity providers.
func (b *IdentityProviderBuilder) LDAP(value *LdapidentityProviderBuilder) *IdentityProviderBuilder {
	b.ldap = value
	return b
}

// OpenID sets the value of the 'open_ID' attribute
// to the given value.
//
// Details for `openid` identity providers.
func (b *IdentityProviderBuilder) OpenID(value *OpenIdidentityProviderBuilder) *IdentityProviderBuilder {
	b.openID = value
	return b
}

// Build creates a 'identity_provider' object using the configuration stored in the builder.
func (b *IdentityProviderBuilder) Build() (object *IdentityProvider, err error) {
	object = new(IdentityProvider)
	object.id = b.id
	object.href = b.href
	object.link = b.link
	if b.type_ != nil {
		object.type_ = b.type_
	}
	if b.name != nil {
		object.name = b.name
	}
	if b.challenge != nil {
		object.challenge = b.challenge
	}
	if b.login != nil {
		object.login = b.login
	}
	if b.mappingMethod != nil {
		object.mappingMethod = b.mappingMethod
	}
	if b.github != nil {
		object.github, err = b.github.Build()
		if err != nil {
			return
		}
	}
	if b.gitlab != nil {
		object.gitlab, err = b.gitlab.Build()
		if err != nil {
			return
		}
	}
	if b.google != nil {
		object.google, err = b.google.Build()
		if err != nil {
			return
		}
	}
	if b.ldap != nil {
		object.ldap, err = b.ldap.Build()
		if err != nil {
			return
		}
	}
	if b.openID != nil {
		object.openID, err = b.openID.Build()
		if err != nil {
			return
		}
	}
	return
}
