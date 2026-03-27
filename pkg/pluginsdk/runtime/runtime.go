package runtime

import (
	"context"
	"fmt"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"

	pluginv1 "github.com/ContinuumApp/continuum-plugin-sdk/pkg/pluginproto/continuum/plugin/v1"
)

const (
	ProtocolVersion  = 1
	MagicCookieKey   = "CONTINUUM_PLUGIN"
	MagicCookieValue = "continuum-rpc-plugin-v1"
	PluginSetName    = "continuum"
)

type CapabilityServers struct {
	Runtime          pluginv1.RuntimeServer
	MetadataProvider pluginv1.MetadataProviderServer
	MediaAnalyzer    pluginv1.MediaAnalyzerServer
	ScheduledTask    pluginv1.ScheduledTaskServer
	EventConsumer    pluginv1.EventConsumerServer
	AuthProvider     pluginv1.AuthProviderServer
	HttpRoutes       pluginv1.HttpRoutesServer
}

type Client struct {
	conn *grpc.ClientConn
}

type ServeConfig struct {
	Plugins plugin.PluginSet
	Logger  hclog.Logger
	Servers CapabilityServers
}

func HandshakeConfig() plugin.HandshakeConfig {
	return plugin.HandshakeConfig{
		ProtocolVersion:  ProtocolVersion,
		MagicCookieKey:   MagicCookieKey,
		MagicCookieValue: MagicCookieValue,
	}
}

func DefaultGRPCServer(opts []grpc.ServerOption) *grpc.Server {
	return grpc.NewServer(opts...)
}

func DefaultPluginSet(servers CapabilityServers) plugin.PluginSet {
	return plugin.PluginSet{
		PluginSetName: &GRPCPlugin{Servers: servers},
	}
}

func NewClient(conn *grpc.ClientConn) *Client {
	return &Client{conn: conn}
}

func (c *Client) Conn() *grpc.ClientConn {
	return c.conn
}

func (c *Client) Runtime() pluginv1.RuntimeClient {
	return pluginv1.NewRuntimeClient(c.conn)
}

func (c *Client) MetadataProvider() pluginv1.MetadataProviderClient {
	return pluginv1.NewMetadataProviderClient(c.conn)
}

func (c *Client) MediaAnalyzer() pluginv1.MediaAnalyzerClient {
	return pluginv1.NewMediaAnalyzerClient(c.conn)
}

func (c *Client) ScheduledTask() pluginv1.ScheduledTaskClient {
	return pluginv1.NewScheduledTaskClient(c.conn)
}

func (c *Client) EventConsumer() pluginv1.EventConsumerClient {
	return pluginv1.NewEventConsumerClient(c.conn)
}

func (c *Client) AuthProvider() pluginv1.AuthProviderClient {
	return pluginv1.NewAuthProviderClient(c.conn)
}

func (c *Client) HttpRoutes() pluginv1.HttpRoutesClient {
	return pluginv1.NewHttpRoutesClient(c.conn)
}

type GRPCPlugin struct {
	plugin.Plugin
	Servers CapabilityServers
}

func (p *GRPCPlugin) GRPCServer(_ *plugin.GRPCBroker, server *grpc.Server) error {
	if p.Servers.Runtime == nil {
		return fmt.Errorf("runtime server is required")
	}

	pluginv1.RegisterRuntimeServer(server, p.Servers.Runtime)
	if p.Servers.MetadataProvider != nil {
		pluginv1.RegisterMetadataProviderServer(server, p.Servers.MetadataProvider)
	}
	if p.Servers.MediaAnalyzer != nil {
		pluginv1.RegisterMediaAnalyzerServer(server, p.Servers.MediaAnalyzer)
	}
	if p.Servers.ScheduledTask != nil {
		pluginv1.RegisterScheduledTaskServer(server, p.Servers.ScheduledTask)
	}
	if p.Servers.EventConsumer != nil {
		pluginv1.RegisterEventConsumerServer(server, p.Servers.EventConsumer)
	}
	if p.Servers.AuthProvider != nil {
		pluginv1.RegisterAuthProviderServer(server, p.Servers.AuthProvider)
	}
	if p.Servers.HttpRoutes != nil {
		pluginv1.RegisterHttpRoutesServer(server, p.Servers.HttpRoutes)
	}
	return nil
}

func (p *GRPCPlugin) GRPCClient(_ context.Context, _ *plugin.GRPCBroker, conn *grpc.ClientConn) (interface{}, error) {
	return NewClient(conn), nil
}

func Serve(cfg ServeConfig) {
	// Handle "manifest" subcommand: print the plugin manifest as JSON and exit.
	if len(os.Args) > 1 && os.Args[1] == "manifest" {
		if cfg.Servers.Runtime == nil {
			fmt.Fprintln(os.Stderr, "runtime server is required to retrieve manifest")
			os.Exit(1)
		}
		resp, err := cfg.Servers.Runtime.GetManifest(context.Background(), &pluginv1.GetManifestRequest{})
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to get manifest: %v\n", err)
			os.Exit(1)
		}
		marshaler := protojson.MarshalOptions{Indent: "  "}
		data, err := marshaler.Marshal(resp.GetManifest())
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to encode manifest: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(data))
		os.Exit(0)
	}

	plugins := cfg.Plugins
	if len(plugins) == 0 {
		plugins = DefaultPluginSet(cfg.Servers)
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: HandshakeConfig(),
		Plugins:         plugins,
		GRPCServer:      plugin.DefaultGRPCServer,
		Logger:          cfg.Logger,
	})
}
