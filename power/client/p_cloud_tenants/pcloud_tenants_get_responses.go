// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_tenants

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

// PcloudTenantsGetReader is a Reader for the PcloudTenantsGet structure.
type PcloudTenantsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudTenantsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudTenantsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudTenantsGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudTenantsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudTenantsGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudTenantsGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudTenantsGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /pcloud/v1/tenants/{tenant_id}] pcloud.tenants.get", response, response.Code())
	}
}

// NewPcloudTenantsGetOK creates a PcloudTenantsGetOK with default headers values
func NewPcloudTenantsGetOK() *PcloudTenantsGetOK {
	return &PcloudTenantsGetOK{}
}

/*
PcloudTenantsGetOK describes a response with status code 200, with default header values.

OK
*/
type PcloudTenantsGetOK struct {
	Payload *models.Tenant
}

// IsSuccess returns true when this pcloud tenants get o k response has a 2xx status code
func (o *PcloudTenantsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud tenants get o k response has a 3xx status code
func (o *PcloudTenantsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud tenants get o k response has a 4xx status code
func (o *PcloudTenantsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud tenants get o k response has a 5xx status code
func (o *PcloudTenantsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud tenants get o k response a status code equal to that given
func (o *PcloudTenantsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud tenants get o k response
func (o *PcloudTenantsGetOK) Code() int {
	return 200
}

func (o *PcloudTenantsGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/tenants/{tenant_id}][%d] pcloudTenantsGetOK %s", 200, payload)
}

func (o *PcloudTenantsGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/tenants/{tenant_id}][%d] pcloudTenantsGetOK %s", 200, payload)
}

func (o *PcloudTenantsGetOK) GetPayload() *models.Tenant {
	return o.Payload
}

func (o *PcloudTenantsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tenant)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTenantsGetBadRequest creates a PcloudTenantsGetBadRequest with default headers values
func NewPcloudTenantsGetBadRequest() *PcloudTenantsGetBadRequest {
	return &PcloudTenantsGetBadRequest{}
}

/*
PcloudTenantsGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudTenantsGetBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud tenants get bad request response has a 2xx status code
func (o *PcloudTenantsGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud tenants get bad request response has a 3xx status code
func (o *PcloudTenantsGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud tenants get bad request response has a 4xx status code
func (o *PcloudTenantsGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud tenants get bad request response has a 5xx status code
func (o *PcloudTenantsGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud tenants get bad request response a status code equal to that given
func (o *PcloudTenantsGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud tenants get bad request response
func (o *PcloudTenantsGetBadRequest) Code() int {
	return 400
}

func (o *PcloudTenantsGetBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/tenants/{tenant_id}][%d] pcloudTenantsGetBadRequest %s", 400, payload)
}

func (o *PcloudTenantsGetBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/tenants/{tenant_id}][%d] pcloudTenantsGetBadRequest %s", 400, payload)
}

func (o *PcloudTenantsGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudTenantsGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTenantsGetUnauthorized creates a PcloudTenantsGetUnauthorized with default headers values
func NewPcloudTenantsGetUnauthorized() *PcloudTenantsGetUnauthorized {
	return &PcloudTenantsGetUnauthorized{}
}

/*
PcloudTenantsGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudTenantsGetUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud tenants get unauthorized response has a 2xx status code
func (o *PcloudTenantsGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud tenants get unauthorized response has a 3xx status code
func (o *PcloudTenantsGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud tenants get unauthorized response has a 4xx status code
func (o *PcloudTenantsGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud tenants get unauthorized response has a 5xx status code
func (o *PcloudTenantsGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud tenants get unauthorized response a status code equal to that given
func (o *PcloudTenantsGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud tenants get unauthorized response
func (o *PcloudTenantsGetUnauthorized) Code() int {
	return 401
}

func (o *PcloudTenantsGetUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/tenants/{tenant_id}][%d] pcloudTenantsGetUnauthorized %s", 401, payload)
}

func (o *PcloudTenantsGetUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/tenants/{tenant_id}][%d] pcloudTenantsGetUnauthorized %s", 401, payload)
}

func (o *PcloudTenantsGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudTenantsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTenantsGetForbidden creates a PcloudTenantsGetForbidden with default headers values
func NewPcloudTenantsGetForbidden() *PcloudTenantsGetForbidden {
	return &PcloudTenantsGetForbidden{}
}

/*
PcloudTenantsGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudTenantsGetForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud tenants get forbidden response has a 2xx status code
func (o *PcloudTenantsGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud tenants get forbidden response has a 3xx status code
func (o *PcloudTenantsGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud tenants get forbidden response has a 4xx status code
func (o *PcloudTenantsGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud tenants get forbidden response has a 5xx status code
func (o *PcloudTenantsGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud tenants get forbidden response a status code equal to that given
func (o *PcloudTenantsGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud tenants get forbidden response
func (o *PcloudTenantsGetForbidden) Code() int {
	return 403
}

func (o *PcloudTenantsGetForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/tenants/{tenant_id}][%d] pcloudTenantsGetForbidden %s", 403, payload)
}

func (o *PcloudTenantsGetForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/tenants/{tenant_id}][%d] pcloudTenantsGetForbidden %s", 403, payload)
}

func (o *PcloudTenantsGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudTenantsGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTenantsGetNotFound creates a PcloudTenantsGetNotFound with default headers values
func NewPcloudTenantsGetNotFound() *PcloudTenantsGetNotFound {
	return &PcloudTenantsGetNotFound{}
}

/*
PcloudTenantsGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudTenantsGetNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud tenants get not found response has a 2xx status code
func (o *PcloudTenantsGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud tenants get not found response has a 3xx status code
func (o *PcloudTenantsGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud tenants get not found response has a 4xx status code
func (o *PcloudTenantsGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud tenants get not found response has a 5xx status code
func (o *PcloudTenantsGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud tenants get not found response a status code equal to that given
func (o *PcloudTenantsGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud tenants get not found response
func (o *PcloudTenantsGetNotFound) Code() int {
	return 404
}

func (o *PcloudTenantsGetNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/tenants/{tenant_id}][%d] pcloudTenantsGetNotFound %s", 404, payload)
}

func (o *PcloudTenantsGetNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/tenants/{tenant_id}][%d] pcloudTenantsGetNotFound %s", 404, payload)
}

func (o *PcloudTenantsGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudTenantsGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTenantsGetInternalServerError creates a PcloudTenantsGetInternalServerError with default headers values
func NewPcloudTenantsGetInternalServerError() *PcloudTenantsGetInternalServerError {
	return &PcloudTenantsGetInternalServerError{}
}

/*
PcloudTenantsGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudTenantsGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud tenants get internal server error response has a 2xx status code
func (o *PcloudTenantsGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud tenants get internal server error response has a 3xx status code
func (o *PcloudTenantsGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud tenants get internal server error response has a 4xx status code
func (o *PcloudTenantsGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud tenants get internal server error response has a 5xx status code
func (o *PcloudTenantsGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud tenants get internal server error response a status code equal to that given
func (o *PcloudTenantsGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud tenants get internal server error response
func (o *PcloudTenantsGetInternalServerError) Code() int {
	return 500
}

func (o *PcloudTenantsGetInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/tenants/{tenant_id}][%d] pcloudTenantsGetInternalServerError %s", 500, payload)
}

func (o *PcloudTenantsGetInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/tenants/{tenant_id}][%d] pcloudTenantsGetInternalServerError %s", 500, payload)
}

func (o *PcloudTenantsGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudTenantsGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
