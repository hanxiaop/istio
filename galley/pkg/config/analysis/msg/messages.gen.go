// GENERATED FILE -- DO NOT EDIT
//

package msg

import (
	"istio.io/api/analysis/v1alpha1"
	"istio.io/istio/galley/pkg/config/analysis/diag"
	"istio.io/istio/pkg/config/resource"
)

var (
	// InternalError defines a NewMessageBase for message "InternalError".
	// Description: There was an internal error in the toolchain. This is almost always a bug in the implementation.
	InternalError = diag.NewMessageBase(diag.Error, "InternalError", "IST0001", "https://istio.io/latest/docs/reference/config/analysis/ist0001/")

	// Deprecated defines a NewMessageBase for message "Deprecated".
	// Description: A feature that the configuration is depending on is now deprecated.
	Deprecated = diag.NewMessageBase(diag.Warning, "Deprecated", "IST0002", "https://istio.io/latest/docs/reference/config/analysis/ist0002/")

	// ReferencedResourceNotFound defines a NewMessageBase for message "ReferencedResourceNotFound".
	// Description: A resource being referenced does not exist.
	ReferencedResourceNotFound = diag.NewMessageBase(diag.Error, "ReferencedResourceNotFound", "IST0101", "https://istio.io/latest/docs/reference/config/analysis/ist0101/")

	// NamespaceNotInjected defines a NewMessageBase for message "NamespaceNotInjected".
	// Description: A namespace is not enabled for Istio injection.
	NamespaceNotInjected = diag.NewMessageBase(diag.Info, "NamespaceNotInjected", "IST0102", "https://istio.io/latest/docs/reference/config/analysis/ist0102/")

	// PodMissingProxy defines a NewMessageBase for message "PodMissingProxy".
	// Description: A pod is missing the Istio proxy.
	PodMissingProxy = diag.NewMessageBase(diag.Warning, "PodMissingProxy", "IST0103", "https://istio.io/latest/docs/reference/config/analysis/ist0103/")

	// GatewayPortNotOnWorkload defines a NewMessageBase for message "GatewayPortNotOnWorkload".
	// Description: Unhandled gateway port
	GatewayPortNotOnWorkload = diag.NewMessageBase(diag.Warning, "GatewayPortNotOnWorkload", "IST0104", "https://istio.io/latest/docs/reference/config/analysis/ist0104/")

	// IstioProxyImageMismatch defines a NewMessageBase for message "IstioProxyImageMismatch".
	// Description: The image of the Istio proxy running on the pod does not match the image defined in the injection configuration.
	IstioProxyImageMismatch = diag.NewMessageBase(diag.Warning, "IstioProxyImageMismatch", "IST0105", "https://istio.io/latest/docs/reference/config/analysis/ist0105/")

	// SchemaValidationError defines a NewMessageBase for message "SchemaValidationError".
	// Description: The resource has a schema validation error.
	SchemaValidationError = diag.NewMessageBase(diag.Error, "SchemaValidationError", "IST0106", "https://istio.io/latest/docs/reference/config/analysis/ist0106/")

	// MisplacedAnnotation defines a NewMessageBase for message "MisplacedAnnotation".
	// Description: An Istio annotation is applied to the wrong kind of resource.
	MisplacedAnnotation = diag.NewMessageBase(diag.Warning, "MisplacedAnnotation", "IST0107", "https://istio.io/latest/docs/reference/config/analysis/ist0107/")

	// UnknownAnnotation defines a NewMessageBase for message "UnknownAnnotation".
	// Description: An Istio annotation is not recognized for any kind of resource
	UnknownAnnotation = diag.NewMessageBase(diag.Warning, "UnknownAnnotation", "IST0108", "https://istio.io/latest/docs/reference/config/analysis/ist0108/")

	// ConflictingMeshGatewayVirtualServiceHosts defines a NewMessageBase for message "ConflictingMeshGatewayVirtualServiceHosts".
	// Description: Conflicting hosts on VirtualServices associated with mesh gateway
	ConflictingMeshGatewayVirtualServiceHosts = diag.NewMessageBase(diag.Error, "ConflictingMeshGatewayVirtualServiceHosts", "IST0109", "https://istio.io/latest/docs/reference/config/analysis/ist0109/")

	// ConflictingSidecarWorkloadSelectors defines a NewMessageBase for message "ConflictingSidecarWorkloadSelectors".
	// Description: A Sidecar resource selects the same workloads as another Sidecar resource
	ConflictingSidecarWorkloadSelectors = diag.NewMessageBase(diag.Error, "ConflictingSidecarWorkloadSelectors", "IST0110", "https://istio.io/latest/docs/reference/config/analysis/ist0110/")

	// MultipleSidecarsWithoutWorkloadSelectors defines a NewMessageBase for message "MultipleSidecarsWithoutWorkloadSelectors".
	// Description: More than one sidecar resource in a namespace has no workload selector
	MultipleSidecarsWithoutWorkloadSelectors = diag.NewMessageBase(diag.Error, "MultipleSidecarsWithoutWorkloadSelectors", "IST0111", "https://istio.io/latest/docs/reference/config/analysis/ist0111/")

	// VirtualServiceDestinationPortSelectorRequired defines a NewMessageBase for message "VirtualServiceDestinationPortSelectorRequired".
	// Description: A VirtualService routes to a service with more than one port exposed, but does not specify which to use.
	VirtualServiceDestinationPortSelectorRequired = diag.NewMessageBase(diag.Error, "VirtualServiceDestinationPortSelectorRequired", "IST0112", "https://istio.io/latest/docs/reference/config/analysis/ist0112/")

	// MTLSPolicyConflict defines a NewMessageBase for message "MTLSPolicyConflict".
	// Description: A DestinationRule and Policy are in conflict with regards to mTLS.
	MTLSPolicyConflict = diag.NewMessageBase(diag.Error, "MTLSPolicyConflict", "IST0113", "https://istio.io/latest/docs/reference/config/analysis/ist0113/")

	// DeploymentAssociatedToMultipleServices defines a NewMessageBase for message "DeploymentAssociatedToMultipleServices".
	// Description: The resulting pods of a service mesh deployment can't be associated with multiple services using the same port but different protocols.
	DeploymentAssociatedToMultipleServices = diag.NewMessageBase(diag.Warning, "DeploymentAssociatedToMultipleServices", "IST0116", "https://istio.io/latest/docs/reference/config/analysis/ist0116/")

	// DeploymentRequiresServiceAssociated defines a NewMessageBase for message "DeploymentRequiresServiceAssociated".
	// Description: The resulting pods of a service mesh deployment must be associated with at least one service.
	DeploymentRequiresServiceAssociated = diag.NewMessageBase(diag.Warning, "DeploymentRequiresServiceAssociated", "IST0117", "https://istio.io/latest/docs/reference/config/analysis/ist0117/")

	// PortNameIsNotUnderNamingConvention defines a NewMessageBase for message "PortNameIsNotUnderNamingConvention".
	// Description: Port name is not under naming convention. Protocol detection is applied to the port.
	PortNameIsNotUnderNamingConvention = diag.NewMessageBase(diag.Info, "PortNameIsNotUnderNamingConvention", "IST0118", "https://istio.io/latest/docs/reference/config/analysis/ist0118/")

	// JwtFailureDueToInvalidServicePortPrefix defines a NewMessageBase for message "JwtFailureDueToInvalidServicePortPrefix".
	// Description: Authentication policy with JWT targets Service with invalid port specification.
	JwtFailureDueToInvalidServicePortPrefix = diag.NewMessageBase(diag.Warning, "JwtFailureDueToInvalidServicePortPrefix", "IST0119", "https://istio.io/latest/docs/reference/config/analysis/ist0119/")

	// InvalidRegexp defines a NewMessageBase for message "InvalidRegexp".
	// Description: Invalid Regex
	InvalidRegexp = diag.NewMessageBase(diag.Warning, "InvalidRegexp", "IST0122", "https://istio.io/latest/docs/reference/config/analysis/ist0122/")

	// NamespaceMultipleInjectionLabels defines a NewMessageBase for message "NamespaceMultipleInjectionLabels".
	// Description: A namespace has both new and legacy injection labels
	NamespaceMultipleInjectionLabels = diag.NewMessageBase(diag.Warning, "NamespaceMultipleInjectionLabels", "IST0123", "https://istio.io/latest/docs/reference/config/analysis/ist0123/")

	// InvalidAnnotation defines a NewMessageBase for message "InvalidAnnotation".
	// Description: An Istio annotation that is not valid
	InvalidAnnotation = diag.NewMessageBase(diag.Warning, "InvalidAnnotation", "IST0125", "https://istio.io/latest/docs/reference/config/analysis/ist0125/")

	// UnknownMeshNetworksServiceRegistry defines a NewMessageBase for message "UnknownMeshNetworksServiceRegistry".
	// Description: A service registry in Mesh Networks is unknown
	UnknownMeshNetworksServiceRegistry = diag.NewMessageBase(diag.Error, "UnknownMeshNetworksServiceRegistry", "IST0126", "")

	// NoMatchingWorkloadsFound defines a NewMessageBase for message "NoMatchingWorkloadsFound".
	// Description: There aren't workloads matching the resource labels
	NoMatchingWorkloadsFound = diag.NewMessageBase(diag.Warning, "NoMatchingWorkloadsFound", "IST0127", "https://istio.io/latest/docs/reference/config/analysis/ist0127/")

	// NoServerCertificateVerificationDestinationLevel defines a NewMessageBase for message "NoServerCertificateVerificationDestinationLevel".
	// Description: No caCertificates are set in DestinationRule, this results in no verification of presented server certificate.
	NoServerCertificateVerificationDestinationLevel = diag.NewMessageBase(diag.Error, "NoServerCertificateVerificationDestinationLevel", "IST0128", "https://istio.io/latest/docs/reference/config/analysis/ist0128/")

	// NoServerCertificateVerificationPortLevel defines a NewMessageBase for message "NoServerCertificateVerificationPortLevel".
	// Description: No caCertificates are set in DestinationRule, this results in no verification of presented server certificate for traffic to a given port.
	NoServerCertificateVerificationPortLevel = diag.NewMessageBase(diag.Warning, "NoServerCertificateVerificationPortLevel", "IST0129", "https://istio.io/latest/docs/reference/config/analysis/ist0129/")

	// VirtualServiceUnreachableRule defines a NewMessageBase for message "VirtualServiceUnreachableRule".
	// Description: A VirtualService rule will never be used because a previous rule uses the same match.
	VirtualServiceUnreachableRule = diag.NewMessageBase(diag.Warning, "VirtualServiceUnreachableRule", "IST0130", "https://istio.io/latest/docs/reference/config/analysis/ist0130/")

	// VirtualServiceIneffectiveMatch defines a NewMessageBase for message "VirtualServiceIneffectiveMatch".
	// Description: A VirtualService rule match duplicates a match in a previous rule.
	VirtualServiceIneffectiveMatch = diag.NewMessageBase(diag.Info, "VirtualServiceIneffectiveMatch", "IST0131", "https://istio.io/latest/docs/reference/config/analysis/ist0131/")

	// VirtualServiceHostNotFoundInGateway defines a NewMessageBase for message "VirtualServiceHostNotFoundInGateway".
	// Description: Host defined in VirtualService not found in Gateway.
	VirtualServiceHostNotFoundInGateway = diag.NewMessageBase(diag.Warning, "VirtualServiceHostNotFoundInGateway", "IST0132", "https://istio.io/latest/docs/reference/config/analysis/ist0132/")

	// SchemaWarning defines a NewMessageBase for message "SchemaWarning".
	// Description: The resource has a schema validation warning.
	SchemaWarning = diag.NewMessageBase(diag.Warning, "SchemaWarning", "IST0133", "")

	// ServiceEntryAddressesRequired defines a NewMessageBase for message "ServiceEntryAddressesRequired".
	// Description: Virtual IP addresses are required for ports serving TCP (or unset) protocol
	ServiceEntryAddressesRequired = diag.NewMessageBase(diag.Warning, "ServiceEntryAddressesRequired", "IST0134", "https://istio.io/latest/docs/reference/config/analysis/ist0134/")

	// DeprecatedAnnotation defines a NewMessageBase for message "DeprecatedAnnotation".
	// Description: A resource is using a deprecated Istio annotation.
	DeprecatedAnnotation = diag.NewMessageBase(diag.Info, "DeprecatedAnnotation", "IST0135", "https://istio.io/latest/docs/reference/config/analysis/ist0135/")

	// AlphaAnnotation defines a NewMessageBase for message "AlphaAnnotation".
	// Description: An Istio annotation may not be suitable for production.
	AlphaAnnotation = diag.NewMessageBase(diag.Info, "AlphaAnnotation", "IST0136", "https://istio.io/latest/docs/reference/config/analysis/ist0136/")

	// DeploymentConflictingPorts defines a NewMessageBase for message "DeploymentConflictingPorts".
	// Description: Two services selecting the same workload with the same targetPort MUST refer to the same port.
	DeploymentConflictingPorts = diag.NewMessageBase(diag.Warning, "DeploymentConflictingPorts", "IST0137", "https://istio.io/latest/docs/reference/config/analysis/ist0137/")

	// GatewayDuplicateCertificate defines a NewMessageBase for message "GatewayDuplicateCertificate".
	// Description: Duplicate certificate in multiple gateways may cause 404s if clients re-use HTTP2 connections.
	GatewayDuplicateCertificate = diag.NewMessageBase(diag.Warning, "GatewayDuplicateCertificate", "IST0138", "")

	// InvalidWebhook defines a NewMessageBase for message "InvalidWebhook".
	// Description: Webhook is invalid or references a control plane service that does not exist.
	InvalidWebhook = diag.NewMessageBase(diag.Error, "InvalidWebhook", "IST0139", "")

	// IngressRouteRulesNotAffected defines a NewMessageBase for message "IngressRouteRulesNotAffected".
	// Description: Route rules have no effect on ingress gateway requests
	IngressRouteRulesNotAffected = diag.NewMessageBase(diag.Warning, "IngressRouteRulesNotAffected", "IST0140", "")

	// InsufficientPermissions defines a NewMessageBase for message "InsufficientPermissions".
	// Description: Required permissions to install Istio are missing.
	InsufficientPermissions = diag.NewMessageBase(diag.Error, "InsufficientPermissions", "IST0141", "")

	// UnsupportedKubernetesVersion defines a NewMessageBase for message "UnsupportedKubernetesVersion".
	// Description: The Kubernetes version is not supported
	UnsupportedKubernetesVersion = diag.NewMessageBase(diag.Error, "UnsupportedKubernetesVersion", "IST0142", "")

	// LocalhostListener defines a NewMessageBase for message "LocalhostListener".
	// Description: A port exposed in a Service is bound to a localhost address
	LocalhostListener = diag.NewMessageBase(diag.Error, "LocalhostListener", "IST0143", "https://istio.io/latest/docs/reference/config/analysis/ist0143/")

	// InvalidApplicationUID defines a NewMessageBase for message "InvalidApplicationUID".
	// Description: Application pods should not run as user ID (UID) 1337
	InvalidApplicationUID = diag.NewMessageBase(diag.Warning, "InvalidApplicationUID", "IST0144", "")

	// ConflictingGateways defines a NewMessageBase for message "ConflictingGateways".
	// Description: Gateway should not have the same selector, port and matched hosts of server
	ConflictingGateways = diag.NewMessageBase(diag.Error, "ConflictingGateways", "IST0145", "")
)

