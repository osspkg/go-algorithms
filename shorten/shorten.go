package shorten

type object struct {
	toStr map[int]string
	toInt map[string]int
	len   int
}

func New(alphabet string) *object {
	v := &object{
		toStr: make(map[int]string),
		toInt: make(map[string]int),
		len:   len(alphabet),
	}

	for i := 0; i < v.len; i++ {
		v.toInt[alphabet[i:i+1]] = i
		v.toStr[i] = alphabet[i : i+1]
	}
	return v
}

func (v *object) Encode(id int) string {
	s := ""
	for id > 0 {
		s = v.toStr[id%v.len] + s
		id /= v.len
	}
	return s
}

func (v *object) Decode(data string) int {
	var id = 0
	for i := 0; i < len(data); i++ {
		id = id*v.len + v.toInt[data[i:i+1]]
	}
	return id
}
