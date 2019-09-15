package main

import (
	"fmt"
	"math/rand"
	"os"
)

// REGISTER
var INSTRUCTION_SIZE = 8
var V = make([]byte, 16) //variables
var (
	NNN         = 0x000
	NN          = 0x00
	N           = 0x0
	PC          = 0           //program counter
	I           = byte(0x000) //address
	ADDRESSES   = make([][3]byte, 4096)
	DELAY_TIMER = byte(0x0)
	SOUND_TIMER = byte(0x0)
)

func INST_0(instruction string) {

	fmt.Println("0x0000")

	switch instruction {
	case "00E0":
		fmt.Print("00E0")
	case "00EE":
		fmt.Print("00EE")
	}
}
func INST_1(instruction string) {
	fmt.Println("0x1111 ", instruction[1:])
	//I = instruction[1:]
}

func INST_2(instruction string) {
	fmt.Println("0x2222 ", instruction[1:])
	//I = instruction[1:]

}

func INST_3(instruction string) {
	fmt.Println("0x3333 ", instruction[1:])
	if V[int(instruction[1])] == fmt.Sprintf("%s%s", instruction[2], instruction[3]) {
		PC++
	}
}

func INST_4(instruction string) {
	fmt.Println("0x4444 ", instruction[1:])
	if V[int(instruction[1])] != fmt.Sprintf("%s%s", instruction[2], instruction[3]) {
		PC++
	}
}

func INST_5(instruction string) {
	fmt.Println("0x5555 ", instruction[1:])
	if V[int(instruction[1])] == V[int(instruction[2])] {
		PC++
	}
}

func INST_6(instruction string) {
	fmt.Println("0x6666 ", instruction[1:])
	V[int(instruction[1])] = fmt.Sprintf("%s%s", instruction[2], instruction[3])
}

func INST_7(instruction string) {
	fmt.Println("0x7777 ", instruction[1:])
	/* V[int(instruction[1])] = V[int(instruction[1])] + strconv.FormatUint(
	fmt.Sprintf("%s%s", instruction[2], instruction[3]),
	16) */

}
func INST_8_0(instruction string) {
	fmt.Println("0x8800 ", instruction[1:])
	V[int(instruction[1])] = V[int(instruction[2])]
}
func INST_8_1(instruction string) {
	V[int(instruction[1])] = V[int(instruction[1])] | V[int(instruction[2])]
	fmt.Println("0x8811 ", instruction[1:])
}

func INST_8_2(instruction string) {
	fmt.Println("0x8822 ", instruction[1:])
	V[int(instruction[1])] = V[int(instruction[1])] & V[int(instruction[2])]

}

func INST_8_3(instruction string) {
	fmt.Println("0x8833 ", instruction[1:])
	V[int(instruction[1])] = V[int(instruction[1])] ^ V[int(instruction[2])]

}

func INST_8_4(instruction string) {

	fmt.Println("0x8844 ", instruction[1:])
	V[int(instruction[1])] = V[int(instruction[1])] + V[int(instruction[2])]
}

func INST_8_5(instruction string) {
	fmt.Println("0x8855 ", instruction[1:])
	//VF
	V[int(instruction[1])] = V[int(instruction[1])] - V[int(instruction[2])]

}

func INST_8_6(instruction string) {
	fmt.Println("0x8866 ", instruction[1:])
	V[15] = V[int(instruction[1])]
	V[int(instruction[1])] = V[int(instruction[1])] >> 1
}

func INST_8_7(instruction string) {
	fmt.Println("0x8877 ", instruction[1:])
	V[int(instruction[1])] = V[int(instruction[2])] - V[int(instruction[1])]

}

func INST_8_E(instruction string) {
	fmt.Println("0x88EE ", instruction[1:])
	V[15] = V[int(instruction[1])]
	V[int(instruction[1])] = V[int(instruction[1])] << 1
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
	if V[int(instruction[1])] != V[int(instruction[2])] {
		PC++
	}
}
func INST_A(instruction string) {
	I = int(V[int(instruction[1])] + V[int(instruction[2])] + V[int(instruction[3])])
}
func INST_B(instruction string) {
	PC = int(V[0]) + int(V[int(instruction[1])]+V[int(instruction[2])]+V[int(instruction[3])])
}
func INST_C(instruction string) {
	V[int(instruction[1])] = rand.Rand(0, 255) * (V[int(instruction[1])] + V[int(instruction[2])])
}

func Draw(x, y, z byte) {
	fmt.Println("z", z)
}
func INST_D(instruction string) {
	Draw(V[int(instruction[1])], V[int(instruction[2])], V[int(instruction[3])])
}

func INST_E(instruction string) {
	if instruction[3] == 0xE {
		if V[instruction[1]] == 0 { //PRESSED

		}
		PC++
	} else if instruction[3] == 0x1 {
		if V[instruction[1]] != 0 { //NOT PRESSED

		}
	}
}
func getKey() {

}
func INST_F(instruction string) {

	if instruction[3] == 0x7 {
		V[instruction[1]] = DELAY_TIMER
	} else if instruction[3] == 0xA {
		getKey()
	} else if instruction[3] == 0x5 && instruction[2] == 0x1 {
		DELAY_TIMER = V[instruction[1]]
	} else if instruction[3] == 0x8 && instruction[2] == 0x1 {
		SOUND_TIMER = V[instruction[1]]
	} else if instruction[3] == 0x8 && instruction[2] == 0x1 {
		SOUND_TIMER = V[instruction[1]]
	} else if instruction[3] == 0xE {
		I = I + V[instruction[1]]
	} else if instruction[3] == 0x9 {
		//I = SPRITE_ADDR[V[instruction[1]]]
	} else if instruction[3] == 0x3 {
		SOUND_TIMER = V[instruction[1]]
	} else if instruction[3] == 0x8 {
		SOUND_TIMER = V[instruction[1]]
	} else if instruction[2] == 0x5 && instruction[3] == 0x5 {
		for i := 0; i < 16; i++ {
			IPC[I+i] = V[i]
		}
	} else if instruction[2] == 0x6 && instruction[3] == 0x5 {
		for i := 0; i < 16; i++ {
			V[i] = I + byte(i)
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
		fmt.Println(fmt.Sprintf("%02X", bytes))
		instruction := fmt.Sprintf("%02X", bytes)

		switch instruction[0] {
		case '0':
			INST_0(instruction)
		case '1':
			INST_1(instruction)
		case '2':
			INST_2(instruction)

		case '3':
			INST_3(instruction)

		case '4':
			INST_4(instruction)

		case '5':
			INST_5(instruction)

		case '6':
			INST_6(instruction)

		case '7':
			INST_7(instruction)

		case '8':
			INST_8(instruction)
		case '9':
			INST_9(instruction)
		case 'A':
			INST_A(instruction)
		case 'B':
			INST_B(instruction)
		case 'C':
			INST_C(instruction)
		case 'D':
			INST_D(instruction)
		case 'E':
			INST_E(instruction)
		case 'F':
			INST_F(instruction)

		default:
			fmt.Println("NOT DEFINED")
		}

		if err != nil {
			break
		}
	}

}
