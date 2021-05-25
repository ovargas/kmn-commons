package authorization

import "sync"

type IIdentity interface {
	Name() *string
	Domain() *string
	ClientID() *string
	AuthenticationType() string
	IsAuthenticated() bool
}

type IPrincipal interface {
	IsInRole(role string) bool
	Identity() IIdentity
}

type Principal struct {
	identity IIdentity
	roles    []string
}

func (p Principal) IsInRole(role string) bool {
	for _, r := range p.roles {
		if r == role {
			return true
		}
	}

	return false
}

func (p Principal) Identity() IIdentity {
	return p.identity
}

func NewPrincipal(roles []string, identity IIdentity) IPrincipal {
	return Principal{
		identity: identity,
		roles:    roles,
	}
}

var (
	once      sync.Once
	anonymous IPrincipal
)

func Anonymous() IPrincipal {
	once.Do(func() {
		anonymous = Principal{
			roles:    []string{},
			identity: anonymousIdentity(),
		}
	})
	return anonymous
}
