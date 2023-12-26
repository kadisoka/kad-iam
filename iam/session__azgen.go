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

// Adjunct-entity Session of Terminal.
//
// A Session represents authorization for a time span. While Terminal
// usually provide longer authorization period, a Session is used to
// break down that authorization into smaller time spans.
//
// If a Session is expired or revoked, the previously authorized
// Application instance (Terminal) could request another Session as long the
// Application's authorization is still valid. There's only one instance
// of Session active at a time for a Terminal.
//
// A Session is represented by an access token.

//region ID

// SessionID is used to identify
// an instance of adjunct entity Session system-wide.
type SessionID struct {
	terminal TerminalID
	idNum    SessionIDNum
}

// The total number of fields in the struct.
const _SessionIDFieldCount = 1 + 1

// NewSessionID returns a new instance
// of SessionID with the provided attribute values.
func NewSessionID(
	terminal TerminalID,
	idNum SessionIDNum,
) SessionID {
	return SessionID{
		terminal: terminal,
		idNum:    idNum,
	}
}

// To ensure that it conforms the interfaces. If any of these is failing,
// there's a bug in the generator.
var _ azid.ID[SessionIDNum] = _SessionIDZero
var _ azid.BinFieldUnmarshalable = &_SessionIDZero
var _ azid.TextUnmarshalable = &_SessionIDZero
var _ azcore.AdjunctEntityID[SessionIDNum] = _SessionIDZero
var _ azcore.ValueObjectAssert[SessionID] = _SessionIDZero
var _ azcore.SessionID[SessionIDNum] = _SessionIDZero

var _SessionIDZero = SessionID{
	terminal: TerminalIDZero(),
	idNum:    SessionIDNumZero,
}

// SessionIDZero returns
// a zero-valued instance of SessionID.
func SessionIDZero() SessionID {
	return _SessionIDZero
}

// Clone returns a copy of self.
func (id SessionID) Clone() SessionID { return id }

// IDNum returns the scoped identifier of the entity.
func (id SessionID) IDNum() SessionIDNum {
	return id.idNum
}

// IDNumPrimitiveValue returns the value in its primitive type. Prefer to use
// this method instead of casting directly.
func (id SessionID) IDNumPrimitiveValue() int32 {
	return int32(id.idNum)
}

// IDNumPtr returns a pointer to a copy of the id-num if it's considered valid
// otherwise it returns nil.
func (id SessionID) IDNumPtr() *SessionIDNum {
	if id.IsNotStaticallyValid() {
		return nil
	}
	i := id.IDNum()
	return &i
}

// AZIDNum is required for conformance with azid.ID.
func (id SessionID) AZIDNum() SessionIDNum {
	return id.idNum
}

// SessionIDNum is required for conformance
// with azcore.SessionID.
func (id SessionID) SessionIDNum() SessionIDNum {
	return id.idNum
}

// IsZero is required as SessionID is a value-object.
func (id SessionID) IsZero() bool {
	return id.terminal.IsZero() &&
		id.idNum == SessionIDNumZero
}

// IsStaticallyValid returns true if this instance is valid as an isolated value
// of SessionID.
// It doesn't tell whether it refers to a valid instance of Session.
func (id SessionID) IsStaticallyValid() bool {
	return id.terminal.IsStaticallyValid() &&
		id.idNum.IsStaticallyValid()
}

// IsNotStaticallyValid returns the negation of value returned by IsStaticallyValid.
func (id SessionID) IsNotStaticallyValid() bool {
	return !id.IsStaticallyValid()
}

// Equals is required for conformance with azcore.AdjunctEntityID.
func (id SessionID) Equals(other interface{}) bool {
	if x, ok := other.(SessionID); ok {
		return id.terminal.EqualsTerminalID(x.terminal) &&
			id.idNum == x.idNum
	}
	if x, _ := other.(*SessionID); x != nil {
		return id.terminal.EqualsTerminalID(x.terminal) &&
			id.idNum == x.idNum
	}
	return false
}

// Equal is required for conformance with azcore.AdjunctEntityID.
func (id SessionID) Equal(other interface{}) bool {
	return id.Equals(other)
}

