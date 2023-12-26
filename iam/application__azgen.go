package iam

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"slices"
	"strings"

	azcore "github.com/alloyzeus/go-azfl/v2/azcore"
	azid "github.com/alloyzeus/go-azfl/v2/azid"
	errors "github.com/alloyzeus/go-azfl/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the azfl package it is being compiled against.
// A compilation error at this line likely means your copy of the
// azfl package needs to be updated.
var _ = azcore.AZCorePackageIsVersion1

// Reference imports to suppress errors if they are not otherwise used.
var _ = azid.BinDataTypeUnspecified
var _ = errors.ErrUnimplemented
var _ = binary.MaxVarintLen16
var _ = context.Canceled
var _ = rand.Reader
var _ = slices.Equal[[]string]
var _ = strings.Compare

// Entity Application.
//
// An Application is a class of application that is part of or communicates
// with our realm. It's coming in form of web apps, mobile apps, servers
// of services from all domains, worker services, etc.
//
// An instance of Application that has made successful authentication with
// the identity service is called Terminal.

//region ID

// ApplicationID is used to identify
// an instance of entity Application.
type ApplicationID ApplicationIDNum

// NewApplicationID returns a new instance
// of ApplicationID with the provided attribute values.
func NewApplicationID(
	idNum ApplicationIDNum,
) ApplicationID {
	return ApplicationID(idNum)
}

// To ensure that it conforms the interfaces. If any of these is failing,
// there's a bug in the generator.
var _ azid.ID[ApplicationIDNum] = _ApplicationIDZero
var _ azid.BinUnmarshalable = &_ApplicationIDZeroVar
var _ azid.BinFieldUnmarshalable = &_ApplicationIDZeroVar
var _ azid.TextUnmarshalable = &_ApplicationIDZeroVar
var _ azcore.EntityID[ApplicationIDNum] = _ApplicationIDZero
var _ azcore.ValueObjectAssert[ApplicationID] = _ApplicationIDZero

const _ApplicationIDZero = ApplicationID(ApplicationIDNumZero)

var _ApplicationIDZeroVar = _ApplicationIDZero

// ApplicationIDZero returns
// a zero-valued instance of ApplicationID.
func ApplicationIDZero() ApplicationID {
	return _ApplicationIDZero
}

// Clone returns a copy of self.
func (id ApplicationID) Clone() ApplicationID { return id }

// PrimitiveValue returns the value in its primitive type. Prefer to use
// this method instead of casting directly.
func (id ApplicationID) PrimitiveValue() int32 {
	return int32(id)
}

// IDNum returns the scoped identifier of the entity.
func (id ApplicationID) IDNum() ApplicationIDNum {
	return ApplicationIDNum(id)
}

// IDNumPtr returns a pointer to a copy of the id-num if it's considered valid
// otherwise it returns nil.
func (id ApplicationID) IDNumPtr() *ApplicationIDNum {
	if id.IsNotStaticallyValid() {
		return nil
	}
	i := id.IDNum()
	return &i
}

// AZIDNum is required for conformance with azid.ID.
func (id ApplicationID) AZIDNum() ApplicationIDNum {
	return ApplicationIDNum(id)
}

// IsZero is required as ApplicationID is a value-object.
func (id ApplicationID) IsZero() bool {
	return ApplicationIDNum(id) == ApplicationIDNumZero
}

// IsStaticallyValid returns true if this instance is valid as an isolated value
// of ApplicationID.
// It doesn't tell whether it refers to a valid instance of Application.
func (id ApplicationID) IsStaticallyValid() bool {
	return ApplicationIDNum(id).IsStaticallyValid()
}

// IsNotStaticallyValid returns the negation of value returned by IsStaticallyValid.
func (id ApplicationID) IsNotStaticallyValid() bool {
	return !id.IsStaticallyValid()
}

