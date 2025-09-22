#!/usr/bin/env bash
set -euo pipefail

TARGET=${1:-windows}
OUTPUT_DIR=${2:-dist}
KEEP_IMAGE=${KEEP_IMAGE:-0}

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "${SCRIPT_DIR}/../.." && pwd)"

case "$TARGET" in
  windows)
    DOCKERFILE="${SCRIPT_DIR}/Dockerfile.windows"
    IMAGE_TAG="ai-launcher:windows"
    ARTIFACT_SRC="/ai-launcher.exe"
    ARTIFACT_NAME="ai-launcher.exe"
    ;;
  linux)
    DOCKERFILE="${SCRIPT_DIR}/Dockerfile"
    IMAGE_TAG="ai-launcher:linux"
    ARTIFACT_SRC="/app/ai-launcher"
    ARTIFACT_NAME="ai-launcher"
    ;;
  *)
    echo "Unsupported target: ${TARGET}" >&2
    exit 1
    ;;
esac

mkdir -p "${OUTPUT_DIR}"

echo "[+] Building Docker image '${IMAGE_TAG}' using ${DOCKERFILE}"
docker build -f "${DOCKERFILE}" -t "${IMAGE_TAG}" "${REPO_ROOT}"

CONTAINER_ID=$(docker create "${IMAGE_TAG}")
cleanup() {
  docker rm -f "${CONTAINER_ID}" >/dev/null 2>&1 || true
}
trap cleanup EXIT

echo "[+] Copying artifact to ${OUTPUT_DIR}/${ARTIFACT_NAME}"
docker cp "${CONTAINER_ID}:${ARTIFACT_SRC}" "${OUTPUT_DIR}/${ARTIFACT_NAME}"

if [[ "${KEEP_IMAGE}" != "1" ]]; then
  echo "[+] Removing intermediate image ${IMAGE_TAG}"
  docker rmi "${IMAGE_TAG}" >/dev/null 2>&1 || true
fi

echo "[?] Build complete. Artifact available at ${OUTPUT_DIR}/${ARTIFACT_NAME}"
