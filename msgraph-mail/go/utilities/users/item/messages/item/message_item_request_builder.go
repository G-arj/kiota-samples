package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models/microsoft/graph"
    i10075314d3b96cd3f0f2195562b0027395b41f6069ced9cc8d93624bbc206528 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/messages/item/extensions"
    i170c26461395ae79c8485bb09f6acf415c54c121851bca73450cfff498f75eb8 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/messages/item/singlevalueextendedproperties"
    i228962c2a18462bcfbde8557f6668f3fd7d4e68d5c7f3f1f3b675bc816866033 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/messages/item/attachments"
    i9550b442511c46b7a974613d5ee3238fb100cef30fff0520a9682fb3d38ad326 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/messages/item/value"
    iadaad8f92ba85254c1d639d8a8cb86def5e843af4bc1062bc836f08d062196d6 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/messages/item/multivalueextendedproperties"
    i361d73846d6331cbbe8151dfaef3455f34e63dfa07626b67576681683b148587 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/messages/item/extensions/item"
    i9f73052b0f535426cd146e9403e4e088995d624a331ea8cf19005e2c6ff46d7f "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/messages/item/singlevalueextendedproperties/item"
    idd4800ce8115a136eb3f42205c28105422c049d9e5b23796f3bdf18e2778f6d6 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/messages/item/attachments/item"
    if06dd322de6c7a1abe876d18d2ddc86a4b8bcf49f2fea78dcba12aff7dcb4994 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/messages/item/multivalueextendedproperties/item"
)

// MessageItemRequestBuilder builds and executes requests for operations under \users\{user-id}\messages\{message-id}
type MessageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// MessageItemRequestBuilderDeleteOptions options for Delete
type MessageItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MessageItemRequestBuilderGetOptions options for Get
type MessageItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *MessageItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MessageItemRequestBuilderGetQueryParameters the messages in a mailbox or folder. Read-only. Nullable.
type MessageItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string;
}
// MessageItemRequestBuilderPatchOptions options for Patch
type MessageItemRequestBuilderPatchOptions struct {
    // 
    Body i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.Messageable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *MessageItemRequestBuilder) Attachments()(*i228962c2a18462bcfbde8557f6668f3fd7d4e68d5c7f3f1f3b675bc816866033.AttachmentsRequestBuilder) {
    return i228962c2a18462bcfbde8557f6668f3fd7d4e68d5c7f3f1f3b675bc816866033.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/.users.item.messages.item.attachments.item collection
func (m *MessageItemRequestBuilder) AttachmentsById(id string)(*idd4800ce8115a136eb3f42205c28105422c049d9e5b23796f3bdf18e2778f6d6.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return idd4800ce8115a136eb3f42205c28105422c049d9e5b23796f3bdf18e2778f6d6.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewMessageItemRequestBuilderInternal instantiates a new MessageItemRequestBuilder and sets the default values.
func NewMessageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MessageItemRequestBuilder) {
    m := &MessageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/messages/{message_id}{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMessageItemRequestBuilder instantiates a new MessageItemRequestBuilder and sets the default values.
func NewMessageItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MessageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMessageItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *MessageItemRequestBuilder) Content()(*i9550b442511c46b7a974613d5ee3238fb100cef30fff0520a9682fb3d38ad326.ContentRequestBuilder) {
    return i9550b442511c46b7a974613d5ee3238fb100cef30fff0520a9682fb3d38ad326.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation the messages in a mailbox or folder. Read-only. Nullable.
func (m *MessageItemRequestBuilder) CreateDeleteRequestInformation(options *MessageItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the messages in a mailbox or folder. Read-only. Nullable.
func (m *MessageItemRequestBuilder) CreateGetRequestInformation(options *MessageItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation the messages in a mailbox or folder. Read-only. Nullable.
func (m *MessageItemRequestBuilder) CreatePatchRequestInformation(options *MessageItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete the messages in a mailbox or folder. Read-only. Nullable.
func (m *MessageItemRequestBuilder) Delete(options *MessageItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *MessageItemRequestBuilder) Extensions()(*i10075314d3b96cd3f0f2195562b0027395b41f6069ced9cc8d93624bbc206528.ExtensionsRequestBuilder) {
    return i10075314d3b96cd3f0f2195562b0027395b41f6069ced9cc8d93624bbc206528.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/.users.item.messages.item.extensions.item collection
func (m *MessageItemRequestBuilder) ExtensionsById(id string)(*i361d73846d6331cbbe8151dfaef3455f34e63dfa07626b67576681683b148587.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i361d73846d6331cbbe8151dfaef3455f34e63dfa07626b67576681683b148587.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the messages in a mailbox or folder. Read-only. Nullable.
func (m *MessageItemRequestBuilder) Get(options *MessageItemRequestBuilderGetOptions)(i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.Messageable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.CreateMessageFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.Messageable), nil
}
func (m *MessageItemRequestBuilder) MultiValueExtendedProperties()(*iadaad8f92ba85254c1d639d8a8cb86def5e843af4bc1062bc836f08d062196d6.MultiValueExtendedPropertiesRequestBuilder) {
    return iadaad8f92ba85254c1d639d8a8cb86def5e843af4bc1062bc836f08d062196d6.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/.users.item.messages.item.multiValueExtendedProperties.item collection
func (m *MessageItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*if06dd322de6c7a1abe876d18d2ddc86a4b8bcf49f2fea78dcba12aff7dcb4994.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return if06dd322de6c7a1abe876d18d2ddc86a4b8bcf49f2fea78dcba12aff7dcb4994.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the messages in a mailbox or folder. Read-only. Nullable.
func (m *MessageItemRequestBuilder) Patch(options *MessageItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *MessageItemRequestBuilder) SingleValueExtendedProperties()(*i170c26461395ae79c8485bb09f6acf415c54c121851bca73450cfff498f75eb8.SingleValueExtendedPropertiesRequestBuilder) {
    return i170c26461395ae79c8485bb09f6acf415c54c121851bca73450cfff498f75eb8.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/.users.item.messages.item.singleValueExtendedProperties.item collection
func (m *MessageItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i9f73052b0f535426cd146e9403e4e088995d624a331ea8cf19005e2c6ff46d7f.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i9f73052b0f535426cd146e9403e4e088995d624a331ea8cf19005e2c6ff46d7f.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
