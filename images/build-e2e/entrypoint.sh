#!/bin/sh

# Vars
BYNARY=e2e.test
if [[ ${PLATOFORM} == 'windows' ]]; then
    BYNARY=e2e.test.exe
fi
BINARY_PATH="/opt/crc/bin/${PLATOFORM}-amd64/${CRC_E2E_BYNARY}"
FEATURES_PATH=/opt/crc/features
RESULTS_PATH="${RESULTS_PATH:-/output}"

if [ "${DEBUG:-}" = "true" ]; then
    set -xuo 
fi

# Validate conf
validate=true
[[ -z "${TARGET_HOST}" ]] \
    && echo "TARGET_HOST requried" \
    && validate=false

[[ -z "${TARGET_HOST_USERNAME}" ]] \
    && echo "TARGET_HOST_USERNAME requried" \
    && validate=false

[[ -z "${TARGET_HOST_KEY_PATH}" && -z "${TARGET_HOST_PASSWORD}" ]] \
    && echo "TARGET_HOST_KEY_PATH or TARGET_HOST_PASSWORD required" \
    && validate=false
[[ $validate == false ]] && exit 1

# Set SCP / SSH command with pass or key
NO_STRICT='-o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null'
[[ ! -z "${TARGET_HOST_KEY_PATH}" ]] \
    && SCP="scp -r ${NO_STRICT} -i ${TARGET_HOST_KEY_PATH}" \
        || SCP="sshpass -p ${TARGET_HOST_PASSWORD} scp -r ${NO_STRICT}" \
    && SSH="ssh ${NO_STRICT} -i ${TARGET_HOST_KEY_PATH}" \
        || SSH="sshpass -p ${TARGET_HOST_PASSWORD} ssh ${NO_STRICT}"

# Create execution folder 
EXECUTION_FOLDER="/Users/${TARGET_HOST_USERNAME}/crc-e2e/${RANDOM}"
if [[ ${PLATOFORM} == 'linux' ]]; then
    EXECUTION_FOLDER="/home/${TARGET_HOST_USERNAME}/crc-e2e/${RANDOM}"
fi
DATA_FOLDER="${EXECUTION_FOLDER}/out"
if [[ ${PLATOFORM} == 'windows' ]]; then
    # Todo change for powershell cmdlet
    $SSH "${TARGET_HOST_USERNAME}@${TARGET_HOST}" "mkdir -p ${EXECUTION_FOLDER} 
else
    $SSH "${TARGET_HOST_USERNAME}@${TARGET_HOST}" "mkdir -p ${EXECUTION_FOLDER} 
fi

# Copy crc-e2e binary and features spec to target host
$SCP "${BINARY_PATH}" "${TARGET_HOST_USERNAME}@${TARGET_HOST}:${EXECUTION_FOLDER}"
$SCP "${FEATURES_PATH}" "${TARGET_HOST_USERNAME}@${TARGET_HOST}:${EXECUTION_FOLDER}"

# Run (one shot) monictl
# TODO add run parameters
# BINARY_EXEC="${EXECUTION_FOLDER}/${BYNARY} 
# $SSH "${TARGET_HOST_USERNAME}@${TARGET_HOST}" "${BINARY_EXEC}"

# # Get results
# mkdir -p "${RESULTS_PATH}"
# $SCP "${TARGET_HOST_USERNAME}@${TARGET_HOST}:${DATA_FOLDER}" "${RESULTS_PATH}"

# # Remove remote execution fodler
# $SSH "${TARGET_HOST_USERNAME}@${TARGET_HOST}" rm -rf "${EXECUTION_FOLDER}"