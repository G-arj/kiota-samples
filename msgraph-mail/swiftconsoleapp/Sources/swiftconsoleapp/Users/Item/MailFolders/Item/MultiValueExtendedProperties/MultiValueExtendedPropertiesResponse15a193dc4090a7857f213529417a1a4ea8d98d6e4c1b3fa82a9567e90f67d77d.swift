import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Users.Item.MailFolders.Item.MultiValueExtendedProperties {
    public class MultiValueExtendedPropertiesResponse : AdditionalDataHolder {
        public var additionalData: [String:Any] = [String:Any]()
        public var nextLink: String?
        public var value: [Swiftconsoleapp.Models.Microsoft.Graph.MultiValueLegacyExtendedProperty]?
    }
}