// All returns a list of all known message types.
func All() []*v1alpha1.AnalysisMessageBase {
	return []*v1alpha1.AnalysisMessageBase{
		InternalError,
		Deprecated,
		ReferencedResourceNotFound,
		NamespaceNotInjected,
		PodMissingProxy,
		GatewayPortNotOnWorkload,
		IstioProxyImageMismatch,
		SchemaValidationError,
		MisplacedAnnotation,
		UnknownAnnotation,
		ConflictingMeshGatewayVirtualServiceHosts,
		ConflictingSidecarWorkloadSelectors,
		MultipleSidecarsWithoutWorkloadSelectors,
		VirtualServiceDestinationPortSelectorRequired,
		MTLSPolicyConflict,
		DeploymentAssociatedToMultipleServices,
		DeploymentRequiresServiceAssociated,
		PortNameIsNotUnderNamingConvention,
		JwtFailureDueToInvalidServicePortPrefix,
		InvalidRegexp,
		NamespaceMultipleInjectionLabels,
		InvalidAnnotation,
		UnknownMeshNetworksServiceRegistry,
		NoMatchingWorkloadsFound,
		NoServerCertificateVerificationDestinationLevel,
		NoServerCertificateVerificationPortLevel,
		VirtualServiceUnreachableRule,
		VirtualServiceIneffectiveMatch,
		VirtualServiceHostNotFoundInGateway,
		SchemaWarning,
		ServiceEntryAddressesRequired,
		DeprecatedAnnotation,
		AlphaAnnotation,
		DeploymentConflictingPorts,
		GatewayDuplicateCertificate,
		InvalidWebhook,
		IngressRouteRulesNotAffected,
		InsufficientPermissions,
		UnsupportedKubernetesVersion,
		LocalhostListener,
		InvalidApplicationUID,
		ConflictingGateways,
	}
}

