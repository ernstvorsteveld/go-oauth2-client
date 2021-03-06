package oauth2

import (
	"github.com/ernstvorsteveld/go-password"
	"github.com/google/uuid"
	"net/url"
)

type ClientType int

const (
	Confidential ClientType = iota
	Public
)

var clientTypes = map[string]ClientType{
	"confidential": Confidential,
	"public":       Public,
}

func (d ClientType) String() string {
	return [...]string{
		"confidential",
		"public"}[d]
}

func ClientTypeValueOf(val string) ClientType {
	clientType := clientTypes[val]
	return clientType
}

type Scope string

type CodeType struct {
	code string
}

type ClientType int

type Client struct {
	clientType           ClientType
	id                   uuid.UUID
	secret               password.Password
	redirect_uris        []url.URL // the URL MAY not contain a fragment and is required and MAY be more than on
	default_redirect_uri url.URL   // use when request does not contain redirect_uri
	scopes               []Scope
}

type TokenEndpointAuthMethodType int

const (
	none TokenEndpointAuthMethodType = iota
	client_secret_post
	client_secret_basic
)

type EmailAddress string

type ClientMetadata struct {
	redirect_uris              []url.URL
	token_endpoint_auth_method TokenEndpointAuthMethodType
	grant_types                []GrantTypeType
	response_types             ResponseType
	client_name                string
	client_uri                 url.URL
	logo_uri                   url.URL
	scope                      []string
	contacts                   []EmailAddress
	tos_uri                    url.URL
	policyt_uri                url.URL
	jwks_uri                   url.URL
	jwks                       url.URL
	software_id                uuid.UUID
	software_version           string
}

type SecondsSinds int64

type ClientInformation struct {
	client_id                uuid.UUID
	client_secret            password.Password
	client_id_issued_at      SecondsSinds
	client_secret_expires_at SecondsSinds
	client_meta_data         ClientMetadata
}

type RegistrationErrorType int

const (
	Invalid_redirect_uri_registration RegistrationErrorType = iota
	Invalid_client_metadata_registration
	Invalid_software_statement
	Uapproved_software_statement
)

var registrationErrorTypes = map[string]RegistrationErrorType{
	"invalid_redirect_uri":    Invalid_redirect_uri_registration,
	"invalid_client_metadata": Invalid_client_metadata_registration,
	"invalid_software":        Invalid_software_statement,
	"uapproved_software":      Uapproved_software_statement,
}

func (r RegistrationErrorType) String() string {
	return [...]string{
		"invalid_redirect_uri",
		"invalid_client_metadata",
		"invalid_software",
		"uapproved_software",
	}[r]
}

func RegistrationErrorTypeValueOf(s string) RegistrationErrorType {
	return registrationErrorTypes[s]
}
