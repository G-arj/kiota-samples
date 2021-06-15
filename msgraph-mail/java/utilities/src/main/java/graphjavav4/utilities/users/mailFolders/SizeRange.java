package graphjavav4.utilities.users.mailFolders;

import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import java.util.function.BiConsumer;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
public class SizeRange implements Parsable {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.  */
    private IDictionary<string, object> _additionalData;
    /** The maximum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.  */
    private Integer _maximumSize;
    /** The minimum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.  */
    private Integer _minimumSize;
    /**
     * Instantiates a new sizeRange and sets the default values.
     * @return a void
     */
    public SizeRange() {
        this._additionalData = new Dictionary<string, object>();
    }
    /**
     * Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @return a IDictionary<string, object>
     */
    @javax.annotation.Nonnull
    public IDictionary<string, object> getAdditionalData() {
        return this._additionalData;
    }
    /**
     * Gets the maximumSize property value. The maximum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
     * @return a integer
     */
    @javax.annotation.Nullable
    public Integer getMaximumSize() {
        return this._maximumSize;
    }
    /**
     * Gets the minimumSize property value. The minimum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
     * @return a integer
     */
    @javax.annotation.Nullable
    public Integer getMinimumSize() {
        return this._minimumSize;
    }
    /**
     * The deserialization information for the current model
     * @return a Map<String, BiConsumer<T, ParseNode>>
     */
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getFieldDeserializers() {
        return new HashMap<>(2) {{
            this.put("maximumSize", (o, n) -> { ((SizeRange)o).setMaximumSize(n.getIntegerValue()); });
            this.put("minimumSize", (o, n) -> { ((SizeRange)o).setMinimumSize(n.getIntegerValue()); });
        }};
    }
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @return a void
     */
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        writer.writeIntegerValue("maximumSize", this.getMaximumSize());
        writer.writeIntegerValue("minimumSize", this.getMinimumSize());
        writer.writeAdditionalData(this.additionalData);
    }
    /**
     * Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @param value Value to set for the AdditionalData property.
     * @return a void
     */
    public void setAdditionalData(@javax.annotation.Nullable final IDictionary<string, object> value) {
        this._additionalData = value;
    }
    /**
     * Sets the maximumSize property value. The maximum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
     * @param value Value to set for the maximumSize property.
     * @return a void
     */
    public void setMaximumSize(@javax.annotation.Nullable final Integer value) {
        this._maximumSize = value;
    }
    /**
     * Sets the minimumSize property value. The minimum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
     * @param value Value to set for the minimumSize property.
     * @return a void
     */
    public void setMinimumSize(@javax.annotation.Nullable final Integer value) {
        this._minimumSize = value;
    }
}
