/*
Copyright 2018 The Kubernetes Authors.

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

package service

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/gosoon/glog"
	"k8s.io/api/admission/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	eksv1 "git.yun.pingan.com/eks/cluster-operator/pkg/apis/eks/v1"
)

const (
	Version  = "v1"
	Resource = "workloadclusters"
	Group    = "eks.yun.pingan.com"

	CREATE = "CREATE"
	UPDATE = "UPDATE"
	DELETE = "DELETE"
)

// AdmitClusterOperator is check all args.
func AdmitClusterOperator(ar v1beta1.AdmissionReview) *v1beta1.AdmissionResponse {
	glog.Info("admitting clusteroperator")

	workloadClusterResource := metav1.GroupVersionResource{Group: Group, Version: Version, Resource: Resource}
	if ar.Request.Resource != workloadClusterResource {
		err := fmt.Errorf("expect resource to be %s", workloadClusterResource)
		glog.Error(err)
		return ToAdmissionResponse(err)
	}

	raw := ar.Request.Object.Raw
	workloadCluster := eksv1.WorkloadCluster{}

	if err := json.Unmarshal(raw, &workloadCluster); err != nil {
		glog.Error(err)
		return ToAdmissionResponse(err)
	}
	reviewResponse := v1beta1.AdmissionResponse{}
	reviewResponse.Allowed = true

	// verify spec
	msg := "plz check,"
	if len(workloadCluster.Spec.Cluster.Masters) == 0 {
		reviewResponse.Allowed = false
		msg += "masters is nil."
	} else {
		master := workloadCluster.Spec.Cluster.Masters[0]
		if len(master.IP) == 0 && len(master.Hostname) == 0 {
			reviewResponse.Allowed = false
			msg += "masters is nil."
		}
	}

	if len(workloadCluster.Spec.Cluster.Workers) == 0 {
		reviewResponse.Allowed = false
		msg += "work node is nil."
	} else {
		worker := workloadCluster.Spec.Cluster.Workers[0]
		if len(worker.IP) == 0 && len(worker.Hostname) == 0 {
			reviewResponse.Allowed = false
			msg += "work node is nil."
		}
	}

	if len(workloadCluster.Spec.Cluster.PrivateSSHKey) == 0 && len(workloadCluster.Spec.Cluster.RootPassword) == 0 {
		reviewResponse.Allowed = false
		msg += "private ssh key and root passwd is nil."
	}

	if ar.Request.Operation == UPDATE {
		oldWorkloadCluster := eksv1.WorkloadCluster{}
		oldRaw := ar.Request.OldObject.Raw
		if err := json.Unmarshal(oldRaw, &oldWorkloadCluster); err != nil {
			glog.Error(err)
			return ToAdmissionResponse(err)
		}
		currentMasters := workloadCluster.Spec.Cluster.Masters
		oldMasters := oldWorkloadCluster.Spec.Cluster.Masters
		glog.Info("master and old master:", currentMasters, oldMasters)

		if !reflect.DeepEqual(currentMasters, oldMasters) {
			reviewResponse.Allowed = false
			msg += "the old master is different from the current master."
		}
	}

	if !reviewResponse.Allowed {
		reviewResponse.Result = &metav1.Status{Message: strings.TrimSpace(msg)}
	}
	return &reviewResponse
}
