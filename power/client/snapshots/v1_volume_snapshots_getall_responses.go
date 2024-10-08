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

// V1VolumeSnapshotsGetallReader is a Reader for the V1VolumeSnapshotsGetall structure.
type V1VolumeSnapshotsGetallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1VolumeSnapshotsGetallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1VolumeSnapshotsGetallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewV1VolumeSnapshotsGetallUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV1VolumeSnapshotsGetallForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewV1VolumeSnapshotsGetallNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV1VolumeSnapshotsGetallInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewV1VolumeSnapshotsGetallServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/volume-snapshots] v1.volume-snapshots.getall", response, response.Code())
	}
}

// NewV1VolumeSnapshotsGetallOK creates a V1VolumeSnapshotsGetallOK with default headers values
func NewV1VolumeSnapshotsGetallOK() *V1VolumeSnapshotsGetallOK {
	return &V1VolumeSnapshotsGetallOK{}
}

/*
V1VolumeSnapshotsGetallOK describes a response with status code 200, with default header values.

OK
*/
type V1VolumeSnapshotsGetallOK struct {
	Payload *models.VolumeSnapshotList
}

// IsSuccess returns true when this v1 volume snapshots getall o k response has a 2xx status code
func (o *V1VolumeSnapshotsGetallOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this v1 volume snapshots getall o k response has a 3xx status code
func (o *V1VolumeSnapshotsGetallOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 volume snapshots getall o k response has a 4xx status code
func (o *V1VolumeSnapshotsGetallOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 volume snapshots getall o k response has a 5xx status code
func (o *V1VolumeSnapshotsGetallOK) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 volume snapshots getall o k response a status code equal to that given
func (o *V1VolumeSnapshotsGetallOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the v1 volume snapshots getall o k response
func (o *V1VolumeSnapshotsGetallOK) Code() int {
	return 200
}

func (o *V1VolumeSnapshotsGetallOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/volume-snapshots][%d] v1VolumeSnapshotsGetallOK %s", 200, payload)
}

func (o *V1VolumeSnapshotsGetallOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/volume-snapshots][%d] v1VolumeSnapshotsGetallOK %s", 200, payload)
}

func (o *V1VolumeSnapshotsGetallOK) GetPayload() *models.VolumeSnapshotList {
	return o.Payload
}

func (o *V1VolumeSnapshotsGetallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumeSnapshotList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1VolumeSnapshotsGetallUnauthorized creates a V1VolumeSnapshotsGetallUnauthorized with default headers values
func NewV1VolumeSnapshotsGetallUnauthorized() *V1VolumeSnapshotsGetallUnauthorized {
	return &V1VolumeSnapshotsGetallUnauthorized{}
}

/*
V1VolumeSnapshotsGetallUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type V1VolumeSnapshotsGetallUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 volume snapshots getall unauthorized response has a 2xx status code
func (o *V1VolumeSnapshotsGetallUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 volume snapshots getall unauthorized response has a 3xx status code
func (o *V1VolumeSnapshotsGetallUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 volume snapshots getall unauthorized response has a 4xx status code
func (o *V1VolumeSnapshotsGetallUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 volume snapshots getall unauthorized response has a 5xx status code
func (o *V1VolumeSnapshotsGetallUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 volume snapshots getall unauthorized response a status code equal to that given
func (o *V1VolumeSnapshotsGetallUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the v1 volume snapshots getall unauthorized response
func (o *V1VolumeSnapshotsGetallUnauthorized) Code() int {
	return 401
}

func (o *V1VolumeSnapshotsGetallUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/volume-snapshots][%d] v1VolumeSnapshotsGetallUnauthorized %s", 401, payload)
}

func (o *V1VolumeSnapshotsGetallUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/volume-snapshots][%d] v1VolumeSnapshotsGetallUnauthorized %s", 401, payload)
}

func (o *V1VolumeSnapshotsGetallUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1VolumeSnapshotsGetallUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1VolumeSnapshotsGetallForbidden creates a V1VolumeSnapshotsGetallForbidden with default headers values
func NewV1VolumeSnapshotsGetallForbidden() *V1VolumeSnapshotsGetallForbidden {
	return &V1VolumeSnapshotsGetallForbidden{}
}

/*
V1VolumeSnapshotsGetallForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type V1VolumeSnapshotsGetallForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 volume snapshots getall forbidden response has a 2xx status code
func (o *V1VolumeSnapshotsGetallForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 volume snapshots getall forbidden response has a 3xx status code
func (o *V1VolumeSnapshotsGetallForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 volume snapshots getall forbidden response has a 4xx status code
func (o *V1VolumeSnapshotsGetallForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 volume snapshots getall forbidden response has a 5xx status code
func (o *V1VolumeSnapshotsGetallForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 volume snapshots getall forbidden response a status code equal to that given
func (o *V1VolumeSnapshotsGetallForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the v1 volume snapshots getall forbidden response
func (o *V1VolumeSnapshotsGetallForbidden) Code() int {
	return 403
}

func (o *V1VolumeSnapshotsGetallForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/volume-snapshots][%d] v1VolumeSnapshotsGetallForbidden %s", 403, payload)
}

func (o *V1VolumeSnapshotsGetallForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/volume-snapshots][%d] v1VolumeSnapshotsGetallForbidden %s", 403, payload)
}

func (o *V1VolumeSnapshotsGetallForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1VolumeSnapshotsGetallForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1VolumeSnapshotsGetallNotFound creates a V1VolumeSnapshotsGetallNotFound with default headers values
func NewV1VolumeSnapshotsGetallNotFound() *V1VolumeSnapshotsGetallNotFound {
	return &V1VolumeSnapshotsGetallNotFound{}
}

/*
V1VolumeSnapshotsGetallNotFound describes a response with status code 404, with default header values.

Not Found
*/
type V1VolumeSnapshotsGetallNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 volume snapshots getall not found response has a 2xx status code
func (o *V1VolumeSnapshotsGetallNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 volume snapshots getall not found response has a 3xx status code
func (o *V1VolumeSnapshotsGetallNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 volume snapshots getall not found response has a 4xx status code
func (o *V1VolumeSnapshotsGetallNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this v1 volume snapshots getall not found response has a 5xx status code
func (o *V1VolumeSnapshotsGetallNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this v1 volume snapshots getall not found response a status code equal to that given
func (o *V1VolumeSnapshotsGetallNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the v1 volume snapshots getall not found response
func (o *V1VolumeSnapshotsGetallNotFound) Code() int {
	return 404
}

func (o *V1VolumeSnapshotsGetallNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/volume-snapshots][%d] v1VolumeSnapshotsGetallNotFound %s", 404, payload)
}

func (o *V1VolumeSnapshotsGetallNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/volume-snapshots][%d] v1VolumeSnapshotsGetallNotFound %s", 404, payload)
}

func (o *V1VolumeSnapshotsGetallNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1VolumeSnapshotsGetallNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1VolumeSnapshotsGetallInternalServerError creates a V1VolumeSnapshotsGetallInternalServerError with default headers values
func NewV1VolumeSnapshotsGetallInternalServerError() *V1VolumeSnapshotsGetallInternalServerError {
	return &V1VolumeSnapshotsGetallInternalServerError{}
}

/*
V1VolumeSnapshotsGetallInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type V1VolumeSnapshotsGetallInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 volume snapshots getall internal server error response has a 2xx status code
func (o *V1VolumeSnapshotsGetallInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 volume snapshots getall internal server error response has a 3xx status code
func (o *V1VolumeSnapshotsGetallInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 volume snapshots getall internal server error response has a 4xx status code
func (o *V1VolumeSnapshotsGetallInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 volume snapshots getall internal server error response has a 5xx status code
func (o *V1VolumeSnapshotsGetallInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this v1 volume snapshots getall internal server error response a status code equal to that given
func (o *V1VolumeSnapshotsGetallInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the v1 volume snapshots getall internal server error response
func (o *V1VolumeSnapshotsGetallInternalServerError) Code() int {
	return 500
}

func (o *V1VolumeSnapshotsGetallInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/volume-snapshots][%d] v1VolumeSnapshotsGetallInternalServerError %s", 500, payload)
}

func (o *V1VolumeSnapshotsGetallInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/volume-snapshots][%d] v1VolumeSnapshotsGetallInternalServerError %s", 500, payload)
}

func (o *V1VolumeSnapshotsGetallInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1VolumeSnapshotsGetallInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV1VolumeSnapshotsGetallServiceUnavailable creates a V1VolumeSnapshotsGetallServiceUnavailable with default headers values
func NewV1VolumeSnapshotsGetallServiceUnavailable() *V1VolumeSnapshotsGetallServiceUnavailable {
	return &V1VolumeSnapshotsGetallServiceUnavailable{}
}

/*
V1VolumeSnapshotsGetallServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable
*/
type V1VolumeSnapshotsGetallServiceUnavailable struct {
	Payload *models.Error
}

// IsSuccess returns true when this v1 volume snapshots getall service unavailable response has a 2xx status code
func (o *V1VolumeSnapshotsGetallServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v1 volume snapshots getall service unavailable response has a 3xx status code
func (o *V1VolumeSnapshotsGetallServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v1 volume snapshots getall service unavailable response has a 4xx status code
func (o *V1VolumeSnapshotsGetallServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this v1 volume snapshots getall service unavailable response has a 5xx status code
func (o *V1VolumeSnapshotsGetallServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this v1 volume snapshots getall service unavailable response a status code equal to that given
func (o *V1VolumeSnapshotsGetallServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the v1 volume snapshots getall service unavailable response
func (o *V1VolumeSnapshotsGetallServiceUnavailable) Code() int {
	return 503
}

func (o *V1VolumeSnapshotsGetallServiceUnavailable) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/volume-snapshots][%d] v1VolumeSnapshotsGetallServiceUnavailable %s", 503, payload)
}

func (o *V1VolumeSnapshotsGetallServiceUnavailable) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/volume-snapshots][%d] v1VolumeSnapshotsGetallServiceUnavailable %s", 503, payload)
}

func (o *V1VolumeSnapshotsGetallServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *V1VolumeSnapshotsGetallServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
