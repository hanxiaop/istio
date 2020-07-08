// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package diag

import (
	"encoding/json"
	"fmt"
	"strings"

	"istio.io/istio/pkg/config/resource"
)

// DocPrefix is the root URL for validation message docs
const DocPrefix = "https://istio.io/docs/reference/config/analysis"

// MessageType is a type of diagnostic message
type MessageType struct {
	// The level of the message.
	level Level

	// The error code of the message
	code string

	// TODO: Make this localizable
	template string
}

// Level returns the level of the MessageType
func (m *MessageType) Level() Level { return m.level }

// Code returns the code of the MessageType
func (m *MessageType) Code() string { return m.code }

// Template returns the message template used by the MessageType
func (m *MessageType) Template() string { return m.template }

// Message is a specific diagnostic message
// TODO: Implement using Analysis message API
type Message struct {
	Type *MessageType

	// The Parameters to the message
	Parameters []interface{}

	// Resource is the underlying resource instance associated with the
	// message, or nil if no resource is associated with it.
	Resource *resource.Instance

	// DocRef is an optional reference tracker for the documentation URL
	DocRef string

	// Exact line number of the error place
	line []int
}

// Unstructured returns this message as a JSON-style unstructured map
func (m *Message) Unstructured(includeOrigin bool) map[string]interface{} {
	result := make(map[string]interface{})

	result["code"] = m.Type.Code()
	result["level"] = m.Type.Level().String()
	if includeOrigin && m.Resource != nil {
		result["origin"] = m.Resource.Origin.FriendlyName()
		if m.Resource.Origin.Reference() != nil {
			var loc string
			if len(m.line) != 0 && m.line[0] != 0 {
				loc = " " + m.Resource.Origin.Reference().String(true) + fmt.Sprintf("%d", m.line[0])
			} else {
				loc = " " + m.Resource.Origin.Reference().String(false)
			}
			result["reference"] = loc
		}
	}
	result["message"] = fmt.Sprintf(m.Type.Template(), m.Parameters...)

	docQueryString := ""
	if m.DocRef != "" {
		docQueryString = fmt.Sprintf("?ref=%s", m.DocRef)
	}
	result["documentation_url"] = fmt.Sprintf("%s/%s%s", DocPrefix, m.Type.Code(), docQueryString)

	return result
}

// String implements io.Stringer
func (m *Message) String() string {
	origin := ""
	if m.Resource != nil {
		loc := ""
		if m.Resource.Origin.Reference() != nil {
			if len(m.line) != 0 && m.line[0] != 0 {
				loc = " " + m.Resource.Origin.Reference().String(true) + fmt.Sprintf("%d", m.line[0])
			} else {
				loc = " " + m.Resource.Origin.Reference().String(false)
			}
		}
		origin = " (" + m.Resource.Origin.FriendlyName() + loc + ")"
	}
	return fmt.Sprintf(
		"%v [%v]%s %s", m.Type.Level(), m.Type.Code(), origin, fmt.Sprintf(m.Type.Template(), m.Parameters...))
}

// MarshalJSON satisfies the Marshaler interface
func (m *Message) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Unstructured(true))
}

// Update line number in the message structure
func (m *Message) UpdateLine(line int) {
	m.line[0] = line
}

// Build maps that contains the error lines and corresponding resource
func (m *Message) FindErrorWord(rMap map[string]*resource.Instance, sMap map[string]map[string]map[string]string) {
	// Skip resource with no reference
	if m.Resource.Origin.Reference() == nil {
		return
	}
	resourceName := m.Resource.Origin.Reference().String(false)
	if _, ok := rMap[resourceName]; !ok {
		mAddress := m.Resource
		rMap[resourceName] = mAddress
	}
	var para []string
	mKey := m.Resource.Origin.Reference().String(false) + m.Type.Code()
	for _, p := range m.Parameters {
		para = append(para, fmt.Sprintf("%v", p))
		mKey += fmt.Sprintf("%v", p)
	}
	resKey, resValue := FindErrorWordHelper(m.Type.code, para, m.Resource)
	if mKey != "" && resKey != "" && resValue != "" {
		if sMap[resourceName] == nil {
			sMap[resourceName] = make(map[string]map[string]string)
		}
		if sMap[resourceName][resKey] == nil {
			sMap[resourceName][resKey] = make(map[string]string)
		}
		sMap[resourceName][resKey][resValue] = mKey
	}
}

func (m *Message) LineNumber() []int {
	return m.line
}

// NewMessageType returns a new MessageType instance.
func NewMessageType(level Level, code, template string) *MessageType {
	return &MessageType{
		level:    level,
		code:     code,
		template: template,
	}
}

// NewMessage returns a new Message instance from an existing type.
func NewMessage(mt *MessageType, r *resource.Instance, p ...interface{}) Message {
	return Message{
		Type:       mt,
		Resource:   r,
		Parameters: p,
		line:       []int{0},
	}
}

// Decide search words to use depending on error message types
// Some of the message do not have reference line number, so no search word also
// Some of the message do not have error parameters, starting line of the yaml chunk will be shown as default
func FindErrorWordHelper(eCode string, para []string, r *resource.Instance) (string, string) {
	var resKey, resValue string
	switch eCode {
	case "IST0101":
		resKey = para[0]
		resValue = para[1]
		if resKey == "gateway" {
			resKey = "gateways"
		} else if resKey == "selector" {
			equalInd := strings.Index(resValue, "=")
			resKey = resValue[:equalInd]
			resValue = resValue[equalInd+1:]
		} else if resKey == "host+subset in destinationrule" {
			resKey = "subset"
			resValue = resValue[strings.Index(resValue, "+")+1:]
		} else if resKey == "host:port" {
			resKey = "port"
			resValue = resValue[strings.Index(resValue, ":")+1:]
		} else {
			resKey = para[0]
		}
	case "IST0104":
		equalInd := strings.Index(para[0], "=")
		resKey = para[0][:equalInd]
		resValue = para[0][equalInd+1:]
	case "IST0110":
		resKey = "name"
		friendlyName := r.Origin.FriendlyName()
		friendlyName = friendlyName[strings.Index(friendlyName, " ")+1:]
		dotIndex := strings.Index(friendlyName, ".")
		if dotIndex > 0 {
			friendlyName = friendlyName[:dotIndex]
		}
		resValue = friendlyName
	case "IST0111":
		resKey = "name"
		friendlyName := r.Origin.FriendlyName()
		friendlyName = friendlyName[strings.Index(friendlyName, " ")+1:]
		dotIndex := strings.Index(friendlyName, ".")
		if dotIndex > 0 {
			friendlyName = friendlyName[:dotIndex]
		}
		resValue = friendlyName
	case "IST0112":
		resKey = "host"
		resValue = para[0]

	case "IST0122":
		resKey = para[0]
		resValue = para[1]
		if resKey == "corsPolicy.allowOrigins" {
			resKey = "allowOrigins"
		}
	default:
		resKey = ""
		resValue = ""
	}
	return resKey, resValue
}