// Equals is required for conformance with azcore.EntityID.
func (id ApplicationID) Equals(other interface{}) bool {
	if x, ok := other.(ApplicationID); ok {
		return x == id
	}
	if x, _ := other.(*ApplicationID); x != nil {
		return *x == id
	}
	return false
}

// Equal is required for conformance with azcore.EntityID.
func (id ApplicationID) Equal(other interface{}) bool {
	return id.Equals(other)
}

// EqualsApplicationID returns true
// if the other value has the same attributes as id.
func (id ApplicationID) EqualsApplicationID(
	other ApplicationID,
) bool {
	return other == id
}

func (id ApplicationID) AZIDBin() []byte {
	b := make([]byte, 4+1)
	b[0] = ApplicationIDNumBinDataType.Byte()
	binary.BigEndian.PutUint32(b[1:], uint32(id))
	return b
}

// ApplicationIDFromAZIDBin creates a new instance of
// ApplicationID from its azid-bin-encoded form.
func ApplicationIDFromAZIDBin(
	input []byte,
) (id ApplicationID, readLen int, err error) {
	typ, err := azid.BinDataTypeFromByte(input[0])
	if err != nil {
		return _ApplicationIDZero, 0,
			errors.ArgMW("input", "type parsing", err)
	}
	if typ != ApplicationIDNumBinDataType {
		return _ApplicationIDZero, 0,
			errors.ArgMsg("input", "type unsupported")
	}

	i, readLen, err := ApplicationIDFromAZIDBinField(input[1:], typ)
	if err != nil {
		return _ApplicationIDZero, 0,
			errors.ArgMW("input", "id-num parsing", err)
	}

	return ApplicationID(i), 1 + readLen, nil
}

// UnmarshalAZIDBin is required for conformance
// with azcore.BinFieldUnmarshalable.
func (id *ApplicationID) UnmarshalAZIDBin(input []byte) (readLen int, err error) {
	i, readLen, err := ApplicationIDFromAZIDBin(input)
	if err == nil {
		*id = i
	}
	return readLen, err
}

func (id ApplicationID) AZIDBinField() ([]byte, azid.BinDataType) {
	return ApplicationIDNum(id).AZIDBinField()
}

func ApplicationIDFromAZIDBinField(
	input []byte, typeHint azid.BinDataType,
) (id ApplicationID, readLen int, err error) {
	idNum, n, err := ApplicationIDNumFromAZIDBinField(input, typeHint)
	if err != nil {
		return _ApplicationIDZero, n, err
	}
	return ApplicationID(idNum), n, nil
}

// UnmarshalAZIDBinField is required for conformance
// with azcore.BinFieldUnmarshalable.
func (id *ApplicationID) UnmarshalAZIDBinField(
	input []byte, typeHint azid.BinDataType,
) (readLen int, err error) {
	i, readLen, err := ApplicationIDFromAZIDBinField(input, typeHint)
	if err == nil {
		*id = i
	}
	return readLen, err
}

const _ApplicationIDAZIDTextPrefix = "KAp0"

// AZIDText is required for conformance
// with azid.ID.
func (id ApplicationID) AZIDText() string {
	if !id.IsStaticallyValid() {
		return ""
	}

	return _ApplicationIDAZIDTextPrefix +
		azid.TextEncode(id.AZIDBin())
}

// ApplicationIDFromAZIDText creates a new instance of
// ApplicationID from its azid-text form.
func ApplicationIDFromAZIDText(input string) (ApplicationID, error) {
	if input == "" {
		return ApplicationIDZero(), nil
	}
	if !strings.HasPrefix(input, _ApplicationIDAZIDTextPrefix) {
		return ApplicationIDZero(),
			errors.ArgMsg("input", "prefix mismatch")
	}
	input = strings.TrimPrefix(input, _ApplicationIDAZIDTextPrefix)
	b, err := azid.TextDecode(input)
	if err != nil {
		return ApplicationIDZero(),
			errors.ArgMW("input", "decoding", err)
	}
	id, _, err := ApplicationIDFromAZIDBin(b)
	if err != nil {
		return ApplicationIDZero(),
			errors.ArgMW("input", "parsing", err)
	}
	return id, nil
}

