package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// REGISTER

var INSTRUCTION_SIZE = 8
var V = make([]int, 16)
var (
	_ii   = 0
	PC    = 0x0
	I     = 0x200
	BK    = 0x0
	PAUSE = false
	STACK = make([]int, 16)
	SP    = 0
	GFX   = make([][]int, 64)
	FONT  = []int{
		0xF0, 0x90, 0x90, 0x90, 0xF0, // 0
		0x20, 0x60, 0x20, 0x20, 0x70, // 1
		0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
		0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
		0x90, 0x90, 0xF0, 0x10, 0x10, // 4
		0xF0, 0x80, 0xF0, 0x10, 0xF0, // 5
		0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
		0xF0, 0x10, 0x20, 0x40, 0x40, // 7
		0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
		0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
		0xF0, 0x90, 0xF0, 0x90, 0x90, // A
		0xE0, 0x90, 0xE0, 0x90, 0xE0, // B
		0xF0, 0x80, 0x80, 0x80, 0xF0, // C
		0xE0, 0x90, 0x90, 0x90, 0xE0, // D
		0xF0, 0x80, 0xF0, 0x80, 0xF0, // E
		0xF0, 0x80, 0xF0, 0x80, 0x80, // F
	}
	MEMORY      = make([]string, 1024)
	DELAY_TIMER = 0x0
	SOUND_TIMER = 0x0
)

func OnByte(X1 byte) int {
	n, e := strconv.ParseInt(fmt.Sprintf("%s", string(X1)), 16, 8)
	if e != nil {
		panic(e)
	}
	return int(n) & 0xFF
}
func ToByte(X1, X2 byte) int {
	n, e := strconv.ParseInt(fmt.Sprintf("%s%s", string(X1), string(X2)), 16, 16)
	if e != nil {
		panic(e)
	}
	return int(n) & 0xFF
}

func ThByte(X1, X2, X3 byte) int {
	n, e := strconv.ParseInt(fmt.Sprintf("%s%s%s", string(X1), string(X2), string(X3)), 16, 16)
	if e != nil {
		panic(e)
	}
	return int(n) & 0xFFF
}

func INST_0(instruction string) {

	switch instruction {
	case "00E0":
		fmt.Println("0x0000 ", "00E0")
		for i := 0; i < 64; i++ {
			for j := 0; j < 32; j++ {
				GFX[i][j] = 0
			}
		}

	case "00EE":
		// return from subroutin
		fmt.Println("0x0000 ", "00EE", " SP ", SP, " STACK[SP] ", STACK[SP])
		PC = STACK[SP-1]
		SP = SP - 1

	}
}
func INST_1(instruction string) {
	fmt.Println("0x1111 ", instruction[1:])
	PC = ThByte(instruction[1], instruction[2], instruction[3])
	PC -= 2
}

func INST_2(instruction string) {
	// call subroutin at NNN
	STACK[SP] = int(PC)
	SP = SP + 1
	PC = ThByte(instruction[1], instruction[2], instruction[3])
	PC -= 2
	fmt.Println("0x2222 ", instruction[1:], " STACK ", STACK, " SP ", SP, " PC ", PC, " TH ", ThByte(instruction[1], instruction[2], instruction[3]))
}

func INST_3(instruction string) {
	fmt.Println("0x3333 ", instruction[1:])
	if V[OnByte(instruction[1])] == ToByte(instruction[2], instruction[3]) {
		PC += 2
	}
}

func INST_4(instruction string) {
	fmt.Println("0x4444 ", instruction[1:])
	if V[OnByte(instruction[1])] != ToByte(instruction[2], instruction[3]) {
		PC += 2
	}
}

func INST_5(instruction string) {
	fmt.Println("0x5555 ", instruction[1:])
	if V[OnByte(instruction[1])] == V[OnByte(instruction[2])] {
		PC += 2
	}
}

