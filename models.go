package main

type asset struct {
	ResourceType string   `json:"resourceType"`
	Content      string   `json:"content"`
	ContentType  string   `json:"contentType"`
	Metadata     metadata `json:"metadata"`
}

type metadata struct {
	ApplicationID string `json:"application_id"`
}

type productValidation struct {
	Action  string  `json:"action"`
	Product product `json:"product"`
}

type product struct {
	Version                          string                      `json:"version"`
	LifecycleStatus                  string                      `json:"lifecycleStatus"`
	IsBundle                         bool                        `json:"isBundle"`
	BundledProductSpecification      []string                    `json:"bundledProductSpecification"`
	ProductSpecCharacteristic        []productSpecCharacteristic `json:"productSpecCharacteristic"`
	ProductSpecificationRelationship []string                    `json:"productSpecificationRelationship"`
	Attachment                       []attachment                `json:"attachment"`
	RelatedParty                     []relatedParty              `json:"relatedParty"`
	Name                             string                      `json:"name"`
	Brand                            string                      `json:"brand"`
	ProductNumber                    string                      `json:"productNumber"`
}

type productSpecCharacteristic struct {
	Name                           string                           `json:"name"`
	Description                    string                           `json:"description"`
	ValueType                      string                           `json:"valueType"`
	Configurable                   bool                             `json:"configurable"`
	ProductSpecCharacteristicValue []productSpecCharacteristicValue `json:"productSpecCharacteristicValue"`
}

type productSpecCharacteristicValue struct {
	Default       bool   `json:"default"`
	UnitOfMeasure string `json:"unitOfMeasure"`
	Value         string `json:"value"`
	ValueFrom     string `json:"valueFrom"`
	ValueTo       string `json:"valueTo"`
}

type attachment struct {
	Type string `json:"type"`
}

type relatedParty struct {
	ID   string `json:"id"`
	Href string `json:"href"`
	Role string `json:"role"`
}
