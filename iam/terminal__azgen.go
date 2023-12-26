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

// Adjunct-entity Terminal of Application, User.
//
// A Terminal is an authorized instance of application. A new instance of
// Terminal will be created upon sucessful authorization-authentication.
//
// As long the the authorization is still valid, a terminal
// could be used to access the service within the limit of the granted
// authorization. An authorization of a Terminal will become invalid when
// it is expired, or revoked by the user, if the terminal is associated
// to a user, or those who have the permission to revoke the authorization.
//
// After a Terminal authorization is invalid, their user must re-authorize
// their instance of Application if they wish to continue using their
// instance of Application to access the service. Re-authorization of an
// instance of an Application will generate a new Terminal. A
// de-authorized Terminal is no longer usable.
//
// A sucessful authorization will generate both a new instance of
// Terminal and an instance of initial Session.
//
// A Terminal lifecycle determines whether Sessions can be renewed.

//region ID

// TerminalID is used to identify
// an instance of adjunct entity Terminal system-wide.
type TerminalID struct {
	application ApplicationID
	user UserID
	idNum TerminalIDNum // Differentiator
}

// The total number of fields in the struct.
const _TerminalIDFieldCount = 2 + 1

// NewTerminalID returns a new instance
// of TerminalID with the provided attribute values.
func NewTerminalID(
	application ApplicationID,
	user UserID,
	idNum TerminalIDNum, // Differentiator
) TerminalID {
	return TerminalID{
		application: application,
		user: user,
		idNum: idNum,
	}
}

// To ensure that it conforms the interfaces. If any of these is failing,
// there's a bug in the generator.
var _ azid.ID[TerminalIDNum] = _TerminalIDZero
var _ azid.BinFieldUnmarshalable = &_TerminalIDZero
var _ azid.TextUnmarshalable = &_TerminalIDZero
var _ azcore.AdjunctEntityID[TerminalIDNum] = _TerminalIDZero
var _ azcore.ValueObjectAssert[TerminalID] = _TerminalIDZero
var _ azcore.TerminalID[TerminalIDNum] = _TerminalIDZero

var _TerminalIDZero = TerminalID{
	application: ApplicationIDZero(),
	user: UserIDZero(),
	idNum: TerminalIDNumZero,
}

// TerminalIDZero returns
// a zero-valued instance of TerminalID.
func TerminalIDZero() TerminalID {
	return _TerminalIDZero
}

// Clone returns a copy of self.
func (id TerminalID) Clone() TerminalID { return id }

// IDNum returns the scoped identifier of the entity.
func (id TerminalID) IDNum() TerminalIDNum {
	return id.idNum
}

// IDNumPrimitiveValue returns the value in its primitive type. Prefer to use
// this method instead of casting directly.
func (id TerminalID) IDNumPrimitiveValue() int64 {
	return int64(id.idNum)
}

// IDNumPtr returns a pointer to a copy of the id-num if it's considered valid
// otherwise it returns nil.
func (id TerminalID) IDNumPtr() *TerminalIDNum {
	if id.IsNotStaticallyValid() {
		return nil
	}
	i := id.IDNum()
	return &i
}

// AZIDNum is required for conformance with azid.ID.
func (id TerminalID) AZIDNum() TerminalIDNum {
	return id.idNum
}

// TerminalIDNum is required for conformance
// with azcore.TerminalID.
func (id TerminalID) TerminalIDNum() TerminalIDNum {
	return id.idNum
}

// IsZero is required as TerminalID is a value-object.
func (id TerminalID) IsZero() bool {
	return id.idNum == TerminalIDNumZero
}

// IsStaticallyValid returns true if this instance is valid as an isolated value
// of TerminalID.
// It doesn't tell whether it refers to a valid instance of Terminal.
func (id TerminalID) IsStaticallyValid() bool {
	return id.idNum.IsStaticallyValid()
}

// IsNotStaticallyValid returns the negation of value returned by IsStaticallyValid.
func (id TerminalID) IsNotStaticallyValid() bool {
	return !id.IsStaticallyValid()
}