// EqualsSessionID returns true
// if the other value has the same attributes as id.
func (id SessionID) EqualsSessionID(
	other SessionID,
) bool {
	return id.terminal.EqualsTerminalID(other.terminal) &&
		id.idNum == other.idNum
}

// AZIDBin is required for conformance
// with azid.ID.
func (id SessionID) AZIDBin() []byte {
	data, typ := id.AZIDBinField()
	out := []byte{typ.Byte()}
	return append(out, data...)
}

// SessionIDFromAZIDBin creates a new instance of
// SessionID from its azid-bin-encoded form.
func SessionIDFromAZIDBin(
	input []byte,
) (id SessionID, readLen int, err error) {
	typ, err := azid.BinDataTypeFromByte(input[0])
	if err != nil {
		return _SessionIDZero, 0,
			errors.ArgMW("input", "type parsing", err)
	}
	if typ != azid.BinDataTypeArray {
		return _SessionIDZero, 0,
			errors.ArgMsg("input", "type unsupported")
	}

	id, readLen, err = SessionIDFromAZIDBinField(input[1:], typ)
	return id, readLen + 1, err
}

// AZIDBinField is required for conformance
// with azid.ID.
func (id SessionID) AZIDBinField() ([]byte, azid.BinDataType) {
	var typesBytes []byte
	var dataBytes []byte
	var fieldBytes []byte
	var fieldType azid.BinDataType

	fieldBytes, fieldType = id.terminal.AZIDBinField()
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

// SessionIDFromAZIDBinField creates SessionID from
// its azid-bin-encoded field form.
func SessionIDFromAZIDBinField(
	input []byte, typeHint azid.BinDataType,
) (id SessionID, readLen int, err error) {
	if typeHint != azid.BinDataTypeArray {
		return SessionIDZero(), 0,
			errors.ArgMsg("typeHint", "value unsupported")
	}

	arrayLen := int(input[0])
	if arrayLen != _SessionIDFieldCount {
		return SessionIDZero(), 0,
			errors.ArgMsg("input", "field count mismatch")
	}

	typeCursor := 1
	dataCursor := typeCursor + arrayLen

	var fieldType azid.BinDataType

	fieldType, err = azid.BinDataTypeFromByte(input[typeCursor])
	if err != nil {
		return SessionIDZero(), 0,
			errors.ArgMW("input", "host terminal ref-key type parsing", err)
	}
	typeCursor++
	terminalID, readLen, err := TerminalIDFromAZIDBinField(
		input[dataCursor:], fieldType)
	if err != nil {
		return SessionIDZero(), 0,
			errors.ArgMW("input", "host terminal ref-key data parsing", err)
	}
	dataCursor += readLen

	fieldType, err = azid.BinDataTypeFromByte(input[typeCursor])
	if err != nil {
		return SessionIDZero(), 0,
			errors.ArgMW("input", "id-num type parsing", err)
	}
	typeCursor++
	idNum, readLen, err := SessionIDNumFromAZIDBinField(
		input[dataCursor:], fieldType)
	if err != nil {
		return SessionIDZero(), 0,
			errors.ArgMW("input", "id-num data parsing", err)
	}
	dataCursor += readLen

	return SessionID{
		terminal: terminalID,
		idNum:    idNum,
	}, dataCursor, nil
}

// UnmarshalAZIDBinField is required for conformance
// with azcore.BinFieldUnmarshalable.
func (id *SessionID) UnmarshalAZIDBinField(
	input []byte, typeHint azid.BinDataType,
) (readLen int, err error) {
	i, readLen, err := SessionIDFromAZIDBinField(input, typeHint)
	if err == nil {
		*id = i
	}
	return readLen, err
}

const _SessionIDAZIDTextPrefix = "KSe0"

// AZIDText is required for conformance
// with azid.ID.
func (id SessionID) AZIDText() string {
	if !id.IsStaticallyValid() {
		return ""
	}

	return _SessionIDAZIDTextPrefix +
		azid.TextEncode(id.AZIDBin())
}

// SessionIDFromAZIDText creates a new instance of
// SessionID from its azid-text form.
func SessionIDFromAZIDText(input string) (SessionID, error) {
	if input == "" {
		return SessionIDZero(), nil
	}
	if !strings.HasPrefix(input, _SessionIDAZIDTextPrefix) {
		return SessionIDZero(),
			errors.ArgMsg("input", "prefix mismatch")
	}
	input = strings.TrimPrefix(input, _SessionIDAZIDTextPrefix)
	b, err := azid.TextDecode(input)
	if err != nil {
		return SessionIDZero(),
			errors.ArgMW("input", "decoding", err)
	}
	id, _, err := SessionIDFromAZIDBin(b)
	if err != nil {
		return SessionIDZero(),
			errors.ArgMW("input", "parsing", err)
	}
	return id, nil
}

// UnmarshalAZIDText is required for conformance
// with azid.TextUnmarshalable.
func (id *SessionID) UnmarshalAZIDText(input string) error {
	r, err := SessionIDFromAZIDText(input)
	if err == nil {
		*id = r
	}
	return err
}

// MarshalText is for compatibility with Go's encoding.TextMarshaler
func (id SessionID) MarshalText() ([]byte, error) {
	return []byte(id.AZIDText()), nil
}

// UnmarshalText is for conformance with Go's encoding.TextUnmarshaler
func (id *SessionID) UnmarshalText(b []byte) error {
	r, err := SessionIDFromAZIDText(string(b))
	if err == nil {
		*id = r
	}
	return err
}

// MarshalJSON makes this type JSON-marshalable.
func (id SessionID) MarshalJSON() ([]byte, error) {
	// We assume that there are no symbols in azid-text
	return []byte("\"" + id.AZIDText() + "\""), nil
}

// UnmarshalJSON parses a JSON value.
func (id *SessionID) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "" {
		*id = SessionIDZero()
		return nil
	}
	i, err := SessionIDFromAZIDText(s)
	if err == nil {
		*id = i
	}
	return err
}

