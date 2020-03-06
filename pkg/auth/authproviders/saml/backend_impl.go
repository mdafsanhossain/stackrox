package saml

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
	saml2 "github.com/russellhaering/gosaml2"
	"github.com/stackrox/rox/pkg/auth/authproviders"
	"github.com/stackrox/rox/pkg/auth/authproviders/idputil"
	"github.com/stackrox/rox/pkg/auth/tokens"
	"github.com/stackrox/rox/pkg/grpc/requestinfo"
	"github.com/stackrox/rox/pkg/httputil"
	"github.com/stackrox/rox/pkg/logging"
	"github.com/stackrox/rox/pkg/stringutils"
)

var (
	log = logging.LoggerForModule()
)

type backendImpl struct {
	factory    *factory
	acsURLPath string
	sp         saml2.SAMLServiceProvider
	id         string

	config map[string]string
}

func (p *backendImpl) OnEnable(provider authproviders.Provider) {
	p.factory.RegisterBackend(p)
}

func (p *backendImpl) OnDisable(provider authproviders.Provider) {
	p.factory.UnregisterBackend(p)
}

func (p *backendImpl) loginURL(clientState string) (string, error) {
	doc, err := p.sp.BuildAuthRequestDocument()
	if err != nil {
		return "", errors.Wrap(err, "could not construct auth request")
	}
	authURL, err := p.sp.BuildAuthURLRedirect(idputil.MakeState(p.id, clientState), doc)
	if err != nil {
		return "", errors.Wrap(err, "could not construct auth URL")
	}
	return authURL, nil
}

func newBackend(ctx context.Context, acsURLPath string, id string, uiEndpoints []string, config map[string]string) (*backendImpl, error) {
	if len(uiEndpoints) != 1 {
		return nil, errors.New("SAML requires exactly one UI endpoint")
	}
	p := &backendImpl{
		acsURLPath: acsURLPath,
		id:         id,
	}

	acsURL := &url.URL{
		Scheme: "https",
		Host:   uiEndpoints[0],
		Path:   acsURLPath,
	}
	p.sp.AssertionConsumerServiceURL = acsURL.String()

	spIssuer := config["sp_issuer"]
	if spIssuer == "" {
		return nil, errors.New("no ServiceProvider issuer specified")
	}
	p.sp.ServiceProviderIssuer = spIssuer

	effectiveConfig := map[string]string{
		"sp_issuer": spIssuer,
	}

	if config["idp_metadata_url"] != "" {
		if !stringutils.AllEmpty(config["idp_issuer"], config["idp_cert_pem"], config["idp_sso_url"], config["idp_nameid_format"]) {
			return nil, errors.New("if IdP metadata URL is set, IdP issuer, SSO URL, certificate data and Name/ID format must be left blank")
		}
		if err := configureIDPFromMetadataURL(ctx, &p.sp, config["idp_metadata_url"]); err != nil {
			return nil, errors.Wrap(err, "could not configure auth provider from IdP metadata URL")
		}
		effectiveConfig["idp_metadata_url"] = config["idp_metadata_url"]
	} else {
		if !stringutils.AllNotEmpty(config["idp_issuer"], config["idp_sso_url"], config["idp_cert_pem"]) {
			return nil, errors.New("if IdP metadata URL is not set, IdP issuer, SSO URL, and certificate data must be specified")
		}
		if err := configureIDPFromSettings(&p.sp, config["idp_issuer"], config["idp_sso_url"], config["idp_cert_pem"], config["idp_nameid_format"]); err != nil {
			return nil, errors.Wrap(err, "could not configure auth provider from settings")
		}
		effectiveConfig["idp_issuer"] = config["idp_issuer"]
		effectiveConfig["idp_sso_url"] = config["idp_sso_url"]
		effectiveConfig["idp_cert_pem"] = config["idp_cert_pem"]
		effectiveConfig["idp_nameid_format"] = config["idp_nameid_format"]
	}

	p.config = effectiveConfig

	return p, nil
}

func (p *backendImpl) Config() map[string]string {
	return p.config
}

func (p *backendImpl) consumeSAMLResponse(samlResponse string) (*authproviders.AuthResponse, error) {
	ai, err := p.sp.RetrieveAssertionInfo(samlResponse)
	if err != nil {
		return nil, err
	}

	var expiry time.Time
	if ai.SessionNotOnOrAfter != nil {
		expiry = *ai.SessionNotOnOrAfter
	}

	claim := saml2AssertionInfoToExternalClaim(ai)
	return &authproviders.AuthResponse{
		Claims:     claim,
		Expiration: expiry,
	}, nil
}

func (p *backendImpl) ProcessHTTPRequest(w http.ResponseWriter, r *http.Request) (*authproviders.AuthResponse, string, error) {
	if r.URL.Path != p.acsURLPath {
		return nil, "", httputil.NewError(http.StatusNotFound, "Not Found")
	}
	if r.Method != http.MethodPost {
		return nil, "", httputil.NewError(http.StatusMethodNotAllowed, "Method Not Allowed")
	}

	samlResponse := r.FormValue("SAMLResponse")
	if samlResponse == "" {
		return nil, "", httputil.NewError(http.StatusBadRequest, "no SAML response transmitted")
	}

	authResp, err := p.consumeSAMLResponse(samlResponse)
	if err != nil {
		return nil, "", err
	}

	relayState := r.FormValue("RelayState")
	_, clientState := idputil.SplitState(relayState)

	return authResp, clientState, err
}

func (p *backendImpl) ExchangeToken(ctx context.Context, externalToken, state string) (*authproviders.AuthResponse, string, error) {
	return nil, "", errors.New("not implemented")
}

func (p *backendImpl) RefreshURL() string {
	return ""
}

func (p *backendImpl) LoginURL(clientState string, _ *requestinfo.RequestInfo) string {
	url, err := p.loginURL(clientState)
	if err != nil {
		log.Errorf("could not obtain the login URL: %v", err)
	}
	return url
}

func (p *backendImpl) Validate(ctx context.Context, claims *tokens.Claims) error {
	return nil
}

// Helpers
//////////

func saml2AssertionInfoToExternalClaim(assertionInfo *saml2.AssertionInfo) *tokens.ExternalUserClaim {
	claim := &tokens.ExternalUserClaim{
		UserID: assertionInfo.NameID,
	}
	claim.Attributes = make(map[string][]string)
	claim.Attributes["userid"] = []string{claim.UserID}

	// We store claims as both friendly name and name for easy of use.
	for _, value := range assertionInfo.Values {
		for _, innerValue := range value.Values {
			if value.Name != "" {
				claim.Attributes[value.Name] = append(claim.Attributes[value.Name], innerValue.Value)
			}
			if value.FriendlyName != "" {
				claim.Attributes[value.FriendlyName] = append(claim.Attributes[value.FriendlyName], innerValue.Value)
			}
		}
	}
	return claim
}
