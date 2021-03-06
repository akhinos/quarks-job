/*

Don't alter this file, it was generated.

*/
// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "code.cloudfoundry.org/quarks-job/pkg/kube/apis/quarksjob/v1alpha1"
	scheme "code.cloudfoundry.org/quarks-job/pkg/kube/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// QuarksJobsGetter has a method to return a QuarksJobInterface.
// A group's client should implement this interface.
type QuarksJobsGetter interface {
	QuarksJobs(namespace string) QuarksJobInterface
}

// QuarksJobInterface has methods to work with QuarksJob resources.
type QuarksJobInterface interface {
	Create(*v1alpha1.QuarksJob) (*v1alpha1.QuarksJob, error)
	Update(*v1alpha1.QuarksJob) (*v1alpha1.QuarksJob, error)
	UpdateStatus(*v1alpha1.QuarksJob) (*v1alpha1.QuarksJob, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.QuarksJob, error)
	List(opts v1.ListOptions) (*v1alpha1.QuarksJobList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.QuarksJob, err error)
	QuarksJobExpansion
}

// quarksJobs implements QuarksJobInterface
type quarksJobs struct {
	client rest.Interface
	ns     string
}

// newQuarksJobs returns a QuarksJobs
func newQuarksJobs(c *QuarksjobV1alpha1Client, namespace string) *quarksJobs {
	return &quarksJobs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the quarksJob, and returns the corresponding quarksJob object, and an error if there is any.
func (c *quarksJobs) Get(name string, options v1.GetOptions) (result *v1alpha1.QuarksJob, err error) {
	result = &v1alpha1.QuarksJob{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("quarksjobs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of QuarksJobs that match those selectors.
func (c *quarksJobs) List(opts v1.ListOptions) (result *v1alpha1.QuarksJobList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.QuarksJobList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("quarksjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested quarksJobs.
func (c *quarksJobs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("quarksjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a quarksJob and creates it.  Returns the server's representation of the quarksJob, and an error, if there is any.
func (c *quarksJobs) Create(quarksJob *v1alpha1.QuarksJob) (result *v1alpha1.QuarksJob, err error) {
	result = &v1alpha1.QuarksJob{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("quarksjobs").
		Body(quarksJob).
		Do().
		Into(result)
	return
}

// Update takes the representation of a quarksJob and updates it. Returns the server's representation of the quarksJob, and an error, if there is any.
func (c *quarksJobs) Update(quarksJob *v1alpha1.QuarksJob) (result *v1alpha1.QuarksJob, err error) {
	result = &v1alpha1.QuarksJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("quarksjobs").
		Name(quarksJob.Name).
		Body(quarksJob).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *quarksJobs) UpdateStatus(quarksJob *v1alpha1.QuarksJob) (result *v1alpha1.QuarksJob, err error) {
	result = &v1alpha1.QuarksJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("quarksjobs").
		Name(quarksJob.Name).
		SubResource("status").
		Body(quarksJob).
		Do().
		Into(result)
	return
}

// Delete takes name of the quarksJob and deletes it. Returns an error if one occurs.
func (c *quarksJobs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("quarksjobs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *quarksJobs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("quarksjobs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched quarksJob.
func (c *quarksJobs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.QuarksJob, err error) {
	result = &v1alpha1.QuarksJob{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("quarksjobs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
