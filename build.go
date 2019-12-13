package samplefile

import (
	"context"
	"samplefile/config"

	"github.com/moby/buildkit/client/llb"
	"github.com/moby/buildkit/frontend/gateway/client"
	"github.com/pkg/errors"
)

func Build(ctx context.Context, c client.Client) (*client.Result, error) {
	opts := c.BuildOpts().Opts
	filename := opts["filename"]
	name := "load Samplefile"
	if filename == "" {
		filename = "Samplefile.yaml"
	}
	if filename != "Samplefile" {
		name += " from " + filename
	}
	src := llb.Local("dockerfile", llb.IncludePatterns([]string{filename}),
		llb.SessionID(c.BuildOpts().SessionID),
		llb.SharedKeyHint("Samplefile.yaml"),
		llb.WithCustomName("[internal] "+name),
	)
	def, err := src.Marshal()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to marshal local source")
	}
	res, err := c.Solve(ctx, client.SolveRequest{
		Definition: def.ToPB(),
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to resolve dockerfile")
	}
	ref, err := res.SingleRef()
	if err != nil {
		return nil, err
	}

	var dtsamplefile []byte
	dtsamplefile, _ = ref.ReadFile(ctx, client.ReadRequest{
		Filename: filename,
	})
	cfg, err := config.NewConfigFromBytes(dtsamplefile)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read dockerfile")
	}
	st := llb.Image(cfg.Image)
	def, err = st.Marshal()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to marshal local source")
	}
	res, err = c.Solve(ctx, client.SolveRequest{
		Definition: def.ToPB(),
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to resolve dockerfile")
	}
	ref, err = res.SingleRef()
	res, _ = c.Solve(ctx, client.SolveRequest{
		Definition: def.ToPB(),
	})
	return res, nil
}
