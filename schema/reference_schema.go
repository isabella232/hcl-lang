package schema

type AttrReference struct {
	ScopeId ScopeId
	Address Address
	Type    ReferenceType

	NotReferable bool
}

type BlockReference struct {
	ScopeId ScopeId
	Address Address
	Type    ReferenceType

	ReferBody    bool
	NotReferable bool
}

// TODO: context-sensitive refs
// - self.<attr>
// - each.key, each.value
// - count.index

type refTypeSigil struct{}

type ReferenceType interface {
	refTypeImpl() refTypeSigil
}

type ReferenceTypes []ReferenceType

func (ReferenceTypes) refTypeImpl() refTypeSigil {
	return refTypeSigil{}
}

type InferredRefType struct {
	AttrName string
}

func (*InferredRefType) refTypeImpl() refTypeSigil {
	return refTypeSigil{}
}

type RefTypeFromConstraint struct {
	AttrName string
}

func (*RefTypeFromConstraint) refTypeImpl() refTypeSigil {
	return refTypeSigil{}
}
