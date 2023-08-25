// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitosync

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/cognitosync/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists datasets for an identity. With Amazon Cognito Sync, each identity has
// access only to its own data. Thus, the credentials used to make this API call
// need to have access to the identity data. ListDatasets can be called with
// temporary user credentials provided by Cognito Identity or with developer
// credentials. You should use the Cognito Identity credentials to make this API
// call. ListDatasets The following examples have been edited for readability. POST
// / HTTP/1.1 CONTENT-TYPE: application/json X-AMZN-REQUESTID:
// 15225768-209f-4078-aaed-7494ace9f2db X-AMZ-TARGET:
// com.amazonaws.cognito.sync.model.AWSCognitoSyncService.ListDatasets HOST:
// cognito-sync.us-east-1.amazonaws.com:443 X-AMZ-DATE: 20141111T215640Z
// AUTHORIZATION: AWS4-HMAC-SHA256 Credential=,
// SignedHeaders=content-type;host;x-amz-date;x-amz-target;x-amzn-requestid,
// Signature= { "Operation": "com.amazonaws.cognito.sync.model#ListDatasets",
// "Service": "com.amazonaws.cognito.sync.model#AWSCognitoSyncService", "Input": {
// "IdentityPoolId": "IDENTITY_POOL_ID", "IdentityId": "IDENTITY_ID", "MaxResults":
// "3" } } 1.1 200 OK x-amzn-requestid: 15225768-209f-4078-aaed-7494ace9f2db,
// 15225768-209f-4078-aaed-7494ace9f2db content-type: application/json
// content-length: 355 date: Tue, 11 Nov 2014 21:56:40 GMT { "Output": { "__type":
// "com.amazonaws.cognito.sync.model#ListDatasetsResponse", "Count": 1, "Datasets":
// [ { "CreationDate": 1.412974057151E9, "DataStorage": 16, "DatasetName":
// "my_list", "IdentityId": "IDENTITY_ID", "LastModifiedBy": "123456789012",
// "LastModifiedDate": 1.412974057244E9, "NumRecords": 1 }], "NextToken": null },
// "Version": "1.0" }
func (c *Client) ListDatasets(ctx context.Context, params *ListDatasetsInput, optFns ...func(*Options)) (*ListDatasetsOutput, error) {
	if params == nil {
		params = &ListDatasetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDatasets", params, optFns, c.addOperationListDatasetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDatasetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request for a list of datasets for an identity.
type ListDatasetsInput struct {

	// A name-spaced GUID (for example,
	// us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE) created by Amazon Cognito. GUID
	// generation is unique within a region.
	//
	// This member is required.
	IdentityId *string

	// A name-spaced GUID (for example,
	// us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE) created by Amazon Cognito. GUID
	// generation is unique within a region.
	//
	// This member is required.
	IdentityPoolId *string

	// The maximum number of results to be returned.
	MaxResults int32

	// A pagination token for obtaining the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

// Returned for a successful ListDatasets request.
type ListDatasetsOutput struct {

	// Number of datasets returned.
	Count int32

	// A set of datasets.
	Datasets []types.Dataset

	// A pagination token for obtaining the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDatasetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDatasets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDatasets{}, middleware.After)
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
	if err = addListDatasetsResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpListDatasetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDatasets(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListDatasets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-sync",
		OperationName: "ListDatasets",
	}
}

type opListDatasetsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opListDatasetsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opListDatasetsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "cognito-sync"
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
				signingName = "cognito-sync"
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
				v4aScheme.SigningName = aws.String("cognito-sync")
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

func addListDatasetsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opListDatasetsResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}