// Equals is required for conformance with azcore.AdjunctEntityID.
func (id TerminalID) Equals(other interface{}) bool {
	if x, ok := other.(TerminalID); ok {
		return id.application.EqualsApplicationID(x.application) &&
			id.user.EqualsUserID(x.user) &&
			id.idNum == x.idNum
	}
	if x, _ := other.(*TerminalID); x != nil {
		return id.application.EqualsApplicationID(x.application) &&
			id.user.EqualsUserID(x.user) &&
			id.idNum == x.idNum
	}
	return false
}

// Equal is required for conformance with azcore.AdjunctEntityID.
func (id TerminalID) Equal(other interface{}) bool {
	return id.Equals(other)
}

// EqualsTerminalID returns true
// if the other value has the same attributes as id.
func (id TerminalID) EqualsTerminalID(
	other TerminalID,
) bool {
	return id.application.EqualsApplicationID(other.application) &&
		id.user.EqualsUserID(other.user) &&
		id.idNum == other.idNum
}

// AZIDBin is required for conformance
// with azid.ID.
func (id TerminalID) AZIDBin() []byte {
	data, typ := id.AZIDBinField()
	out := []byte{typ.Byte()}
	return append(out, data...)
}

// TerminalIDFromAZIDBin creates a new instance of
// TerminalID from its azid-bin-encoded form.
func TerminalIDFromAZIDBin(
	input []byte,
) (id TerminalID, readLen int, err error) {
	typ, err := azid.BinDataTypeFromByte(input[0])
	if err != nil {
		return _TerminalIDZero, 0,
			errors.ArgMW("input", "type parsing", err)
	}
	if typ != azid.BinDataTypeArray {
		return _TerminalIDZero, 0,
			errors.ArgMsg("input", "type unsupported")
	}

	id, readLen, err = TerminalIDFromAZIDBinField(input[1:], typ)
	return id, readLen + 1, err
}

// AZIDBinField is required for conformance
// with azid.ID.
func (id TerminalID) AZIDBinField() ([]byte, azid.BinDataType) {
	var typesBytes []byte
	var dataBytes []byte
	var fieldBytes []byte
	var fieldType azid.BinDataType

	fieldBytes, fieldType = id.application.AZIDBinField()
	typesBytes = append(typesBytes, fieldType.Byte())
	dataBytes = append(dataBytes, fieldBytes...)

	fieldBytes, fieldType = id.user.AZIDBinField()
	typesBytes = append(typesBytes, fieldType.Byte())
	dataBytes = append(dataBytes, fieldBytes...)

	fieldBytes, fieldType = id.idNum.AZIDBinField()
	typesBytes = append(typesBytes, fieldType.Byte())
	dataBytes = append(dataBytes, fieldBytes...)

	var out = []byte{byte(len(typesBytes))}
	out = append(out, typesBytes...)
	out = append(out, dataBytes...)
	return out, azid.BinDataTypeArray
}

// TerminalIDFromAZIDBinField creates TerminalID from
// its azid-bin-encoded field form.
func TerminalIDFromAZIDBinField(
	input []byte, typeHint azid.BinDataType,
) (id TerminalID, readLen int, err error) {
	if typeHint != azid.BinDataTypeArray {
		return TerminalIDZero(), 0,
			errors.ArgMsg("typeHint", "value unsupported")
	}

	arrayLen := int(input[0])
	if arrayLen != _TerminalIDFieldCount {
		return TerminalIDZero(), 0,
			errors.ArgMsg("input", "field count mismatch")
	}

	typeCursor := 1
	dataCursor := typeCursor + arrayLen

	var fieldType azid.BinDataType

	fieldType, err = azid.BinDataTypeFromByte(input[typeCursor])
	if err != nil {
		return TerminalIDZero(), 0,
			errors.ArgMW("input", "host application ref-key type parsing", err)
	}
	typeCursor++
	applicationID, readLen, err := ApplicationIDFromAZIDBinField(
		input[dataCursor:], fieldType)
	if err != nil {
		return TerminalIDZero(), 0,
			errors.ArgMW("input", "host application ref-key data parsing", err)
	}
	dataCursor += readLen

	fieldType, err = azid.BinDataTypeFromByte(input[typeCursor])
	if err != nil {
		return TerminalIDZero(), 0,
			errors.ArgMW("input", "host user ref-key type parsing", err)
	}
	typeCursor++
	userID, readLen, err := UserIDFromAZIDBinField(
		input[dataCursor:], fieldType)
	if err != nil {
		return TerminalIDZero(), 0,
			errors.ArgMW("input", "host user ref-key data parsing", err)
	}
	dataCursor += readLen

	fieldType, err = azid.BinDataTypeFromByte(input[typeCursor])
	if err != nil {
		return TerminalIDZero(), 0,
			errors.ArgMW("input", "id-num type parsing", err)
	}
	typeCursor++
	idNum, readLen, err := TerminalIDNumFromAZIDBinField(
		input[dataCursor:], fieldType)
	if err != nil {
		return TerminalIDZero(), 0,
			errors.ArgMW("input", "id-num data parsing", err)
	}
	dataCursor += readLen

	return TerminalID{
		application: applicationID,
		user: userID,
		idNum: idNum,
	}, dataCursor, nil
}

