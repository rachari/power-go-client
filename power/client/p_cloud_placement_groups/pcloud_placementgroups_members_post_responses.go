// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_placement_groups

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

// PcloudPlacementgroupsMembersPostReader is a Reader for the PcloudPlacementgroupsMembersPost structure.
type PcloudPlacementgroupsMembersPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPlacementgroupsMembersPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudPlacementgroupsMembersPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudPlacementgroupsMembersPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudPlacementgroupsMembersPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudPlacementgroupsMembersPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudPlacementgroupsMembersPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPcloudPlacementgroupsMembersPostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudPlacementgroupsMembersPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudPlacementgroupsMembersPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}/members] pcloud.placementgroups.members.post", response, response.Code())
	}
}

// NewPcloudPlacementgroupsMembersPostOK creates a PcloudPlacementgroupsMembersPostOK with default headers values
func NewPcloudPlacementgroupsMembersPostOK() *PcloudPlacementgroupsMembersPostOK {
	return &PcloudPlacementgroupsMembersPostOK{}
}

/*
PcloudPlacementgroupsMembersPostOK describes a response with status code 200, with default header values.

OK
*/
type PcloudPlacementgroupsMembersPostOK struct {
	Payload *models.PlacementGroup
}

// IsSuccess returns true when this pcloud placementgroups members post o k response has a 2xx status code
func (o *PcloudPlacementgroupsMembersPostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud placementgroups members post o k response has a 3xx status code
func (o *PcloudPlacementgroupsMembersPostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud placementgroups members post o k response has a 4xx status code
func (o *PcloudPlacementgroupsMembersPostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud placementgroups members post o k response has a 5xx status code
func (o *PcloudPlacementgroupsMembersPostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud placementgroups members post o k response a status code equal to that given
func (o *PcloudPlacementgroupsMembersPostOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud placementgroups members post o k response
func (o *PcloudPlacementgroupsMembersPostOK) Code() int {
	return 200
}

func (o *PcloudPlacementgroupsMembersPostOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}/members][%d] pcloudPlacementgroupsMembersPostOK %s", 200, payload)
}

func (o *PcloudPlacementgroupsMembersPostOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}/members][%d] pcloudPlacementgroupsMembersPostOK %s", 200, payload)
}

func (o *PcloudPlacementgroupsMembersPostOK) GetPayload() *models.PlacementGroup {
	return o.Payload
}

func (o *PcloudPlacementgroupsMembersPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PlacementGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPlacementgroupsMembersPostBadRequest creates a PcloudPlacementgroupsMembersPostBadRequest with default headers values
func NewPcloudPlacementgroupsMembersPostBadRequest() *PcloudPlacementgroupsMembersPostBadRequest {
	return &PcloudPlacementgroupsMembersPostBadRequest{}
}

/*
PcloudPlacementgroupsMembersPostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudPlacementgroupsMembersPostBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud placementgroups members post bad request response has a 2xx status code
func (o *PcloudPlacementgroupsMembersPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud placementgroups members post bad request response has a 3xx status code
func (o *PcloudPlacementgroupsMembersPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud placementgroups members post bad request response has a 4xx status code
func (o *PcloudPlacementgroupsMembersPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud placementgroups members post bad request response has a 5xx status code
func (o *PcloudPlacementgroupsMembersPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud placementgroups members post bad request response a status code equal to that given
func (o *PcloudPlacementgroupsMembersPostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud placementgroups members post bad request response
func (o *PcloudPlacementgroupsMembersPostBadRequest) Code() int {
	return 400
}

func (o *PcloudPlacementgroupsMembersPostBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}/members][%d] pcloudPlacementgroupsMembersPostBadRequest %s", 400, payload)
}

func (o *PcloudPlacementgroupsMembersPostBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}/members][%d] pcloudPlacementgroupsMembersPostBadRequest %s", 400, payload)
}

func (o *PcloudPlacementgroupsMembersPostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPlacementgroupsMembersPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPlacementgroupsMembersPostUnauthorized creates a PcloudPlacementgroupsMembersPostUnauthorized with default headers values
func NewPcloudPlacementgroupsMembersPostUnauthorized() *PcloudPlacementgroupsMembersPostUnauthorized {
	return &PcloudPlacementgroupsMembersPostUnauthorized{}
}

/*
PcloudPlacementgroupsMembersPostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudPlacementgroupsMembersPostUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud placementgroups members post unauthorized response has a 2xx status code
func (o *PcloudPlacementgroupsMembersPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud placementgroups members post unauthorized response has a 3xx status code
func (o *PcloudPlacementgroupsMembersPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud placementgroups members post unauthorized response has a 4xx status code
func (o *PcloudPlacementgroupsMembersPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud placementgroups members post unauthorized response has a 5xx status code
func (o *PcloudPlacementgroupsMembersPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud placementgroups members post unauthorized response a status code equal to that given
func (o *PcloudPlacementgroupsMembersPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud placementgroups members post unauthorized response
func (o *PcloudPlacementgroupsMembersPostUnauthorized) Code() int {
	return 401
}

func (o *PcloudPlacementgroupsMembersPostUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}/members][%d] pcloudPlacementgroupsMembersPostUnauthorized %s", 401, payload)
}

func (o *PcloudPlacementgroupsMembersPostUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}/members][%d] pcloudPlacementgroupsMembersPostUnauthorized %s", 401, payload)
}

func (o *PcloudPlacementgroupsMembersPostUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPlacementgroupsMembersPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPlacementgroupsMembersPostForbidden creates a PcloudPlacementgroupsMembersPostForbidden with default headers values
func NewPcloudPlacementgroupsMembersPostForbidden() *PcloudPlacementgroupsMembersPostForbidden {
	return &PcloudPlacementgroupsMembersPostForbidden{}
}

/*
PcloudPlacementgroupsMembersPostForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudPlacementgroupsMembersPostForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud placementgroups members post forbidden response has a 2xx status code
func (o *PcloudPlacementgroupsMembersPostForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud placementgroups members post forbidden response has a 3xx status code
func (o *PcloudPlacementgroupsMembersPostForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud placementgroups members post forbidden response has a 4xx status code
func (o *PcloudPlacementgroupsMembersPostForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud placementgroups members post forbidden response has a 5xx status code
func (o *PcloudPlacementgroupsMembersPostForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud placementgroups members post forbidden response a status code equal to that given
func (o *PcloudPlacementgroupsMembersPostForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud placementgroups members post forbidden response
func (o *PcloudPlacementgroupsMembersPostForbidden) Code() int {
	return 403
}

func (o *PcloudPlacementgroupsMembersPostForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}/members][%d] pcloudPlacementgroupsMembersPostForbidden %s", 403, payload)
}

func (o *PcloudPlacementgroupsMembersPostForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}/members][%d] pcloudPlacementgroupsMembersPostForbidden %s", 403, payload)
}

func (o *PcloudPlacementgroupsMembersPostForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPlacementgroupsMembersPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPlacementgroupsMembersPostNotFound creates a PcloudPlacementgroupsMembersPostNotFound with default headers values
func NewPcloudPlacementgroupsMembersPostNotFound() *PcloudPlacementgroupsMembersPostNotFound {
	return &PcloudPlacementgroupsMembersPostNotFound{}
}

/*
PcloudPlacementgroupsMembersPostNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudPlacementgroupsMembersPostNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud placementgroups members post not found response has a 2xx status code
func (o *PcloudPlacementgroupsMembersPostNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud placementgroups members post not found response has a 3xx status code
func (o *PcloudPlacementgroupsMembersPostNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud placementgroups members post not found response has a 4xx status code
func (o *PcloudPlacementgroupsMembersPostNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud placementgroups members post not found response has a 5xx status code
func (o *PcloudPlacementgroupsMembersPostNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud placementgroups members post not found response a status code equal to that given
func (o *PcloudPlacementgroupsMembersPostNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud placementgroups members post not found response
func (o *PcloudPlacementgroupsMembersPostNotFound) Code() int {
	return 404
}

func (o *PcloudPlacementgroupsMembersPostNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}/members][%d] pcloudPlacementgroupsMembersPostNotFound %s", 404, payload)
}

func (o *PcloudPlacementgroupsMembersPostNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}/members][%d] pcloudPlacementgroupsMembersPostNotFound %s", 404, payload)
}

func (o *PcloudPlacementgroupsMembersPostNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPlacementgroupsMembersPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPlacementgroupsMembersPostConflict creates a PcloudPlacementgroupsMembersPostConflict with default headers values
func NewPcloudPlacementgroupsMembersPostConflict() *PcloudPlacementgroupsMembersPostConflict {
	return &PcloudPlacementgroupsMembersPostConflict{}
}

/*
PcloudPlacementgroupsMembersPostConflict describes a response with status code 409, with default header values.

Conflict
*/
type PcloudPlacementgroupsMembersPostConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud placementgroups members post conflict response has a 2xx status code
func (o *PcloudPlacementgroupsMembersPostConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud placementgroups members post conflict response has a 3xx status code
func (o *PcloudPlacementgroupsMembersPostConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud placementgroups members post conflict response has a 4xx status code
func (o *PcloudPlacementgroupsMembersPostConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud placementgroups members post conflict response has a 5xx status code
func (o *PcloudPlacementgroupsMembersPostConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud placementgroups members post conflict response a status code equal to that given
func (o *PcloudPlacementgroupsMembersPostConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the pcloud placementgroups members post conflict response
func (o *PcloudPlacementgroupsMembersPostConflict) Code() int {
	return 409
}

func (o *PcloudPlacementgroupsMembersPostConflict) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}/members][%d] pcloudPlacementgroupsMembersPostConflict %s", 409, payload)
}

func (o *PcloudPlacementgroupsMembersPostConflict) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}/members][%d] pcloudPlacementgroupsMembersPostConflict %s", 409, payload)
}

func (o *PcloudPlacementgroupsMembersPostConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPlacementgroupsMembersPostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPlacementgroupsMembersPostUnprocessableEntity creates a PcloudPlacementgroupsMembersPostUnprocessableEntity with default headers values
func NewPcloudPlacementgroupsMembersPostUnprocessableEntity() *PcloudPlacementgroupsMembersPostUnprocessableEntity {
	return &PcloudPlacementgroupsMembersPostUnprocessableEntity{}
}

/*
PcloudPlacementgroupsMembersPostUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudPlacementgroupsMembersPostUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud placementgroups members post unprocessable entity response has a 2xx status code
func (o *PcloudPlacementgroupsMembersPostUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud placementgroups members post unprocessable entity response has a 3xx status code
func (o *PcloudPlacementgroupsMembersPostUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud placementgroups members post unprocessable entity response has a 4xx status code
func (o *PcloudPlacementgroupsMembersPostUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud placementgroups members post unprocessable entity response has a 5xx status code
func (o *PcloudPlacementgroupsMembersPostUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud placementgroups members post unprocessable entity response a status code equal to that given
func (o *PcloudPlacementgroupsMembersPostUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the pcloud placementgroups members post unprocessable entity response
func (o *PcloudPlacementgroupsMembersPostUnprocessableEntity) Code() int {
	return 422
}

func (o *PcloudPlacementgroupsMembersPostUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}/members][%d] pcloudPlacementgroupsMembersPostUnprocessableEntity %s", 422, payload)
}

func (o *PcloudPlacementgroupsMembersPostUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}/members][%d] pcloudPlacementgroupsMembersPostUnprocessableEntity %s", 422, payload)
}

func (o *PcloudPlacementgroupsMembersPostUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPlacementgroupsMembersPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPlacementgroupsMembersPostInternalServerError creates a PcloudPlacementgroupsMembersPostInternalServerError with default headers values
func NewPcloudPlacementgroupsMembersPostInternalServerError() *PcloudPlacementgroupsMembersPostInternalServerError {
	return &PcloudPlacementgroupsMembersPostInternalServerError{}
}

/*
PcloudPlacementgroupsMembersPostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudPlacementgroupsMembersPostInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud placementgroups members post internal server error response has a 2xx status code
func (o *PcloudPlacementgroupsMembersPostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud placementgroups members post internal server error response has a 3xx status code
func (o *PcloudPlacementgroupsMembersPostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud placementgroups members post internal server error response has a 4xx status code
func (o *PcloudPlacementgroupsMembersPostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud placementgroups members post internal server error response has a 5xx status code
func (o *PcloudPlacementgroupsMembersPostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud placementgroups members post internal server error response a status code equal to that given
func (o *PcloudPlacementgroupsMembersPostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud placementgroups members post internal server error response
func (o *PcloudPlacementgroupsMembersPostInternalServerError) Code() int {
	return 500
}

func (o *PcloudPlacementgroupsMembersPostInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}/members][%d] pcloudPlacementgroupsMembersPostInternalServerError %s", 500, payload)
}

func (o *PcloudPlacementgroupsMembersPostInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}/members][%d] pcloudPlacementgroupsMembersPostInternalServerError %s", 500, payload)
}

func (o *PcloudPlacementgroupsMembersPostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPlacementgroupsMembersPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