// NewInternalError returns a new diag.Message based on InternalError.
func NewInternalError(r *resource.Instance, detail string) diag.Message {
	return diag.NewMessage(
		InternalError,
		"There was an internal error in the toolchain. This is almost always a bug in the implementation.",
		"Internal error: %v",
		r,
		[]*v1alpha1.AnalysisMessageWeakSchema_ArgType{
			{
				Name:   "detail",
				GoType: "string",
			},
		},
		detail,
	)
}

// NewDeprecated returns a new diag.Message based on Deprecated.
func NewDeprecated(r *resource.Instance, detail string) diag.Message {
	return diag.NewMessage(
		Deprecated,
		"A feature that the configuration is depending on is now deprecated.",
		"Deprecated: %s",
		r,
		[]*v1alpha1.AnalysisMessageWeakSchema_ArgType{
			{
				Name:   "detail",
				GoType: "string",
			},
		},
		detail,
	)
}

// NewReferencedResourceNotFound returns a new diag.Message based on ReferencedResourceNotFound.
func NewReferencedResourceNotFound(r *resource.Instance, reftype string, refval string) diag.Message {
	return diag.NewMessage(
		ReferencedResourceNotFound,
		"A resource being referenced does not exist.",
		"Referenced %s not found: %q",
		r,
		[]*v1alpha1.AnalysisMessageWeakSchema_ArgType{
			{
				Name:   "reftype",
				GoType: "string",
			},
			{
				Name:   "refval",
				GoType: "string",
			},
		},
		reftype,
		refval,
	)
}

