// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// The geographic extent of the Earth Observation job.
//
// The following types satisfy this interface:
//
//	AreaOfInterestMemberAreaOfInterestGeometry
type AreaOfInterest interface {
	isAreaOfInterest()
}

// A GeoJSON object representing the geographic extent in the coordinate space.
type AreaOfInterestMemberAreaOfInterestGeometry struct {
	Value AreaOfInterestGeometry

	noSmithyDocumentSerde
}

func (*AreaOfInterestMemberAreaOfInterestGeometry) isAreaOfInterest() {}

// A GeoJSON object representing the geographic extent in the coordinate space.
//
// The following types satisfy this interface:
//
//	AreaOfInterestGeometryMemberMultiPolygonGeometry
//	AreaOfInterestGeometryMemberPolygonGeometry
type AreaOfInterestGeometry interface {
	isAreaOfInterestGeometry()
}

// The structure representing the MultiPolygon Geometry.
type AreaOfInterestGeometryMemberMultiPolygonGeometry struct {
	Value MultiPolygonGeometryInput

	noSmithyDocumentSerde
}

func (*AreaOfInterestGeometryMemberMultiPolygonGeometry) isAreaOfInterestGeometry() {}

// The structure representing Polygon Geometry.
type AreaOfInterestGeometryMemberPolygonGeometry struct {
	Value PolygonGeometryInput

	noSmithyDocumentSerde
}

func (*AreaOfInterestGeometryMemberPolygonGeometry) isAreaOfInterestGeometry() {}

// The structure containing the asset properties.
type AssetValue struct {

	// Link to the asset object.
	Href *string

	noSmithyDocumentSerde
}

// Input structure for the BandMath operation type. Defines Predefined and
// CustomIndices to be computed using BandMath.
type BandMathConfigInput struct {

	// CustomIndices that are computed.
	CustomIndices *CustomIndicesInput

	// One or many of the supported predefined indices to compute. Allowed values: NDVI
	// , EVI2 , MSAVI , NDWI , NDMI , NDSI , and WDRVI .
	PredefinedIndices []string

	noSmithyDocumentSerde
}

// Input structure for CloudMasking operation type.
type CloudMaskingConfigInput struct {
	noSmithyDocumentSerde
}

// Input structure for Cloud Removal Operation type
type CloudRemovalConfigInput struct {

	// The name of the algorithm used for cloud removal.
	AlgorithmName AlgorithmNameCloudRemoval

	// The interpolation value you provide for cloud removal.
	InterpolationValue *string

	// TargetBands to be returned in the output of CloudRemoval operation.
	TargetBands []string

	noSmithyDocumentSerde
}

// Input object defining the custom BandMath indices to compute.
type CustomIndicesInput struct {

	// A list of BandMath indices to compute.
	Operations []Operation

	noSmithyDocumentSerde
}

// The structure representing the errors in an EarthObservationJob.
type EarthObservationJobErrorDetails struct {

	// A detailed message describing the error in an Earth Observation job.
	Message *string

	// The type of error in an Earth Observation job.
	Type EarthObservationJobErrorType

	noSmithyDocumentSerde
}

// The structure representing the EoCloudCover filter.
type EoCloudCoverInput struct {

	// Lower bound for EoCloudCover.
	//
	// This member is required.
	LowerBound *float32

	// Upper bound for EoCloudCover.
	//
	// This member is required.
	UpperBound *float32

	noSmithyDocumentSerde
}

// The structure for returning the export error details in a
// GetEarthObservationJob.
type ExportErrorDetails struct {

	// The structure for returning the export error details while exporting results of
	// an Earth Observation job.
	ExportResults *ExportErrorDetailsOutput

	// The structure for returning the export error details while exporting the source
	// images of an Earth Observation job.
	ExportSourceImages *ExportErrorDetailsOutput

	noSmithyDocumentSerde
}

// The structure representing the errors in an export EarthObservationJob
// operation.
type ExportErrorDetailsOutput struct {

	// A detailed message describing the error in an export EarthObservationJob
	// operation.
	Message *string

	// The type of error in an export EarthObservationJob operation.
	Type ExportErrorType

	noSmithyDocumentSerde
}

