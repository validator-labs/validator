package test

import (
	"context"

	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type SubResourceMock struct {
	UpdateErrors []error
}

func (m SubResourceMock) Create(ctx context.Context, obj client.Object, subResource client.Object, opts ...client.SubResourceCreateOption) error {
	return nil
}
func (m SubResourceMock) Update(ctx context.Context, obj client.Object, opts ...client.SubResourceUpdateOption) error {
	var err error
	var errs []error
	if m.UpdateErrors != nil {
		err, errs = m.UpdateErrors[0], m.UpdateErrors[1:]
	}
	m = SubResourceMock{
		UpdateErrors: errs,
	}
	return err
}
func (m SubResourceMock) Patch(ctx context.Context, obj client.Object, patch client.Patch, opts ...client.SubResourcePatchOption) error {
	return nil
}

type ClientMock struct {
	CreateErrors []error
	GetErrors    []error
	UpdateErrors []error
	SubResourceMock
}

func (m ClientMock) Get(ctx context.Context, key client.ObjectKey, obj client.Object, opts ...client.GetOption) error {
	var err error
	var errs []error
	if m.GetErrors != nil {
		err, errs = m.GetErrors[0], m.GetErrors[1:]
	}
	m = ClientMock{
		CreateErrors: m.CreateErrors,
		GetErrors:    errs,
		UpdateErrors: m.UpdateErrors,
	}
	return err
}
func (m ClientMock) List(ctx context.Context, list client.ObjectList, opts ...client.ListOption) error {
	return nil
}
func (m ClientMock) Create(ctx context.Context, obj client.Object, opts ...client.CreateOption) error {
	var err error
	var errs []error
	if m.CreateErrors != nil {
		err, errs = m.CreateErrors[0], m.CreateErrors[1:]
	}
	m = ClientMock{
		CreateErrors: errs,
		GetErrors:    m.GetErrors,
		UpdateErrors: m.UpdateErrors,
	}
	return err
}
func (m ClientMock) Delete(ctx context.Context, obj client.Object, opts ...client.DeleteOption) error {
	return nil
}
func (m ClientMock) Update(ctx context.Context, obj client.Object, opts ...client.UpdateOption) error {
	var err error
	var errs []error
	if m.UpdateErrors != nil {
		err, errs = m.UpdateErrors[0], m.UpdateErrors[1:]
	}
	m = ClientMock{
		CreateErrors: m.CreateErrors,
		GetErrors:    m.GetErrors,
		UpdateErrors: errs,
	}
	return err
}
func (m ClientMock) Patch(ctx context.Context, obj client.Object, patch client.Patch, opts ...client.PatchOption) error {
	return nil
}
func (m ClientMock) DeleteAllOf(ctx context.Context, obj client.Object, opts ...client.DeleteAllOfOption) error {
	return nil
}
func (m ClientMock) Status() client.SubResourceWriter                        { return m.SubResourceMock }
func (m ClientMock) SubResource(subResource string) client.SubResourceClient { return nil }
func (m ClientMock) Scheme() *runtime.Scheme                                 { return nil }
func (m ClientMock) RESTMapper() meta.RESTMapper                             { return nil }
func (m ClientMock) GroupVersionKindFor(obj runtime.Object) (schema.GroupVersionKind, error) {
	return schema.GroupVersionKind{}, nil
}
func (m ClientMock) IsObjectNamespaced(obj runtime.Object) (bool, error) { return false, nil }
