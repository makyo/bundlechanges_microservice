// Package params holds all of the request and response parameters for endpoints, along
// with any required structs.
package params

import "github.com/juju/httprequest"

// StatusUnprocessableEntity describes a request that is correct, but the Entity
// provided in the request is not able to be processed.
const StatusUnprocessableEntity = 422

// ErrorResponse represents an error encountered by the server.
type ErrorResponse struct {
	Message string
	Code    ErrorCode
}

// ErrorCode holds the code of an error returned from the API.
type ErrorCode string

// Error implements errrors.Error.
func (code ErrorCode) Error() string {
	return string(code)
}

// Possible ErrorCodes from the API.
const (
	ErrUnparsable          ErrorCode = "unparsable data"
	ErrVerificationFailure ErrorCode = "data cannot be verified"
)

// Change represents one change in the change set that the GUI needs to execute
// to deploy the bundle.
type Change struct {
	Args     []interface{} `json:"args"`
	Id       string        `json:"id"`
	Requires []string      `json:"requires"`
	Method   string        `json:"method"`
}

// ChangesResponse contains the results of parsing a bundle into a list of
// changes.
type ChangesResponse struct {
	Changes []Change `json:"changes"`
}

// ChangesRequest contains the bundle as a YAML-encoded string which is to be
// parsed into a list of changes.
type ChangesRequest struct {
	Bundle string `json:"bundle"`
}

// ChangesFromYAMLParams contains the parameters required for passing a bundle
// to the API and recieving a list of changes in return
type ChangesFromYAMLParams struct {
	httprequest.Route `httprequest:"POST /bundlechanges/fromYAML"`
	NicelyFormatted   bool           `httprequest:"nice,form"`
	Body              ChangesRequest `httprequest:",body"`
}