// Terminal returns instance's Terminal value.
func (id SessionID) Terminal() TerminalID {
	return id.terminal
}

// TerminalPtr returns a pointer to a copy of
// TerminalID if it's considered valid.
func (id SessionID) TerminalPtr() *TerminalID {
	if id.terminal.IsStaticallyValid() {
		rk := id.terminal
		return &rk
	}
	return nil
}

// WithTerminal returns a copy
// of SessionID
// with its terminal attribute set to the provided value.
func (id SessionID) WithTerminal(
	terminal TerminalID,
) SessionID {
	return SessionID{
		terminal: terminal,
		idNum:    id.idNum,
	}
}

//endregion

//region IDNum

// SessionIDNum is a scoped identifier
// used to identify an instance of adjunct entity Session
// scoped within its host entity(s).
type SessionIDNum int32

// SessionIDNumBinDataType is the type when it is in its
// az-bin-encoded form.
const SessionIDNumBinDataType = azid.BinDataTypeInt32

// To ensure that it conforms the interfaces. If any of these is failing,
// there's a bug in the generator.
var _ azid.IDNumMethods = SessionIDNumZero
var _ azid.BinFieldUnmarshalable = &_SessionIDNumZeroVar
var _ azcore.AdjunctEntityIDNumMethods = SessionIDNumZero
var _ azcore.ValueObjectAssert[SessionIDNum] = SessionIDNumZero
var _ azcore.SessionIDNumMethods = SessionIDNumZero

// SessionIDNumIdentifierBitsMask is used to
// extract identifier bits from an instance of SessionIDNum.
const SessionIDNumIdentifierBitsMask uint32 = 0b_00000000_11111111_11111111_11111111

// SessionIDNumZero is the zero value for SessionIDNum.
const SessionIDNumZero = SessionIDNum(0)

// _SessionIDNumZeroVar is used for testing
// pointer-based interfaces conformance.
var _SessionIDNumZeroVar = SessionIDNumZero

// SessionIDNumFromPrimitiveValue creates an instance
// of SessionIDNum from its primitive value.
func SessionIDNumFromPrimitiveValue(v int32) SessionIDNum {
	return SessionIDNum(v)
}

