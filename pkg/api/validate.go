package api

import (
	//	"errors"
	"fmt"
	//	"net"
	//	"net/url"
	//	"regexp"
	//	"time"

	"github.com/Azure/acs-engine/pkg/api/common"
	//	validator "gopkg.in/go-playground/validator.v9"
)

// Validate implements APIObject
func (o *OrchestratorProfile) Validate() error {
	switch o.OrchestratorType {
	case DCOS:
		switch o.OrchestratorRelease {
		case common.DCOSRelease1Dot7:
		case common.DCOSRelease1Dot8:
		case common.DCOSRelease1Dot9:
		case "":
		default:
			return fmt.Errorf("OrchestratorProfile has unknown orchestrator release: %s", o.OrchestratorRelease)
		}

	case Swarm:
	case SwarmMode:

	case Kubernetes:
		switch o.OrchestratorRelease {
		case common.KubernetesRelease1Dot7:
		case common.KubernetesRelease1Dot6:
		case common.KubernetesRelease1Dot5:
		case "":
		default:
			return fmt.Errorf("OrchestratorProfile has unknown orchestrator release: %s", o.OrchestratorRelease)
		}

		/* TODO: restore
		if o.KubernetesConfig != nil {
			err := o.KubernetesConfig.Validate(o.OrchestratorRelease)
			if err != nil {
				return err
			}
		}
		*/

	default:
		return fmt.Errorf("OrchestratorProfile has unknown orchestrator: %s", o.OrchestratorType)
	}

	if o.OrchestratorType != Kubernetes && o.KubernetesConfig != nil && (*o.KubernetesConfig != KubernetesConfig{}) {
		return fmt.Errorf("KubernetesConfig can be specified only when OrchestratorType is Kubernetes")
	}

	return nil
}

/*
// Validate implements APIObject
func (m *MasterProfile) Validate() error {
}

// Validate implements APIObject
func (a *AgentPoolProfile) Validate() error {
}

// Validate implements APIObject
func (l *LinuxProfile) Validate() error {
}

// Validate implements APIObject
func (a *Properties) Validate() error {
}
*/