func INST_6(instruction string) {
	fmt.Println("0x6666 ", instruction[1:], instruction[1], ToByte(instruction[2], instruction[3]))
	V[OnByte(instruction[1])] = ToByte(instruction[2], instruction[3]) & 0xFF

}

func INST_7(instruction string) {
	fmt.Println("0x7777 ", instruction)
	V[OnByte(instruction[1])] = (V[OnByte(instruction[1])] + ToByte(instruction[2], instruction[3])) & 0xFF

}
func INST_8_0(instruction string) {
	fmt.Println("0x8800 ", instruction[1:])
	V[OnByte(instruction[1])] = V[OnByte(instruction[2])]
}
func INST_8_1(instruction string) {
	V[OnByte(instruction[1])] = (V[OnByte(instruction[1])] | V[OnByte(instruction[2])]) & 0xFF
	fmt.Println("0x8811 ", instruction[1:])
}

func INST_8_2(instruction string) {
	fmt.Println("0x8822 ", instruction[1:])
	V[OnByte(instruction[1])] = (V[OnByte(instruction[1])] & V[OnByte(instruction[2])]) & 0xFF

}

func INST_8_3(instruction string) {
	fmt.Println("0x8833 ", instruction[1:])
	V[OnByte(instruction[1])] = (V[OnByte(instruction[1])] ^ V[OnByte(instruction[2])]) & 0xFF

}

func INST_8_4(instruction string) {

	fmt.Println("0x8844 ", instruction[1:])
	intsum := (V[OnByte(instruction[1])] + V[OnByte(instruction[2])])
	boo := (intsum & 0xFFFFFF00)
	if boo != 0 {
		V[0xF] = 0x1
	} else {
		V[0xF] = 0x0
	}
	V[OnByte(instruction[1])] = (V[OnByte(instruction[1])] + V[OnByte(instruction[2])]) & 0xFF

}

func INST_8_5(instruction string) {
	fmt.Println("0x8855 ", instruction[1:])
	intsum := (V[OnByte(instruction[1])] - V[OnByte(instruction[2])])

	if intsum < 0 {
		V[0xF] = 0x1
	} else {
		V[0xF] = 0x0
	}
	V[OnByte(instruction[1])] = (V[OnByte(instruction[1])] - V[OnByte(instruction[2])]) & 0xFF

}

func INST_8_6(instruction string) {
	fmt.Println("0x8866 ", instruction[1:])
	V[0xF] = V[OnByte(instruction[1])] & 0x1
	V[OnByte(instruction[1])] = (V[OnByte(instruction[1])] >> 1) & 0xFF
}

func INST_8_7(instruction string) {
	fmt.Println("0x8877 ", instruction[1:])
	V[0xF] = ((V[OnByte(instruction[2])] - V[OnByte(instruction[1])]) >> 8) ^ 0x1
	V[OnByte(instruction[1])] = (V[OnByte(instruction[2])] - V[OnByte(instruction[1])]) & 0xFF

}

func INST_8_E(instruction string) {
	os.Exit(2)
	fmt.Println("0x88EE ", instruction[1:])
	V[0xF] = (V[OnByte(instruction[1])] >> 7) & 0x1
	V[OnByte(instruction[1])] = (V[OnByte(instruction[1])] << 1) & 0xFF
}

func INST_8(instruction string) {
	switch instruction[3] {
	case '0':
		INST_8_0(instruction)
	case '1':
		INST_8_1(instruction)
	case '2':
		INST_8_2(instruction)
	case '3':
		INST_8_3(instruction)
	case '4':
		INST_8_4(instruction)
	case 5:
		INST_8_5(instruction)
	case '6':
		INST_8_6(instruction)
	case '7':
		INST_8_7(instruction)
	case 'E':
		INST_8_E(instruction)

	}
}

func INST_9(instruction string) {
	fmt.Println("0x9999 ", instruction[1:])

	if V[OnByte(instruction[1])] != V[OnByte(instruction[2])] {
		PC += 2
	}
}

