package agones

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"os"

	"github.com/ribincao/ribin-dev-box/ribin-common/config"
	"github.com/ribincao/ribin-dev-box/ribin-common/constant"

	pb "agones.dev/agones/pkg/allocation/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type GameServer struct {
	Ip      string
	Port    int32
	PodName string
}

type agonesClient struct {
	client       pb.AllocationServiceClient
	region       string
	namespace    string
	multiCluster bool
}

var client map[string]*agonesClient // map[region]client

func InitAgones() {
	client = make(map[string]*agonesClient)
	// TODO 本地启动适配，后续删除
	for _, region := range config.GlobalConfig.AgonesConfig {
		cert, err := os.ReadFile("./cert/" + region.Prefix + "-client.crt")
		if err != nil {
			panic(err)
		}
		key, err := os.ReadFile("./cert/" + region.Prefix + "-client.key")
		if err != nil {
			panic(err)
		}
		caCert, err := os.ReadFile("./cert/" + region.Prefix + "-ca.crt")
		if err != nil {
			panic(err)
		}

		dialOpts, err := createRemoteClusterDialOption(cert, key, caCert)
		if err != nil {
			panic(err)
		}
		conn, err := grpc.Dial(region.AgonesAddr, dialOpts)
		if err != nil {
			panic(err)
		}

		client[region.Region] = &agonesClient{
			client:       pb.NewAllocationServiceClient(conn),
			region:       region.Region,
			namespace:    region.NameSpace,
			multiCluster: region.MultiCluster,
		}
	}

	fmt.Println("[Match-Agones] Agones Connect Success")
}

// createRemoteClusterDialOption creates a grpc client dial option with TLS configuration.
func createRemoteClusterDialOption(clientCert, clientKey, caCert []byte) (grpc.DialOption, error) {
	// Load client cert
	cert, err := tls.X509KeyPair(clientCert, clientKey)
	if err != nil {
		return nil, err
	}

	tlsConfig := &tls.Config{Certificates: []tls.Certificate{cert}}
	if len(caCert) != 0 {
		// Load CA cert, if provided and trust the server certificate.
		// This is required for self-signed certs.
		tlsConfig.RootCAs = x509.NewCertPool()
		if !tlsConfig.RootCAs.AppendCertsFromPEM(caCert) {
			return nil, errors.New("only PEM format is accepted for server CA")
		}
	}

	return grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)), nil
}

func MockLocalGameServer() (*GameServer, error) {
	gameServer := &GameServer{
		Ip:      "127.0.0.1",
		Port:    8888,
		PodName: "engine-server-us-v1-61-0-prod-19262-keyframe-vztkv-gkj4j",
	}
	return gameServer, nil
}

func CreateGameRoom(region, fleetId, mapId, platform, version string) (*GameServer, error) {
	if config.GlobalConfig.ServiceConfig.Env == constant.ENV_LOCAL {
		return MockLocalGameServer()
	}
	matchLabels := make(map[string]string)
	matchLabels["agones.dev/fleet"] = fleetId

	selectors := make([]*pb.GameServerSelector, 1)
	selectors[0] = &pb.GameServerSelector{
		MatchLabels:     matchLabels,
		GameServerState: pb.GameServerSelector_READY,
		Players:         nil,
	}

	c := client[region]

	request := &pb.AllocationRequest{
		Namespace: c.namespace, // k8s 的命名空间
		MultiClusterSetting: &pb.MultiClusterSetting{ // 多集群策略
			Enabled: c.multiCluster,
		},
		Scheduling:          pb.AllocationRequest_Packed, // 调度策略
		GameServerSelectors: selectors,
		Metadata: &pb.MetaPatch{
			Labels: map[string]string{"mapId": mapId, "platform": platform, "version": version},
		},
	}

	response, err := c.client.Allocate(context.Background(), request)
	if err != nil {
		return nil, err
	}

	gameServer := &GameServer{
		Ip:      response.Address,
		PodName: response.GameServerName,
	}

	for _, ports := range response.Ports {
		switch ports.Name {
		case "server":
			gameServer.Port = ports.Port
		default:
			err = errors.New("get ports err: port.Name = " + ports.Name)
			return nil, err
		}
	}

	return gameServer, nil
}
