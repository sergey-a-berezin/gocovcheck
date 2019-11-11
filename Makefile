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

ROOT := $(shell source ./common.sh; echo "$${ROOT}")

all:
	@echo "Please pick a target:"
	@echo "  make build    - compile (but do not install) the package"
	@echo "  make install  - compile and install gocovcheck executable"
	@echo "  make init     - initialize the development environment"
	@echo "  make test     - run tests"
	@echo "  make clean    - delete object files and other temporary files"
	@echo "  make pristine - clean + delete everything created by bootstrap"

install:
	(source "$(ROOT)/bin/activate"; cd "$(ROOT)/src/gocovcheck"; go install)

build:
	(source bin/activate; cd src/gocovcheck; go build)

test:
	./runtests

init:
	./bootstrap
	(source "$(ROOT)/bin/activate"; cd "$(ROOT)/src/gocovcheck"; \
		go install github.com/smartystreets/goconvey; \
		go install golang.org/x/lint/golint; \
		go install; \
		cd jsonread; go install)
	@echo "Bootstrap done!"

clean:
	chmod -R u+w "$(ROOT)/pkg"
	rm -rf "$(ROOT)/pkg" "$(ROOT)/src/.coverage"

pristine: clean
	rm -rf "$(ROOT)/bin"