// UnmarshalAZIDText is required for conformance
// with azid.TextUnmarshalable.
func (id *ApplicationID) UnmarshalAZIDText(input string) error {
	r, err := ApplicationIDFromAZIDText(input)
	if err == nil {
		*id = r
	}
	return err
}

// MarshalText is for compatibility with Go's encoding.TextMarshaler
func (id ApplicationID) MarshalText() ([]byte, error) {
	return []byte(id.AZIDText()), nil
}

// UnmarshalText is for conformance with Go's encoding.TextUnmarshaler
func (id *ApplicationID) UnmarshalText(b []byte) error {
	r, err := ApplicationIDFromAZIDText(string(b))
	if err == nil {
		*id = r
	}
	return err
}

// MarshalJSON makes this type JSON-marshalable.
func (id ApplicationID) MarshalJSON() ([]byte, error) {
	// We assume that there are no symbols in azid-text
	return []byte("\"" + id.AZIDText() + "\""), nil
}

// UnmarshalJSON parses a JSON value.
func (id *ApplicationID) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "" {
		*id = ApplicationIDZero()
		return nil
	}
	i, err := ApplicationIDFromAZIDText(s)
	if err == nil {
		*id = i
	}
	return err
}

// ApplicationIDService abstracts
// ApplicationID-related services.
type ApplicationIDService interface {
	// IsApplicationID is to check if the ref-key is
	// trully registered to system. It does not check whether the instance
	// is active or not.
	IsApplicationIDRegistered(id ApplicationID) bool
}

//endregion

//region IDNum

// ApplicationIDNum is a scoped identifier
// used to identify an instance of entity Application.
type ApplicationIDNum int32

// ApplicationIDNumBinDataType is the type when it is in its
// az-bin-encoded form.
const ApplicationIDNumBinDataType = azid.BinDataTypeInt32

// To ensure that it conforms the interfaces. If any of these is failing,
// there's a bug in the generator.
var _ azid.IDNumMethods = ApplicationIDNumZero
var _ azid.BinFieldUnmarshalable = &_ApplicationIDNumZeroVar
var _ azcore.EntityIDNumMethods = ApplicationIDNumZero
var _ azcore.ValueObjectAssert[ApplicationIDNum] = ApplicationIDNumZero

// ApplicationIDNumIdentifierBitsMask is used to
// extract identifier bits from an instance of ApplicationIDNum.
const ApplicationIDNumIdentifierBitsMask uint32 = 0b_00000011_11111111_11111111_11111111

// ApplicationIDNumZero is the zero value
// for ApplicationIDNum.
const ApplicationIDNumZero = ApplicationIDNum(0)

// _ApplicationIDNumZeroVar is used for testing
// pointer-based interfaces conformance.
var _ApplicationIDNumZeroVar = ApplicationIDNumZero

// ApplicationIDNumFromPrimitiveValue creates an instance
// of ApplicationIDNum from its primitive value.
func ApplicationIDNumFromPrimitiveValue(v int32) ApplicationIDNum {
	return ApplicationIDNum(v)
}

// ApplicationIDNumFromAZIDBinField creates ApplicationIDNum from
// its azid-bin-field form.
func ApplicationIDNumFromAZIDBinField(
	input []byte, typeHint azid.BinDataType,
) (idNum ApplicationIDNum, readLen int, err error) {
	if typeHint != azid.BinDataTypeUnspecified && typeHint != ApplicationIDNumBinDataType {
		return ApplicationIDNum(0), 0,
			errors.ArgMsg("typeHint", "value unsupported")
	}
	i := binary.BigEndian.Uint32(input)
	return ApplicationIDNum(i), 4, nil
}

