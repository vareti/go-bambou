// Copyright (c) 2015, Alcatel-Lucent Inc.
//
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
//   list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
//   this list of conditions and the following disclaimer in the documentation
//   and/or other materials provided with the distribution.
//
// * Neither the name of bambou nor the names of its
//   contributors may be used to endorse or promote products derived from
//   this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package bambou

import (
	"fmt"
)

// Represents a list of Exposables.
type ExposablesList []Exposable

// Interface of a Exposable object.
//
// An Exposable also implements the Identifiable and Operationable interfaces.
type Exposable interface {
	Identifiable
	Operationable

	GetURL() string
	GetURLForChildrenIdentity(RESTIdentity) string
}

// Represents a list of Rootables
type RootablesList []Rootable

// Interface of a Rootable object.
//
// An Rootable also implements the Identifiable and Exposable. Rootable
// is the interface an object must implement in order to be able to act
// as a root api object. For instance for "/auth".
type Rootable interface {
	Exposable

	GetAPIKey() string
	SetAPIKey(string)
}

// Represents an object that is exposed throught the ReST api.
//
// This struct must be embedded into all objects that are available
// throught the ReST api.
type ExposedObject struct {
	ID           string       `json:"ID,omitempty"`
	ParentID     string       `json:"parentID,omitempty"`
	ParentType   string       `json:"parentType,omitempty"`
	Owner        string       `json:"owner,omitempty"`
	ParentObject string       `json:"-"`
	Identity     RESTIdentity `json:"-"`
}

// Returns the Identity
func (o *ExposedObject) GetIdentity() RESTIdentity {

	return o.Identity
}

// Sets the Identity
func (o *ExposedObject) SetIdentity(identity RESTIdentity) {

	o.Identity = identity
}

// Returns the URL that holds the information about the object.
func (o *ExposedObject) GetURL() string {

	session := CurrentSession()
	baseURL := session.URL

	if o.ID != "" {
		return baseURL + "/" + o.Identity.ResourceName + "/" + o.ID
	}

	return baseURL + "/" + o.Identity.ResourceName
}

// Returns the URL of children with the given identity
//
// The URL will be constructed based on the current object URL and the identity of
// the children.
func (o *ExposedObject) GetURLForChildrenIdentity(identity RESTIdentity) string {

	return o.GetURL() + "/" + identity.ResourceName
}

// Returns the string representation of the object
func (o *ExposedObject) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity.RESTName, o.ID)
}
