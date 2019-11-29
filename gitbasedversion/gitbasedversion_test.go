// Copyright 2019 Sergey Berezin

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// jsonread tool takes a file and a key, reads the file as a flat JSON format,
// and prints out the value of the key. If key is not present, prints the default value.

package main

import (
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {
	t.Parallel()
	Convey("Version works", t, func() {
		r := strings.NewReader(`37b9613
3c00ed6
a53dd34
2d33be8
d76538d
3e618f8
1cd90a2
c7056bf
4755409
`)
		So(Version(r), ShouldEqual, "00009-37b9613")
	})
}