// PrimitiveValue returns the value in its primitive type. Prefer to use
// this method instead of casting directly.
func (idNum ApplicationIDNum) PrimitiveValue() int32 {
	return int32(idNum)
}

// Clone returns a copy of self.
func (idNum ApplicationIDNum) Clone() ApplicationIDNum { return idNum }

// IsZero is required as ApplicationIDNum is a value-object.
func (idNum ApplicationIDNum) IsZero() bool {
	return idNum == ApplicationIDNumZero
}

// IsStaticallyValid returns true if this instance is valid as an isolated value
// of ApplicationIDNum. It doesn't tell whether it refers to
// a valid instance of Application.
//
// What is considered valid in this context here is that the data
// contained in this instance doesn't break any rule for an instance of
// ApplicationIDNum. Whether the instance is valid in a certain context,
// it requires case-by-case validation which is out of the scope of this
// method.
//
// For example, age 1000 is a valid as an instance of age, but in the context
// of human living age, we can consider it as invalid.
//
// Another example, a ticket has a date of validity for today, but
// after it got checked in to the counter, it turns out that its serial number
// is not registered in the issuer's database. The ticket claims that it's
// valid, but it's considered invalid because it's a fake.
func (idNum ApplicationIDNum) IsStaticallyValid() bool {
	return int32(idNum) > 0 &&
		(uint32(idNum)&ApplicationIDNumIdentifierBitsMask) != 0
}

// IsNotStaticallyValid returns the negation of value returned by IsStaticallyValid.
func (idNum ApplicationIDNum) IsNotStaticallyValid() bool {
	return !idNum.IsStaticallyValid()
}

// Equals is required as ApplicationIDNum is a value-object.
//
// Use EqualsApplicationIDNum method if the other value
// has the same type.
func (idNum ApplicationIDNum) Equals(other interface{}) bool {
	if x, ok := other.(ApplicationIDNum); ok {
		return x == idNum
	}
	if x, _ := other.(*ApplicationIDNum); x != nil {
		return *x == idNum
	}
	return false
}

// Equal is a wrapper for Equals method. It is required for
// compatibility with github.com/google/go-cmp
func (idNum ApplicationIDNum) Equal(other interface{}) bool {
	return idNum.Equals(other)
}

// EqualsApplicationIDNum determines if the other instance is equal
// to this instance.
func (idNum ApplicationIDNum) EqualsApplicationIDNum(
	other ApplicationIDNum,
) bool {
	return idNum == other
}

// AZIDBinField is required for conformance
// with azid.IDNum.
func (idNum ApplicationIDNum) AZIDBinField() ([]byte, azid.BinDataType) {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(idNum))
	return b, ApplicationIDNumBinDataType
}

// UnmarshalAZIDBinField is required for conformance
// with azid.BinFieldUnmarshalable.
func (idNum *ApplicationIDNum) UnmarshalAZIDBinField(
	input []byte, typeHint azid.BinDataType,
) (readLen int, err error) {
	i, readLen, err := ApplicationIDNumFromAZIDBinField(input, typeHint)
	if err == nil {
		*idNum = i
	}
	return readLen, err
}

// Embedded fields
const (
	ApplicationIDNumEmbeddedFieldsMask = 0b_01110000_00000000_00000000_00000000

	ApplicationIDNumFirstPartyMask                         = 0b_01000000_00000000_00000000_00000000
	ApplicationIDNumFirstPartyBits                         = 0b_01000000_00000000_00000000_00000000
	ApplicationIDNumServiceMask                            = 0b_00100000_00000000_00000000_00000000
	ApplicationIDNumServiceBits                            = 0b_00000000_00000000_00000000_00000000
	ApplicationIDNumUserAgentMask                          = 0b_00100000_00000000_00000000_00000000
	ApplicationIDNumUserAgentBits                          = 0b_00100000_00000000_00000000_00000000
	ApplicationIDNumUserAgentAuthorizationPublicMask       = 0b_00110000_00000000_00000000_00000000
	ApplicationIDNumUserAgentAuthorizationPublicBits       = 0b_00100000_00000000_00000000_00000000
	ApplicationIDNumUserAgentAuthorizationConfidentialMask = 0b_00110000_00000000_00000000_00000000
	ApplicationIDNumUserAgentAuthorizationConfidentialBits = 0b_00110000_00000000_00000000_00000000
)

