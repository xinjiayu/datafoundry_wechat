package api

import (
	"fmt"
)

type Error struct {
	code    uint
	message string
}

var (
	Errors            [NumErrors]*Error
	ErrorNone         *Error
	ErrorUnkown       *Error
	ErrorJsonBuilding *Error
)

const (
	ErrorCodeAny = -1 // this is for testing only

	ErrorCodeNone = 0

	ErrorCodeUnkown            = 1300
	ErrorCodeJsonBuilding      = 1301
	ErrorCodeParseJsonFailed   = 1302
	ErrorCodeUrlNotSupported   = 1303
	ErrorCodeDbNotInitlized    = 1304
	ErrorCodeAuthFailed        = 1305
	ErrorCodePermissionDenied  = 1306
	ErrorCodeInvalidParameters = 1307
	ErrorCodeCreateOrder       = 1308
	ErrorCodeQueryOrder        = 1309
	ErrorCodeDeleteCoupon      = 1310
	ErrorCodeGetCoupon         = 1311
	ErrorCodeQueryCoupons      = 1312
	ErrorCodeCallRecharge      = 1313
	ErrorCodeGetCouponById     = 1314
	ErrorCodeCouponHasUsed     = 1315
	ErrorCodeCouponHasExpired  = 1316
	ErrorCodeCouponUnavailable = 1317
	ErrorCodeGetCouponNotExsit = 1318
	ErrorCodeProvideCoupons    = 1319
	ErrorCouponHasProvided     = 1320
	ErrorNoMoreCoupon          = 1321
	ErrorCodeInputParam        = 1322

	NumErrors = 1500 // about 12k memroy wasted
)

func init() {
	initError(ErrorCodeNone, "OK")
	initError(ErrorCodeUnkown, "unknown error")
	initError(ErrorCodeJsonBuilding, "json building error")
	initError(ErrorCodeParseJsonFailed, "parse json failed")

	initError(ErrorCodeUrlNotSupported, "unsupported url")
	initError(ErrorCodeDbNotInitlized, "db is not inited")
	initError(ErrorCodeAuthFailed, "auth failed")
	initError(ErrorCodePermissionDenied, "permission denied")
	initError(ErrorCodeInvalidParameters, "invalid parameters")

	initError(ErrorCodeCreateOrder, "failed to create order")
	initError(ErrorCodeQueryOrder, "failed to query a order")
	initError(ErrorCodeProvideCoupons, "failed to provide coupons")
	initError(ErrorCodeGetCoupon, "failed to retrieve coupon")
	initError(ErrorCodeQueryCoupons, "failed to query coupons")
	initError(ErrorCouponHasProvided, "the user has provided")
	initError(ErrorNoMoreCoupon, "no more coupon can provide")

	initError(ErrorCodeInputParam, "input params is not correct")

	ErrorNone = GetError(ErrorCodeNone)
	ErrorUnkown = GetError(ErrorCodeUnkown)
	ErrorJsonBuilding = GetError(ErrorCodeJsonBuilding)
}

func initError(code uint, message string) {
	if code < NumErrors {
		Errors[code] = newError(code, message)
	}
}

func GetError(code uint) *Error {
	if code > NumErrors {
		return Errors[ErrorCodeUnkown]
	}

	return Errors[code]
}

func GetError2(code uint, message string) *Error {
	e := GetError(code)
	if e == nil {
		return newError(code, message)
	} else {
		return newError(code, fmt.Sprintf("%s (%s)", e.message, message))
	}
}

func newError(code uint, message string) *Error {
	return &Error{code: code, message: message}
}

func newUnknownError(message string) *Error {
	return &Error{
		code:    ErrorCodeUnkown,
		message: message,
	}
}

func newInvalidParameterError(paramName string) *Error {
	return &Error{
		code:    ErrorCodeInvalidParameters,
		message: fmt.Sprintf("%s: %s", GetError(ErrorCodeInvalidParameters).message, paramName),
	}
}
