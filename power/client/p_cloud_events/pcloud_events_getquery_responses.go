// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudEventsGetqueryReader is a Reader for the PcloudEventsGetquery structure.
type PcloudEventsGetqueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudEventsGetqueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudEventsGetqueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudEventsGetqueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudEventsGetqueryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudEventsGetqueryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudEventsGetqueryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudEventsGetqueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/events] pcloud.events.getquery", response, response.Code())
	}
}

// NewPcloudEventsGetqueryOK creates a PcloudEventsGetqueryOK with default headers values
func NewPcloudEventsGetqueryOK() *PcloudEventsGetqueryOK {
	return &PcloudEventsGetqueryOK{}
}

/*
PcloudEventsGetqueryOK describes a response with status code 200, with default header values.

OK
*/
type PcloudEventsGetqueryOK struct {
	Payload *models.Events
}

// IsSuccess returns true when this pcloud events getquery o k response has a 2xx status code
func (o *PcloudEventsGetqueryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud events getquery o k response has a 3xx status code
func (o *PcloudEventsGetqueryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud events getquery o k response has a 4xx status code
func (o *PcloudEventsGetqueryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud events getquery o k response has a 5xx status code
func (o *PcloudEventsGetqueryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud events getquery o k response a status code equal to that given
func (o *PcloudEventsGetqueryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud events getquery o k response
func (o *PcloudEventsGetqueryOK) Code() int {
	return 200
}

func (o *PcloudEventsGetqueryOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/events][%d] pcloudEventsGetqueryOK %s", 200, payload)
}

func (o *PcloudEventsGetqueryOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/events][%d] pcloudEventsGetqueryOK %s", 200, payload)
}

func (o *PcloudEventsGetqueryOK) GetPayload() *models.Events {
	return o.Payload
}

func (o *PcloudEventsGetqueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Events)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudEventsGetqueryBadRequest creates a PcloudEventsGetqueryBadRequest with default headers values
func NewPcloudEventsGetqueryBadRequest() *PcloudEventsGetqueryBadRequest {
	return &PcloudEventsGetqueryBadRequest{}
}

/*
PcloudEventsGetqueryBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudEventsGetqueryBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud events getquery bad request response has a 2xx status code
func (o *PcloudEventsGetqueryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud events getquery bad request response has a 3xx status code
func (o *PcloudEventsGetqueryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud events getquery bad request response has a 4xx status code
func (o *PcloudEventsGetqueryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud events getquery bad request response has a 5xx status code
func (o *PcloudEventsGetqueryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud events getquery bad request response a status code equal to that given
func (o *PcloudEventsGetqueryBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud events getquery bad request response
func (o *PcloudEventsGetqueryBadRequest) Code() int {
	return 400
}

func (o *PcloudEventsGetqueryBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/events][%d] pcloudEventsGetqueryBadRequest %s", 400, payload)
}

func (o *PcloudEventsGetqueryBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/events][%d] pcloudEventsGetqueryBadRequest %s", 400, payload)
}

func (o *PcloudEventsGetqueryBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudEventsGetqueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudEventsGetqueryUnauthorized creates a PcloudEventsGetqueryUnauthorized with default headers values
func NewPcloudEventsGetqueryUnauthorized() *PcloudEventsGetqueryUnauthorized {
	return &PcloudEventsGetqueryUnauthorized{}
}

/*
PcloudEventsGetqueryUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudEventsGetqueryUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud events getquery unauthorized response has a 2xx status code
func (o *PcloudEventsGetqueryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud events getquery unauthorized response has a 3xx status code
func (o *PcloudEventsGetqueryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud events getquery unauthorized response has a 4xx status code
func (o *PcloudEventsGetqueryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud events getquery unauthorized response has a 5xx status code
func (o *PcloudEventsGetqueryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud events getquery unauthorized response a status code equal to that given
func (o *PcloudEventsGetqueryUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud events getquery unauthorized response
func (o *PcloudEventsGetqueryUnauthorized) Code() int {
	return 401
}

func (o *PcloudEventsGetqueryUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/events][%d] pcloudEventsGetqueryUnauthorized %s", 401, payload)
}

func (o *PcloudEventsGetqueryUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/events][%d] pcloudEventsGetqueryUnauthorized %s", 401, payload)
}

func (o *PcloudEventsGetqueryUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudEventsGetqueryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudEventsGetqueryForbidden creates a PcloudEventsGetqueryForbidden with default headers values
func NewPcloudEventsGetqueryForbidden() *PcloudEventsGetqueryForbidden {
	return &PcloudEventsGetqueryForbidden{}
}

/*
PcloudEventsGetqueryForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudEventsGetqueryForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud events getquery forbidden response has a 2xx status code
func (o *PcloudEventsGetqueryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud events getquery forbidden response has a 3xx status code
func (o *PcloudEventsGetqueryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud events getquery forbidden response has a 4xx status code
func (o *PcloudEventsGetqueryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud events getquery forbidden response has a 5xx status code
func (o *PcloudEventsGetqueryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud events getquery forbidden response a status code equal to that given
func (o *PcloudEventsGetqueryForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud events getquery forbidden response
func (o *PcloudEventsGetqueryForbidden) Code() int {
	return 403
}

func (o *PcloudEventsGetqueryForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/events][%d] pcloudEventsGetqueryForbidden %s", 403, payload)
}

func (o *PcloudEventsGetqueryForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/events][%d] pcloudEventsGetqueryForbidden %s", 403, payload)
}

func (o *PcloudEventsGetqueryForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudEventsGetqueryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudEventsGetqueryNotFound creates a PcloudEventsGetqueryNotFound with default headers values
func NewPcloudEventsGetqueryNotFound() *PcloudEventsGetqueryNotFound {
	return &PcloudEventsGetqueryNotFound{}
}

/*
PcloudEventsGetqueryNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudEventsGetqueryNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud events getquery not found response has a 2xx status code
func (o *PcloudEventsGetqueryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud events getquery not found response has a 3xx status code
func (o *PcloudEventsGetqueryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud events getquery not found response has a 4xx status code
func (o *PcloudEventsGetqueryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud events getquery not found response has a 5xx status code
func (o *PcloudEventsGetqueryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud events getquery not found response a status code equal to that given
func (o *PcloudEventsGetqueryNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud events getquery not found response
func (o *PcloudEventsGetqueryNotFound) Code() int {
	return 404
}

func (o *PcloudEventsGetqueryNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/events][%d] pcloudEventsGetqueryNotFound %s", 404, payload)
}

func (o *PcloudEventsGetqueryNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/events][%d] pcloudEventsGetqueryNotFound %s", 404, payload)
}

func (o *PcloudEventsGetqueryNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudEventsGetqueryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudEventsGetqueryInternalServerError creates a PcloudEventsGetqueryInternalServerError with default headers values
func NewPcloudEventsGetqueryInternalServerError() *PcloudEventsGetqueryInternalServerError {
	return &PcloudEventsGetqueryInternalServerError{}
}

/*
PcloudEventsGetqueryInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudEventsGetqueryInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud events getquery internal server error response has a 2xx status code
func (o *PcloudEventsGetqueryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud events getquery internal server error response has a 3xx status code
func (o *PcloudEventsGetqueryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud events getquery internal server error response has a 4xx status code
func (o *PcloudEventsGetqueryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud events getquery internal server error response has a 5xx status code
func (o *PcloudEventsGetqueryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud events getquery internal server error response a status code equal to that given
func (o *PcloudEventsGetqueryInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud events getquery internal server error response
func (o *PcloudEventsGetqueryInternalServerError) Code() int {
	return 500
}

func (o *PcloudEventsGetqueryInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/events][%d] pcloudEventsGetqueryInternalServerError %s", 500, payload)
}

func (o *PcloudEventsGetqueryInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/events][%d] pcloudEventsGetqueryInternalServerError %s", 500, payload)
}

func (o *PcloudEventsGetqueryInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudEventsGetqueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
