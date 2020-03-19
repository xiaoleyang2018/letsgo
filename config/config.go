// Package config wraps ini-style config file and command-line args.
package config

import (
	"flag"
	"fmt"

	ini "gopkg.in/ini.v1"
)

var (
	runModeOptions   = []string{"dev", "prod"}
	logFormatOptions = []string{"text", "json"}
)

var (
	defaultName           = "app"
	defaultPort           = 8000
	defaultRunMode        = "dev"
	defaultLogPath        = "/var/log/app-server.log"
	defaultHandlerLogPath = "/var/log/app-handler.log"
	defaultLogFormat      = "text"
)

var (
	defaultDBType        = "mysql"
	defaultDBHost        = ""
	defaultDBPort        = 3306
	defaultDBUser        = "root"
	defaultDBPassword    = "123456"
	defaultDBName        = "users"
	defaultDBTablePrefix = ""
	defaultDBLogPath     = "/var/log/app-xorm.log"
)

// Version information.
var (
	ReleaseVersion = "None"
	BuildTS        = "None"
	GitHash        = "None"
	GitBranch      = "None"
)

// Version prints the version information.
func Version() {
	fmt.Println("Release Version:", ReleaseVersion)
	fmt.Println("Git Commit Hash:", GitHash)
	fmt.Println("Git Branch:", GitBranch)
	fmt.Println("UTC Build Time: ", BuildTS)
}

// Config set app config.
type Config struct {
	*flag.FlagSet
	configFile string
	Version    bool

	Name           string
	Port           int
	RunMode        string
	LogPath        string
	HandlerLogPath string
	LogFormat      string

	DB DBConfig
}

// DBConfig sets database config.
type DBConfig struct {
	Type        string
	Host        string
	Port        int
	User        string
	Password    string
	Name        string
	TablePrefix string
	LogPath     string
}

// New returns a new Config.
func New() *Config {
	cfg := &Config{}
	cfg.FlagSet = flag.NewFlagSet("app", flag.ContinueOnError)
	fs := cfg.FlagSet

	fs.BoolVar(&cfg.Version, "version", false, "print version information and exit")
	fs.StringVar(&cfg.configFile, "config", "./config/app.conf", "config file")
	fs.StringVar(&cfg.Name, "name", defaultName, "app name")
	fs.IntVar(&cfg.Port, "port", defaultPort, "port to listen")
	fs.StringVar(&cfg.RunMode, "run-mode", defaultRunMode, "run mode[dev,prod]")
	fs.StringVar(&cfg.LogPath, "log", defaultLogPath, "server log path")
	fs.StringVar(&cfg.HandlerLogPath, "handler-log", defaultHandlerLogPath, "handler log path")
	fs.StringVar(&cfg.LogFormat, "log-format", defaultLogFormat, "server log format[text, json]")
	fs.StringVar(&cfg.DB.Type, "db-type", defaultDBType, "database type")
	fs.StringVar(&cfg.DB.Host, "db-host", defaultDBHost, "database host")
	fs.IntVar(&cfg.DB.Port, "db-port", defaultDBPort, "database port")
	fs.StringVar(&cfg.DB.User, "db-user", defaultDBUser, "database user")
	fs.StringVar(&cfg.DB.Password, "db-password", defaultDBPassword, "database password")
	fs.StringVar(&cfg.DB.Name, "db-name", defaultDBName, "database name")
	fs.StringVar(&cfg.DB.TablePrefix, "db-table-prefix", defaultDBTablePrefix, "database table prefix")
	fs.StringVar(&cfg.DB.LogPath, "db-log", defaultDBLogPath, "database log path")

	return cfg
}

func (c *Config) parse(sec *ini.Section) {
	c.Name = sec.Key("name").MustString(defaultName)
	c.Port = sec.Key("port").MustInt(defaultPort)
	c.RunMode = sec.Key("runMode").In(defaultRunMode, runModeOptions)
	c.LogPath = sec.Key("logPath").MustString(defaultLogPath)
	c.HandlerLogPath = sec.Key("handlerLogPath").MustString(defaultHandlerLogPath)
	c.LogFormat = sec.Key("logFormat").In(defaultLogFormat, logFormatOptions)
}

func (c *DBConfig) parse(sec *ini.Section) {
	c.Type = sec.Key("type").MustString(defaultDBType)
	c.Host = sec.Key("host").MustString(defaultDBHost)
	c.Port = sec.Key("port").MustInt(defaultDBPort)
	c.User = sec.Key("user").MustString(defaultDBUser)
	c.Password = sec.Key("password").MustString(defaultDBPassword)
	c.Name = sec.Key("name").MustString(defaultDBName)
	c.TablePrefix = sec.Key("prefix").MustString(defaultDBTablePrefix)
	c.LogPath = sec.Key("logPath").MustString(defaultDBLogPath)
}

// Parse parse config file.
func (c *Config) Parse(args []string) error {
	// Parse first to get config file.
	err := c.FlagSet.Parse(args)
	if err != nil {
		return err
	}

	// Load config file if specified.
	if c.configFile != "" {
		cfg, err := ini.Load(c.configFile)
		if err != nil {
			return err
		}
		cfg.BlockMode = false // read only

		app, err := cfg.GetSection("app")
		if err != nil {
			return err
		}
		c.parse(app)

		db, err := cfg.GetSection("db")
		if err != nil {
			return err
		}
		c.DB.parse(db)
	}

	// Parse again to replace with command line options.
	return c.FlagSet.Parse(args)
}
