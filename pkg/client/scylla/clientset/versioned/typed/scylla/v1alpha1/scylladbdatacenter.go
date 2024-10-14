// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/scylladb/scylla-operator/pkg/api/scylla/v1alpha1"
	scheme "github.com/scylladb/scylla-operator/pkg/client/scylla/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ScyllaDBDatacentersGetter has a method to return a ScyllaDBDatacenterInterface.
// A group's client should implement this interface.
type ScyllaDBDatacentersGetter interface {
	ScyllaDBDatacenters(namespace string) ScyllaDBDatacenterInterface
}

// ScyllaDBDatacenterInterface has methods to work with ScyllaDBDatacenter resources.
type ScyllaDBDatacenterInterface interface {
	Create(ctx context.Context, scyllaDBDatacenter *v1alpha1.ScyllaDBDatacenter, opts v1.CreateOptions) (*v1alpha1.ScyllaDBDatacenter, error)
	Update(ctx context.Context, scyllaDBDatacenter *v1alpha1.ScyllaDBDatacenter, opts v1.UpdateOptions) (*v1alpha1.ScyllaDBDatacenter, error)
	UpdateStatus(ctx context.Context, scyllaDBDatacenter *v1alpha1.ScyllaDBDatacenter, opts v1.UpdateOptions) (*v1alpha1.ScyllaDBDatacenter, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ScyllaDBDatacenter, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ScyllaDBDatacenterList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ScyllaDBDatacenter, err error)
	ScyllaDBDatacenterExpansion
}

// scyllaDBDatacenters implements ScyllaDBDatacenterInterface
type scyllaDBDatacenters struct {
	client rest.Interface
	ns     string
}

// newScyllaDBDatacenters returns a ScyllaDBDatacenters
func newScyllaDBDatacenters(c *ScyllaV1alpha1Client, namespace string) *scyllaDBDatacenters {
	return &scyllaDBDatacenters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the scyllaDBDatacenter, and returns the corresponding scyllaDBDatacenter object, and an error if there is any.
func (c *scyllaDBDatacenters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ScyllaDBDatacenter, err error) {
	result = &v1alpha1.ScyllaDBDatacenter{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("scylladbdatacenters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ScyllaDBDatacenters that match those selectors.
func (c *scyllaDBDatacenters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ScyllaDBDatacenterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ScyllaDBDatacenterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("scylladbdatacenters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested scyllaDBDatacenters.
func (c *scyllaDBDatacenters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("scylladbdatacenters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a scyllaDBDatacenter and creates it.  Returns the server's representation of the scyllaDBDatacenter, and an error, if there is any.
func (c *scyllaDBDatacenters) Create(ctx context.Context, scyllaDBDatacenter *v1alpha1.ScyllaDBDatacenter, opts v1.CreateOptions) (result *v1alpha1.ScyllaDBDatacenter, err error) {
	result = &v1alpha1.ScyllaDBDatacenter{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("scylladbdatacenters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(scyllaDBDatacenter).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a scyllaDBDatacenter and updates it. Returns the server's representation of the scyllaDBDatacenter, and an error, if there is any.
func (c *scyllaDBDatacenters) Update(ctx context.Context, scyllaDBDatacenter *v1alpha1.ScyllaDBDatacenter, opts v1.UpdateOptions) (result *v1alpha1.ScyllaDBDatacenter, err error) {
	result = &v1alpha1.ScyllaDBDatacenter{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("scylladbdatacenters").
		Name(scyllaDBDatacenter.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(scyllaDBDatacenter).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *scyllaDBDatacenters) UpdateStatus(ctx context.Context, scyllaDBDatacenter *v1alpha1.ScyllaDBDatacenter, opts v1.UpdateOptions) (result *v1alpha1.ScyllaDBDatacenter, err error) {
	result = &v1alpha1.ScyllaDBDatacenter{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("scylladbdatacenters").
		Name(scyllaDBDatacenter.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(scyllaDBDatacenter).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the scyllaDBDatacenter and deletes it. Returns an error if one occurs.
func (c *scyllaDBDatacenters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("scylladbdatacenters").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *scyllaDBDatacenters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("scylladbdatacenters").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched scyllaDBDatacenter.
func (c *scyllaDBDatacenters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ScyllaDBDatacenter, err error) {
	result = &v1alpha1.ScyllaDBDatacenter{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("scylladbdatacenters").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}