// NewNamespaceNotInjected returns a new diag.Message based on NamespaceNotInjected.
func NewNamespaceNotInjected(r *resource.Instance, namespace string, namespace2 string) diag.Message {
	return diag.NewMessage(
		NamespaceNotInjected,
		"A namespace is not enabled for Istio injection.",
		"The namespace is not enabled for Istio injection. Run 'kubectl label namespace %s istio-injection=enabled' to enable it, or 'kubectl label namespace %s istio-injection=disabled' to explicitly mark it as not needing injection.",
		r,
		[]*v1alpha1.AnalysisMessageWeakSchema_ArgType{
			{
				Name:   "namespace",
				GoType: "string",
			},
			{
				Name:   "namespace2",
				GoType: "string",
			},
		},
		namespace,
		namespace2,
	)
}

// NewPodMissingProxy returns a new diag.Message based on PodMissingProxy.
func NewPodMissingProxy(r *resource.Instance) diag.Message {
	return diag.NewMessage(
		PodMissingProxy,
		"A pod is missing the Istio proxy.",
		"The pod is missing the Istio proxy. This can often be resolved by restarting or redeploying the workload.",
		r,
		[]*v1alpha1.AnalysisMessageWeakSchema_ArgType{},
	)
}

// NewGatewayPortNotOnWorkload returns a new diag.Message based on GatewayPortNotOnWorkload.
func NewGatewayPortNotOnWorkload(r *resource.Instance, selector string, port int) diag.Message {
	return diag.NewMessage(
		GatewayPortNotOnWorkload,
		"Unhandled gateway port",
		"The gateway refers to a port that is not exposed on the workload (pod selector %s; port %d)",
		r,
		[]*v1alpha1.AnalysisMessageWeakSchema_ArgType{
			{
				Name:   "selector",
				GoType: "string",
			},
			{
				Name:   "port",
				GoType: "int",
			},
		},
		selector,
		port,
	)
}

// NewIstioProxyImageMismatch returns a new diag.Message based on IstioProxyImageMismatch.
func NewIstioProxyImageMismatch(r *resource.Instance, proxyImage string, injectionImage string) diag.Message {
	return diag.NewMessage(
		IstioProxyImageMismatch,
		"The image of the Istio proxy running on the pod does not match the image defined in the injection configuration.",
		"The image of the Istio proxy running on the pod does not match the image defined in the injection configuration (pod image: %s; injection configuration image: %s). This often happens after upgrading the Istio control-plane and can be fixed by redeploying the pod.",
		r,
		[]*v1alpha1.AnalysisMessageWeakSchema_ArgType{
			{
				Name:   "proxyImage",
				GoType: "string",
			},
			{
				Name:   "injectionImage",
				GoType: "string",
			},
		},
		proxyImage,
		injectionImage,
	)
}

// NewSchemaValidationError returns a new diag.Message based on SchemaValidationError.
func NewSchemaValidationError(r *resource.Instance, err error) diag.Message {
	return diag.NewMessage(
		SchemaValidationError,
		"The resource has a schema validation error.",
		"Schema validation error: %v",
		r,
		[]*v1alpha1.AnalysisMessageWeakSchema_ArgType{
			{
				Name:   "err",
				GoType: "error",
			},
		},
		err,
	)
}

// NewMisplacedAnnotation returns a new diag.Message based on MisplacedAnnotation.
func NewMisplacedAnnotation(r *resource.Instance, annotation string, kind string) diag.Message {
	return diag.NewMessage(
		MisplacedAnnotation,
		"An Istio annotation is applied to the wrong kind of resource.",
		"Misplaced annotation: %s can only be applied to %s",
		r,
		[]*v1alpha1.AnalysisMessageWeakSchema_ArgType{
			{
				Name:   "annotation",
				GoType: "string",
			},
			{
				Name:   "kind",
				GoType: "string",
			},
		},
		annotation,
		kind,
	)
}

// NewUnknownAnnotation returns a new diag.Message based on UnknownAnnotation.
func NewUnknownAnnotation(r *resource.Instance, annotation string) diag.Message {
	return diag.NewMessage(
		UnknownAnnotation,
		"An Istio annotation is not recognized for any kind of resource",
		"Unknown annotation: %s",
		r,
		[]*v1alpha1.AnalysisMessageWeakSchema_ArgType{
			{
				Name:   "annotation",
				GoType: "string",
			},
		},
		annotation,
	)
}

