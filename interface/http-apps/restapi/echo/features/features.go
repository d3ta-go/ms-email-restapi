package features

import (
	"github.com/d3ta-go/ms-email-restapi/interface/http-apps/restapi/echo/features/email"
	"github.com/d3ta-go/ms-email-restapi/interface/http-apps/restapi/echo/features/openapis"
	"github.com/d3ta-go/system/interface/http-apps/restapi/echo/features/system"
	"github.com/d3ta-go/system/system/handler"
)

// NewFeatures create new Features
func NewFeatures(h *handler.Handler) (*Features, error) {
	var err error

	f := new(Features)
	f.handler = h

	if f.System, err = system.NewSystem(h); err != nil {
		return nil, err
	}
	if f.OpenAPIs, err = openapis.NewOpenAPIs(h); err != nil {
		return nil, err
	}

	if f.Email, err = email.NewFEmail(h); err != nil {
		return nil, err
	}

	return f, nil
}

// Features represent Features
type Features struct {
	handler *handler.Handler

	System   *system.FSystem
	OpenAPIs *openapis.FOpenAPIs
	Email    *email.FEmail
}
