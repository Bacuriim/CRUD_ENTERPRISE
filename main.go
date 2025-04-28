package main

import (
	"crud-enterprise/models"
	"fmt"
)

func main() {
	// 1. Criando funcionários
	hiel := models.Funcionario{
		ID:             "A",
		Nome:           "Hiel",
		CPF:            "123.456.789-00",
		CEP:            "61635-025",
		Salario:        "5000",
		DataNascimento: "2000-11-20",
		Sexo:           "M",
		DepartamentoID: 1,
	}

	conrado := models.Funcionario{
		ID:             "B",
		Nome:           "Conrado",
		CPF:            "234.567.890-11",
		CEP:            "60142-412",
		Salario:        "3467",
		DataNascimento: "2002-10-05",
		Sexo:           "M",
		DepartamentoID: 2,
	}

	pimenta := models.Funcionario{
		ID:             "C",
		Nome:           "Pimenta",
		CPF:            "345.678.901-22",
		CEP:            "62431-931",
		Salario:        "10045",
		DataNascimento: "2004-08-05",
		Sexo:           "M",
		DepartamentoID: 3,
	}

	// Create
	hiel.Salvar()
	conrado.Salvar()
	pimenta.Salvar()
	fmt.Println()

	// Read
	models.ListarFuncionarios()
	fmt.Println()

	// Update
	conradoAtualizado := models.Funcionario{
		ID:             "B",
		Nome:           "Conrado Silva",
		CPF:            "234.567.890-11",
		CEP:            "60142-412",
		Salario:        "4000",
		DataNascimento: "2002-10-05",
		Sexo:           "M",
		DepartamentoID: 2,
	}
	conradoAtualizado.Atualizar()
	fmt.Println()

	// Read
	models.ListarFuncionarios()
	fmt.Println()

	// Delete
	pimenta.Deletar()
	fmt.Println()

	// Read
	models.ListarFuncionarios()
	fmt.Println()

	// 2. Criando departamentos
	vendas := models.Departamento{
		ID:      1,
		Nome:    "Vendas",
		ChefeID: 4,
	}

	marketing := models.Departamento{
		ID:      2,
		Nome:    "Marketing",
		ChefeID: 5,
	}

	pesquisa := models.Departamento{
		ID:      3,
		Nome:    "P&D",
		ChefeID: 6,
	}

	// Create
	vendas.Salvar()
	marketing.Salvar()
	pesquisa.Salvar()
	fmt.Println()

	// Read
	models.ListarDepartamentos()
	fmt.Println()

	// Update
	marketingAtualizado := models.Departamento{
		ID:      2,
		Nome:    "Marketing Digital",
		ChefeID: 5,
	}
	marketingAtualizado.Atualizar()
	fmt.Println()

	// Read
	models.ListarDepartamentos()
	fmt.Println()

	// Delete
	pesquisa.Deletar()
	fmt.Println()

	// Read
	models.ListarDepartamentos()
	fmt.Println()

	// Limpando arquivos .json
	// utils.EsvaziarArquivoJSON("data/funcionarios.json")
	// utils.EsvaziarArquivoJSON("data/departamentos.json")
}
