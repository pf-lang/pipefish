package pf_os_2fexec

import (
	"github.com/google/shlex"
	"os/exec"
)

var PIPEFISH_FUNCTION_CONVERTER = map[string](func(t uint32, v any) any){}

var PIPEFISH_VALUE_CONVERTER = map[string]any{}

func GoGet(command string) any {
	bits, err1 := shlex.Split(command)
	if err1 != nil {
		return err1
	}
	c := bits[0]
	a := bits[1:]
	bytes, err2 := exec.Command(c, a...).Output()
	if err2 != nil {
		return err2
	}
	return string(bytes)
}

func GoPost(command string) any {
	bits, err1 := shlex.Split(command)
	if err1 != nil {
		return err1
	}
	c := bits[0]
	a := bits[1:]
	err2 := exec.Command(c, a...).Run()
	if err2 != nil {
		return err2
	}
	return struct{}{}
}
