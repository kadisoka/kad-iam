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

// Entity User.

//region ID

// UserID is used to identify
// an instance of entity User.
type UserID UserIDNum

// NewUserID returns a new instance
// of UserID with the provided attribute values.
func NewUserID(
	idNum UserIDNum,
) UserID {
	return UserID(idNum)
}

// To ensure that it conforms the interfaces. If any of these is failing,
// there's a bug in the generator.
var _ azid.ID[UserIDNum] = _UserIDZero
var _ azid.BinUnmarshalable = &_UserIDZeroVar
var _ azid.BinFieldUnmarshalable = &_UserIDZeroVar
var _ azid.TextUnmarshalable = &_UserIDZeroVar
var _ azcore.EntityID[UserIDNum] = _UserIDZero
var _ azcore.ValueObjectAssert[UserID] = _UserIDZero
var _ azcore.UserID[UserIDNum] = _UserIDZero

const _UserIDZero = UserID(UserIDNumZero)

var _UserIDZeroVar = _UserIDZero

// UserIDZero returns
// a zero-valued instance of UserID.
func UserIDZero() UserID {
	return _UserIDZero
}

// Clone returns a copy of self.
func (id UserID) Clone() UserID { return id }

// PrimitiveValue returns the value in its primitive type. Prefer to use
// this method instead of casting directly.
func (id UserID) PrimitiveValue() int64 {
	return int64(id)
}

// IDNum returns the scoped identifier of the entity.
func (id UserID) IDNum() UserIDNum {
	return UserIDNum(id)
}

// IDNumPtr returns a pointer to a copy of the id-num if it's considered valid
// otherwise it returns nil.
func (id UserID) IDNumPtr() *UserIDNum {
	if id.IsNotStaticallyValid() {
		return nil
	}
	i := id.IDNum()
	return &i
}

// AZIDNum is required for conformance with azid.ID.
func (id UserID) AZIDNum() UserIDNum {
	return UserIDNum(id)
}

// UserIDNum is required for conformance
// with azcore.UserID.
func (id UserID) UserIDNum() UserIDNum {
	return UserIDNum(id)
}

// IsZero is required as UserID is a value-object.
func (id UserID) IsZero() bool {
	return UserIDNum(id) == UserIDNumZero
}

// IsStaticallyValid returns true if this instance is valid as an isolated value
// of UserID.
// It doesn't tell whether it refers to a valid instance of User.
func (id UserID) IsStaticallyValid() bool {
	return UserIDNum(id).IsStaticallyValid()
}

// IsNotStaticallyValid returns the negation of value returned by IsStaticallyValid.
func (id UserID) IsNotStaticallyValid() bool {
	return !id.IsStaticallyValid()
}

// Equals is required for conformance with azcore.EntityID.
func (id UserID) Equals(other interface{}) bool {
	if x, ok := other.(UserID); ok {
		return x == id
	}
	if x, _ := other.(*UserID); x != nil {
		return *x == id
	}
	return false
}

// Equal is required for conformance with azcore.EntityID.
func (id UserID) Equal(other interface{}) bool {
	return id.Equals(other)
}

// EqualsUserID returns true
// if the other value has the same attributes as id.
func (id UserID) EqualsUserID(
	other UserID,
) bool {
	return other == id
}

func (id UserID) AZIDBin() []byte {
	b := make([]byte, 8+1)
	b[0] = UserIDNumBinDataType.Byte()
	binary.BigEndian.PutUint64(b[1:], uint64(id))
	return b
}

// UserIDFromAZIDBin creates a new instance of
// UserID from its azid-bin-encoded form.
func UserIDFromAZIDBin(
	input []byte,
) (id UserID, readLen int, err error) {
	typ, err := azid.BinDataTypeFromByte(input[0])
	if err != nil {
		return _UserIDZero, 0,
			errors.ArgMW("input", "type parsing", err)
	}
	if typ != UserIDNumBinDataType {
		return _UserIDZero, 0,
			errors.ArgMsg("input", "type unsupported")
	}

	i, readLen, err := UserIDFromAZIDBinField(input[1:], typ)
	if err != nil {
		return _UserIDZero, 0,
			errors.ArgMW("input", "id-num parsing", err)
	}

	return UserID(i), 1 + readLen, nil
}

// UnmarshalAZIDBin is required for conformance
// with azcore.BinFieldUnmarshalable.
func (id *UserID) UnmarshalAZIDBin(input []byte) (readLen int, err error) {
	i, readLen, err := UserIDFromAZIDBin(input)
	if err == nil {
		*id = i
	}
	return readLen, err
}

func (id UserID) AZIDBinField() ([]byte, azid.BinDataType) {
	return UserIDNum(id).AZIDBinField()
}

