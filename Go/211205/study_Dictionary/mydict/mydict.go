package mydict

import "errors"

//bankAccount에서 struct 사용한 것과 달리 type 사용

//Dictionary type
type Dictionary map[string]string

//   =========
//     별명

var (
	errNotFound  = errors.New("해당 단어를 찾을 수 없습니다.")
	errCanUpdate = errors.New("존재하지 않는 단어는 업데이트 불가")
	errWordExist = errors.New("해당 단어가 이미 존재합니다.")
)

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

//Add a word
func (d Dictionary) Add(word, def string) error {
	//에러만 보면 되기 떄문에 definition _처리
	_, err := d.Search(word)
	if err == errNotFound {
		d[word] = def
	} else if err == nil {
		return errWordExist
	}
	return nil
}

//Update a word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCanUpdate
	}
	return nil
}

//Delete a word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