// SessionIDNumFromAZIDBinField creates SessionIDNum from
// its azid-bin-encoded form.
func SessionIDNumFromAZIDBinField(
	input []byte, typeHint azid.BinDataType,
) (idNum SessionIDNum, readLen int, err error) {
	if typeHint != azid.BinDataTypeUnspecified && typeHint != SessionIDNumBinDataType {
		return SessionIDNum(0), 0,
			errors.ArgMsg("typeHint", "value unsupported")
	}
	i := binary.BigEndian.Uint32(input)
	return SessionIDNum(i), 4, nil
}

// PrimitiveValue returns the value in its primitive type. Prefer to use
// this method instead of casting directly.
func (idNum SessionIDNum) PrimitiveValue() int32 {
	return int32(idNum)
}

// Clone returns a copy of self.
func (idNum SessionIDNum) Clone() SessionIDNum { return idNum }

// IsZero is required as SessionIDNum is a value-object.
func (idNum SessionIDNum) IsZero() bool {
	return idNum == SessionIDNumZero
}

// IsStaticallyValid returns true if this instance is valid as an isolated value
// of SessionIDNum. It doesn't tell whether it refers to
// a valid instance of Session.
func (idNum SessionIDNum) IsStaticallyValid() bool {
	return int32(idNum) > 0 &&
		(uint32(idNum)&SessionIDNumIdentifierBitsMask) != 0
}

// IsNotStaticallyValid returns the negation of value returned by IsStaticallyValid.
func (idNum SessionIDNum) IsNotStaticallyValid() bool {
	return !idNum.IsStaticallyValid()
}

// Equals is required as SessionIDNum is a value-object.
//
// Use EqualsSessionIDNum method if the other value
// has the same type.
func (idNum SessionIDNum) Equals(other interface{}) bool {
	if x, ok := other.(SessionIDNum); ok {
		return x == idNum
	}
	if x, _ := other.(*SessionIDNum); x != nil {
		return *x == idNum
	}
	return false
}

// Equal is a wrapper for Equals method. It is required for
// compatibility with github.com/google/go-cmp
func (idNum SessionIDNum) Equal(other interface{}) bool {
	return idNum.Equals(other)
}

// EqualsSessionIDNum determines if the other instance
// is equal to this instance.
func (idNum SessionIDNum) EqualsSessionIDNum(
	other SessionIDNum,
) bool {
	return idNum == other
}

// AZIDBinField is required for conformance
// with azid.IDNum.
func (idNum SessionIDNum) AZIDBinField() ([]byte, azid.BinDataType) {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(idNum))
	return b, SessionIDNumBinDataType
}

// UnmarshalAZIDBinField is required for conformance
// with azid.BinFieldUnmarshalable.
func (idNum *SessionIDNum) UnmarshalAZIDBinField(
	input []byte, typeHint azid.BinDataType,
) (readLen int, err error) {
	i, readLen, err := SessionIDNumFromAZIDBinField(input, typeHint)
	if err == nil {
		*idNum = i
	}
	return readLen, err
}

// Embedded fields
const (
	SessionIDNumEmbeddedFieldsMask = 0b_00000000_00000000_00000000_00000000
)

//endregion

//region AttrSet

// SessionAttributes contains
// attributes for adjunct Session.
type SessionAttributes struct {
	// AZxAdjunctEntityAttrSetBase

}

// NewSessionAttributes returns a new instance
// of SessionAttributes with the provided attribute values.
func NewSessionAttributes() SessionAttributes {
	return SessionAttributes{}
}

var _ azcore.ValueObjectAssert[SessionAttributes] = SessionAttributes{}

// Clone returns a copy of SessionAttributes
func (attrs SessionAttributes) Clone() SessionAttributes {
	return SessionAttributes{}
}

func (attrs SessionAttributes) Equals(
	other interface{},
) bool {
	if x, ok := other.(SessionAttributes); ok {
		return attrs.EqualsSessionAttributesPtr(&x)
	}
	if x, _ := other.(*SessionAttributes); x != nil {
		return attrs.EqualsSessionAttributesPtr(x)
	}
	return false
}

