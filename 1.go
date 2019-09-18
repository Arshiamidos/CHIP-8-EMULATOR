package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
)

// REGISTER
var INSTRUCTION_SIZE = 8
var V = make([]int, 16)
var (
	PC          = 0x0
	I           = 0x200
	BK          = 0x0
	FONT =[]int{ 
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
		0xF0, 0x80, 0xF0, 0x80, 0x80  // F
  };
	MEMORY      = make([]string, 4096)
	DELAY_TIMER = 0x0
	SOUND_TIMER = 0x0
)

func OnByte(X1 byte) int {
	n, e := strconv.ParseInt(fmt.Sprintf("%s", string(X1)), 16, 8)
	if e != nil {
		fmt.Println(">")
		panic(e)
	}
	return int(n) & 0xFF
}
func ToByte(X1, X2 byte) int {
	n, e := strconv.ParseInt(fmt.Sprintf("%s%s", string(X1), string(X2)), 16, 8)
	if e != nil {
		fmt.Println(">>")
		panic(e)
	}
	return int(n) & 0xFF
}

func ThByte(X1, X2, X3 byte) int {
	n, e := strconv.ParseInt(fmt.Sprintf("%s%s%s", string(X1), string(X2), string(X3)), 16, 12)
	if e != nil {
		fmt.Println(">>>")
		panic(e)
	}
	return int(n) & 0xFFF
}
func DrawClr() {

}

func INST_0(instruction string) {

	switch instruction {
	case "00E0":
		fmt.Println("0x0000 ", "00E0")
		DrawClr()
	case "00EE":
		fmt.Println("0x0000 ", "00EE")
		I = BK
	}
}
func INST_1(instruction string) {
	fmt.Println("0x1111 ", instruction[1:])
	I = ThByte(instruction[1], instruction[2], instruction[3])
}

func INST_2(instruction string) {
	fmt.Println("0x2222 ", instruction[1:], ThByte(instruction[1], instruction[2], instruction[3]))
	I = 512 + (ThByte(instruction[1], instruction[2], instruction[3])-512)/2
	BK = I
}

func INST_3(instruction string) {
	fmt.Println("0x3333 ", instruction[1:])
	if V[OnByte(instruction[1])] == ToByte(instruction[2], instruction[3]) {
		PC++
	}
}

func INST_4(instruction string) {
	fmt.Println("0x4444 ", instruction[1:])
	if V[OnByte(instruction[1])] != ToByte(instruction[2], instruction[3]) {
		PC++
	}
}

func INST_5(instruction string) {
	fmt.Println("0x5555 ", instruction[1:])
	if V[OnByte(instruction[1])] == V[OnByte(instruction[2])] {
		PC++
	}
}

func INST_6(instruction string) {
	fmt.Println("0x6666 ", instruction[1:], instruction[1], ToByte(instruction[2], instruction[3]))
	V[OnByte(instruction[1])] = ToByte(instruction[2], instruction[3])
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
	V[OnByte(instruction[1])] = (V[OnByte(instruction[1])] + V[OnByte(instruction[2])]) & 0xFF
}

func INST_8_5(instruction string) {
	fmt.Println("0x8855 ", instruction[1:])
	//VF
	V[OnByte(instruction[1])] = (V[OnByte(instruction[1])] - V[OnByte(instruction[2])]) & 0xFF

}

func INST_8_6(instruction string) {
	fmt.Println("0x8866 ", instruction[1:])
	V[15] = V[OnByte(instruction[1])]
	V[OnByte(instruction[1])] = (V[OnByte(instruction[1])] >> 1) & 0xFF
}

func INST_8_7(instruction string) {
	fmt.Println("0x8877 ", instruction[1:])
	V[OnByte(instruction[1])] = (V[OnByte(instruction[2])] - V[OnByte(instruction[1])]) & 0xFF

}

func INST_8_E(instruction string) {
	fmt.Println("0x88EE ", instruction[1:])
	V[15] = V[OnByte(instruction[1])]
	V[OnByte(instruction[1])] = (V[OnByte(instruction[1])] << 1) & 0xFF
}

