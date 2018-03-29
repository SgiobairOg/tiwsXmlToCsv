package main_test

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
	main "github.com/sgiobairog/tiwsXmlToCsv"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request events.APIGatewayProxyRequest
		expect  string
		err     error
	}{
		{
			// Test for correct response
			request: events.APIGatewayProxyRequest{Body: "Paul"},
			expect:  "Hello Paul",
			err:     nil,
		},
		{
			// Test for ErrNameNotProvided
			request: events.APIGatewayProxyRequest{Body: ""},
			expect:  "",
			err:     main.ErrNameNotProvided,
		},
	}

	for _, test := range tests {
		response, err := main.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect, response.Body)
	}
}
