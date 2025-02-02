package et

import (
	"encore.dev/beta/auth"
)

// OverrideAuthInfo overrides the auth information for the current request.
// Subsequent calls to auth.UserID and auth.Data() within the same request
// will return the given uid and data, and API calls made from the request
// will propagate the newly set user info.
//
// Passing in an empty string as the uid results in unsetting the auth information,
// causing future API calls to behave as if there was no authenticated user.
//
// If the application's auth handler returns custom auth data, two additional
// requirements exist. First, the data parameter passed to WithContext must be of
// the same type as the auth handler returns. Second, if the uid argument is not
// the empty string then data may not be nil. If these requirements are not met,
// API calls made with these options will not be made and will immediately return
// a client-side error.
//
// OverrideAuthInfo is not safe for concurrent use with code that invokes
// auth.UserID or auth.Data() within the same request.
func OverrideAuthInfo(uid auth.UID, data any) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.27.0/runtime/et/pkgfn.go#L26-L28
	doPanic("encore apps must be run using the encore command")
	return
}
