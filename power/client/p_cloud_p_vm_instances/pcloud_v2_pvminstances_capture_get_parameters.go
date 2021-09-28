// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_p_vm_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPcloudV2PvminstancesCaptureGetParams creates a new PcloudV2PvminstancesCaptureGetParams object
// with the default values initialized.
func NewPcloudV2PvminstancesCaptureGetParams() *PcloudV2PvminstancesCaptureGetParams {
	var ()
	return &PcloudV2PvminstancesCaptureGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudV2PvminstancesCaptureGetParamsWithTimeout creates a new PcloudV2PvminstancesCaptureGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPcloudV2PvminstancesCaptureGetParamsWithTimeout(timeout time.Duration) *PcloudV2PvminstancesCaptureGetParams {
	var ()
	return &PcloudV2PvminstancesCaptureGetParams{

		timeout: timeout,
	}
}

// NewPcloudV2PvminstancesCaptureGetParamsWithContext creates a new PcloudV2PvminstancesCaptureGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewPcloudV2PvminstancesCaptureGetParamsWithContext(ctx context.Context) *PcloudV2PvminstancesCaptureGetParams {
	var ()
	return &PcloudV2PvminstancesCaptureGetParams{

		Context: ctx,
	}
}

// NewPcloudV2PvminstancesCaptureGetParamsWithHTTPClient creates a new PcloudV2PvminstancesCaptureGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPcloudV2PvminstancesCaptureGetParamsWithHTTPClient(client *http.Client) *PcloudV2PvminstancesCaptureGetParams {
	var ()
	return &PcloudV2PvminstancesCaptureGetParams{
		HTTPClient: client,
	}
}

/*PcloudV2PvminstancesCaptureGetParams contains all the parameters to send to the API endpoint
for the pcloud v2 pvminstances capture get operation typically these are written to a http.Request
*/
type PcloudV2PvminstancesCaptureGetParams struct {

	/*CloudInstanceID
	  Cloud Instance ID of a PCloud Instance

	*/
	CloudInstanceID string
	/*PvmInstanceID
	  PCloud PVM Instance ID

	*/
	PvmInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pcloud v2 pvminstances capture get params
func (o *PcloudV2PvminstancesCaptureGetParams) WithTimeout(timeout time.Duration) *PcloudV2PvminstancesCaptureGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud v2 pvminstances capture get params
func (o *PcloudV2PvminstancesCaptureGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud v2 pvminstances capture get params
func (o *PcloudV2PvminstancesCaptureGetParams) WithContext(ctx context.Context) *PcloudV2PvminstancesCaptureGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud v2 pvminstances capture get params
func (o *PcloudV2PvminstancesCaptureGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud v2 pvminstances capture get params
func (o *PcloudV2PvminstancesCaptureGetParams) WithHTTPClient(client *http.Client) *PcloudV2PvminstancesCaptureGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud v2 pvminstances capture get params
func (o *PcloudV2PvminstancesCaptureGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud v2 pvminstances capture get params
func (o *PcloudV2PvminstancesCaptureGetParams) WithCloudInstanceID(cloudInstanceID string) *PcloudV2PvminstancesCaptureGetParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud v2 pvminstances capture get params
func (o *PcloudV2PvminstancesCaptureGetParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithPvmInstanceID adds the pvmInstanceID to the pcloud v2 pvminstances capture get params
func (o *PcloudV2PvminstancesCaptureGetParams) WithPvmInstanceID(pvmInstanceID string) *PcloudV2PvminstancesCaptureGetParams {
	o.SetPvmInstanceID(pvmInstanceID)
	return o
}

// SetPvmInstanceID adds the pvmInstanceId to the pcloud v2 pvminstances capture get params
func (o *PcloudV2PvminstancesCaptureGetParams) SetPvmInstanceID(pvmInstanceID string) {
	o.PvmInstanceID = pvmInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudV2PvminstancesCaptureGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	// path param pvm_instance_id
	if err := r.SetPathParam("pvm_instance_id", o.PvmInstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
