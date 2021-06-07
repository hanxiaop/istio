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
	"testing"

	. "github.com/onsi/gomega"

	"istio.io/istio/pkg/config/resource"
	"istio.io/istio/pkg/url"
)

func TestMessage_String(t *testing.T) {
	g := NewWithT(t)
	mt := NewMessageBase(Error, "CheeseNotFoundError", "IST-0042", "")
	m := NewMessage(mt, "Cheese type is not found", "Cheese type not found: %q", nil, nil, "Feta")

	g.Expect(m.String()).To(Equal(`Error [IST-0042] Cheese type not found: "Feta"`))
}

func TestMessageWithResource_String(t *testing.T) {
	g := NewWithT(t)
	mt := NewMessageBase(Error, "NotFoundError", "IST-0042", "")
	m := NewMessage(mt, "", "Cheese type not found: %q",
		&resource.Instance{Origin: testOrigin{name: "toppings/cheese", ref: testReference{"path/to/file"}}},
		nil, "Feta")

	g.Expect(m.String()).To(Equal(`Error [IST-0042] (toppings/cheese path/to/file) Cheese type not found: "Feta"`))
}

func TestMessage_Unstructured(t *testing.T) {
	g := NewWithT(t)
	mt := NewMessageBase(Error, "NotFoundError", "IST-0042", "")
	m := NewMessage(mt, "", "Cheese type not found: %q", nil, nil, "Feta")

	g.Expect(m.Unstructured(true)).To(Not(HaveKey("origin")))
	g.Expect(m.Unstructured(false)).To(Not(HaveKey("origin")))

	m = NewMessage(mt, "", "Cheese type not found: %q",
		 &resource.Instance{Origin: testOrigin{name: "toppings/cheese"}}, nil, "Feta")

	g.Expect(m.Unstructured(true)).To((HaveKey("origin")))
	g.Expect(m.Unstructured(false)).To(Not(HaveKey("origin")))
}

func TestMessageWithDocRef(t *testing.T) {
	g := NewWithT(t)
	mt := NewMessageBase(Error, "NotFoundError", "IST-0042", "")
	m := NewMessage(mt, "", "Cheese type not found: %q", nil, nil, "Feta")
	m.Schema.MessageBase.DocumentationUrl = "test-ref"
	g.Expect(m.Unstructured(false)["documentationUrl"]).To(Equal(url.ConfigAnalysis + "/ist0042/?ref=test-ref"))
}

func TestMessage_JSON(t *testing.T) {
	g := NewWithT(t)
	mt := NewMessageBase(Error, "NotFoundError", "IST-0042", "")
	m := NewMessage(mt, "",  "Cheese type not found: %q",
		&resource.Instance{Origin: testOrigin{name: "toppings/cheese", ref: testReference{"path/to/file"}}},
		nil, "Feta")

	j, _ := json.Marshal(&m)
	g.Expect(string(j)).To(Equal(`{"code":"IST0042","documentationUrl":"` + url.ConfigAnalysis + `/ist0042/"` +
		`,"level":"Error","message":"Cheese type not found: \"Feta\"","origin":"toppings/cheese","reference":"path/to/file"}`))
}

func TestMessage_ReplaceLine(t *testing.T) {
	testCases := []string{"test.yaml", "test.yaml:1", "test.yaml:10", "test.yaml: 10", "test", "test:10", "123:10", "123"}
	result := make([]string, 0)
	g := NewGomegaWithT(t)
	m := &Message{Line: 321}
	for _, v := range testCases {
		result = append(result, m.ReplaceLine(v))
	}
	g.Expect(result).To(Equal([]string{"test.yaml", "test.yaml:321", "test.yaml:321", "test.yaml:321", "test", "test:321", "123:321", "123"}))
}
