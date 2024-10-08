// Code generated by go-swagger; DO NOT EDIT.

package snapshots

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

// V1VolumeSnapshotsGetReader is a Reader for the V1VolumeSnapshotsGet structure.
type V1VolumeSnapshotsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1VolumeSnapshotsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1VolumeSnapshotsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewV1VolumeSnapshotsGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewV1VolumeSnapshotsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV1VolumeSnapshotsGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewV1VolumeSnapshotsGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV1VolumeSnapshotsGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewV1VolumeSnapshotsGetServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/volume-snapshots/{volume_snapshot_uuid}] v1.volume-snapshots.get", response, response.Code())
	}
}

// NewV1VolumeSnapshotsGetOK creates a V1VolumeSnapshotsGetOK with default headers values
func NewV1VolumeSnapshotsGetOK() *V1VolumeSnapshotsGetOK {
	return &V1VolumeSnapshotsGetOK{}
}

/*
V1VolumeSnapshotsGetOK describes a response with status code 200, with default header values.

OK
*/
type V1VolumeSnapshotsGetOK struct {
	Payload *models.SnapshotV1
}

// IsSuccess returns true when this v1 volume snapshots get o k response has a 2xx status code
func (o *V1VolumeSnapshotsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this v1 volume snapshots get o k response has a 3xx status code
func (o *V1VolumeSnapshotsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 volume snapshots get o k response has a 4xx status code
func (o *V1VolumeSnapshotsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 volume snapshots get o k response has a 5xx status code
func (o *V1VolumeSnapshotsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 volume snapshots get o k response a status code equal to that given
func (o *V1VolumeSnapshotsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the v1 volume snapshots get o k response
func (o *V1VolumeSnapshotsGetOK) Code() int {
	return 200
}

func (o *V1VolumeSnapshotsGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/volume-snapshots/{volume_snapshot_uuid}][%d] v1VolumeSnapshotsGetOK %s", 200, payload)
}

func (o *V1VolumeSnapshotsGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/volume-snapshots/{volume_snapshot_uuid}][%d] v1VolumeSnapshotsGetOK %s", 200, payload)
}

func (o *V1VolumeSnapshotsGetOK) GetPayload() *models.SnapshotV1 {
	return o.Payload
}

func (o *V1VolumeSnapshotsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnapshotV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1VolumeSnapshotsGetBadRequest creates a V1VolumeSnapshotsGetBadRequest with default headers values
func NewV1VolumeSnapshotsGetBadRequest() *V1VolumeSnapshotsGetBadRequest {
	return &V1VolumeSnapshotsGetBadRequest{}
}

/*
V1VolumeSnapshotsGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type V1VolumeSnapshotsGetBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 volume snapshots get bad request response has a 2xx status code
func (o *V1VolumeSnapshotsGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 volume snapshots get bad request response has a 3xx status code
func (o *V1VolumeSnapshotsGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 volume snapshots get bad request response has a 4xx status code
func (o *V1VolumeSnapshotsGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 volume snapshots get bad request response has a 5xx status code
func (o *V1VolumeSnapshotsGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 volume snapshots get bad request response a status code equal to that given
func (o *V1VolumeSnapshotsGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the v1 volume snapshots get bad request response
func (o *V1VolumeSnapshotsGetBadRequest) Code() int {
	return 400
}

func (o *V1VolumeSnapshotsGetBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/volume-snapshots/{volume_snapshot_uuid}][%d] v1VolumeSnapshotsGetBadRequest %s", 400, payload)
}

func (o *V1VolumeSnapshotsGetBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/volume-snapshots/{volume_snapshot_uuid}][%d] v1VolumeSnapshotsGetBadRequest %s", 400, payload)
}

func (o *V1VolumeSnapshotsGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1VolumeSnapshotsGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1VolumeSnapshotsGetUnauthorized creates a V1VolumeSnapshotsGetUnauthorized with default headers values
func NewV1VolumeSnapshotsGetUnauthorized() *V1VolumeSnapshotsGetUnauthorized {
	return &V1VolumeSnapshotsGetUnauthorized{}
}

/*
V1VolumeSnapshotsGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type V1VolumeSnapshotsGetUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 volume snapshots get unauthorized response has a 2xx status code
func (o *V1VolumeSnapshotsGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 volume snapshots get unauthorized response has a 3xx status code
func (o *V1VolumeSnapshotsGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 volume snapshots get unauthorized response has a 4xx status code
func (o *V1VolumeSnapshotsGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 volume snapshots get unauthorized response has a 5xx status code
func (o *V1VolumeSnapshotsGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 volume snapshots get unauthorized response a status code equal to that given
func (o *V1VolumeSnapshotsGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the v1 volume snapshots get unauthorized response
func (o *V1VolumeSnapshotsGetUnauthorized) Code() int {
	return 401
}

func (o *V1VolumeSnapshotsGetUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/volume-snapshots/{volume_snapshot_uuid}][%d] v1VolumeSnapshotsGetUnauthorized %s", 401, payload)
}

func (o *V1VolumeSnapshotsGetUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/volume-snapshots/{volume_snapshot_uuid}][%d] v1VolumeSnapshotsGetUnauthorized %s", 401, payload)
}

func (o *V1VolumeSnapshotsGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1VolumeSnapshotsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1VolumeSnapshotsGetForbidden creates a V1VolumeSnapshotsGetForbidden with default headers values
func NewV1VolumeSnapshotsGetForbidden() *V1VolumeSnapshotsGetForbidden {
	return &V1VolumeSnapshotsGetForbidden{}
}

/*
V1VolumeSnapshotsGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type V1VolumeSnapshotsGetForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 volume snapshots get forbidden response has a 2xx status code
func (o *V1VolumeSnapshotsGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 volume snapshots get forbidden response has a 3xx status code
func (o *V1VolumeSnapshotsGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 volume snapshots get forbidden response has a 4xx status code
func (o *V1VolumeSnapshotsGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 volume snapshots get forbidden response has a 5xx status code
func (o *V1VolumeSnapshotsGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 volume snapshots get forbidden response a status code equal to that given
func (o *V1VolumeSnapshotsGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the v1 volume snapshots get forbidden response
func (o *V1VolumeSnapshotsGetForbidden) Code() int {
	return 403
}

func (o *V1VolumeSnapshotsGetForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/volume-snapshots/{volume_snapshot_uuid}][%d] v1VolumeSnapshotsGetForbidden %s", 403, payload)
}

func (o *V1VolumeSnapshotsGetForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/volume-snapshots/{volume_snapshot_uuid}][%d] v1VolumeSnapshotsGetForbidden %s", 403, payload)
}

func (o *V1VolumeSnapshotsGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1VolumeSnapshotsGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1VolumeSnapshotsGetNotFound creates a V1VolumeSnapshotsGetNotFound with default headers values
func NewV1VolumeSnapshotsGetNotFound() *V1VolumeSnapshotsGetNotFound {
	return &V1VolumeSnapshotsGetNotFound{}
}

/*
V1VolumeSnapshotsGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type V1VolumeSnapshotsGetNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 volume snapshots get not found response has a 2xx status code
func (o *V1VolumeSnapshotsGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 volume snapshots get not found response has a 3xx status code
func (o *V1VolumeSnapshotsGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 volume snapshots get not found response has a 4xx status code
func (o *V1VolumeSnapshotsGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 volume snapshots get not found response has a 5xx status code
func (o *V1VolumeSnapshotsGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 volume snapshots get not found response a status code equal to that given
func (o *V1VolumeSnapshotsGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the v1 volume snapshots get not found response
func (o *V1VolumeSnapshotsGetNotFound) Code() int {
	return 404
}

func (o *V1VolumeSnapshotsGetNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/volume-snapshots/{volume_snapshot_uuid}][%d] v1VolumeSnapshotsGetNotFound %s", 404, payload)
}

func (o *V1VolumeSnapshotsGetNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/volume-snapshots/{volume_snapshot_uuid}][%d] v1VolumeSnapshotsGetNotFound %s", 404, payload)
}

func (o *V1VolumeSnapshotsGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1VolumeSnapshotsGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1VolumeSnapshotsGetInternalServerError creates a V1VolumeSnapshotsGetInternalServerError with default headers values
func NewV1VolumeSnapshotsGetInternalServerError() *V1VolumeSnapshotsGetInternalServerError {
	return &V1VolumeSnapshotsGetInternalServerError{}
}

/*
V1VolumeSnapshotsGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type V1VolumeSnapshotsGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 volume snapshots get internal server error response has a 2xx status code
func (o *V1VolumeSnapshotsGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 volume snapshots get internal server error response has a 3xx status code
func (o *V1VolumeSnapshotsGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 volume snapshots get internal server error response has a 4xx status code
func (o *V1VolumeSnapshotsGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 volume snapshots get internal server error response has a 5xx status code
func (o *V1VolumeSnapshotsGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this v1 volume snapshots get internal server error response a status code equal to that given
func (o *V1VolumeSnapshotsGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the v1 volume snapshots get internal server error response
func (o *V1VolumeSnapshotsGetInternalServerError) Code() int {
	return 500
}

func (o *V1VolumeSnapshotsGetInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/volume-snapshots/{volume_snapshot_uuid}][%d] v1VolumeSnapshotsGetInternalServerError %s", 500, payload)
}

func (o *V1VolumeSnapshotsGetInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/volume-snapshots/{volume_snapshot_uuid}][%d] v1VolumeSnapshotsGetInternalServerError %s", 500, payload)
}

func (o *V1VolumeSnapshotsGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1VolumeSnapshotsGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1VolumeSnapshotsGetServiceUnavailable creates a V1VolumeSnapshotsGetServiceUnavailable with default headers values
func NewV1VolumeSnapshotsGetServiceUnavailable() *V1VolumeSnapshotsGetServiceUnavailable {
	return &V1VolumeSnapshotsGetServiceUnavailable{}
}

/*
V1VolumeSnapshotsGetServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable
*/
type V1VolumeSnapshotsGetServiceUnavailable struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 volume snapshots get service unavailable response has a 2xx status code
func (o *V1VolumeSnapshotsGetServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 volume snapshots get service unavailable response has a 3xx status code
func (o *V1VolumeSnapshotsGetServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 volume snapshots get service unavailable response has a 4xx status code
func (o *V1VolumeSnapshotsGetServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 volume snapshots get service unavailable response has a 5xx status code
func (o *V1VolumeSnapshotsGetServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this v1 volume snapshots get service unavailable response a status code equal to that given
func (o *V1VolumeSnapshotsGetServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the v1 volume snapshots get service unavailable response
func (o *V1VolumeSnapshotsGetServiceUnavailable) Code() int {
	return 503
}

func (o *V1VolumeSnapshotsGetServiceUnavailable) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/volume-snapshots/{volume_snapshot_uuid}][%d] v1VolumeSnapshotsGetServiceUnavailable %s", 503, payload)
}

func (o *V1VolumeSnapshotsGetServiceUnavailable) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/volume-snapshots/{volume_snapshot_uuid}][%d] v1VolumeSnapshotsGetServiceUnavailable %s", 503, payload)
}

func (o *V1VolumeSnapshotsGetServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1VolumeSnapshotsGetServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