// NewConflictingMeshGatewayVirtualServiceHosts returns a new diag.Message based on ConflictingMeshGatewayVirtualServiceHosts.
func NewConflictingMeshGatewayVirtualServiceHosts(r *resource.Instance, virtualServices string, host string) diag.Message {
	return diag.NewMessage(
		ConflictingMeshGatewayVirtualServiceHosts,
		"Conflicting hosts on VirtualServices associated with mesh gateway",
		"The VirtualServices %s associated with mesh gateway define the same host %s which can lead to undefined behavior. This can be fixed by merging the conflicting VirtualServices into a single resource.",
		r,
		[]*v1alpha1.AnalysisMessageWeakSchema_ArgType{
			{
				Name:   "virtualServices",
				GoType: "string",
			},
			{
				Name:   "host",
				GoType: "string",
			},
		},
		virtualServices,
		host,
	)
}

// NewConflictingSidecarWorkloadSelectors returns a new diag.Message based on ConflictingSidecarWorkloadSelectors.
func NewConflictingSidecarWorkloadSelectors(r *resource.Instance, conflictingSidecars []string, namespace string, workloadPod string) diag.Message {
	return diag.NewMessage(
		ConflictingSidecarWorkloadSelectors,
		"A Sidecar resource selects the same workloads as another Sidecar resource",
		"The Sidecars %v in namespace %q select the same workload pod %q, which can lead to undefined behavior.",
		r,
		[]*v1alpha1.AnalysisMessageWeakSchema_ArgType{
			{
				Name:   "conflictingSidecars",
				GoType: "[]string",
			},
			{
				Name:   "namespace",
				GoType: "string",
			},
			{
				Name:   "workloadPod",
				GoType: "string",
			},
		},
		conflictingSidecars,
		namespace,
		workloadPod,
	)
}

// NewMultipleSidecarsWithoutWorkloadSelectors returns a new diag.Message based on MultipleSidecarsWithoutWorkloadSelectors.
func NewMultipleSidecarsWithoutWorkloadSelectors(r *resource.Instance, conflictingSidecars []string, namespace string) diag.Message {
	return diag.NewMessage(
		MultipleSidecarsWithoutWorkloadSelectors,
		"More than one sidecar resource in a namespace has no workload selector",
		"The Sidecars %v in namespace %q have no workload selector, which can lead to undefined behavior.",
		r,
		[]*v1alpha1.AnalysisMessageWeakSchema_ArgType{
			{
				Name:   "conflictingSidecars",
				GoType: "[]string",
			},
			{
				Name:   "namespace",
				GoType: "string",
			},
		},
		conflictingSidecars,
		namespace,
	)
}

// NewVirtualServiceDestinationPortSelectorRequired returns a new diag.Message based on VirtualServiceDestinationPortSelectorRequired.
func NewVirtualServiceDestinationPortSelectorRequired(r *resource.Instance, destHost string, destPorts []int) diag.Message {
	return diag.NewMessage(
		VirtualServiceDestinationPortSelectorRequired,
		"A VirtualService routes to a service with more than one port exposed, but does not specify which to use.",
		"This VirtualService routes to a service %q that exposes multiple ports %v. Specifying a port in the destination is required to disambiguate.",
		r,
		[]*v1alpha1.AnalysisMessageWeakSchema_ArgType{
			{
				Name:   "destHost",
				GoType: "string",
			},
			{
				Name:   "destPorts",
				GoType: "[]int",
			},
		},
		destHost,
		destPorts,
	)
}

// NewMTLSPolicyConflict returns a new diag.Message based on MTLSPolicyConflict.
func NewMTLSPolicyConflict(r *resource.Instance, host string, destinationRuleName string, destinationRuleMTLSMode bool, policyName string, policyMTLSMode string) diag.Message {
	return diag.NewMessage(
		MTLSPolicyConflict,
		"A DestinationRule and Policy are in conflict with regards to mTLS.",
		"A DestinationRule and Policy are in conflict with regards to mTLS for host %s. The DestinationRule %q specifies that mTLS must be %t but the Policy object %q specifies %s.",
		r,
		[]*v1alpha1.AnalysisMessageWeakSchema_ArgType{
			{
				Name:   "host",
				GoType: "string",
			},
			{
				Name:   "destinationRuleName",
				GoType: "string",
			},
			{
				Name:   "destinationRuleMTLSMode",
				GoType: "bool",
			},
			{
				Name:   "policyName",
				GoType: "string",
			},
			{
				Name:   "policyMTLSMode",
				GoType: "string",
			},
		},
		host,
		destinationRuleName,
		destinationRuleMTLSMode,
		policyName,
		policyMTLSMode,
	)
}

// NewDeploymentAssociatedToMultipleServices returns a new diag.Message based on DeploymentAssociatedToMultipleServices.
func NewDeploymentAssociatedToMultipleServices(r *resource.Instance, deployment string, port int32, services []string) diag.Message {
	return diag.NewMessage(
		DeploymentAssociatedToMultipleServices,
		"The resulting pods of a service mesh deployment can't be associated with multiple services using the same port but different protocols.",
		"This deployment %s is associated with multiple services using port %d but different protocols: %v",
		r,
		[]*v1alpha1.AnalysisMessageWeakSchema_ArgType{
			{
				Name:   "deployment",
				GoType: "string",
			},
			{
				Name:   "port",
				GoType: "int32",
			},
			{
				Name:   "services",
				GoType: "[]string",
			},
		},
		deployment,
		port,
		services,
	)
}

// NewDeploymentRequiresServiceAssociated returns a new diag.Message based on DeploymentRequiresServiceAssociated.
func NewDeploymentRequiresServiceAssociated(r *resource.Instance) diag.Message {
	return diag.NewMessage(
		DeploymentRequiresServiceAssociated,
		"The resulting pods of a service mesh deployment must be associated with at least one service.",
		"No service associated with this deployment. Service mesh deployments must be associated with a service.",
		r,
		[]*v1alpha1.AnalysisMessageWeakSchema_ArgType{},
	)
}

