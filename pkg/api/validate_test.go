package api

import (
	"testing"

	"github.com/Azure/acs-engine/pkg/api/common"
)

func Test_OrchestratorProfile_ShouldErrorWithNoType(t *testing.T) {
	o := &OrchestratorProfile{}
	if err := o.Validate(); err == nil {
		t.Errorf("should error with no orchestrator type")
	}
}

func Test_OrchestratorProfile_DCOS_NoReleaseOK(t *testing.T) {
	o := &OrchestratorProfile{
		OrchestratorType: "DCOS",
	}
	if err := o.Validate(); err != nil {
		t.Errorf("should not error with orchestrator type '%s' and release '%s'", o.OrchestratorType, o.OrchestratorRelease)
	}
}

func Test_OrchestratorProfile_DCOS_SupportedReleasesOK(t *testing.T) {
	for _, r := range []string{common.DCOSRelease1Dot7, common.DCOSRelease1Dot8, common.DCOSRelease1Dot9} {
		o := &OrchestratorProfile{
			OrchestratorType:    "DCOS",
			OrchestratorRelease: r,
		}
		if err := o.Validate(); err != nil {
			t.Errorf("should not error with orchestrator type '%s' and release '%s'", o.OrchestratorType, o.OrchestratorRelease)
		}
	}
}

func Test_OrchestratorProfile_DCOS_UnsupportedReleasesFailValidation(t *testing.T) {
	for _, r := range []string{"0.3.141"} {
		o := &OrchestratorProfile{
			OrchestratorType:    "DCOS",
			OrchestratorRelease: r,
		}
		if err := o.Validate(); err == nil {
			t.Errorf("should error with orchestrator type '%s' and release '%s'", o.OrchestratorType, o.OrchestratorRelease)
		}
	}
}

func Test_OrchestratorProfile_Kubernetes_SupportedReleasesOK(t *testing.T) {
	for _, r := range []string{common.KubernetesRelease1Dot5, common.KubernetesRelease1Dot6, common.KubernetesRelease1Dot7} {
		o := &OrchestratorProfile{
			OrchestratorType:    "Kubernetes",
			OrchestratorRelease: r,
		}
		if err := o.Validate(); err != nil {
			t.Errorf("should not error with orchestrator type '%s' and release '%s'", o.OrchestratorType, o.OrchestratorRelease)
		}
	}
}

func Test_OrchestratorProfile_Kubernetes_UnsupportedReleasesFailValidation(t *testing.T) {
	for _, r := range []string{"1.1"} {
		o := &OrchestratorProfile{
			OrchestratorType:    "Kubernetes",
			OrchestratorRelease: r,
		}
		if err := o.Validate(); err == nil {
			t.Errorf("should error with orchestrator type '%s' and release '%s'", o.OrchestratorType, o.OrchestratorRelease)
		}
	}
}

func Test_OrchestratorProfile_FailsWhenKubernetesConfigPopulatedForNonKubernetesOrchestrator(t *testing.T) {
	for _, ot := range []string{Swarm, SwarmMode, DCOS, DockerCE} {
		o := &OrchestratorProfile{
			OrchestratorType:    ot,
			OrchestratorRelease: common.DCOSRelease1Dot9,
			KubernetesConfig: &KubernetesConfig{
				ClusterSubnet: "10.0.0.0/16",
			},
		}
		if err := o.Validate(); err == nil {
			t.Errorf("should error when KubernetesConfig populated for orchestrator type '%s'", o.OrchestratorType)
		}
	}
}

func Test_OrchestratorProfile_KubernetesConfigOKForKubernetes(t *testing.T) {
	o := &OrchestratorProfile{
		OrchestratorType:    Kubernetes,
		OrchestratorRelease: common.KubernetesRelease1Dot7,
		KubernetesConfig: &KubernetesConfig{
			ClusterSubnet: "10.0.0.0/16",
		},
	}
	if err := o.Validate(); err != nil {
		t.Errorf("should not error when KubernetesConfig populated for orchestrator type '%s'", o.OrchestratorType)
	}
}
