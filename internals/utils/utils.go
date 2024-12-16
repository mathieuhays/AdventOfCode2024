package utils

import (
	"errors"
	"io"
	"os"
	"strconv"
)

func Validate() error {
	if len(os.Args) != 3 {
		return errors.New("invalid number of args")
	}

	part, err := GetPart()
	if err != nil {
		return err
	}

	if part < 1 || part > 2 {
		return errors.New("invalid part number")
	}

	return nil
}

func GetPart() (int, error) {
	return strconv.Atoi(os.Args[1])
}

func GetInput() (io.Reader, error) {
	return os.Open(os.Args[2])
}
