package main1

import (
	"fmt"
	"strconv"
)

func main() {
	//v := 0xFFFFF
	/* v := []byte("abcdZz")
	fmt.Println(float64(14))
	i, _ := strconv.ParseInt("F", 16, 8)
	fmt.Println(v, i, fmt.Sprintf("%02X", 240))
	fmt.Println(0x0, 0x00, 0x000)
	fmt.Println(0x0+1, 0x00+1, 0x000+1)
	fmt.Println(0xF+1, (0xFFE+1)/3, 0xFFFFFF+1)

	f, err := os.Open("GUESS")
	bytes := make([]byte, 2)

	//fmt.Println(0x0 + 1)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	for {
		_, err := f.Read(bytes)
		fmt.Println(fmt.Sprintf("%02X", bytes))
		fmt.Sprintf("%02X", bytes)

		if err != nil {
			break
		}
	} */
	fmt.Println(">>>" + "<<<<")
	n, _ := strconv.ParseInt("FF", 16, 8)
	fmt.Println(n + 1)
}
