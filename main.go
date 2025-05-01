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

	// 3. Criando chefes de departamento
	chefe1 := models.ChefeDepartamento{
		ID:            1,
		FuncionarioID: "A", // ID de um funcionário existente
	}

	chefe2 := models.ChefeDepartamento{
		ID:            2,
		FuncionarioID: "B", // ID de um funcionário existente
	}

	chefe3 := models.ChefeDepartamento{
		ID:            3,
		FuncionarioID: "C", // ID de um funcionário existente
	}

	// Create
	chefe1.Salvar()
	chefe2.Salvar()
	chefe3.Salvar()
	fmt.Println()

	// Read
	models.ListarChefesDepartamento()
	fmt.Println()

	// Update
	chefeAtualizado := models.ChefeDepartamento{
		ID:            2,
		FuncionarioID: "D", // Atualizando o FuncionarioID
	}
	chefeAtualizado.Atualizar()
	fmt.Println()

	// Read
	models.ListarChefesDepartamento()
	fmt.Println()

	// Delete
	chefe3.Deletar()
	fmt.Println()

	// Read
	models.ListarChefesDepartamento()
	fmt.Println()

	// 4. Criando relações funcionário-projeto
	relacao1 := models.FuncionarioProjeto{
		ID:            "D",
		FuncionarioID: "A", // ID de um funcionário existente
		ProjetoID:     201, // ID de um projeto existente
	}

	relacao2 := models.FuncionarioProjeto{
		ID:            "E",
		FuncionarioID: "B", // ID de um funcionário existente
		ProjetoID:     202, // ID de um projeto existente
	}

	relacao3 := models.FuncionarioProjeto{
		ID:            "F",
		FuncionarioID: "C", // ID de um funcionário existente
		ProjetoID:     203, // ID de um projeto existente
	}

	// Create
	relacao1.Salvar()
	relacao2.Salvar()
	relacao3.Salvar()
	fmt.Println()

	// Read
	models.ListarFuncionariosProjetos()
	fmt.Println()

	// Update
	relacaoAtualizada := models.FuncionarioProjeto{
		ID:            "G",
		FuncionarioID: "D", // Atualizando o FuncionarioID
		ProjetoID:     204, // Atualizando o ProjetoID
	}
	relacaoAtualizada.Atualizar()
	fmt.Println()

	// Read
	models.ListarFuncionariosProjetos()
	fmt.Println()

	// Delete
	relacao3.Deletar()
	fmt.Println()

	// Read
	models.ListarFuncionariosProjetos()
	fmt.Println()

	// 5. Criando projetos
	projeto1 := models.Projeto{
		ID:                     1,
		Nome:                   "Projeto Alpha",
		Local:                  "São Paulo",
		DepartamentoID:         10,  // ID de um departamento existente
		FuncionariosProjetosID: "D", // ID de uma relação funcionário-projeto existente
	}

	projeto2 := models.Projeto{
		ID:                     2,
		Nome:                   "Projeto Beta",
		Local:                  "Rio de Janeiro",
		DepartamentoID:         20,  // ID de um departamento existente
		FuncionariosProjetosID: "E", // ID de uma relação funcionário-projeto existente
	}

	projeto3 := models.Projeto{
		ID:                     3,
		Nome:                   "Projeto Gamma",
		Local:                  "Belo Horizonte",
		DepartamentoID:         30,  // ID de um departamento existente
		FuncionariosProjetosID: "F", // ID de uma relação funcionário-projeto existente
	}

	// Create
	projeto1.Salvar()
	projeto2.Salvar()
	projeto3.Salvar()
	fmt.Println()

	// Read
	models.ListarProjetos()
	fmt.Println()

	// Update
	projetoAtualizado := models.Projeto{
		ID:                     2,
		Nome:                   "Projeto Beta Atualizado",
		Local:                  "Curitiba",
		DepartamentoID:         25,  // Atualizando o DepartamentoID
		FuncionariosProjetosID: "G", // Atualizando o FuncionariosProjetosID
	}
	projetoAtualizado.Atualizar()
	fmt.Println()

	// Read
	models.ListarProjetos()
	fmt.Println()

	// Delete
	projeto3.Deletar()
	fmt.Println()

	// Read
	models.ListarProjetos()
	fmt.Println()

	// Limpando arquivos .json
	// utils.EsvaziarArquivoJSON("data/json/funcionarios.json")
	// utils.EsvaziarArquivoJSON("data/json/departamentos.json")
	// utils.EsvaziarArquivoJSON("data/json/chefes_departamento.json")
	// utils.EsvaziarArquivoJSON("data/json/funcionarios_projetos.json")
	// utils.EsvaziarArquivoJSON("data/json/projetos.json")
}
