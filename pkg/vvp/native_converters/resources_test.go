package nativeconverters

import (
	"github.com/fintechstudios/ververica-platform-k8s-operator/api/v1beta2"
	"github.com/fintechstudios/ververica-platform-k8s-operator/pkg/vvp/appmanager-api"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/api/resource"
)

var _ = Describe("Resources", func() {
	Describe("ResourcesToNative", func() {
		memory := "2g"
		cpu := 2.0
		resources := map[string]appmanagerapi.ResourceSpec{
			"jobmanager": {
				Cpu:    cpu,
				Memory: memory,
			},
			"taskmanager": {
				Cpu:    cpu,
				Memory: memory,
			},
		}

		It("should map a API resource to K8s native", func() {
			vpResources, err := ResourcesToNative(resources)
			Expect(err).ToNot(HaveOccurred())
			Expect(len(vpResources)).To(Equal(2))
			Expect(vpResources).To(HaveKey("jobmanager"))
			Expect(vpResources).To(HaveKey("taskmanager"))
			for _, resource := range vpResources {
				fmtCPU := resource.CPU.MilliValue() / 1000
				Expect(float64(fmtCPU)).To(Equal(cpu))
				Expect(*resource.Memory).To(Equal(memory))
			}
		})
	})

	Describe("ResourcesFromNative", func() {
		memory := "2g"
		cpu := resource.MustParse("2.0")
		vpResources := map[string]v1beta2.VpResourceSpec{
			"jobmanager": {
				CPU:    cpu,
				Memory: &memory,
			},
			"taskmanager": {
				CPU:    cpu,
				Memory: &memory,
			},
		}

		It("should map K8s native resources to API", func() {
			resources, err := ResourcesFromNative(vpResources)
			Expect(err).ToNot(HaveOccurred())
			for _, resource := range resources {
				fmtCPU := cpu.MilliValue() / 1000
				Expect(resource.Cpu).To(Equal(float64(fmtCPU)))
				Expect(resource.Memory).To(Equal(memory))
			}
		})
	})

})
