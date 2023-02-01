package ascii

import (
	"ascii-art-web/cmd/ascii-art/funcs"
	"fmt"
)

func Ascii(text string, font string) (string, error) {
	args, err := funcs.ReadArgs(text)
	if err != nil {
		return "", err
	}
	art, err := funcs.Readfile(font)
	if err != nil {
		return "", err
	}

	sign, err := funcs.Arrayart(art)
	if err != nil {
		return "", err
	}
	result := funcs.Compare(sign, args)
	fmt.Println(result)
	return result, nil
}
