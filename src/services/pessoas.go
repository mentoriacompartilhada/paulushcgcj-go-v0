package services

import (
	"github.com/mentoriacompartilhada/paulushcgcj-go-v0/src/models"
)

var idGen int = 1
var pessoasMap []models.Pessoa = []models.Pessoa{}

func ListPessoas(page int, size int) models.Page {

	_offset := page * size
	_offsetEnd := _offset + size

	var pageItems []models.Pessoa = []models.Pessoa{}

	for index, pessoa := range pessoasMap {
		if index >= _offset && index < _offsetEnd {
			pageItems = append(pageItems, pessoa)
		}
	}

	return models.Page{Page: page, Size: len(pageItems), Total: len(pessoasMap), Items: pageItems}
}

func AddPessoa(pessoa models.Pessoa) int {
	pessoa.Id = idGen
	idGen++
	pessoasMap = append(pessoasMap, pessoa)
	return pessoa.Id
}

func Getpessoa(pessoaId int) models.Pessoa {
	return pessoasMap[pessoaId-1]
}

func UpdatePessoa(pessoaId int, pessoa models.Pessoa) {
	pessoa.Id = pessoaId
	pessoasMap[pessoaId-1] = pessoa
}

func DeletePessoa(pessoaId int) {
	pessoasMap = append(pessoasMap[:pessoaId-1], pessoasMap[pessoaId:]...)
}

func HasPessoa(pessoaId int) bool {

	for _, pessoa := range pessoasMap {
		if pessoa.Id == pessoaId {
			return true
		}
	}

	return false
}
