// Copyright (c) 2020-2024 Tigera, Inc. All rights reserved.
/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IntrusionDetectionSpec defines the desired state of Tigera intrusion detection capabilities.
type IntrusionDetectionSpec struct {
	// ComponentResources can be used to customize the resource requirements for each component.
	// Only DeepPacketInspection is supported for this spec.
	// +optional
	ComponentResources []IntrusionDetectionComponentResource `json:"componentResources,omitempty"`

	// AnomalyDetection is now deprecated, and configuring it has no effect.
	// +optional
	AnomalyDetection AnomalyDetectionSpec `json:"anomalyDetection,omitempty"`

	// IntrusionDetectionControllerDeployment configures the IntrusionDetection Controller Deployment.
	// +optional
	IntrusionDetectionControllerDeployment *IntrusionDetectionControllerDeployment `json:"intrusionDetectionControllerDeployment,omitempty"`
}

type AnomalyDetectionSpec struct {

	// StorageClassName is now deprecated, and configuring it has no effect.
	// +optional
	StorageClassName string `json:"storageClassName,omitempty"`
}

// IntrusionDetectionStatus defines the observed state of Tigera intrusion detection capabilities.
type IntrusionDetectionStatus struct {
	// State provides user-readable status.
	State string `json:"state,omitempty"`

	// Conditions represents the latest observed set of conditions for the component. A component may be one or more of
	// Ready, Progressing, Degraded or other customer types.
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster

// IntrusionDetection installs the components required for Tigera intrusion detection. At most one instance
// of this resource is supported. It must be named "tigera-secure".
type IntrusionDetection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Specification of the desired state for Tigera intrusion detection.
	Spec IntrusionDetectionSpec `json:"spec,omitempty"`
	// Most recently observed state for Tigera intrusion detection.
	Status IntrusionDetectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IntrusionDetectionList contains a list of IntrusionDetection
type IntrusionDetectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IntrusionDetection `json:"items"`
}

type IntrusionDetectionComponentName string

const (
	ComponentNameDeepPacketInspection IntrusionDetectionComponentName = "DeepPacketInspection"
)

// The ComponentResource struct associates a ResourceRequirements with a component by name
type IntrusionDetectionComponentResource struct {
	// ComponentName is an enum which identifies the component
	// +kubebuilder:validation:Enum=DeepPacketInspection
	ComponentName IntrusionDetectionComponentName `json:"componentName"`
	// ResourceRequirements allows customization of limits and requests for compute resources such as cpu and memory.
	ResourceRequirements *corev1.ResourceRequirements `json:"resourceRequirements"`
}

// IntrusionDetectionControllerDeployment is the configuration for the IntrusionDetectionController Deployment.
type IntrusionDetectionControllerDeployment struct {

	// Spec is the specification of the IntrusionDetectionController Deployment.
	// +optional
	Spec *IntrusionDetectionControllerDeploymentSpec `json:"spec,omitempty"`
}

// IntrusionDetectionControllerDeploymentSpec defines configuration for the IntrusionDetectionController Deployment.
type IntrusionDetectionControllerDeploymentSpec struct {

	// Template describes the IntrusionDetectionController Deployment pod that will be created.
	// +optional
	Template *IntrusionDetectionControllerDeploymentPodTemplateSpec `json:"template,omitempty"`
}

// IntrusionDetectionControllerDeploymentPodTemplateSpec is the IntrusionDetectionController Deployment's PodTemplateSpec
type IntrusionDetectionControllerDeploymentPodTemplateSpec struct {

	// Spec is the IntrusionDetectionController Deployment's PodSpec.
	// +optional
	Spec *IntrusionDetectionControllerDeploymentPodSpec `json:"spec,omitempty"`
}

