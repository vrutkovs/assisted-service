package nutanix

import (
	"github.com/openshift/assisted-service/internal/common"
)

func (p nutanixProvider) PreCreateManifestsHook(_ *common.Cluster, _ *[]string, workDir string) error {
	return nil
}
func (p nutanixProvider) PostCreateManifestsHook(_ *common.Cluster, _ *[]string, workDir string) error {
	return nil
}
