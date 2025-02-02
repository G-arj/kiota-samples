<?php

namespace Microsoft\Graph\Models\Microsoft\Graph;

use Microsoft\Kiota\Abstractions\Serialization\Parsable;
use Microsoft\Kiota\Abstractions\Serialization\ParseNode;
use Microsoft\Kiota\Abstractions\Serialization\SerializationWriter;

class OutlookItem extends Entity implements Parsable 
{
    /** @var array<string>|null $categories The categories associated with the item */
    private ?array $categories;
    
    /** @var string|null $changeKey Identifies the version of the item. Every time the item is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object. Read-only. */
    private ?string $changeKey;
    
    /** @var DateTimeOffset|null $createdDateTime The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z */
    private ?DateTimeOffset $createdDateTime;
    
    /** @var DateTimeOffset|null $lastModifiedDateTime The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z */
    private ?DateTimeOffset $lastModifiedDateTime;
    
    /**
     * Instantiates a new outlookItem and sets the default values.
    */
    public function __construct() {
        parent::__construct();
    }

    /**
     * Gets the categories property value. The categories associated with the item
     * @return array<string>|null
    */
    public function getCategories(): ?array {
        return $this->categories;
    }

    /**
     * Gets the changeKey property value. Identifies the version of the item. Every time the item is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object. Read-only.
     * @return string|null
    */
    public function getChangeKey(): ?string {
        return $this->changeKey;
    }

    /**
     * Gets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
     * @return DateTimeOffset|null
    */
    public function getCreatedDateTime(): ?DateTimeOffset {
        return $this->createdDateTime;
    }

    /**
     * Gets the lastModifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
     * @return DateTimeOffset|null
    */
    public function getLastModifiedDateTime(): ?DateTimeOffset {
        return $this->lastModifiedDateTime;
    }

    /**
     * The deserialization information for the current model
     * @return array<string, callable>
    */
    public function getFieldDeserializers(): array {
        return array_merge(parent::getFieldDeserializers(), [
            'categories' => function (OutlookItem $o, array $n) { $o->setCategories($n); },
            'changeKey' => function (OutlookItem $o, string $n) { $o->setChangeKey($n); },
            'createdDateTime' => function (OutlookItem $o, DateTimeOffset $n) { $o->setCreatedDateTime($n); },
            'lastModifiedDateTime' => function (OutlookItem $o, DateTimeOffset $n) { $o->setLastModifiedDateTime($n); },
        ]);
    }

    /**
     * Serializes information the current object
     * @param SerializationWriter $writer Serialization writer to use to serialize this model
    */
    public function serialize(SerializationWriter $writer): void {
        parent::serialize($writer);
        $writer->writeCollectionOfObjectValues('categories', $this->categories);
        $writer->writeStringValue('changeKey', $this->changeKey);
        $writer->writeObjectValue('createdDateTime', $this->createdDateTime);
        $writer->writeObjectValue('lastModifiedDateTime', $this->lastModifiedDateTime);
    }

    /**
     * Sets the categories property value. The categories associated with the item
     *  @param array<string>|null $value Value to set for the categories property.
    */
    public function setCategories(?array $value ): void {
        $this->categories = $value;
    }

    /**
     * Sets the changeKey property value. Identifies the version of the item. Every time the item is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object. Read-only.
     *  @param string|null $value Value to set for the changeKey property.
    */
    public function setChangeKey(?string $value ): void {
        $this->changeKey = $value;
    }

    /**
     * Sets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
     *  @param DateTimeOffset|null $value Value to set for the createdDateTime property.
    */
    public function setCreatedDateTime(?DateTimeOffset $value ): void {
        $this->createdDateTime = $value;
    }

    /**
     * Sets the lastModifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
     *  @param DateTimeOffset|null $value Value to set for the lastModifiedDateTime property.
    */
    public function setLastModifiedDateTime(?DateTimeOffset $value ): void {
        $this->lastModifiedDateTime = $value;
    }

}
