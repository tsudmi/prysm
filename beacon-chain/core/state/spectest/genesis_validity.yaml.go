// Code generated by yaml_to_go. DO NOT EDIT.
// source: genesis_initialization_minimal.yaml

package spectest

import pb "github.com/prysmaticlabs/prysm/proto/beacon/p2p/v1"

type GensisValidityTest struct {
	Title         string   `json:"title"`
	Summary       string   `json:"summary"`
	ForksTimeline string   `json:"forks_timeline"`
	Forks         []string `json:"forks"`
	Config        string   `json:"config"`
	Runner        string   `json:"runner"`
	Handler       string   `json:"handler"`
	TestCases     []struct {
		Description string          `json:"description"`
		Genesis     *pb.BeaconState `json:"genesis"`
		IsValid     bool            `json:"is_valid"`
	} `json:"test_cases"`
}
