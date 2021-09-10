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
	idGen = idGen + 1
	pessoasMap = append(pessoasMap, pessoa)
	return pessoa.Id
}

func Getpessoa(pessoaId int) *models.Pessoa {
	for _, pessoa := range pessoasMap {
		if pessoa.Id == pessoaId {
			return &pessoa
		}
	}
	return nil
}

func UpdatePessoa(pessoaId int, pessoa models.Pessoa) {
	pessoa.Id = pessoaId
	for _index, _pessoa := range pessoasMap {
		if _pessoa.Id == pessoaId {
			pessoasMap[_index] = pessoa
		}
	}
}

func DeletePessoa(pessoaId int) {
	var _pessoasMap []models.Pessoa = []models.Pessoa{}
	for _, pessoa := range pessoasMap {
		if pessoa.Id != pessoaId {
			pessoasMap = append(_pessoasMap, pessoa)
		}
	}
	pessoasMap = _pessoasMap
}

func HasPessoa(pessoaId int) bool {

	for _, pessoa := range pessoasMap {
		if pessoa.Id == pessoaId {
			return true
		}
	}

	return false
}