func (attrs *SessionAttributes) EqualsSessionAttributesPtr(
	other *SessionAttributes,
) bool {
	return attrs == other || (other != nil)
}

//endregion

//region Instance

// SessionInstanceService is a service which
// provides methods to manipulate an instance of Session.
type SessionInstanceService interface {
	SessionInstanceStateService
}

// SessionInstanceStateService is a service which
// provides access to instances metadata.
type SessionInstanceStateService interface {
	// GetSessionInstanceState checks if the provided
	// ref-key is valid and whether the instance is deleted.
	//
	// This method returns nil if the id is not referencing to any valid
	// instance.
	GetSessionInstanceState(
		ctx context.Context,
		id SessionID,
	) (*SessionInstanceState, error)
}

// SessionInstanceState holds information about
// an instance of Session.
type SessionInstanceState struct {
	RevisionNumber_ int32

	// Deletion_ holds information about the deletion of the instance. If
	// the instance has not been deleted, this field value will be nil.
	Deletion_ *SessionDeletionState
}

var _ azcore.EntityInstanceInfo[
	int32, SessionDeletionState,
] = SessionInstanceState{}
var _ azcore.ValueObjectAssert[SessionDeletionState] = SessionDeletionState{}

// SessionInstanceStateZero returns an instance of
// SessionInstanceState with attributes set their respective zero
// value.
func SessionInstanceStateZero() SessionInstanceState {
	return SessionInstanceState{}
}

func (instInfo SessionInstanceState) Clone() SessionInstanceState {
	if instInfo.Deletion_ != nil {
		cp := instInfo
		delInfo := cp.Deletion_.Clone()
		cp.Deletion_ = &delInfo
		return cp
	}
	// Already a copy and there's no shared underlying data instance
	return instInfo
}

func (instInfo SessionInstanceState) RevisionNumber() int32 { return instInfo.RevisionNumber_ }
func (instInfo SessionInstanceState) Deletion() *SessionDeletionState {
	return instInfo.Deletion_
}

// IsActive returns true if the instance is considered as active.
func (instInfo SessionInstanceState) IsActive() bool {
	// Note: we will check other flags in the future, but that's said,
	// deleted instance is considered inactive.
	return !instInfo.IsDeleted()
}

// IsDeleted returns true if the instance was deleted.
func (instInfo SessionInstanceState) IsDeleted() bool {
	if delInfo := instInfo.Deletion(); delInfo != nil {
		return delInfo.Deleted()
	}
	return false
}

//----

// SessionDeletionState holds information about
// the deletion of an instance if the instance has been deleted.
type SessionDeletionState struct {
	Deleted_ bool
}

var _ azcore.EntityDeletionInfo = SessionDeletionState{}
var _ azcore.ValueObjectAssert[SessionDeletionState] = SessionDeletionState{}

func (instDelInfo SessionDeletionState) Clone() SessionDeletionState {
	// Already a copy and there's no shared underlying data instance
	return instDelInfo
}

func (instDelInfo SessionDeletionState) Deleted() bool {
	return instDelInfo.Deleted_
}

//----

// SessionInstanceServiceInternal is a service which provides
// methods for manipulating instances of Session. Declared for
// internal use within a process, this interface contains methods that
// available to be called from another part of a process.
type SessionInstanceServiceInternal interface {
}

// SessionCreationParams contains data to be passed
// as an argument when invoking the method CreateSessionInternal
// of SessionInstanceServiceInternal.
type SessionCreationParams struct {
}

// SessionDeletionParams contains data to be passed
// as an argument when invoking the method DeleteSessionInternal
// of SessionInstanceServiceInternal.
type SessionDeletionParams struct {
}

//endregion

//region Service

// SessionService provides a contract
// for methods related to entity Session.
type SessionService interface {
	// AZxEntityService

	SessionInstanceService
}

// SessionServiceClient is the interface for
// clients of SessionService.
type SessionServiceClient interface {
	SessionService
}

//endregion
