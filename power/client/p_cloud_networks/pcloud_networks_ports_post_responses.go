// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_networks

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

// PcloudNetworksPortsPostReader is a Reader for the PcloudNetworksPortsPost structure.
type PcloudNetworksPortsPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudNetworksPortsPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPcloudNetworksPortsPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudNetworksPortsPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudNetworksPortsPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudNetworksPortsPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudNetworksPortsPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPcloudNetworksPortsPostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudNetworksPortsPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudNetworksPortsPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports] pcloud.networks.ports.post", response, response.Code())
	}
}

// NewPcloudNetworksPortsPostCreated creates a PcloudNetworksPortsPostCreated with default headers values
func NewPcloudNetworksPortsPostCreated() *PcloudNetworksPortsPostCreated {
	return &PcloudNetworksPortsPostCreated{}
}

/*
PcloudNetworksPortsPostCreated describes a response with status code 201, with default header values.

Created
*/
type PcloudNetworksPortsPostCreated struct {
	Payload *models.NetworkPort
}

// IsSuccess returns true when this pcloud networks ports post created response has a 2xx status code
func (o *PcloudNetworksPortsPostCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud networks ports post created response has a 3xx status code
func (o *PcloudNetworksPortsPostCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks ports post created response has a 4xx status code
func (o *PcloudNetworksPortsPostCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud networks ports post created response has a 5xx status code
func (o *PcloudNetworksPortsPostCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks ports post created response a status code equal to that given
func (o *PcloudNetworksPortsPostCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the pcloud networks ports post created response
func (o *PcloudNetworksPortsPostCreated) Code() int {
	return 201
}

func (o *PcloudNetworksPortsPostCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports][%d] pcloudNetworksPortsPostCreated %s", 201, payload)
}

func (o *PcloudNetworksPortsPostCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports][%d] pcloudNetworksPortsPostCreated %s", 201, payload)
}

func (o *PcloudNetworksPortsPostCreated) GetPayload() *models.NetworkPort {
	return o.Payload
}

func (o *PcloudNetworksPortsPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksPortsPostBadRequest creates a PcloudNetworksPortsPostBadRequest with default headers values
func NewPcloudNetworksPortsPostBadRequest() *PcloudNetworksPortsPostBadRequest {
	return &PcloudNetworksPortsPostBadRequest{}
}

/*
PcloudNetworksPortsPostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudNetworksPortsPostBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks ports post bad request response has a 2xx status code
func (o *PcloudNetworksPortsPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks ports post bad request response has a 3xx status code
func (o *PcloudNetworksPortsPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks ports post bad request response has a 4xx status code
func (o *PcloudNetworksPortsPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud networks ports post bad request response has a 5xx status code
func (o *PcloudNetworksPortsPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks ports post bad request response a status code equal to that given
func (o *PcloudNetworksPortsPostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud networks ports post bad request response
func (o *PcloudNetworksPortsPostBadRequest) Code() int {
	return 400
}

func (o *PcloudNetworksPortsPostBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports][%d] pcloudNetworksPortsPostBadRequest %s", 400, payload)
}

func (o *PcloudNetworksPortsPostBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports][%d] pcloudNetworksPortsPostBadRequest %s", 400, payload)
}

func (o *PcloudNetworksPortsPostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksPortsPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksPortsPostUnauthorized creates a PcloudNetworksPortsPostUnauthorized with default headers values
func NewPcloudNetworksPortsPostUnauthorized() *PcloudNetworksPortsPostUnauthorized {
	return &PcloudNetworksPortsPostUnauthorized{}
}

/*
PcloudNetworksPortsPostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudNetworksPortsPostUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks ports post unauthorized response has a 2xx status code
func (o *PcloudNetworksPortsPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks ports post unauthorized response has a 3xx status code
func (o *PcloudNetworksPortsPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks ports post unauthorized response has a 4xx status code
func (o *PcloudNetworksPortsPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud networks ports post unauthorized response has a 5xx status code
func (o *PcloudNetworksPortsPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks ports post unauthorized response a status code equal to that given
func (o *PcloudNetworksPortsPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud networks ports post unauthorized response
func (o *PcloudNetworksPortsPostUnauthorized) Code() int {
	return 401
}

func (o *PcloudNetworksPortsPostUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports][%d] pcloudNetworksPortsPostUnauthorized %s", 401, payload)
}

func (o *PcloudNetworksPortsPostUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports][%d] pcloudNetworksPortsPostUnauthorized %s", 401, payload)
}

func (o *PcloudNetworksPortsPostUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksPortsPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksPortsPostForbidden creates a PcloudNetworksPortsPostForbidden with default headers values
func NewPcloudNetworksPortsPostForbidden() *PcloudNetworksPortsPostForbidden {
	return &PcloudNetworksPortsPostForbidden{}
}

/*
PcloudNetworksPortsPostForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudNetworksPortsPostForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks ports post forbidden response has a 2xx status code
func (o *PcloudNetworksPortsPostForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks ports post forbidden response has a 3xx status code
func (o *PcloudNetworksPortsPostForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks ports post forbidden response has a 4xx status code
func (o *PcloudNetworksPortsPostForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud networks ports post forbidden response has a 5xx status code
func (o *PcloudNetworksPortsPostForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks ports post forbidden response a status code equal to that given
func (o *PcloudNetworksPortsPostForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud networks ports post forbidden response
func (o *PcloudNetworksPortsPostForbidden) Code() int {
	return 403
}

func (o *PcloudNetworksPortsPostForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports][%d] pcloudNetworksPortsPostForbidden %s", 403, payload)
}

func (o *PcloudNetworksPortsPostForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports][%d] pcloudNetworksPortsPostForbidden %s", 403, payload)
}

func (o *PcloudNetworksPortsPostForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksPortsPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksPortsPostNotFound creates a PcloudNetworksPortsPostNotFound with default headers values
func NewPcloudNetworksPortsPostNotFound() *PcloudNetworksPortsPostNotFound {
	return &PcloudNetworksPortsPostNotFound{}
}

/*
PcloudNetworksPortsPostNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudNetworksPortsPostNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks ports post not found response has a 2xx status code
func (o *PcloudNetworksPortsPostNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks ports post not found response has a 3xx status code
func (o *PcloudNetworksPortsPostNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks ports post not found response has a 4xx status code
func (o *PcloudNetworksPortsPostNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud networks ports post not found response has a 5xx status code
func (o *PcloudNetworksPortsPostNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks ports post not found response a status code equal to that given
func (o *PcloudNetworksPortsPostNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud networks ports post not found response
func (o *PcloudNetworksPortsPostNotFound) Code() int {
	return 404
}

func (o *PcloudNetworksPortsPostNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports][%d] pcloudNetworksPortsPostNotFound %s", 404, payload)
}

func (o *PcloudNetworksPortsPostNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports][%d] pcloudNetworksPortsPostNotFound %s", 404, payload)
}

func (o *PcloudNetworksPortsPostNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksPortsPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksPortsPostConflict creates a PcloudNetworksPortsPostConflict with default headers values
func NewPcloudNetworksPortsPostConflict() *PcloudNetworksPortsPostConflict {
	return &PcloudNetworksPortsPostConflict{}
}

/*
PcloudNetworksPortsPostConflict describes a response with status code 409, with default header values.

Conflict
*/
type PcloudNetworksPortsPostConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks ports post conflict response has a 2xx status code
func (o *PcloudNetworksPortsPostConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks ports post conflict response has a 3xx status code
func (o *PcloudNetworksPortsPostConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks ports post conflict response has a 4xx status code
func (o *PcloudNetworksPortsPostConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud networks ports post conflict response has a 5xx status code
func (o *PcloudNetworksPortsPostConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks ports post conflict response a status code equal to that given
func (o *PcloudNetworksPortsPostConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the pcloud networks ports post conflict response
func (o *PcloudNetworksPortsPostConflict) Code() int {
	return 409
}

func (o *PcloudNetworksPortsPostConflict) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports][%d] pcloudNetworksPortsPostConflict %s", 409, payload)
}

func (o *PcloudNetworksPortsPostConflict) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports][%d] pcloudNetworksPortsPostConflict %s", 409, payload)
}

func (o *PcloudNetworksPortsPostConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksPortsPostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksPortsPostUnprocessableEntity creates a PcloudNetworksPortsPostUnprocessableEntity with default headers values
func NewPcloudNetworksPortsPostUnprocessableEntity() *PcloudNetworksPortsPostUnprocessableEntity {
	return &PcloudNetworksPortsPostUnprocessableEntity{}
}

/*
PcloudNetworksPortsPostUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudNetworksPortsPostUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks ports post unprocessable entity response has a 2xx status code
func (o *PcloudNetworksPortsPostUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks ports post unprocessable entity response has a 3xx status code
func (o *PcloudNetworksPortsPostUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks ports post unprocessable entity response has a 4xx status code
func (o *PcloudNetworksPortsPostUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud networks ports post unprocessable entity response has a 5xx status code
func (o *PcloudNetworksPortsPostUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks ports post unprocessable entity response a status code equal to that given
func (o *PcloudNetworksPortsPostUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the pcloud networks ports post unprocessable entity response
func (o *PcloudNetworksPortsPostUnprocessableEntity) Code() int {
	return 422
}

func (o *PcloudNetworksPortsPostUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports][%d] pcloudNetworksPortsPostUnprocessableEntity %s", 422, payload)
}

func (o *PcloudNetworksPortsPostUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports][%d] pcloudNetworksPortsPostUnprocessableEntity %s", 422, payload)
}

func (o *PcloudNetworksPortsPostUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksPortsPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksPortsPostInternalServerError creates a PcloudNetworksPortsPostInternalServerError with default headers values
func NewPcloudNetworksPortsPostInternalServerError() *PcloudNetworksPortsPostInternalServerError {
	return &PcloudNetworksPortsPostInternalServerError{}
}

/*
PcloudNetworksPortsPostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudNetworksPortsPostInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks ports post internal server error response has a 2xx status code
func (o *PcloudNetworksPortsPostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks ports post internal server error response has a 3xx status code
func (o *PcloudNetworksPortsPostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks ports post internal server error response has a 4xx status code
func (o *PcloudNetworksPortsPostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud networks ports post internal server error response has a 5xx status code
func (o *PcloudNetworksPortsPostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud networks ports post internal server error response a status code equal to that given
func (o *PcloudNetworksPortsPostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud networks ports post internal server error response
func (o *PcloudNetworksPortsPostInternalServerError) Code() int {
	return 500
}

func (o *PcloudNetworksPortsPostInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports][%d] pcloudNetworksPortsPostInternalServerError %s", 500, payload)
}

func (o *PcloudNetworksPortsPostInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports][%d] pcloudNetworksPortsPostInternalServerError %s", 500, payload)
}

func (o *PcloudNetworksPortsPostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksPortsPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
