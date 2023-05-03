package Vector

type Vector[T any] struct {
	data []T
}

// #region Constructors

func MakeVector[T any](length uint64) Vector[T] {
	return Vector[T]{
		data: make([]T, length),
	}
}

func MakeVector2[T any](length, capacity uint64) Vector[T] {
	return Vector[T]{
		data: make([]T, length, capacity),
	}
}

// #endregion

// #region Setters

func (vector *Vector[T]) SetSlice(data []T) {
	vector.data = data
}

func (vector *Vector[T]) Push(element T) {
	vector.data = append(vector.data, element)
}

func (vector *Vector[T]) Pop() T {
	var element T
	element, vector.data = vector.data[len(vector.data)-1], vector.data[:len(vector.data)-1]
	return element
}

// #endregion

// #region Getters

func (vector *Vector[T]) GetAll() []T {
	return vector.data[:]
}

// #endregion
