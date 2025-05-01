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

	windowMain.SetCloseIntercept(func() {
		utils.EsvaziarArquivoJSON("data/funcionarios.json")
		windowMain.Close()
	})
	windowMain.ShowAndRun()
}
