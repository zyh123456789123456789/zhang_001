// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Detailed information about the agent.
type AgentDetails struct {

	// Current agent version.
	//
	// This member is required.
	AgentVersion *string

	// List of versions being used by agent components.
	//
	// This member is required.
	ComponentVersions []ComponentVersion

	// ID of EC2 instance agent is running on.
	//
	// This member is required.
	InstanceId *string

	// Type of EC2 instance agent is running on.
	//
	// This member is required.
	InstanceType *string

	// List of CPU cores reserved for the agent.
	AgentCpuCores []int32

	// This field should not be used. Use agentCpuCores instead. List of CPU cores
	// reserved for processes other than the agent running on the EC2 instance.
	ReservedCpuCores []int32

	noSmithyDocumentSerde
}

// Aggregate status of Agent components.
type AggregateStatus struct {

	// Aggregate status.
	//
	// This member is required.
	Status AgentStatus

	// Sparse map of failure signatures.
	SignatureMap map[string]bool

	noSmithyDocumentSerde
}

// Details about an antenna demod decode Config used in a contact.
type AntennaDemodDecodeDetails struct {

	// Name of an antenna demod decode output node used in a contact.
	OutputNode *string

	noSmithyDocumentSerde
}

// Information about how AWS Ground Station should configure an antenna for
// downlink during a contact.
type AntennaDownlinkConfig struct {

	// Object that describes a spectral Config .
	//
	// This member is required.
	SpectrumConfig *SpectrumConfig

	noSmithyDocumentSerde
}

// Information about how AWS Ground Station should conﬁgure an antenna for
// downlink demod decode during a contact.
type AntennaDownlinkDemodDecodeConfig struct {

	// Information about the decode Config .
	//
	// This member is required.
	DecodeConfig *DecodeConfig

	// Information about the demodulation Config .
	//
	// This member is required.
	DemodulationConfig *DemodulationConfig

	// Information about the spectral Config .
	//
	// This member is required.
	SpectrumConfig *SpectrumConfig

	noSmithyDocumentSerde
}

// Information about the uplink Config of an antenna.
type AntennaUplinkConfig struct {

	// Information about the uplink spectral Config .
	//
	// This member is required.
	SpectrumConfig *UplinkSpectrumConfig

	// EIRP of the target.
	//
	// This member is required.
	TargetEirp *Eirp

	// Whether or not uplink transmit is disabled.
	TransmitDisabled *bool

	noSmithyDocumentSerde
}

// Information about AwsGroundStationAgentEndpoint.
type AwsGroundStationAgentEndpoint struct {

	// The egress address of AgentEndpoint.
	//
	// This member is required.
	EgressAddress *ConnectionDetails

	// The ingress address of AgentEndpoint.
	//
	// This member is required.
	IngressAddress *RangedConnectionDetails

	// Name string associated with AgentEndpoint. Used as a human-readable identifier
	// for AgentEndpoint.
	//
	// This member is required.
	Name *string

	// The status of AgentEndpoint.
	AgentStatus AgentStatus

	// The results of the audit.
	AuditResults AuditResults

	noSmithyDocumentSerde
}

// Data on the status of agent components.
type ComponentStatusData struct {

	// Capability ARN of the component.
	//
	// This member is required.
	CapabilityArn *string

	// The Component type.
	//
	// This member is required.
	ComponentType *string

	// Dataflow UUID associated with the component.
	//
	// This member is required.
	DataflowId *string

	// Component status.
	//
	// This member is required.
	Status AgentStatus

	// Bytes received by the component.
	BytesReceived *int64

	// Bytes sent by the component.
	BytesSent *int64

	// Packets dropped by component.
	PacketsDropped *int64

	noSmithyDocumentSerde
}

// Version information for agent components.
type ComponentVersion struct {

	// Component type.
	//
	// This member is required.
	ComponentType *string

	// List of versions.
	//
	// This member is required.
	Versions []string

	noSmithyDocumentSerde
}

// Details for certain Config object types in a contact.
//
// The following types satisfy this interface:
//
//	ConfigDetailsMemberAntennaDemodDecodeDetails
//	ConfigDetailsMemberEndpointDetails
//	ConfigDetailsMemberS3RecordingDetails
type ConfigDetails interface {
	isConfigDetails()
}

// Details for antenna demod decode Config in a contact.
type ConfigDetailsMemberAntennaDemodDecodeDetails struct {
	Value AntennaDemodDecodeDetails

	noSmithyDocumentSerde
}

