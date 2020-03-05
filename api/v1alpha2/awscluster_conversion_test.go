package v1alpha2

import (
	"testing"

	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	infrav1alpha3 "sigs.k8s.io/cluster-api-provider-aws/api/v1alpha3"
)

func TestConvertAWSCluster(t *testing.T) {
	g := NewWithT(t)

	t.Run("from hub", func(t *testing.T) {
		t.Run("should restore SSHKeyName, converting a nil value to an empty string", func(t *testing.T) {
			src := &infrav1alpha3.AWSCluster{
				ObjectMeta: metav1.ObjectMeta{},
				Spec: infrav1alpha3.AWSClusterSpec{
					SSHKeyName: nil,
				},
			}
			dst := &AWSCluster{}
			g.Expect(dst.ConvertFrom(src)).To(Succeed())
			restored := &infrav1alpha3.AWSCluster{}
			g.Expect(dst.ConvertTo(restored)).To(Succeed())
			g.Expect(restored.Spec.SSHKeyName).To(Equal(""))
		})
	})

}
