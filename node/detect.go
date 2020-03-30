package node

import (
	"github.com/cloudfoundry/packit"
)

func Detect() packit.DetectFunc {
	return func(context packit.DetectContext) (packit.DetectResult, error) {
		return packit.DetectResult{
			Plan: packit.BuildPlan{
				Provides: []packit.BuildPlanProvision{
					{Name: "node"},
				},
				Requires: []packit.BuildPlanRequirement{
					{Name: "node"},
				},
			},
		}, nil
	}
}
