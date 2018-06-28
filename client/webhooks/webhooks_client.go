// Code generated by go-swagger; DO NOT EDIT.

package webhooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new webhooks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for webhooks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
APIV1WebhooksByIDDelete deletes webhook by id
*/
func (a *Client) APIV1WebhooksByIDDelete(params *APIV1WebhooksByIDDeleteParams) (*APIV1WebhooksByIDDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAPIV1WebhooksByIDDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ApiV1WebhooksByIdDelete",
		Method:             "DELETE",
		PathPattern:        "/api/v1/Webhooks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &APIV1WebhooksByIDDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*APIV1WebhooksByIDDeleteNoContent), nil

}

/*
APIV1WebhooksByIDGet gets webhook
*/
func (a *Client) APIV1WebhooksByIDGet(params *APIV1WebhooksByIDGetParams) (*APIV1WebhooksByIDGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAPIV1WebhooksByIDGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ApiV1WebhooksByIdGet",
		Method:             "GET",
		PathPattern:        "/api/v1/Webhooks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &APIV1WebhooksByIDGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*APIV1WebhooksByIDGetOK), nil

}

/*
APIV1WebhooksByIDPut updates webhook
*/
func (a *Client) APIV1WebhooksByIDPut(params *APIV1WebhooksByIDPutParams) (*APIV1WebhooksByIDPutNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAPIV1WebhooksByIDPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ApiV1WebhooksByIdPut",
		Method:             "PUT",
		PathPattern:        "/api/v1/Webhooks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &APIV1WebhooksByIDPutReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*APIV1WebhooksByIDPutNoContent), nil

}

/*
APIV1WebhooksDelete deletes all webhooks
*/
func (a *Client) APIV1WebhooksDelete(params *APIV1WebhooksDeleteParams) (*APIV1WebhooksDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAPIV1WebhooksDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ApiV1WebhooksDelete",
		Method:             "DELETE",
		PathPattern:        "/api/v1/Webhooks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &APIV1WebhooksDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*APIV1WebhooksDeleteNoContent), nil

}

/*
APIV1WebhooksPost registers webhook for given event
*/
func (a *Client) APIV1WebhooksPost(params *APIV1WebhooksPostParams) (*APIV1WebhooksPostCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAPIV1WebhooksPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ApiV1WebhooksPost",
		Method:             "POST",
		PathPattern:        "/api/v1/Webhooks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &APIV1WebhooksPostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*APIV1WebhooksPostCreated), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
