import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
import {MailFolder} from '../mailFolder';

export class ChildFoldersResponse implements Parsable<ChildFoldersResponse> {
    public value?: MailFolder[] | undefined;
    public nextLink?: string | undefined;
    public serialize (writer: SerializationWriter) : void {
        writer.writeCollectionOfObjectValues<MailFolder>("value", this.value);
        writer.writeStringValue("nextLink", this.nextLink);
    };
    public deserializeFields () : Map<string, (item: ChildFoldersResponse, node: ParseNode) => void> {
        return new Map<string, (item: ChildFoldersResponse, node: ParseNode) => void>([
            ["value", (o, n) => { o.value = n.getCollectionOfObjectValues<MailFolder>(MailFolder); }],
            ["nextLink", (o, n) => { o.nextLink = n.getStringValue(); }],
        ]);
    };
}
