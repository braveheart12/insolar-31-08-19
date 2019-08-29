// +build apitests

package apitests

// The identifier is to be incremented in every request and each response will contain a corresponding one.

const (
	// JSON RPC protocol version:
	JSONRPCVersion = "2.0"
	//  Endpoint URL for local deployment (is to be changed to a production URL):
	url = "http://localhost:19101/api/"
)

type nodePorts struct {
	Role              string
	ApiPort           string
	IntrospectionPort string
	Discovery         bool
}

type nodesPorts struct {
	Nodes []nodePorts
}

func getLaunchnet5Nodes() nodesPorts {
	ports := nodesPorts{
		Nodes: nil,
	}
	ports.Nodes = append(ports.Nodes, nodePorts{
		Role:              "Heavy",
		ApiPort:           "19101",
		IntrospectionPort: "55501",
		Discovery:         true,
	})
	ports.Nodes = append(ports.Nodes, nodePorts{
		Role:              "Virtual",
		ApiPort:           "19102",
		IntrospectionPort: "55502",
		Discovery:         true,
	})
	ports.Nodes = append(ports.Nodes, nodePorts{
		Role:              "Light",
		ApiPort:           "19103",
		IntrospectionPort: "55503",
		Discovery:         true,
	})
	ports.Nodes = append(ports.Nodes, nodePorts{
		Role:              "Virtual",
		ApiPort:           "19104",
		IntrospectionPort: "55504",
		Discovery:         true,
	})
	ports.Nodes = append(ports.Nodes, nodePorts{
		Role:              "Light",
		ApiPort:           "19105",
		IntrospectionPort: "55505",
		Discovery:         true,
	})

	return ports
}
