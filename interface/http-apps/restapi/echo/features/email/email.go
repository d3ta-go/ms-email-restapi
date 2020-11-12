package email

import (
	"net/http"

	appEmail "github.com/d3ta-go/ddd-mod-email/modules/email/application"
	appEmailDTO "github.com/d3ta-go/ddd-mod-email/modules/email/application/dto/email"
	appEmailDTOET "github.com/d3ta-go/ddd-mod-email/modules/email/application/dto/email_template"
	"github.com/d3ta-go/system/interface/http-apps/restapi/echo/features"
	"github.com/d3ta-go/system/interface/http-apps/restapi/echo/response"
	"github.com/d3ta-go/system/system/handler"
	"github.com/labstack/echo/v4"
)

// NewFEmail new FEmail
func NewFEmail(h *handler.Handler) (*FEmail, error) {
	var err error

	f := new(FEmail)
	f.SetHandler(h)

	if f.appEmail, err = appEmail.NewEmailApp(h); err != nil {
		return nil, err
	}

	return f, nil
}

// FEmail represent Email Feature
type FEmail struct {
	features.BaseFeature
	appEmail *appEmail.EmailApp
}

// ListAllEmailTemplate list all EmailTemplate
func (f *FEmail) ListAllEmailTemplate(c echo.Context) error {
	// identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.TranslateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	resp, err := f.appEmail.EmailTemplateSvc.ListAll(i)
	if err != nil {
		return f.TranslateErrorMessage(err, c)
	}

	return response.OKWithData(resp, c)
}

// FindEmailTemplateByCode find EmailTemplateByCode
func (f *FEmail) FindEmailTemplateByCode(c echo.Context) error {
	// identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.TranslateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	// params
	code := c.Param("code")

	req := new(appEmailDTOET.ETFindByCodeReqDTO)
	req.Code = code

	resp, err := f.appEmail.EmailTemplateSvc.FindByCode(req, i)
	if err != nil {
		return f.TranslateErrorMessage(err, c)
	}

	return response.OKWithData(resp, c)
}

// CreateEmailTemplate create EmailTemplate
func (f *FEmail) CreateEmailTemplate(c echo.Context) error {
	// identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.TranslateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	req := new(appEmailDTOET.ETCreateReqDTO)
	if err := c.Bind(req); err != nil {
		return f.TranslateErrorMessage(err, c)
	}

	resp, err := f.appEmail.EmailTemplateSvc.Create(req, i)
	if err != nil {
		return f.TranslateErrorMessage(err, c)
	}

	return response.OKWithData(resp, c)
}

// UpdateEmailTemplate update existing EmailTemplate
func (f *FEmail) UpdateEmailTemplate(c echo.Context) error {
	// identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.TranslateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	// param
	code := c.Param("code")

	reqKeys := new(appEmailDTOET.ETUpdateKeysDTO)
	reqKeys.Code = code

	reqData := new(appEmailDTOET.ETUpdateDataDTO)
	if err := c.Bind(reqData); err != nil {
		return f.TranslateErrorMessage(err, c)
	}

	req := &appEmailDTOET.ETUpdateReqDTO{
		Keys: reqKeys,
		Data: reqData,
	}

	resp, err := f.appEmail.EmailTemplateSvc.Update(req, i)
	if err != nil {
		return f.TranslateErrorMessage(err, c)
	}

	return response.OKWithData(resp, c)
}

// SetActiveEmailTemplate set existing EmailTemplate active status
func (f *FEmail) SetActiveEmailTemplate(c echo.Context) error {
	// identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.TranslateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	// param
	code := c.Param("code")

	reqKeys := new(appEmailDTOET.ETSetActiveKeysDTO)
	reqKeys.Code = code

	reqData := new(appEmailDTOET.ETSetActiveDataDTO)
	if err := c.Bind(reqData); err != nil {
		return f.TranslateErrorMessage(err, c)
	}

	req := &appEmailDTOET.ETSetActiveReqDTO{
		Keys: reqKeys,
		Data: reqData,
	}

	resp, err := f.appEmail.EmailTemplateSvc.SetActive(req, i)
	if err != nil {
		return f.TranslateErrorMessage(err, c)
	}

	return response.OKWithData(resp, c)
}

// DeleteEmailTemplate delete existing EmailTemplate with template version
func (f *FEmail) DeleteEmailTemplate(c echo.Context) error {
	// identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.TranslateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	// param
	code := c.Param("code")

	req := new(appEmailDTOET.ETDeleteReqDTO)
	req.Code = code

	resp, err := f.appEmail.EmailTemplateSvc.Delete(req, i)
	if err != nil {
		return f.TranslateErrorMessage(err, c)
	}

	return response.OKWithData(resp, c)
}

// SendEmail send Email
func (f *FEmail) SendEmail(c echo.Context) error {
	// identity
	i, err := f.SetIdentity(c)
	if err != nil {
		return f.TranslateErrorMessage(err, c)
	}
	if !i.IsLogin || i.IsAnonymous {
		return response.FailWithMessageWithCode(http.StatusForbidden, "Forbidden Access", c)
	}

	req := new(appEmailDTO.SendEmailReqDTO)
	if err := c.Bind(req); err != nil {
		return f.TranslateErrorMessage(err, c)
	}

	resp, err := f.appEmail.EmailSvc.Send(req, i)
	if err != nil {
		return f.TranslateErrorMessage(err, c)
	}

	return response.OKWithData(resp, c)
}
