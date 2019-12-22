package code

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Instructions []byte

func (instructions Instructions) String() string {
	var out bytes.Buffer

	i := 0
	for i < len(instructions) {
		definition, err := Lookup(instructions[i])
		if err != nil {
			fmt.Fprintf(&out, "ERROR: %s\n", err)
			continue
		}

		operands, read := ReadOperands(definition, instructions[i+1:])

		fmt.Fprintf(&out, "%04d %s\n", i, instructions.fmtInstruction(definition, operands))

		i += 1 + read
	}

	return out.String()
}

type Opcode byte

const (
	OpConstant Opcode = iota
	OpAdd
)

type Definition struct {
	Name          string
	OperandWidths []int
}

var definitions = map[Opcode]*Definition{
	OpConstant: {"OpConstant", []int{2}}, // OpConstant definiton: push a constant (single operand, which is 2 bytes long) to the stack
	OpAdd:      {"OpAdd", []int{}},       // OpAdd: pop the two topmost stack items, add them, and push the result (no operands)
}

func Lookup(op byte) (*Definition, error) {
	definition, ok := definitions[Opcode(op)]
	if !ok {
		return nil, fmt.Errorf("opcode %d is undefined", op)
	}

	return definition, nil
}

func Make(op Opcode, operands ...int) []byte {
	definition, ok := definitions[op]
	if !ok {
		return []byte{}
	}

	instructionLength := 1
	for _, width := range definition.OperandWidths {
		instructionLength += width
	}

	instruction := make([]byte, instructionLength)
	instruction[0] = byte(op)

	offset := 1
	for i, operand := range operands {
		width := definition.OperandWidths[i]
		switch width {
		case 2:
			binary.BigEndian.PutUint16(instruction[offset:], uint16(operand))
		}
		offset += width
	}

	return instruction
}

func ReadOperands(definition *Definition, instructions Instructions) ([]int, int) {
	operands := make([]int, len(definition.OperandWidths))
	offset := 0

	for i, width := range definition.OperandWidths {
		switch width {
		case 2:
			operands[i] = int(ReadUint16(instructions[offset:]))
		}

		offset += width
	}

	return operands, offset
}

func ReadUint16(Instructions Instructions) uint16 {
	return binary.BigEndian.Uint16(Instructions)
}

func (ins Instructions) fmtInstruction(definition *Definition, operands []int) string {
	operandCount := len(definition.OperandWidths)

	if len(operands) != operandCount {
		return fmt.Sprintf("ERROR: oparand length %d does not match the definition %d\n", len(operands), operandCount)
	}

	switch operandCount {
	case 0:
		return definition.Name
	case 1:
		return fmt.Sprintf("%s %d", definition.Name, operands[0])
	}

	return fmt.Sprintf("ERROR: number of operands for %s\n", definition.Name)
}