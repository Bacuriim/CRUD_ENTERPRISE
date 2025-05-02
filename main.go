package main

import (
	"crud-enterprise/models"
	"crud-enterprise/utils"
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New() // cria com ID fixo
	windowMain := a.NewWindow("CRUD Empresarial")
	windowMain.Resize(fyne.NewSize(800, 600))
	windowList := a.NewWindow("Funcionários Existentes")
	windowList.Resize(fyne.NewSize(800, 600))
	entryID := widget.NewEntry()
	entryName := widget.NewEntry()
	entryCPF := widget.NewEntry()
	entryCEP := widget.NewEntry()
	entrySalario := widget.NewEntry()
	entryDataNascimento := widget.NewEntry()
	entrySexo := widget.NewEntry()
	entryDepartamentoID := widget.NewEntry()

	entryID.SetPlaceHolder("Digite o ID aqui")
	entryName.SetPlaceHolder("Digite o nome aqui")
	entryCPF.SetPlaceHolder("Digite o CPF aqui")
	entryCEP.SetPlaceHolder("Digite o CEP aqui")
	entrySalario.SetPlaceHolder("Digite o salário aqui")
	entryDataNascimento.SetPlaceHolder("Digite a data de nascimento aqui")
	entrySexo.SetPlaceHolder("Digite o sexo aqui")
	entryDepartamentoID.SetPlaceHolder("Digite o ID do departamento aqui")

	listaFuncionarios := container.NewVBox()
	btnCriar := widget.NewButton("Criar funcionário", func() {
		depID, err := strconv.Atoi(entryDepartamentoID.Text)
		if err != nil {
			fmt.Println("Erro ao converter ID do departamento:", err)
			return
		} else {
			entryDepartamentoID.SetText(strconv.Itoa(depID))
		}
		funcionario := models.Funcionario{
			ID:             entryID.Text,
			Nome:           entryName.Text,
			CPF:            entryCPF.Text,
			CEP:            entryCEP.Text,
			Salario:        entrySalario.Text,
			DataNascimento: entryDataNascimento.Text,
			Sexo:           entrySexo.Text,
			DepartamentoID: depID,
		}
		funcionario.Salvar()
		utils.LimparCampos(entryID, entryName, entryCPF, entryCEP,
			entrySalario, entryDataNascimento, entrySexo, entryDepartamentoID)
		fmt.Println("Funcionário criado:", funcionario)
	})
	btnListar := widget.NewButton("Listar Funcionários", func() {
		funcionarios, err := models.CarregarFuncionarios()
		if err != nil {
			fmt.Println("Erro:", err)
			return
		}

		listaFuncionarios.RemoveAll()

		for _, f := range funcionarios {
			card := widget.NewCard(f.Nome,
				"",
				widget.NewLabel(fmt.Sprintf("ID: %s\nDepartamento: %d\nCPF: %s\nCEP: %s\nSalário: %s\nNascimento: %s\nSexo: %s",
					f.ID, f.DepartamentoID, f.CPF, f.CEP, f.Salario, f.DataNascimento, f.Sexo)))
			listaFuncionarios.Add(card)
		}
		windowList.Show()

		listaFuncionarios.Refresh()
		windowList.SetCloseIntercept(func() {
			windowList.Hide()
		})
	})

	windowList.SetContent(container.NewScroll(
		listaFuncionarios,
	))

	windowMain.SetContent(container.NewVBox(
		entryID,
		entryName,
		entryCPF,
		entryCEP,
		entrySalario,
		entryDataNascimento,
		entrySexo,
		entryDepartamentoID,

		container.NewVBox(
			btnCriar,
			btnListar,
		),
	))

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

	windowMain.SetCloseIntercept(func() {
		utils.EsvaziarArquivoJSON("data/funcionarios.json")
		utils.EsvaziarArquivoJSON("data/departamentos.json")
		utils.EsvaziarArquivoJSON("data/chefes_departamento.json")
		utils.EsvaziarArquivoJSON("data/funcionarios_projetos.json")
		utils.EsvaziarArquivoJSON("data/projetos.json")
		windowMain.Close()
	})
	windowMain.ShowAndRun()
}
