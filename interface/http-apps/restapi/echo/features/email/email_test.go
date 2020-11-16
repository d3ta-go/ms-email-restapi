package email

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	ht "github.com/d3ta-go/ms-email-restapi/interface/http-apps/restapi/echo/features/helper_test"
	"github.com/d3ta-go/system/system/initialize"
	"github.com/d3ta-go/system/system/utils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestEmail_ListAllEmailTemplate(t *testing.T) {
	// client request
	// none

	// setup echo
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/email/templates/list-all", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)

	// handler
	handler := ht.NewHandler()
	if err := initialize.LoadAllDatabaseConnection(handler); err != nil {
		t.Errorf("initialize.LoadAllDatabaseConnection: %s", err.Error())
		return
	}

	// set identity (test only)
	token, claims, err := ht.GenerateUserTestToken(handler, t)
	if err != nil {
		t.Errorf("generateUserTestToken: %s", err.Error())
		return
	}
	c.Set("identity.token.jwt", token)
	c.Set("identity.token.jwt.claims", claims)

	// test feature
	email, err := NewFEmail(handler)
	if err != nil {
		t.Errorf("NewFEmail: %s", err.Error())
		return
	}

	if assert.NoError(t, email.ListAllEmailTemplate(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		t.Logf("RESPONSE.Email.ListAllEmailTemplate: %s", res.Body.String())
	}
}

