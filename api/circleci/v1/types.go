package v1

// Auth configure authentication for CircleCI API.
type Auth struct {
	Token *SecretSource `json:"token,omitempty"`
}

// Owner represents the owner of the CircleCI Organization. it can be either ID or vcs + name for slug
type Owner struct {
	Id string `json:"id,omitempty"`
	// +kubebuilder:validation:Enum:={"gh","bb"}
	// +kubebuilder:default:="gh"
	VCS  string `json:"vcs,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

// LocalObjectRef represents the reference to the object in the same namespace
type LocalObjectRef struct {
	Name string `json:"name"`
}

// NamespacedLocalObjectRef represents the reference to the namespaced object
type NamespacedObjectRef struct {
	LocalObjectRef `json:",inline"`
	Namespace      string `json:"namespace,omitempty"`
}

// SecretSource represents a source for the value of a Secrets.
type SecretSource struct {
	SecretKeyRef *SecretKeySelector `json:"secretKeyRef,omitempty" protobuf:"bytes,4,opt,name=secretKeyRef"`
}

// SecretKeySelector selects a key of a Secret
// +structType=atomic
type SecretKeySelector struct {
	// The name of the secret to select from.
	Name string `json:"name"`
	// The namespace of the secret to select from.
	Namespace string `json:"namespace,omitempty"`
	// The key of the secret to select from.  Must be a valid secret key.
	Key string `json:"key"`
}

// EnvVar represents an environment variable that can be used in a Context and Project.
type EnvVar struct {
	// The name of the environment variable.
	Name string `json:"name"`
	// The value of the environment variable.
	Value string `json:"value,omitempty"`
}

// EnvVar represents an environment variable that can be used in a Context and Project.
type EnvVarFrom struct {
	EnvVar `json:",inline"`
	// The value of the environment variable as a reference to a secret key.
	ValueFrom *SecretSource `json:"valueFrom,omitempty"`
}

// SecretEnvSource
type SecretEnvSource struct {
	NamespacedObjectRef `json:",inline"`
}

// EnvFrom
type EnvFrom struct {
	SecretRef *SecretEnvSource `json:"secretRef,omitempty" protobuf:"bytes,3,opt,name=secretRef"`
}
