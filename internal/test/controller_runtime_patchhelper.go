package test

import (
	"context"

	"sigs.k8s.io/cluster-api/util/patch"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type PatchHelperMock struct {
	PatchErrors []error
}

func (m PatchHelperMock) Patch(ctx context.Context, obj client.Object, opts ...patch.Option) error {
	var err error
	var errs []error
	if m.PatchErrors != nil {
		err, errs = m.PatchErrors[0], m.PatchErrors[1:]
	}
	m = PatchHelperMock{
		PatchErrors: errs,
	}
	return err
}
