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
	"container/list"
	"fmt"
	"path/filepath"
	"strings"

	"istio.io/istio/pkg/config/resource"
	"istio.io/istio/pkg/config/schema/collection"
	"istio.io/istio/pkg/config/schema/collections"
)

// Origin is a K8s specific implementation of resource.Origin
type Origin struct {
	Collection collection.Name
	Kind       string
	FullName   resource.FullName
	Version    resource.Version
	Ref        resource.Reference
}

var _ resource.Origin = &Origin{}
var _ resource.Reference = &Position{}

// FriendlyName implements resource.Origin
func (o *Origin) FriendlyName() string {
	parts := strings.Split(o.FullName.String(), "/")
	if len(parts) == 2 {
		// The istioctl convention is <type> <name>[.<namespace>].
		// This code has no notion of a default and always shows the namespace.
		return fmt.Sprintf("%s %s.%s", o.Kind, parts[1], parts[0])
	}
	return fmt.Sprintf("%s %s", o.Kind, o.FullName.String())
}

// Namespace implements resource.Origin
func (o *Origin) Namespace() resource.Namespace {
	// Special case: the namespace of a namespace resource is its own name
	if o.Collection == collections.K8SCoreV1Namespaces.Name() {
		return resource.Namespace(o.FullName.Name)
	}

	return o.FullName.Namespace
}

// Reference implements resource.Origin
func (o *Origin) Reference() resource.Reference {
	return o.Ref
}

// Position is a representation of the location of a source.
type Position struct {
	Filename string // filename, if any
	Line     int    // line number, starting at 1
	ChunkMap    map[string]interface{} // map with key-value pairs from yaml chunk
}

// String outputs the string representation of the position.
func (p *Position) String(hasLineNumber bool) string {
	s := p.Filename
	// TODO: support json file position.
	if p.isValid() && filepath.Ext(p.Filename) != ".json" {
		if s != "" {
			s += ":"
		}
		if hasLineNumber == false {
			s += fmt.Sprintf("%d", p.Line)
		}
	}
	return s
}

func (p *Position) isValid() bool {
	return p.Line > 0 && p.Filename != ""
}

func (p *Position) ProcessMap(yamlChunk []byte) {
	yamlInfo :=	string(yamlChunk)
	yamlInfoArray := strings.Split(yamlInfo, "\n")
	keyMap, _ := MapHelper(yamlInfoArray, 0, 0, p.Line)
	p.ChunkMap = keyMap
}

func (p *Position) YamlMap() map[string]interface{} {
	return p.ChunkMap
}

func (p *Position) FindErrors(sMap map[string]map[string]string) ([]string, []int) {
	queue := list.New()
	var messageRes []string
	var lineRes []int
	if p.ChunkMap == nil {
		return messageRes, lineRes
	}

	// check duplicate results
	hasResult := make(map[string]bool)

	queue.PushBack(p.ChunkMap)
	for queue.Len() > 0 {
		front := queue.Front()
		frontValue := front.Value
		if fmt.Sprintf("%T", frontValue) == "map[string]interface {}" {
			temp := frontValue.(map[string]interface{})
			for k := range  temp{
				if v, ok := sMap[k]; ok {
					mRes, lRes := FindErrorsHelper(v, temp[k])
					if len(mRes) > 0 {
						for i, v := range mRes {
							if _, ok := hasResult[v]; ok {
								continue
							}
							hasResult[v] = true
							messageRes = append(messageRes, mRes[i])
							lineRes = append(lineRes, lRes[i])
						}
					}
				}
				if fmt.Sprintf("%T", temp[k]) == "map[string]interface {}" {
					queue.PushBack(temp[k])
				}else if fmt.Sprintf("%T", temp[k]) == "[]map[string]interface {}" {
					newVal := temp[k].([]map[string]interface {})
					for _, v := range newVal {
						queue.PushBack(v)
					}
				}
			}
		}

		queue.Remove(front)
	}
	return messageRes, lineRes
}

func FindErrorsHelper(sMap map[string]string, para interface{}) ([]string, []int) {
	var mRes []string
	var lRes []int
	if fmt.Sprintf("%T", para) == "map[string]int" {
		resMap := para.(map[string]int)
		for k, v := range resMap {
			if sv, ok := sMap[k]; ok {
				mRes = append(mRes, sv)
				lRes = append(lRes, v)
			}
		}
	}else if fmt.Sprintf("%T", para) == "map[string]interface {}" {
		resMap := para.(map[string]interface{})
		for _, v := range resMap {
			mV, lV := FindErrorsHelper(sMap, v)
			if len(mV) > 0 {
				mRes = append(mRes, mV...)
				lRes = append(lRes, lV...)
			}
		}
	}else if fmt.Sprintf("%T", para) == "[]map[string]interface {}" {
		resMap := para.([]map[string]interface {})
		for _, v := range resMap {
			mV, lV := FindErrorsHelper(sMap, v)
			if len(mV) > 0 {
				mRes = append(mRes, mV...)
				lRes = append(lRes, lV...)
			}
		}
	} else if fmt.Sprintf("%T", para) == "[]map[string]int" {
		resMap := para.([]map[string]int)
		for _, v := range resMap {
			mV, lV := FindErrorsHelper(sMap, v)
			if len(mV) > 0 {
				mRes = append(mRes, mV...)
				lRes = append(lRes, lV...)
			}
		}
	}
	return mRes, lRes
}

