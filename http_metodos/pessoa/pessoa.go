package pessoa

type Pessoa struct {
	Nome  string `json:"nome"`
	Idade string `json:"idade"`
}

func ListaPessoa() []Pessoa {
	pessoas := []Pessoa{
		{Nome: "Nati, idade: 21"},
		{Nome: "Maria, idade: 25"},
	}

	return pessoas

}