// IntrusionDetectionControllerDeploymentPodSpec is the IntrusionDetectionController Deployment's PodSpec.
type IntrusionDetectionControllerDeploymentPodSpec struct {
	// InitContainers is a list of IntrusionDetectionController init containers.
	// If specified, this overrides the specified IntrusionDetectionController Deployment init containers.
	// If omitted, the IntrusionDetectionController Deployment will use its default values for its init containers.
	// +optional
	InitContainers []IntrusionDetectionControllerDeploymentInitContainer `json:"initContainers,omitempty"`

	// Containers is a list of IntrusionDetectionController containers.
	// If specified, this overrides the specified IntrusionDetectionController Deployment containers.
	// If omitted, the IntrusionDetectionController Deployment will use its default values for its containers.
	// +optional
	Containers []IntrusionDetectionControllerDeploymentContainer `json:"containers,omitempty"`

	// +optional
	Affinity *corev1.Affinity `json:"affinity,omitempty"`
	// +optional
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`
	// +optional
	Tolerations []corev1.Toleration `json:"tolerations,omitempty"`
}

// IntrusionDetectionControllerDeploymentContainer is a IntrusionDetectionController Deployment container.
type IntrusionDetectionControllerDeploymentContainer struct {
	// Name is an enum which identifies the IntrusionDetectionController Deployment container by name.
	// Supported values are: controller, webhooks-processor
	// +kubebuilder:validation:Enum=controller;webhooks-processor
	Name string `json:"name"`

	// Resources allows customization of limits and requests for compute resources such as cpu and memory.
	// If specified, this overrides the named IntrusionDetectionController Deployment container's resources.
	// If omitted, the IntrusionDetection Deployment will use its default value for this container's resources.
	// +optional
	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`
}

// IntrusionDetectionControllerDeploymentInitContainer is a IntrusionDetectionController Deployment init container.
type IntrusionDetectionControllerDeploymentInitContainer struct {
	// Name is an enum which identifies the IntrusionDetectionController Deployment init container by name.
	// Supported values are: intrusion-detection-tls-key-cert-provisioner
	// +kubebuilder:validation:Enum=intrusion-detection-tls-key-cert-provisioner
	Name string `json:"name"`

	// Resources allows customization of limits and requests for compute resources such as cpu and memory.
	// If specified, this overrides the named IntrusionDetectionController Deployment init container's resources.
	// If omitted, the IntrusionDetectionController Deployment will use its default value for this init container's resources.
	// +optional
	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`
}

func (c *IntrusionDetectionControllerDeployment) GetMetadata() *Metadata {
	return nil
}

func (c *IntrusionDetectionControllerDeployment) GetMinReadySeconds() *int32 {
	return nil
}

func (c *IntrusionDetectionControllerDeployment) GetPodTemplateMetadata() *Metadata {
	return nil
}

func (c *IntrusionDetectionControllerDeployment) GetInitContainers() []corev1.Container {
	if c != nil {
		if c.Spec.Template != nil {
			if c.Spec.Template.Spec != nil {
				if c.Spec.Template.Spec.InitContainers != nil {
					cs := make([]corev1.Container, len(c.Spec.Template.Spec.InitContainers))
					for i, v := range c.Spec.Template.Spec.InitContainers {
						// Only copy and return the init container if it has resources set.
						if v.Resources == nil {
							continue
						}
						c := corev1.Container{Name: v.Name, Resources: *v.Resources}
						cs[i] = c
					}
					return cs
				}
			}
		}
	}

	return nil
}

func (c *IntrusionDetectionControllerDeployment) GetContainers() []corev1.Container {
	if c != nil {
		if c.Spec != nil {
			if c.Spec.Template != nil {
				if c.Spec.Template.Spec != nil {
					if c.Spec.Template.Spec.Containers != nil {
						cs := make([]corev1.Container, len(c.Spec.Template.Spec.Containers))
						for i, v := range c.Spec.Template.Spec.Containers {
							// Only copy and return the init container if it has resources set.
							if v.Resources == nil {
								continue
							}
							c := corev1.Container{Name: v.Name, Resources: *v.Resources}
							cs[i] = c
						}
						return cs
					}
				}
			}
		}
	}
	return nil
}

func (c *IntrusionDetectionControllerDeployment) GetAffinity() *corev1.Affinity {
	return nil
}

func (c *IntrusionDetectionControllerDeployment) GetTopologySpreadConstraints() []corev1.TopologySpreadConstraint {
	return nil
}

func (c *IntrusionDetectionControllerDeployment) GetNodeSelector() map[string]string {
	return nil
}

func (c *IntrusionDetectionControllerDeployment) GetTolerations() []corev1.Toleration {
	return nil
}

func (c *IntrusionDetectionControllerDeployment) GetTerminationGracePeriodSeconds() *int64 {
	return nil
}

func (c *IntrusionDetectionControllerDeployment) GetDeploymentStrategy() *appsv1.DeploymentStrategy {
	return nil
}

func (c *IntrusionDetectionControllerDeployment) GetPriorityClassName() string {
	return ""
}

func init() {
	SchemeBuilder.Register(&IntrusionDetection{}, &IntrusionDetectionList{})
}
