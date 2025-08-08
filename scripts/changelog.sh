#!/bin/bash
# Copyright © 2022 sealos.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
TAG=$1

if [ -z "${TAG}" ]; then
    echo "Usage: $0 <repo> <tag>"
    exit 1
fi

if [ ! -d CHANGELOG ]; then
    mkdir CHANGELOG
fi

git-chglog --output CHANGELOG/CHANGELOG-"${TAG#v}".md --next-tag ${TAG} ${TAG}

echo "# Changelog" > CHANGELOG/CHANGELOG.md
echo -e "\nAll notable changes to this project will be documented in this file.\n" >> CHANGELOG/CHANGELOG.md

for file in $(ls CHANGELOG |grep -v '^CHANGELOG.md$' | sort -V -r); do
    version=$(echo $file | sed -E 's/CHANGELOG-(.*)\.md/\1/')
    echo -e "- [CHANGELOG-${version}.md](./CHANGELOG-${version}.md)" >> CHANGELOG/CHANGELOG.md
done
