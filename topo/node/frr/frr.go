package frr

import (
	"fmt"

	tpb "github.com/openconfig/kne/proto/topo"
	"github.com/openconfig/kne/topo/node"
)

func New(nodeImpl *node.Impl) (node.Node, error) {
	if nodeImpl == nil {
		return nil, fmt.Errorf("nodeImpl cannot be nil")
	}
	if nodeImpl.Proto == nil {
		return nil, fmt.Errorf("nodeImpl.Proto cannot be nil")
	}
	defaults(nodeImpl.Proto)
	n := &Node{
		Impl: nodeImpl,
	}
	return n, nil
}

type Node struct {
	*node.Impl
}

func defaults(pb *tpb.Node) *tpb.Node {
	if pb.Config == nil {
		pb.Config = &tpb.Config{}
	}
	if pb.Config.Image == "" {
		pb.Config.Image = "frrouting/frr:latest"
	}
	// if len(pb.GetConfig().GetCommand()) == 0 {
	// 	pb.Config.Command = []string{"/bin/sh"}
	// 	pb.Config.Args = []string{"/init.sh"}
	// }
	if pb.Config.EntryCommand == "" {
		pb.Config.EntryCommand = fmt.Sprintf("kubectl exec -it %s -- /bin/bash", pb.Name)
	}
	if pb.Config.ConfigPath == "" {
		pb.Config.ConfigPath = "/"
	}
	if pb.Config.ConfigFile == "" {
		pb.Config.ConfigFile = "init.sh"
	}
	return pb
}

func init() {
	node.Vendor(tpb.Vendor_FRR, New)
}
