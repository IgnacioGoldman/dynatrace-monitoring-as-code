// @license
// Copyright 2021 Dynatrace LLC
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build unit

package coordinate

import (
	"gotest.tools/assert"
	"testing"
)

func TestMatch(t *testing.T) {
	coordinate := Coordinate{
		Project: "project1",
		Api:     "dashboard",
		Config:  "dashboard1",
	}

	result := coordinate.Match(Coordinate{
		Project: "project1",
		Api:     "dashboard",
		Config:  "dashboard1",
	})

	assert.Assert(t, result, "should match")
}

func TestMatchShouldReturnFalseOnNonMatching(t *testing.T) {
	coordinate := Coordinate{
		Project: "project1",
		Api:     "dashboard",
		Config:  "dashboard1",
	}

	result := coordinate.Match(Coordinate{
		Project: "project1",
		Api:     "auto-tag",
		Config:  "tags",
	})

	assert.Assert(t, !result, "shouldn't match")
}

func TestMatchShouldReturnFalseOnNonMatchingSameApi(t *testing.T) {
	coordinate := Coordinate{
		Project: "project1",
		Api:     "dashboard",
		Config:  "dashboard1",
	}

	result := coordinate.Match(Coordinate{
		Project: "project1",
		Api:     "dashboard",
		Config:  "dashboard2",
	})

	assert.Assert(t, !result, "shouldn't match")
}