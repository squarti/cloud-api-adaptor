// (C) Copyright Confidential Containers Contributors
// SPDX-License-Identifier: Apache-2.0

package ibm

import (
	"github.com/confidential-containers/cloud-api-adaptor/src/cloud-providers/util"
)

type Config struct {
}

func (c Config) Redact() Config {
	return *util.RedactStruct(&c).(*Config)
}
