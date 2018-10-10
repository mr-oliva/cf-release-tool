package client

import (
	"os/exec"

	"github.com/cloudfoundry/cli/plugin"
)

// Client operates cloudfoundry API.
type Client struct {
	cc plugin.CliConnection
}

// NewClient init Client
func NewClient(cc plugin.CliConnection) *Client {
	return &Client{
		cc: cc,
	}
}

// Init prepare material, git branch, and cf target.
func (c *Client) Init(materialDir, branch, org, space string) error {
	exec.Command("rm", "-rf", "./.bp-config").Run()
	if err := exec.Command("cp", "-rf", materialDir, "./.bp-config").Run(); err != nil {
		return err
	}
	if err := exec.Command("git", "checkout", branch).Run(); err != nil {
		return err
	}
	if err := exec.Command("git", "pull", "origin", branch).Run(); err != nil {
		return err
	}
	if _, err := c.cc.CliCommand("target", "-o", org, "-s", space); err != nil {
		return err
	}
	return nil
}

// Push executes cf push.
func (c *Client) Push(app, manifestFile string) error {
	if _, err := c.cc.CliCommand("push", app, "-f", manifestFile); err != nil {
		return err
	}
	return nil
}

// Rename executes cf rename.
func (c *Client) Rename(oldApp, newApp string) error {
	if _, err := c.cc.CliCommand("rename", oldApp, newApp); err != nil {
		return err
	}
	return nil
}

// Delete executes cf delete
func (c *Client) Delete(app string) error {
	if _, err := c.cc.CliCommand("delete", app); err != nil {
		return err
	}
	return nil
}

// MapRoute executes cf map-route
func (c *Client) MapRoute(app, domain, host string) error {
	if host != "" {
		if _, err := c.cc.CliCommand("map-route", app, domain, "--hostname", host); err != nil {
			return err
		}
	} else {
		if _, err := c.cc.CliCommand("map-route", app, domain); err != nil {
			return err
		}
	}
	return nil
}

// UnMapRoute executes cf unmap-route
func (c *Client) UnMapRoute(app, domain, host string) error {
	if host != "" {
		if _, err := c.cc.CliCommand("unmap-route", app, domain, "--hostname", host); err != nil {
			return err
		}
	} else {
		if _, err := c.cc.CliCommand("unmap-route", app, domain); err != nil {
			return err
		}
	}
	return nil
}

// AppExists check if there is a app in your space
func (c *Client) AppExists(app string) error {
	_, err := c.cc.GetApp(app)
	if err != nil {
		return err
	}
	return nil
}