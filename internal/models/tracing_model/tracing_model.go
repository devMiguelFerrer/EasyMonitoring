package tracingmodel

type TracingModel struct {
	Id           string `bson:"-"`
	CreatedAt    string `bson:"createdAt"`
	Url          string `bson:"url"`
	StatusCode   int    `bson:"statusCode"`
	Method       string `bson:"method"`
	RequestBody  string `bson:"requestBody"`
	ResponseBody string `bson:"responseBody"`
	ResponseTime int    `bson:"responseTime"`
}