func UserIDFromAZIDBinField(
	input []byte, typeHint azid.BinDataType,
) (id UserID, readLen int, err error) {
	idNum, n, err := UserIDNumFromAZIDBinField(input, typeHint)
	if err != nil {
		return _UserIDZero, n, err
	}
	return UserID(idNum), n, nil
}

// UnmarshalAZIDBinField is required for conformance
// with azcore.BinFieldUnmarshalable.
func (id *UserID) UnmarshalAZIDBinField(
	input []byte, typeHint azid.BinDataType,
) (readLen int, err error) {
	i, readLen, err := UserIDFromAZIDBinField(input, typeHint)
	if err == nil {
		*id = i
	}
	return readLen, err
}

const _UserIDAZIDTextPrefix = "KUs0"

// AZIDText is required for conformance
// with azid.ID.
func (id UserID) AZIDText() string {
	if !id.IsStaticallyValid() {
		return ""
	}

	return _UserIDAZIDTextPrefix +
		azid.TextEncode(id.AZIDBin())
}

// UserIDFromAZIDText creates a new instance of
// UserID from its azid-text form.
func UserIDFromAZIDText(input string) (UserID, error) {
	if input == "" {
		return UserIDZero(), nil
	}
	if !strings.HasPrefix(input, _UserIDAZIDTextPrefix) {
		return UserIDZero(),
			errors.ArgMsg("input", "prefix mismatch")
	}
	input = strings.TrimPrefix(input, _UserIDAZIDTextPrefix)
	b, err := azid.TextDecode(input)
	if err != nil {
		return UserIDZero(),
			errors.ArgMW("input", "decoding", err)
	}
	id, _, err := UserIDFromAZIDBin(b)
	if err != nil {
		return UserIDZero(),
			errors.ArgMW("input", "parsing", err)
	}
	return id, nil
}

// UnmarshalAZIDText is required for conformance
// with azid.TextUnmarshalable.
func (id *UserID) UnmarshalAZIDText(input string) error {
	r, err := UserIDFromAZIDText(input)
	if err == nil {
		*id = r
	}
	return err
}

// MarshalText is for compatibility with Go's encoding.TextMarshaler
func (id UserID) MarshalText() ([]byte, error) {
	return []byte(id.AZIDText()), nil
}

// UnmarshalText is for conformance with Go's encoding.TextUnmarshaler
func (id *UserID) UnmarshalText(b []byte) error {
	r, err := UserIDFromAZIDText(string(b))
	if err == nil {
		*id = r
	}
	return err
}

// MarshalJSON makes this type JSON-marshalable.
func (id UserID) MarshalJSON() ([]byte, error) {
	// We assume that there are no symbols in azid-text
	return []byte("\"" + id.AZIDText() + "\""), nil
}

// UnmarshalJSON parses a JSON value.
func (id *UserID) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "" {
		*id = UserIDZero()
		return nil
	}
	i, err := UserIDFromAZIDText(s)
	if err == nil {
		*id = i
	}
	return err
}

// UserIDService abstracts
// UserID-related services.
type UserIDService interface {
	// IsUserID is to check if the ref-key is
	// trully registered to system. It does not check whether the instance
	// is active or not.
	IsUserIDRegistered(id UserID) bool
}

//endregion

//region IDNum

// UserIDNum is a scoped identifier
// used to identify an instance of entity User.
type UserIDNum int64

// UserIDNumBinDataType is the type when it is in its
// az-bin-encoded form.
const UserIDNumBinDataType = azid.BinDataTypeInt64

// To ensure that it conforms the interfaces. If any of these is failing,
// there's a bug in the generator.
var _ azid.IDNumMethods = UserIDNumZero
var _ azid.BinFieldUnmarshalable = &_UserIDNumZeroVar
var _ azcore.EntityIDNumMethods = UserIDNumZero
var _ azcore.ValueObjectAssert[UserIDNum] = UserIDNumZero
var _ azcore.UserIDNumMethods = UserIDNumZero

// UserIDNumIdentifierBitsMask is used to
// extract identifier bits from an instance of UserIDNum.
const UserIDNumIdentifierBitsMask uint64 = 0b_00000000_00000000_11111111_11111111_11111111_11111111_11111111_11111111

// UserIDNumZero is the zero value
// for UserIDNum.
const UserIDNumZero = UserIDNum(0)

// _UserIDNumZeroVar is used for testing
// pointer-based interfaces conformance.
var _UserIDNumZeroVar = UserIDNumZero

// UserIDNumFromPrimitiveValue creates an instance
// of UserIDNum from its primitive value.
func UserIDNumFromPrimitiveValue(v int64) UserIDNum {
	return UserIDNum(v)
}

