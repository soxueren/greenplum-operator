package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// GPDBClusterSpec defines the desired state of GPDBCluster
// +k8s:openapi-gen=true
type GPDBClusterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
   masterSelector string `json:"masterselect,omitempty"`
   masterAndStandby MasterAndStandby `json:"masterselect,omitempty"`
   segments Segments `json:"segments,omitempty"`
   mirrors Mirrors `json:"mirrors,omitempty"`
}

type MasterAndStandby struct{
	Replicas int32 `json:"replicas,omitempty"`
	Image string `json:"image,omitempty"`
	hostBasedAuthentication []string `json:"hostauth,omitempty"`
	StorageClassName string `json:"storage_class_name,omitempty"`
	Storage resource.Quantity `json:"storage"`
}

type Segments struct {
    Replicas int32 `json:"replicas,omitempty"`
	Image string `json:"image,omitempty"`
	StorageClassName string `json:"storage_class_name,omitempty"`
	Storage resource.Quantity `json:"storage"`
}

type Mirrors struct {
	Image string `json:"image,omitempty"`
	StorageClassName string `json:"storage_class_name,omitempty"`
	Storage resource.Quantity `json:"storage"`
}
// GPDBClusterStatus defines the observed state of GPDBCluster
// +k8s:openapi-gen=true
type GPDBClusterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	TimeStarted metav1.Time `json:"timeStarted"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GPDBCluster is the Schema for the gpdbclusters API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=gpdbclusters,scope=Namespaced
type GPDBCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GPDBClusterSpec   `json:"spec,omitempty"`
	Status GPDBClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GPDBClusterList contains a list of GPDBCluster
type GPDBClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GPDBCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GPDBCluster{}, &GPDBClusterList{})
}
