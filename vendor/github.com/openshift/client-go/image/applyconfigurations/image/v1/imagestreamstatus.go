// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// ImageStreamStatusApplyConfiguration represents an declarative configuration of the ImageStreamStatus type for use
// with apply.
type ImageStreamStatusApplyConfiguration struct {
	DockerImageRepository       *string                               `json:"dockerImageRepository,omitempty"`
	PublicDockerImageRepository *string                               `json:"publicDockerImageRepository,omitempty"`
	Tags                        []NamedTagEventListApplyConfiguration `json:"tags,omitempty"`
}

// ImageStreamStatusApplyConfiguration constructs an declarative configuration of the ImageStreamStatus type for use with
// apply.
func ImageStreamStatus() *ImageStreamStatusApplyConfiguration {
	return &ImageStreamStatusApplyConfiguration{}
}

// WithDockerImageRepository sets the DockerImageRepository field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DockerImageRepository field is set to the value of the last call.
func (b *ImageStreamStatusApplyConfiguration) WithDockerImageRepository(value string) *ImageStreamStatusApplyConfiguration {
	b.DockerImageRepository = &value
	return b
}

// WithPublicDockerImageRepository sets the PublicDockerImageRepository field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PublicDockerImageRepository field is set to the value of the last call.
func (b *ImageStreamStatusApplyConfiguration) WithPublicDockerImageRepository(value string) *ImageStreamStatusApplyConfiguration {
	b.PublicDockerImageRepository = &value
	return b
}

// WithTags adds the given value to the Tags field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Tags field.
func (b *ImageStreamStatusApplyConfiguration) WithTags(values ...*NamedTagEventListApplyConfiguration) *ImageStreamStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithTags")
		}
		b.Tags = append(b.Tags, *values[i])
	}
	return b
}
