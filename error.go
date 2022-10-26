// Copyright 2022 Marcelo Cantos
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package linq

import "fmt"

type Error string

func errorf(format string, args ...any) Error {
	return Error(fmt.Sprintf(format, args...))
}

func (e Error) Error() string {
	return string(e)
}

const (
	EmptySourceError  Error = "empty source"
	NoFastCountError  Error = "fast count unavailable"
	ZeroIotaStepError Error = "iota step is zero"
	NotSingleError    Error = "count != 1"
	NoValueError      Error = "no value"
)
