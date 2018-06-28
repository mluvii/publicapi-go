// Code generated by go-swagger; DO NOT EDIT.

package webhooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewAPIV1WebhooksDeleteParams creates a new APIV1WebhooksDeleteParams object
// with the default values initialized.
func NewAPIV1WebhooksDeleteParams() *APIV1WebhooksDeleteParams {

	return &APIV1WebhooksDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIV1WebhooksDeleteParamsWithTimeout creates a new APIV1WebhooksDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIV1WebhooksDeleteParamsWithTimeout(timeout time.Duration) *APIV1WebhooksDeleteParams {

	return &APIV1WebhooksDeleteParams{

		timeout: timeout,
	}
}

// NewAPIV1WebhooksDeleteParamsWithContext creates a new APIV1WebhooksDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIV1WebhooksDeleteParamsWithContext(ctx context.Context) *APIV1WebhooksDeleteParams {

	return &APIV1WebhooksDeleteParams{

		Context: ctx,
	}
}

// NewAPIV1WebhooksDeleteParamsWithHTTPClient creates a new APIV1WebhooksDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIV1WebhooksDeleteParamsWithHTTPClient(client *http.Client) *APIV1WebhooksDeleteParams {

	return &APIV1WebhooksDeleteParams{
		HTTPClient: client,
	}
}

/*APIV1WebhooksDeleteParams contains all the parameters to send to the API endpoint
for the Api v1 webhooks delete operation typically these are written to a http.Request
*/
type APIV1WebhooksDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the Api v1 webhooks delete params
func (o *APIV1WebhooksDeleteParams) WithTimeout(timeout time.Duration) *APIV1WebhooksDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Api v1 webhooks delete params
func (o *APIV1WebhooksDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Api v1 webhooks delete params
func (o *APIV1WebhooksDeleteParams) WithContext(ctx context.Context) *APIV1WebhooksDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Api v1 webhooks delete params
func (o *APIV1WebhooksDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Api v1 webhooks delete params
func (o *APIV1WebhooksDeleteParams) WithHTTPClient(client *http.Client) *APIV1WebhooksDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Api v1 webhooks delete params
func (o *APIV1WebhooksDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *APIV1WebhooksDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}