package terminal

import (
	"github.com/devtron-labs/devtron/internal/sql/repository"
	"github.com/devtron-labs/devtron/pkg/clusterTerminalAccess"
	"github.com/google/wire"
)

var terminalWireBaseSet = wire.NewSet(
	NewUserTerminalAccessRouterImpl,
	wire.Bind(new(UserTerminalAccessRouter), new(*UserTerminalAccessRouterImpl)),
	NewUserTerminalAccessRestHandlerImpl,
	wire.Bind(new(UserTerminalAccessRestHandler), new(*UserTerminalAccessRestHandlerImpl)),
	clusterTerminalAccess.GetTerminalAccessConfig,
	clusterTerminalAccess.NewUserTerminalAccessServiceImpl,
	wire.Bind(new(clusterTerminalAccess.UserTerminalAccessService), new(*clusterTerminalAccess.UserTerminalAccessServiceImpl)),
)

var TerminalWireSet = wire.NewSet(
	terminalWireBaseSet,
	repository.NewTerminalAccessRepositoryImpl,
	wire.Bind(new(repository.TerminalAccessRepository), new(*repository.TerminalAccessRepositoryImpl)),
)

var TerminalWireSetK8sClient = wire.NewSet(
	terminalWireBaseSet,
	repository.NewTerminalAccessFileBasedRepository,
	wire.Bind(new(repository.TerminalAccessRepository), new(*repository.TerminalAccessFileBasedRepository)),
)
