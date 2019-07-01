package server

import "github.com/jarcoal/httpmock"

func TestServeClusterOperator(t *testing) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://127.0.0.1:8443/ecs/operator/cluster",
		httpmock.NewStringResponder(200, `{"Allowed":"true","Result":{"status":"","message":""}}`))

}
