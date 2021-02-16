package pier

type Targets []*Target

type Target struct {
	MainDir    string
	Files      []string
	Dockerfile string
	Context    string
}
