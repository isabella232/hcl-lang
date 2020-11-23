package schema

import (
	"fmt"
	"strings"

	"github.com/zclconf/go-cty/cty"
)

type ScopeId string

type ExprSchema []Expr

func (es ExprSchema) TypeHint() (cty.Type, bool) {
	if es == nil {
		return cty.NilType, false
	}

	for _, expr := range es {
		switch e := expr.(type) {
		case LiteralValueExpr:
			if e.Type != cty.NilType {
				return e.Type, true
			}
		case TupleExpr:
			return cty.Tuple([]cty.Type{}), true
		}
	}

	return cty.NilType, false
}

func (es ExprSchema) FriendlyNames() []string {
	if es == nil {
		return []string{}
	}

	return friendlyNames(es)
}

func friendlyNames(expressions []Expr) []string {
	names := make([]string, 0)

	for _, expr := range expressions {
		switch e := expr.(type) {
		case LiteralValueExpr:
			if e.Type != cty.NilType {
				names = append(names, e.Type.FriendlyName())
			}
		case TupleExpr:
			ofType := ""
			types := friendlyNames(e.Exprs)
			if len(types) > 0 {
				ofType = strings.Join(types, " or ")
			}
			names = append(names, fmt.Sprintf("tuple of %s", ofType))
		case ObjectExpr:
			names = append(names, "object")
		case MapExpr:
			names = append(names, "map")
		}
		// TODO: ScopeTraversalExpr (scope names?)
	}

	return names
}

type exprSigil struct{}

type Expr interface {
	isExprImpl() exprSigil
}

// LiteralValueExpr expresses literal value of the given type
// (if provided), such as "str", 43 or true
type LiteralValueExpr struct {
	Type cty.Type
	// TODO: StaticCandidates / DynamicCandidates
}

func (LiteralValueExpr) isExprImpl() exprSigil {
	return exprSigil{}
}

// ScopeTraversalExpr expresses traversal of the given scope and type,
// such as `data.aws_instance.public_ip` or `var.name`
type ScopeTraversalExpr struct {
	ScopeId ScopeId
	OfType  cty.Type
}

func (ScopeTraversalExpr) isExprImpl() exprSigil {
	return exprSigil{}
}

// TupleExpr expresses a tuple of given expressions,
// such as `[ aws_instance.vpn, data.eip.vpn ]`
// TODO: Change to VariadicSetExpr ?
type TupleExpr struct {
	Exprs []Expr
}

func (TupleExpr) isExprImpl() exprSigil {
	return exprSigil{}
}

// ObjectExpr expresses an object with predeclared keys & expressions, e.g.
//   provider_requirements = {
//	   source  = "registry.terraform.io/hashicorp/aws"
//     version = "~> 1.0.0"
//   }
type ObjectExpr struct {
	Attributes map[string]ExprSchema
}

func (ObjectExpr) isExprImpl() exprSigil {
	return exprSigil{}
}

// MapExpr expresses a map with predeclared key and expression, e.g.
//   providers = {
//     google  = google.uswest
//     azurerm = azurerm.euwest
//   }
type MapExpr struct {
	Key  Expr
	Expr Expr
}

func (MapExpr) isExprImpl() exprSigil {
	return exprSigil{}
}
