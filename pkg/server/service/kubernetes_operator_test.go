package service

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	ecsv1 "git.yun.pingan.com/ecs/cluster-operator/pkg/apis/ecs/v1"
	"k8s.io/api/admission/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestAdmitClusterOperator(t *testing.T) {
	ar := v1beta1.AdmissionReview{
		TypeMeta: metav1.TypeMeta{
			Kind:       "AdmissionReview",
			APIVersion: "admission.k8s.io/v1beta1",
		},
		Request: &v1beta1.AdmissionRequest{
			Kind: metav1.GroupVersionKind{
				Group:   "ecs.yun.pingan.com",
				Version: "v1",
				Kind:    "WorkloadCluster",
			},
			Resource: metav1.GroupVersionResource{
				Group:    "ecs.yun.pingan.com",
				Version:  "v1",
				Resource: "workloadclusters",
			},
			Operation: "CREATE",
		},
	}

	testCases := []ecsv1.WorkloadCluster{
		{
			//ObjectMeta: ecsv1.ObjectMeta{
			//Labels: map[string]string{"Allow": "false"},
			//},
			Spec: ecsv1.WorkloadClusterSpec{
				Cluster: ecsv1.Cluster{
					Masters: []ecsv1.VmNode{
						{
							IP: "10.8.215.6",
						},
					},
					Workers: []ecsv1.VmNode{
						{
							IP: "10.8.215.6",
						},
					},
					RootPassword:  "",
					PrivateSSHKey: "",
					BearerToken:   "",
					KubeConfig:    "",
				},
			},
		},
	}

	// create
	for idx, test := range testCases {
		test.APIVersion = "ecs.yun.pingan.com/v1"
		test.Kind = "WorkloadCluster"
		test.Name = fmt.Sprintf("test-%v", idx)

		w, _ := json.Marshal(test)
		ar.Request.Object.Raw = w

		ap := AdmitClusterOperator(ar)
		if !reflect.DeepEqual(ap.Allowed, false) {
			t.Logf("except ,got %v,err is %v", ap.Allowed, ap.Result.Message)
		}
	}
	// update
	ar.Request.Operation = "UPDATE"
	for idx, test := range testCases {
		test.APIVersion = "ecs.yun.pingan.com/v1"
		test.Kind = "WorkloadCluster"
		test.Name = fmt.Sprintf("test-%v", idx)

		w, _ := json.Marshal(test)
		ar.Request.Object.Raw = w

		ap := AdmitClusterOperator(ar)
		if !reflect.DeepEqual(ap.Allowed, false) {
			t.Logf("except ,got %v,err is %v", ap.Allowed, ap.Result.Message)
		}
	}
}
