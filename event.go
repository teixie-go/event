package event

var (
	_ Event = (*event)(nil)
)

type ListenerFunc func(payload interface{})

// 全局监听方法
var _listeners = make([]ListenerFunc, 0)

type Event interface {
	// 派发事件
	Dispatch(payload interface{})

	// 注册事件监听方法
	Listen(listeners ...ListenerFunc)
}

type event struct {
	listeners []ListenerFunc
}

func (e *event) Dispatch(payload interface{}) {
	// 激活绑定监听方法
	for _, listener := range e.listeners {
		listener(payload)
	}

	// 激活全局监听方法
	for _, listener := range _listeners {
		listener(payload)
	}
}

func (e *event) Listen(listeners ...ListenerFunc) {
	e.listeners = append(e.listeners, listeners...)
}

// 创建实例
func New() *event {
	return &event{listeners: make([]ListenerFunc, 0)}
}

// 注册全局事件监听方法
func Listen(listeners ...ListenerFunc) {
	_listeners = append(_listeners, listeners...)
}
