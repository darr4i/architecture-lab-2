package lab2

import "io"

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	data := make([]byte, 0)
	buffer := make([]byte, 8)
	for {
		n, err := ch.Input.Read(buffer)
		data = append(data, buffer[:n]...)
		if err == io.EOF {
			break
		}
	}
	str := string(data)
	res, err := CalculatePrefix(str)
	if err != nil {
		return err
	}
	ch.Output.Write([]byte(res))
	return nil
}
