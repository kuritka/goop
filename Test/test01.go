package Test


type a1 struct {
  aa int
}


type b1 struct {
	bb int
}


func X() int {
	return new(b1).bb * new(a1).aa
}


