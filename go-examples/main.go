package main

import (
	"fmt"

	blablo "github.com/oleksii-honchar/blablo"
	c "github.com/oleksii-honchar/coteco"
)

var f = fmt.Sprintf

func main() {
	useColors := false

	logger := blablo.NewLogger("main", string(blablo.LevelDebug), useColors)
	logger.Info(c.WithGreenCyan49("Blablo going Green&Cyan"))
	logger.Warn("Example warning message")
	logger.Error("Example error message")

	data := struct {
		key1 string
		key2 string
	}{
		key1: "value1",
		key2: "value2",
	}

	logger.Info(f("Colored struct message: %s%+v%s", c.Yellow(), data, c.Reset()))
	logger.Info(f("Struct message: %+v", data))
	logger.Info(f("Struct message: %#v", data))

	logger1 := blablo.NewLogger("pref----10", string(blablo.LevelDebug), useColors)
	logger1.Info("Blablo 10 char prefix 'pref----10' example")
	logger2 := blablo.NewLogger("pref5", string(blablo.LevelDebug), useColors)
	logger2.Info("Blablo 5 char prefix 'pref5' example")
	logger3 := blablo.NewLogger("pref15---0----5", string(blablo.LevelDebug), useColors)
	logger3.Info("Blablo 5 char prefix 'pref15---0----5' example")

}
