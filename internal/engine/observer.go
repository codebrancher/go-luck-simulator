package engine

type ObserverManager[T any] struct {
	observers []Observer[T]
}

func NewObserverManager[T any]() *ObserverManager[T] {
	return &ObserverManager[T]{}
}

func (m *ObserverManager[T]) RegisterObserver(o Observer[T]) {
	m.observers = append(m.observers, o)
}

func (m *ObserverManager[T]) UnregisterObserver(o Observer[T]) {
	for i, observer := range m.observers {
		if observer == o {
			m.observers = append(m.observers[:i], m.observers[i+1:]...)
			break
		}
	}
}

func (m *ObserverManager[T]) NotifyObservers(state T) {
	for _, observer := range m.observers {
		observer.Update(state)
	}
}
