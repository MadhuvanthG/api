#!/usr/bin/env bash

function index_js() {
cat <<EOF
const parent = require("../");

module.exports = parent.${1};
EOF
}

function index_test_ts() {
cat <<EOF
describe("${1}", () => {
    test("require", () => {
        const schema = require("./index");

        Object.keys(schema).forEach((key) => {
            console.log(key)
        })
    });
});
EOF
}

## TODO: EVENTUALLY REMOVE THESE

function file_d_ts() {
cat <<EOF
export * from "./index"
EOF
}

function file_js() {
cat <<EOF | tee "${directory}/${file_base}.js"
module.exports = require("./index");
EOF
}

function root_js() {
  root_pkg="${1}"
  shift

  filenames=""
  for file in "$@"; do
    filenames="${filenames}"$'\n'"filenames.push(path.join(__dirname, ${file//\//\", \"}));"
  done

cat <<EOF
const path = require('path');
const protoLoader = require('@grpc/proto-loader');
const grpc = require('grpc');

const filenames = [];
${filenames}

const packageDefinition = protoLoader.loadSync(
    filenames,
    {
        keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true,
        includeDirs: [
            __dirname,
        ]
    }
);

const descriptor = grpc.loadPackageDefinition(packageDefinition);

module.exports = descriptor.${root_pkg};
EOF
}

## MAIN EXECUTION

for file in $(find . -name *.proto | cut -c 3-); do
  directory="$(dirname "${file}")"
  base="$(basename "${directory}")"
  file_base="$(basename "${file}" .proto)"

  if [[ ! -f "${directory}/index.d.ts" ]]; then
    touch "${directory}/index.d.ts"
  fi

  index_js "${base}"      > "${directory}/index.js"
  index_test_ts "${base}" > "${directory}/index_test.ts"
  file_d_ts               > "${directory}/${file_base}.d.ts"
  file_js                 > "${directory}/${file_base}.js"
done

root_js "cloud.deps.api" $(find . -name *.proto | cut -c 3- | xargs -I {} echo '"{}"') > index.js
index_test_ts "api" > index_test.ts