// IsFirstParty returns true if
// the Application instance this ApplicationIDNum is for
// is a FirstParty Application.
//
// FirstParty indicates that the application is a first-party
// application, i.e., the application was provided by the realm.
//
// First-party applications, in contrast to third party
// applications, are those of official applications, officially
// supported public applications, and internal-use only
// applications.
//
// First-party applications could be in various forms, e.g.,
// official mobile app, web app, or system dashboard.
func (idNum ApplicationIDNum) IsFirstParty() bool {
	return idNum.IsStaticallyValid() &&
		idNum.HasFirstPartyBits()
}

// HasFirstPartyBits is only checking the bits
// without validating other information.
func (idNum ApplicationIDNum) HasFirstPartyBits() bool {
	return (uint32(idNum) &
		ApplicationIDNumFirstPartyMask) ==
		ApplicationIDNumFirstPartyBits
}

// IsService returns true if
// the Application instance this ApplicationIDNum is for
// is a Service Application.
//
// A service application is an application which does not
// represent user. Note that this is different from, e.g.,
// bot, where it's a specialization of User; a bot is a User.
//
// All service applications are confidential OAuth 2.0
// clients (RFC6749 section 2.1).
func (idNum ApplicationIDNum) IsService() bool {
	return idNum.IsStaticallyValid() &&
		idNum.HasServiceBits()
}

// HasServiceBits is only checking the bits
// without validating other information.
func (idNum ApplicationIDNum) HasServiceBits() bool {
	return (uint32(idNum) &
		ApplicationIDNumServiceMask) ==
		ApplicationIDNumServiceBits
}

// IsUserAgent returns true if
// the Application instance this ApplicationIDNum is for
// is a UserAgent Application.
//
// A user-agent application is an application which represents
// the user it got authorized by.
//
// There are two types of user agent based on the flow they use
// to obtain authorization: public and confidential.
// These types align with OAuth 2.0 client types (RFC6749
// section 2.1).
func (idNum ApplicationIDNum) IsUserAgent() bool {
	return idNum.IsStaticallyValid() &&
		idNum.HasUserAgentBits()
}

// HasUserAgentBits is only checking the bits
// without validating other information.
func (idNum ApplicationIDNum) HasUserAgentBits() bool {
	return (uint32(idNum) &
		ApplicationIDNumUserAgentMask) ==
		ApplicationIDNumUserAgentBits
}

// IsUserAgentAuthorizationPublic returns true if
// the Application instance this ApplicationIDNum is for
// is a UserAgentAuthorizationPublic Application.
//
// A public-authorization user-agent application is an
// application which can be used by users to authenticate
// themselves to the system. The application will automatically
// receive authorization upon successful authentication of
// the user.
//
// A public-authorization user-agent credentials should
// never be assumed to be secure, i.e., it can not securely
// authenticate itself. An authorized application of this
// kind has no strong identity. Access control checks should
// be focused on testing the user's claims and less of the
// application's claims.
func (idNum ApplicationIDNum) IsUserAgentAuthorizationPublic() bool {
	return idNum.IsStaticallyValid() &&
		idNum.HasUserAgentAuthorizationPublicBits()
}

// HasUserAgentAuthorizationPublicBits is only checking the bits
// without validating other information.
func (idNum ApplicationIDNum) HasUserAgentAuthorizationPublicBits() bool {
	return (uint32(idNum) &
		ApplicationIDNumUserAgentAuthorizationPublicMask) ==
		ApplicationIDNumUserAgentAuthorizationPublicBits
}

