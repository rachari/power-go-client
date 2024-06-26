// Code generated by go-swagger; DO NOT EDIT.

package datacenters

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

// V1DatacentersGetReader is a Reader for the V1DatacentersGet structure.
type V1DatacentersGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1DatacentersGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1DatacentersGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewV1DatacentersGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewV1DatacentersGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV1DatacentersGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV1DatacentersGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/datacenters/{datacenter_region}] v1.datacenters.get", response, response.Code())
	}
}

// NewV1DatacentersGetOK creates a V1DatacentersGetOK with default headers values
func NewV1DatacentersGetOK() *V1DatacentersGetOK {
	return &V1DatacentersGetOK{}
}

/*
V1DatacentersGetOK describes a response with status code 200, with default header values.

OK
*/
type V1DatacentersGetOK struct {
	Payload *models.Datacenter
}

// IsSuccess returns true when this v1 datacenters get o k response has a 2xx status code
func (o *V1DatacentersGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this v1 datacenters get o k response has a 3xx status code
func (o *V1DatacentersGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 datacenters get o k response has a 4xx status code
func (o *V1DatacentersGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 datacenters get o k response has a 5xx status code
func (o *V1DatacentersGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 datacenters get o k response a status code equal to that given
func (o *V1DatacentersGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the v1 datacenters get o k response
func (o *V1DatacentersGetOK) Code() int {
	return 200
}

func (o *V1DatacentersGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/datacenters/{datacenter_region}][%d] v1DatacentersGetOK %s", 200, payload)
}

func (o *V1DatacentersGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/datacenters/{datacenter_region}][%d] v1DatacentersGetOK %s", 200, payload)
}

func (o *V1DatacentersGetOK) GetPayload() *models.Datacenter {
	return o.Payload
}

func (o *V1DatacentersGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Datacenter)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1DatacentersGetBadRequest creates a V1DatacentersGetBadRequest with default headers values
func NewV1DatacentersGetBadRequest() *V1DatacentersGetBadRequest {
	return &V1DatacentersGetBadRequest{}
}

/*
V1DatacentersGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type V1DatacentersGetBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 datacenters get bad request response has a 2xx status code
func (o *V1DatacentersGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 datacenters get bad request response has a 3xx status code
func (o *V1DatacentersGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 datacenters get bad request response has a 4xx status code
func (o *V1DatacentersGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 datacenters get bad request response has a 5xx status code
func (o *V1DatacentersGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 datacenters get bad request response a status code equal to that given
func (o *V1DatacentersGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the v1 datacenters get bad request response
func (o *V1DatacentersGetBadRequest) Code() int {
	return 400
}

func (o *V1DatacentersGetBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/datacenters/{datacenter_region}][%d] v1DatacentersGetBadRequest %s", 400, payload)
}

func (o *V1DatacentersGetBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/datacenters/{datacenter_region}][%d] v1DatacentersGetBadRequest %s", 400, payload)
}

func (o *V1DatacentersGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1DatacentersGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1DatacentersGetUnauthorized creates a V1DatacentersGetUnauthorized with default headers values
func NewV1DatacentersGetUnauthorized() *V1DatacentersGetUnauthorized {
	return &V1DatacentersGetUnauthorized{}
}

/*
V1DatacentersGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type V1DatacentersGetUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 datacenters get unauthorized response has a 2xx status code
func (o *V1DatacentersGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 datacenters get unauthorized response has a 3xx status code
func (o *V1DatacentersGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 datacenters get unauthorized response has a 4xx status code
func (o *V1DatacentersGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 datacenters get unauthorized response has a 5xx status code
func (o *V1DatacentersGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 datacenters get unauthorized response a status code equal to that given
func (o *V1DatacentersGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the v1 datacenters get unauthorized response
func (o *V1DatacentersGetUnauthorized) Code() int {
	return 401
}

func (o *V1DatacentersGetUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/datacenters/{datacenter_region}][%d] v1DatacentersGetUnauthorized %s", 401, payload)
}

func (o *V1DatacentersGetUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/datacenters/{datacenter_region}][%d] v1DatacentersGetUnauthorized %s", 401, payload)
}

func (o *V1DatacentersGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1DatacentersGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1DatacentersGetForbidden creates a V1DatacentersGetForbidden with default headers values
func NewV1DatacentersGetForbidden() *V1DatacentersGetForbidden {
	return &V1DatacentersGetForbidden{}
}

/*
V1DatacentersGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type V1DatacentersGetForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 datacenters get forbidden response has a 2xx status code
func (o *V1DatacentersGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 datacenters get forbidden response has a 3xx status code
func (o *V1DatacentersGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 datacenters get forbidden response has a 4xx status code
func (o *V1DatacentersGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 datacenters get forbidden response has a 5xx status code
func (o *V1DatacentersGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 datacenters get forbidden response a status code equal to that given
func (o *V1DatacentersGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the v1 datacenters get forbidden response
func (o *V1DatacentersGetForbidden) Code() int {
	return 403
}

func (o *V1DatacentersGetForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/datacenters/{datacenter_region}][%d] v1DatacentersGetForbidden %s", 403, payload)
}

func (o *V1DatacentersGetForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/datacenters/{datacenter_region}][%d] v1DatacentersGetForbidden %s", 403, payload)
}

func (o *V1DatacentersGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1DatacentersGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1DatacentersGetInternalServerError creates a V1DatacentersGetInternalServerError with default headers values
func NewV1DatacentersGetInternalServerError() *V1DatacentersGetInternalServerError {
	return &V1DatacentersGetInternalServerError{}
}

/*
V1DatacentersGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type V1DatacentersGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 datacenters get internal server error response has a 2xx status code
func (o *V1DatacentersGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 datacenters get internal server error response has a 3xx status code
func (o *V1DatacentersGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 datacenters get internal server error response has a 4xx status code
func (o *V1DatacentersGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 datacenters get internal server error response has a 5xx status code
func (o *V1DatacentersGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this v1 datacenters get internal server error response a status code equal to that given
func (o *V1DatacentersGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the v1 datacenters get internal server error response
func (o *V1DatacentersGetInternalServerError) Code() int {
	return 500
}

func (o *V1DatacentersGetInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/datacenters/{datacenter_region}][%d] v1DatacentersGetInternalServerError %s", 500, payload)
}

func (o *V1DatacentersGetInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/datacenters/{datacenter_region}][%d] v1DatacentersGetInternalServerError %s", 500, payload)
}

func (o *V1DatacentersGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1DatacentersGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
