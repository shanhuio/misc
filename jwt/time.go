// Copyright (C) 2021  Shanhu Tech Inc.
//
// This program is free software: you can redistribute it and/or modify it
// under the terms of the GNU Affero General Public License as published by the
// Free Software Foundation, either version 3 of the License, or (at your
// option) any later version.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.  See the GNU Affero General Public License
// for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package jwt

import (
	"time"

	"shanhu.io/misc/errcode"
)

// CheckTime checks if the token's claims is in valid at time now.
func CheckTime(claims *ClaimSet, now time.Time) (time.Duration, error) {
	// Issued time must be after now.
	// In case of small clock error, gives a 5 minute grace period.
	issued := time.Unix(claims.Iat, 0).Add(-5 * time.Minute)
	if !issued.Before(now) {
		return 0, errcode.Unauthorizedf("token issued in the future")
	}

	expires := time.Unix(claims.Exp, 0)
	if now.After(expires) {
		return 0, errcode.Unauthorizedf("token expired")
	}
	return expires.Sub(now), nil
}
