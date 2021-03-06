// Code generated by go-swagger; DO NOT EDIT.

package bar_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// BarReader is a Reader for the Bar structure.
type BarReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BarReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewBarOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBarOK creates a BarOK with default headers values
func NewBarOK() *BarOK {
	return &BarOK{}
}

/*BarOK handles this case with default header values.

A successful response.
*/
type BarOK struct {
	Payload interface{}
}

func (o *BarOK) Error() string {
	return fmt.Sprintf("[POST /Bar][%d] barOK  %+v", 200, o.Payload)
}

func (o *BarOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*BarBody bar body
swagger:model BarBody
*/
type BarBody struct {

	// common
	Common *BarParamsBodyCommon `json:"common,omitempty"`
}

// Validate validates this bar body
func (o *BarBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCommon(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *BarBody) validateCommon(formats strfmt.Registry) error {

	if swag.IsZero(o.Common) { // not required
		return nil
	}

	if o.Common != nil {
		if err := o.Common.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "common")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *BarBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *BarBody) UnmarshalBinary(b []byte) error {
	var res BarBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*BarParamsBodyCommon bar params body common
swagger:model BarParamsBodyCommon
*/
type BarParamsBodyCommon struct {

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this bar params body common
func (o *BarParamsBodyCommon) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *BarParamsBodyCommon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *BarParamsBodyCommon) UnmarshalBinary(b []byte) error {
	var res BarParamsBodyCommon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
