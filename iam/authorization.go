package iam

import (
	"github.com/alloyzeus/go-azfl/v2/azcore"
	"github.com/alloyzeus/go-azfl/v2/errors"
)

// Used in API call metadata: HTTP header and gRPC call metadata
const (
	AuthorizationMetadataKey    = "Authorization"
	AuthorizationMetadataKeyAlt = "authorization"
)

// Authorization is generally used to provide authorization information
// for call or request. An Authorization is usually obtained from authorization
// token (access token) provided along the request/call.
// TODO: include the application ref if it's using client authentication.
type Authorization struct {
	// DelegateAuthorization is valid when this Authorization is a
	// delegated one. This field holds the authorization information
	// of the delegate.
	DelegateAuthorization *Authorization

	SessionID SessionID

	// Scope, expiry time

	rawToken string
}

var _ azcore.Session[
	SessionIDNum, SessionID,
	TerminalIDNum, TerminalID,
	UserIDNum, UserID,
	AuthorizationSubject,
] = Authorization{}

// newEmptyAuthorization creates a new instance of Authorization without
// any data.
func newEmptyAuthorization() *Authorization {
	return &Authorization{}
}

func (authz Authorization) DelegateSession() azcore.Session[
	SessionIDNum, SessionID,
	TerminalIDNum, TerminalID,
	UserIDNum, UserID,
	AuthorizationSubject,
] {
	return authz.DelegateAuthorization
}

func (authz Authorization) ImpersonatorSession() azcore.Session[
	SessionIDNum, SessionID,
	TerminalIDNum, TerminalID,
	UserIDNum, UserID,
	AuthorizationSubject,
] {
	panic(errors.ErrUnimplemented)
}

func (authz Authorization) ID() SessionID { return authz.SessionID }

func (authz Authorization) Subject() AuthorizationSubject {
	return NewAuthorizationSubject(
		authz.SessionID.terminal, authz.SessionID.terminal.user)
}

func (authz Authorization) IsStaticallyValid() bool {
	return authz.SessionID.IsStaticallyValid()
}

func (authz Authorization) IsNotStaticallyValid() bool {
	return !authz.IsStaticallyValid()
}

// IsTerminal returns true if the authorized terminal is the same as terminalID.
func (authz Authorization) IsTerminal(terminalID TerminalID) bool {
	ctxTerm := authz.SessionID.terminal
	return ctxTerm.IsStaticallyValid() && ctxTerm.EqualsTerminalID(terminalID)
}

// IsUser checks if this authorization is represeting a particular user.
func (authz Authorization) IsUser(userID UserID) bool {
	return authz.clientApplicationIDNum().IsUserAgent() &&
		authz.SessionID.terminal.user.EqualsUserID(userID)
}

// HasUserAsSubject is used to determine if this authorization represents a user.
func (authz Authorization) HasUserAsSubject() bool {
	if authz.clientApplicationIDNum().IsUserAgent() &&
		authz.SessionID.terminal.user.IsStaticallyValid() {
		return true
	}
	return false
}

func (authz Authorization) IsServiceClientContext() bool {
	if authz.clientApplicationIDNum().IsService() &&
		authz.SessionID.terminal.user.IsNotStaticallyValid() {
		return true
	}
	return false
}

func (authz Authorization) UserID() UserID {
	return authz.SessionID.terminal.user
}

func (authz Authorization) TerminalID() TerminalID {
	return authz.SessionID.terminal
}

func (authz Authorization) clientApplicationIDNum() ApplicationIDNum {
	return authz.SessionID.terminal.application.IDNum()
}

// RawToken returns the token where this instance of Authorization
// was parsed from.
func (authz Authorization) RawToken() string {
	return authz.rawToken
}

// AuthorizationSubject provides information about who or what an authorization
// represents.
type AuthorizationSubject struct {
	// terminalID is the ID of the terminal where the action was
	// initiated from.
	terminalID TerminalID
	// userID is the ID of the user who performed the action. This might
	// be empty if the action was performed by non-user-representing agent.
	userID UserID
}

func NewAuthorizationSubject(
	terminalID TerminalID,
	userID UserID,
) AuthorizationSubject {
	return AuthorizationSubject{userID: userID, terminalID: terminalID}
}

var _ azcore.SessionSubject[
	TerminalIDNum, TerminalID, UserIDNum, UserID] = AuthorizationSubject{}

// AZSubject is required by azcore.Subject
func (AuthorizationSubject) AZSessionSubject() {}

// IsRepresentingAUser is required by azcore.Subject
func (actor AuthorizationSubject) IsRepresentingAUser() bool {
	return actor.userID.IsStaticallyValid()
}

// TerminalID is required by azcore.Subject
func (actor AuthorizationSubject) TerminalID() TerminalID { return actor.terminalID }

// UserID is required by azcore.Subject
func (actor AuthorizationSubject) UserID() UserID { return actor.userID }
