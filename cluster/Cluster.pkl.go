// Code generated from Pkl module `mattsre.flyhouse.Cluster`. DO NOT EDIT.
package cluster

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type Cluster struct {
	// The name of the Clickhouse Cluster to be deployed
	Name string `pkl:"name"`

	// Whether this cluster should have a public IPv4 address or not
	Public bool `pkl:"public"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a Cluster
func LoadFromPath(ctx context.Context, path string) (ret *Cluster, err error) {
	evaluator, err := pkl.NewEvaluator(ctx, pkl.PreconfiguredOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := evaluator.Close()
		if err == nil {
			err = cerr
		}
	}()
	ret, err = Load(ctx, evaluator, pkl.FileSource(path))
	return ret, err
}

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a Cluster
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*Cluster, error) {
	var ret Cluster
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
