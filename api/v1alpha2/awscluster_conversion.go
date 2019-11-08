/*
Copyright 2019 The Kubernetes Authors.

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

package v1alpha2

import (
	apiconversion "k8s.io/apimachinery/pkg/conversion"
	infrav1alpha3 "sigs.k8s.io/cluster-api-provider-aws/api/v1alpha3"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

const (
	unsupportedAWSClusterFieldsPrefix     = "unsupported-fields/awscluster.infrastructure.cluster.x-k8s.io/"
	awsClusterImageLookupOrgAnnotation    = unsupportedAWSClusterFieldsPrefix + "ImageLookupOrg"
	awsClusterImageLookupBaseOSAnnotation = unsupportedAWSClusterFieldsPrefix + "ImageLookupBaseOS"
)

// ConvertTo converts this AWSCluster to the Hub version (v1alpha3).
func (src *AWSCluster) ConvertTo(dstRaw conversion.Hub) error { // nolint
	dst := dstRaw.(*infrav1alpha3.AWSCluster)
	if err := Convert_v1alpha2_AWSCluster_To_v1alpha3_AWSCluster(src, dst, nil); err != nil {
		return err
	}

	// Restore the value of awsCluster.Spec.ImageLookupOrg if the annotation is present
	imageLookupOrg, ok := src.Annotations[awsClusterImageLookupOrgAnnotation]
	if ok {
		dst.Spec.ImageLookupOrg = imageLookupOrg
	}

	// Restore the value of awsCluster.Spec.ImageLookupBaseOS if the annotation is present
	imageLookupBaseOS, ok := src.Annotations[awsClusterImageLookupBaseOSAnnotation]
	if ok {
		dst.Spec.ImageLookupBaseOS = imageLookupBaseOS
	}

	return nil
}

// ConvertFrom converts from the Hub version (v1alpha3) to this version.
func (dst *AWSCluster) ConvertFrom(srcRaw conversion.Hub) error { // nolint
	src := srcRaw.(*infrav1alpha3.AWSCluster)
	if err := Convert_v1alpha3_AWSCluster_To_v1alpha2_AWSCluster(src, dst, nil); err != nil {
		return err
	}

	if src.Annotations == nil {
		src.Annotations = make(map[string]string)
	}

	// Preserve the value of awsCluster.Spec.ImageLookupOrg as an annotation
	src.Annotations[awsClusterImageLookupOrgAnnotation] = src.Spec.ImageLookupOrg

	// Preserve the value of awsCluster.Spec.ImageLookupBaseOS as an annotation
	src.Annotations[awsClusterImageLookupBaseOSAnnotation] = src.Spec.ImageLookupBaseOS

	return nil
}

// ConvertTo converts this AWSClusterList to the Hub version (v1alpha3).
func (src *AWSClusterList) ConvertTo(dstRaw conversion.Hub) error { // nolint
	dst := dstRaw.(*infrav1alpha3.AWSClusterList)
	return Convert_v1alpha2_AWSClusterList_To_v1alpha3_AWSClusterList(src, dst, nil)
}

// ConvertFrom converts from the Hub version (v1alpha3) to this version.
func (dst *AWSClusterList) ConvertFrom(srcRaw conversion.Hub) error { // nolint
	src := srcRaw.(*infrav1alpha3.AWSClusterList)
	return Convert_v1alpha3_AWSClusterList_To_v1alpha2_AWSClusterList(src, dst, nil)
}

// Convert_v1alpha3_AWSClusterSpec_To_v1alpha2_AWSClusterSpec converts from the Hub version (v1alpha3) of the AWSClusterSpec to this version.
// Requires manual conversion as infrav1alpha3.AWSClusterSpec.ImageLookupOrg and infrav1alpha3.AWSClusterSpec.ImageLookupBaseOS do not exist in AWSClusterSpec.
func Convert_v1alpha3_AWSClusterSpec_To_v1alpha2_AWSClusterSpec(in *infrav1alpha3.AWSClusterSpec, out *AWSClusterSpec, s apiconversion.Scope) error { // nolint
	if err := autoConvert_v1alpha3_AWSClusterSpec_To_v1alpha2_AWSClusterSpec(in, out, s); err != nil {
		return err
	}

	// ImageLookupOrg and ImageLookupBaseOS are preserved by the AWSCluster ConvertTo/ConvertFrom methods

	return nil
}
