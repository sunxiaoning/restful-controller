/*
Copyright 2017 The Kubernetes Authors.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RestfulJob is a specification for a RestfulJob resource
type RestfulJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RestfulJobSpec   `json:"spec"`
	Status RestfulJobStatus `json:"status"`
}

// RestfulJobSpec is the spec for a RestfulJob resource

type RestfulTask struct {
	Url           string   `json:"url"`
	Method        string   `json:"method"`
	Headers       []string `json:"headers"`
	Params        string   `json:"params"`
	ResultStorage string   `json:"resultStorage"`
}

type RestfulJobSpec struct {
	Task     RestfulTask `json:"task"`
	MaxRetry *int32      `json:"maxRetry"`
	Suspend  *bool       `json:"suspend"`
}

// RestfulJobStatus is the status for a RestfulJob resource

type JobStatus string

var (
	JobCreated JobStatus = "created"
	JobPending JobStatus = "pending"
	JobRunning JobStatus = "running"
	JobSucceed JobStatus = "succeed"
	JobFailed  JobStatus = "failed"
	JobUnknown JobStatus = "unknown"
)

type RestfulJobStatus struct {
	Status   JobStatus `json:"status"`
	LogFile  string    `json:"logFile"`
	DataFile string    `json:"dataFile"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RestfulJobList is a list of RestfulJob resources
type RestfulJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []RestfulJob `json:"items"`
}
