// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_p_vm_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudPvminstancesConsoleGetReader is a Reader for the PcloudPvminstancesConsoleGet structure.
type PcloudPvminstancesConsoleGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPvminstancesConsoleGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPcloudPvminstancesConsoleGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPcloudPvminstancesConsoleGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPcloudPvminstancesConsoleGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPcloudPvminstancesConsoleGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPcloudPvminstancesConsoleGetOK creates a PcloudPvminstancesConsoleGetOK with default headers values
func NewPcloudPvminstancesConsoleGetOK() *PcloudPvminstancesConsoleGetOK {
	return &PcloudPvminstancesConsoleGetOK{}
}

/*PcloudPvminstancesConsoleGetOK handles this case with default header values.

OK
*/
type PcloudPvminstancesConsoleGetOK struct {
	Payload *models.ConsoleLanguages
}

func (o *PcloudPvminstancesConsoleGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsoleGetOK  %+v", 200, o.Payload)
}

func (o *PcloudPvminstancesConsoleGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsoleLanguages)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesConsoleGetBadRequest creates a PcloudPvminstancesConsoleGetBadRequest with default headers values
func NewPcloudPvminstancesConsoleGetBadRequest() *PcloudPvminstancesConsoleGetBadRequest {
	return &PcloudPvminstancesConsoleGetBadRequest{}
}

/*PcloudPvminstancesConsoleGetBadRequest handles this case with default header values.

Bad Request
*/
type PcloudPvminstancesConsoleGetBadRequest struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesConsoleGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsoleGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudPvminstancesConsoleGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesConsoleGetUnauthorized creates a PcloudPvminstancesConsoleGetUnauthorized with default headers values
func NewPcloudPvminstancesConsoleGetUnauthorized() *PcloudPvminstancesConsoleGetUnauthorized {
	return &PcloudPvminstancesConsoleGetUnauthorized{}
}

/*PcloudPvminstancesConsoleGetUnauthorized handles this case with default header values.

Unauthorized
*/
type PcloudPvminstancesConsoleGetUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesConsoleGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsoleGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudPvminstancesConsoleGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesConsoleGetInternalServerError creates a PcloudPvminstancesConsoleGetInternalServerError with default headers values
func NewPcloudPvminstancesConsoleGetInternalServerError() *PcloudPvminstancesConsoleGetInternalServerError {
	return &PcloudPvminstancesConsoleGetInternalServerError{}
}

/*PcloudPvminstancesConsoleGetInternalServerError handles this case with default header values.

Internal Server Error
*/
type PcloudPvminstancesConsoleGetInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesConsoleGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsoleGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPvminstancesConsoleGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
