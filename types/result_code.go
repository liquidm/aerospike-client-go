// Copyright 2013-2014 Aerospike, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import "errors"

// ResultCode signifies the database operation error codes.
// The positive numbers align with the server side file proto.h.
type ResultCode int

const (
	// Asynchronous max concurrent database commands have been exceeded and therefore rejected.
	TYPE_NOT_SUPPORTED ResultCode = -7

	// Asynchronous max concurrent database commands have been exceeded and therefore rejected.
	COMMAND_REJECTED ResultCode = -6

	// Query was terminated by user.
	QUERY_TERMINATED ResultCode = -5

	// Scan was terminated by user.
	SCAN_TERMINATED ResultCode = -4

	// Chosen node is not currently active.
	INVALID_NODE_ERROR ResultCode = -3

	// Client parse error.
	PARSE_ERROR ResultCode = -2

	// Client serialization error.
	SERIALIZE_ERROR ResultCode = -1

	// Operation was successful.
	OK ResultCode = 0

	// Unknown server failure.
	SERVER_ERROR ResultCode = 1

	// On retrieving, touching or replacing a record that doesn't exist.
	KEY_NOT_FOUND_ERROR ResultCode = 2

	// On modifying a record with unexpected generation.
	GENERATION_ERROR ResultCode = 3

	// Bad parameter(s) were passed in database operation call.
	PARAMETER_ERROR ResultCode = 4

	// On create-only (write unique) operations on a record that already
	// exists.
	KEY_EXISTS_ERROR ResultCode = 5

	// On create-only (write unique) operations on a bin that already
	// exists.
	BIN_EXISTS_ERROR ResultCode = 6

	// Expected cluster ID was not received.
	CLUSTER_KEY_MISMATCH ResultCode = 7

	// Server has run out of memory.
	SERVER_MEM_ERROR ResultCode = 8

	// Client or server has timed out.
	TIMEOUT ResultCode = 9

	// XDS product is not available.
	NO_XDS ResultCode = 10

	// Server is not accepting requests.
	SERVER_NOT_AVAILABLE ResultCode = 11

	// Operation is not supported with configured bin type (single-bin or
	// multi-bin).
	BIN_TYPE_ERROR ResultCode = 12

	// Record size exceeds limit.
	RECORD_TOO_BIG ResultCode = 13

	// Too many concurrent operations on the same record.
	KEY_BUSY ResultCode = 14

	// Scan aborted by server.
	SCAN_ABORT ResultCode = 15

	// Unsupported Server Feature (e.g. Scan + UDF)
	UNSUPPORTED_FEATURE ResultCode = 16

	// Specified bin name does not exist in record.
	BIN_NOT_FOUND ResultCode = 17

	// Specified bin name does not exist in record.
	DEVICE_OVERLOAD ResultCode = 18

	// Key type mismatch.
	KEY_MISMATCH ResultCode = 19

	// Invalid namespace.
	INVALID_NAMESPACE ResultCode = 20

	// Bin name length greater than 14 characters.
	BIN_NAME_TOO_LONG ResultCode = 21

	// Operation not allowed at this time.
	FAIL_FORBIDDEN ResultCode = 22

	// There are no more records left for query.
	QUERY_END ResultCode = 50

	SECURITY_NOT_SUPPORTED        ResultCode = 51
	SECURITY_NOT_ENABLED          ResultCode = 52
	SECURITY_SCHEME_NOT_SUPPORTED ResultCode = 53

	// Administration command is invalid.
	INVALID_COMMAND ResultCode = 54

	// Administration field is invalid.
	INVALID_FIELD ResultCode = 55

	ILLEGAL_STATE ResultCode = 56

	// User name is invalid.
	INVALID_USER ResultCode = 60

	// User was previously created.
	USER_ALREADY_EXISTS ResultCode = 61

	// Password is invalid.
	INVALID_PASSWORD ResultCode = 62

	// Security credential is invalid.
	EXPIRED_PASSWORD ResultCode = 63

	// Forbidden password (e.g. recently used)
	FORBIDDEN_PASSWORD ResultCode = 64

	// Security credential is invalid.
	INVALID_CREDENTIAL ResultCode = 65

	// Role name is invalid.
	INVALID_ROLE ResultCode = 70

	// Role already exists.
	ROLE_ALREADY_EXISTS ResultCode = 71

	// Privilege is invalid.
	INVALID_PRIVILEGE ResultCode = 72

	// User must be authentication before performing database operations.
	NOT_AUTHENTICATED ResultCode = 80

	// User does not posses the required role to perform the database operation.
	ROLE_VIOLATION ResultCode = 81

	// A user defined function returned an error code.
	UDF_BAD_RESPONSE ResultCode = 100

	// The requested item in a large collection was not found.
	LARGE_ITEM_NOT_FOUND ResultCode = 125

	// Secondary index already exists.
	INDEX_FOUND ResultCode = 200

	// Requested secondary index does not exist.
	INDEX_NOTFOUND ResultCode = 201

	// Secondary index memory space exceeded.
	INDEX_OOM ResultCode = 202

	// Secondary index not available.
	INDEX_NOTREADABLE ResultCode = 203

	// Generic secondary index error.
	INDEX_GENERIC ResultCode = 204

	// Index name maximum length exceeded.
	INDEX_NAME_MAXLEN ResultCode = 205

	// Maximum number of indicies exceeded.
	INDEX_MAXCOUNT ResultCode = 206

	// Secondary index query aborted.
	QUERY_ABORTED ResultCode = 210

	// Secondary index queue full.
	QUERY_QUEUEFULL ResultCode = 211

	// Secondary index query timed out on server.
	QUERY_TIMEOUT ResultCode = 212

	// Generic query error.
	QUERY_GENERIC ResultCode = 213
)

