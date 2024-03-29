# Copyright 2017 Sergey Berezin

# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at

#     http://www.apache.org/licenses/LICENSE-2.0

# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

INSTALLS=. ./jsonread
GOPATH=$(shell go env GOPATH)

all:
	@echo "Please pick a target:"
	@echo "  make install  - compile and install gocovcheck executables"
	@echo "  make init     - initialize the development environment"
	@echo "  make test     - run tests"
	@echo "  make gofmt     - format all *.go files"
	@echo "  make goconvey  - start a goconvey session (Crtl-C to exit)"
	@echo "  make clean    - delete generated files"

install:
	go install $(INSTALLS)

test:
	./runtests

init:
	go install github.com/smartystreets/goconvey@v1.7.2
	go install honnef.co/go/tools/cmd/staticcheck@2021.1.2
	go install $(INSTALLS)
	@echo "Bootstrap done!"

gofmt:
	gofmt -s -w .

goconvey:
	$(GOPATH)/bin/goconvey

clean:
	rm -f ".coverage"
	rm -f "coverage.html"