// IsUserAgentAuthorizationConfidential returns true if
// the Application instance this ApplicationIDNum is for
// is a UserAgentAuthorizationConfidential Application.
//
// A confidential-authorization user-agent application
// is a user-agent application which could receive authorization
// from a consenting user through 3-legged authorization flow.
// A confidential user-agent can not be used for directly
// performing user authentication.
func (idNum ApplicationIDNum) IsUserAgentAuthorizationConfidential() bool {
	return idNum.IsStaticallyValid() &&
		idNum.HasUserAgentAuthorizationConfidentialBits()
}

// HasUserAgentAuthorizationConfidentialBits is only checking the bits
// without validating other information.
func (idNum ApplicationIDNum) HasUserAgentAuthorizationConfidentialBits() bool {
	return (uint32(idNum) &
		ApplicationIDNumUserAgentAuthorizationConfidentialMask) ==
		ApplicationIDNumUserAgentAuthorizationConfidentialBits
}

//endregion

//region AttrSet

// ApplicationAttrSet is a value-object which contains
// entity's built-in attributes.
type ApplicationAttrSet struct {
	// AZxEntityAttrSetBase

	displayName       string
	secret            string
	platformType      string
	requiredScopes    []string
	oAuth2RedirectURI []string
}

// NewApplicationAttrSet returns a new instance
// of ApplicationAttrSet with the provided attribute values.
func NewApplicationAttrSet(
	displayName string,
	secret string,
	platformType string,
	requiredScopes []string,
	oAuth2RedirectURI []string,
) ApplicationAttrSet {
	return ApplicationAttrSet{
		displayName:       displayName,
		secret:            secret,
		platformType:      platformType,
		requiredScopes:    slices.Clone(requiredScopes),
		oAuth2RedirectURI: slices.Clone(oAuth2RedirectURI),
	}
}

var _ azcore.ValueObjectAssert[ApplicationAttrSet] = ApplicationAttrSet{}

// Clone returns a copy of ApplicationAttrSet
func (attrs ApplicationAttrSet) Clone() ApplicationAttrSet {
	return ApplicationAttrSet{
		displayName:       attrs.displayName,
		secret:            attrs.secret,
		platformType:      attrs.platformType,
		requiredScopes:    slices.Clone(attrs.requiredScopes),
		oAuth2RedirectURI: slices.Clone(attrs.oAuth2RedirectURI),
	}
}

// DisplayName returns instance's DisplayName value.
func (attrs ApplicationAttrSet) DisplayName() string {
	return attrs.displayName
}

// WithDisplayName returns a copy
// of ApplicationAttrSet
// with its displayName attribute set to the provided value.
func (attrs ApplicationAttrSet) WithDisplayName(
	displayName string,
) ApplicationAttrSet {
	return ApplicationAttrSet{
		displayName:       displayName,
		secret:            attrs.secret,
		platformType:      attrs.platformType,
		requiredScopes:    slices.Clone(attrs.requiredScopes),
		oAuth2RedirectURI: slices.Clone(attrs.oAuth2RedirectURI),
	}
}

// Secret returns instance's Secret value.
func (attrs ApplicationAttrSet) Secret() string {
	return attrs.secret
}

// WithSecret returns a copy
// of ApplicationAttrSet
// with its secret attribute set to the provided value.
func (attrs ApplicationAttrSet) WithSecret(
	secret string,
) ApplicationAttrSet {
	return ApplicationAttrSet{
		displayName:       attrs.displayName,
		secret:            secret,
		platformType:      attrs.platformType,
		requiredScopes:    slices.Clone(attrs.requiredScopes),
		oAuth2RedirectURI: slices.Clone(attrs.oAuth2RedirectURI),
	}
}

