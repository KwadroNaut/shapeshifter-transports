package replicant

import (
	"github.com/OperatorFoundation/shapeshifter-transports/transports/Replicant/polish"
	"github.com/OperatorFoundation/shapeshifter-transports/transports/Replicant/toneburst"
)

type Config struct {
	toneburst toneburst.Config
	polish    polish.Config
}
