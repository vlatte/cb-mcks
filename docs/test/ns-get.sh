#!/bin/bash
# ------------------------------------------------------------------------------
# usage
if [ "$#" -lt 1 ]; then 
	echo "./ns-get.sh <namespace>"
	echo "./ns-get.sh cb-mcks-ns "
	exit 0
fi

source ./conf.env

# ------------------------------------------------------------------------------
# const

# ------------------------------------------------------------------------------
# variables

v_NAMESPACE="$1"
if [ "${v_NAMESPACE}" == "" ]; then read -e -p "namespace ? : "  v_NAMESPACE;	fi
if [ "${v_NAMESPACE}" == "" ]; then echo "[ERROR] missing <namespace>"; exit -1; fi


# ------------------------------------------------------------------------------
# print info.
echo ""
echo "[INFO]"
echo "- (Name of namespace)        is '${v_NAMESPACE}'"


# ------------------------------------------------------------------------------
# get
get() {
	curl -sX GET ${c_URL_TUMBLEBUG}/ns/${v_NAMESPACE}          -H "${c_AUTH}" -H "${c_CT}" | jq
}

# ------------------------------------------------------------------------------
if [ "$1" != "-h" ]; then 
	echo ""
	echo "------------------------------------------------------------------------------"
	get;
fi
