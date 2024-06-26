// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_v_p_n_connections

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

// PcloudVpnconnectionsPeersubnetsDeleteReader is a Reader for the PcloudVpnconnectionsPeersubnetsDelete structure.
type PcloudVpnconnectionsPeersubnetsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudVpnconnectionsPeersubnetsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudVpnconnectionsPeersubnetsDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudVpnconnectionsPeersubnetsDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudVpnconnectionsPeersubnetsDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudVpnconnectionsPeersubnetsDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudVpnconnectionsPeersubnetsDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudVpnconnectionsPeersubnetsDeleteUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudVpnconnectionsPeersubnetsDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets] pcloud.vpnconnections.peersubnets.delete", response, response.Code())
	}
}

// NewPcloudVpnconnectionsPeersubnetsDeleteOK creates a PcloudVpnconnectionsPeersubnetsDeleteOK with default headers values
func NewPcloudVpnconnectionsPeersubnetsDeleteOK() *PcloudVpnconnectionsPeersubnetsDeleteOK {
	return &PcloudVpnconnectionsPeersubnetsDeleteOK{}
}

/*
PcloudVpnconnectionsPeersubnetsDeleteOK describes a response with status code 200, with default header values.

OK
*/
type PcloudVpnconnectionsPeersubnetsDeleteOK struct {
	Payload *models.PeerSubnets
}

