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
	"testing"

	. "github.com/onsi/gomega"

	"istio.io/istio/pkg/url"
)

func TestMessages_Sort(t *testing.T) {
	g := NewWithT(t)

	firstMsg := NewMessage(
		NewMessageBase(Error, "", "B1", ""),
		"",
		"Template: %q",
		MockResource("B"),
		nil,
		"B",
	)
	secondMsg := NewMessage(
		NewMessageBase(Warning, "", "A1", ""),
		"",
		"Template: %q",
		MockResource("B"),
		nil,
		"B",
	)
	thirdMsg := NewMessage(
		NewMessageBase(Warning, "", "B1", ""),
		"",
		"Template: %q",
		MockResource("A"),
		nil,
		"B",
	)
	fourthMsg := NewMessage(
		NewMessageBase(Warning, "", "B1", ""),
		"",
		"Template: %q",
		MockResource("B"),
		nil,
		"A",
	)
	fifthMsg := NewMessage(
		NewMessageBase(Warning, "", "B1", ""),
		"",
		"Template: %q",
		MockResource("B"),
		nil,
		"B",
	)

	msgs := Messages{fifthMsg, fourthMsg, thirdMsg, secondMsg, firstMsg}
	expectedMsgs := Messages{firstMsg, secondMsg, thirdMsg, fourthMsg, fifthMsg}

	msgs.Sort()

	g.Expect(msgs).To(Equal(expectedMsgs))
}

func TestMessages_SortWithNilOrigin(t *testing.T) {
	g := NewWithT(t)

	firstMsg := NewMessage(
		NewMessageBase(Warning, "", "B1", ""),
		"",
		"Template: %q",
		nil,
		nil,
		"B",
	)
	secondMsg := NewMessage(
		NewMessageBase(Warning, "", "B1", ""),
		"",
		"Template: %q",
		nil,
		nil,
		"C",
	)
	thirdMsg := NewMessage(
		NewMessageBase(Warning, "", "B1", ""),
		"",
		"Template: %q",
		MockResource("B"),
		nil,
		"B",
	)

	msgs := Messages{thirdMsg, secondMsg, firstMsg}
	expectedMsgs := Messages{firstMsg, secondMsg, thirdMsg}

	msgs.Sort()

	g.Expect(msgs).To(Equal(expectedMsgs))
}

func TestMessages_SortedCopy(t *testing.T) {
	g := NewWithT(t)

	firstMsg := NewMessage(
		NewMessageBase(Warning, "", "B1", ""),
		"",
		"Template: %q",
		MockResource("B"),
		nil,
		"B",
	)
	secondMsg := NewMessage(
		NewMessageBase(Warning, "", "A1", ""),
		"",
		"Template: %q",
		MockResource("B"),
		nil,
		"B",
	)
	// Oops, we have a duplicate (identical to firstMsg) - it should be removed.
	thirdMsg := NewMessage(
		NewMessageBase(Warning, "", "B1", ""),
		"",
		"Template: %q",
		MockResource("B"),
		nil,
		"B",
	)

	msgs := Messages{thirdMsg, secondMsg, firstMsg}
	expectedMsgs := Messages{firstMsg, secondMsg}

	newMsgs := msgs.SortedDedupedCopy()

	g.Expect(newMsgs).To(Equal(expectedMsgs))
}

func TestMessages_SetRefDoc(t *testing.T) {
	g := NewWithT(t)

	firstMsg := NewMessage(
		NewMessageBase(Warning, "", "B1", ""),
		"",
		"Template: %q",
		MockResource("B"),
		nil,
		"B",
	)
	secondMsg := NewMessage(
		NewMessageBase(Warning, "", "C1", ""),
		"",
		"Template: %q",
		MockResource("B"),
		nil,
		"B",
	)

	msgs := Messages{firstMsg, secondMsg}
	msgs.SetDocRef("istioctl-awesome")

	getDocURL := func(msg Message) string {
		return msg.Unstructured(false)["documentationUrl"].(string)
	}

	g.Expect(getDocURL(msgs[0])).To(Equal(url.ConfigAnalysis + "/b1/?ref=istioctl-awesome"))
	g.Expect(getDocURL(msgs[1])).To(Equal(url.ConfigAnalysis + "/c1/?ref=istioctl-awesome"))
}

func TestMessages_Filter(t *testing.T) {
	g := NewWithT(t)

	firstMsg := NewMessage(
		NewMessageBase(Warning, "", "B1", ""),
		"",
		"Template: %q",
		MockResource("B"),
		nil,
		"B",
	)
	secondMsg := NewMessage(
		NewMessageBase(Warning, "", "A1", ""),
		"",
		"Template: %q",
		MockResource("B"),
		nil,
		"B",
	)
	thirdMsg := NewMessage(
		NewMessageBase(Warning, "", "C1", ""),
		"",
		"Template: %q",
		MockResource("B"),
		nil,
		"B",
	)

	msgs := Messages{firstMsg, secondMsg, thirdMsg}
	filteredMsgs := msgs.FilterOutLowerThan(Warning)
	expectedMsgs := Messages{firstMsg, thirdMsg}

	g.Expect(filteredMsgs).To(Equal(expectedMsgs))
}

func TestMessages_FilterOutAll(t *testing.T) {
	g := NewWithT(t)

	firstMsg := NewMessage(
		NewMessageBase(Warning, "", "A1", ""),
		"",
		"Template: %q",
		MockResource("B"),
		nil,
		"B",
	)
	secondMsg := NewMessage(
		NewMessageBase(Warning, "", "C1", ""),
		"",
		"Template: %q",
		MockResource("B"),
		nil,
		"B",
	)

	msgs := Messages{firstMsg, secondMsg}
	filteredMsgs := msgs.FilterOutLowerThan(Error)
	expectedMsgs := Messages{}

	g.Expect(filteredMsgs).To(Equal(expectedMsgs))
}
