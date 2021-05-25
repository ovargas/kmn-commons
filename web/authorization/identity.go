package authorization

const (
	AuthenticationTypeAnonymous = "ANONYMOUS"
	AuthenticationTypeUser      = "FULL_AUTHENTICATED_USER"
	AuthenticationTypeClient    = "CLIENT_CREDENTIALS"
)

type Identity struct {
	name               *string
	domain             *string
	clientId           *string
	authenticationType string
	authenticated      bool
}

func (i Identity) Name() *string {
	return i.name
}

func (i Identity) Domain() *string {
	return i.domain
}

func (i Identity) ClientID() *string {
	return i.clientId
}

func (i Identity) AuthenticationType() string {
	return i.authenticationType
}

func (i Identity) IsAuthenticated() bool {
	return i.authenticated
}

func anonymousIdentity() Identity {
	return Identity{
		authenticationType: AuthenticationTypeAnonymous,
	}
}

func NewUserIdentity(name string, domain string, clientId string) Identity {
	return Identity{
		name:               &name,
		domain:             &domain,
		clientId:           &clientId,
		authenticationType: AuthenticationTypeUser,
	}
}

func NewClientCredentialsIdentity(clientId string) Identity {
	return Identity{
		clientId:           &clientId,
		authenticationType: AuthenticationTypeClient,
	}
}
