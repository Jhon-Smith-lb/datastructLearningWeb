package dmresources

type Resources struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func NewResources() *Resources {
	return &Resources{}
}
