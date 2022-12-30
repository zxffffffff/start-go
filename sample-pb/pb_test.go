package pb

import (
	"testing"

	"google.golang.org/protobuf/proto"
	pb "github.com/zxffffffff/start-go/sample-pb/pb_go"
)

func TestHeartbeatReq(t *testing.T) {
	req1 := pb.HeartbeatReq{
		ReqIndex: 123,
	}
	buf, err := proto.Marshal(&req1)
	if err != nil {
		t.Fatalf("TestHeartbeatReq Marshal error %q %v", buf, err)
	}
	req2 := pb.HeartbeatReq{}
	err = proto.Unmarshal(buf, &req2)
	if err != nil {
		t.Fatalf("TestHeartbeatReq Unmarshal error %q %v", buf, err)
	}
	if req1.ReqIndex != req2.ReqIndex {
		t.Fatalf("TestHeartbeatReq error %q %q", req1.String(), req2.String())
	}
}
