package utils

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Exec(ctx context.Context, envs map[string]string, name string, args ...[]string) ([]byte, error) {
	cmd := exec.CommandContext(ctx, name, args[0]...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if envs != nil {
		for k, v := range envs {
			cmd.Env = append(os.Environ(), strings.Join([]string{k, v}, "="))
		}
	}
	if err := cmd.Run(); err != nil {
		if ctx.Err() != nil {
			return nil, fmt.Errorf("%w", ctx.Err())
		}
		return nil, fmt.Errorf("%v", strings.TrimSpace(stderr.String()))
	}
	return stdout.Bytes(), nil
}