// IsSuccess returns true when this pcloud vpnconnections peersubnets delete o k response has a 2xx status code
func (o *PcloudVpnconnectionsPeersubnetsDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud vpnconnections peersubnets delete o k response has a 3xx status code
func (o *PcloudVpnconnectionsPeersubnetsDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud vpnconnections peersubnets delete o k response has a 4xx status code
func (o *PcloudVpnconnectionsPeersubnetsDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud vpnconnections peersubnets delete o k response has a 5xx status code
func (o *PcloudVpnconnectionsPeersubnetsDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud vpnconnections peersubnets delete o k response a status code equal to that given
func (o *PcloudVpnconnectionsPeersubnetsDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud vpnconnections peersubnets delete o k response
func (o *PcloudVpnconnectionsPeersubnetsDeleteOK) Code() int {
	return 200
}

func (o *PcloudVpnconnectionsPeersubnetsDeleteOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsDeleteOK %s", 200, payload)
}

func (o *PcloudVpnconnectionsPeersubnetsDeleteOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsDeleteOK %s", 200, payload)
}

func (o *PcloudVpnconnectionsPeersubnetsDeleteOK) GetPayload() *models.PeerSubnets {
	return o.Payload
}

func (o *PcloudVpnconnectionsPeersubnetsDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PeerSubnets)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPeersubnetsDeleteBadRequest creates a PcloudVpnconnectionsPeersubnetsDeleteBadRequest with default headers values
func NewPcloudVpnconnectionsPeersubnetsDeleteBadRequest() *PcloudVpnconnectionsPeersubnetsDeleteBadRequest {
	return &PcloudVpnconnectionsPeersubnetsDeleteBadRequest{}
}

/*
PcloudVpnconnectionsPeersubnetsDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudVpnconnectionsPeersubnetsDeleteBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud vpnconnections peersubnets delete bad request response has a 2xx status code
func (o *PcloudVpnconnectionsPeersubnetsDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud vpnconnections peersubnets delete bad request response has a 3xx status code
func (o *PcloudVpnconnectionsPeersubnetsDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud vpnconnections peersubnets delete bad request response has a 4xx status code
func (o *PcloudVpnconnectionsPeersubnetsDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud vpnconnections peersubnets delete bad request response has a 5xx status code
func (o *PcloudVpnconnectionsPeersubnetsDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud vpnconnections peersubnets delete bad request response a status code equal to that given
func (o *PcloudVpnconnectionsPeersubnetsDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud vpnconnections peersubnets delete bad request response
func (o *PcloudVpnconnectionsPeersubnetsDeleteBadRequest) Code() int {
	return 400
}

func (o *PcloudVpnconnectionsPeersubnetsDeleteBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsDeleteBadRequest %s", 400, payload)
}

func (o *PcloudVpnconnectionsPeersubnetsDeleteBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsDeleteBadRequest %s", 400, payload)
}

func (o *PcloudVpnconnectionsPeersubnetsDeleteBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsPeersubnetsDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPeersubnetsDeleteUnauthorized creates a PcloudVpnconnectionsPeersubnetsDeleteUnauthorized with default headers values
func NewPcloudVpnconnectionsPeersubnetsDeleteUnauthorized() *PcloudVpnconnectionsPeersubnetsDeleteUnauthorized {
	return &PcloudVpnconnectionsPeersubnetsDeleteUnauthorized{}
}

/*
PcloudVpnconnectionsPeersubnetsDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudVpnconnectionsPeersubnetsDeleteUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud vpnconnections peersubnets delete unauthorized response has a 2xx status code
func (o *PcloudVpnconnectionsPeersubnetsDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud vpnconnections peersubnets delete unauthorized response has a 3xx status code
func (o *PcloudVpnconnectionsPeersubnetsDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud vpnconnections peersubnets delete unauthorized response has a 4xx status code
func (o *PcloudVpnconnectionsPeersubnetsDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud vpnconnections peersubnets delete unauthorized response has a 5xx status code
func (o *PcloudVpnconnectionsPeersubnetsDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud vpnconnections peersubnets delete unauthorized response a status code equal to that given
func (o *PcloudVpnconnectionsPeersubnetsDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud vpnconnections peersubnets delete unauthorized response
func (o *PcloudVpnconnectionsPeersubnetsDeleteUnauthorized) Code() int {
	return 401
}

func (o *PcloudVpnconnectionsPeersubnetsDeleteUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsDeleteUnauthorized %s", 401, payload)
}

func (o *PcloudVpnconnectionsPeersubnetsDeleteUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsDeleteUnauthorized %s", 401, payload)
}

func (o *PcloudVpnconnectionsPeersubnetsDeleteUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsPeersubnetsDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPeersubnetsDeleteForbidden creates a PcloudVpnconnectionsPeersubnetsDeleteForbidden with default headers values
func NewPcloudVpnconnectionsPeersubnetsDeleteForbidden() *PcloudVpnconnectionsPeersubnetsDeleteForbidden {
	return &PcloudVpnconnectionsPeersubnetsDeleteForbidden{}
}

/*
PcloudVpnconnectionsPeersubnetsDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudVpnconnectionsPeersubnetsDeleteForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud vpnconnections peersubnets delete forbidden response has a 2xx status code
func (o *PcloudVpnconnectionsPeersubnetsDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud vpnconnections peersubnets delete forbidden response has a 3xx status code
func (o *PcloudVpnconnectionsPeersubnetsDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud vpnconnections peersubnets delete forbidden response has a 4xx status code
func (o *PcloudVpnconnectionsPeersubnetsDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud vpnconnections peersubnets delete forbidden response has a 5xx status code
func (o *PcloudVpnconnectionsPeersubnetsDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud vpnconnections peersubnets delete forbidden response a status code equal to that given
func (o *PcloudVpnconnectionsPeersubnetsDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud vpnconnections peersubnets delete forbidden response
func (o *PcloudVpnconnectionsPeersubnetsDeleteForbidden) Code() int {
	return 403
}

func (o *PcloudVpnconnectionsPeersubnetsDeleteForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsDeleteForbidden %s", 403, payload)
}

func (o *PcloudVpnconnectionsPeersubnetsDeleteForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsDeleteForbidden %s", 403, payload)
}

func (o *PcloudVpnconnectionsPeersubnetsDeleteForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsPeersubnetsDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPeersubnetsDeleteNotFound creates a PcloudVpnconnectionsPeersubnetsDeleteNotFound with default headers values
func NewPcloudVpnconnectionsPeersubnetsDeleteNotFound() *PcloudVpnconnectionsPeersubnetsDeleteNotFound {
	return &PcloudVpnconnectionsPeersubnetsDeleteNotFound{}
}

/*
PcloudVpnconnectionsPeersubnetsDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudVpnconnectionsPeersubnetsDeleteNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud vpnconnections peersubnets delete not found response has a 2xx status code
func (o *PcloudVpnconnectionsPeersubnetsDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud vpnconnections peersubnets delete not found response has a 3xx status code
func (o *PcloudVpnconnectionsPeersubnetsDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud vpnconnections peersubnets delete not found response has a 4xx status code
func (o *PcloudVpnconnectionsPeersubnetsDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud vpnconnections peersubnets delete not found response has a 5xx status code
func (o *PcloudVpnconnectionsPeersubnetsDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud vpnconnections peersubnets delete not found response a status code equal to that given
func (o *PcloudVpnconnectionsPeersubnetsDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud vpnconnections peersubnets delete not found response
func (o *PcloudVpnconnectionsPeersubnetsDeleteNotFound) Code() int {
	return 404
}

func (o *PcloudVpnconnectionsPeersubnetsDeleteNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsDeleteNotFound %s", 404, payload)
}

func (o *PcloudVpnconnectionsPeersubnetsDeleteNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsDeleteNotFound %s", 404, payload)
}

func (o *PcloudVpnconnectionsPeersubnetsDeleteNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsPeersubnetsDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPeersubnetsDeleteUnprocessableEntity creates a PcloudVpnconnectionsPeersubnetsDeleteUnprocessableEntity with default headers values
func NewPcloudVpnconnectionsPeersubnetsDeleteUnprocessableEntity() *PcloudVpnconnectionsPeersubnetsDeleteUnprocessableEntity {
	return &PcloudVpnconnectionsPeersubnetsDeleteUnprocessableEntity{}
}

/*
PcloudVpnconnectionsPeersubnetsDeleteUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudVpnconnectionsPeersubnetsDeleteUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud vpnconnections peersubnets delete unprocessable entity response has a 2xx status code
func (o *PcloudVpnconnectionsPeersubnetsDeleteUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud vpnconnections peersubnets delete unprocessable entity response has a 3xx status code
func (o *PcloudVpnconnectionsPeersubnetsDeleteUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud vpnconnections peersubnets delete unprocessable entity response has a 4xx status code
func (o *PcloudVpnconnectionsPeersubnetsDeleteUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud vpnconnections peersubnets delete unprocessable entity response has a 5xx status code
func (o *PcloudVpnconnectionsPeersubnetsDeleteUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud vpnconnections peersubnets delete unprocessable entity response a status code equal to that given
func (o *PcloudVpnconnectionsPeersubnetsDeleteUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the pcloud vpnconnections peersubnets delete unprocessable entity response
func (o *PcloudVpnconnectionsPeersubnetsDeleteUnprocessableEntity) Code() int {
	return 422
}

func (o *PcloudVpnconnectionsPeersubnetsDeleteUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsDeleteUnprocessableEntity %s", 422, payload)
}

func (o *PcloudVpnconnectionsPeersubnetsDeleteUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsDeleteUnprocessableEntity %s", 422, payload)
}

func (o *PcloudVpnconnectionsPeersubnetsDeleteUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsPeersubnetsDeleteUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPeersubnetsDeleteInternalServerError creates a PcloudVpnconnectionsPeersubnetsDeleteInternalServerError with default headers values
func NewPcloudVpnconnectionsPeersubnetsDeleteInternalServerError() *PcloudVpnconnectionsPeersubnetsDeleteInternalServerError {
	return &PcloudVpnconnectionsPeersubnetsDeleteInternalServerError{}
}

/*
PcloudVpnconnectionsPeersubnetsDeleteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudVpnconnectionsPeersubnetsDeleteInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud vpnconnections peersubnets delete internal server error response has a 2xx status code
func (o *PcloudVpnconnectionsPeersubnetsDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud vpnconnections peersubnets delete internal server error response has a 3xx status code
func (o *PcloudVpnconnectionsPeersubnetsDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud vpnconnections peersubnets delete internal server error response has a 4xx status code
func (o *PcloudVpnconnectionsPeersubnetsDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud vpnconnections peersubnets delete internal server error response has a 5xx status code
func (o *PcloudVpnconnectionsPeersubnetsDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud vpnconnections peersubnets delete internal server error response a status code equal to that given
func (o *PcloudVpnconnectionsPeersubnetsDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud vpnconnections peersubnets delete internal server error response
func (o *PcloudVpnconnectionsPeersubnetsDeleteInternalServerError) Code() int {
	return 500
}

func (o *PcloudVpnconnectionsPeersubnetsDeleteInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsDeleteInternalServerError %s", 500, payload)
}

func (o *PcloudVpnconnectionsPeersubnetsDeleteInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsDeleteInternalServerError %s", 500, payload)
}

func (o *PcloudVpnconnectionsPeersubnetsDeleteInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsPeersubnetsDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