func MapHelper(yamlChunk []string, index int, space int, startLine int) (map[string]interface{}, int) {
	res := make(map[string]interface{})
	nextInd := index
	for nextInd < len(yamlChunk) {
		nextInd = CheckComment(nextInd, yamlChunk)

		if nextInd >= len(yamlChunk) {
			return res, nextInd
		}

		curStr := yamlChunk[nextInd]
		numSpace, _, newString := FindSpaceNum(curStr)
		if numSpace < space {
			return res, nextInd
		}

		colonInd := strings.Index(newString, ":")

		if colonInd != len(newString) - 1 {
			pair := strings.Split(newString, ":")
			pair[0], pair[1] = RemoveQuotation(pair[0]), RemoveQuotation(pair[1])
			if val, ok := res[pair[0]]; ok {
				if fmt.Sprintf("%T", res[pair[0]]) == "map[string]int" {
					firstMap := val
					res[pair[0]] = []map[string]int{}
					res[pair[0]] = append(res[pair[0]].([]map[string]int), firstMap.(map[string]int))
				}
				res[pair[0]] = append(res[pair[0]].([]map[string]int), map[string]int{pair[1]:nextInd + startLine})
			}else {
				res[pair[0]] = make(map[string]int)
				res[pair[0]].(map[string]int)[pair[1]] = nextInd + startLine
			}
			nextInd += 1
		}else {
			key := RemoveQuotation(newString[:len(newString) - 1])
			nextInd += 1

			if nextInd >= len(yamlChunk) {
				return res, nextInd
			}

			nextInd = CheckComment(nextInd, yamlChunk)

			if nextInd >= len(yamlChunk) {
				return res, nextInd
			}

			curStr := yamlChunk[nextInd]
			numSpace, _, newString = FindSpaceNum(curStr)
			newString = RemoveQuotation(newString)
			if strings.Index(newString, ":") < 0 {
				res[key] = []map[string]int{}
				for strings.Index(newString, ":") < 0 {
					tMap := map[string]int{newString:nextInd + startLine}
					res[key] = append(res[key].([]map[string]int), tMap)
					nextInd += 1

					if nextInd >= len(yamlChunk) {
						return res, nextInd
					}

					nextInd = CheckComment(nextInd, yamlChunk)

					if nextInd >= len(yamlChunk) {
						return res, nextInd
					}

					curStr := yamlChunk[nextInd]
					numSpace, _, newString = FindSpaceNum(curStr)
				}
			}else {
				newRes, newInd := MapHelper(yamlChunk, nextInd, numSpace, startLine)
				if val, ok := res[key]; ok {
					if fmt.Sprintf("%T", res[key]) == "map[string]interface {}" {
						firstMap := val
						res[key] = []map[string]interface{}{}
						res[key] = append(res[key].([]map[string]interface{}), firstMap.(map[string]interface{}))
					}
					res[key] = append(res[key].([]map[string]interface{}), newRes)
				}else {
					res[key] = newRes
				}
				res[key] = newRes
				nextInd = newInd
			}
		}

	}
	return res, nextInd

}

// return space number in front, flag to show if it is an array, and new trimed string
func FindSpaceNum(s string)  (int, bool, string) {
	newS := strings.TrimSpace(s)
	isArray := false
	if newS[0] == '-' {
		isArray = true
		newS = newS[2:]
	}
	commentInd := strings.Index(newS, "#")
	if commentInd > 0{
		newS = newS[:commentInd]
	}
	numSpace := strings.Index(s, newS[:1])
	return numSpace, isArray, newS
}

func RemoveQuotation(s string) string {
	newS := strings.TrimSpace(s)
	if len(newS) > 0 && newS[0] == '"' {
		newS = newS[1:]
	}
	if len(newS) > 0 && newS[len(newS) - 1] == '"' {
		newS = newS[:len(newS) - 1]
	}
	return newS
}

func CheckComment(curInd int, yamlChunk []string) int {
	curStr := yamlChunk[curInd]
	curI := curInd
	tempInd := CheckCommentHelper(curStr, curI)
	for tempInd < len(yamlChunk) && tempInd != curI {
		curI = tempInd
		curStr = yamlChunk[curI]
		tempInd = CheckCommentHelper(curStr, curI)
	}
	return tempInd
}

func CheckCommentHelper(s string, ind int) int  {
	newS := strings.TrimSpace(s)
	if len(newS) == 0 || newS[0] == '#' {
		return ind + 1
	}
	return ind
}