// The structure containing the Amazon S3 path to export the Earth Observation job
// output.
type ExportS3DataInput struct {

	// The URL to the Amazon S3 data input.
	//
	// This member is required.
	S3Uri *string

	// The Key Management Service key ID for server-side encryption.
	KmsKeyId *string

	noSmithyDocumentSerde
}

// An object containing information about the output file.
type ExportVectorEnrichmentJobOutputConfig struct {

	// The input structure for Amazon S3 data; representing the Amazon S3 location of
	// the input data objects.
	//
	// This member is required.
	S3Data *VectorEnrichmentJobS3Data

	noSmithyDocumentSerde
}

// The structure representing the filters supported by a RasterDataCollection.
type Filter struct {

	// The name of the filter.
	//
	// This member is required.
	Name *string

	// The type of the filter being used.
	//
	// This member is required.
	Type *string

	// The maximum value of the filter.
	Maximum *float32

	// The minimum value of the filter.
	Minimum *float32

	noSmithyDocumentSerde
}

// The structure representing a Geometry in terms of Type and Coordinates as per
// GeoJson spec.
type Geometry struct {

	// The coordinates of the GeoJson Geometry.
	//
	// This member is required.
	Coordinates [][][]float64

	// GeoJson Geometry types like Polygon and MultiPolygon.
	//
	// This member is required.
	Type *string

	noSmithyDocumentSerde
}

// Input configuration information for the geomosaic.
type GeoMosaicConfigInput struct {

	// The name of the algorithm being used for geomosaic.
	AlgorithmName AlgorithmNameGeoMosaic

	// The target bands for geomosaic.
	TargetBands []string

	noSmithyDocumentSerde
}

// Input configuration information.
type InputConfigInput struct {

	// The Amazon Resource Name (ARN) of the previous Earth Observation job.
	PreviousEarthObservationJobArn *string

	// The structure representing the RasterDataCollection Query consisting of the
	// Area of Interest, RasterDataCollectionArn,TimeRange and Property Filters.
	RasterDataCollectionQuery *RasterDataCollectionQueryInput

	noSmithyDocumentSerde
}

// The InputConfig for an EarthObservationJob response.
type InputConfigOutput struct {

	// The Amazon Resource Name (ARN) of the previous Earth Observation job.
	PreviousEarthObservationJobArn *string

	// The structure representing the RasterDataCollection Query consisting of the
	// Area of Interest, RasterDataCollectionArn, RasterDataCollectionName, TimeRange,
	// and Property Filters.
	RasterDataCollectionQuery *RasterDataCollectionQueryOutput

	noSmithyDocumentSerde
}

// The structure representing the items in the response for
// SearchRasterDataCollection.
type ItemSource struct {

	// The searchable date and time of the item, in UTC.
	//
	// This member is required.
	DateTime *time.Time

	// The item Geometry in GeoJson format.
	//
	// This member is required.
	Geometry *Geometry

	// A unique Id for the source item.
	//
	// This member is required.
	Id *string

	// This is a dictionary of Asset Objects data associated with the Item that can be
	// downloaded or streamed, each with a unique key.
	Assets map[string]AssetValue

	// This field contains additional properties of the item.
	Properties *Properties

	noSmithyDocumentSerde
}

// The input structure for the JobConfig in an EarthObservationJob.
//
// The following types satisfy this interface:
//
//	JobConfigInputMemberBandMathConfig
//	JobConfigInputMemberCloudMaskingConfig
//	JobConfigInputMemberCloudRemovalConfig
//	JobConfigInputMemberGeoMosaicConfig
//	JobConfigInputMemberLandCoverSegmentationConfig
//	JobConfigInputMemberResamplingConfig
//	JobConfigInputMemberStackConfig
//	JobConfigInputMemberTemporalStatisticsConfig
//	JobConfigInputMemberZonalStatisticsConfig
type JobConfigInput interface {
	isJobConfigInput()
}