// UnmarshalAZIDBinField is required for conformance
// with azcore.BinFieldUnmarshalable.
func (id *TerminalID) UnmarshalAZIDBinField(
	input []byte, typeHint azid.BinDataType,
) (readLen int, err error) {
	i, readLen, err := TerminalIDFromAZIDBinField(input, typeHint)
	if err == nil {
		*id = i
	}
	return readLen, err
}

const _TerminalIDAZIDTextPrefix = "KTx0"

// AZIDText is required for conformance
// with azid.ID.
func (id TerminalID) AZIDText() string {
	if !id.IsStaticallyValid() {
		return ""
	}

	return _TerminalIDAZIDTextPrefix +
		azid.TextEncode(id.AZIDBin())
}

// TerminalIDFromAZIDText creates a new instance of
// TerminalID from its azid-text form.
func TerminalIDFromAZIDText(input string) (TerminalID, error) {
	if input == "" {
		return TerminalIDZero(), nil
	}
	if !strings.HasPrefix(input, _TerminalIDAZIDTextPrefix) {
		return TerminalIDZero(),
			errors.ArgMsg("input", "prefix mismatch")
	}
	input = strings.TrimPrefix(input, _TerminalIDAZIDTextPrefix)
	b, err := azid.TextDecode(input)
	if err != nil {
		return TerminalIDZero(),
			errors.ArgMW("input", "decoding", err)
	}
	id, _, err := TerminalIDFromAZIDBin(b)
	if err != nil {
		return TerminalIDZero(),
			errors.ArgMW("input", "parsing", err)
	}
	return id, nil
}

// UnmarshalAZIDText is required for conformance
// with azid.TextUnmarshalable.
func (id *TerminalID) UnmarshalAZIDText(input string) error {
	r, err := TerminalIDFromAZIDText(input)
	if err == nil {
		*id = r
	}
	return err
}

// MarshalText is for compatibility with Go's encoding.TextMarshaler
func (id TerminalID) MarshalText() ([]byte, error) {
	return []byte(id.AZIDText()), nil
}

// UnmarshalText is for conformance with Go's encoding.TextUnmarshaler
func (id *TerminalID) UnmarshalText(b []byte) error {
	r, err := TerminalIDFromAZIDText(string(b))
	if err == nil {
		*id = r
	}
	return err
}

// MarshalJSON makes this type JSON-marshalable.
func (id TerminalID) MarshalJSON() ([]byte, error) {
	// We assume that there are no symbols in azid-text
	return []byte("\"" + id.AZIDText() + "\""), nil
}

// UnmarshalJSON parses a JSON value.
func (id *TerminalID) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "" {
		*id = TerminalIDZero()
		return nil
	}
	i, err := TerminalIDFromAZIDText(s)
	if err == nil {
		*id = i
	}
	return err
}

// Application returns instance's Application value.
func (id TerminalID) Application() ApplicationID {
	return id.application
}

// ApplicationPtr returns a pointer to a copy of
// ApplicationID if it's considered valid.
func (id TerminalID) ApplicationPtr() *ApplicationID {
	if id.application.IsStaticallyValid() {
		rk := id.application
		return &rk
	}
	return nil
}

// WithApplication returns a copy
// of TerminalID
// with its application attribute set to the provided value.
func (id TerminalID) WithApplication(
	application ApplicationID,
) TerminalID {
	return TerminalID{
		application: application,
		user: id.user,
		idNum: id.idNum,
	}
}

