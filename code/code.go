package code

import "fmt"

type Instructions []byte
type Opcode byte

const (
	OpConstant Opcode = iota
)

type Definition struct {
	Name          string
	OperandWidths []int
}

var definitions = map[Opcode]*Definition{
	OpConstant: {"OpConstant", []int{2}}, // OpConstant definiton: my name is OpConstant, and I have one singe operand, which is 2 bytes long
}

func Lookup(op byte) (*Definition, error) {
	definition, ok := definitions[Opcode(op)]
	if !ok {
		return nil, fmt.Errorf("Opcode %d is undefined.", op)
	}

	return definition, nil
}