// An object containing information about the job configuration for BandMath.
type JobConfigInputMemberBandMathConfig struct {
	Value BandMathConfigInput

	noSmithyDocumentSerde
}

func (*JobConfigInputMemberBandMathConfig) isJobConfigInput() {}

// An object containing information about the job configuration for cloud masking.
type JobConfigInputMemberCloudMaskingConfig struct {
	Value CloudMaskingConfigInput

	noSmithyDocumentSerde
}

func (*JobConfigInputMemberCloudMaskingConfig) isJobConfigInput() {}

// An object containing information about the job configuration for cloud removal.
type JobConfigInputMemberCloudRemovalConfig struct {
	Value CloudRemovalConfigInput

	noSmithyDocumentSerde
}

func (*JobConfigInputMemberCloudRemovalConfig) isJobConfigInput() {}

// An object containing information about the job configuration for geomosaic.
type JobConfigInputMemberGeoMosaicConfig struct {
	Value GeoMosaicConfigInput

	noSmithyDocumentSerde
}

func (*JobConfigInputMemberGeoMosaicConfig) isJobConfigInput() {}

// An object containing information about the job configuration for land cover
// segmentation.
type JobConfigInputMemberLandCoverSegmentationConfig struct {
	Value LandCoverSegmentationConfigInput

	noSmithyDocumentSerde
}

func (*JobConfigInputMemberLandCoverSegmentationConfig) isJobConfigInput() {}

// An object containing information about the job configuration for resampling.
type JobConfigInputMemberResamplingConfig struct {
	Value ResamplingConfigInput

	noSmithyDocumentSerde
}

func (*JobConfigInputMemberResamplingConfig) isJobConfigInput() {}

// An object containing information about the job configuration for a Stacking
// Earth Observation job.
type JobConfigInputMemberStackConfig struct {
	Value StackConfigInput

	noSmithyDocumentSerde
}

func (*JobConfigInputMemberStackConfig) isJobConfigInput() {}

// An object containing information about the job configuration for temporal
// statistics.
type JobConfigInputMemberTemporalStatisticsConfig struct {
	Value TemporalStatisticsConfigInput

	noSmithyDocumentSerde
}

func (*JobConfigInputMemberTemporalStatisticsConfig) isJobConfigInput() {}

// An object containing information about the job configuration for zonal
// statistics.
type JobConfigInputMemberZonalStatisticsConfig struct {
	Value ZonalStatisticsConfigInput

	noSmithyDocumentSerde
}

func (*JobConfigInputMemberZonalStatisticsConfig) isJobConfigInput() {}

// The input structure for Land Cover Operation type.
type LandCoverSegmentationConfigInput struct {
	noSmithyDocumentSerde
}

// The structure representing Land Cloud Cover property for Landsat data
// collection.
type LandsatCloudCoverLandInput struct {

	// The minimum value for Land Cloud Cover property filter. This will filter items
	// having Land Cloud Cover greater than or equal to this value.
	//
	// This member is required.
	LowerBound *float32

	// The maximum value for Land Cloud Cover property filter. This will filter items
	// having Land Cloud Cover less than or equal to this value.
	//
	// This member is required.
	UpperBound *float32

	noSmithyDocumentSerde
}

// An object containing information about the output file.
type ListEarthObservationJobOutputConfig struct {

	// The Amazon Resource Name (ARN) of the list of the Earth Observation jobs.
	//
	// This member is required.
	Arn *string

	// The creation time.
	//
	// This member is required.
	CreationTime *time.Time

	// The duration of the session, in seconds.
	//
	// This member is required.
	DurationInSeconds *int32

	// The names of the Earth Observation jobs in the list.
	//
	// This member is required.
	Name *string

	// The operation type for an Earth Observation job.
	//
	// This member is required.
	OperationType *string

	// The status of the list of the Earth Observation jobs.
	//
	// This member is required.
	Status EarthObservationJobStatus

	// Each tag consists of a key and a value.
	Tags map[string]string

	noSmithyDocumentSerde
}

