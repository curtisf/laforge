// Code generated by entc, DO NOT EDIT.

package ent

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
)

// Option function to configure the client.
type Option func(*config)

// Config is the configuration for the client and its builder.
type config struct {
	// driver used for executing database requests.
	driver dialect.Driver
	// debug enable a debug logging.
	debug bool
	// log used for logging on debug mode.
	log func(...interface{})
	// hooks to execute on mutations.
	hooks *hooks
}

// hooks per client, for fast access.
type hooks struct {
	AdhocPlan          []ent.Hook
	AgentStatus        []ent.Hook
	AgentTask          []ent.Hook
	AuthUser           []ent.Hook
	Build              []ent.Hook
	BuildCommit        []ent.Hook
	Command            []ent.Hook
	Competition        []ent.Hook
	DNS                []ent.Hook
	DNSRecord          []ent.Hook
	Disk               []ent.Hook
	Environment        []ent.Hook
	FileDelete         []ent.Hook
	FileDownload       []ent.Hook
	FileExtract        []ent.Hook
	Finding            []ent.Hook
	GinFileMiddleware  []ent.Hook
	Host               []ent.Hook
	HostDependency     []ent.Hook
	Identity           []ent.Hook
	IncludedNetwork    []ent.Hook
	Network            []ent.Hook
	Plan               []ent.Hook
	PlanDiff           []ent.Hook
	ProvisionedHost    []ent.Hook
	ProvisionedNetwork []ent.Hook
	ProvisioningStep   []ent.Hook
	Repository         []ent.Hook
	Script             []ent.Hook
	ServerTask         []ent.Hook
	Status             []ent.Hook
	Tag                []ent.Hook
	Team               []ent.Hook
	Token              []ent.Hook
	User               []ent.Hook
}

// Options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...interface{})) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}