// NewPortNameIsNotUnderNamingConvention returns a new diag.Message based on PortNameIsNotUnderNamingConvention.
func NewPortNameIsNotUnderNamingConvention(r *resource.Instance, portName string, port int, targetPort string) diag.Message {
	return diag.NewMessage(
		PortNameIsNotUnderNamingConvention,
		"Port name is not under naming convention. Protocol detection is applied to the port.",
		"Port name %s (port: %d, targetPort: %s) doesn't follow the naming convention of Istio port.",
		r,
		[]*v1alpha1.AnalysisMessageWeakSchema_ArgType{
			{
				Name:   "portName",
				GoType: "string",
			},
			{
				Name:   "port",
				GoType: "int",
			},
			{
				Name:   "targetPort",
				GoType: "string",
			},
		},
		portName,
		port,
		targetPort,
	)
}

// NewJwtFailureDueToInvalidServicePortPrefix returns a new diag.Message based on JwtFailureDueToInvalidServicePortPrefix.
func NewJwtFailureDueToInvalidServicePortPrefix(r *resource.Instance, port int, portName string, protocol string, targetPort string) diag.Message {
	return diag.NewMessage(
		JwtFailureDueToInvalidServicePortPrefix,
		"Authentication policy with JWT targets Service with invalid port specification.",
		"Authentication policy with JWT targets Service with invalid port specification (port: %d, name: %s, protocol: %s, targetPort: %s).",
		r,
		[]*v1alpha1.AnalysisMessageWeakSchema_ArgType{
			{
				Name:   "port",
				GoType: "int",
			},
			{
				Name:   "portName",
				GoType: "string",
			},
			{
				Name:   "protocol",
				GoType: "string",
			},
			{
				Name:   "targetPort",
				GoType: "string",
			},
		},
		port,
		portName,
		protocol,
		targetPort,
	)
}

// NewInvalidRegexp returns a new diag.Message based on InvalidRegexp.
func NewInvalidRegexp(r *resource.Instance, where string, re string, problem string) diag.Message {
	return diag.NewMessage(
		InvalidRegexp,
		"Invalid Regex",
		"Field %q regular expression invalid: %q (%s)",
		r,
		[]*v1alpha1.AnalysisMessageWeakSchema_ArgType{
			{
				Name:   "where",
				GoType: "string",
			},
			{
				Name:   "re",
				GoType: "string",
			},
			{
				Name:   "problem",
				GoType: "string",
			},
		},
		where,
		re,
		problem,
	)
}

// NewNamespaceMultipleInjectionLabels returns a new diag.Message based on NamespaceMultipleInjectionLabels.
func NewNamespaceMultipleInjectionLabels(r *resource.Instance, namespace string, namespace2 string) diag.Message {
	return diag.NewMessage(
		NamespaceMultipleInjectionLabels,
		"A namespace has both new and legacy injection labels",
		"The namespace has both new and legacy injection labels. Run 'kubectl label namespace %s istio.io/rev-' or 'kubectl label namespace %s istio-injection-'",
		r,
		[]*v1alpha1.AnalysisMessageWeakSchema_ArgType{
			{
				Name:   "namespace",
				GoType: "string",
			},
			{
				Name:   "namespace2",
				GoType: "string",
			},
		},
		namespace,
		namespace2,
	)
}

// NewInvalidAnnotation returns a new diag.Message based on InvalidAnnotation.
func NewInvalidAnnotation(r *resource.Instance, annotation string, problem string) diag.Message {
	return diag.NewMessage(
		InvalidAnnotation,
		"An Istio annotation that is not valid",
		"Invalid annotation %s: %s",
		r,
		[]*v1alpha1.AnalysisMessageWeakSchema_ArgType{
			{
				Name:   "annotation",
				GoType: "string",
			},
			{
				Name:   "problem",
				GoType: "string",
			},
		},
		annotation,
		problem,
	)
}

// NewUnknownMeshNetworksServiceRegistry returns a new diag.Message based on UnknownMeshNetworksServiceRegistry.
func NewUnknownMeshNetworksServiceRegistry(r *resource.Instance, serviceregistry string, network string) diag.Message {
	return diag.NewMessage(
		UnknownMeshNetworksServiceRegistry,
		"A service registry in Mesh Networks is unknown",
		"Unknown service registry %s in network %s",
		r,
		[]*v1alpha1.AnalysisMessageWeakSchema_ArgType{
			{
				Name:   "serviceregistry",
				GoType: "string",
			},
			{
				Name:   "network",
				GoType: "string",
			},
		},
		serviceregistry,
		network,
	)
}

// NewNoMatchingWorkloadsFound returns a new diag.Message based on NoMatchingWorkloadsFound.
func NewNoMatchingWorkloadsFound(r *resource.Instance, labels string) diag.Message {
	return diag.NewMessage(
		NoMatchingWorkloadsFound,
		"There aren't workloads matching the resource labels",
		"No matching workloads for this resource with the following labels: %s",
		r,
		[]*v1alpha1.AnalysisMessageWeakSchema_ArgType{
			{
				Name:   "labels",
				GoType: "string",
			},
		},
		labels,
	)
}

// NewNoServerCertificateVerificationDestinationLevel returns a new diag.Message based on NoServerCertificateVerificationDestinationLevel.
func NewNoServerCertificateVerificationDestinationLevel(r *resource.Instance, destinationrule string, namespace string, mode string, host string) diag.Message {
	return diag.NewMessage(
		NoServerCertificateVerificationDestinationLevel,
		"No caCertificates are set in DestinationRule, this results in no verification of presented server certificate.",
		"DestinationRule %s in namespace %s has TLS mode set to %s but no caCertificates are set to validate server identity for host: %s",
		r,
		[]*v1alpha1.AnalysisMessageWeakSchema_ArgType{
			{
				Name:   "destinationrule",
				GoType: "string",
			},
			{
				Name:   "namespace",
				GoType: "string",
			},
			{
				Name:   "mode",
				GoType: "string",
			},
			{
				Name:   "host",
				GoType: "string",
			},
		},
		destinationrule,
		namespace,
		mode,
		host,
	)
}

