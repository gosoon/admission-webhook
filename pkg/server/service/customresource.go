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

//const (
//customResourcePatch1 string = `[
//{ "op": "add", "path": "/data/mutation-stage-1", "value": "yes" }
//]`
//customResourcePatch2 string = `[
//{ "op": "add", "path": "/data/mutation-stage-2", "value": "yes" }
//]`
//)

//func MutateCustomResource(ar v1beta1.AdmissionReview) *v1beta1.AdmissionResponse {
//glog.V(2).Info("mutating custom resource")
//cr := struct {
//metav1.ObjectMeta
//Data map[string]string
//}{}

//raw := ar.Request.Object.Raw
//err := json.Unmarshal(raw, &cr)
//if err != nil {
//glog.Error(err)
//return server.ToAdmissionResponse(err)
//}

//reviewResponse := v1beta1.AdmissionResponse{}
//reviewResponse.Allowed = true

//if cr.Data["mutation-start"] == "yes" {
//reviewResponse.Patch = []byte(customResourcePatch1)
//}
//if cr.Data["mutation-stage-1"] == "yes" {
//reviewResponse.Patch = []byte(customResourcePatch2)
//}
//pt := v1beta1.PatchTypeJSONPatch
//reviewResponse.PatchType = &pt
//return &reviewResponse
//}

//func AdmitCustomResource(ar v1beta1.AdmissionReview) *v1beta1.AdmissionResponse {
//glog.V(2).Info("admitting custom resource")
//cr := struct {
//metav1.ObjectMeta
//Data map[string]string
//}{}

//var raw []byte
//if ar.Request.Operation == v1beta1.Delete {
//raw = ar.Request.OldObject.Raw
//} else {
//raw = ar.Request.Object.Raw
//}
//err := json.Unmarshal(raw, &cr)
//if err != nil {
//glog.Error(err)
//return server.ToAdmissionResponse(err)
//}

//reviewResponse := v1beta1.AdmissionResponse{}
//reviewResponse.Allowed = true
//for k, v := range cr.Data {
//if k == "webhook-e2e-test" && v == "webhook-disallow" &&
//(ar.Request.Operation == v1beta1.Create || ar.Request.Operation == v1beta1.Update) {
//reviewResponse.Allowed = false
//reviewResponse.Result = &metav1.Status{
//Reason: "the custom resource contains unwanted data",
//}
//}
//if k == "webhook-e2e-test" && v == "webhook-nondeletable" && ar.Request.Operation == v1beta1.Delete {
//reviewResponse.Allowed = false
//reviewResponse.Result = &metav1.Status{
//Reason: "the custom resource cannot be deleted because it contains unwanted key and value",
//}
//}
//}
//return &reviewResponse
//}
