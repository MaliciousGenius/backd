package structs

// RBAC is the struct that defines how the role permissions are set on the db
type RBAC struct {
	ID           string `json:"_id" bson:"_id"`
	DomainID     string `json:"domain_id" bson:"domain_id"`
	IdentityID   string `json:"identity_id" bson:"identity_id"`
	Collection   string `json:"collection" bson:"collection"`
	CollectionID string `json:"collection_id" bson:"collection_id"`
	Permission   string `json:"perm" bson:"perm"`
}

// RBACValidator is the JSON schema validation for the domains collection
func RBACValidator() map[string]interface{} {

	return map[string]interface{}{
		"$jsonSchema": map[string]interface{}{
			"bsonType": "object",
			"required": []string{"_id", "domain_id", "identity_id", "collection", "collection_id", "perm"},
			"properties": map[string]interface{}{
				"_id": map[string]interface{}{
					"bsonType": "string",
					"pattern":  "^[a-zA-Z0-9]{20}$",
				},
				"domain_id": map[string]interface{}{
					"bsonType":  "string",
					"pattern":   "[a-zA-Z0-9]+",
					"maxLength": 32,
				},
				"identity_id": map[string]interface{}{
					"bsonType": "string",
					"pattern":  "^[a-zA-Z0-9]{20}$",
				},
				"collection": map[string]interface{}{
					"bsonType":  "string",
					"pattern":   "[a-zA-Z0-9*]+",
					"maxLength": 32,
				},
				"collection_id": map[string]interface{}{
					"bsonType":  "string",
					"pattern":   "[a-zA-Z0-9*]+",
					"maxLength": 20,
				},
				"perm": map[string]interface{}{
					"bsonType":  "string",
					"pattern":   "[a-zA-Z0-9]+",
					"maxLength": 32,
				},
			},
		},
	}

}

// Indexes
var (
	RBACIndexes = []Index{
		{
			Fields: []string{"_id"},
			Unique: true,
		},
		{
			Fields: []string{"domain_id", "identity_id", "collection", "collection_id", "perm"}, // can queries
			Unique: false,
		},
		{
			Fields: []string{"domain_id", "identity_id", "collection", "perm"}, // visibleIDs queries
			Unique: false,
		},
	}
)
