package main

import (
	"crud-enterprise/models"
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New() // cria com ID fixo
	windowMain := a.NewWindow("Meu App Fyne")
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
	btnListar := widget.NewButton("Listar Funcionários", func() {
		windowList.Show()
		funcionarios, err := models.CarregarFuncionarios()
		if err != nil {
			fmt.Println("Erro:", err)
			return
		}
		listaFuncionarios.Objects = nil // limpa
		for _, f := range funcionarios {
			card := widget.NewCard(f.Nome,
				fmt.Sprintf("ID: %s | Departamento: %d", f.ID, f.DepartamentoID),
				widget.NewLabel(fmt.Sprintf("CPF: %s\nCEP: %s\nSalário: %s\nNascimento: %s\nSexo: %s",
					f.CPF, f.CEP, f.Salario, f.DataNascimento, f.Sexo)))
			listaFuncionarios.Add(card)
		}

		listaFuncionarios.Refresh()
		windowList.SetCloseIntercept(func() {
			windowList.Hide()
		})
	})

	windowList.SetContent(container.NewVBox(
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
		widget.NewButton("Criar funcionário", func() {
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
			fmt.Println("Funcionário criado:", funcionario)
		}),

		container.NewVBox(
			btnListar,
		),
	))

	// Roda a lógica do CRUD em paralelo
	go func() {
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

		// CRUD funcionários
		hiel.Salvar()
		conrado.Salvar()
		pimenta.Salvar()
		fmt.Println()
		models.ListarFuncionarios()
		fmt.Println()

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
		models.ListarFuncionarios()
		fmt.Println()
		pimenta.Deletar()
		fmt.Println()
		models.ListarFuncionarios()
		fmt.Println()

		// CRUD departamentos
		vendas := models.Departamento{ID: 1, Nome: "Vendas", ChefeID: 4}
		marketing := models.Departamento{ID: 2, Nome: "Marketing", ChefeID: 5}
		pesquisa := models.Departamento{ID: 3, Nome: "P&D", ChefeID: 6}

		vendas.Salvar()
		marketing.Salvar()
		pesquisa.Salvar()
		fmt.Println()
		models.ListarDepartamentos()
		fmt.Println()

		marketingAtualizado := models.Departamento{ID: 2, Nome: "Marketing Digital", ChefeID: 5}
		marketingAtualizado.Atualizar()
		fmt.Println()
		models.ListarDepartamentos()
		fmt.Println()
		pesquisa.Deletar()
		fmt.Println()
		models.ListarDepartamentos()
		fmt.Println()

		// utils.EsvaziarArquivoJSON("data/funcionarios.json")
		// utils.EsvaziarArquivoJSON("data/departamentos.json")
	}()

	windowMain.ShowAndRun()
}
