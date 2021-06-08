/*
 Copyright 2021. The KubeVela Authors.

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

package v1beta1

import (
	runtimev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
)

// WorkflowStepDefinitionSpec defines the desired state of WorkflowStepDefinition
type WorkflowStepDefinitionSpec struct {
	// Reference to the CustomResourceDefinition that defines this trait kind.
	Reference common.DefinitionReference `json:"definitionRef,omitempty"`

	// Schematic defines the data format and template of the encapsulation of the workflow step definition
	// +optional
	Schematic *common.Schematic `json:"schematic,omitempty"`
}

// WorkflowStepDefinitionStatus is the status of WorkflowStepDefinition
type WorkflowStepDefinitionStatus struct {
	// ConditionedStatus reflects the observed status of a resource
	runtimev1alpha1.ConditionedStatus `json:",inline"`

	// LatestRevision of the component definition
	// +optional
	LatestRevision *common.Revision `json:"latestRevision,omitempty"`
}

// SetConditions set condition for WorkflowStepDefinition
func (d *WorkflowStepDefinition) SetConditions(c ...runtimev1alpha1.Condition) {
	d.Status.SetConditions(c...)
}

// GetCondition gets condition from WorkflowStepDefinition
func (d *WorkflowStepDefinition) GetCondition(conditionType runtimev1alpha1.ConditionType) runtimev1alpha1.Condition {
	return d.Status.GetCondition(conditionType)
}

// +kubebuilder:object:root=true

// WorkflowStepDefinition is the Schema for the workflowstepdefinitions API
// +kubebuilder:resource:scope=Namespaced,categories={oam},shortName=workflowstep
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
type WorkflowStepDefinition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WorkflowStepDefinitionSpec   `json:"spec,omitempty"`
	Status WorkflowStepDefinitionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkflowStepDefinitionList contains a list of WorkflowStepDefinition
type WorkflowStepDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorkflowStepDefinition `json:"items"`
}