var (
	TYPE_NOT_SUPPORTED_ERROR            = errors.New("Type cannot be converted to Value Type.")
	COMMAND_REJECTED_ERROR              = errors.New("command rejected")
	QUERY_TERMINATED_ERROR              = errors.New("Query terminated")
	SCAN_TERMINATED_ERROR               = errors.New("Scan terminated")
	INVALID_NODE_ERROR_ERROR            = errors.New("Invalid node")
	PARSE_ERROR_ERROR                   = errors.New("Parse error")
	SERIALIZE_ERROR_ERROR               = errors.New("Serialize error")
	OK_ERROR                            = errors.New("ok")
	SERVER_ERROR_ERROR                  = errors.New("Server error")
	KEY_NOT_FOUND_ERROR_ERROR           = errors.New("Key not found")
	GENERATION_ERROR_ERROR              = errors.New("Generation error")
	PARAMETER_ERROR_ERROR               = errors.New("Parameter error")
	KEY_EXISTS_ERROR_ERROR              = errors.New("Key already exists")
	BIN_EXISTS_ERROR_ERROR              = errors.New("Bin already exists")
	CLUSTER_KEY_MISMATCH_ERROR          = errors.New("Cluster key mismatch")
	SERVER_MEM_ERROR_ERROR              = errors.New("Server memory error")
	TIMEOUT_ERROR                       = errors.New("Timeout")
	NO_XDS_ERROR                        = errors.New("XDS not available")
	SERVER_NOT_AVAILABLE_ERROR          = errors.New("Server not available")
	BIN_TYPE_ERROR_ERROR                = errors.New("Bin type error")
	RECORD_TOO_BIG_ERROR                = errors.New("Record too big")
	KEY_BUSY_ERROR                      = errors.New("Hot key")
	SCAN_ABORT_ERROR                    = errors.New("Scan aborted")
	UNSUPPORTED_FEATURE_ERROR           = errors.New("Unsupported Server Feature")
	BIN_NOT_FOUND_ERROR                 = errors.New("Bin not found")
	DEVICE_OVERLOAD_ERROR               = errors.New("Device overload")
	KEY_MISMATCH_ERROR                  = errors.New("Key mismatch")
	INVALID_NAMESPACE_ERROR             = errors.New("Namespace not found")
	BIN_NAME_TOO_LONG_ERROR             = errors.New("Bin name length greater than 14 characters")
	FAIL_FORBIDDEN_ERROR                = errors.New("Operation not allowed at this time")
	QUERY_END_ERROR                     = errors.New("Query end")
	SECURITY_NOT_SUPPORTED_ERROR        = errors.New("Security not supported")
	SECURITY_NOT_ENABLED_ERROR          = errors.New("Security not enabled")
	SECURITY_SCHEME_NOT_SUPPORTED_ERROR = errors.New("Security scheme not supported")
	INVALID_COMMAND_ERROR               = errors.New("Invalid command")
	INVALID_FIELD_ERROR                 = errors.New("Invalid field")
	ILLEGAL_STATE_ERROR                 = errors.New("Illegal state")
	INVALID_USER_ERROR                  = errors.New("Invalid user")
	USER_ALREADY_EXISTS_ERROR           = errors.New("User already exists")
	INVALID_PASSWORD_ERROR              = errors.New("Invalid password")
	EXPIRED_PASSWORD_ERROR              = errors.New("Expired password")
	FORBIDDEN_PASSWORD_ERROR            = errors.New("Forbidden password")
	INVALID_CREDENTIAL_ERROR            = errors.New("Invalid credential")
	INVALID_ROLE_ERROR                  = errors.New("Invalid role")
	ROLE_ALREADY_EXISTS_ERROR           = errors.New("Role already exists")
	INVALID_PRIVILEGE_ERROR             = errors.New("Invalid privilege")
	NOT_AUTHENTICATED_ERROR             = errors.New("Not authenticated")
	ROLE_VIOLATION_ERROR                = errors.New("Role violation")
	UDF_BAD_RESPONSE_ERROR              = errors.New("UDF returned error")
	LARGE_ITEM_NOT_FOUND_ERROR          = errors.New("Large collection item not found")
	INDEX_FOUND_ERROR                   = errors.New("Index already exists")
	INDEX_NOTFOUND_ERROR                = errors.New("Index not found")
	INDEX_OOM_ERROR                     = errors.New("Index out of memory")
	INDEX_NOTREADABLE_ERROR             = errors.New("Index not readable")
	INDEX_GENERIC_ERROR                 = errors.New("Index error")
	INDEX_NAME_MAXLEN_ERROR             = errors.New("Index name max length exceeded")
	INDEX_MAXCOUNT_ERROR                = errors.New("Index count exceeds max")
	QUERY_ABORTED_ERROR                 = errors.New("Query aborted")
	QUERY_QUEUEFULL_ERROR               = errors.New("Query queue full")
	QUERY_TIMEOUT_ERROR                 = errors.New("Query timeout")
	QUERY_GENERIC_ERROR                 = errors.New("Query error")
	NO_ERROR_MESSAGE_ERROR              = errors.New("Error message not available yet - please file an issue on github.")
)

