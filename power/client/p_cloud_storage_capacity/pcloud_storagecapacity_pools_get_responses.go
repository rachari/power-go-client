// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_storage_capacity

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

// PcloudStoragecapacityPoolsGetReader is a Reader for the PcloudStoragecapacityPoolsGet structure.
type PcloudStoragecapacityPoolsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudStoragecapacityPoolsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudStoragecapacityPoolsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudStoragecapacityPoolsGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudStoragecapacityPoolsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudStoragecapacityPoolsGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudStoragecapacityPoolsGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudStoragecapacityPoolsGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/storage-capacity/storage-pools/{storage_pool_name}] pcloud.storagecapacity.pools.get", response, response.Code())
	}
}

// NewPcloudStoragecapacityPoolsGetOK creates a PcloudStoragecapacityPoolsGetOK with default headers values
func NewPcloudStoragecapacityPoolsGetOK() *PcloudStoragecapacityPoolsGetOK {
	return &PcloudStoragecapacityPoolsGetOK{}
}

/*
PcloudStoragecapacityPoolsGetOK describes a response with status code 200, with default header values.

OK
*/
type PcloudStoragecapacityPoolsGetOK struct {
	Payload *models.StoragePoolCapacity
}

// IsSuccess returns true when this pcloud storagecapacity pools get o k response has a 2xx status code
func (o *PcloudStoragecapacityPoolsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud storagecapacity pools get o k response has a 3xx status code
func (o *PcloudStoragecapacityPoolsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud storagecapacity pools get o k response has a 4xx status code
func (o *PcloudStoragecapacityPoolsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud storagecapacity pools get o k response has a 5xx status code
func (o *PcloudStoragecapacityPoolsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud storagecapacity pools get o k response a status code equal to that given
func (o *PcloudStoragecapacityPoolsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud storagecapacity pools get o k response
func (o *PcloudStoragecapacityPoolsGetOK) Code() int {
	return 200
}

func (o *PcloudStoragecapacityPoolsGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/storage-capacity/storage-pools/{storage_pool_name}][%d] pcloudStoragecapacityPoolsGetOK %s", 200, payload)
}

func (o *PcloudStoragecapacityPoolsGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/storage-capacity/storage-pools/{storage_pool_name}][%d] pcloudStoragecapacityPoolsGetOK %s", 200, payload)
}

func (o *PcloudStoragecapacityPoolsGetOK) GetPayload() *models.StoragePoolCapacity {
	return o.Payload
}

func (o *PcloudStoragecapacityPoolsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StoragePoolCapacity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudStoragecapacityPoolsGetBadRequest creates a PcloudStoragecapacityPoolsGetBadRequest with default headers values
func NewPcloudStoragecapacityPoolsGetBadRequest() *PcloudStoragecapacityPoolsGetBadRequest {
	return &PcloudStoragecapacityPoolsGetBadRequest{}
}

/*
PcloudStoragecapacityPoolsGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudStoragecapacityPoolsGetBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud storagecapacity pools get bad request response has a 2xx status code
func (o *PcloudStoragecapacityPoolsGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud storagecapacity pools get bad request response has a 3xx status code
func (o *PcloudStoragecapacityPoolsGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud storagecapacity pools get bad request response has a 4xx status code
func (o *PcloudStoragecapacityPoolsGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud storagecapacity pools get bad request response has a 5xx status code
func (o *PcloudStoragecapacityPoolsGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud storagecapacity pools get bad request response a status code equal to that given
func (o *PcloudStoragecapacityPoolsGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud storagecapacity pools get bad request response
func (o *PcloudStoragecapacityPoolsGetBadRequest) Code() int {
	return 400
}

func (o *PcloudStoragecapacityPoolsGetBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/storage-capacity/storage-pools/{storage_pool_name}][%d] pcloudStoragecapacityPoolsGetBadRequest %s", 400, payload)
}

func (o *PcloudStoragecapacityPoolsGetBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/storage-capacity/storage-pools/{storage_pool_name}][%d] pcloudStoragecapacityPoolsGetBadRequest %s", 400, payload)
}

func (o *PcloudStoragecapacityPoolsGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudStoragecapacityPoolsGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudStoragecapacityPoolsGetUnauthorized creates a PcloudStoragecapacityPoolsGetUnauthorized with default headers values
func NewPcloudStoragecapacityPoolsGetUnauthorized() *PcloudStoragecapacityPoolsGetUnauthorized {
	return &PcloudStoragecapacityPoolsGetUnauthorized{}
}

/*
PcloudStoragecapacityPoolsGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudStoragecapacityPoolsGetUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud storagecapacity pools get unauthorized response has a 2xx status code
func (o *PcloudStoragecapacityPoolsGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud storagecapacity pools get unauthorized response has a 3xx status code
func (o *PcloudStoragecapacityPoolsGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud storagecapacity pools get unauthorized response has a 4xx status code
func (o *PcloudStoragecapacityPoolsGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud storagecapacity pools get unauthorized response has a 5xx status code
func (o *PcloudStoragecapacityPoolsGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud storagecapacity pools get unauthorized response a status code equal to that given
func (o *PcloudStoragecapacityPoolsGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud storagecapacity pools get unauthorized response
func (o *PcloudStoragecapacityPoolsGetUnauthorized) Code() int {
	return 401
}

func (o *PcloudStoragecapacityPoolsGetUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/storage-capacity/storage-pools/{storage_pool_name}][%d] pcloudStoragecapacityPoolsGetUnauthorized %s", 401, payload)
}

func (o *PcloudStoragecapacityPoolsGetUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/storage-capacity/storage-pools/{storage_pool_name}][%d] pcloudStoragecapacityPoolsGetUnauthorized %s", 401, payload)
}

func (o *PcloudStoragecapacityPoolsGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudStoragecapacityPoolsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudStoragecapacityPoolsGetForbidden creates a PcloudStoragecapacityPoolsGetForbidden with default headers values
func NewPcloudStoragecapacityPoolsGetForbidden() *PcloudStoragecapacityPoolsGetForbidden {
	return &PcloudStoragecapacityPoolsGetForbidden{}
}

/*
PcloudStoragecapacityPoolsGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudStoragecapacityPoolsGetForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud storagecapacity pools get forbidden response has a 2xx status code
func (o *PcloudStoragecapacityPoolsGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud storagecapacity pools get forbidden response has a 3xx status code
func (o *PcloudStoragecapacityPoolsGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud storagecapacity pools get forbidden response has a 4xx status code
func (o *PcloudStoragecapacityPoolsGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud storagecapacity pools get forbidden response has a 5xx status code
func (o *PcloudStoragecapacityPoolsGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud storagecapacity pools get forbidden response a status code equal to that given
func (o *PcloudStoragecapacityPoolsGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud storagecapacity pools get forbidden response
func (o *PcloudStoragecapacityPoolsGetForbidden) Code() int {
	return 403
}

func (o *PcloudStoragecapacityPoolsGetForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/storage-capacity/storage-pools/{storage_pool_name}][%d] pcloudStoragecapacityPoolsGetForbidden %s", 403, payload)
}

func (o *PcloudStoragecapacityPoolsGetForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/storage-capacity/storage-pools/{storage_pool_name}][%d] pcloudStoragecapacityPoolsGetForbidden %s", 403, payload)
}

func (o *PcloudStoragecapacityPoolsGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudStoragecapacityPoolsGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudStoragecapacityPoolsGetNotFound creates a PcloudStoragecapacityPoolsGetNotFound with default headers values
func NewPcloudStoragecapacityPoolsGetNotFound() *PcloudStoragecapacityPoolsGetNotFound {
	return &PcloudStoragecapacityPoolsGetNotFound{}
}

/*
PcloudStoragecapacityPoolsGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudStoragecapacityPoolsGetNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud storagecapacity pools get not found response has a 2xx status code
func (o *PcloudStoragecapacityPoolsGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud storagecapacity pools get not found response has a 3xx status code
func (o *PcloudStoragecapacityPoolsGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud storagecapacity pools get not found response has a 4xx status code
func (o *PcloudStoragecapacityPoolsGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud storagecapacity pools get not found response has a 5xx status code
func (o *PcloudStoragecapacityPoolsGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud storagecapacity pools get not found response a status code equal to that given
func (o *PcloudStoragecapacityPoolsGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud storagecapacity pools get not found response
func (o *PcloudStoragecapacityPoolsGetNotFound) Code() int {
	return 404
}

func (o *PcloudStoragecapacityPoolsGetNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/storage-capacity/storage-pools/{storage_pool_name}][%d] pcloudStoragecapacityPoolsGetNotFound %s", 404, payload)
}

func (o *PcloudStoragecapacityPoolsGetNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/storage-capacity/storage-pools/{storage_pool_name}][%d] pcloudStoragecapacityPoolsGetNotFound %s", 404, payload)
}

func (o *PcloudStoragecapacityPoolsGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudStoragecapacityPoolsGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudStoragecapacityPoolsGetInternalServerError creates a PcloudStoragecapacityPoolsGetInternalServerError with default headers values
func NewPcloudStoragecapacityPoolsGetInternalServerError() *PcloudStoragecapacityPoolsGetInternalServerError {
	return &PcloudStoragecapacityPoolsGetInternalServerError{}
}

/*
PcloudStoragecapacityPoolsGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudStoragecapacityPoolsGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud storagecapacity pools get internal server error response has a 2xx status code
func (o *PcloudStoragecapacityPoolsGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud storagecapacity pools get internal server error response has a 3xx status code
func (o *PcloudStoragecapacityPoolsGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud storagecapacity pools get internal server error response has a 4xx status code
func (o *PcloudStoragecapacityPoolsGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud storagecapacity pools get internal server error response has a 5xx status code
func (o *PcloudStoragecapacityPoolsGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud storagecapacity pools get internal server error response a status code equal to that given
func (o *PcloudStoragecapacityPoolsGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud storagecapacity pools get internal server error response
func (o *PcloudStoragecapacityPoolsGetInternalServerError) Code() int {
	return 500
}

func (o *PcloudStoragecapacityPoolsGetInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/storage-capacity/storage-pools/{storage_pool_name}][%d] pcloudStoragecapacityPoolsGetInternalServerError %s", 500, payload)
}

func (o *PcloudStoragecapacityPoolsGetInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/storage-capacity/storage-pools/{storage_pool_name}][%d] pcloudStoragecapacityPoolsGetInternalServerError %s", 500, payload)
}

func (o *PcloudStoragecapacityPoolsGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudStoragecapacityPoolsGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
