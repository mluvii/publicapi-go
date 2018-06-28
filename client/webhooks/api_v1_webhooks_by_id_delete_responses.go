// Code generated by go-swagger; DO NOT EDIT.

package webhooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// APIV1WebhooksByIDDeleteReader is a Reader for the APIV1WebhooksByIDDelete structure.
type APIV1WebhooksByIDDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIV1WebhooksByIDDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewAPIV1WebhooksByIDDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAPIV1WebhooksByIDDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewAPIV1WebhooksByIDDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewAPIV1WebhooksByIDDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAPIV1WebhooksByIDDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAPIV1WebhooksByIDDeleteNoContent creates a APIV1WebhooksByIDDeleteNoContent with default headers values
func NewAPIV1WebhooksByIDDeleteNoContent() *APIV1WebhooksByIDDeleteNoContent {
	return &APIV1WebhooksByIDDeleteNoContent{}
}

/*APIV1WebhooksByIDDeleteNoContent handles this case with default header values.

Webhook deleted
*/
type APIV1WebhooksByIDDeleteNoContent struct {
}

func (o *APIV1WebhooksByIDDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/Webhooks/{id}][%d] apiV1WebhooksByIdDeleteNoContent ", 204)
}

func (o *APIV1WebhooksByIDDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAPIV1WebhooksByIDDeleteBadRequest creates a APIV1WebhooksByIDDeleteBadRequest with default headers values
func NewAPIV1WebhooksByIDDeleteBadRequest() *APIV1WebhooksByIDDeleteBadRequest {
	return &APIV1WebhooksByIDDeleteBadRequest{}
}

/*APIV1WebhooksByIDDeleteBadRequest handles this case with default header values.

Bad Request
*/
type APIV1WebhooksByIDDeleteBadRequest struct {
	Payload APIV1WebhooksByIDDeleteBadRequestBody
}

func (o *APIV1WebhooksByIDDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/Webhooks/{id}][%d] apiV1WebhooksByIdDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *APIV1WebhooksByIDDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIV1WebhooksByIDDeleteUnauthorized creates a APIV1WebhooksByIDDeleteUnauthorized with default headers values
func NewAPIV1WebhooksByIDDeleteUnauthorized() *APIV1WebhooksByIDDeleteUnauthorized {
	return &APIV1WebhooksByIDDeleteUnauthorized{}
}

/*APIV1WebhooksByIDDeleteUnauthorized handles this case with default header values.

Unauthorized
*/
type APIV1WebhooksByIDDeleteUnauthorized struct {
}

func (o *APIV1WebhooksByIDDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/Webhooks/{id}][%d] apiV1WebhooksByIdDeleteUnauthorized ", 401)
}

func (o *APIV1WebhooksByIDDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAPIV1WebhooksByIDDeleteForbidden creates a APIV1WebhooksByIDDeleteForbidden with default headers values
func NewAPIV1WebhooksByIDDeleteForbidden() *APIV1WebhooksByIDDeleteForbidden {
	return &APIV1WebhooksByIDDeleteForbidden{}
}

/*APIV1WebhooksByIDDeleteForbidden handles this case with default header values.

Forbidden
*/
type APIV1WebhooksByIDDeleteForbidden struct {
}

func (o *APIV1WebhooksByIDDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/Webhooks/{id}][%d] apiV1WebhooksByIdDeleteForbidden ", 403)
}

func (o *APIV1WebhooksByIDDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAPIV1WebhooksByIDDeleteNotFound creates a APIV1WebhooksByIDDeleteNotFound with default headers values
func NewAPIV1WebhooksByIDDeleteNotFound() *APIV1WebhooksByIDDeleteNotFound {
	return &APIV1WebhooksByIDDeleteNotFound{}
}

/*APIV1WebhooksByIDDeleteNotFound handles this case with default header values.

Not Found
*/
type APIV1WebhooksByIDDeleteNotFound struct {
}

func (o *APIV1WebhooksByIDDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/Webhooks/{id}][%d] apiV1WebhooksByIdDeleteNotFound ", 404)
}

func (o *APIV1WebhooksByIDDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*APIV1WebhooksByIDDeleteBadRequestBody API v1 webhooks by ID delete bad request body
swagger:model APIV1WebhooksByIDDeleteBadRequestBody
*/
type APIV1WebhooksByIDDeleteBadRequestBody map[string]string

// Validate validates this API v1 webhooks by ID delete bad request body
func (o APIV1WebhooksByIDDeleteBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
