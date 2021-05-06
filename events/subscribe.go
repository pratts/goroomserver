package events

type SubscribeEvent struct {
	name string
}

func (e *SubscribeEvent) getName() string {
	return e.name
}

func (e *SubscribeEvent) handle(params map[string]interface{}) {

}