// PlatformType returns instance's PlatformType value.
func (attrs ApplicationAttrSet) PlatformType() string {
	return attrs.platformType
}

// WithPlatformType returns a copy
// of ApplicationAttrSet
// with its platformType attribute set to the provided value.
func (attrs ApplicationAttrSet) WithPlatformType(
	platformType string,
) ApplicationAttrSet {
	return ApplicationAttrSet{
		displayName:       attrs.displayName,
		secret:            attrs.secret,
		platformType:      platformType,
		requiredScopes:    slices.Clone(attrs.requiredScopes),
		oAuth2RedirectURI: slices.Clone(attrs.oAuth2RedirectURI),
	}
}

// RequiredScopes returns instance's RequiredScopes value.
func (attrs ApplicationAttrSet) RequiredScopes() []string {
	return attrs.requiredScopes
}

// WithRequiredScopes returns a copy
// of ApplicationAttrSet
// with its requiredScopes attribute set to the provided value.
func (attrs ApplicationAttrSet) WithRequiredScopes(
	requiredScopes []string,
) ApplicationAttrSet {
	return ApplicationAttrSet{
		displayName:       attrs.displayName,
		secret:            attrs.secret,
		platformType:      attrs.platformType,
		requiredScopes:    slices.Clone(requiredScopes),
		oAuth2RedirectURI: slices.Clone(attrs.oAuth2RedirectURI),
	}
}

// OAuth2RedirectURI returns instance's OAuth2RedirectURI value.
func (attrs ApplicationAttrSet) OAuth2RedirectURI() []string {
	return attrs.oAuth2RedirectURI
}

// WithOAuth2RedirectURI returns a copy
// of ApplicationAttrSet
// with its oAuth2RedirectURI attribute set to the provided value.
func (attrs ApplicationAttrSet) WithOAuth2RedirectURI(
	oAuth2RedirectURI []string,
) ApplicationAttrSet {
	return ApplicationAttrSet{
		displayName:       attrs.displayName,
		secret:            attrs.secret,
		platformType:      attrs.platformType,
		requiredScopes:    slices.Clone(attrs.requiredScopes),
		oAuth2RedirectURI: slices.Clone(oAuth2RedirectURI),
	}
}

func (attrs ApplicationAttrSet) Equals(
	other interface{},
) bool {
	if x, ok := other.(ApplicationAttrSet); ok {
		return attrs.EqualsApplicationAttrSetPtr(&x)
	}
	if x, _ := other.(*ApplicationAttrSet); x != nil {
		return attrs.EqualsApplicationAttrSetPtr(x)
	}
	return false
}

func (attrs *ApplicationAttrSet) EqualsApplicationAttrSetPtr(
	other *ApplicationAttrSet,
) bool {
	return attrs == other || (other != nil &&
		attrs.displayName == other.displayName &&
		attrs.secret == other.secret &&
		attrs.platformType == other.platformType &&
		slices.Equal(attrs.requiredScopes, other.requiredScopes) &&
		slices.Equal(attrs.oAuth2RedirectURI, other.oAuth2RedirectURI))
}

//endregion

//region Instance

// ApplicationInstanceService is a service which
// provides methods to manipulate an instance of Application.
type ApplicationInstanceService interface {
	ApplicationInstanceStateService
}

// ApplicationInstanceStateService is a service which
// provides access to instances metadata.
type ApplicationInstanceStateService interface {
	// GetApplicationInstanceState checks if the provided
	// ref-key is valid and whether the instance is deleted.
	//
	// This method returns nil if the id is not referencing to any valid
	// instance.
	GetApplicationInstanceState(
		ctx context.Context,
		id ApplicationID,
	) (*ApplicationInstanceState, error)
}

// ApplicationInstanceState holds information about
// an instance of Application.
type ApplicationInstanceState struct {
	RevisionNumber_ int32

	// Deletion_ holds information about the deletion of the instance. If
	// the instance has not been deleted, this field value will be nil.
	Deletion_ *ApplicationDeletionState
}

