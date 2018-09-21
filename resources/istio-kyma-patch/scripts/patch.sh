#!/usr/bin/env bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"

if [[ -z ${CONFIG_DIR} ]]; then
    CONFIG_DIR=${DIR}
fi

function run_all_patches() {
  echo "--> Patch resources"
  for f in $(find ${CONFIG_DIR} -name '*\.patch\.json'); do
    local type=$(basename ${f} | cut -d. -f2)
    local name=$(basename ${f} | cut -d. -f1)
    echo "    Patch $type $name from: ${f}"
    local patch=$(cat ${f})
    kubectl patch ${type} -n istio-system ${name} --patch "$patch" --type json
  done
}

function remove_not_used() {
  echo "--> Delete resources"
  while read line; do
    echo "    Delete $line"
    local type=$(cut -d' ' -f1 <<< ${line})
    local name=$(cut -d' ' -f2 <<< ${line})
    kubectl delete ${type} ${name} -n istio-system
  done <${CONFIG_DIR}/delete
}

function configure_sidecar_injector() {
  echo "--> Configure sidecar injector"
  local configmap=$(kubectl -n istio-system get configmap istio-sidecar-injector -o jsonpath='{.data.config}')
  local alreadyEnabled=$(grep "policy: enabled" <<< "$configmap")
  if [[ -n ${alreadyEnabled} ]]; then
    # Disable automatic injecting
    configmap=$(sed 's/policy: enabled/policy: disabled/' <<< "$configmap")

    # Set limits for sidecar. Our namespaces have resource quota set thus every container needs to have limits defined.
    # Add limits to already existing resources sections
    configmap=$(sed 's|    resources:|    resources:\'$'\n      limits: { memory: 50Mi }|' <<< "$configmap")
    # In case there is no limits section add one at the beginning of container definition. It serves as default.
    configmap=$(sed 's|  - name: istio-\(.*\)|  - name: istio-\1\'$'\n    resources: { limits: { memory: 50Mi } }|' <<< "$configmap")

    # Escape new lines and double quotes for kubectl
    configmap=$(sed -e ':a' -e 'N' -e '$!ba' -e 's/\n/\\n/g' <<< "$configmap")
    configmap=$(sed 's/"/\\"/g' <<< "$configmap")

    kubectl patch -n istio-system configmap istio-sidecar-injector --type merge -p '{"data": {"config":"'"$configmap"'"}}'
  fi
}

function apply_all() {
  echo "--> Apply resources"
  for f in $(find ${DIR} -name '*\.yaml' | xargs -I '{}' basename '{}'); do
    echo "    Apply $f"
    kubectl apply -f ${CONFIG_DIR}/${f}
  done
}

function check_requirements() {
  while read crd; do
    echo "    Require CRD crd"
    kubectl get crd ${crd}
    if [ $? -ne 0 ]; then
        echo "Cannot find required CRD $crd"
    fi
  done <${CONFIG_DIR}/required-crds
}

configure_sidecar_injector
check_requirements
run_all_patches
remove_not_used
apply_all
