package main

import (
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	alias := cmd[0]
	var args []string
	if len(cmd) > 1 {
		args = cmd[1:]
	}
	changeEnvs(env)
	command := exec.Command(alias, args...)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Run()
	return command.ProcessState.ExitCode()
}

func changeEnvs(envs Environment) {
	for name, value := range envs {
		if value.NeedRemove {
			os.Unsetenv(name)
			delete(envs, name)
			continue
		}
		os.Setenv(name, value.Value)
	}
}
