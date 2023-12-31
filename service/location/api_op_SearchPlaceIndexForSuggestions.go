// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/location/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Generates suggestions for addresses and points of interest based on partial or
// misspelled free-form text. This operation is also known as autocomplete,
// autosuggest, or fuzzy matching. Optional parameters let you narrow your search
// results by bounding box or country, or bias your search toward a specific
// position on the globe. You can search for suggested place names near a specified
// position by using BiasPosition , or filter results within a bounding box by
// using FilterBBox . These parameters are mutually exclusive; using both
// BiasPosition and FilterBBox in the same command returns an error.
func (c *Client) SearchPlaceIndexForSuggestions(ctx context.Context, params *SearchPlaceIndexForSuggestionsInput, optFns ...func(*Options)) (*SearchPlaceIndexForSuggestionsOutput, error) {
	if params == nil {
		params = &SearchPlaceIndexForSuggestionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchPlaceIndexForSuggestions", params, optFns, c.addOperationSearchPlaceIndexForSuggestionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchPlaceIndexForSuggestionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchPlaceIndexForSuggestionsInput struct {

	// The name of the place index resource you want to use for the search.
	//
	// This member is required.
	IndexName *string

	// The free-form partial text to use to generate place suggestions. For example,
	// eiffel tow .
	//
	// This member is required.
	Text *string

	// An optional parameter that indicates a preference for place suggestions that
	// are closer to a specified position. If provided, this parameter must contain a
	// pair of numbers. The first number represents the X coordinate, or longitude; the
	// second number represents the Y coordinate, or latitude. For example,
	// [-123.1174, 49.2847] represents the position with longitude -123.1174 and
	// latitude 49.2847 . BiasPosition and FilterBBox are mutually exclusive.
	// Specifying both options results in an error.
	BiasPosition []float64

	// An optional parameter that limits the search results by returning only
	// suggestions within a specified bounding box. If provided, this parameter must
	// contain a total of four consecutive numbers in two pairs. The first pair of
	// numbers represents the X and Y coordinates (longitude and latitude,
	// respectively) of the southwest corner of the bounding box; the second pair of
	// numbers represents the X and Y coordinates (longitude and latitude,
	// respectively) of the northeast corner of the bounding box. For example,
	// [-12.7935, -37.4835, -12.0684, -36.9542] represents a bounding box where the
	// southwest corner has longitude -12.7935 and latitude -37.4835 , and the
	// northeast corner has longitude -12.0684 and latitude -36.9542 . FilterBBox and
	// BiasPosition are mutually exclusive. Specifying both options results in an error.
	FilterBBox []float64

	// A list of one or more Amazon Location categories to filter the returned places.
	// If you include more than one category, the results will include results that
	// match any of the categories listed. For more information about using categories,
	// including a list of Amazon Location categories, see Categories and filtering (https://docs.aws.amazon.com/location/latest/developerguide/category-filtering.html)
	// , in the Amazon Location Service Developer Guide.
	FilterCategories []string

	// An optional parameter that limits the search results by returning only
	// suggestions within the provided list of countries.
	//   - Use the ISO 3166 (https://www.iso.org/iso-3166-country-codes.html) 3-digit
	//   country code. For example, Australia uses three upper-case characters: AUS .
	FilterCountries []string

	// The optional API key (https://docs.aws.amazon.com/location/latest/developerguide/using-apikeys.html)
	// to authorize the request.
	Key *string

	// The preferred language used to return results. The value must be a valid BCP 47 (https://tools.ietf.org/search/bcp47)
	// language tag, for example, en for English. This setting affects the languages
	// used in the results. If no language is specified, or not supported for a
	// particular result, the partner automatically chooses a language for the result.
	// For an example, we'll use the Greek language. You search for Athens, Gr to get
	// suggestions with the language parameter set to en . The results found will most
	// likely be returned as Athens, Greece . If you set the language parameter to el ,
	// for Greek, then the result found will more likely be returned as Αθήνα, Ελλάδα .
	// If the data provider does not have a value for Greek, the result will be in a
	// language that the provider does support.
	Language *string

	// An optional parameter. The maximum number of results returned per request. The
	// default: 5
	MaxResults *int32

	noSmithyDocumentSerde
}

type SearchPlaceIndexForSuggestionsOutput struct {

	// A list of place suggestions that best match the search text.
	//
	// This member is required.
	Results []types.SearchForSuggestionsResult

	// Contains a summary of the request. Echoes the input values for BiasPosition ,
	// FilterBBox , FilterCountries , Language , MaxResults , and Text . Also includes
	// the DataSource of the place index.
	//
	// This member is required.
	Summary *types.SearchPlaceIndexForSuggestionsSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchPlaceIndexForSuggestionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchPlaceIndexForSuggestions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchPlaceIndexForSuggestions{}, middleware.After)
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
	if err = addEndpointPrefix_opSearchPlaceIndexForSuggestionsMiddleware(stack); err != nil {
		return err
	}
	if err = addSearchPlaceIndexForSuggestionsResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpSearchPlaceIndexForSuggestionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchPlaceIndexForSuggestions(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opSearchPlaceIndexForSuggestionsMiddleware struct {
}

func (*endpointPrefix_opSearchPlaceIndexForSuggestionsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opSearchPlaceIndexForSuggestionsMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "places." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opSearchPlaceIndexForSuggestionsMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opSearchPlaceIndexForSuggestionsMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opSearchPlaceIndexForSuggestions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "geo",
		OperationName: "SearchPlaceIndexForSuggestions",
	}
}

type opSearchPlaceIndexForSuggestionsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opSearchPlaceIndexForSuggestionsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opSearchPlaceIndexForSuggestionsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "geo"
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
				signingName = "geo"
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
				v4aScheme.SigningName = aws.String("geo")
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

func addSearchPlaceIndexForSuggestionsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opSearchPlaceIndexForSuggestionsResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
