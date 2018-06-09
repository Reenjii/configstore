package configstore

type ErrItemNotFound string
type ErrUninitializedItemList string
type ErrAmbigousItem string
type ErrProvider string

func (e ErrItemNotFound) Error() string {
	return string(e)
}

func (e ErrUninitializedItemList) Error() string {
	return string(e)
}

func (e ErrAmbigousItem) Error() string {
	return string(e)
}

func (e ErrProvider) Error() string {
	return string(e)
}
