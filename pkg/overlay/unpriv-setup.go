package overlay

import (
	"github.com/raharper/stacker/pkg/types"
)

func UnprivSetup(config types.StackerConfig, uid, gid int) error {
	return Check(config)
}