func (*ConfigDetailsMemberAntennaDemodDecodeDetails) isConfigDetails() {}

// Information about the endpoint details.
type ConfigDetailsMemberEndpointDetails struct {
	Value EndpointDetails

	noSmithyDocumentSerde
}

func (*ConfigDetailsMemberEndpointDetails) isConfigDetails() {}

// Details for an S3 recording Config in a contact.
type ConfigDetailsMemberS3RecordingDetails struct {
	Value S3RecordingDetails

	noSmithyDocumentSerde
}

func (*ConfigDetailsMemberS3RecordingDetails) isConfigDetails() {}

// An item in a list of Config objects.
type ConfigListItem struct {

	// ARN of a Config .
	ConfigArn *string

	// UUID of a Config .
	ConfigId *string

	// Type of a Config .
	ConfigType ConfigCapabilityType

	// Name of a Config .
	Name *string

	noSmithyDocumentSerde
}

// Object containing the parameters of a Config . See the subtype definitions for
// what each type of Config contains.
//
// The following types satisfy this interface:
//
//	ConfigTypeDataMemberAntennaDownlinkConfig
//	ConfigTypeDataMemberAntennaDownlinkDemodDecodeConfig
//	ConfigTypeDataMemberAntennaUplinkConfig
//	ConfigTypeDataMemberDataflowEndpointConfig
//	ConfigTypeDataMemberS3RecordingConfig
//	ConfigTypeDataMemberTrackingConfig
//	ConfigTypeDataMemberUplinkEchoConfig
type ConfigTypeData interface {
	isConfigTypeData()
}

// Information about how AWS Ground Station should configure an antenna for
// downlink during a contact.
type ConfigTypeDataMemberAntennaDownlinkConfig struct {
	Value AntennaDownlinkConfig

	noSmithyDocumentSerde
}

func (*ConfigTypeDataMemberAntennaDownlinkConfig) isConfigTypeData() {}

// Information about how AWS Ground Station should conﬁgure an antenna for
// downlink demod decode during a contact.
type ConfigTypeDataMemberAntennaDownlinkDemodDecodeConfig struct {
	Value AntennaDownlinkDemodDecodeConfig

	noSmithyDocumentSerde
}

func (*ConfigTypeDataMemberAntennaDownlinkDemodDecodeConfig) isConfigTypeData() {}

// Information about how AWS Ground Station should conﬁgure an antenna for uplink
// during a contact.
type ConfigTypeDataMemberAntennaUplinkConfig struct {
	Value AntennaUplinkConfig

	noSmithyDocumentSerde
}

func (*ConfigTypeDataMemberAntennaUplinkConfig) isConfigTypeData() {}

// Information about the dataflow endpoint Config .
type ConfigTypeDataMemberDataflowEndpointConfig struct {
	Value DataflowEndpointConfig

	noSmithyDocumentSerde
}

func (*ConfigTypeDataMemberDataflowEndpointConfig) isConfigTypeData() {}

// Information about an S3 recording Config .
type ConfigTypeDataMemberS3RecordingConfig struct {
	Value S3RecordingConfig

	noSmithyDocumentSerde
}

func (*ConfigTypeDataMemberS3RecordingConfig) isConfigTypeData() {}

// Object that determines whether tracking should be used during a contact
// executed with this Config in the mission profile.
type ConfigTypeDataMemberTrackingConfig struct {
	Value TrackingConfig

	noSmithyDocumentSerde
}

func (*ConfigTypeDataMemberTrackingConfig) isConfigTypeData() {}

// Information about an uplink echo Config . Parameters from the
// AntennaUplinkConfig , corresponding to the specified AntennaUplinkConfigArn ,
// are used when this UplinkEchoConfig is used in a contact.
type ConfigTypeDataMemberUplinkEchoConfig struct {
	Value UplinkEchoConfig

	noSmithyDocumentSerde
}

func (*ConfigTypeDataMemberUplinkEchoConfig) isConfigTypeData() {}

// Egress address of AgentEndpoint with an optional mtu.
type ConnectionDetails struct {

	// A socket address.
	//
	// This member is required.
	SocketAddress *SocketAddress

	// Maximum transmission unit (MTU) size in bytes of a dataflow endpoint.
	Mtu *int32

	noSmithyDocumentSerde
}

