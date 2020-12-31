// Code generated by go-swagger; DO NOT EDIT.

package bootfiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// DownloadBootFilesOKCode is the HTTP code returned for type DownloadBootFilesOK
const DownloadBootFilesOKCode int = 200

/*DownloadBootFilesOK Success.

swagger:response downloadBootFilesOK
*/
type DownloadBootFilesOK struct {

	/*
	  In: Body
	*/
	Payload io.ReadCloser `json:"body,omitempty"`
}

// NewDownloadBootFilesOK creates DownloadBootFilesOK with default headers values
func NewDownloadBootFilesOK() *DownloadBootFilesOK {

	return &DownloadBootFilesOK{}
}

// WithPayload adds the payload to the download boot files o k response
func (o *DownloadBootFilesOK) WithPayload(payload io.ReadCloser) *DownloadBootFilesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download boot files o k response
func (o *DownloadBootFilesOK) SetPayload(payload io.ReadCloser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadBootFilesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// DownloadBootFilesTemporaryRedirectCode is the HTTP code returned for type DownloadBootFilesTemporaryRedirect
const DownloadBootFilesTemporaryRedirectCode int = 307

/*DownloadBootFilesTemporaryRedirect Redirect.

swagger:response downloadBootFilesTemporaryRedirect
*/
type DownloadBootFilesTemporaryRedirect struct {
	/*

	 */
	Location string `json:"Location"`
}

// NewDownloadBootFilesTemporaryRedirect creates DownloadBootFilesTemporaryRedirect with default headers values
func NewDownloadBootFilesTemporaryRedirect() *DownloadBootFilesTemporaryRedirect {

	return &DownloadBootFilesTemporaryRedirect{}
}

// WithLocation adds the location to the download boot files temporary redirect response
func (o *DownloadBootFilesTemporaryRedirect) WithLocation(location string) *DownloadBootFilesTemporaryRedirect {
	o.Location = location
	return o
}

// SetLocation sets the location to the download boot files temporary redirect response
func (o *DownloadBootFilesTemporaryRedirect) SetLocation(location string) {
	o.Location = location
}

// WriteResponse to the client
func (o *DownloadBootFilesTemporaryRedirect) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Location

	location := o.Location
	if location != "" {
		rw.Header().Set("Location", location)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(307)
}

// DownloadBootFilesUnauthorizedCode is the HTTP code returned for type DownloadBootFilesUnauthorized
const DownloadBootFilesUnauthorizedCode int = 401

/*DownloadBootFilesUnauthorized Unauthorized.

swagger:response downloadBootFilesUnauthorized
*/
type DownloadBootFilesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewDownloadBootFilesUnauthorized creates DownloadBootFilesUnauthorized with default headers values
func NewDownloadBootFilesUnauthorized() *DownloadBootFilesUnauthorized {

	return &DownloadBootFilesUnauthorized{}
}

// WithPayload adds the payload to the download boot files unauthorized response
func (o *DownloadBootFilesUnauthorized) WithPayload(payload *models.InfraError) *DownloadBootFilesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download boot files unauthorized response
func (o *DownloadBootFilesUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadBootFilesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadBootFilesForbiddenCode is the HTTP code returned for type DownloadBootFilesForbidden
const DownloadBootFilesForbiddenCode int = 403

/*DownloadBootFilesForbidden Forbidden.

swagger:response downloadBootFilesForbidden
*/
type DownloadBootFilesForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewDownloadBootFilesForbidden creates DownloadBootFilesForbidden with default headers values
func NewDownloadBootFilesForbidden() *DownloadBootFilesForbidden {

	return &DownloadBootFilesForbidden{}
}

// WithPayload adds the payload to the download boot files forbidden response
func (o *DownloadBootFilesForbidden) WithPayload(payload *models.InfraError) *DownloadBootFilesForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download boot files forbidden response
func (o *DownloadBootFilesForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadBootFilesForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadBootFilesInternalServerErrorCode is the HTTP code returned for type DownloadBootFilesInternalServerError
const DownloadBootFilesInternalServerErrorCode int = 500

/*DownloadBootFilesInternalServerError Error.

swagger:response downloadBootFilesInternalServerError
*/
type DownloadBootFilesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDownloadBootFilesInternalServerError creates DownloadBootFilesInternalServerError with default headers values
func NewDownloadBootFilesInternalServerError() *DownloadBootFilesInternalServerError {

	return &DownloadBootFilesInternalServerError{}
}

// WithPayload adds the payload to the download boot files internal server error response
func (o *DownloadBootFilesInternalServerError) WithPayload(payload *models.Error) *DownloadBootFilesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download boot files internal server error response
func (o *DownloadBootFilesInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadBootFilesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
