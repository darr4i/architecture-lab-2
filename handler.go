package lab2

import (
	"bytes"
	"io"
	"strconv"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	bufRead := new(bytes.Buffer)
	_, err := io.Copy(bufRead, ch.Input)
	if err != nil {
		return err
	}

	input := bufRead.String()

	res, err := CalculatePrefix(input)
	if err != nil {
		return err
	}
	ch.Output.Write([]byte(strconv.FormatInt(int64(res), 10)))
	return nil
}
