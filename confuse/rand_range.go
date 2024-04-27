// Copyright [2022] [thinkgos] thinkgo@aliyun.com
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

package confuse

import "math/rand/v2"

// IntR returns, as an int, a non-negative pseudo-random number in the half-open interval [min,max).
// It panics if max < min.
func IntR(min, max int) int {
	if min > max {
		panic("confuse: invalid argument to IntR")
	}
	if min == max {
		return min
	}
	return rand.IntN(max-min) + min
}

// Int32R returns, as an int32, a non-negative pseudo-random number in the half-open interval [min,max).
// It panics if max < min.
func Int32R(min, max int32) int32 {
	if min > max {
		panic("confuse: invalid argument to Int32R")
	}
	if min == max {
		return min
	}
	return rand.Int32N(max-min) + min
}

// Int64R returns, as an int64, a non-negative pseudo-random number in the half-open interval [min,max).
// It panics if max < min.
func Int64R(min, max int64) int64 {
	if min > max {
		panic("confuse: invalid argument to Int64R")
	}
	if min == max {
		return min
	}
	return rand.Int64N(max-min) + min
}

// Float64R returns, as an float64, a non-negative pseudo-random number in the half-open interval [min,max).
// It panics if max < min.
func Float64R(min, max float64) float64 {
	if min > max {
		panic("confuse: invalid argument to Float64R")
	}
	return min + (max-min)*rand.Float64()
}

// UintR returns, as an uint, a non-negative pseudo-random number in the half-open interval [min,max).
// It panics if max < min.
func UintR(min, max uint) uint {
	if min > max {
		panic("confuse: invalid argument to UintR")
	}
	if min == max {
		return min
	}
	return rand.UintN(max-min) + min
}

// Uint32R returns, as an uint32, a non-negative pseudo-random number in the half-open interval [min,max).
// It panics if max < min.
func Uint32R(min, max uint32) uint32 {
	if min > max {
		panic("confuse: invalid argument to Uint32R")
	}
	if min == max {
		return min
	}
	return rand.Uint32N(max-min) + min
}

// Uint64R returns, as an uint64, a non-negative pseudo-random number in the half-open interval [min,max).
// It panics if max < min.
func Uint64R(min, max uint64) uint64 {
	if min > max {
		panic("confuse: invalid argument to Uint64R")
	}
	if min == max {
		return min
	}
	return rand.Uint64N(max-min) + min
}

type intType interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// NR returns a pseudo-random number in the half-open interval [min,max) from the default Source.
// The type parameter Int can be any integer type.
// It panics if max < min.
func NR[Int intType](min, max Int) Int {
	if min > max {
		panic("confuse: invalid argument to NR")
	}
	if min == max {
		return min
	}
	return rand.N(max-min) + min
}