// UserIDNumFromAZIDBinField creates UserIDNum from
// its azid-bin-field form.
func UserIDNumFromAZIDBinField(
	input []byte, typeHint azid.BinDataType,
) (idNum UserIDNum, readLen int, err error) {
	if typeHint != azid.BinDataTypeUnspecified && typeHint != UserIDNumBinDataType {
		return UserIDNum(0), 0,
			errors.ArgMsg("typeHint", "value unsupported")
	}
	i := binary.BigEndian.Uint64(input)
	return UserIDNum(i), 8, nil
}

// PrimitiveValue returns the value in its primitive type. Prefer to use
// this method instead of casting directly.
func (idNum UserIDNum) PrimitiveValue() int64 {
	return int64(idNum)
}

// Clone returns a copy of self.
func (idNum UserIDNum) Clone() UserIDNum { return idNum }

// IsZero is required as UserIDNum is a value-object.
func (idNum UserIDNum) IsZero() bool {
	return idNum == UserIDNumZero
}

// IsStaticallyValid returns true if this instance is valid as an isolated value
// of UserIDNum. It doesn't tell whether it refers to
// a valid instance of User.
//
// What is considered valid in this context here is that the data
// contained in this instance doesn't break any rule for an instance of
// UserIDNum. Whether the instance is valid in a certain context,
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
func (idNum UserIDNum) IsStaticallyValid() bool {
	return int64(idNum) > 0 &&
		(uint64(idNum)&UserIDNumIdentifierBitsMask) != 0
}

// IsNotStaticallyValid returns the negation of value returned by IsStaticallyValid.
func (idNum UserIDNum) IsNotStaticallyValid() bool {
	return !idNum.IsStaticallyValid()
}

// Equals is required as UserIDNum is a value-object.
//
// Use EqualsUserIDNum method if the other value
// has the same type.
func (idNum UserIDNum) Equals(other interface{}) bool {
	if x, ok := other.(UserIDNum); ok {
		return x == idNum
	}
	if x, _ := other.(*UserIDNum); x != nil {
		return *x == idNum
	}
	return false
}

// Equal is a wrapper for Equals method. It is required for
// compatibility with github.com/google/go-cmp
func (idNum UserIDNum) Equal(other interface{}) bool {
	return idNum.Equals(other)
}

// EqualsUserIDNum determines if the other instance is equal
// to this instance.
func (idNum UserIDNum) EqualsUserIDNum(
	other UserIDNum,
) bool {
	return idNum == other
}

// AZIDBinField is required for conformance
// with azid.IDNum.
func (idNum UserIDNum) AZIDBinField() ([]byte, azid.BinDataType) {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(idNum))
	return b, UserIDNumBinDataType
}

// UnmarshalAZIDBinField is required for conformance
// with azid.BinFieldUnmarshalable.
func (idNum *UserIDNum) UnmarshalAZIDBinField(
	input []byte, typeHint azid.BinDataType,
) (readLen int, err error) {
	i, readLen, err := UserIDNumFromAZIDBinField(input, typeHint)
	if err == nil {
		*idNum = i
	}
	return readLen, err
}

// Embedded fields
const (
	UserIDNumEmbeddedFieldsMask = 0b_01000000_00000000_00000000_00000000_00000000_00000000_00000000_00000000

	UserIDNumBotMask = 0b_01000000_00000000_00000000_00000000_00000000_00000000_00000000_00000000
	UserIDNumBotBits = 0b_01000000_00000000_00000000_00000000_00000000_00000000_00000000_00000000
)

// IsBot returns true if
// the User instance this UserIDNum is for
// is a Bot User.
//
// Bot account is ....
func (idNum UserIDNum) IsBot() bool {
	return idNum.IsStaticallyValid() &&
		idNum.HasBotBits()
}

// HasBotBits is only checking the bits
// without validating other information.
func (idNum UserIDNum) HasBotBits() bool {
	return (uint64(idNum) &
		UserIDNumBotMask) ==
		UserIDNumBotBits
}

//endregion

//region AttrSet

// UserAttrSet is a value-object which contains
// entity's built-in attributes.
type UserAttrSet struct {
	// AZxEntityAttrSetBase

}

// NewUserAttrSet returns a new instance
// of UserAttrSet with the provided attribute values.
func NewUserAttrSet() UserAttrSet {
	return UserAttrSet{}
}

var _ azcore.ValueObjectAssert[UserAttrSet] = UserAttrSet{}

// Clone returns a copy of UserAttrSet
func (attrs UserAttrSet) Clone() UserAttrSet {
	return UserAttrSet{}
}

func (attrs UserAttrSet) Equals(
	other interface{},
) bool {
	if x, ok := other.(UserAttrSet); ok {
		return attrs.EqualsUserAttrSetPtr(&x)
	}
	if x, _ := other.(*UserAttrSet); x != nil {
		return attrs.EqualsUserAttrSetPtr(x)
	}
	return false
}

