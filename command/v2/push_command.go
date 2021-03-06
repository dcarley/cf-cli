package v2

import (
	"os"

	"code.cloudfoundry.org/cli/cf/cmd"
	"code.cloudfoundry.org/cli/command"
	"code.cloudfoundry.org/cli/command/flag"
)

type PushCommand struct {
	AppPorts             string               `long:"app-ports" description:"Comma delimited list of ports the application may listen on" hidden:"true"` //TODO: Custom AppPorts flag
	BuildpackName        string               `short:"b" description:"Custom buildpack by name (e.g. my-buildpack) or Git URL (e.g. 'https://github.com/cloudfoundry/java-buildpack.git') or Git URL with a branch or tag (e.g. 'https://github.com/cloudfoundry/java-buildpack.git#v3.3.0' for 'v3.3.0' tag). To use built-in buildpacks only, specify 'default' or 'null'"`
	StartupCommand       string               `short:"c" description:"Startup command, set to null to reset to default start command"`
	Domain               string               `short:"d" description:"Domain (e.g. example.com)"`
	DockerImage          string               `long:"docker-image" short:"o" description:"Docker-image to be used (e.g. user/docker-image-name)"`
	PathToManifest       flag.Filename        `short:"f" description:"Path to manifest"`
	HealthCheckType      flag.HealthCheckType `long:"health-check-type" short:"u" description:"Application health check type (Default: 'port', 'none' accepted for 'process', 'http' implies endpoint '/')"`
	Hostname             string               `long:"hostname" short:"n" description:"Hostname (e.g. my-subdomain)"`
	NumInstances         int                  `short:"i" description:"Number of instances"`
	DiskLimit            string               `short:"k" description:"Disk limit (e.g. 256M, 1024M, 1G)"`
	MemoryLimit          string               `short:"m" description:"Memory limit (e.g. 256M, 1024M, 1G)"`
	NoHostname           bool                 `long:"no-hostname" description:"Map the root domain to this app"`
	NoManifest           bool                 `long:"no-manifest" description:"Ignore manifest file"`
	NoRoute              bool                 `long:"no-route" description:"Do not map a route to this app and remove routes from previous pushes of this app"`
	NoStart              bool                 `long:"no-start" description:"Do not start an app after pushing"`
	DirectoryPath        flag.Filename        `short:"p" description:"Path to app directory or to a zip file of the contents of the app directory"`
	RandomRoute          bool                 `long:"random-route" description:"Create a random route for this app"`
	RoutePath            string               `long:"route-path" description:"Path for the route"`
	Stack                string               `short:"s" description:"Stack to use (a stack is a pre-built file system, including an operating system, that can run apps)"`
	ApplicationStartTime int                  `short:"t" description:"Time (in seconds) allowed to elapse between starting up an app and the first healthy response from the app"`
	usage                interface{}          `usage:"Push a single app (with or without a manifest):\n   CF_NAME push APP_NAME [-b BUILDPACK_NAME] [-c COMMAND] [-d DOMAIN] [-f MANIFEST_PATH] [--docker-image DOCKER_IMAGE]\n   [-i NUM_INSTANCES] [-k DISK] [-m MEMORY] [--hostname HOST] [-p PATH] [-s STACK] [-t TIMEOUT] [-u (process | port | http)] [--route-path ROUTE_PATH]\n   [--no-hostname] [--no-manifest] [--no-route] [--no-start] [--random-route]\n\n   Push multiple apps with a manifest:\n   cf push [-f MANIFEST_PATH]"`
	envCFStagingTimeout  interface{}          `environmentName:"CF_STAGING_TIMEOUT" environmentDescription:"Max wait time for buildpack staging, in minutes" environmentDefault:"15"`
	envCFStartupTimeout  interface{}          `environmentName:"CF_STARTUP_TIMEOUT" environmentDescription:"Max wait time for app instance startup, in minutes" environmentDefault:"5"`
	relatedCommands      interface{}          `related_commands:"apps, create-app-manifest, logs, ssh, start"`
}

func (_ PushCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (_ PushCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
