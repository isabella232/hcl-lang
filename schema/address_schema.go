package schema

type Address []AddrStep

type addrStepSigil struct{}

type AddrStep interface {
	addrStepImpl() addrStepSigil
}

type StaticStep struct {
	Value string
}

func (StaticStep) addrStepImpl() addrStepSigil {
	return addrStepSigil{}
}

type LabelValueStep struct {
	Index int
}

func (LabelValueStep) addrStepImpl() addrStepSigil {
	return addrStepSigil{}
}

type AttrNameStep struct{}

func (AttrNameStep) addrStepImpl() addrStepSigil {
	return addrStepSigil{}
}

type AttrValueStep struct {
	AttrName string
}

func (AttrValueStep) addrStepImpl() addrStepSigil {
	return addrStepSigil{}
}
