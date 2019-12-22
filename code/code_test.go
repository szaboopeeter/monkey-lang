package code

import (
	"testing"
)

func TestMake(t *testing.T) {
	tests := []struct {
		op       Opcode
		operands []int
		expected []byte
	}{
		{OpConstant, []int{65534}, []byte{byte(OpConstant), 255, 254}},
		{OpAdd, []int{}, []byte{byte(OpAdd)}},
	}

	for _, tt := range tests {
		instruction := Make(tt.op, tt.operands...)

		if len(instruction) != len(tt.expected) {
			t.Errorf("Instruction has incorrect length. Expected=%d, got=%d", len(tt.expected), len(instruction))
		}

		for i, b := range tt.expected {
			if instruction[i] != tt.expected[i] {
				t.Errorf("Incorrect byte at position %d. Expected=%d, got=%d", i, b, instruction[i])
			}
		}
	}
}

func TestInstructionsString(t *testing.T) {
	instructions := []Instructions{
		Make(OpAdd),
		Make(OpConstant, 2),
		Make(OpConstant, 65535),
	}

	expected := `0000 OpAdd
0001 OpConstant 2
0004 OpConstant 65535
`

	concatted := Instructions{}
	for _, instruction := range instructions {
		concatted = append(concatted, instruction...)
	}

	if concatted.String() != expected {
		t.Errorf("incorrectly formatted instructions.\nexpected=%q\nactual=%q", expected, concatted.String())
	}
}

func TestReadOperands(t *testing.T) {
	tests := []struct {
		op        Opcode
		operands  []int
		bytesRead int
	}{
		{OpConstant, []int{65535}, 2},
	}

	for _, tt := range tests {
		instruction := Make(tt.op, tt.operands...)

		def, err := Lookup(byte(tt.op))
		if err != nil {
			t.Fatalf("definition not found: %q", err)
		}

		operandsRead, offset := ReadOperands(def, instruction[1:])
		if offset != tt.bytesRead {
			t.Fatalf("incorrect offset. expected=%d, actual=%d", tt.bytesRead, offset)
		}

		for i, expected := range tt.operands {
			if operandsRead[i] != expected {
				t.Errorf("incorrect operand. expected=%d, actual=%d", expected, operandsRead[i])
			}
		}
	}
}
