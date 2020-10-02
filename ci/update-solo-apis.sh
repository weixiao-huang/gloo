#!/bin/bash

cd ../
git clone https://github.com/solo-io/solo-apis.git
cd solo-apis
./hack/sync-gloo-apis.sh;make generate -B
git checkout -b update-solo-apis
git add .
git commit -m "update to latest gloo version"
git push origin master

