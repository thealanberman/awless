package aws

import (
	"fmt"

	"github.com/wallix/awless/cloud/rdf"
)

type property struct {
	AwlessLabel   string
	RDFLabel      string
	RDFType       string
	RdfsDefinedBy string
	RdfsDataType  string
}

var PropertiesDefinitions = []property{
	{AwlessLabel: "Actions", RDFLabel: fmt.Sprintf("%s:actions", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsList, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Affinity", RDFLabel: fmt.Sprintf("%s:affinity", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "ApproximateMessageCount", RDFLabel: fmt.Sprintf("%s:approximateMessageCount", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdInt},
	{AwlessLabel: "Architecture", RDFLabel: fmt.Sprintf("%s:architecture", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Arn", RDFLabel: fmt.Sprintf("%s:arn", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Attachable", RDFLabel: fmt.Sprintf("%s:attachable", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdBoolean},
	{AwlessLabel: "AutoUpgrade", RDFLabel: fmt.Sprintf("%s:autoUpgrade", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdBoolean},
	{AwlessLabel: "AvailabilityZone", RDFLabel: fmt.Sprintf("%s:availabilityZone", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsClass, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "AvailabilityZones", RDFLabel: fmt.Sprintf("%s:availabilityZones", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsList, RdfsDataType: rdf.RdfsClass},
	{AwlessLabel: "BackupRetentionPeriod", RDFLabel: fmt.Sprintf("%s:backupRetentionPeriod", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdDateTime},
	{AwlessLabel: "Bucket", RDFLabel: fmt.Sprintf("%s:bucketName", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsClass, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "CallerReference", RDFLabel: fmt.Sprintf("%s:callerReference", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "CertificateAuthority", RDFLabel: fmt.Sprintf("%s:certificateAuthority", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Certificates", RDFLabel: fmt.Sprintf("%s:certificates", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsList, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Charset", RDFLabel: fmt.Sprintf("%s:charset", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "CheckHTTPCode", RDFLabel: fmt.Sprintf("%s:checkHTTPCode", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "CheckInterval", RDFLabel: fmt.Sprintf("%s:checkInterval", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdInt},
	{AwlessLabel: "CheckPath", RDFLabel: fmt.Sprintf("%s:checkPath", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "CheckPort", RDFLabel: fmt.Sprintf("%s:checkPort", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "CheckProtocol", RDFLabel: fmt.Sprintf("%s:checkProtocol", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "CheckTimeout", RDFLabel: fmt.Sprintf("%s:checkTimeout", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdInt},
	{AwlessLabel: "CIDR", RDFLabel: fmt.Sprintf("%s:cidr", rdf.NetNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "CIDRv6", RDFLabel: fmt.Sprintf("%s:cidrv6", rdf.NetNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "CipherSuite", RDFLabel: fmt.Sprintf("%s:cipherSuite", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Class", RDFLabel: fmt.Sprintf("%s:class", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Cluster", RDFLabel: fmt.Sprintf("%s:cluster", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Comment", RDFLabel: rdf.RdfsComment, RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Continent", RDFLabel: fmt.Sprintf("%s:continent", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "CopyTagsToSnapshot", RDFLabel: fmt.Sprintf("%s:copyTagsToSnapshot", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Country", RDFLabel: fmt.Sprintf("%s:country", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Created", RDFLabel: fmt.Sprintf("%s:created", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdDateTime},
	{AwlessLabel: "DBSecurityGroups", RDFLabel: fmt.Sprintf("%s:dbSecurityGroups", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsList, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "DBSubnetGroup", RDFLabel: fmt.Sprintf("%s:dbSubnetGroup", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Default", RDFLabel: fmt.Sprintf("%s:default", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdBoolean},
	{AwlessLabel: "Delay", RDFLabel: fmt.Sprintf("%s:delaySeconds", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdInt},
	{AwlessLabel: "Description", RDFLabel: fmt.Sprintf("%s:description", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Encrypted", RDFLabel: fmt.Sprintf("%s:encrypted", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdBoolean},
	{AwlessLabel: "Endpoint", RDFLabel: fmt.Sprintf("%s:endpoint", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Engine", RDFLabel: fmt.Sprintf("%s:engine", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "EngineVersion", RDFLabel: fmt.Sprintf("%s:engineVersion", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Failover", RDFLabel: fmt.Sprintf("%s:failover", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Fingerprint", RDFLabel: fmt.Sprintf("%s:fingerprint", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "GlobalID", RDFLabel: fmt.Sprintf("%s:globalID", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "GranteeType", RDFLabel: fmt.Sprintf("%s:granteeType", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Grants", RDFLabel: fmt.Sprintf("%s:grants", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsList, RdfsDataType: rdf.Grant},
	{AwlessLabel: "Handler", RDFLabel: fmt.Sprintf("%s:handler", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Hash", RDFLabel: fmt.Sprintf("%s:hash", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "HealthCheck", RDFLabel: fmt.Sprintf("%s:healthCheck", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "HealthyThresholdCount", RDFLabel: fmt.Sprintf("%s:healthyThresholdCount", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdInt},
	{AwlessLabel: "Host", RDFLabel: fmt.Sprintf("%s:host", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Hypervisor", RDFLabel: fmt.Sprintf("%s:hypervisor", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "ID", RDFLabel: fmt.Sprintf("%s:id", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Image", RDFLabel: fmt.Sprintf("%s:image", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "InboundRules", RDFLabel: fmt.Sprintf("%s:inboundRules", rdf.NetNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsList, RdfsDataType: rdf.NetFirewallRule},
	{AwlessLabel: "InlinePolicies", RDFLabel: fmt.Sprintf("%s:inlinePolicies", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsList, RdfsDataType: rdf.RdfsClass},
	{AwlessLabel: "IOPS", RDFLabel: fmt.Sprintf("%s:iops", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdInt},
	{AwlessLabel: "IPType", RDFLabel: fmt.Sprintf("%s:ipType", rdf.NetNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Key", RDFLabel: fmt.Sprintf("%s:key", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "KeyPair", RDFLabel: fmt.Sprintf("%s:keyPair", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsClass, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "LatestRestorableTime", RDFLabel: fmt.Sprintf("%s:latestRestorableTime", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdDateTime},
	{AwlessLabel: "Launched", RDFLabel: fmt.Sprintf("%s:launched", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdDateTime},
	{AwlessLabel: "License", RDFLabel: fmt.Sprintf("%s:license", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Lifecycle", RDFLabel: fmt.Sprintf("%s:lifecycle", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "LoadBalancer", RDFLabel: fmt.Sprintf("%s:loadBalancer", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsClass, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Main", RDFLabel: fmt.Sprintf("%s:main", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdBoolean},
	{AwlessLabel: "Memory", RDFLabel: fmt.Sprintf("%s:memory", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdInt},
	{AwlessLabel: "Messages", RDFLabel: fmt.Sprintf("%s:messages", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsList, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Modified", RDFLabel: fmt.Sprintf("%s:modified", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdDateTime},
	{AwlessLabel: "MonitoringInterval", RDFLabel: fmt.Sprintf("%s:monitoringInterval", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "MonitoringRole", RDFLabel: fmt.Sprintf("%s:monitoringRole", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "MultiAZ", RDFLabel: fmt.Sprintf("%s:multiAZ", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Name", RDFLabel: fmt.Sprintf("%s:name", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "NetworkInterfaces", RDFLabel: fmt.Sprintf("%s:networkInterfaces", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsList, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "OptionGroups", RDFLabel: fmt.Sprintf("%s:optionGroups", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsList, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "OutboundRules", RDFLabel: fmt.Sprintf("%s:outboundRules", rdf.NetNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsList, RdfsDataType: rdf.NetFirewallRule},
	{AwlessLabel: "Owner", RDFLabel: fmt.Sprintf("%s:owner", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "ParameterGroups", RDFLabel: fmt.Sprintf("%s:parameterGroups", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsList, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "PasswordLastUsed", RDFLabel: fmt.Sprintf("%s:passwordLastUsed", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdDateTime},
	{AwlessLabel: "Path", RDFLabel: fmt.Sprintf("%s:path", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "PlacementGroup", RDFLabel: fmt.Sprintf("%s:placementGroup", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Port", RDFLabel: fmt.Sprintf("%s:port", rdf.NetNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdInt},
	{AwlessLabel: "PortRange", RDFLabel: fmt.Sprintf("%s:portRange", rdf.NetNS), RDFType: rdf.RdfsSubProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "PreferredBackupDate", RDFLabel: fmt.Sprintf("%s:preferredBackupDate", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "PreferredMaintenanceDate", RDFLabel: fmt.Sprintf("%s:preferredMaintenanceDate", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Private", RDFLabel: fmt.Sprintf("%s:private", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "PrivateIP", RDFLabel: fmt.Sprintf("%s:privateIP", rdf.NetNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Profile", RDFLabel: fmt.Sprintf("%s:profile", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Protocol", RDFLabel: fmt.Sprintf("%s:protocol", rdf.NetNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Public", RDFLabel: fmt.Sprintf("%s:public", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdBoolean},
	{AwlessLabel: "PublicDNS", RDFLabel: fmt.Sprintf("%s:publicDNS", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "PublicIP", RDFLabel: fmt.Sprintf("%s:publicIP", rdf.NetNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "RecordCount", RDFLabel: fmt.Sprintf("%s:records", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdInt},
	{AwlessLabel: "Records", RDFLabel: fmt.Sprintf("%s:recordCount", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsList, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Region", RDFLabel: fmt.Sprintf("%s:region", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Role", RDFLabel: fmt.Sprintf("%s:rootDeviceType", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsClass, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "RootDevice", RDFLabel: fmt.Sprintf("%s:role", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "RootDeviceType", RDFLabel: fmt.Sprintf("%s:rootDevice", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Routes", RDFLabel: fmt.Sprintf("%s:routes", rdf.NetNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsList, RdfsDataType: rdf.NetRoute},
	{AwlessLabel: "Runtime", RDFLabel: fmt.Sprintf("%s:runtime", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Scheme", RDFLabel: fmt.Sprintf("%s:scheme", rdf.NetNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "SecondaryAvailabilityZone", RDFLabel: fmt.Sprintf("%s:secondaryAvailabilityZone", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "SecurityGroups", RDFLabel: fmt.Sprintf("%s:securityGroups", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsList, RdfsDataType: rdf.RdfsClass},
	{AwlessLabel: "Set", RDFLabel: fmt.Sprintf("%s:set", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Size", RDFLabel: fmt.Sprintf("%s:size", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdInt},
	{AwlessLabel: "SpotInstanceRequestId", RDFLabel: fmt.Sprintf("%s:spotInstanceRequestId", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "SpotPrice", RDFLabel: fmt.Sprintf("%s:spotPrice", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "State", RDFLabel: fmt.Sprintf("%s:state", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Storage", RDFLabel: fmt.Sprintf("%s:storage", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "StorageType", RDFLabel: fmt.Sprintf("%s:storageType", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Subnet", RDFLabel: fmt.Sprintf("%s:subnet", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsClass, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Subnets", RDFLabel: fmt.Sprintf("%s:subnets", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsList, RdfsDataType: rdf.RdfsClass},
	{AwlessLabel: "Timeout", RDFLabel: fmt.Sprintf("%s:timezone", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdInt},
	{AwlessLabel: "Timezone", RDFLabel: fmt.Sprintf("%s:timeout", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Topic", RDFLabel: fmt.Sprintf("%s:topic", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsClass, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "TrafficPolicyInstance", RDFLabel: fmt.Sprintf("%s:trafficPolicyInstance", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "TTL", RDFLabel: fmt.Sprintf("%s:ttl", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdInt},
	{AwlessLabel: "Type", RDFLabel: fmt.Sprintf("%s:type", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "UnhealthyThresholdCount", RDFLabel: fmt.Sprintf("%s:unhealthyThresholdCount", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdInt},
	{AwlessLabel: "Updated", RDFLabel: fmt.Sprintf("%s:updated", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "UserData", RDFLabel: fmt.Sprintf("%s:userData", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Username", RDFLabel: fmt.Sprintf("%s:username", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Version", RDFLabel: fmt.Sprintf("%s:version", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Vpc", RDFLabel: fmt.Sprintf("%s:vpc", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsClass, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Vpcs", RDFLabel: fmt.Sprintf("%s:vpcs", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsList, RdfsDataType: rdf.RdfsClass},
	{AwlessLabel: "Weight", RDFLabel: fmt.Sprintf("%s:weight", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsLiteral, RdfsDataType: rdf.XsdString},
	{AwlessLabel: "Zone", RDFLabel: fmt.Sprintf("%s:zone", rdf.CloudNS), RDFType: rdf.RdfProperty, RdfsDefinedBy: rdf.RdfsClass, RdfsDataType: rdf.XsdString},
}