// NewNoServerCertificateVerificationPortLevel returns a new diag.Message based on NoServerCertificateVerificationPortLevel.
func NewNoServerCertificateVerificationPortLevel(r *resource.Instance, destinationrule string, namespace string, mode string, host string, port string) diag.Message {
	return diag.NewMessage(
		NoServerCertificateVerificationPortLevel,
		"No caCertificates are set in DestinationRule, this results in no verification of presented server certificate for traffic to a given port.",
		"DestinationRule %s in namespace %s has TLS mode set to %s but no caCertificates are set to validate server identity for host: %s at port %s",
		r,
		[]*v1alpha1.AnalysisMessageWeakSchema_ArgType{
			{
				Name:   "destinationrule",
				GoType: "string",
			},
			{
				Name:   "namespace",
				GoType: "string",
			},
			{
				Name:   "mode",
				GoType: "string",
			},
			{
				Name:   "host",
				GoType: "string",
			},
			{
				Name:   "port",
				GoType: "string",
			},
		},
		destinationrule,
		namespace,
		mode,
		host,
		port,
	)
}

// NewVirtualServiceUnreachableRule returns a new diag.Message based on VirtualServiceUnreachableRule.
func NewVirtualServiceUnreachableRule(r *resource.Instance, ruleno string, reason string) diag.Message {
	return diag.NewMessage(
		VirtualServiceUnreachableRule,
		"A VirtualService rule will never be used because a previous rule uses the same match.",
		"VirtualService rule %v not used (%s).",
		r,
		[]*v1alpha1.AnalysisMessageWeakSchema_ArgType{
			{
				Name:   "ruleno",
				GoType: "string",
			},
			{
				Name:   "reason",
				GoType: "string",
			},
		},
		ruleno,
		reason,
	)
}

// NewVirtualServiceIneffectiveMatch returns a new diag.Message based on VirtualServiceIneffectiveMatch.
func NewVirtualServiceIneffectiveMatch(r *resource.Instance, ruleno string, matchno string, dupno string) diag.Message {
	return diag.NewMessage(
		VirtualServiceIneffectiveMatch,
		"A VirtualService rule match duplicates a match in a previous rule.",
		"VirtualService rule %v match %v is not used (duplicates a match in rule %v).",
		r,
		[]*v1alpha1.AnalysisMessageWeakSchema_ArgType{
			{
				Name:   "ruleno",
				GoType: "string",
			},
			{
				Name:   "matchno",
				GoType: "string",
			},
			{
				Name:   "dupno",
				GoType: "string",
			},
		},
		ruleno,
		matchno,
		dupno,
	)
}

// NewVirtualServiceHostNotFoundInGateway returns a new diag.Message based on VirtualServiceHostNotFoundInGateway.
func NewVirtualServiceHostNotFoundInGateway(r *resource.Instance, host []string, virtualservice string, gateway string) diag.Message {
	return diag.NewMessage(
		VirtualServiceHostNotFoundInGateway,
		"Host defined in VirtualService not found in Gateway.",
		"one or more host %v defined in VirtualService %s not found in Gateway %s.",
		r,
		[]*v1alpha1.AnalysisMessageWeakSchema_ArgType{
			{
				Name:   "host",
				GoType: "[]string",
			},
			{
				Name:   "virtualservice",
				GoType: "string",
			},
			{
				Name:   "gateway",
				GoType: "string",
			},
		},
		host,
		virtualservice,
		gateway,
	)
}

// NewSchemaWarning returns a new diag.Message based on SchemaWarning.
func NewSchemaWarning(r *resource.Instance, err error) diag.Message {
	return diag.NewMessage(
		SchemaWarning,
		"The resource has a schema validation warning.",
		"Schema validation warning: %v",
		r,
		[]*v1alpha1.AnalysisMessageWeakSchema_ArgType{
			{
				Name:   "err",
				GoType: "error",
			},
		},
		err,
	)
}

// NewServiceEntryAddressesRequired returns a new diag.Message based on ServiceEntryAddressesRequired.
func NewServiceEntryAddressesRequired(r *resource.Instance) diag.Message {
	return diag.NewMessage(
		ServiceEntryAddressesRequired,
		"Virtual IP addresses are required for ports serving TCP (or unset) protocol",
		"ServiceEntry addresses are required for this protocol.",
		r,
		[]*v1alpha1.AnalysisMessageWeakSchema_ArgType{},
	)
}

// NewDeprecatedAnnotation returns a new diag.Message based on DeprecatedAnnotation.
func NewDeprecatedAnnotation(r *resource.Instance, annotation string, extra string) diag.Message {
	return diag.NewMessage(
		DeprecatedAnnotation,
		"A resource is using a deprecated Istio annotation.",
		"Annotation %q has been deprecated%s and may not work in future Istio versions.",
		r,
		[]*v1alpha1.AnalysisMessageWeakSchema_ArgType{
			{
				Name:   "annotation",
				GoType: "string",
			},
			{
				Name:   "extra",
				GoType: "string",
			},
		},
		annotation,
		extra,
	)
}

// NewAlphaAnnotation returns a new diag.Message based on AlphaAnnotation.
func NewAlphaAnnotation(r *resource.Instance, annotation string) diag.Message {
	return diag.NewMessage(
		AlphaAnnotation,
		"An Istio annotation may not be suitable for production.",
		"Annotation %q is part of an alpha-phase feature and may be incompletely supported.",
		r,
		[]*v1alpha1.AnalysisMessageWeakSchema_ArgType{
			{
				Name:   "annotation",
				GoType: "string",
			},
		},
		annotation,
	)
}

