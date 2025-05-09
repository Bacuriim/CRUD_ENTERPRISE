package main

import (
	"crud-enterprise/controllers"
	"crud-enterprise/models"
)

func main() {

	go func() {
		// 1. Criando funcionários
		funcionario1 := models.Funcionario{
			ID:             "A",
			Nome:           "Alice",
			CPF:            "123.456.789-00",
			CEP:            "12345-678",
			Salario:        "5000.00",
			DataNascimento: "1990-01-01",
			Sexo:           "Feminino",
			DepartamentoID: 1,
		}

		funcionario2 := models.Funcionario{
			ID:             "B",
			Nome:           "Bob",
			CPF:            "987.654.321-00",
			CEP:            "87654-321",
			Salario:        "6000.00",
			DataNascimento: "1985-05-05",
			Sexo:           "Masculino",
			DepartamentoID: 2,
		}

		// Salvando
		funcionario1.Salvar(models.GetDepartamentosIDs())
		funcionario2.Salvar(models.GetDepartamentosIDs())

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

		// Update
		marketingAtualizado := models.Departamento{
			ID:      2,
			Nome:    "Marketing Digital",
			ChefeID: 5,
		}
		marketingAtualizado.Atualizar()

		// Delete
		pesquisa.Deletar()

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
		chefe1.Salvar(models.GetFuncionariosIDs())
		chefe2.Salvar(models.GetFuncionariosIDs())
		chefe3.Salvar(models.GetFuncionariosIDs())

		// Update
		chefeAtualizado := models.ChefeDepartamento{
			ID:            2,
			FuncionarioID: "D", // Atualizando o FuncionarioID
		}
		chefeAtualizado.Atualizar()

		// Delete
		chefe3.Deletar()

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

		// Update
		relacaoAtualizada := models.FuncionarioProjeto{
			ID:            "G",
			FuncionarioID: "D", // Atualizando o FuncionarioID
			ProjetoID:     204, // Atualizando o ProjetoID
		}
		relacaoAtualizada.Atualizar()

		// Delete
		relacao3.Deletar()

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
		projeto1.Salvar(models.GetDepartamentosIDs(), models.GetFuncionariosIDs())
		projeto2.Salvar(models.GetDepartamentosIDs(), models.GetFuncionariosIDs())
		projeto3.Salvar(models.GetDepartamentosIDs(), models.GetFuncionariosIDs())

		// Update
		projetoAtualizado := models.Projeto{
			ID:                     2,
			Nome:                   "Projeto Beta Atualizado",
			Local:                  "Curitiba",
			DepartamentoID:         25,  // Atualizando o DepartamentoID
			FuncionariosProjetosID: "G", // Atualizando o FuncionariosProjetosID
		}
		projetoAtualizado.Atualizar()

		// Delete
		projeto3.Deletar()
	}()

	controllers.Init()
}
