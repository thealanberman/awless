package ast

type Entity string

var entities = map[Entity]struct{}{
	"none": {},

	"accesskey":           {},
	"alarm":               {},
	"appscalingtarget":    {},
	"appscalingpolicy":    {},
	"scalinggroup":        {},
	"bucket":              {},
	"container":           {},
	"containercluster":    {},
	"containerservice":    {},
	"containertask":       {},
	"database":            {},
	"distribution":        {},
	"dbsubnetgroup":       {},
	"elasticip":           {},
	"function":            {},
	"group":               {},
	"instance":            {},
	"image":               {},
	"internetgateway":     {},
	"natgateway":          {},
	"instanceprofile":     {},
	"keypair":             {},
	"launchconfiguration": {},
	"listener":            {},
	"loadbalancer":        {},
	"loginprofile":        {},
	"policy":              {},
	"queue":               {},
	"record":              {},
	"registry":            {},
	"repository":          {},
	"role":                {},
	"route":               {},
	"routetable":          {},
	"s3object":            {},
	"scalingpolicy":       {},
	"securitygroup":       {},
	"snapshot":            {},
	"stack":               {},
	"subnet":              {},
	"subscription":        {},
	"tag":                 {},
	"targetgroup":         {},
	"topic":               {},
	"user":                {},
	"volume":              {},
	"vpc":                 {},
	"zone":                {},
}

func IsInvalidEntity(s string) bool {
	_, ok := entities[Entity(s)]
	return !ok
}
