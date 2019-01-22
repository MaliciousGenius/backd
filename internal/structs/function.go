package structs

// Function defines a function to be called from other places API, scheduler, etc
type Function struct {
	ID       string `json:"_id" bson:"_id"`
	Name     string `json:"name" bson:"name"`
	API      bool   `json:"api" bson:"api"`
	RunAs    string `json:"run_as" bson:"run_as"`
	Code     string `json:"code" bson:"code"`
	Metadata `json:"meta" bson:"meta"`
}

// FunctionValidator is the JSON schema validation for the functions collection
func FunctionValidator() map[string]interface{} {

	return BuildValidator(
		map[string]interface{}{
			"_id": map[string]interface{}{
				"bsonType": "string",
				"pattern":  "^[a-zA-Z0-9]{20}$",
			},
			"name": map[string]interface{}{
				"bsonType": "string",
				"pattern":  "^[a-zA-Z0-9]{2,32}$",
			},
			"api": map[string]interface{}{
				"bsonType": "bool",
			},
			"run_as": map[string]interface{}{
				"bsonType": "string",
			},
			"code": map[string]interface{}{
				"bsonType": "string",
			},
		},
		[]string{"_id", "name", "code"},
	)

}

// Indexes
var (
	FunctionIndexes = []Index{
		{
			Fields: []string{"_id"},
			Unique: true,
		},
		{
			Fields: []string{"name"},
			Unique: true,
		},
	}
)
