package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/cilium/cilium/api/v1/models"
)

// NewGetIdentityParams creates a new GetIdentityParams object
// with the default values initialized.
func NewGetIdentityParams() *GetIdentityParams {
	var ()
	return &GetIdentityParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIdentityParamsWithTimeout creates a new GetIdentityParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIdentityParamsWithTimeout(timeout time.Duration) *GetIdentityParams {
	var ()
	return &GetIdentityParams{

		timeout: timeout,
	}
}

// NewGetIdentityParamsWithContext creates a new GetIdentityParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIdentityParamsWithContext(ctx context.Context) *GetIdentityParams {
	var ()
	return &GetIdentityParams{

		Context: ctx,
	}
}

// NewGetIdentityParamsWithHTTPClient creates a new GetIdentityParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIdentityParamsWithHTTPClient(client *http.Client) *GetIdentityParams {
	var ()
	return &GetIdentityParams{
		HTTPClient: client,
	}
}

/*GetIdentityParams contains all the parameters to send to the API endpoint
for the get identity operation typically these are written to a http.Request
*/
type GetIdentityParams struct {

	/*Labels*/
	Labels models.Labels

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get identity params
func (o *GetIdentityParams) WithTimeout(timeout time.Duration) *GetIdentityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get identity params
func (o *GetIdentityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get identity params
func (o *GetIdentityParams) WithContext(ctx context.Context) *GetIdentityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get identity params
func (o *GetIdentityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get identity params
func (o *GetIdentityParams) WithHTTPClient(client *http.Client) *GetIdentityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get identity params
func (o *GetIdentityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLabels adds the labels to the get identity params
func (o *GetIdentityParams) WithLabels(labels models.Labels) *GetIdentityParams {
	o.SetLabels(labels)
	return o
}

// SetLabels adds the labels to the get identity params
func (o *GetIdentityParams) SetLabels(labels models.Labels) {
	o.Labels = labels
}

// WriteToRequest writes these params to a swagger request
func (o *GetIdentityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Labels); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
