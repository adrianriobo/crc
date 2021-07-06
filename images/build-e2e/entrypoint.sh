#!/bin/sh

# Vars
BINARY=e2e.test
if [[ ${PLATFORM} == 'windows' ]]; then
    BINARY=e2e.test.exe
fi
BINARY_PATH="/opt/crc/bin/${PLATFORM}-amd64/${BINARY}"
# Review this when go 1.16 with embed support
FEATURES_PATH=/opt/crc/features
TESTDATA_PATH=/opt/crc/testdata
APPLESCRIPTS_PATH=/opt/crc/applescripts
# Results
RESULTS_PATH="${RESULTS_PATH:-/output}"

if [ "${DEBUG:-}" = "true" ]; then
    set -xuo 
fi

# Validate conf
validate=true
[[ -z "${TARGET_HOST+x}" ]] \
    && echo "TARGET_HOST required" \
    && validate=false

[[ -z "${TARGET_HOST_USERNAME+x}" ]] \
    && echo "TARGET_HOST_USERNAME required" \
    && validate=false

[[ -z "${TARGET_HOST_KEY_PATH+x}" && -z "${TARGET_HOST_PASSWORD+x}" ]] \
    && echo "TARGET_HOST_KEY_PATH or TARGET_HOST_PASSWORD required" \
    && validate=false
[[ $validate == false ]] && exit 1

[[ -z "${PULL_SECRET_FILE_PATH+x}" ]] \
    && echo "PULL_SECRET_FILE_PATH required" \
    && validate=false

# Define remote connection
REMOTE="${TARGET_HOST_USERNAME}@${TARGET_HOST}"
if [[ ! -z "${TARGET_HOST_DOMAIN+x}" ]]; then
    REMOTE="${TARGET_HOST_USERNAME}@${TARGET_HOST_DOMAIN}@${TARGET_HOST}"
fi

# Set SCP / SSH command with pass or key
NO_STRICT='-o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null'
if [[ ! -z "${TARGET_HOST_KEY_PATH+x}" ]]; then
    SCP="scp -r ${NO_STRICT} -i ${TARGET_HOST_KEY_PATH}"
    SSH="ssh ${NO_STRICT} -i ${TARGET_HOST_KEY_PATH}"
else
    SCP="sshpass -p ${TARGET_HOST_PASSWORD} scp -r ${NO_STRICT}" \
    SSH="sshpass -p ${TARGET_HOST_PASSWORD} ssh ${NO_STRICT}"
fi

# Create execution folder 
EXECUTION_FOLDER="/Users/${TARGET_HOST_USERNAME}/crc-e2e/${RANDOM}"
if [[ ${PLATFORM} == 'linux' ]]; then
    EXECUTION_FOLDER="/home/${TARGET_HOST_USERNAME}/crc-e2e/${RANDOM}"
fi
DATA_FOLDER="${EXECUTION_FOLDER}/out"
if [[ ${PLATFORM} == 'windows' ]]; then
    # Todo change for powershell cmdlet
    $SSH "${REMOTE}" "powershell.exe -c New-Item -ItemType directory -Path ${EXECUTION_FOLDER}"
else
    $SSH "${REMOTE}" "mkdir -p ${EXECUTION_FOLDER}"
fi

# Copy crc-e2e binary and pull-secret
$SCP "${BINARY_PATH}" "${REMOTE}:${EXECUTION_FOLDER}"
$SCP "${PULL_SECRET_FILE_PATH}" "${REMOTE}:${EXECUTION_FOLDER}"
# Review this when go 1.16 with embed support
$SCP "${FEATURES_PATH}" "${REMOTE}:${EXECUTION_FOLDER}"
$SCP "${TESTDATA_PATH}" "${REMOTE}:${EXECUTION_FOLDER}"
$SCP "${APPLESCRIPTS_PATH}" "${REMOTE}:${EXECUTION_FOLDER}"

# Run e2e
# TODO add run parameters
# BINARY_EXEC="${EXECUTION_FOLDER}/${BYNARY} 
# $SSH "${TARGET_HOST_USERNAME}@${TARGET_HOST}" "${BINARY_EXEC}"

# # Get results
# mkdir -p "${RESULTS_PATH}"
# $SCP "${TARGET_HOST_USERNAME}@${TARGET_HOST}:${DATA_FOLDER}" "${RESULTS_PATH}"

# # Remove remote execution fodler
# $SSH "${TARGET_HOST_USERNAME}@${TARGET_HOST}" rm -rf "${EXECUTION_FOLDER}"