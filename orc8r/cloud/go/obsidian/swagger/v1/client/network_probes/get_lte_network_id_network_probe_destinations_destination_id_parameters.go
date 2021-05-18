// Code generated by go-swagger; DO NOT EDIT.

package network_probes

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

// NewGetLTENetworkIDNetworkProbeDestinationsDestinationIDParams creates a new GetLTENetworkIDNetworkProbeDestinationsDestinationIDParams object
// with the default values initialized.
func NewGetLTENetworkIDNetworkProbeDestinationsDestinationIDParams() *GetLTENetworkIDNetworkProbeDestinationsDestinationIDParams {
	var ()
	return &GetLTENetworkIDNetworkProbeDestinationsDestinationIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLTENetworkIDNetworkProbeDestinationsDestinationIDParamsWithTimeout creates a new GetLTENetworkIDNetworkProbeDestinationsDestinationIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLTENetworkIDNetworkProbeDestinationsDestinationIDParamsWithTimeout(timeout time.Duration) *GetLTENetworkIDNetworkProbeDestinationsDestinationIDParams {
	var ()
	return &GetLTENetworkIDNetworkProbeDestinationsDestinationIDParams{

		timeout: timeout,
	}
}

// NewGetLTENetworkIDNetworkProbeDestinationsDestinationIDParamsWithContext creates a new GetLTENetworkIDNetworkProbeDestinationsDestinationIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLTENetworkIDNetworkProbeDestinationsDestinationIDParamsWithContext(ctx context.Context) *GetLTENetworkIDNetworkProbeDestinationsDestinationIDParams {
	var ()
	return &GetLTENetworkIDNetworkProbeDestinationsDestinationIDParams{

		Context: ctx,
	}
}

// NewGetLTENetworkIDNetworkProbeDestinationsDestinationIDParamsWithHTTPClient creates a new GetLTENetworkIDNetworkProbeDestinationsDestinationIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLTENetworkIDNetworkProbeDestinationsDestinationIDParamsWithHTTPClient(client *http.Client) *GetLTENetworkIDNetworkProbeDestinationsDestinationIDParams {
	var ()
	return &GetLTENetworkIDNetworkProbeDestinationsDestinationIDParams{
		HTTPClient: client,
	}
}

/*GetLTENetworkIDNetworkProbeDestinationsDestinationIDParams contains all the parameters to send to the API endpoint
for the get LTE network ID network probe destinations destination ID operation typically these are written to a http.Request
*/
type GetLTENetworkIDNetworkProbeDestinationsDestinationIDParams struct {

	/*DestinationID
	  Network Probe Destination ID

	*/
	DestinationID string
	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get LTE network ID network probe destinations destination ID params
func (o *GetLTENetworkIDNetworkProbeDestinationsDestinationIDParams) WithTimeout(timeout time.Duration) *GetLTENetworkIDNetworkProbeDestinationsDestinationIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get LTE network ID network probe destinations destination ID params
func (o *GetLTENetworkIDNetworkProbeDestinationsDestinationIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get LTE network ID network probe destinations destination ID params
func (o *GetLTENetworkIDNetworkProbeDestinationsDestinationIDParams) WithContext(ctx context.Context) *GetLTENetworkIDNetworkProbeDestinationsDestinationIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get LTE network ID network probe destinations destination ID params
func (o *GetLTENetworkIDNetworkProbeDestinationsDestinationIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get LTE network ID network probe destinations destination ID params
func (o *GetLTENetworkIDNetworkProbeDestinationsDestinationIDParams) WithHTTPClient(client *http.Client) *GetLTENetworkIDNetworkProbeDestinationsDestinationIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get LTE network ID network probe destinations destination ID params
func (o *GetLTENetworkIDNetworkProbeDestinationsDestinationIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDestinationID adds the destinationID to the get LTE network ID network probe destinations destination ID params
func (o *GetLTENetworkIDNetworkProbeDestinationsDestinationIDParams) WithDestinationID(destinationID string) *GetLTENetworkIDNetworkProbeDestinationsDestinationIDParams {
	o.SetDestinationID(destinationID)
	return o
}

// SetDestinationID adds the destinationId to the get LTE network ID network probe destinations destination ID params
func (o *GetLTENetworkIDNetworkProbeDestinationsDestinationIDParams) SetDestinationID(destinationID string) {
	o.DestinationID = destinationID
}

// WithNetworkID adds the networkID to the get LTE network ID network probe destinations destination ID params
func (o *GetLTENetworkIDNetworkProbeDestinationsDestinationIDParams) WithNetworkID(networkID string) *GetLTENetworkIDNetworkProbeDestinationsDestinationIDParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get LTE network ID network probe destinations destination ID params
func (o *GetLTENetworkIDNetworkProbeDestinationsDestinationIDParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLTENetworkIDNetworkProbeDestinationsDestinationIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param destination_id
	if err := r.SetPathParam("destination_id", o.DestinationID); err != nil {
		return err
	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}