package verif_module

import (
	"context"

	pb "github.com/BCinterfaceModified/verif-module/bcinterface"
)

type DecisionCriteria struct {
	Matric string
}

type vrfValue struct {
	Val    string `json:"val"`
	Proof  []byte `json:"pi"`
	PubKey []byte `json:"pk"`
}

type CommitteeNodeInfo struct {
	Round     int32
	Address   string
	VrfPubKey []byte
	VrfResult vrfValue
}

type CandidateNode struct {
	Criteria      DecisionCriteria
	CommitteeList []CommitteeNodeInfo
}

var CandidatesChannel = make(chan CandidateNode)

func (s *Server) RequestLeader(ctx context.Context, in *pb.RequestLeaderRequest) (*pb.RequestLeaderResponse, error) {
	go func() {
		task := CandidateNode{
			Criteria:      in.Nodeinfo,
			CommitteeList: in.Nodeinfo,
		}

		CandidatesChannel <- task
	}()

	return &pb.RequestLeaderResponse{
		Primarynode: "Node1",
		Code:        200,
	}, nil
}

func InitCommuModule() {
	startGrpcServer()
}

func ReturnPrimaryNode(PrimaryNodeInfo string) {

}
