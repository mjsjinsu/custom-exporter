package collector

type Vpc struct {
	GetVpcListResponse GetVpcListResponse `json:"getVpcListResponse"`
}
type VpcStatus struct {
	Code     string `json:"code"`
	CodeName string `json:"codeName"`
}
type VpcList struct {
	VpcNo         string    `json:"vpcNo"`
	VpcName       string    `json:"vpcName"`
	Ipv4CidrBlock string    `json:"ipv4CidrBlock"`
	VpcStatus     VpcStatus `json:"vpcStatus"`
	RegionCode    string    `json:"regionCode"`
	CreateDate    string    `json:"createDate"`
}
type GetVpcListResponse struct {
	RequestID     string    `json:"requestId"`
	ReturnCode    string    `json:"returnCode"`
	ReturnMessage string    `json:"returnMessage"`
	TotalRows     int       `json:"totalRows"`
	VpcList       []VpcList `json:"vpcList"`
}

type Subnet struct {
	GetSubnetListResponse GetSubnetListResponse `json:"getSubnetListResponse"`
}
type SubnetStatus struct {
	Code     string `json:"code"`
	CodeName string `json:"codeName"`
}
type SubnetType struct {
	Code     string `json:"code"`
	CodeName string `json:"codeName"`
}
type UsageType struct {
	Code     string `json:"code"`
	CodeName string `json:"codeName"`
}
type SubnetList struct {
	SubnetNo     string       `json:"subnetNo"`
	VpcNo        string       `json:"vpcNo"`
	ZoneCode     string       `json:"zoneCode"`
	SubnetName   string       `json:"subnetName"`
	Subnet       string       `json:"subnet"`
	SubnetStatus SubnetStatus `json:"subnetStatus"`
	CreateDate   string       `json:"createDate"`
	SubnetType   SubnetType   `json:"subnetType"`
	UsageType    UsageType    `json:"usageType"`
	NetworkACLNo string       `json:"networkAclNo"`
}
type GetSubnetListResponse struct {
	RequestID     string       `json:"requestId"`
	ReturnCode    string       `json:"returnCode"`
	ReturnMessage string       `json:"returnMessage"`
	TotalRows     int          `json:"totalRows"`
	SubnetList    []SubnetList `json:"subnetList"`
}

type NatGW struct {
	GetNatGatewayInstanceListResponse GetNatGatewayInstanceListResponse `json:"getNatGatewayInstanceListResponse"`
}
type NatGatewayInstanceStatus struct {
	Code     string `json:"code"`
	CodeName string `json:"codeName"`
}
type NatGatewayInstanceOperation struct {
}
type NatGatewayInstanceList struct {
	VpcName                      string                      `json:"vpcName"`
	NatGatewayInstanceNo         string                      `json:"natGatewayInstanceNo"`
	NatGatewayName               string                      `json:"natGatewayName"`
	PublicIP                     string                      `json:"publicIp"`
	NatGatewayInstanceStatus     NatGatewayInstanceStatus    `json:"natGatewayInstanceStatus"`
	NatGatewayInstanceStatusName string                      `json:"natGatewayInstanceStatusName"`
	NatGatewayInstanceOperation  NatGatewayInstanceOperation `json:"natGatewayInstanceOperation"`
	CreateDate                   string                      `json:"createDate"`
	NatGatewayDescription        string                      `json:"natGatewayDescription"`
	ZoneCode                     string                      `json:"zoneCode"`
}
type GetNatGatewayInstanceListResponse struct {
	RequestID              string                   `json:"requestId"`
	ReturnCode             string                   `json:"returnCode"`
	ReturnMessage          string                   `json:"returnMessage"`
	TotalRows              int                      `json:"totalRows"`
	NatGatewayInstanceList []NatGatewayInstanceList `json:"natGatewayInstanceList"`
}

type NAS struct {
	GetNasVolumeInstanceListResponse GetNasVolumeInstanceListResponse `json:"getNasVolumeInstanceListResponse"`
}
type NasVolumeInstanceStatus struct {
	Code     string `json:"code"`
	CodeName string `json:"codeName"`
}
type NasVolumeInstanceOperation struct {
	Code     string `json:"code"`
	CodeName string `json:"codeName"`
}
type VolumeAllotmentProtocolType struct {
	Code     string `json:"code"`
	CodeName string `json:"codeName"`
}
type NasVolumeServerInstanceNoList struct {
}
type NasVolumeInstanceList struct {
	NasVolumeInstanceNo              string                        `json:"nasVolumeInstanceNo"`
	NasVolumeInstanceStatus          NasVolumeInstanceStatus       `json:"nasVolumeInstanceStatus"`
	NasVolumeInstanceOperation       NasVolumeInstanceOperation    `json:"nasVolumeInstanceOperation"`
	NasVolumeInstanceStatusName      string                        `json:"nasVolumeInstanceStatusName"`
	CreateDate                       string                        `json:"createDate"`
	NasVolumeDescription             string                        `json:"nasVolumeDescription"`
	MountInformation                 string                        `json:"mountInformation"`
	VolumeAllotmentProtocolType      VolumeAllotmentProtocolType   `json:"volumeAllotmentProtocolType"`
	VolumeName                       string                        `json:"volumeName"`
	VolumeTotalSize                  int64                         `json:"volumeTotalSize"`
	VolumeSize                       int64                         `json:"volumeSize"`
	VolumeUseSize                    int                           `json:"volumeUseSize"`
	VolumeUseRatio                   int                           `json:"volumeUseRatio"`
	SnapshotVolumeConfigurationRatio int                           `json:"snapshotVolumeConfigurationRatio"`
	SnapshotVolumeSize               int                           `json:"snapshotVolumeSize"`
	SnapshotVolumeUseSize            int                           `json:"snapshotVolumeUseSize"`
	SnapshotVolumeUseRatio           int                           `json:"snapshotVolumeUseRatio"`
	IsSnapshotConfiguration          bool                          `json:"isSnapshotConfiguration"`
	IsEventConfiguration             bool                          `json:"isEventConfiguration"`
	RegionCode                       string                        `json:"regionCode"`
	ZoneCode                         string                        `json:"zoneCode"`
	NasVolumeServerInstanceNoList    NasVolumeServerInstanceNoList `json:"nasVolumeServerInstanceNoList"`
	IsEncryptedVolume                bool                          `json:"isEncryptedVolume"`
}
type GetNasVolumeInstanceListResponse struct {
	RequestID             string                  `json:"requestId"`
	ReturnCode            string                  `json:"returnCode"`
	ReturnMessage         string                  `json:"returnMessage"`
	TotalRows             int                     `json:"totalRows"`
	NasVolumeInstanceList []NasVolumeInstanceList `json:"nasVolumeInstanceList"`
}
