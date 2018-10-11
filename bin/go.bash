#!/bin/bash

set -e

if [ -z $GOOGLEAPIS ]; then
	echo >&2 "need GOOGLEAPIS variable pointed to a copy of master branch of https://github.com/googleapis/googleapis"
	exit 1
fi

if [ -z $COMMON_PROTO ]; then
	echo >&2 "need COMMON_PROTO variable pointed to a copy of https://github.com/googleapis/api-common-protos"
	exit 1
fi

if ! which gen-go-sample > /dev/null; then
	cat >&2 <<EOF
need gen-go-sample binary:
download https://github.com/googleapis/gapic-generator-go
$ cd gapic-generator-go
$ go install ./cmd/gen-go-sample
EOF
	exit 1
fi

# Paging methods don't work yet
#
# gen-go-sample \
#   -clientpkg 'cloud.google.com/go/dialogflow/apiv2;dialogflow' \
#   -gapic gapic_configs/dialogflow/v2/dialogflow_gapic.yaml \
#   -o dialogflow/v2/go \
#   -desc <(protoc -o /dev/stdout --include_imports -I "$COMMON_PROTO" -I "$GOOGLEAPIS" "$GOOGLEAPIS"/google/cloud/dialogflow/v2/*.proto)


# Arrays don't work yet
#
# gen-go-sample \
#   -clientpkg 'cloud.google.com/go/dlp/apiv2;dlp' \
#   -gapic gapic_configs/dlp/v2/dlp_gapic.yaml \
#   -o dlp/v2/go \
#   -desc <(protoc -o /dev/stdout --include_imports -I "$COMMON_PROTO" -I "$GOOGLEAPIS" "$GOOGLEAPIS"/google/privacy/dlp/v2/*.proto)

# Bug: nested message types don't work
#
# gen-go-sample \
#   -clientpkg 'cloud.google.com/go/kms/apiv1;kms' \
#   -gapic gapic_configs/kms/v1/cloudkms_gapic.yaml \
#   -o kms/v1/go \
#   -desc <(protoc -o /dev/stdout --include_imports -I "$COMMON_PROTO" -I "$GOOGLEAPIS" "$GOOGLEAPIS"/google/cloud/kms/v1/*.proto)

gen-go-sample \
  -clientpkg 'cloud.google.com/go/language/apiv1;language' \
  -gapic gapic_configs/language/v1/language_gapic.yaml \
  -o language/v1/go \
  -desc <(protoc -o /dev/stdout --include_imports -I "$COMMON_PROTO" -I "$GOOGLEAPIS" "$GOOGLEAPIS"/google/cloud/language/v1/*.proto)

gen-go-sample \
  -clientpkg 'cloud.google.com/go/speech/apiv1;speech' \
  -gapic gapic_configs/speech/v1/cloud_speech_gapic.yaml \
  -o speech/v1/go \
  -desc <(protoc -o /dev/stdout --include_imports -I "$COMMON_PROTO" -I "$GOOGLEAPIS" "$GOOGLEAPIS"/google/cloud/speech/v1/*.proto)

cat >&2 <<'EOF'
done generating; to test compile:
find -name '*.go' | xargs -P $(nproc) -IF go build -o /dev/null F
EOF
