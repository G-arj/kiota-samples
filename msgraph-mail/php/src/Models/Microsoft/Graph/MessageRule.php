<?php

namespace Microsoft\Graph\Models\Microsoft\Graph;

use Microsoft\Kiota\Abstractions\Serialization\Parsable;
use Microsoft\Kiota\Abstractions\Serialization\ParseNode;
use Microsoft\Kiota\Abstractions\Serialization\SerializationWriter;

class MessageRule extends Entity implements Parsable 
{
    /** @var MessageRuleActions|null $actions  */
    private ?MessageRuleActions $actions;
    
    /** @var MessageRulePredicates|null $conditions  */
    private ?MessageRulePredicates $conditions;
    
    /** @var string|null $displayName The display name of the rule. */
    private ?string $displayName;
    
    /** @var MessageRulePredicates|null $exceptions  */
    private ?MessageRulePredicates $exceptions;
    
    /** @var bool|null $hasError Indicates whether the rule is in an error condition. Read-only. */
    private ?bool $hasError;
    
    /** @var bool|null $isEnabled Indicates whether the rule is enabled to be applied to messages. */
    private ?bool $isEnabled;
    
    /** @var bool|null $isReadOnly Indicates if the rule is read-only and cannot be modified or deleted by the rules REST API. */
    private ?bool $isReadOnly;
    
    /** @var int|null $sequence Indicates the order in which the rule is executed, among other rules. */
    private ?int $sequence;
    
    /**
     * Instantiates a new messageRule and sets the default values.
    */
    public function __construct() {
        parent::__construct();
    }

    /**
     * Gets the actions property value. 
     * @return MessageRuleActions|null
    */
    public function getActions(): ?MessageRuleActions {
        return $this->actions;
    }

    /**
     * Gets the conditions property value. 
     * @return MessageRulePredicates|null
    */
    public function getConditions(): ?MessageRulePredicates {
        return $this->conditions;
    }

    /**
     * Gets the displayName property value. The display name of the rule.
     * @return string|null
    */
    public function getDisplayName(): ?string {
        return $this->displayName;
    }

    /**
     * Gets the exceptions property value. 
     * @return MessageRulePredicates|null
    */
    public function getExceptions(): ?MessageRulePredicates {
        return $this->exceptions;
    }

    /**
     * Gets the hasError property value. Indicates whether the rule is in an error condition. Read-only.
     * @return bool|null
    */
    public function getHasError(): ?bool {
        return $this->hasError;
    }

    /**
     * Gets the isEnabled property value. Indicates whether the rule is enabled to be applied to messages.
     * @return bool|null
    */
    public function getIsEnabled(): ?bool {
        return $this->isEnabled;
    }

    /**
     * Gets the isReadOnly property value. Indicates if the rule is read-only and cannot be modified or deleted by the rules REST API.
     * @return bool|null
    */
    public function getIsReadOnly(): ?bool {
        return $this->isReadOnly;
    }

    /**
     * Gets the sequence property value. Indicates the order in which the rule is executed, among other rules.
     * @return int|null
    */
    public function getSequence(): ?int {
        return $this->sequence;
    }

    /**
     * The deserialization information for the current model
     * @return array<string, callable>
    */
    public function getFieldDeserializers(): array {
        return array_merge(parent::getFieldDeserializers(), [
            'actions' => function (MessageRule $o, MessageRuleActions $n) { $o->setActions($n); },
            'conditions' => function (MessageRule $o, MessageRulePredicates $n) { $o->setConditions($n); },
            'displayName' => function (MessageRule $o, string $n) { $o->setDisplayName($n); },
            'exceptions' => function (MessageRule $o, MessageRulePredicates $n) { $o->setExceptions($n); },
            'hasError' => function (MessageRule $o, bool $n) { $o->setHasError($n); },
            'isEnabled' => function (MessageRule $o, bool $n) { $o->setIsEnabled($n); },
            'isReadOnly' => function (MessageRule $o, bool $n) { $o->setIsReadOnly($n); },
            'sequence' => function (MessageRule $o, int $n) { $o->setSequence($n); },
        ]);
    }

    /**
     * Serializes information the current object
     * @param SerializationWriter $writer Serialization writer to use to serialize this model
    */
    public function serialize(SerializationWriter $writer): void {
        parent::serialize($writer);
        $writer->writeObjectValue('actions', $this->actions);
        $writer->writeObjectValue('conditions', $this->conditions);
        $writer->writeStringValue('displayName', $this->displayName);
        $writer->writeObjectValue('exceptions', $this->exceptions);
        $writer->writeBooleanValue('hasError', $this->hasError);
        $writer->writeBooleanValue('isEnabled', $this->isEnabled);
        $writer->writeBooleanValue('isReadOnly', $this->isReadOnly);
        $writer->writeObjectValue('sequence', $this->sequence);
    }

    /**
     * Sets the actions property value. 
     *  @param MessageRuleActions|null $value Value to set for the actions property.
    */
    public function setActions(?MessageRuleActions $value ): void {
        $this->actions = $value;
    }

    /**
     * Sets the conditions property value. 
     *  @param MessageRulePredicates|null $value Value to set for the conditions property.
    */
    public function setConditions(?MessageRulePredicates $value ): void {
        $this->conditions = $value;
    }

    /**
     * Sets the displayName property value. The display name of the rule.
     *  @param string|null $value Value to set for the displayName property.
    */
    public function setDisplayName(?string $value ): void {
        $this->displayName = $value;
    }

    /**
     * Sets the exceptions property value. 
     *  @param MessageRulePredicates|null $value Value to set for the exceptions property.
    */
    public function setExceptions(?MessageRulePredicates $value ): void {
        $this->exceptions = $value;
    }

    /**
     * Sets the hasError property value. Indicates whether the rule is in an error condition. Read-only.
     *  @param bool|null $value Value to set for the hasError property.
    */
    public function setHasError(?bool $value ): void {
        $this->hasError = $value;
    }

    /**
     * Sets the isEnabled property value. Indicates whether the rule is enabled to be applied to messages.
     *  @param bool|null $value Value to set for the isEnabled property.
    */
    public function setIsEnabled(?bool $value ): void {
        $this->isEnabled = $value;
    }

    /**
     * Sets the isReadOnly property value. Indicates if the rule is read-only and cannot be modified or deleted by the rules REST API.
     *  @param bool|null $value Value to set for the isReadOnly property.
    */
    public function setIsReadOnly(?bool $value ): void {
        $this->isReadOnly = $value;
    }

    /**
     * Sets the sequence property value. Indicates the order in which the rule is executed, among other rules.
     *  @param int|null $value Value to set for the sequence property.
    */
    public function setSequence(?int $value ): void {
        $this->sequence = $value;
    }

}