// User returns instance's User value.
func (id TerminalID) User() UserID {
	return id.user
}

// UserPtr returns a pointer to a copy of
// UserID if it's considered valid.
func (id TerminalID) UserPtr() *UserID {
	if id.user.IsStaticallyValid() {
		rk := id.user
		return &rk
	}
	return nil
}

// WithUser returns a copy
// of TerminalID
// with its user attribute set to the provided value.
func (id TerminalID) WithUser(
	user UserID,
) TerminalID {
	return TerminalID{
		application: id.application,
		user: user,
		idNum: id.idNum,
	}
}

//endregion

//region IDNum

// TerminalIDNum is a scoped identifier
// used to identify an instance of adjunct entity Terminal
// scoped within its host entity(s).
type TerminalIDNum int64

// TerminalIDNumBinDataType is the type when it is in its
// az-bin-encoded form.
const TerminalIDNumBinDataType = azid.BinDataTypeInt64

// To ensure that it conforms the interfaces. If any of these is failing,
// there's a bug in the generator.
var _ azid.IDNumMethods = TerminalIDNumZero
var _ azid.BinFieldUnmarshalable = &_TerminalIDNumZeroVar
var _ azcore.AdjunctEntityIDNumMethods = TerminalIDNumZero
var _ azcore.ValueObjectAssert[TerminalIDNum] = TerminalIDNumZero
var _ azcore.TerminalIDNumMethods = TerminalIDNumZero

// TerminalIDNumIdentifierBitsMask is used to
// extract identifier bits from an instance of TerminalIDNum.
const TerminalIDNumIdentifierBitsMask uint64 =
	0b_00000000_11111111_11111111_11111111_11111111_11111111_11111111_11111111

// TerminalIDNumZero is the zero value for TerminalIDNum.
const TerminalIDNumZero =
	TerminalIDNum(0)

// _TerminalIDNumZeroVar is used for testing
// pointer-based interfaces conformance.
var _TerminalIDNumZeroVar =
	TerminalIDNumZero

// TerminalIDNumFromPrimitiveValue creates an instance
// of TerminalIDNum from its primitive value.
func TerminalIDNumFromPrimitiveValue(v int64) TerminalIDNum {
	return TerminalIDNum(v)
}

// TerminalIDNumFromAZIDBinField creates TerminalIDNum from
// its azid-bin-encoded form.
func TerminalIDNumFromAZIDBinField(
	input []byte, typeHint azid.BinDataType,
) (idNum TerminalIDNum, readLen int, err error) {
	if typeHint != azid.BinDataTypeUnspecified && typeHint != TerminalIDNumBinDataType {
		return TerminalIDNum(0), 0,
			errors.ArgMsg("typeHint", "value unsupported")
	}
	i := binary.BigEndian.Uint64(input)
	return TerminalIDNum(i), 8, nil
}

// PrimitiveValue returns the value in its primitive type. Prefer to use
// this method instead of casting directly.
func (idNum TerminalIDNum) PrimitiveValue() int64 {
	return int64(idNum)
}

// Clone returns a copy of self.
func (idNum TerminalIDNum) Clone() TerminalIDNum { return idNum }

// IsZero is required as TerminalIDNum is a value-object.
func (idNum TerminalIDNum) IsZero() bool {
	return idNum == TerminalIDNumZero
}

// IsStaticallyValid returns true if this instance is valid as an isolated value
// of TerminalIDNum. It doesn't tell whether it refers to
// a valid instance of Terminal.
func (idNum TerminalIDNum) IsStaticallyValid() bool {
	return int64(idNum) > 0 &&
		(uint64(idNum) & TerminalIDNumIdentifierBitsMask) != 0
}

// IsNotStaticallyValid returns the negation of value returned by IsStaticallyValid.
func (idNum TerminalIDNum) IsNotStaticallyValid() bool {
	return !idNum.IsStaticallyValid()
}

// Equals is required as TerminalIDNum is a value-object.
//
// Use EqualsTerminalIDNum method if the other value
// has the same type.
func (idNum TerminalIDNum) Equals(other interface{}) bool {
	if x, ok := other.(TerminalIDNum); ok {
		return x == idNum
	}
	if x, _ := other.(*TerminalIDNum); x != nil {
		return *x == idNum
	}
	return false
}