// An object containing information about the output file.
type ListVectorEnrichmentJobOutputConfig struct {

	// The Amazon Resource Name (ARN) of the list of the Vector Enrichment jobs.
	//
	// This member is required.
	Arn *string

	// The creation time.
	//
	// This member is required.
	CreationTime *time.Time

	// The duration of the session, in seconds.
	//
	// This member is required.
	DurationInSeconds *int32

	// The names of the Vector Enrichment jobs in the list.
	//
	// This member is required.
	Name *string

	// The status of the Vector Enrichment jobs list.
	//
	// This member is required.
	Status VectorEnrichmentJobStatus

	// The type of the list of Vector Enrichment jobs.
	//
	// This member is required.
	Type VectorEnrichmentJobType

	// Each tag consists of a key and a value.
	Tags map[string]string

	noSmithyDocumentSerde
}

// The input structure for Map Matching operation type.
type MapMatchingConfig struct {

	// The field name for the data that describes the identifier representing a
	// collection of GPS points belonging to an individual trace.
	//
	// This member is required.
	IdAttributeName *string

	// The name of the timestamp attribute.
	//
	// This member is required.
	TimestampAttributeName *string

	// The name of the X-attribute
	//
	// This member is required.
	XAttributeName *string

	// The name of the Y-attribute
	//
	// This member is required.
	YAttributeName *string

	noSmithyDocumentSerde
}

// The structure representing Polygon Geometry based on the GeoJson spec (https://www.rfc-editor.org/rfc/rfc7946#section-3.1.6)
// .
type MultiPolygonGeometryInput struct {

	// The coordinates of the multipolygon geometry.
	//
	// This member is required.
	Coordinates [][][][]float64

	noSmithyDocumentSerde
}

// Represents an arithmetic operation to compute spectral index.
type Operation struct {

	// Textual representation of the math operation; Equation used to compute the
	// spectral index.
	//
	// This member is required.
	Equation *string

	// The name of the operation.
	//
	// This member is required.
	Name *string

	// The type of the operation.
	OutputType OutputType

	noSmithyDocumentSerde
}

// A single EarthObservationJob output band.
type OutputBand struct {

	// The name of the band.
	//
	// This member is required.
	BandName *string

	// The datatype of the output band.
	//
	// This member is required.
	OutputDataType OutputType

	noSmithyDocumentSerde
}

// The response structure for an OutputConfig returned by an
// ExportEarthObservationJob.
type OutputConfigInput struct {

	// Path to Amazon S3 storage location for the output configuration file.
	//
	// This member is required.
	S3Data *ExportS3DataInput

	noSmithyDocumentSerde
}

// OutputResolution Configuration indicating the target resolution for the output
// of Resampling operation.
type OutputResolutionResamplingInput struct {

	// User Defined Resolution for the output of Resampling operation defined by value
	// and unit.
	//
	// This member is required.
	UserDefined *UserDefined

	noSmithyDocumentSerde
}

// The input structure representing Output Resolution for Stacking Operation.
type OutputResolutionStackInput struct {

	// A string value representing Predefined Output Resolution for a stacking
	// operation. Allowed values are HIGHEST , LOWEST , and AVERAGE .
	Predefined PredefinedResolution

	// The structure representing User Output Resolution for a Stacking operation
	// defined as a value and unit.
	UserDefined *UserDefined

	noSmithyDocumentSerde
}

// The input structure for specifying Platform. Platform refers to the unique name
// of the specific platform the instrument is attached to. For satellites it is the
// name of the satellite, eg. landsat-8 (Landsat-8), sentinel-2a.
type PlatformInput struct {

	// The value of the platform.
	//
	// This member is required.
	Value *string

	// The ComparisonOperator to use with PlatformInput.
	ComparisonOperator ComparisonOperator

	noSmithyDocumentSerde
}

