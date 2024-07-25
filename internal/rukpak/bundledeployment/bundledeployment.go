/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package bundledeployment

// BundleDeploymentSpec defines the desired state of BundleDeployment
type BundleDeploymentSpec struct {
	// installNamespace is the namespace where the bundle should be installed. However, note that
	// the bundle may contain resources that are cluster-scoped or that are
	// installed in a different namespace. This namespace is expected to exist.
	InstallNamespace string

	// provisionerClassName sets the name of the provisioner that should reconcile this BundleDeployment.
	ProvisionerClassName string

	// source defines the configuration for the underlying Bundle content.
	Source BundleSource
}

// BundleDeployment is the Schema for the bundledeployments API
type BundleDeployment struct {
	Name string

	Spec BundleDeploymentSpec
}

type SourceType string

const SourceTypeImage SourceType = "image"

type BundleSource struct {
	// Type defines the kind of Bundle content being sourced.
	Type SourceType
	// Image is the bundle image that backs the content of this bundle.
	Image *ImageSource
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BundleDeploymentStatus.
// This is generated code copied from rukpak.
func (in *BundleSource) DeepCopy() *BundleSource {
	if in == nil {
		return nil
	}
	out := new(BundleSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
// This is generated code copied from rukpak.
func (in *BundleSource) DeepCopyInto(out *BundleSource) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(ImageSource)
		**out = **in
	}
}

type ImageSource struct {
	// Ref contains the reference to a container image containing Bundle contents.
	Ref string
	// ImagePullSecretName contains the name of the image pull secret in the namespace that the provisioner is deployed.
	ImagePullSecretName string
	// InsecureSkipTLSVerify indicates that TLS certificate validation should be skipped.
	// If this option is specified, the HTTPS protocol will still be used to
	// fetch the specified image reference.
	// This should not be used in a production environment.
	InsecureSkipTLSVerify bool
}
