// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of supported geographic locations. Countries are listed first,
// and continents are listed last. If Amazon Route 53 supports subdivisions for a
// country (for example, states or provinces), the subdivisions for that country
// are listed in alphabetical order immediately after the corresponding country.
// Route 53 does not perform authorization for this API because it retrieves
// information that is already available to the public. For a list of supported
// geolocation codes, see the GeoLocation (https://docs.aws.amazon.com/Route53/latest/APIReference/API_GeoLocation.html)
// data type.
func (c *Client) ListGeoLocations(ctx context.Context, params *ListGeoLocationsInput, optFns ...func(*Options)) (*ListGeoLocationsOutput, error) {
	if params == nil {
		params = &ListGeoLocationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListGeoLocations", params, optFns, c.addOperationListGeoLocationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListGeoLocationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to get a list of geographic locations that Amazon Route 53 supports
// for geolocation resource record sets.
type ListGeoLocationsInput struct {

	// (Optional) The maximum number of geolocations to be included in the response
	// body for this request. If more than maxitems geolocations remain to be listed,
	// then the value of the IsTruncated element in the response is true .
	MaxItems *int32

	// The code for the continent with which you want to start listing locations that
	// Amazon Route 53 supports for geolocation. If Route 53 has already returned a
	// page or more of results, if IsTruncated is true, and if NextContinentCode from
	// the previous response has a value, enter that value in startcontinentcode to
	// return the next page of results. Include startcontinentcode only if you want to
	// list continents. Don't include startcontinentcode when you're listing countries
	// or countries with their subdivisions.
	StartContinentCode *string

	// The code for the country with which you want to start listing locations that
	// Amazon Route 53 supports for geolocation. If Route 53 has already returned a
	// page or more of results, if IsTruncated is true , and if NextCountryCode from
	// the previous response has a value, enter that value in startcountrycode to
	// return the next page of results.
	StartCountryCode *string

	// The code for the state of the United States with which you want to start
	// listing locations that Amazon Route 53 supports for geolocation. If Route 53 has
	// already returned a page or more of results, if IsTruncated is true , and if
	// NextSubdivisionCode from the previous response has a value, enter that value in
	// startsubdivisioncode to return the next page of results. To list subdivisions
	// (U.S. states), you must include both startcountrycode and startsubdivisioncode .
	StartSubdivisionCode *string

	noSmithyDocumentSerde
}

// A complex type containing the response information for the request.
type ListGeoLocationsOutput struct {

	// A complex type that contains one GeoLocationDetails element for each location
	// that Amazon Route 53 supports for geolocation.
	//
	// This member is required.
	GeoLocationDetailsList []types.GeoLocationDetails

	// A value that indicates whether more locations remain to be listed after the
	// last location in this response. If so, the value of IsTruncated is true . To get
	// more values, submit another request and include the values of NextContinentCode
	// , NextCountryCode , and NextSubdivisionCode in the startcontinentcode ,
	// startcountrycode , and startsubdivisioncode , as applicable.
	//
	// This member is required.
	IsTruncated bool

	// The value that you specified for MaxItems in the request.
	//
	// This member is required.
	MaxItems *int32

	// If IsTruncated is true , you can make a follow-up request to display more
	// locations. Enter the value of NextContinentCode in the startcontinentcode
	// parameter in another ListGeoLocations request.
	NextContinentCode *string

	// If IsTruncated is true , you can make a follow-up request to display more
	// locations. Enter the value of NextCountryCode in the startcountrycode parameter
	// in another ListGeoLocations request.
	NextCountryCode *string

	// If IsTruncated is true , you can make a follow-up request to display more
	// locations. Enter the value of NextSubdivisionCode in the startsubdivisioncode
	// parameter in another ListGeoLocations request.
	NextSubdivisionCode *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListGeoLocationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListGeoLocations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListGeoLocations{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addListGeoLocationsResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListGeoLocations(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opListGeoLocations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "ListGeoLocations",
	}
}

type opListGeoLocationsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opListGeoLocationsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opListGeoLocationsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "route53"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "route53"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("route53")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addListGeoLocationsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opListGeoLocationsResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