// The structure representing Polygon Geometry based on the GeoJson spec (https://www.rfc-editor.org/rfc/rfc7946#section-3.1.6)
// .
type PolygonGeometryInput struct {

	// Coordinates representing a Polygon based on the GeoJson spec (https://www.rfc-editor.org/rfc/rfc7946#section-3.1.6)
	// .
	//
	// This member is required.
	Coordinates [][][]float64

	noSmithyDocumentSerde
}

// Properties associated with the Item.
type Properties struct {

	// Estimate of cloud cover.
	EoCloudCover *float32

	// Land cloud cover for Landsat Data Collection.
	LandsatCloudCoverLand *float32

	// Platform property. Platform refers to the unique name of the specific platform
	// the instrument is attached to. For satellites it is the name of the satellite,
	// eg. landsat-8 (Landsat-8), sentinel-2a.
	Platform *string

	// The angle from the sensor between nadir (straight down) and the scene center.
	// Measured in degrees (0-90).
	ViewOffNadir *float32

	// The sun azimuth angle. From the scene center point on the ground, this is the
	// angle between truth north and the sun. Measured clockwise in degrees (0-360).
	ViewSunAzimuth *float32

	// The sun elevation angle. The angle from the tangent of the scene center point
	// to the sun. Measured from the horizon in degrees (-90-90). Negative values
	// indicate the sun is below the horizon, e.g. sun elevation of -10° means the data
	// was captured during nautical twilight (https://www.timeanddate.com/astronomy/different-types-twilight.html)
	// .
	ViewSunElevation *float32

	noSmithyDocumentSerde
}

// Represents a single searchable property to search on.
//
// The following types satisfy this interface:
//
//	PropertyMemberEoCloudCover
//	PropertyMemberLandsatCloudCoverLand
//	PropertyMemberPlatform
//	PropertyMemberViewOffNadir
//	PropertyMemberViewSunAzimuth
//	PropertyMemberViewSunElevation
type Property interface {
	isProperty()
}

// The structure representing EoCloudCover property filter containing a lower
// bound and upper bound.
type PropertyMemberEoCloudCover struct {
	Value EoCloudCoverInput

	noSmithyDocumentSerde
}

func (*PropertyMemberEoCloudCover) isProperty() {}

// The structure representing Land Cloud Cover property filter for Landsat
// collection containing a lower bound and upper bound.
type PropertyMemberLandsatCloudCoverLand struct {
	Value LandsatCloudCoverLandInput

	noSmithyDocumentSerde
}

func (*PropertyMemberLandsatCloudCoverLand) isProperty() {}

// The structure representing Platform property filter consisting of value and
// comparison operator.
type PropertyMemberPlatform struct {
	Value PlatformInput

	noSmithyDocumentSerde
}

func (*PropertyMemberPlatform) isProperty() {}

// The structure representing ViewOffNadir property filter containing a lower
// bound and upper bound.
type PropertyMemberViewOffNadir struct {
	Value ViewOffNadirInput

	noSmithyDocumentSerde
}

func (*PropertyMemberViewOffNadir) isProperty() {}

// The structure representing ViewSunAzimuth property filter containing a lower
// bound and upper bound.
type PropertyMemberViewSunAzimuth struct {
	Value ViewSunAzimuthInput

	noSmithyDocumentSerde
}

func (*PropertyMemberViewSunAzimuth) isProperty() {}

// The structure representing ViewSunElevation property filter containing a lower
// bound and upper bound.
type PropertyMemberViewSunElevation struct {
	Value ViewSunElevationInput

	noSmithyDocumentSerde
}

func (*PropertyMemberViewSunElevation) isProperty() {}

// The structure representing a single PropertyFilter.
type PropertyFilter struct {

	// Represents a single property to match with when searching a raster data
	// collection.
	//
	// This member is required.
	Property Property

	noSmithyDocumentSerde
}

// A list of PropertyFilter objects.
type PropertyFilters struct {

	// The Logical Operator used to combine the Property Filters.
	LogicalOperator LogicalOperator

	// A list of Property Filters.
	Properties []PropertyFilter

	noSmithyDocumentSerde
}

