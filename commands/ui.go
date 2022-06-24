package commands

import (
	"fmt"
	"strconv"

	"github.com/playmean/guest/ui"

	"github.com/urfave/cli/v2"
)

func StartUI(c *cli.Context) error {
	s := ui.NewServer()

	port, err := resolveUIPort(c, 3080)
	if err != nil {
		return err
	}

	address := fmt.Sprintf("%s:%d", "localhost", port)

	// TODO 13.06.2022 get state from channel and open browser
	go ui.WaitEndpoint("http://" + address)

	err = s.Start(address)
	if err != nil {
		return err
	}

	return nil
}

func resolveUIPort(c *cli.Context, defaultPort int) (int, error) {
	portArg := c.Args().First()
	if portArg != "" {
		port, err := strconv.Atoi(portArg)
		if err != nil {
			return defaultPort, err
		}

		return port, nil
	}

	return defaultPort, nil
}