// NewDeploymentConflictingPorts returns a new diag.Message based on DeploymentConflictingPorts.
func NewDeploymentConflictingPorts(r *resource.Instance, deployment string, services []string, targetPort string, ports []int32) diag.Message {
	return diag.NewMessage(
		DeploymentConflictingPorts,
		"Two services selecting the same workload with the same targetPort MUST refer to the same port.",
		"This deployment %s is associated with multiple services %v using targetPort %q but different ports: %v.",
		r,
		[]*v1alpha1.AnalysisMessageWeakSchema_ArgType{
			{
				Name:   "deployment",
				GoType: "string",
			},
			{
				Name:   "services",
				GoType: "[]string",
			},
			{
				Name:   "targetPort",
				GoType: "string",
			},
			{
				Name:   "ports",
				GoType: "[]int32",
			},
		},
		deployment,
		services,
		targetPort,
		ports,
	)
}

// NewGatewayDuplicateCertificate returns a new diag.Message based on GatewayDuplicateCertificate.
func NewGatewayDuplicateCertificate(r *resource.Instance, gateways []string) diag.Message {
	return diag.NewMessage(
		GatewayDuplicateCertificate,
		"Duplicate certificate in multiple gateways may cause 404s if clients re-use HTTP2 connections.",
		"Duplicate certificate in multiple gateways %v may cause 404s if clients re-use HTTP2 connections.",
		r,
		[]*v1alpha1.AnalysisMessageWeakSchema_ArgType{
			{
				Name:   "gateways",
				GoType: "[]string",
			},
		},
		gateways,
	)
}

// NewInvalidWebhook returns a new diag.Message based on InvalidWebhook.
func NewInvalidWebhook(r *resource.Instance, error string) diag.Message {
	return diag.NewMessage(
		InvalidWebhook,
		"Webhook is invalid or references a control plane service that does not exist.",
		"%v",
		r,
		[]*v1alpha1.AnalysisMessageWeakSchema_ArgType{
			{
				Name:   "error",
				GoType: "string",
			},
		},
		error,
	)
}

// NewIngressRouteRulesNotAffected returns a new diag.Message based on IngressRouteRulesNotAffected.
func NewIngressRouteRulesNotAffected(r *resource.Instance, virtualservicesubset string, virtualservice string) diag.Message {
	return diag.NewMessage(
		IngressRouteRulesNotAffected,
		"Route rules have no effect on ingress gateway requests",
		"Subset in virtual service %s has no effect on ingress gateway %s requests",
		r,
		[]*v1alpha1.AnalysisMessageWeakSchema_ArgType{
			{
				Name:   "virtualservicesubset",
				GoType: "string",
			},
			{
				Name:   "virtualservice",
				GoType: "string",
			},
		},
		virtualservicesubset,
		virtualservice,
	)
}

// NewInsufficientPermissions returns a new diag.Message based on InsufficientPermissions.
func NewInsufficientPermissions(r *resource.Instance, resource string, error string) diag.Message {
	return diag.NewMessage(
		InsufficientPermissions,
		"Required permissions to install Istio are missing.",
		"Missing required permission to create resource %v (%v)",
		r,
		[]*v1alpha1.AnalysisMessageWeakSchema_ArgType{
			{
				Name:   "resource",
				GoType: "string",
			},
			{
				Name:   "error",
				GoType: "string",
			},
		},
		resource,
		error,
	)
}

// NewUnsupportedKubernetesVersion returns a new diag.Message based on UnsupportedKubernetesVersion.
func NewUnsupportedKubernetesVersion(r *resource.Instance, version string, minimumVersion string) diag.Message {
	return diag.NewMessage(
		UnsupportedKubernetesVersion,
		"The Kubernetes version is not supported",
		"The Kubernetes Version %q is lower than the minimum version: %v",
		r,
		[]*v1alpha1.AnalysisMessageWeakSchema_ArgType{
			{
				Name:   "version",
				GoType: "string",
			},
			{
				Name:   "minimumVersion",
				GoType: "string",
			},
		},
		version,
		minimumVersion,
	)
}

// NewLocalhostListener returns a new diag.Message based on LocalhostListener.
func NewLocalhostListener(r *resource.Instance, port string) diag.Message {
	return diag.NewMessage(
		LocalhostListener,
		"A port exposed in a Service is bound to a localhost address",
		"Port %v is exposed in a Service but listens on localhost. It will not be exposed to other pods.",
		r,
		[]*v1alpha1.AnalysisMessageWeakSchema_ArgType{
			{
				Name:   "port",
				GoType: "string",
			},
		},
		port,
	)
}

// NewInvalidApplicationUID returns a new diag.Message based on InvalidApplicationUID.
func NewInvalidApplicationUID(r *resource.Instance) diag.Message {
	return diag.NewMessage(
		InvalidApplicationUID,
		"Application pods should not run as user ID (UID) 1337",
		"User ID (UID) 1337 is reserved for the sidecar proxy.",
		r,
		[]*v1alpha1.AnalysisMessageWeakSchema_ArgType{},
	)
}

// NewConflictingGateways returns a new diag.Message based on ConflictingGateways.
func NewConflictingGateways(r *resource.Instance, gateway string, selector string, portnumber string, hosts string) diag.Message {
	return diag.NewMessage(
		ConflictingGateways,
		"Gateway should not have the same selector, port and matched hosts of server",
		"Conflict with gateways %s (workload selector %s, port %s, hosts %v).",
		r,
		[]*v1alpha1.AnalysisMessageWeakSchema_ArgType{
			{
				Name:   "gateway",
				GoType: "string",
			},
			{
				Name:   "selector",
				GoType: "string",
			},
			{
				Name:   "portnumber",
				GoType: "string",
			},
			{
				Name:   "hosts",
				GoType: "string",
			},
		},
		gateway,
		selector,
		portnumber,
		hosts,
	)
}