var _ azcore.EntityInstanceInfo[
	int32, ApplicationDeletionState,
] = ApplicationInstanceState{}
var _ azcore.ValueObjectAssert[ApplicationDeletionState] = ApplicationDeletionState{}

// ApplicationInstanceStateZero returns an instance of
// ApplicationInstanceState with attributes set their respective zero
// value.
func ApplicationInstanceStateZero() ApplicationInstanceState {
	return ApplicationInstanceState{}
}

func (instInfo ApplicationInstanceState) Clone() ApplicationInstanceState {
	if instInfo.Deletion_ != nil {
		cp := instInfo
		delInfo := cp.Deletion_.Clone()
		cp.Deletion_ = &delInfo
		return cp
	}
	// Already a copy and there's no shared underlying data instance
	return instInfo
}

func (instInfo ApplicationInstanceState) RevisionNumber() int32 { return instInfo.RevisionNumber_ }
func (instInfo ApplicationInstanceState) Deletion() *ApplicationDeletionState {
	return instInfo.Deletion_
}

// IsActive returns true if the instance is considered as active.
func (instInfo ApplicationInstanceState) IsActive() bool {
	// Note: we will check other flags in the future, but that's said,
	// deleted instance is considered inactive.
	return !instInfo.IsDeleted()
}

// IsDeleted returns true if the instance was deleted.
func (instInfo ApplicationInstanceState) IsDeleted() bool {
	if delInfo := instInfo.Deletion(); delInfo != nil {
		return delInfo.Deleted()
	}
	return false
}

//----

// ApplicationDeletionState holds information about
// the deletion of an instance if the instance has been deleted.
type ApplicationDeletionState struct {
	Deleted_ bool
}

var _ azcore.EntityDeletionInfo = ApplicationDeletionState{}
var _ azcore.ValueObjectAssert[ApplicationDeletionState] = ApplicationDeletionState{}

func (instDelInfo ApplicationDeletionState) Clone() ApplicationDeletionState {
	// Already a copy and there's no shared underlying data instance
	return instDelInfo
}

func (instDelInfo ApplicationDeletionState) Deleted() bool {
	return instDelInfo.Deleted_
}

//----

// ApplicationInstanceServiceInternal is a service which provides
// methods for manipulating instances of Application. Declared for
// internal use within a process, this interface contains methods that
// available to be called from another part of a process.
type ApplicationInstanceServiceInternal interface {
	CreateApplicationInternal(
		ctx context.Context,
		input ApplicationCreationParams,
	) (id ApplicationID, initialState ApplicationInstanceState, err error)

	// DeleteApplicationInternal deletes an instance of
	// Application entity based identfied by idOfInstToDel.
	// The returned justDeleted will have the value of
	// true if this particular call resulted the deletion of the instance and
	// it will have the value of false of subsequent calls to this method.
	DeleteApplicationInternal(
		ctx context.Context,
		idOfInstToDel ApplicationID,
		input ApplicationDeletionParams,
	) (justDeleted bool, currentState ApplicationInstanceState, err error)
}

// ApplicationCreationParams contains data to be passed
// as an argument when invoking the method CreateApplicationInternal
// of ApplicationInstanceServiceInternal.
type ApplicationCreationParams struct {
}

// ApplicationDeletionParams contains data to be passed
// as an argument when invoking the method DeleteApplicationInternal
// of ApplicationInstanceServiceInternal.
type ApplicationDeletionParams struct {
}

//endregion

//region Service

// ApplicationService provides a contract
// for methods related to entity Application.
type ApplicationService interface {
	// AZxEntityService

	ApplicationIDService
	ApplicationInstanceService
	ApplicationInstanceStateService
}

// ApplicationServiceClient is the interface for
// clients of ApplicationService.
type ApplicationServiceClient interface {
	ApplicationService
}

//endregion
