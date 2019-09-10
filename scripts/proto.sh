#!/bin/bash

set -e

# Replace the generated .pb.go files in the repository with the ones generated
# by the build. The generated versions are committed for IDE usage.

wd="$(cd "$( dirname "${BASH_SOURCE[0]}")" >/dev/null && pwd )"
cd $wd/..

# Build all the Go protobuf libraries.
echo "Building Go protobuf libraries"

bazel build $(bazel query 'kind("go_proto_library", //... except //vendor/...)') || code=$?

if [ "$code" != "" ]
then
    echo "If Bazel failed with the following error try 'brew uninstall protobuf' before rerunning this script"
    echo "ERROR: /private/var/tmp/{user}/{build_id}/external/com_google_protobuf/BUILD:120:1: C++ compilation of rule '@com_google_protobuf//:protobuf' failed (Exit 1) cc_wrapper.sh failed: error executing command external/local_config_cc/cc_wrapper.sh -U_FORTIFY_SOURCE -fstack-protector -Wall -Wthread-safety -Wself-assign -fcolor-diagnostics -fno-omit-frame-pointer -g0 -O2 '-D_FORTIFY_SOURCE=1' -DNDEBUG ... (remaining 37 argument(s) skipped)"
    exit $code
fi

# Hacky way to attempt copying the results for the platform we just built for.
filter='linux_amd64'
if [[ "$(uname)" != "Darwin" ]]; then
    filter='darwin_amd64'
fi

# Copy them from the build output into the repo. They're architecture
# neutral so we don't care what architecture they're under.
echo "Copying protobuf files:"
export base=$(basename $(git rev-parse --show-toplevel))
find bazel-out/ -type f -name "*.pb*.go" \
    | grep -v $filter \
    | grep $base \
    | perl -ne 'chomp; $src=$_; s|.*$ENV{base}/||; print "\t$_\n"; `cp $src $_ && chmod 644 $_`'
