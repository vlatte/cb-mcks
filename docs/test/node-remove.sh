#!/bin/bash
# -----------------------------------------------------------------
# usage
if [ "$#" -lt 1 ]; then 
	echo "./node-remove.sh <namespace> <clsuter name> <node name>"
	echo "./node-remove.sh cb-mcks-ns cluster-01 cluster-01-w-1-asdflk"
	exit 0; 
fi

source ./conf.env

# ------------------------------------------------------------------------------
# const


# -----------------------------------------------------------------
# parameter

# 1. namespace
if [ "$#" -gt 0 ]; then v_NAMESPACE="$1"; else	v_NAMESPACE="${NAMESPACE}"; fi
if [ "${v_NAMESPACE}" == "" ]; then 
	read -e -p "Namespace ? : " v_NAMESPACE
fi
if [ "${v_NAMESPACE}" == "" ]; then echo "[ERROR] missing <namespace>"; exit -1; fi

# 2. Cluster Name
if [ "$#" -gt 1 ]; then v_CLUSTER_NAME="$2"; else	v_CLUSTER_NAME="${CLUSTER_NAME}"; fi
if [ "${v_CLUSTER_NAME}" == "" ]; then 
	read -e -p "Cluster name  ? : "  v_CLUSTER_NAME
fi
if [ "${v_CLUSTER_NAME}" == "" ]; then echo "[ERROR] missing <cluster name>"; exit -1; fi

# 3. Node Name
if [ "$#" -gt 2 ]; then v_NODE_NAME="$3"; else	v_NODE_NAME="${NODE_NAME}"; fi
if [ "${v_NODE_NAME}" == "" ]; then 
	read -e -p "Node name  ? : "  v_NODE_NAME
fi
if [ "${v_NODE_NAME}" == "" ]; then echo "[ERROR] missing <node name>"; exit -1; fi

c_URL_MCKS_NS="${c_URL_MCKS}/ns/${v_NAMESPACE}"


# ------------------------------------------------------------------------------
# print info.
echo ""
echo "[INFO]"
echo "- Namespace                  is '${v_NAMESPACE}'"
echo "- Cluster name               is '${v_CLUSTER_NAME}'"
echo "- Node name                  is '${v_NODE_NAME}'"


# ------------------------------------------------------------------------------
# remove node
delete() {

	if [ "$MCKS_CALL_METHOD" == "REST" ]; then
		
		curl -sX DELETE ${c_URL_MCKS_NS}/clusters/${v_CLUSTER_NAME}/nodes/${v_NODE_NAME}    -H "${c_CT}" | jq;

	elif [ "$MCKS_CALL_METHOD" == "GRPC" ]; then

		$APP_ROOT/src/grpc-api/cbadm/cbadm node remove --config $APP_ROOT/src/grpc-api/cbadm/grpc_conf.yaml -o json --ns ${v_NAMESPACE} --cluster ${v_CLUSTER_NAME} --node ${v_NODE_NAME}
		
	else
		echo "[ERROR] missing MCKS_CALL_METHOD"; exit -1;
	fi
	
}

# ------------------------------------------------------------------------------
if [ "$1" != "-h" ]; then 
	echo ""
	echo "------------------------------------------------------------------------------"
	delete;
fi
