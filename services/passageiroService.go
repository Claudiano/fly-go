package services

import (
	"fly/dtos"
	"fly/models"
	"fly/repositories"
	"fmt"
)

// Repository responsavel pela conexão com o banco de dados
var passageiroRepository = repositories.PassageiroRepository{}

type PassageiroService struct {
}

func (PassageiroService) CadastrarPassageiro(passageiro models.Passageiro) {
	passageiroRepository.Save(passageiro)
}

func (PassageiroService) BuscarPassageiros() []models.Passageiro {
	fmt.Println("buscar passageiros")
	passageiros := passageiroRepository.FindByAll()
	return passageiros
}

func (PassageiroService) BuscarPassageiro(idPassageiro uint64) models.Passageiro {
	return passageiroRepository.FindById(idPassageiro)
}

func (PassageiroService) AtualizarPassageiro(passageiro models.Passageiro) {
	passageiroRepository.Update(passageiro)
}

func (PassageiroService) ExcluirPassageiro(passageiro models.Passageiro) {
	passageiroRepository.Delete(passageiro)
}

func (PassageiroService) BuscarPassageiroLogin(passageiroDto dtos.PassageiroDto) models.Passageiro {
	passageiro := passageiroRepository.FIndByPassageiro(passageiroDto)
	return passageiro
}