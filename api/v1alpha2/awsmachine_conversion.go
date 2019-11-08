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
	unsupportedAWSMachineFieldsPrefix     = "unsupported-fields/awscluster.infrastructure.cluster.x-k8s.io/"
	awsMachineImageLookupBaseOSAnnotation = unsupportedAWSMachineFieldsPrefix + "ImageLookupBaseOS"
)

// ConvertTo converts this AWSMachine to the Hub version (v1alpha3).
func (src *AWSMachine) ConvertTo(dstRaw conversion.Hub) error { // nolint
	dst := dstRaw.(*infrav1alpha3.AWSMachine)
	if err := Convert_v1alpha2_AWSMachine_To_v1alpha3_AWSMachine(src, dst, nil); err != nil {
		return err
	}

	// Restore the value of awsCluster.Spec.ImageLookupBaseOS if the annotation is present
	imageLookupBaseOS, ok := src.Annotations[awsMachineImageLookupBaseOSAnnotation]
	if ok {
		dst.Spec.ImageLookupBaseOS = imageLookupBaseOS
	}

	return nil
}

// ConvertFrom converts from the Hub version (v1alpha3) to this version.
func (dst *AWSMachine) ConvertFrom(srcRaw conversion.Hub) error { // nolint
	src := srcRaw.(*infrav1alpha3.AWSMachine)
	if err := Convert_v1alpha3_AWSMachine_To_v1alpha2_AWSMachine(src, dst, nil); err != nil {
		return err
	}

	if src.Annotations == nil {
		src.Annotations = make(map[string]string)
	}

	// Preserve the value of awsMachine.Spec.ImageLookupBaseOS as an annotation
	src.Annotations[awsMachineImageLookupBaseOSAnnotation] = src.Spec.ImageLookupBaseOS

	return nil
}

// ConvertTo converts this AWSMachineList to the Hub version (v1alpha3).
func (src *AWSMachineList) ConvertTo(dstRaw conversion.Hub) error { // nolint
	dst := dstRaw.(*infrav1alpha3.AWSMachineList)
	return Convert_v1alpha2_AWSMachineList_To_v1alpha3_AWSMachineList(src, dst, nil)
}

// ConvertFrom converts from the Hub version (v1alpha3) to this version.
func (dst *AWSMachineList) ConvertFrom(srcRaw conversion.Hub) error { // nolint
	src := srcRaw.(*infrav1alpha3.AWSMachineList)
	return Convert_v1alpha3_AWSMachineList_To_v1alpha2_AWSMachineList(src, dst, nil)
}

// Convert_v1alpha3_AWSMachineSpec_To_v1alpha2_AWSMachineSpec converts from the Hub version (v1alpha3) of the AWSMachineSpec to this version.
// Requires manual conversion as infrav1alpha3.AWSMachineSpec.ImageLookupBaseOS does not exist in AWSMachineSpec.
func Convert_v1alpha3_AWSMachineSpec_To_v1alpha2_AWSMachineSpec(in *infrav1alpha3.AWSMachineSpec, out *AWSMachineSpec, s apiconversion.Scope) error { // nolint
	if err := autoConvert_v1alpha3_AWSMachineSpec_To_v1alpha2_AWSMachineSpec(in, out, s); err != nil {
		return err
	}

	// ImageLookupBaseOS is preserved by the AWSMachine ConvertTo/ConvertFrom methods

	return nil
}