// Response object containing details for a specific RasterDataCollection.
type RasterDataCollectionMetadata struct {

	// The Amazon Resource Name (ARN) of the raster data collection.
	//
	// This member is required.
	Arn *string

	// A description of the raster data collection.
	//
	// This member is required.
	Description *string

	// The name of the raster data collection.
	//
	// This member is required.
	Name *string

	// The list of filters supported by the raster data collection.
	//
	// This member is required.
	SupportedFilters []Filter

	// The type of raster data collection.
	//
	// This member is required.
	Type DataCollectionType

	// The description URL of the raster data collection.
	DescriptionPageUrl *string

	// Each tag consists of a key and a value.
	Tags map[string]string

	noSmithyDocumentSerde
}

// The input structure for Raster Data Collection Query containing the Area of
// Interest, TimeRange Filters, and Property Filters.
type RasterDataCollectionQueryInput struct {

	// The Amazon Resource Name (ARN) of the raster data collection.
	//
	// This member is required.
	RasterDataCollectionArn *string

	// The TimeRange Filter used in the RasterDataCollection Query.
	//
	// This member is required.
	TimeRangeFilter *TimeRangeFilterInput

	// The area of interest being queried for the raster data collection.
	AreaOfInterest AreaOfInterest

	// The list of Property filters used in the Raster Data Collection Query.
	PropertyFilters *PropertyFilters

	noSmithyDocumentSerde
}

// The output structure contains the Raster Data Collection Query input along with
// some additional metadata.
type RasterDataCollectionQueryOutput struct {

	// The ARN of the Raster Data Collection against which the search is done.
	//
	// This member is required.
	RasterDataCollectionArn *string

	// The name of the raster data collection.
	//
	// This member is required.
	RasterDataCollectionName *string

	// The TimeRange filter used in the search.
	//
	// This member is required.
	TimeRangeFilter *TimeRangeFilterOutput

	// The Area of Interest used in the search.
	AreaOfInterest AreaOfInterest

	// Property filters used in the search.
	PropertyFilters *PropertyFilters

	noSmithyDocumentSerde
}

// This is a RasterDataCollectionQueryInput containing AreaOfInterest, Time Range
// filter and Property filters.
type RasterDataCollectionQueryWithBandFilterInput struct {

	// The TimeRange Filter used in the search query.
	//
	// This member is required.
	TimeRangeFilter *TimeRangeFilterInput

	// The Area of interest to be used in the search query.
	AreaOfInterest AreaOfInterest

	// The list of Bands to be displayed in the result for each item.
	BandFilter []string

	// The Property Filters used in the search query.
	PropertyFilters *PropertyFilters

	noSmithyDocumentSerde
}

// The structure representing input for resampling operation.
type ResamplingConfigInput struct {

	// The structure representing output resolution (in target georeferenced units) of
	// the result of resampling operation.
	//
	// This member is required.
	OutputResolution *OutputResolutionResamplingInput

	// The name of the algorithm used for resampling.
	AlgorithmName AlgorithmNameResampling

	// Bands used in the operation. If no target bands are specified, it uses all
	// bands available in the input.
	TargetBands []string

	noSmithyDocumentSerde
}

// The input structure for Reverse Geocoding operation type.
type ReverseGeocodingConfig struct {

	// The field name for the data that describes x-axis coordinate, eg. longitude of
	// a point.
	//
	// This member is required.
	XAttributeName *string

	// The field name for the data that describes y-axis coordinate, eg. latitude of a
	// point.
	//
	// This member is required.
	YAttributeName *string

	noSmithyDocumentSerde
}

// The input structure for Stacking Operation.
type StackConfigInput struct {

	// The structure representing output resolution (in target georeferenced units) of
	// the result of stacking operation.
	OutputResolution *OutputResolutionStackInput

	// A list of bands to be stacked in the specified order. When the parameter is not
	// provided, all the available bands in the data collection are stacked in the
	// alphabetical order of their asset names.
	TargetBands []string

	noSmithyDocumentSerde
}

