// +build !ignore_autogenerated

// autogenerated by controller-gen object, do not modify manually

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterDocsTopic) DeepCopyInto(out *ClusterDocsTopic) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterDocsTopic.
func (in *ClusterDocsTopic) DeepCopy() *ClusterDocsTopic {
	if in == nil {
		return nil
	}
	out := new(ClusterDocsTopic)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterDocsTopic) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterDocsTopicList) DeepCopyInto(out *ClusterDocsTopicList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterDocsTopic, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterDocsTopicList.
func (in *ClusterDocsTopicList) DeepCopy() *ClusterDocsTopicList {
	if in == nil {
		return nil
	}
	out := new(ClusterDocsTopicList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterDocsTopicList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterDocsTopicSpec) DeepCopyInto(out *ClusterDocsTopicSpec) {
	*out = *in
	in.CommonDocsTopicSpec.DeepCopyInto(&out.CommonDocsTopicSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterDocsTopicSpec.
func (in *ClusterDocsTopicSpec) DeepCopy() *ClusterDocsTopicSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterDocsTopicSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterDocsTopicStatus) DeepCopyInto(out *ClusterDocsTopicStatus) {
	*out = *in
	in.CommonDocsTopicStatus.DeepCopyInto(&out.CommonDocsTopicStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterDocsTopicStatus.
func (in *ClusterDocsTopicStatus) DeepCopy() *ClusterDocsTopicStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterDocsTopicStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonDocsTopicSpec) DeepCopyInto(out *CommonDocsTopicSpec) {
	*out = *in
	if in.Sources != nil {
		in, out := &in.Sources, &out.Sources
		*out = make([]Source, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonDocsTopicSpec.
func (in *CommonDocsTopicSpec) DeepCopy() *CommonDocsTopicSpec {
	if in == nil {
		return nil
	}
	out := new(CommonDocsTopicSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonDocsTopicStatus) DeepCopyInto(out *CommonDocsTopicStatus) {
	*out = *in
	in.LastHeartbeatTime.DeepCopyInto(&out.LastHeartbeatTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonDocsTopicStatus.
func (in *CommonDocsTopicStatus) DeepCopy() *CommonDocsTopicStatus {
	if in == nil {
		return nil
	}
	out := new(CommonDocsTopicStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DocsTopic) DeepCopyInto(out *DocsTopic) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DocsTopic.
func (in *DocsTopic) DeepCopy() *DocsTopic {
	if in == nil {
		return nil
	}
	out := new(DocsTopic)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DocsTopic) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DocsTopicList) DeepCopyInto(out *DocsTopicList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DocsTopic, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DocsTopicList.
func (in *DocsTopicList) DeepCopy() *DocsTopicList {
	if in == nil {
		return nil
	}
	out := new(DocsTopicList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DocsTopicList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DocsTopicSpec) DeepCopyInto(out *DocsTopicSpec) {
	*out = *in
	in.CommonDocsTopicSpec.DeepCopyInto(&out.CommonDocsTopicSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DocsTopicSpec.
func (in *DocsTopicSpec) DeepCopy() *DocsTopicSpec {
	if in == nil {
		return nil
	}
	out := new(DocsTopicSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DocsTopicStatus) DeepCopyInto(out *DocsTopicStatus) {
	*out = *in
	in.CommonDocsTopicStatus.DeepCopyInto(&out.CommonDocsTopicStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DocsTopicStatus.
func (in *DocsTopicStatus) DeepCopy() *DocsTopicStatus {
	if in == nil {
		return nil
	}
	out := new(DocsTopicStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Source) DeepCopyInto(out *Source) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Source.
func (in *Source) DeepCopy() *Source {
	if in == nil {
		return nil
	}
	out := new(Source)
	in.DeepCopyInto(out)
	return out
}
