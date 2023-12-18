package exec

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"strings"

	"github.com/go-courier/logr"
	"github.com/octohelm/piper/pkg/cueflow"
	"github.com/octohelm/x/ptr"
	"github.com/pkg/errors"

	"github.com/octohelm/piper/pkg/engine/task"
	"github.com/octohelm/piper/pkg/engine/task/client"
	taskwd "github.com/octohelm/piper/pkg/engine/task/wd"
	"github.com/octohelm/piper/pkg/wd"
)

func init() {
	cueflow.RegisterTask(task.Factory, &Run{})
}

// Run some cmd
type Run struct {
	task.Task

	taskwd.CurrentWorkDir

	// cmd for executing
	Command client.StringOrSlice `json:"cmd"`
	// env vars
	Env map[string]client.SecretOrString `json:"env,omitempty"`
	// executing user
	User string `json:"user,omitempty"`

	// other setting
	With RunOption `json:"with,omitempty"`

	// result
	RunResult `json:"-" output:"result"`
}

type RunOption struct {
	// when enabled
	// once executed failed, will break whole pipeline
	// otherwise, just set result
	Failfast bool `json:"failfast,omitempty" default:"true"`

	// when enabled
	// `result.stdout` should be with the string value
	// otherwise, just log stdout
	Stdout bool `json:"stdout,omitempty" default:"false"`

	// when enabled
	// `result.ok` will not set be false if empty stdout
	StdoutOmitempty bool `json:"stdoutOmitempty,omitempty" default:"false"`

	// when enabled
	// `result.stderr` should be with the string value
	// otherwise, just log stderr
	Stderr bool `json:"stderr,omitempty" default:"false"`
}

// result of the executing
type RunResult struct {
	cueflow.Result
	// exists when `with.stdout` enabled
	Stdout *string `json:"stdout,omitempty"`
	// exists when `with.stdout` enabled
	Stderr *string `json:"stderr,omitempty"`
}

func (r *RunResult) ResultValue() any {
	return r
}

func (r RunResult) Success() bool {
	return r.Ok
}

func (r *Run) Do(ctx context.Context) error {
	return r.Cwd.Do(ctx, func(ctx context.Context, cwd wd.WorkDir) error {
		cmd := strings.Join(r.Command, " ")

		env := map[string]string{}

		for k, v := range r.Env {
			if s := v.Secret; s != nil {
				if secret, ok := s.Value(ctx); ok {
					env[k] = secret.Value
				} else {
					return errors.Errorf("not found secret for %s", k)
				}
			} else {
				env[k] = v.Value
			}
		}

		stdout := bytes.NewBuffer(nil)
		stderr := bytes.NewBuffer(nil)

		opts := []wd.OptionFunc{
			wd.WithEnv(env),
			wd.WithUser(r.User),
		}

		if r.With.Stdout {
			opts = append(opts, wd.WithStdout(stdout))
		} else {
			w := (&forward{l: logr.FromContext(ctx)}).NewWriter()
			defer w.Close()
			opts = append(opts, wd.WithStdout(w))
		}

		if r.With.Stderr {
			opts = append(opts, wd.WithStderr(stderr))
		} else {
			w := (&forward{l: logr.FromContext(ctx), stderr: true}).NewWriter()
			defer w.Close()
			opts = append(opts, wd.WithStderr(w))
		}

		err := cwd.Exec(ctx, cmd, opts...)

		if err != nil {
			if r.With.Failfast {
				return err
			}
		} else {
			r.Result.Ok = true
		}

		if r.With.Stdout {
			r.Stdout = ptr.Ptr(stdout.String())

			if !r.With.StdoutOmitempty {
				r.Result.Ok = stdout.Len() > 0

				if !r.Result.Ok && r.With.Failfast {
					return errors.New("empty stdout")
				}
			}
		}

		if r.With.Stdout {
			r.Stderr = ptr.Ptr(stderr.String())
		}

		return nil
	})
}

type forward struct {
	l      logr.Logger
	stderr bool
}

func (f *forward) NewWriter() io.WriteCloser {
	r, w := io.Pipe()

	go func() {
		s := bufio.NewScanner(r)

		for s.Scan() {
			if f.stderr {
				f.l.Warn(errors.New(s.Text()))
			} else {
				f.l.Info(s.Text())
			}
		}
	}()

	return w
}