// The structure representing the configuration for Temporal Statistics operation.
type TemporalStatisticsConfigInput struct {

	// The list of the statistics method options.
	//
	// This member is required.
	Statistics []TemporalStatistics

	// The input for the temporal statistics grouping by time frequency option.
	GroupBy GroupBy

	// The list of target band names for the temporal statistic to calculate.
	TargetBands []string

	noSmithyDocumentSerde
}

// The input for the time-range filter.
type TimeRangeFilterInput struct {

	// The end time for the time-range filter.
	//
	// This member is required.
	EndTime *time.Time

	// The start time for the time-range filter.
	//
	// This member is required.
	StartTime *time.Time

	noSmithyDocumentSerde
}

// The output structure of the time range filter.
type TimeRangeFilterOutput struct {

	// The ending time for the time range filter.
	//
	// This member is required.
	EndTime *time.Time

	// The starting time for the time range filter.
	//
	// This member is required.
	StartTime *time.Time

	noSmithyDocumentSerde
}

// The output resolution (in target georeferenced units) of the result of the
// operation
type UserDefined struct {

	// The units for output resolution of the result.
	//
	// This member is required.
	Unit Unit

	// The value for output resolution of the result.
	//
	// This member is required.
	Value *float32

	noSmithyDocumentSerde
}

// It contains configs such as ReverseGeocodingConfig and MapMatchingConfig.
//
// The following types satisfy this interface:
//
//	VectorEnrichmentJobConfigMemberMapMatchingConfig
//	VectorEnrichmentJobConfigMemberReverseGeocodingConfig
type VectorEnrichmentJobConfig interface {
	isVectorEnrichmentJobConfig()
}

// The input structure for Map Matching operation type.
type VectorEnrichmentJobConfigMemberMapMatchingConfig struct {
	Value MapMatchingConfig

	noSmithyDocumentSerde
}

func (*VectorEnrichmentJobConfigMemberMapMatchingConfig) isVectorEnrichmentJobConfig() {}

// The input structure for Reverse Geocoding operation type.
type VectorEnrichmentJobConfigMemberReverseGeocodingConfig struct {
	Value ReverseGeocodingConfig

	noSmithyDocumentSerde
}

func (*VectorEnrichmentJobConfigMemberReverseGeocodingConfig) isVectorEnrichmentJobConfig() {}

// The input structure for the data source that represents the storage type of the
// input data objects.
//
// The following types satisfy this interface:
//
//	VectorEnrichmentJobDataSourceConfigInputMemberS3Data
type VectorEnrichmentJobDataSourceConfigInput interface {
	isVectorEnrichmentJobDataSourceConfigInput()
}

// The input structure for the Amazon S3 data that represents the Amazon S3
// location of the input data objects.
type VectorEnrichmentJobDataSourceConfigInputMemberS3Data struct {
	Value VectorEnrichmentJobS3Data

	noSmithyDocumentSerde
}

func (*VectorEnrichmentJobDataSourceConfigInputMemberS3Data) isVectorEnrichmentJobDataSourceConfigInput() {
}

// VectorEnrichmentJob error details in response from GetVectorEnrichmentJob.
type VectorEnrichmentJobErrorDetails struct {

	// A message that you define and then is processed and rendered by the Vector
	// Enrichment job when the error occurs.
	ErrorMessage *string

	// The type of error generated during the Vector Enrichment job.
	ErrorType VectorEnrichmentJobErrorType

	noSmithyDocumentSerde
}

// VectorEnrichmentJob export error details in response from
// GetVectorEnrichmentJob.
type VectorEnrichmentJobExportErrorDetails struct {

	// The message providing details about the errors generated during the Vector
	// Enrichment job.
	Message *string

	// The output error details for an Export operation on a Vector Enrichment job.
	Type VectorEnrichmentJobExportErrorType

	noSmithyDocumentSerde
}

