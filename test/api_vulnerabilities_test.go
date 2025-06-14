/*
FOSSA API

Testing VulnerabilitiesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package fossa

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/RobertKielty/fossa-go"
)

func Test_fossa_VulnerabilitiesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test VulnerabilitiesAPIService GetRemediationGuidance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var vulnId string
		var revisionId string

		resp, httpRes, err := apiClient.VulnerabilitiesAPI.GetRemediationGuidance(context.Background(), vulnId, revisionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