// Should connection be put back into pool.
func KeepConnection(resultCode int) bool {
	switch ResultCode(resultCode) {
	case OK, // Exception did not originate on server.
		QUERY_TERMINATED,
		SCAN_TERMINATED,
		INVALID_NODE_ERROR,
		PARSE_ERROR,
		SERIALIZE_ERROR,
		SERVER_MEM_ERROR,
		TIMEOUT,
		SERVER_NOT_AVAILABLE,
		SCAN_ABORT,
		INDEX_OOM,
		QUERY_ABORTED,
		QUERY_TIMEOUT:
		return false

	default:
		return true
	}
}

// Return result code as an error.
func ResultCodeToError(resultCode ResultCode) error {
	switch ResultCode(resultCode) {
	case TYPE_NOT_SUPPORTED:
		return TYPE_NOT_SUPPORTED_ERROR

	case COMMAND_REJECTED:
		return COMMAND_REJECTED_ERROR

	case QUERY_TERMINATED:
		return QUERY_TERMINATED_ERROR

	case SCAN_TERMINATED:
		return SCAN_TERMINATED_ERROR

	case INVALID_NODE_ERROR:
		return INVALID_NODE_ERROR_ERROR

	case PARSE_ERROR:
		return PARSE_ERROR_ERROR

	case SERIALIZE_ERROR:
		return SERIALIZE_ERROR_ERROR

	case OK:
		return OK_ERROR

	case SERVER_ERROR:
		return SERVER_ERROR_ERROR

	case KEY_NOT_FOUND_ERROR:
		return KEY_NOT_FOUND_ERROR_ERROR

	case GENERATION_ERROR:
		return GENERATION_ERROR_ERROR

	case PARAMETER_ERROR:
		return PARAMETER_ERROR_ERROR

	case KEY_EXISTS_ERROR:
		return KEY_EXISTS_ERROR_ERROR

	case BIN_EXISTS_ERROR:
		return BIN_EXISTS_ERROR_ERROR

	case CLUSTER_KEY_MISMATCH:
		return CLUSTER_KEY_MISMATCH_ERROR

	case SERVER_MEM_ERROR:
		return SERVER_MEM_ERROR_ERROR

	case TIMEOUT:
		return TIMEOUT_ERROR

	case NO_XDS:
		return NO_XDS_ERROR

	case SERVER_NOT_AVAILABLE:
		return SERVER_NOT_AVAILABLE_ERROR

	case BIN_TYPE_ERROR:
		return BIN_TYPE_ERROR_ERROR

	case RECORD_TOO_BIG:
		return RECORD_TOO_BIG_ERROR

	case KEY_BUSY:
		return KEY_BUSY_ERROR

	case SCAN_ABORT:
		return SCAN_ABORT_ERROR

	case UNSUPPORTED_FEATURE:
		return UNSUPPORTED_FEATURE_ERROR

	case BIN_NOT_FOUND:
		return BIN_NOT_FOUND_ERROR

	case DEVICE_OVERLOAD:
		return DEVICE_OVERLOAD_ERROR

	case KEY_MISMATCH:
		return KEY_MISMATCH_ERROR

	case INVALID_NAMESPACE:
		return INVALID_NAMESPACE_ERROR

	case BIN_NAME_TOO_LONG:
		return BIN_NAME_TOO_LONG_ERROR

	case FAIL_FORBIDDEN:
		return FAIL_FORBIDDEN_ERROR

	case QUERY_END:
		return QUERY_END_ERROR

	case SECURITY_NOT_SUPPORTED:
		return SECURITY_NOT_SUPPORTED_ERROR

	case SECURITY_NOT_ENABLED:
		return SECURITY_NOT_ENABLED_ERROR

	case SECURITY_SCHEME_NOT_SUPPORTED:
		return SECURITY_SCHEME_NOT_SUPPORTED_ERROR

	case INVALID_COMMAND:
		return INVALID_COMMAND_ERROR

	case INVALID_FIELD:
		return INVALID_FIELD_ERROR

	case ILLEGAL_STATE:
		return ILLEGAL_STATE_ERROR

	case INVALID_USER:
		return INVALID_USER_ERROR

	case USER_ALREADY_EXISTS:
		return USER_ALREADY_EXISTS_ERROR

	case INVALID_PASSWORD:
		return INVALID_PASSWORD_ERROR

	case EXPIRED_PASSWORD:
		return EXPIRED_PASSWORD_ERROR

	case FORBIDDEN_PASSWORD:
		return FORBIDDEN_PASSWORD_ERROR

	case INVALID_CREDENTIAL:
		return INVALID_CREDENTIAL_ERROR

	case INVALID_ROLE:
		return INVALID_ROLE_ERROR

	case ROLE_ALREADY_EXISTS:
		return ROLE_ALREADY_EXISTS_ERROR

	case INVALID_PRIVILEGE:
		return INVALID_PRIVILEGE_ERROR

	case NOT_AUTHENTICATED:
		return NOT_AUTHENTICATED_ERROR

	case ROLE_VIOLATION:
		return ROLE_VIOLATION_ERROR

	case UDF_BAD_RESPONSE:
		return UDF_BAD_RESPONSE_ERROR

	case LARGE_ITEM_NOT_FOUND:
		return LARGE_ITEM_NOT_FOUND_ERROR

	case INDEX_FOUND:
		return INDEX_FOUND_ERROR

	case INDEX_NOTFOUND:
		return INDEX_NOTFOUND_ERROR

	case INDEX_OOM:
		return INDEX_OOM_ERROR

	case INDEX_NOTREADABLE:
		return INDEX_NOTREADABLE_ERROR

	case INDEX_GENERIC:
		return INDEX_GENERIC_ERROR

	case INDEX_NAME_MAXLEN:
		return INDEX_NAME_MAXLEN_ERROR

	case INDEX_MAXCOUNT:
		return INDEX_MAXCOUNT_ERROR

	case QUERY_ABORTED:
		return QUERY_ABORTED_ERROR

	case QUERY_QUEUEFULL:
		return QUERY_QUEUEFULL_ERROR

	case QUERY_TIMEOUT:
		return QUERY_TIMEOUT_ERROR

	case QUERY_GENERIC:
		return QUERY_GENERIC_ERROR

	default:
		return NO_ERROR_MESSAGE_ERROR
	}
}