func INST_A(instruction string) {
	fmt.Println("0xAAAA ", instruction[1:], ThByte(instruction[1], instruction[2], instruction[3]))

	I = ThByte(instruction[1], instruction[2], instruction[3]) & 0x0FFF
}

func INST_B(instruction string) {
	fmt.Println("0xBBBB ", instruction[1:])

	PC = (V[0] + V[OnByte(instruction[1])] + V[OnByte(instruction[2])] + V[OnByte(instruction[3])]) & 0xFF
}

func INST_C(instruction string) {
	fmt.Println("0xCCCC ", instruction[1:])

	V[OnByte(instruction[1])] = (rand.Intn(255) * (V[OnByte(instruction[1])] + V[OnByte(instruction[2])])) & 0xFF
}

func Draw(x, y, z int) {

	_x := V[x]
	_y := V[y]

	_height := z

	V[0xF] = 0x0
	for yline := _y; yline < _y+_height; yline++ {
		n, _ := strconv.ParseInt(MEMORY[I+yline-_y], 16, 16)
		pixel := int(n) //& 0x0000FFFF
		for xline := _x; xline < _x+8; xline++ {

			if (pixel & (0x80 >> uint(xline-_x))) != 0 {
				if xline > 63 || yline > 31 {
					fmt.Println("XXX ", " YYY ", yline)
					os.Exit(2)
					continue
				}
				if GFX[xline][yline] == 1 {
					GFX[xline][yline] = 0
					V[0xF] = 0x1
				} else {
					GFX[xline][yline] = 1
				}
			}
		}
	}

	//os.Exit(2)
}
func INST_DRAW(pc int) {

	fmt.Print("\033[0m\033[1J \033[0;0H \033[J \033[0m\033[1J\n")
	for j := 0; j < 32; j++ {
		for i := 0; i < 64; i++ {
			if GFX[i][j] == 1 {
				fmt.Print(" ⬜️")
			} else {
				//fmt.Print("0")
				fmt.Print("  ")

			}
			//			fmt.Print(GFX[i][j])
		}
		fmt.Println()
	}
	fmt.Println("----", _ii, "----", pc, " ", MEMORY[pc], MEMORY[pc+1], " V ", V, "------")

}

func INST_D(instruction string) {
	fmt.Println("0xDDDD ", instruction[1:])

	Draw(OnByte(instruction[1]), OnByte(instruction[2]), OnByte(instruction[3]))
}

func INST_E(instruction string) {
	fmt.Println("0xEEEE ", instruction[1:])

	if instruction[3] == 0xE {
		if V[OnByte(instruction[1])] == 0 { //PRESSED

		}
	} else if instruction[3] == 0x1 {
		if V[OnByte(instruction[1])] != 0 { //NOT PRESSED

		}
	}
}

