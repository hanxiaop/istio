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

	"istio.io/istio/pkg/config/resource"

	. "github.com/onsi/gomega"
)

func TestMessage_String(t *testing.T) {
	g := NewGomegaWithT(t)
	mt := NewMessageType(Error, "IST-0042", "Cheese type not found: %q")
	m := NewMessage(mt, nil, "Feta")

	g.Expect(m.String()).To(Equal(`Error [IST-0042] Cheese type not found: "Feta"`))
}

func TestMessageWithResource_String(t *testing.T) {
	g := NewGomegaWithT(t)
	mt := NewMessageType(Error, "IST-0042", "Cheese type not found: %q")
	m := NewMessage(mt, &resource.Instance{Origin: testOrigin{name: "toppings/cheese", ref: testReference{"path/to/file"}}}, "Feta")

	g.Expect(m.String()).To(Equal(`Error [IST-0042] (toppings/cheese path/to/file) Cheese type not found: "Feta"`))
}

func TestMessage_Unstructured(t *testing.T) {
	g := NewGomegaWithT(t)
	mt := NewMessageType(Error, "IST-0042", "Cheese type not found: %q")
	m := NewMessage(mt, nil, "Feta")

	g.Expect(m.Unstructured(true)).To(Not(HaveKey("origin")))
	g.Expect(m.Unstructured(false)).To(Not(HaveKey("origin")))

	m = NewMessage(mt, &resource.Instance{Origin: testOrigin{name: "toppings/cheese"}}, "Feta")

	g.Expect(m.Unstructured(true)).To((HaveKey("origin")))
	g.Expect(m.Unstructured(false)).To(Not(HaveKey("origin")))
}

func TestMessageWithDocRef(t *testing.T) {
	g := NewGomegaWithT(t)
	mt := NewMessageType(Error, "IST-0042", "Cheese type not found: %q")
	m := NewMessage(mt, nil, "Feta")
	m.DocRef = "test-ref"
	g.Expect(m.Unstructured(false)["documentation_url"]).To(Equal("https://istio.io/docs/reference/config/analysis/IST-0042?ref=test-ref"))
}

func TestMessage_JSON(t *testing.T) {
	g := NewGomegaWithT(t)
	mt := NewMessageType(Error, "IST-0042", "Cheese type not found: %q")
	m := NewMessage(mt, &resource.Instance{Origin: testOrigin{name: "toppings/cheese", ref: testReference{"path/to/file"}}}, "Feta")

	j, _ := json.Marshal(&m)
	g.Expect(string(j)).To(Equal(`{"code":"IST-0042","documentation_url":"https://istio.io/docs/reference/config/analysis/IST-0042"` +
		`,"level":"Error","message":"Cheese type not found: \"Feta\"","origin":"toppings/cheese","reference":"path/to/file"}`))
}

func TestMessage_UpdateLine(t *testing.T) {
	g := NewGomegaWithT(t)
	mt := NewMessageType(Error, "IST-0042", "Cheese type not found: %q")
	m := NewMessage(mt, nil, "Feta")
	m.DocRef = "test-ref"
	m.line = []int{0}
	m.UpdateLine(5)
	g.Expect(m.line[0]).To(Equal(5))
}

func TestMessage_FindErrorWord(t *testing.T) {
	g := NewGomegaWithT(t)
	mt := NewMessageType(Error, "IST-0042", "Cheese type not found: %q")
	m := NewMessage(mt, &resource.Instance{Origin: testOrigin{name: "toppings/cheese", ref: testReference{"path/to/file"}}}, "Feta")
	rMap := make(map[string]*resource.Instance)
	sMap := make(map[string]map[string]map[string]string)
	m.FindErrorWord(rMap, sMap)
	key := m.Resource.Origin.Reference().String(false) + m.Type.code + "Feta"
	g.Expect(key).To(Equal("path/to/fileIST-0042Feta"))
	_, ok := rMap[key]
	g.Expect(ok).To(Equal(false))
}

func TestFindErrorWordHelper(t *testing.T) {
	g := NewGomegaWithT(t)
	mt := NewMessageType(Error, "IST-0042", "Cheese type not found: %q")
	m := NewMessage(mt, &resource.Instance{Origin: testOrigin{name: "toppings/cheese", ref: testReference{"path/to/file"}}}, "Feta")
	key, value := FindErrorWordHelper(mt.code, []string{"Feta"}, m.Resource)
	g.Expect(key).To(Equal(""))
	g.Expect(value).To(Equal(""))
}
