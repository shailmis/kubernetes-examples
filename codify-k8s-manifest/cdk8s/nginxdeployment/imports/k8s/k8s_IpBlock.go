// k8s
package k8s


// IPBlock describes a particular CIDR (Ex.
//
// "192.168.1.1/24","2001:db9::/64") that is allowed to the pods matched by a NetworkPolicySpec's podSelector. The except entry describes CIDRs that should not be included within this rule.
type IpBlock struct {
	// CIDR is a string representing the IP Block Valid examples are "192.168.1.1/24" or "2001:db9::/64".
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
	// Except is a slice of CIDRs that should not be included within an IP Block Valid examples are "192.168.1.1/24" or "2001:db9::/64" Except values will be rejected if they are outside the CIDR range.
	Except *[]*string `field:"optional" json:"except" yaml:"except"`
}

