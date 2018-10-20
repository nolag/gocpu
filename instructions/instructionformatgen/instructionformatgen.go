package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/nolag/gocpu/instructions/shared"
)

type formatPart struct {
	name     string
	capsName string
	offset   uint8
	size     uint8
	signed   bool
}

func writeGetter(part *formatPart, instruction string, builder bool, sw *shared.SimpleWriter) {
	name := instruction
	if builder {
		name = fmt.Sprintf("*%vBuilder", instruction)
	}

	sw.Writeln("// Get%v gets the %v of %v", part.capsName, part.name, instruction)
	sw.Writeln("func (instruction %v) Get%v {", name, part.capsName)
	sw.Write("return ")
	if builder {
		sw.Write("*")
	}
	sw.Writeln("instruciton >>")
	sw.Writeln("}")
}

func writeSetter(part *formatPart, instruction string, sw *shared.SimpleWriter) {
	sw.Writeln(
		"// Set%v sets the %v portion of the %v",
		part.capsName,
		part.name,
		instruction)
	sw.Writeln("func (builder *%vBuilder) Set%v {", instruction, part.capsName)
	// TODO xor with part that I want, shift one way then shift back for signed, sift for the one.
	sw.Writeln("}")
}

func verifyArgForNoError(err error, offset int, name string) {
	if err != nil {
		_ = fmt.Errorf("Error reading argument %v as a %v: %v", offset, name, err)
		os.Exit(1)
	}
}

func readSizeFromArgs(offset int, name string) uint8 {
	val, err := strconv.ParseUint(os.Args[offset], 10, 8)
	verifyArgForNoError(err, offset, name)
	return uint8(val)
}

// This file is tested via tests of the output
func main() {
	args := os.Args
	if len(args) < 7 || (len(args)-4)%3 != 0 {
		_ = fmt.Errorf(
			"Expected useage %v <output_file> <package> <format_name> <base_struct> [<part> <offset> <size> <signed>...]", os.Args[0])
		return
	}

	outputFile := args[1]
	sw := shared.CreateSimpleWriter(outputFile)
	defer sw.Close()

	sw.WriteGenHeader(
		"instructionformatgen",
		"https://github.com/nolag/gocpu/tree/master/instructions/instructionformatgen",
		args[2])

	instructionName := args[3]
	baseName := args[4]
	sw.Writeln("// %v reperesents an instruction of the %v format", instructionName, instructionName)
	sw.Writeln("struct %v %v", instructionName, baseName)
	sw.Writeln(
		"// %vBuilder reperesents a builder for an instruction instruction of the %v format",
		instructionName,
		instructionName)
	sw.Writeln("struct %vBuilider %v", instructionName, baseName)
	for i := 5; i < len(os.Args); i += 4 {
		format := formatPart{}
		format.name = args[i]
		format.capsName = strings.Title(format.name)
		format.offset = readSizeFromArgs(i+1, "offset")
		format.size = readSizeFromArgs(i+2, "size")
		signed, err := strconv.ParseBool(args[i+3])
		verifyArgForNoError(err, i+3, "signed")
		format.signed = signed
		writeGetter(&format, instructionName, false, sw)
		writeGetter(&format, instructionName, true, sw)
		writeSetter(&format, instructionName, sw)
	}
}
