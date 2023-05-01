// k8s
package k8s


// APIService represents a server for a particular GroupVersion.
//
// Name must be "version.group".
type KubeApiServiceProps struct {
	// Standard object's metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
	// Spec contains information for locating and communicating with a server.
	Spec *ApiServiceSpec `field:"optional" json:"spec" yaml:"spec"`
}

