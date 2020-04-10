// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"time"

	v1beta1 "gitlab.alibaba-inc.com/cos/application-center-manager/pkg/apis/alibabacloud/v1beta1"
	scheme "gitlab.alibaba-inc.com/cos/application-center-manager/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CanariesGetter has a method to return a CanaryInterface.
// A group's client should implement this interface.
type CanariesGetter interface {
	Canaries(namespace string) CanaryInterface
}

// CanaryInterface has methods to work with Canary resources.
type CanaryInterface interface {
	Create(*v1beta1.Canary) (*v1beta1.Canary, error)
	Update(*v1beta1.Canary) (*v1beta1.Canary, error)
	UpdateStatus(*v1beta1.Canary) (*v1beta1.Canary, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.Canary, error)
	List(opts v1.ListOptions) (*v1beta1.CanaryList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Canary, err error)
	CanaryExpansion
}

// canaries implements CanaryInterface
type canaries struct {
	client rest.Interface
	ns     string
}

// newCanaries returns a Canaries
func newCanaries(c *AppV1beta1Client, namespace string) *canaries {
	return &canaries{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the canary, and returns the corresponding canary object, and an error if there is any.
func (c *canaries) Get(name string, options v1.GetOptions) (result *v1beta1.Canary, err error) {
	result = &v1beta1.Canary{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("canaries").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Canaries that match those selectors.
func (c *canaries) List(opts v1.ListOptions) (result *v1beta1.CanaryList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.CanaryList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("canaries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested canaries.
func (c *canaries) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("canaries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a canary and creates it.  Returns the server's representation of the canary, and an error, if there is any.
func (c *canaries) Create(canary *v1beta1.Canary) (result *v1beta1.Canary, err error) {
	result = &v1beta1.Canary{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("canaries").
		Body(canary).
		Do().
		Into(result)
	return
}

// Update takes the representation of a canary and updates it. Returns the server's representation of the canary, and an error, if there is any.
func (c *canaries) Update(canary *v1beta1.Canary) (result *v1beta1.Canary, err error) {
	result = &v1beta1.Canary{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("canaries").
		Name(canary.Name).
		Body(canary).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *canaries) UpdateStatus(canary *v1beta1.Canary) (result *v1beta1.Canary, err error) {
	result = &v1beta1.Canary{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("canaries").
		Name(canary.Name).
		SubResource("status").
		Body(canary).
		Do().
		Into(result)
	return
}

// Delete takes name of the canary and deletes it. Returns an error if one occurs.
func (c *canaries) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("canaries").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *canaries) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("canaries").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched canary.
func (c *canaries) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Canary, err error) {
	result = &v1beta1.Canary{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("canaries").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
