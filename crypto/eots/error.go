package eots

import (
	ecdsa_schnorr "github.com/decred/dcrd/dcrec/secp256k1/v4/schnorr"
)

// KindOfError identifies a kind of error. It has full support for errors.Is
// and errors.As, so the caller can directly check against an error kind
// when determining the reason for an error.
type KindOfError = ecdsa_schnorr.ErrorKind

// Error identifies an error related to a schnorr signature. It has full
// support for errors.Is and errors.As, so the caller can ascertain the
// specific reason for the error by checking the underlying error.
type Error = ecdsa_schnorr.Error

// signatureError creates an Error given a set of arguments.
func signatureError(kind KindOfError, desc string) Error {
	return Error{Err: kind, Description: desc}
}