// The input structure for the InputConfig in a VectorEnrichmentJob.
type VectorEnrichmentJobInputConfig struct {

	// The input structure for the data source that represents the storage type of the
	// input data objects.
	//
	// This member is required.
	DataSourceConfig VectorEnrichmentJobDataSourceConfigInput

	// The input structure that defines the data source file type.
	//
	// This member is required.
	DocumentType VectorEnrichmentJobDocumentType

	noSmithyDocumentSerde
}

// The Amazon S3 data for the Vector Enrichment job.
type VectorEnrichmentJobS3Data struct {

	// The URL to the Amazon S3 data for the Vector Enrichment job.
	//
	// This member is required.
	S3Uri *string

	// The Key Management Service key ID for server-side encryption.
	KmsKeyId *string

	noSmithyDocumentSerde
}

// The input structure for specifying ViewOffNadir property filter. ViewOffNadir
// refers to the angle from the sensor between nadir (straight down) and the scene
// center. Measured in degrees (0-90).
type ViewOffNadirInput struct {

	// The minimum value for ViewOffNadir property filter. This filters items having
	// ViewOffNadir greater than or equal to this value.
	//
	// This member is required.
	LowerBound *float32

	// The maximum value for ViewOffNadir property filter. This filters items having
	// ViewOffNadir lesser than or equal to this value.
	//
	// This member is required.
	UpperBound *float32

	noSmithyDocumentSerde
}

// The input structure for specifying ViewSunAzimuth property filter.
// ViewSunAzimuth refers to the Sun azimuth angle. From the scene center point on
// the ground, this is the angle between truth north and the sun. Measured
// clockwise in degrees (0-360).
type ViewSunAzimuthInput struct {

	// The minimum value for ViewSunAzimuth property filter. This filters items having
	// ViewSunAzimuth greater than or equal to this value.
	//
	// This member is required.
	LowerBound *float32

	// The maximum value for ViewSunAzimuth property filter. This filters items having
	// ViewSunAzimuth lesser than or equal to this value.
	//
	// This member is required.
	UpperBound *float32

	noSmithyDocumentSerde
}

// The input structure for specifying ViewSunElevation angle property filter.
type ViewSunElevationInput struct {

	// The lower bound to view the sun elevation.
	//
	// This member is required.
	LowerBound *float32

	// The upper bound to view the sun elevation.
	//
	// This member is required.
	UpperBound *float32

	noSmithyDocumentSerde
}

// The structure representing input configuration of ZonalStatistics operation.
type ZonalStatisticsConfigInput struct {

	// List of zonal statistics to compute.
	//
	// This member is required.
	Statistics []ZonalStatistics

	// The Amazon S3 path pointing to the GeoJSON containing the polygonal zones.
	//
	// This member is required.
	ZoneS3Path *string

	// Bands used in the operation. If no target bands are specified, it uses all
	// bands available input.
	TargetBands []string

	// The Amazon Resource Name (ARN) or an ID of a Amazon Web Services Key Management
	// Service (Amazon Web Services KMS) key that Amazon SageMaker uses to decrypt your
	// output artifacts with Amazon S3 server-side encryption. The SageMaker execution
	// role must have kms:GenerateDataKey permission. The KmsKeyId can be any of the
	// following formats:
	//   - // KMS Key ID "1234abcd-12ab-34cd-56ef-1234567890ab"
	//   - // Amazon Resource Name (ARN) of a KMS Key
	//   "arn:aws:kms:<region>:<account>:key/<key-id-12ab-34cd-56ef-1234567890ab>"
	// For more information about key identifiers, see Key identifiers (KeyID) (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id-key-id)
	// in the Amazon Web Services Key Management Service (Amazon Web Services KMS)
	// documentation.
	ZoneS3PathKmsKeyId *string

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

func (*UnknownUnionMember) isAreaOfInterest()                           {}
func (*UnknownUnionMember) isAreaOfInterestGeometry()                   {}
func (*UnknownUnionMember) isJobConfigInput()                           {}
func (*UnknownUnionMember) isProperty()                                 {}
func (*UnknownUnionMember) isVectorEnrichmentJobConfig()                {}
func (*UnknownUnionMember) isVectorEnrichmentJobDataSourceConfigInput() {}