func INST_8(instruction string) {
	switch instruction[3] {
	case 0:
		INST_8_0(instruction)
	case 1:
		INST_8_1(instruction)
	case 2:
		INST_8_2(instruction)
	case 3:
		INST_8_3(instruction)
	case 4:
		INST_8_4(instruction)
	case 5:
		INST_8_5(instruction)
	case '6':
		INST_8_6(instruction)
	case 7:
		INST_8_7(instruction)
	case 'E':
		INST_8_E(instruction)

	}
}

func INST_9(instruction string) {
	if V[OnByte(instruction[1])] != V[OnByte(instruction[2])] {
		PC++
	}
}

func INST_A(instruction string) {
	I = 512 + (ThByte(instruction[1], instruction[2], instruction[3])-512)/2

	fmt.Println(">>", string(instruction[1]), string(instruction[2]), string(instruction[3]), ThByte(instruction[1], instruction[2], instruction[3]), I)
}

func INST_B(instruction string) {
	PC = (V[0] + V[OnByte(instruction[1])] + V[OnByte(instruction[2])] + V[OnByte(instruction[3])]) & 0xFF
}

func INST_C(instruction string) {
	V[OnByte(instruction[1])] = (rand.Intn(255) * (V[OnByte(instruction[1])] + V[OnByte(instruction[2])])) & 0xFF
}

func Draw(x, y, z int) {
	fmt.Println("Drawed", x, y, z)
}

func INST_D(instruction string) {
	Draw(V[OnByte(instruction[1])], V[OnByte(instruction[2])], V[OnByte(instruction[3])])
}

func INST_E(instruction string) {
	if instruction[3] == 0xE {
		if V[OnByte(instruction[1])] == 0 { //PRESSED

		}
		PC++
	} else if instruction[3] == 0x1 {
		if V[OnByte(instruction[1])] != 0 { //NOT PRESSED

		}
	}
}

func getKey() {

}
func INST_F(instruction string) {

	if instruction[3] == 0x7 {
		V[OnByte(instruction[1])] = DELAY_TIMER
	} else if instruction[3] == 0xA {
		getKey()
	} else if instruction[3] == 0x5 && instruction[2] == 0x1 {
		DELAY_TIMER = V[OnByte(instruction[1])]
	} else if instruction[3] == 0x8 && instruction[2] == 0x1 {
		SOUND_TIMER = V[OnByte(instruction[1])]
	} else if instruction[3] == 0xE {
		I = (I + V[OnByte(instruction[1])]) & 0xFFF
	} else if instruction[3] == 0x9 {
		//I = SPRITE_ADDR[V[OnByte(instruction[1])]]
	} else if instruction[3] == 0x3 {
		num := float64(V[OnByte(instruction[1])])
		MEMORY[I+0] = fmt.Sprintf("%02X", int(num-10.0*math.Ceil(num/10.0)))
		MEMORY[I+1] = fmt.Sprintf("%02X", int(math.Ceil((100*math.Ceil(num/100.0)-num)/10.0)))
		MEMORY[I+2] = fmt.Sprintf("%02X", int(math.Ceil(num/100.0)))

	} else if instruction[3] == 0x8 {
		SOUND_TIMER = V[OnByte(instruction[1])]
	} else if instruction[2] == 0x5 && instruction[3] == 0x5 {
		for i := 0; i < 16; i++ {
			MEMORY[I+i] = fmt.Sprintf("%02X", V[i])
		}
	} else if instruction[2] == 0x6 && instruction[3] == 0x5 {
		for i := 0; i < 16; i++ {
			V[i] = I + i
		}
	}
}

func main() {
	f, err := os.Open("GUESS")
	bytes := make([]byte, 2)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	for {
		_, err := f.Read(bytes)
		MEMORY[I] = fmt.Sprintf("%02X", bytes)
		I = I + 1
		if err != nil {
			break
		}
	}
	I = 0x200
	fmt.Println("::: ", MEMORY)
	os.Exit(1)
	for {
		ins := MEMORY[I]
		fmt.Println("::: ", "I:", I, " V: ", V, " ", ins, fmt.Sprintf("%X", ins[0]))

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

		PC += 2
	}

}