// Equal is a wrapper for Equals method. It is required for
// compatibility with github.com/google/go-cmp
func (idNum TerminalIDNum) Equal(other interface{}) bool {
	return idNum.Equals(other)
}

// EqualsTerminalIDNum determines if the other instance
// is equal to this instance.
func (idNum TerminalIDNum) EqualsTerminalIDNum(
	other TerminalIDNum,
) bool {
	return idNum == other
}

// AZIDBinField is required for conformance
// with azid.IDNum.
func (idNum TerminalIDNum) AZIDBinField() ([]byte, azid.BinDataType) {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(idNum))
	return b, TerminalIDNumBinDataType
}

// UnmarshalAZIDBinField is required for conformance
// with azid.BinFieldUnmarshalable.
func (idNum *TerminalIDNum) UnmarshalAZIDBinField(
	input []byte, typeHint azid.BinDataType,
) (readLen int, err error) {
	i, readLen, err := TerminalIDNumFromAZIDBinField(input, typeHint)
	if err == nil {
		*idNum = i
	}
	return readLen, err
}

// Embedded fields
const (
	TerminalIDNumEmbeddedFieldsMask = 0b_00000000_00000000_00000000_00000000_00000000_00000000_00000000_00000000


)

//endregion

//region AttrSet

// TerminalAttributes contains
// attributes for adjunct Terminal.
type TerminalAttributes struct {
	// AZxAdjunctEntityAttrSetBase
	
	secret string
	displayName string
	acceptLanguage string
}

// NewTerminalAttributes returns a new instance
// of TerminalAttributes with the provided attribute values.
func NewTerminalAttributes(
	secret string,
	displayName string,
	acceptLanguage string,
) TerminalAttributes {
	return TerminalAttributes{
		secret: secret,
		displayName: displayName,
		acceptLanguage: acceptLanguage,
	}
}

var _ azcore.ValueObjectAssert[TerminalAttributes] = TerminalAttributes{}

// Clone returns a copy of TerminalAttributes
func (attrs TerminalAttributes) Clone() TerminalAttributes {
	return TerminalAttributes{
		secret: attrs.secret,
		displayName: attrs.displayName,
		acceptLanguage: attrs.acceptLanguage,
	}
}

// Secret returns instance's Secret value.
func (attrs TerminalAttributes) Secret() string {
	return attrs.secret
}

// WithSecret returns a copy
// of TerminalAttributes
// with its secret attribute set to the provided value.
func (attrs TerminalAttributes) WithSecret(
	secret string,
) TerminalAttributes {
	return TerminalAttributes{
		secret: secret,
		displayName: attrs.displayName,
		acceptLanguage: attrs.acceptLanguage,
	}
}

// DisplayName returns instance's DisplayName value.
func (attrs TerminalAttributes) DisplayName() string {
	return attrs.displayName
}

// WithDisplayName returns a copy
// of TerminalAttributes
// with its displayName attribute set to the provided value.
func (attrs TerminalAttributes) WithDisplayName(
	displayName string,
) TerminalAttributes {
	return TerminalAttributes{
		secret: attrs.secret,
		displayName: displayName,
		acceptLanguage: attrs.acceptLanguage,
	}
}

// AcceptLanguage returns instance's AcceptLanguage value.
func (attrs TerminalAttributes) AcceptLanguage() string {
	return attrs.acceptLanguage
}

// WithAcceptLanguage returns a copy
// of TerminalAttributes
// with its acceptLanguage attribute set to the provided value.
func (attrs TerminalAttributes) WithAcceptLanguage(
	acceptLanguage string,
) TerminalAttributes {
	return TerminalAttributes{
		secret: attrs.secret,
		displayName: attrs.displayName,
		acceptLanguage: acceptLanguage,
	}
}

func (attrs TerminalAttributes) Equals(
	other interface{},
) bool {
	if x, ok := other.(TerminalAttributes); ok {
		return attrs.EqualsTerminalAttributesPtr(&x)
	}
	if x, _ := other.(*TerminalAttributes); x != nil {
		return attrs.EqualsTerminalAttributesPtr(x)
	}
	return false
}

