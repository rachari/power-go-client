// Code generated by go-swagger; DO NOT EDIT.

package storage_types

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

// ServiceBrokerStoragetypesGetReader is a Reader for the ServiceBrokerStoragetypesGet structure.
type ServiceBrokerStoragetypesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceBrokerStoragetypesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceBrokerStoragetypesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewServiceBrokerStoragetypesGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewServiceBrokerStoragetypesGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewServiceBrokerStoragetypesGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewServiceBrokerStoragetypesGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewServiceBrokerStoragetypesGetUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewServiceBrokerStoragetypesGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /broker/v1/storage-types] serviceBroker.storagetypes.get", response, response.Code())
	}
}

// NewServiceBrokerStoragetypesGetOK creates a ServiceBrokerStoragetypesGetOK with default headers values
func NewServiceBrokerStoragetypesGetOK() *ServiceBrokerStoragetypesGetOK {
	return &ServiceBrokerStoragetypesGetOK{}
}

/*
ServiceBrokerStoragetypesGetOK describes a response with status code 200, with default header values.

OK
*/
type ServiceBrokerStoragetypesGetOK struct {
	Payload models.StorageTypes
}

// IsSuccess returns true when this service broker storagetypes get o k response has a 2xx status code
func (o *ServiceBrokerStoragetypesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service broker storagetypes get o k response has a 3xx status code
func (o *ServiceBrokerStoragetypesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker storagetypes get o k response has a 4xx status code
func (o *ServiceBrokerStoragetypesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service broker storagetypes get o k response has a 5xx status code
func (o *ServiceBrokerStoragetypesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker storagetypes get o k response a status code equal to that given
func (o *ServiceBrokerStoragetypesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the service broker storagetypes get o k response
func (o *ServiceBrokerStoragetypesGetOK) Code() int {
	return 200
}

func (o *ServiceBrokerStoragetypesGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /broker/v1/storage-types][%d] serviceBrokerStoragetypesGetOK %s", 200, payload)
}

func (o *ServiceBrokerStoragetypesGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /broker/v1/storage-types][%d] serviceBrokerStoragetypesGetOK %s", 200, payload)
}

func (o *ServiceBrokerStoragetypesGetOK) GetPayload() models.StorageTypes {
	return o.Payload
}

func (o *ServiceBrokerStoragetypesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerStoragetypesGetBadRequest creates a ServiceBrokerStoragetypesGetBadRequest with default headers values
func NewServiceBrokerStoragetypesGetBadRequest() *ServiceBrokerStoragetypesGetBadRequest {
	return &ServiceBrokerStoragetypesGetBadRequest{}
}

/*
ServiceBrokerStoragetypesGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ServiceBrokerStoragetypesGetBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker storagetypes get bad request response has a 2xx status code
func (o *ServiceBrokerStoragetypesGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker storagetypes get bad request response has a 3xx status code
func (o *ServiceBrokerStoragetypesGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker storagetypes get bad request response has a 4xx status code
func (o *ServiceBrokerStoragetypesGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker storagetypes get bad request response has a 5xx status code
func (o *ServiceBrokerStoragetypesGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker storagetypes get bad request response a status code equal to that given
func (o *ServiceBrokerStoragetypesGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the service broker storagetypes get bad request response
func (o *ServiceBrokerStoragetypesGetBadRequest) Code() int {
	return 400
}

func (o *ServiceBrokerStoragetypesGetBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /broker/v1/storage-types][%d] serviceBrokerStoragetypesGetBadRequest %s", 400, payload)
}

func (o *ServiceBrokerStoragetypesGetBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /broker/v1/storage-types][%d] serviceBrokerStoragetypesGetBadRequest %s", 400, payload)
}

func (o *ServiceBrokerStoragetypesGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerStoragetypesGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerStoragetypesGetUnauthorized creates a ServiceBrokerStoragetypesGetUnauthorized with default headers values
func NewServiceBrokerStoragetypesGetUnauthorized() *ServiceBrokerStoragetypesGetUnauthorized {
	return &ServiceBrokerStoragetypesGetUnauthorized{}
}

/*
ServiceBrokerStoragetypesGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ServiceBrokerStoragetypesGetUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker storagetypes get unauthorized response has a 2xx status code
func (o *ServiceBrokerStoragetypesGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker storagetypes get unauthorized response has a 3xx status code
func (o *ServiceBrokerStoragetypesGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker storagetypes get unauthorized response has a 4xx status code
func (o *ServiceBrokerStoragetypesGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker storagetypes get unauthorized response has a 5xx status code
func (o *ServiceBrokerStoragetypesGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker storagetypes get unauthorized response a status code equal to that given
func (o *ServiceBrokerStoragetypesGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the service broker storagetypes get unauthorized response
func (o *ServiceBrokerStoragetypesGetUnauthorized) Code() int {
	return 401
}

func (o *ServiceBrokerStoragetypesGetUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /broker/v1/storage-types][%d] serviceBrokerStoragetypesGetUnauthorized %s", 401, payload)
}

func (o *ServiceBrokerStoragetypesGetUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /broker/v1/storage-types][%d] serviceBrokerStoragetypesGetUnauthorized %s", 401, payload)
}

func (o *ServiceBrokerStoragetypesGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerStoragetypesGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerStoragetypesGetForbidden creates a ServiceBrokerStoragetypesGetForbidden with default headers values
func NewServiceBrokerStoragetypesGetForbidden() *ServiceBrokerStoragetypesGetForbidden {
	return &ServiceBrokerStoragetypesGetForbidden{}
}

/*
ServiceBrokerStoragetypesGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ServiceBrokerStoragetypesGetForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker storagetypes get forbidden response has a 2xx status code
func (o *ServiceBrokerStoragetypesGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker storagetypes get forbidden response has a 3xx status code
func (o *ServiceBrokerStoragetypesGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker storagetypes get forbidden response has a 4xx status code
func (o *ServiceBrokerStoragetypesGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker storagetypes get forbidden response has a 5xx status code
func (o *ServiceBrokerStoragetypesGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker storagetypes get forbidden response a status code equal to that given
func (o *ServiceBrokerStoragetypesGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the service broker storagetypes get forbidden response
func (o *ServiceBrokerStoragetypesGetForbidden) Code() int {
	return 403
}

func (o *ServiceBrokerStoragetypesGetForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /broker/v1/storage-types][%d] serviceBrokerStoragetypesGetForbidden %s", 403, payload)
}

func (o *ServiceBrokerStoragetypesGetForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /broker/v1/storage-types][%d] serviceBrokerStoragetypesGetForbidden %s", 403, payload)
}

func (o *ServiceBrokerStoragetypesGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerStoragetypesGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerStoragetypesGetNotFound creates a ServiceBrokerStoragetypesGetNotFound with default headers values
func NewServiceBrokerStoragetypesGetNotFound() *ServiceBrokerStoragetypesGetNotFound {
	return &ServiceBrokerStoragetypesGetNotFound{}
}

/*
ServiceBrokerStoragetypesGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ServiceBrokerStoragetypesGetNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker storagetypes get not found response has a 2xx status code
func (o *ServiceBrokerStoragetypesGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker storagetypes get not found response has a 3xx status code
func (o *ServiceBrokerStoragetypesGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker storagetypes get not found response has a 4xx status code
func (o *ServiceBrokerStoragetypesGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker storagetypes get not found response has a 5xx status code
func (o *ServiceBrokerStoragetypesGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker storagetypes get not found response a status code equal to that given
func (o *ServiceBrokerStoragetypesGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the service broker storagetypes get not found response
func (o *ServiceBrokerStoragetypesGetNotFound) Code() int {
	return 404
}

func (o *ServiceBrokerStoragetypesGetNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /broker/v1/storage-types][%d] serviceBrokerStoragetypesGetNotFound %s", 404, payload)
}

func (o *ServiceBrokerStoragetypesGetNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /broker/v1/storage-types][%d] serviceBrokerStoragetypesGetNotFound %s", 404, payload)
}

func (o *ServiceBrokerStoragetypesGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerStoragetypesGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerStoragetypesGetUnprocessableEntity creates a ServiceBrokerStoragetypesGetUnprocessableEntity with default headers values
func NewServiceBrokerStoragetypesGetUnprocessableEntity() *ServiceBrokerStoragetypesGetUnprocessableEntity {
	return &ServiceBrokerStoragetypesGetUnprocessableEntity{}
}

/*
ServiceBrokerStoragetypesGetUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type ServiceBrokerStoragetypesGetUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker storagetypes get unprocessable entity response has a 2xx status code
func (o *ServiceBrokerStoragetypesGetUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker storagetypes get unprocessable entity response has a 3xx status code
func (o *ServiceBrokerStoragetypesGetUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker storagetypes get unprocessable entity response has a 4xx status code
func (o *ServiceBrokerStoragetypesGetUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker storagetypes get unprocessable entity response has a 5xx status code
func (o *ServiceBrokerStoragetypesGetUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker storagetypes get unprocessable entity response a status code equal to that given
func (o *ServiceBrokerStoragetypesGetUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the service broker storagetypes get unprocessable entity response
func (o *ServiceBrokerStoragetypesGetUnprocessableEntity) Code() int {
	return 422
}

func (o *ServiceBrokerStoragetypesGetUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /broker/v1/storage-types][%d] serviceBrokerStoragetypesGetUnprocessableEntity %s", 422, payload)
}

func (o *ServiceBrokerStoragetypesGetUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /broker/v1/storage-types][%d] serviceBrokerStoragetypesGetUnprocessableEntity %s", 422, payload)
}

func (o *ServiceBrokerStoragetypesGetUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerStoragetypesGetUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerStoragetypesGetInternalServerError creates a ServiceBrokerStoragetypesGetInternalServerError with default headers values
func NewServiceBrokerStoragetypesGetInternalServerError() *ServiceBrokerStoragetypesGetInternalServerError {
	return &ServiceBrokerStoragetypesGetInternalServerError{}
}

/*
ServiceBrokerStoragetypesGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ServiceBrokerStoragetypesGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker storagetypes get internal server error response has a 2xx status code
func (o *ServiceBrokerStoragetypesGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker storagetypes get internal server error response has a 3xx status code
func (o *ServiceBrokerStoragetypesGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker storagetypes get internal server error response has a 4xx status code
func (o *ServiceBrokerStoragetypesGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this service broker storagetypes get internal server error response has a 5xx status code
func (o *ServiceBrokerStoragetypesGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this service broker storagetypes get internal server error response a status code equal to that given
func (o *ServiceBrokerStoragetypesGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the service broker storagetypes get internal server error response
func (o *ServiceBrokerStoragetypesGetInternalServerError) Code() int {
	return 500
}

func (o *ServiceBrokerStoragetypesGetInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /broker/v1/storage-types][%d] serviceBrokerStoragetypesGetInternalServerError %s", 500, payload)
}

func (o *ServiceBrokerStoragetypesGetInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /broker/v1/storage-types][%d] serviceBrokerStoragetypesGetInternalServerError %s", 500, payload)
}

func (o *ServiceBrokerStoragetypesGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerStoragetypesGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