func getKey() {

}
func INST_F(instruction string) {

	if instruction[3] == '7' {
		fmt.Println("0xFX07 ", instruction[1:])

		V[OnByte(instruction[1])] = DELAY_TIMER
	} else if instruction[3] == 'A' {
		fmt.Println("0xFX0A ", instruction[1:])
		getKey()
		PAUSE = true
	} else if instruction[3] == '5' && instruction[2] == '1' {
		fmt.Println("0xFX15 ", instruction[1:])

		DELAY_TIMER = V[OnByte(instruction[1])]
	} else if instruction[3] == '8' && instruction[2] == '1' {
		fmt.Println("0xFX18 ", instruction[1:])

		SOUND_TIMER = V[OnByte(instruction[1])]
	} else if instruction[3] == 'E' {
		fmt.Println("0xFX1E ", instruction[1:])

		I = (I + V[OnByte(instruction[1])]) & 0xFFF
	} else if instruction[3] == '9' {
		fmt.Println("0xFX29 ", instruction[1:])

		I = V[OnByte(instruction[1])] & 0xFFF
	} else if instruction[3] == '3' {
		fmt.Println("0xFX33 ", instruction[1:])

		num := float64(V[OnByte(instruction[1])])

		MEMORY[I+0] = fmt.Sprintf("%02X", int(num/100)%10)
		MEMORY[I+1] = fmt.Sprintf("%02X", int(num/10)%10)
		MEMORY[I+2] = fmt.Sprintf("%02X", int(num)%10)
		fmt.Println("BCD ::::::::> ", float64(V[OnByte(instruction[1])]), " -- ", MEMORY[I]+MEMORY[I+1]+MEMORY[I+2])
	} else if instruction[3] == '8' {
		fmt.Println("0xFX18 ", instruction[1:])

		SOUND_TIMER = V[OnByte(instruction[1])]
	} else if instruction[2] == '5' && instruction[3] == '5' {
		fmt.Println("0xFX55 ", instruction[1:])

		for i := 0; i <= OnByte(instruction[1]); i++ {
			MEMORY[I+i] = fmt.Sprintf("%02X", V[i])
		}
	} else if instruction[2] == '6' && instruction[3] == '5' {
		fmt.Println("0xFX65 ", instruction[1:])

		for i := 0; i <= OnByte(instruction[1]); i++ {
			n, _ := strconv.ParseInt(MEMORY[I+i], 16, 8)
			fmt.Println(MEMORY[660:670])
			fmt.Println(I+i, "__", "<<< I:", (I+i)&0xFFF, " <<", MEMORY[I+i], "<<", n)
			V[i] = int(n) & 0xFF
		}
	} else {
		errors.New("math: square root of negative number")
	}
}

func main() {
	f, err := os.Open("ROMS/TICTAC")

	if err != nil {
		panic(err)
	}
	defer f.Close()

	for i := 0; i < 1024; i++ {
		MEMORY[i] = "00"
	}

	for i := 0; i < 16; i++ {
		V[i] = 0
		STACK[i] = 0
	}

	for i := 0; i < 64; i++ {
		GFX[i] = make([]int, 32)
	}
	for i := 0; i < 64; i++ {
		for j := 0; j < 32; j++ {
			GFX[i][j] = 0
		}
	}

	for i := 0; i < len(FONT); i++ {
		MEMORY[i] = fmt.Sprintf("%02X", FONT[i])
	}

	PC = 0x200
	bufr := bufio.NewReader(f)
	bytes := make([]byte, 1)

	for {

		_, err := bufr.Read(bytes)
		if err != nil {
			break
		}
		MEMORY[PC] = fmt.Sprintf("%02X", bytes[0])
		PC = PC + 1

	}

	PC = 0x200

	for {

		ins := MEMORY[PC] + MEMORY[PC+1]

		_ii = _ii + 1
		fmt.Println("::: PC ", PC, " V: ", V, " _ii ", _ii, " ins ", ins, " SP ", SP, " stack ", STACK)
		if PAUSE {
			continue
		}
		switch ins[0] {
		case '0':
			INST_0(ins)

		case '1':
			INST_1(ins)

		case '2':
			INST_2(ins)

		case '3':
			INST_3(ins)

		case '4':
			INST_4(ins)

		case '5':
			INST_5(ins)

		case '6':
			INST_6(ins)

		case '7':
			INST_7(ins)

		case '8':
			INST_8(ins)

		case '9':
			INST_9(ins)

		case 'A':
			INST_A(ins)

		case 'B':
			INST_B(ins)

		case 'C':
			INST_C(ins)

		case 'D':
			INST_D(ins)

		case 'E':
			INST_E(ins)

		case 'F':
			INST_F(ins)

		default:
			fmt.Println("NOT DEFINED")
		}
		if true {
			INST_DRAW(PC)
		}
		time.Sleep(time.Millisecond * 60)

		PC += 2
	}

}
