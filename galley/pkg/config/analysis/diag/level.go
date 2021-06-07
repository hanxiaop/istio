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
	"strings"

	"istio.io/api/analysis/v1alpha1"
)

var (
	// Unknown level is for unknown messages
	Unknown = v1alpha1.AnalysisMessageBase_UNKNOWN

	// Info level is for informational messages
	Info = v1alpha1.AnalysisMessageBase_INFO

	// Warning level is for warning messages
	Warning = v1alpha1.AnalysisMessageBase_WARNING

	// Error level is for error messages
	Error = v1alpha1.AnalysisMessageBase_ERROR
)

// GetAllLevels returns an arbitrarily ordered slice of all Levels defined.
func GetAllLevels() []v1alpha1.AnalysisMessageBase_Level {
	return []v1alpha1.AnalysisMessageBase_Level{Info, Warning, Error, Unknown}
}

// GetAllLevelStrings returns a list of strings representing the names of all Levels defined. The order is arbitrary but
// should be the same as GetAllLevels.
func GetAllLevelStrings() []string {
	levels := GetAllLevels()
	var s []string
	for _, l := range levels {
		s = append(s, l.String())
	}
	return s
}

// GetUppercaseStringToLevelMap returns a mapping of uppercase strings to Level structs. This function is intended to be
// used to convert user input to structs.
func GetUppercaseStringToLevelMap() map[string]v1alpha1.AnalysisMessageBase_Level {
	m := make(map[string]v1alpha1.AnalysisMessageBase_Level)
	for _, l := range GetAllLevels() {
		m[strings.ToUpper(l.String())] = l
	}
	return m
}
