package container

import (
	"context"
	"dagger.io/dagger"
	"fmt"
	"github.com/go-courier/logr"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/octohelm/piper/pkg/cueflow"
	piperdagger "github.com/octohelm/piper/pkg/dagger"
	"github.com/octohelm/piper/pkg/engine/task"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"strings"
	"sync"
)

func init() {
	cueflow.RegisterTask(task.Factory, &Push{})
}

// Push image to registry
type Push struct {
	task.Task
	// image tag
	Dest string `json:"dest"`
	// images for push
	// [Platform]: _
	Images map[string]Container `json:"images"`

	// registry auth
	Auth *Auth `json:"auth,omitempty"`

	// image pushed result
	Result string `json:"-" output:"result"`
}

func (x *Push) ResultValue() any {
	return map[string]any{
		"result": x.Result,
	}
}

func (x *Push) Do(ctx context.Context) error {
	if len(x.Images) == 0 {
		return errors.New("Push at least one image")
	}

	eg := &errgroup.Group{}
	published := &sync.Map{}

	for platform, container := range x.Images {
		eg.Go(func() error {
			return container.Select(ctx).Do(ctx, func(ctx context.Context, c *piperdagger.Client) error {
				cc := container.Container(c)

				p, err := v1.ParsePlatform(platform)
				if err != nil {
					return errors.Wrapf(err, "parse platform failed: %s", p)
				}

				// push without tag
				image := strings.Split(x.Dest, ":")[0]

				logr.FromContext(ctx).WithValues("platform", p).Info(fmt.Sprintf("publishing %s", image))

				cc = RegistryAuthStoreContext.From(ctx).ApplyTo(ctx, c, cc, image, x.Auth)

				tag, err := cc.Publish(ctx, image)
				if err != nil {
					return fmt.Errorf("published image failed")
				}

				published.Store(p, tag)
				return nil
			})
		})
	}

	if err := eg.Wait(); err != nil {
		return err
	}

	return piperdagger.Select(ctx, piperdagger.Scope{}).Do(ctx, func(ctx context.Context, c *piperdagger.Client) error {
		opt := dagger.ContainerPublishOpts{}

		for k, v := range published.Range {
			platform := k.(*v1.Platform)
			image := v.(string)

			cc := c.Container(dagger.ContainerOpts{Platform: dagger.Platform(platform.String())})
			cc = RegistryAuthStoreContext.From(ctx).ApplyTo(ctx, c, cc, image, x.Auth)
			opt.PlatformVariants = append(opt.PlatformVariants, cc.From(image))
		}

		result, err := c.Container().Publish(ctx, x.Dest, opt)
		if err != nil {
			return fmt.Errorf("published manifest list failed")
		}

		x.Result = result

		return nil
	})
}