func (attrs *UserAttrSet) EqualsUserAttrSetPtr(
	other *UserAttrSet,
) bool {
	return attrs == other || (other != nil)
}

//endregion

//region Instance

// UserInstanceService is a service which
// provides methods to manipulate an instance of User.
type UserInstanceService interface {
	UserInstanceStateService
}

// UserInstanceStateService is a service which
// provides access to instances metadata.
type UserInstanceStateService interface {
	// GetUserInstanceState checks if the provided
	// ref-key is valid and whether the instance is deleted.
	//
	// This method returns nil if the id is not referencing to any valid
	// instance.
	GetUserInstanceState(
		ctx context.Context,
		id UserID,
	) (*UserInstanceState, error)
}

// UserInstanceState holds information about
// an instance of User.
type UserInstanceState struct {
	RevisionNumber_ int32

	// Deletion_ holds information about the deletion of the instance. If
	// the instance has not been deleted, this field value will be nil.
	Deletion_ *UserDeletionState
}

var _ azcore.EntityInstanceInfo[
	int32, UserDeletionState,
] = UserInstanceState{}
var _ azcore.ValueObjectAssert[UserDeletionState] = UserDeletionState{}

// UserInstanceStateZero returns an instance of
// UserInstanceState with attributes set their respective zero
// value.
func UserInstanceStateZero() UserInstanceState {
	return UserInstanceState{}
}

func (instInfo UserInstanceState) Clone() UserInstanceState {
	if instInfo.Deletion_ != nil {
		cp := instInfo
		delInfo := cp.Deletion_.Clone()
		cp.Deletion_ = &delInfo
		return cp
	}
	// Already a copy and there's no shared underlying data instance
	return instInfo
}

func (instInfo UserInstanceState) RevisionNumber() int32 { return instInfo.RevisionNumber_ }
func (instInfo UserInstanceState) Deletion() *UserDeletionState {
	return instInfo.Deletion_
}

// IsActive returns true if the instance is considered as active.
func (instInfo UserInstanceState) IsActive() bool {
	// Note: we will check other flags in the future, but that's said,
	// deleted instance is considered inactive.
	return !instInfo.IsDeleted()
}

// IsDeleted returns true if the instance was deleted.
func (instInfo UserInstanceState) IsDeleted() bool {
	if delInfo := instInfo.Deletion(); delInfo != nil {
		return delInfo.Deleted()
	}
	return false
}

//----

// UserDeletionState holds information about
// the deletion of an instance if the instance has been deleted.
type UserDeletionState struct {
	Deleted_       bool
	DeletionNotes_ string
}

var _ azcore.EntityDeletionInfo = UserDeletionState{}
var _ azcore.ValueObjectAssert[UserDeletionState] = UserDeletionState{}

func (instDelInfo UserDeletionState) Clone() UserDeletionState {
	// Already a copy and there's no shared underlying data instance
	return instDelInfo
}

func (instDelInfo UserDeletionState) Deleted() bool {
	return instDelInfo.Deleted_
}
func (instDelInfo UserDeletionState) DeletionNotes() string {
	return instDelInfo.DeletionNotes_
}

//----

// UserInstanceServiceInternal is a service which provides
// methods for manipulating instances of User. Declared for
// internal use within a process, this interface contains methods that
// available to be called from another part of a process.
type UserInstanceServiceInternal interface {
	CreateUserInternal(
		ctx context.Context,
		input UserCreationParams,
	) (id UserID, initialState UserInstanceState, err error)

	// DeleteUserInternal deletes an instance of
	// User entity based identfied by idOfInstToDel.
	// The returned justDeleted will have the value of
	// true if this particular call resulted the deletion of the instance and
	// it will have the value of false of subsequent calls to this method.
	DeleteUserInternal(
		ctx context.Context,
		idOfInstToDel UserID,
		input UserDeletionParams,
	) (justDeleted bool, currentState UserInstanceState, err error)
}

// UserCreationParams contains data to be passed
// as an argument when invoking the method CreateUserInternal
// of UserInstanceServiceInternal.
type UserCreationParams struct {
}

// UserDeletionParams contains data to be passed
// as an argument when invoking the method DeleteUserInternal
// of UserInstanceServiceInternal.
type UserDeletionParams struct {
	DeletionNotes string
}

//endregion

//region Service

// UserService provides a contract
// for methods related to entity User.
type UserService interface {
	// AZxEntityService

	UserIDService
	UserInstanceService
	UserInstanceStateService
}

// UserServiceClient is the interface for
// clients of UserService.
type UserServiceClient interface {
	UserService
}

//endregion