func TestEmail_CreateEmailTemplate(t *testing.T) {
	h := ht.NewHandler()

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.email.email-template.interface-layer.features.create.request")

	unique := utils.GenerateUUID()
	etCode := fmt.Sprintf(testData["et-code"], unique)
	etName := fmt.Sprintf(testData["et-name"], unique)
	// client request
	reqDTO := `{
	"code": "` + etCode + `",
	"name": "` + etName + `",
	"isActive": ` + testData["et-is-active"] + `,
	"emailFormat": "` + testData["et-email-format"] + `",
	"template": {
		"subjectTpl": "` + testData["et-tpl-subject"] + `",
		"bodyTpl": "` + testData["et-tpl-body"] + `"
	}
}`

	// setup echo
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/email/template", strings.NewReader(reqDTO))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)

	// handler
	handler := ht.NewHandler()
	if err := initialize.LoadAllDatabaseConnection(handler); err != nil {
		t.Errorf("initialize.LoadAllDatabaseConnection: %s", err.Error())
		return
	}

	// set identity (test only)
	token, claims, err := ht.GenerateUserTestToken(handler, t)
	if err != nil {
		t.Errorf("generateUserTestToken: %s", err.Error())
		return
	}
	c.Set("identity.token.jwt", token)
	c.Set("identity.token.jwt.claims", claims)

	// test feature
	email, err := NewFEmail(handler)
	if err != nil {
		t.Errorf("NewFEmail: %s", err.Error())
		return
	}

	if assert.NoError(t, email.CreateEmailTemplate(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		// save to test-data
		// save result for next test
		viper.Set("test-data.email.email-template.interface-layer.features.create.response.json", res.Body.String())
		viper.Set("test-data.email.email-template.interface-layer.features.find-by-code.request.et-code", etCode)

		viper.Set("test-data.email.email-template.interface-layer.features.update.request.et-code", etCode)
		viper.Set("test-data.email.email-template.interface-layer.features.update.request.et-name", etName)

		viper.Set("test-data.email.email-template.interface-layer.features.delete.request.et-code", etCode)
		viper.Set("test-data.email.email-template.interface-layer.features.set-active.request.et-code", etCode)

		if err := viper.WriteConfig(); err != nil {
			t.Errorf("Error: viper.WriteConfig(), %s", err.Error())
		}
		t.Logf("RESPONSE.Email.CreateEmailTemplate: %s", res.Body.String())
	}
}

func TestEmail_FindEmailTemplateByCode(t *testing.T) {
	h := ht.NewHandler()

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.email.email-template.interface-layer.features.find-by-code.request")

	// client request
	// --> set on context param [http method = GET]

	// setup echo
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/email/template/:code", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)
	c.SetParamNames("code")
	c.SetParamValues(testData["et-code"])

	// handler
	handler := ht.NewHandler()
	if err := initialize.LoadAllDatabaseConnection(handler); err != nil {
		t.Errorf("initialize.LoadAllDatabaseConnection: %s", err.Error())
		return
	}

	// set identity (test only)
	token, claims, err := ht.GenerateUserTestToken(handler, t)
	if err != nil {
		t.Errorf("generateUserTestToken: %s", err.Error())
		return
	}
	c.Set("identity.token.jwt", token)
	c.Set("identity.token.jwt.claims", claims)

	// test feature
	email, err := NewFEmail(handler)
	if err != nil {
		t.Errorf("NewFEmail: %s", err.Error())
		return
	}

	if assert.NoError(t, email.FindEmailTemplateByCode(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		// save to test-data
		// save result for next test
		viper.Set("test-data.email.email-template.interface-layer.features.find-by-code.response.json", res.Body.String())
		if err := viper.WriteConfig(); err != nil {
			t.Errorf("Error: viper.WriteConfig(), %s", err.Error())
		}
		t.Logf("RESPONSE.Email.FindEmailTemplateByCode: %s", res.Body.String())
	}
}

func TestEmail_UpdateEmailTemplate(t *testing.T) {
	h := ht.NewHandler()

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.email.email-template.interface-layer.features.update.request")

	// client request
	reqDTO := `{
	"name": "` + testData["et-name"] + ` Updated",
	"isActive": ` + testData["et-is-active"] + `,
	"emailFormat": "` + testData["et-email-format"] + `",
	"template": {
		"subjectTpl": "` + testData["et-tpl-subject"] + `",
		"bodyTpl": "` + testData["et-tpl-body"] + `"
	}
}`

	// setup echo
	e := echo.New()

	req := httptest.NewRequest(http.MethodPut, "/api/v1/email/template/update/:code", strings.NewReader(reqDTO))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)
	c.SetParamNames("code")
	c.SetParamValues(testData["et-code"])

	// handler
	handler := ht.NewHandler()
	if err := initialize.LoadAllDatabaseConnection(handler); err != nil {
		t.Errorf("initialize.LoadAllDatabaseConnection: %s", err.Error())
		return
	}

	// set identity (test only)
	token, claims, err := ht.GenerateUserTestToken(handler, t)
	if err != nil {
		t.Errorf("generateUserTestToken: %s", err.Error())
		return
	}
	c.Set("identity.token.jwt", token)
	c.Set("identity.token.jwt.claims", claims)

	// test feature
	email, err := NewFEmail(handler)
	if err != nil {
		t.Errorf("NewFEmail: %s", err.Error())
		return
	}

	if assert.NoError(t, email.UpdateEmailTemplate(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		// save to test-data
		// save result for next test
		viper.Set("test-data.email.email-template.interface-layer.features.update.response.json", res.Body.String())
		if err := viper.WriteConfig(); err != nil {
			t.Errorf("Error: viper.WriteConfig(), %s", err.Error())
		}
		t.Logf("RESPONSE.Email.UpdateEmailTemplate: %s", res.Body.String())
	}
}

func TestEmail_SetActiveEmailTemplate(t *testing.T) {
	h := ht.NewHandler()

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.email.email-template.interface-layer.features.set-active.request")

	// client request
	reqDTO := `{
	"isActive": ` + testData["et-is-active"] + `
}`

	// setup echo
	e := echo.New()

	req := httptest.NewRequest(http.MethodPut, "/api/v1/email/template/set-active/:code", strings.NewReader(reqDTO))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)
	c.SetParamNames("code")
	c.SetParamValues(testData["et-code"])

	// handler
	handler := ht.NewHandler()
	if err := initialize.LoadAllDatabaseConnection(handler); err != nil {
		t.Errorf("initialize.LoadAllDatabaseConnection: %s", err.Error())
		return
	}

	// set identity (test only)
	token, claims, err := ht.GenerateUserTestToken(handler, t)
	if err != nil {
		t.Errorf("generateUserTestToken: %s", err.Error())
		return
	}
	c.Set("identity.token.jwt", token)
	c.Set("identity.token.jwt.claims", claims)

	// test feature
	email, err := NewFEmail(handler)
	if err != nil {
		t.Errorf("NewFEmail: %s", err.Error())
		return
	}

	if assert.NoError(t, email.SetActiveEmailTemplate(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		// save to test-data
		// save result for next test
		viper.Set("test-data.email.email-template.interface-layer.features.set-active.response.json", res.Body.String())
		if err := viper.WriteConfig(); err != nil {
			t.Errorf("Error: viper.WriteConfig(), %s", err.Error())
		}
		t.Logf("RESPONSE.Email.SetActiveEmailTemplate: %s", res.Body.String())
	}
}

func TestEmail_DeleteEmailTemplate(t *testing.T) {
	h := ht.NewHandler()

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.email.email-template.interface-layer.features.delete.request")

	// client request
	// none

	// setup echo
	e := echo.New()

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/email/template/:code", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)
	c.SetParamNames("code")
	c.SetParamValues(testData["et-code"])

	// handler
	handler := ht.NewHandler()
	if err := initialize.LoadAllDatabaseConnection(handler); err != nil {
		t.Errorf("initialize.LoadAllDatabaseConnection: %s", err.Error())
		return
	}

	// set identity (test only)
	token, claims, err := ht.GenerateUserTestToken(handler, t)
	if err != nil {
		t.Errorf("generateUserTestToken: %s", err.Error())
		return
	}
	c.Set("identity.token.jwt", token)
	c.Set("identity.token.jwt.claims", claims)

	// test feature
	email, err := NewFEmail(handler)
	if err != nil {
		t.Errorf("NewFEmail: %s", err.Error())
		return
	}

	if assert.NoError(t, email.DeleteEmailTemplate(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		// save to test-data
		// save result for next test
		viper.Set("test-data.email.email-template.interface-layer.features.delete.response.json", res.Body.String())
		if err := viper.WriteConfig(); err != nil {
			t.Errorf("Error: viper.WriteConfig(), %s", err.Error())
		}
		t.Logf("RESPONSE.Email.DeleteEmailTemplate: %s", res.Body.String())
	}
}

func TestEmail_SendEmail(t *testing.T) {
	h := ht.NewHandler()

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.email.email.interface-layer.features.send.request")
	testDataET := viper.GetStringMapString("test-data.email.email.interface-layer.features.send.request.email-template-data")

	// client request
	reqDTO := `{
    "templateCode": "` + testData["email-template-code"] + `",
    "from": { "email": "` + testData["from-email"] + `", "name": "` + testData["from-name"] + `" },
    "to": { "email": "` + testData["to-email"] + `", "name": "` + testData["to-name"] + `" },
    "cc": [
        { "email": "` + testData["cc-email-01"] + `", "name": "` + testData["cc-name-01"] + `" },
        { "email": "` + testData["cc-email-02"] + `", "name": "` + testData["cc-name-02"] + `" }
    ],
    "bcc": [
        { "email": "` + testData["bcc-email-01"] + `", "name": "` + testData["bcc-name-01"] + `" },
		{ "email": "` + testData["bcc-email-02"] + `", "name": "` + testData["bcc-name-02"] + `" }
    ],
    "templateData": {
		"Header.Name": "` + testDataET["header-name"] + `",
		"Body.UserAccount": "` + testDataET["body-user-account"] + `",
		"Body.ActivationURL": "` + testDataET["body-activation-url"] + `",
        "Footer.Name": "` + testDataET["footer-name"] + `"
	},
	"processingType": "` + testData["processing-type"] + `"
}`

	// setup echo
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/email/send", strings.NewReader(reqDTO))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)

	// handler
	handler := ht.NewHandler()
	if err := initialize.LoadAllDatabaseConnection(handler); err != nil {
		t.Errorf("initialize.LoadAllDatabaseConnection: %s", err.Error())
		return
	}

	// set identity (test only)
	token, claims, err := ht.GenerateUserTestToken(handler, t)
	if err != nil {
		t.Errorf("generateUserTestToken: %s", err.Error())
		return
	}
	c.Set("identity.token.jwt", token)
	c.Set("identity.token.jwt.claims", claims)

	// test feature
	email, err := NewFEmail(handler)
	if err != nil {
		t.Errorf("NewFEmail: %s", err.Error())
		return
	}

	if assert.NoError(t, email.SendEmail(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		// save to test-data
		// save result for next test
		viper.Set("test-data.email.email.interface-layer.features.send.response.json", res.Body.String())
		if err := viper.WriteConfig(); err != nil {
			t.Errorf("Error: viper.WriteConfig(), %s", err.Error())
		}
		t.Logf("RESPONSE.Email.SendEmail: %s", res.Body.String())
	}
}
