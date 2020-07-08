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

package rt

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func TestPositionString(t *testing.T) {
	testcases := []struct {
		filename string
		line     int
		output   string
	}{
		{
			filename: "test.yaml",
			line:     1,
			output:   "test.yaml:1",
		},
		{
			filename: "test.yaml",
			line:     0,
			output:   "test.yaml",
		},
		{
			filename: "test.json",
			line:     1,
			output:   "test.json",
		},
		{
			filename: "-",
			line:     1,
			output:   "-:1",
		},
		{
			filename: "",
			line:     1,
			output:   "",
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			g := NewGomegaWithT(t)

			p := Position{Filename: tc.filename, Line: tc.line}
			g.Expect(p.String(false)).To(Equal(tc.output))
		})
	}
}

func TestCheckComment(t *testing.T) {
	yaml := []string{"\n", "# line \n", "abc\n"}
	g := NewGomegaWithT(t)
	g.Expect(CheckComment(0, yaml)).To(Equal(2))
}

func TestRemoveQuotation(t *testing.T) {
	s := "\"test\""
	g := NewGomegaWithT(t)
	g.Expect(RemoveQuotation(s)).To(Equal("test"))
}

func TestMapHelper(t *testing.T) {
	yaml := []string{"test1: content1", "test2:", "  test3: content3"}
	g := NewGomegaWithT(t)
	c1 := map[string]int{"content1": 1}
	c3 := map[string]int{"content3": 3}
	smallMap := map[string]interface{}{"test3": c3}
	resultMap := map[string]interface{}{"test1": c1, "test2": smallMap}
	resMap, index := MapHelper(yaml, 0, 0, 1)
	g.Expect(resMap).To(Equal(resultMap))
	g.Expect(index).To(Equal(3))
}

func TestFindErrorsHelper(t *testing.T) {
	map1 := map[string]string{"test": "result"}
	para := map[string]int{"test": 1}
	r1, r2 := FindErrorsHelper(map1, para)
	g := NewGomegaWithT(t)
	g.Expect(r1).To(Equal([]string{"result"}))
	g.Expect(r2).To(Equal([]int{1}))
}

func TestPosition_FindErrors(t *testing.T) {
	p := Position{Filename: "test", Line: 1}
	c1 := map[string]int{"content1": 1}
	c3 := map[string]int{"content3": 3}
	smallMap := map[string]interface{}{"test3": c3}
	resultMap := map[string]interface{}{"test1": c1, "test2": smallMap}
	p.ChunkMap = resultMap

	inputNestMap := map[string]string{"content3": "testResult"}
	outerMap := map[string]map[string]string{"test2": inputNestMap}

	messageRes, lineRes := p.FindErrors(outerMap)
	g := NewGomegaWithT(t)
	g.Expect(messageRes).To(Equal([]string{"testResult"}))
	g.Expect(lineRes).To(Equal([]int{3}))
}
