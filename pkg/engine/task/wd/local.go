package wd

import (
	"context"

	"github.com/octohelm/piper/pkg/cueflow"
	"github.com/octohelm/piper/pkg/engine/task"
)

func init() {
	cueflow.RegisterTask(task.Factory, &Local{})
}

// Local
// create a local work dir
type Local struct {
	task.Task

	// related dir on the root of project
	Dir string `json:"dir" default:"."`

	// the local work dir
	WorkDir WorkDir `json:"-" output:"wd"`
}

func (local *Local) ResultValue() any {
	return local.WorkDir
}

func (local *Local) Do(ctx context.Context) error {
	cwd, err := task.ClientContext.From(ctx).SourceDir(ctx)
	if err != nil {
		return err
	}
	return local.WorkDir.Sync(ctx, cwd)
}
