package lang

import (
	"github.com/hashicorp/hcl/v2"
)

type Link struct {
	URI   string
	Range hcl.Range
}
