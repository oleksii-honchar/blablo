package main

import (
	blablo "github.com/oleksii-honchar/blablo"
	c "github.com/oleksii-honchar/coteco"
)

func main() {
	logger := blablo.NewLogger("main", string(blablo.LevelDebug))
	logger.Info(c.WithGreenCyan49("Blablo going Green&Cyan"))
	logger.Warn("Example warning message")
	logger.Error("Example error message")
}
