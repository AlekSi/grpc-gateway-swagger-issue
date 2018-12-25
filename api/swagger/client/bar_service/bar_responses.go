// Code generated by go-swagger; DO NOT EDIT.

package bar_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

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
	Common interface{} `json:"common,omitempty"`
}

// Validate validates this bar body
func (o *BarBody) Validate(formats strfmt.Registry) error {
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