// Data describing a contact.
type ContactData struct {

	// UUID of a contact.
	ContactId *string

	// Status of a contact.
	ContactStatus ContactStatus

	// End time of a contact in UTC.
	EndTime *time.Time

	// Error message of a contact.
	ErrorMessage *string

	// Name of a ground station.
	GroundStation *string

	// Maximum elevation angle of a contact.
	MaximumElevation *Elevation

	// ARN of a mission profile.
	MissionProfileArn *string

	// Amount of time after a contact ends that you’d like to receive a CloudWatch
	// event indicating the pass has finished.
	PostPassEndTime *time.Time

	// Amount of time prior to contact start you’d like to receive a CloudWatch event
	// indicating an upcoming pass.
	PrePassStartTime *time.Time

	// Region of a contact.
	Region *string

	// ARN of a satellite.
	SatelliteArn *string

	// Start time of a contact in UTC.
	StartTime *time.Time

	// Tags assigned to a contact.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Information about a dataflow edge used in a contact.
type DataflowDetail struct {

	// Dataflow details for the destination side.
	Destination *Destination

	// Error message for a dataflow.
	ErrorMessage *string

	// Dataflow details for the source side.
	Source *Source

	noSmithyDocumentSerde
}

// Information about a dataflow endpoint.
type DataflowEndpoint struct {

	// Socket address of a dataflow endpoint.
	Address *SocketAddress

	// Maximum transmission unit (MTU) size in bytes of a dataflow endpoint.
	Mtu *int32

	// Name of a dataflow endpoint.
	Name *string

	// Status of a dataflow endpoint.
	Status EndpointStatus

	noSmithyDocumentSerde
}

// Information about the dataflow endpoint Config .
type DataflowEndpointConfig struct {

	// Name of a dataflow endpoint.
	//
	// This member is required.
	DataflowEndpointName *string

	// Region of a dataflow endpoint.
	DataflowEndpointRegion *string

	noSmithyDocumentSerde
}

// Item in a list of DataflowEndpoint groups.
type DataflowEndpointListItem struct {

	// ARN of a dataflow endpoint group.
	DataflowEndpointGroupArn *string

	// UUID of a dataflow endpoint group.
	DataflowEndpointGroupId *string

	noSmithyDocumentSerde
}

// Information about the decode Config .
type DecodeConfig struct {

	// Unvalidated JSON of a decode Config .
	//
	// This member is required.
	UnvalidatedJSON *string

	noSmithyDocumentSerde
}

// Information about the demodulation Config .
type DemodulationConfig struct {

	// Unvalidated JSON of a demodulation Config .
	//
	// This member is required.
	UnvalidatedJSON *string

	noSmithyDocumentSerde
}

// Dataflow details for the destination side.
type Destination struct {

	// Additional details for a Config , if type is dataflow endpoint or antenna demod
	// decode.
	ConfigDetails ConfigDetails

	// UUID of a Config .
	ConfigId *string

	// Type of a Config .
	ConfigType ConfigCapabilityType

	// Region of a dataflow destination.
	DataflowDestinationRegion *string

	noSmithyDocumentSerde
}

// Data for agent discovery.
type DiscoveryData struct {

	// List of capabilities to associate with agent.
	//
	// This member is required.
	CapabilityArns []string

	// List of private IP addresses to associate with agent.
	//
	// This member is required.
	PrivateIpAddresses []string

	// List of public IP addresses to associate with agent.
	//
	// This member is required.
	PublicIpAddresses []string

	noSmithyDocumentSerde
}

// Object that represents EIRP.
type Eirp struct {

	// Units of an EIRP.
	//
	// This member is required.
	Units EirpUnits

	// Value of an EIRP. Valid values are between 20.0 to 50.0 dBW.
	//
	// This member is required.
	Value *float64

	noSmithyDocumentSerde
}

// Elevation angle of the satellite in the sky during a contact.
type Elevation struct {

	// Elevation angle units.
	//
	// This member is required.
	Unit AngleUnits

	// Elevation angle value.
	//
	// This member is required.
	Value *float64

	noSmithyDocumentSerde
}

// Information about the endpoint details.
type EndpointDetails struct {

	// An agent endpoint.
	AwsGroundStationAgentEndpoint *AwsGroundStationAgentEndpoint

	// A dataflow endpoint.
	Endpoint *DataflowEndpoint

	// Health reasons for a dataflow endpoint. This field is ignored when calling
	// CreateDataflowEndpointGroup .
	HealthReasons []CapabilityHealthReason

	// A dataflow endpoint health status. This field is ignored when calling
	// CreateDataflowEndpointGroup .
	HealthStatus CapabilityHealth

	// Endpoint security details including a list of subnets, a list of security
	// groups and a role to connect streams to instances.
	SecurityDetails *SecurityDetails

	noSmithyDocumentSerde
}

// Ephemeris data.
//
// The following types satisfy this interface:
//
//	EphemerisDataMemberOem
//	EphemerisDataMemberTle
type EphemerisData interface {
	isEphemerisData()
}

// Ephemeris data in Orbit Ephemeris Message (OEM) format.
type EphemerisDataMemberOem struct {
	Value OEMEphemeris

	noSmithyDocumentSerde
}

func (*EphemerisDataMemberOem) isEphemerisData() {}

// Two-line element set (TLE) ephemeris.
type EphemerisDataMemberTle struct {
	Value TLEEphemeris

	noSmithyDocumentSerde
}

func (*EphemerisDataMemberTle) isEphemerisData() {}

// Description of ephemeris.
type EphemerisDescription struct {

	// Supplied ephemeris data.
	EphemerisData *string

	// Source S3 object used for the ephemeris.
	SourceS3Object *S3Object

	noSmithyDocumentSerde
}

// Ephemeris item.
type EphemerisItem struct {

	// The time the ephemeris was uploaded in UTC.
	CreationTime *time.Time

	// Whether or not the ephemeris is enabled.
	Enabled *bool

	// The AWS Ground Station ephemeris ID.
	EphemerisId *string

	// A name string associated with the ephemeris. Used as a human-readable
	// identifier for the ephemeris.
	Name *string

	// Customer-provided priority score to establish the order in which overlapping
	// ephemerides should be used. The default for customer-provided ephemeris priority
	// is 1, and higher numbers take precedence. Priority must be 1 or greater
	Priority *int32

	// Source S3 object used for the ephemeris.
	SourceS3Object *S3Object

	// The status of the ephemeris.
	Status EphemerisStatus

	noSmithyDocumentSerde
}

// Metadata describing a particular ephemeris.
type EphemerisMetaData struct {

	// The EphemerisSource that generated a given ephemeris.
	//
	// This member is required.
	Source EphemerisSource

	// UUID of a customer-provided ephemeris. This field is not populated for default
	// ephemerides from Space Track.
	EphemerisId *string

	// The epoch of a default, ephemeris from Space Track in UTC. This field is not
	// populated for customer-provided ephemerides.
	Epoch *time.Time

	// A name string associated with the ephemeris. Used as a human-readable
	// identifier for the ephemeris. A name is only returned for customer-provider
	// ephemerides that have a name associated.
	Name *string

	noSmithyDocumentSerde
}

// The following types satisfy this interface:
//
//	EphemerisTypeDescriptionMemberOem
//	EphemerisTypeDescriptionMemberTle
type EphemerisTypeDescription interface {
	isEphemerisTypeDescription()
}

// Description of ephemeris.
type EphemerisTypeDescriptionMemberOem struct {
	Value EphemerisDescription

	noSmithyDocumentSerde
}

func (*EphemerisTypeDescriptionMemberOem) isEphemerisTypeDescription() {}

// Description of ephemeris.
type EphemerisTypeDescriptionMemberTle struct {
	Value EphemerisDescription

	noSmithyDocumentSerde
}

func (*EphemerisTypeDescriptionMemberTle) isEphemerisTypeDescription() {}

// Object that describes the frequency.
type Frequency struct {

	// Frequency units.
	//
	// This member is required.
	Units FrequencyUnits

	// Frequency value. Valid values are between 2200 to 2300 MHz and 7750 to 8400 MHz
	// for downlink and 2025 to 2120 MHz for uplink.
	//
	// This member is required.
	Value *float64

	noSmithyDocumentSerde
}

// Object that describes the frequency bandwidth.
type FrequencyBandwidth struct {

	// Frequency bandwidth units.
	//
	// This member is required.
	Units BandwidthUnits

	// Frequency bandwidth value. AWS Ground Station currently has the following
	// bandwidth limitations:
	//   - For AntennaDownlinkDemodDecodeconfig , valid values are between 125 kHz to
	//   650 MHz.
	//   - For AntennaDownlinkconfig , valid values are between 10 kHz to 54 MHz.
	//   - For AntennaUplinkConfig , valid values are between 10 kHz to 54 MHz.
	//
	// This member is required.
	Value *float64

	noSmithyDocumentSerde
}

// Information about the ground station data.
type GroundStationData struct {

	// UUID of a ground station.
	GroundStationId *string

	// Name of a ground station.
	GroundStationName *string

	// Ground station Region.
	Region *string

	noSmithyDocumentSerde
}

// An integer range that has a minimum and maximum value.
type IntegerRange struct {

	// A maximum value.
	//
	// This member is required.
	Maximum *int32

	// A minimum value.
	//
	// This member is required.
	Minimum *int32

	noSmithyDocumentSerde
}

// AWS Key Management Service (KMS) Key.
//
// The following types satisfy this interface:
//
//	KmsKeyMemberKmsAliasArn
//	KmsKeyMemberKmsKeyArn
type KmsKey interface {
	isKmsKey()
}

// KMS Alias Arn.
type KmsKeyMemberKmsAliasArn struct {
	Value string

	noSmithyDocumentSerde
}

func (*KmsKeyMemberKmsAliasArn) isKmsKey() {}

// KMS Key Arn.
type KmsKeyMemberKmsKeyArn struct {
	Value string

	noSmithyDocumentSerde
}

func (*KmsKeyMemberKmsKeyArn) isKmsKey() {}

// Item in a list of mission profiles.
type MissionProfileListItem struct {

	// ARN of a mission profile.
	MissionProfileArn *string

	// UUID of a mission profile.
	MissionProfileId *string

	// Name of a mission profile.
	Name *string

	// Region of a mission profile.
	Region *string

	noSmithyDocumentSerde
}

// Ephemeris data in Orbit Ephemeris Message (OEM) format.
type OEMEphemeris struct {

	// The data for an OEM ephemeris, supplied directly in the request rather than
	// through an S3 object.
	OemData *string

	// Identifies the S3 object to be used as the ephemeris.
	S3Object *S3Object

	noSmithyDocumentSerde
}

// Ingress address of AgentEndpoint with a port range and an optional mtu.
type RangedConnectionDetails struct {

	// A ranged socket address.
	//
	// This member is required.
	SocketAddress *RangedSocketAddress

	// Maximum transmission unit (MTU) size in bytes of a dataflow endpoint.
	Mtu *int32

	noSmithyDocumentSerde
}

// A socket address with a port range.
type RangedSocketAddress struct {

	// IPv4 socket address.
	//
	// This member is required.
	Name *string

	// Port range of a socket address.
	//
	// This member is required.
	PortRange *IntegerRange

	noSmithyDocumentSerde
}

// Object stored in S3 containing ephemeris data.
type S3Object struct {

	// An Amazon S3 Bucket name.
	Bucket *string

	// An Amazon S3 key for the ephemeris.
	Key *string

	// For versioned S3 objects, the version to use for the ephemeris.
	Version *string

	noSmithyDocumentSerde
}

// Information about an S3 recording Config .
type S3RecordingConfig struct {

	// ARN of the bucket to record to.
	//
	// This member is required.
	BucketArn *string

	// ARN of the role Ground Station assumes to write data to the bucket.
	//
	// This member is required.
	RoleArn *string

	// S3 Key prefix to prefice data files.
	Prefix *string

	noSmithyDocumentSerde
}

// Details about an S3 recording Config used in a contact.
type S3RecordingDetails struct {

	// ARN of the bucket used.
	BucketArn *string

	// Key template used for the S3 Recording Configuration
	KeyTemplate *string

	noSmithyDocumentSerde
}

// Item in a list of satellites.
type SatelliteListItem struct {

	// The current ephemeris being used to compute the trajectory of the satellite.
	CurrentEphemeris *EphemerisMetaData

	// A list of ground stations to which the satellite is on-boarded.
	GroundStations []string

	// NORAD satellite ID number.
	NoradSatelliteID int32

	// ARN of a satellite.
	SatelliteArn *string

	// UUID of a satellite.
	SatelliteId *string

	noSmithyDocumentSerde
}

// Information about endpoints.
type SecurityDetails struct {

	// ARN to a role needed for connecting streams to your instances.
	//
	// This member is required.
	RoleArn *string

	// The security groups to attach to the elastic network interfaces.
	//
	// This member is required.
	SecurityGroupIds []string

	// A list of subnets where AWS Ground Station places elastic network interfaces to
	// send streams to your instances.
	//
	// This member is required.
	SubnetIds []string

	noSmithyDocumentSerde
}

// Information about the socket address.
type SocketAddress struct {

	// Name of a socket address.
	//
	// This member is required.
	Name *string

	// Port of a socket address.
	//
	// This member is required.
	Port *int32

	noSmithyDocumentSerde
}

// Dataflow details for the source side.
type Source struct {

	// Additional details for a Config , if type is dataflow-endpoint or
	// antenna-downlink-demod-decode
	ConfigDetails ConfigDetails

	// UUID of a Config .
	ConfigId *string

	// Type of a Config .
	ConfigType ConfigCapabilityType

	// Region of a dataflow source.
	DataflowSourceRegion *string

	noSmithyDocumentSerde
}

// Object that describes a spectral Config .
type SpectrumConfig struct {

	// Bandwidth of a spectral Config . AWS Ground Station currently has the following
	// bandwidth limitations:
	//   - For AntennaDownlinkDemodDecodeconfig , valid values are between 125 kHz to
	//   650 MHz.
	//   - For AntennaDownlinkconfig valid values are between 10 kHz to 54 MHz.
	//   - For AntennaUplinkConfig , valid values are between 10 kHz to 54 MHz.
	//
	// This member is required.
	Bandwidth *FrequencyBandwidth

	// Center frequency of a spectral Config . Valid values are between 2200 to 2300
	// MHz and 7750 to 8400 MHz for downlink and 2025 to 2120 MHz for uplink.
	//
	// This member is required.
	CenterFrequency *Frequency

	// Polarization of a spectral Config . Capturing both "RIGHT_HAND" and "LEFT_HAND"
	// polarization requires two separate configs.
	Polarization Polarization

	noSmithyDocumentSerde
}

// A time range with a start and end time.
type TimeRange struct {

	// Time in UTC at which the time range ends.
	//
	// This member is required.
	EndTime *time.Time

	// Time in UTC at which the time range starts.
	//
	// This member is required.
	StartTime *time.Time

	noSmithyDocumentSerde
}

// Two-line element set (TLE) data.
type TLEData struct {

	// First line of two-line element set (TLE) data.
	//
	// This member is required.
	TleLine1 *string

	// Second line of two-line element set (TLE) data.
	//
	// This member is required.
	TleLine2 *string

	// The valid time range for the TLE. Gaps or overlap are not permitted.
	//
	// This member is required.
	ValidTimeRange *TimeRange

	noSmithyDocumentSerde
}

// Two-line element set (TLE) ephemeris.
type TLEEphemeris struct {

	// Identifies the S3 object to be used as the ephemeris.
	S3Object *S3Object

	// The data for a TLE ephemeris, supplied directly in the request rather than
	// through an S3 object.
	TleData []TLEData

	noSmithyDocumentSerde
}

// Object that determines whether tracking should be used during a contact
// executed with this Config in the mission profile.
type TrackingConfig struct {

	// Current setting for autotrack.
	//
	// This member is required.
	Autotrack Criticality

	noSmithyDocumentSerde
}

// Information about an uplink echo Config . Parameters from the
// AntennaUplinkConfig , corresponding to the specified AntennaUplinkConfigArn ,
// are used when this UplinkEchoConfig is used in a contact.
type UplinkEchoConfig struct {

	// ARN of an uplink Config .
	//
	// This member is required.
	AntennaUplinkConfigArn *string

	// Whether or not an uplink Config is enabled.
	//
	// This member is required.
	Enabled *bool

	noSmithyDocumentSerde
}

// Information about the uplink spectral Config .
type UplinkSpectrumConfig struct {

	// Center frequency of an uplink spectral Config . Valid values are between 2025 to
	// 2120 MHz.
	//
	// This member is required.
	CenterFrequency *Frequency

	// Polarization of an uplink spectral Config . Capturing both "RIGHT_HAND" and
	// "LEFT_HAND" polarization requires two separate configs.
	Polarization Polarization

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

// UnknownUnionMember is returned when a union member is returned over the wire,
// but has an unknown tag.
type UnknownUnionMember struct {
	Tag   string
	Value []byte

	noSmithyDocumentSerde
}

func (*UnknownUnionMember) isConfigDetails()            {}
func (*UnknownUnionMember) isConfigTypeData()           {}
func (*UnknownUnionMember) isEphemerisData()            {}
func (*UnknownUnionMember) isEphemerisTypeDescription() {}
func (*UnknownUnionMember) isKmsKey()                   {}