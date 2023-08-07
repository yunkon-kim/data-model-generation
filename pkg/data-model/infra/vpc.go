package infra

// Ref.: https://github.com/cloud-barista/cb-spider/blob/master/cloud-control-manager/cloud-driver/drivers/azure/main/Test_Resources.go

type VPC struct {
	NameId   string `json:"nameId"  yaml:"nameId"  validate:"required"`
	IPv4CIDR string `json:"iPv4CIDR"  yaml:"iPv4CIDR"  validate:"cidrv4"`
	Subnets  []struct {
		NameId   string `json:"nameId"  yaml:"nameId"  validate:"required"`
		IPv4CIDR string `json:"iPv4CIDR"  yaml:"iPv4CIDR"  validate:"cidrv4"`
	} `json:"subnets"  yaml:"subnets"  validate:"required,min=1"`
}
