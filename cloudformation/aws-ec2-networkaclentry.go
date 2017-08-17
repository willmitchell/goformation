package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSEC2NetworkAclEntry AWS CloudFormation Resource (AWS::EC2::NetworkAclEntry)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl-entry.html
type AWSEC2NetworkAclEntry struct {

	// CidrBlock AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl-entry.html#cfn-ec2-networkaclentry-cidrblock
	CidrBlock string `json:"CidrBlock,omitempty"`

	// Egress AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl-entry.html#cfn-ec2-networkaclentry-egress
	Egress bool `json:"Egress,omitempty"`

	// Icmp AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl-entry.html#cfn-ec2-networkaclentry-icmp
	Icmp *AWSEC2NetworkAclEntry_Icmp `json:"Icmp,omitempty"`

	// Ipv6CidrBlock AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl-entry.html#cfn-ec2-networkaclentry-ipv6cidrblock
	Ipv6CidrBlock string `json:"Ipv6CidrBlock,omitempty"`

	// NetworkAclId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl-entry.html#cfn-ec2-networkaclentry-networkaclid
	NetworkAclId string `json:"NetworkAclId,omitempty"`

	// PortRange AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl-entry.html#cfn-ec2-networkaclentry-portrange
	PortRange *AWSEC2NetworkAclEntry_PortRange `json:"PortRange,omitempty"`

	// Protocol AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl-entry.html#cfn-ec2-networkaclentry-protocol
	Protocol int `json:"Protocol,omitempty"`

	// RuleAction AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl-entry.html#cfn-ec2-networkaclentry-ruleaction
	RuleAction string `json:"RuleAction,omitempty"`

	// RuleNumber AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl-entry.html#cfn-ec2-networkaclentry-rulenumber
	RuleNumber int `json:"RuleNumber,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2NetworkAclEntry) AWSCloudFormationType() string {
	return "AWS::EC2::NetworkAclEntry"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2NetworkAclEntry) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSEC2NetworkAclEntry) MarshalJSON() ([]byte, error) {
	type Properties AWSEC2NetworkAclEntry
	return json.Marshal(&struct {
		Type       string
		Properties Properties
	}{
		Type:       r.AWSCloudFormationType(),
		Properties: (Properties)(*r),
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *AWSEC2NetworkAclEntry) UnmarshalJSON(b []byte) error {
	type Properties AWSEC2NetworkAclEntry
	res := &struct {
		Type       string
		Properties *Properties
	}{}
	if err := json.Unmarshal(b, &res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}
	*r = AWSEC2NetworkAclEntry(*res.Properties)
	return nil
}

// GetAllAWSEC2NetworkAclEntryResources retrieves all AWSEC2NetworkAclEntry items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2NetworkAclEntryResources() map[string]AWSEC2NetworkAclEntry {
	results := map[string]AWSEC2NetworkAclEntry{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSEC2NetworkAclEntry:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::EC2::NetworkAclEntry" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSEC2NetworkAclEntry
						if err := json.Unmarshal(b, &result); err == nil {
							results[name] = result
						}
					}
				}
			}
		}
	}
	return results
}

// GetAWSEC2NetworkAclEntryWithName retrieves all AWSEC2NetworkAclEntry items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2NetworkAclEntryWithName(name string) (AWSEC2NetworkAclEntry, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSEC2NetworkAclEntry:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::EC2::NetworkAclEntry" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSEC2NetworkAclEntry
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSEC2NetworkAclEntry{}, errors.New("resource not found")
}