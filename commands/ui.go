package commands

import (
	"fmt"
	"log"

	"github.com/playmean/guest/ui"

	"github.com/urfave/cli/v2"
)

func StartUI(c *cli.Context) error {
	vars, err := resolveVariables(c)
	if err != nil {
		return err
	}

	w, err := resolveWorkspace(c, []string{})
	if err != nil {
		log.Println(err)
	}

	s := ui.NewServer(w, vars)

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
	portArg := c.Int("port")
	if portArg != 0 {
		return portArg, nil
	}

	return defaultPort, nil
}
