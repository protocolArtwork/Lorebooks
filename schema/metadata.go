package schema

// this is not a gorm Model, but is used inside of other gorm models
// DO NOT MARK THIS AS A GORM MODEL. IT WILL CREATE A WHOLE NEW TABLE IN THE SQL SCHEMA
type Metadata struct {
	Author     string
	ApprovedBy string
}
