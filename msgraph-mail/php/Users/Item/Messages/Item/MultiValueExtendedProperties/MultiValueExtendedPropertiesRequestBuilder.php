<?php

namespace Microsoft\Graph\Users\Item\Messages\Item\MultiValueExtendedProperties;

use Microsoft\Graph\Models\Microsoft\Graph\MultiValueLegacyExtendedProperty;
use Microsoft\Graph\Users\Item\Messages\Item\MultiValueExtendedProperties\Item\MultiValueLegacyExtendedPropertyRequestBuilder;
use Microsoft\Kiota\Abstractions\HttpMethod;
use Microsoft\Kiota\Abstractions\MiddlewareOption;
use Microsoft\Kiota\Abstractions\RequestAdapter;
use Microsoft\Kiota\Abstractions\RequestInformation;
use Microsoft\Kiota\Abstractions\ResponseHandler;
use Microsoft\Kiota\Abstractions\Serialization\Parsable;

class MultiValueExtendedPropertiesRequestBuilder 
{
    /** @var string $currentPath Current path for the request */
    private string $currentPath;
    
    /** @var bool $isRawUrl Whether the current path is a raw URL */
    private bool $isRawUrl;
    
    /** @var string $pathSegment Path segment to use to build the URL for the current request builder */
    private string $pathSegment;
    
    private RequestAdapter $requestAdapter;
    
    /**
     * Instantiates a new MultiValueExtendedPropertiesRequestBuilder and sets the default values.
     * @param string $currentPath Current path for the request
     * @param bool|null $isRawUrl Whether the current path is a raw URL
     * @param RequestAdapter $requestAdapter The http core service to use to execute the requests.
    */
    public function __construct(string $currentPath, ?bool $isRawUrl, RequestAdapter $requestAdapter) {
        if (is_null($currentPath)) {
            throw new \Exception('$currentPath cannot be null');
        }
        if (is_null($requestAdapter)) {
            throw new \Exception('$requestAdapter cannot be null');
        }
        $this->pathSegment = '/multiValueExtendedProperties';
        $this->requestAdapter = $requestAdapter;
        $this->currentPath = $currentPath;
    }

    /**
     * The collection of multi-value extended properties defined for the message. Nullable.
     * @param GetQueryParameters|null $queryParameters Request query parameters
     * @param array|null $headers Request headers
     * @param array|null $options Request options
     * @return RequestInformation
    */
    public function createGetRequestInformation(?GetQueryParameters $queryParameters, ?array $headers, ?array $options): RequestInformation {
        $requestInfo = new RequestInformation();
        $requestInfo->setUri($this->currentPath, $this->pathSegment, $this->isRawUrl);
        $requestInfo->httpMethod = HttpMethod::GET;
        $requestInfo->setHeadersFromRawObject($headers);
        $requestInfo->setQueryStringParametersFromRawObject($queryStringParams);
        $requestInfo->addRequestOptions(...$options);
        return $requestInfo;
    }

    /**
     * The collection of multi-value extended properties defined for the message. Nullable.
     * @param MultiValueLegacyExtendedProperty $body 
     * @param array|null $headers Request headers
     * @param array|null $options Request options
     * @return RequestInformation
    */
    public function createPostRequestInformation(MultiValueLegacyExtendedProperty $body, ?array $headers, ?array $options): RequestInformation {
        if (is_null($body)) {
            throw new \Exception('$body cannot be null');
        }
        $requestInfo = new RequestInformation();
        $requestInfo->setUri($this->currentPath, $this->pathSegment, $this->isRawUrl);
        $requestInfo->httpMethod = HttpMethod::POST;
        $requestInfo->setHeadersFromRawObject($headers);
        $requestInfo->setContentFromParsable($this->requestAdapter, "application/json", $body);
        $requestInfo->addRequestOptions(...$options);
        return $requestInfo;
    }

    /**
     * The collection of multi-value extended properties defined for the message. Nullable.
     * @param GetQueryParameters|null $queryParameters Request query parameters
     * @param array|null $headers Request headers
     * @param array|null $options Request options
     * @param ResponseHandler|null $responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return MultiValueExtendedPropertiesResponse|null
    */
    public function get(?GetQueryParameters $queryParameters, ?array $headers, ?array $options, ?ResponseHandler $responseHandler): ?MultiValueExtendedPropertiesResponse {
    }

    /**
     * The collection of multi-value extended properties defined for the message. Nullable.
     * @param MultiValueLegacyExtendedProperty $body 
     * @param array|null $headers Request headers
     * @param array|null $options Request options
     * @param ResponseHandler|null $responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return MultiValueLegacyExtendedProperty|null
    */
    public function post(MultiValueLegacyExtendedProperty $body, ?array $headers, ?array $options, ?ResponseHandler $responseHandler): ?MultiValueLegacyExtendedProperty {
        if (is_null($body)) {
            throw new \Exception('$body cannot be null');
        }
    }

}