func (attrs *TerminalAttributes) EqualsTerminalAttributesPtr(
	other *TerminalAttributes,
) bool {
	return attrs == other || (other != nil &&
		attrs.secret == other.secret &&
		attrs.displayName == other.displayName &&
		attrs.acceptLanguage == other.acceptLanguage)
}

//endregion

//region Instance

// TerminalInstanceService is a service which
// provides methods to manipulate an instance of Terminal.
type TerminalInstanceService interface {
	TerminalInstanceStateService
}

// TerminalInstanceStateService is a service which
// provides access to instances metadata.
type TerminalInstanceStateService interface {
	// GetTerminalInstanceState checks if the provided
    // ref-key is valid and whether the instance is deleted.
	//
	// This method returns nil if the id is not referencing to any valid
	// instance.
	GetTerminalInstanceState(
		ctx context.Context,
		id TerminalID,
	) (*TerminalInstanceState, error)
}

// TerminalInstanceState holds information about
// an instance of Terminal.
type TerminalInstanceState struct {
    RevisionNumber_ int32

    // Deletion_ holds information about the deletion of the instance. If
    // the instance has not been deleted, this field value will be nil.
	Deletion_ *TerminalDeletionState
}

var _ azcore.EntityInstanceInfo[
	int32, TerminalDeletionState,
] = TerminalInstanceState{}
var _ azcore.ValueObjectAssert[
	TerminalDeletionState,
] = TerminalDeletionState{}

// TerminalInstanceStateZero returns an instance of
// TerminalInstanceState with attributes set their respective zero
// value.
func TerminalInstanceStateZero() TerminalInstanceState {
	return TerminalInstanceState{}
}

func (instInfo TerminalInstanceState) Clone() TerminalInstanceState {
	if instInfo.Deletion_ != nil {
		cp := instInfo
		delInfo := cp.Deletion_.Clone()
		cp.Deletion_ = &delInfo
		return cp
	}
	// Already a copy and there's no shared underlying data instance
	return instInfo
}

func (instInfo TerminalInstanceState) RevisionNumber() int32 { return instInfo.RevisionNumber_ }
func (instInfo TerminalInstanceState) Deletion() *TerminalDeletionState {
	return instInfo.Deletion_
}

// IsActive returns true if the instance is considered as active.
func (instInfo TerminalInstanceState) IsActive() bool {
    // Note: we will check other flags in the future, but that's said,
    // deleted instance is considered inactive.
    return !instInfo.IsDeleted()
}

// IsDeleted returns true if the instance was deleted.
func (instInfo TerminalInstanceState) IsDeleted() bool {
	if delInfo := instInfo.Deletion(); delInfo != nil {
		return delInfo.Deleted()
	}
	return false
}

//----

// TerminalDeletionState holds information about
// the deletion of an instance if the instance has been deleted.
type TerminalDeletionState struct {
	Deleted_       bool
}

var _ azcore.EntityDeletionInfo = TerminalDeletionState{}
var _ azcore.ValueObjectAssert[
	TerminalDeletionState,
] = TerminalDeletionState{}

func (instDelInfo TerminalDeletionState) Clone() TerminalDeletionState {
	// Already a copy and there's no shared underlying data instance
	return instDelInfo
}

func (instDelInfo TerminalDeletionState) Deleted() bool {
	return instDelInfo.Deleted_
}

//----

// TerminalInstanceServiceInternal is a service which provides
// methods for manipulating instances of Terminal. Declared for
// internal use within a process, this interface contains methods that
// available to be called from another part of a process.
type TerminalInstanceServiceInternal interface {
}

// TerminalCreationParams contains data to be passed
// as an argument when invoking the method CreateTerminalInternal
// of TerminalInstanceServiceInternal.
type TerminalCreationParams struct {
}

// TerminalDeletionParams contains data to be passed
// as an argument when invoking the method DeleteTerminalInternal
// of TerminalInstanceServiceInternal.
type TerminalDeletionParams struct {
}

//endregion

//region Service

// TerminalService provides a contract
// for methods related to entity Terminal.
type TerminalService interface {
	// AZxEntityService

	TerminalInstanceService
}

// TerminalServiceClient is the interface for
// clients of TerminalService.
type TerminalServiceClient interface {
	TerminalService
}

//endregion
