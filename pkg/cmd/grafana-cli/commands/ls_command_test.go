package commands

import (
	"errors"
	"github.com/grafana/grafana/pkg/cmd/grafana-cli/commands/commandstest"
	s "github.com/grafana/grafana/pkg/cmd/grafana-cli/services"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMissingPath(t *testing.T) {
	var org = validateLsCommand

	Convey("ls command", t, func() {
		validateLsCommand = org

		Convey("Missing path flag", func() {
			cmd := Command{
				Client: nil,
			}
			c, err := commandstest.NewCliContext([]string{"ls"})
			So(err, ShouldBeNil)
			s.IoHelper = &commandstest.FakeIoUtil{}

			Convey("should return error", func() {
				err := cmd.lsCommand(c)
				So(err, ShouldBeError, "missing path flag")
			})
		})

		Convey("Path is not a directory", func() {
			c, err := commandstest.NewCliContext([]string{"ls", "--path", "/var/lib/grafana/plugins"})
			So(err, ShouldBeNil)
			s.IoHelper = &commandstest.FakeIoUtil{
				FakeIsDirectory: false,
			}
			cmd := Command{
				Client: nil,
			}

			Convey("should return error", func() {
				err := cmd.lsCommand(c)
				So(err, ShouldNotBeNil)
			})
		})

		Convey("can override validateLsCommand", func() {
			c, err := commandstest.NewCliContext([]string{"ls", "--path", "/var/lib/grafana/plugins"})
			So(err, ShouldBeNil)

			validateLsCommand = func(pluginDir string) error {
				return errors.New("dummy error")
			}

			Convey("should return error", func() {
				cmd := Command{
					Client: nil,
				}
				err := cmd.lsCommand(c)
				So(err.Error(), ShouldEqual, "dummy error")
			})
		})

		Convey("Validate that validateLsCommand is reset", func() {
			c, err := commandstest.NewCliContext([]string{"ls", "--path", "/var/lib/grafana/plugins"})
			So(err, ShouldBeNil)
			cmd := Command{
				Client: nil,
			}

			Convey("should return error", func() {
				err := cmd.lsCommand(c)
				So(err.Error(), ShouldNotEqual, "dummy error")
			})
		})
	})
}
