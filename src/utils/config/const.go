package config

type CSP string
type VMStatus string

const (
	CONTROL_PLANE = "control-plane"
	WORKER        = "worker"

	CSP_AWS       CSP = "aws"
	CSP_GCP       CSP = "gcp"
	CSP_AZURE     CSP = "azure"
	CSP_ALIBABA   CSP = "alibaba"
	CSP_TENCENT   CSP = "tencent"
	CSP_OPENSTACK CSP = "openstack"
	CSP_IBM       CSP = "ibm"

	NETWORKCNI_KILO  = "kilo"
	NETWORKCNI_CANAL = "canal"

	POD_CIDR       = "10.244.0.0/16"
	SERVICE_CIDR   = "10.96.0.0/12"
	SERVICE_DOMAIN = "cluster.local"

	Creating    VMStatus = "Creating" // from launch to running
	Running     VMStatus = "Running"
	Suspending  VMStatus = "Suspending" // from running to suspended
	Suspended   VMStatus = "Suspended"
	Resuming    VMStatus = "Resuming"    // from suspended to running
	Rebooting   VMStatus = "Rebooting"   // from running to running
	Terminating VMStatus = "Terminating" // from running, suspended to terminated
	Terminated  VMStatus = "Terminated"
	NotExist    VMStatus = "NotExist" // VM does not exist
	Failed      VMStatus = "Failed"

	LABEL_KEY_CSP    = "topology.cloud-barista.github.io/csp"
	LABEL_KEY_REGION = "topology.kubernetes.io/region"
	LABEL_KEY_ZONE   = "topology.kubernetes.io/zone"

	MCIS_LABEL       = "mcks"
	MCIS_SYSTEMLABEL = "Managed by MCKS"
)
