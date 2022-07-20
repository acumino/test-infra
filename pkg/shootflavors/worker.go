package shootflavors

import (
	gardencorev1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	v1beta1constants "github.com/gardener/gardener/pkg/apis/core/v1beta1/constants"
	"k8s.io/utils/pointer"

	"github.com/gardener/test-infra/pkg/common"
	"github.com/gardener/test-infra/pkg/util"
)

func SetupWorker(cloudprofile gardencorev1beta1.CloudProfile, workers []gardencorev1beta1.Worker) ([]gardencorev1beta1.Worker, error) {
	res := make([]gardencorev1beta1.Worker, len(workers))
	for i, w := range workers {
		worker := w.DeepCopy()
		if worker.Machine.Image != nil && (worker.Machine.Image.Version == nil || *worker.Machine.Image.Version == common.PatternLatest) {
			if worker.Machine.Architecture == nil {
				worker.Machine.Architecture = pointer.String(v1beta1constants.ArchitectureAMD64)
			}
			version, err := util.GetLatestMachineImageVersion(cloudprofile, worker.Machine.Image.Name, *worker.Machine.Architecture)
			if err != nil {
				return nil, err
			}
			worker.Machine.Image.Version = &version.Version
		}
		res[i] = *worker
	}
	return res, nil
}
