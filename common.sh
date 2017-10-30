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

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "$DIR"

PARENT_PATH="src/github.com/sergey-a-berezin/gocovcheck"
# ${DIR} is expected to be ${ROOT}/src/github.com/sergey-a-berezin/gocovcheck

ROOT=${DIR}
curr_parent="${PARENT_PATH}"

while [ ${curr_parent} != "." ]; do
    curr_parent="$( dirname "${curr_parent}")"
    ROOT="$( dirname "${ROOT}" )"
done

if [ "${ROOT}/${PARENT_PATH}" != "${DIR}" ]; then
    echo "  ERROR: your checkout is not inside $PARENT_PATH."
    exit 1
fi
