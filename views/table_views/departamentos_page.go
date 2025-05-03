package table_views

import (
	"crud-enterprise/models"
	"crud-enterprise/utils"
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func DepartamentosPage(myApp fyne.App, mainPage fyne.Window) {
	departmentsMainPage := myApp.NewWindow("Tabela de Departamentos")
	departmentsMainPage.Resize(fyne.NewSize(800, 600))

	departmentsListPage := myApp.NewWindow("Departamentos Existentes")
	departmentsListPage.Resize(fyne.NewSize(800, 600))

	icon, err := fyne.LoadResourceFromPath("assets/imgs/CRUD_IMAGE.png") // Load the icon from a file
	if err != nil {
		fmt.Println("Erro ao carregar o ícone:", err)
	} else {
		departmentsMainPage.SetIcon(icon)
		departmentsListPage.SetIcon(icon) // Set the icon for the main window
	}

	entryID := utils.CriarEntryNumeros("ID")
	entryName := utils.CriarEntryLetrasNumeros("Nome do Departamento")
	entryBossID := utils.CriarEntryNumeros("ID do Chefe")

	lbResultado := widget.NewLabel("Resultado: nenhum")

	listaDepartamentos := container.NewVBox()
	btnCriar := widget.NewButton("Criar departamento", func() {
		chefeID, err := strconv.Atoi(entryBossID.Text)
		if err != nil {
			fmt.Println("Erro ao converter ID do departamento:", err)
			return
		} else {
			entryBossID.SetText(strconv.Itoa(chefeID))
		}

		depID, err := strconv.Atoi(entryID.Text)
		if err != nil {
			fmt.Println("Erro ao converter ID do departamento:", err)
			return
		} else {
			entryID.SetText(strconv.Itoa(depID))
		}

		Departamento := models.Departamento{
			ID:      depID,
			Nome:    entryName.Text,
			ChefeID: chefeID,
		}

		lbResultado.SetText("Resultado: " + Departamento.Salvar())
		utils.LimparCampos(entryID, entryName, entryBossID)
	})

	btnListar := widget.NewButton("Listar Departamentos", func() {
		Departamentos, err := models.CarregarDepartamentos()
		if err != nil {
			fmt.Println("Erro:", err)
			return
		}

		listaDepartamentos.RemoveAll()

		for _, d := range Departamentos {
			card := widget.NewCard(d.Nome,
				"",
				widget.NewLabel(fmt.Sprintf("ID: %d\nChefeID: %d",
					d.ID, d.ChefeID)))
			listaDepartamentos.Add(card)
		}
		departmentsListPage.Show()

		listaDepartamentos.Refresh()
		departmentsListPage.SetCloseIntercept(func() {
			departmentsListPage.Hide()
		})
	})

	departmentsListPage.SetContent(container.NewScroll(
		listaDepartamentos,
	))

	rodape := widget.NewButton("Voltar", func() {
		departmentsMainPage.Hide()
		departmentsListPage.Hide()
		mainPage.Show()
	})

	departmentsMainPage.SetContent(container.NewBorder(
		container.NewVBox(
			entryID,
			entryName,
			entryBossID,
		),
		rodape,
		nil,
		nil,
		container.NewVBox(
			btnCriar,
			btnListar,
			lbResultado,
		),
	))

	departmentsMainPage.SetCloseIntercept(func() {
		// utils.EsvaziarArquivoJSON("data/Departamentos.json")
		// utils.EsvaziarArquivoJSON("data/departamentos.json")
		// utils.EsvaziarArquivoJSON("data/chefes_departamento.json")
		// utils.EsvaziarArquivoJSON("data/Departamentos_projetos.json")
		// utils.EsvaziarArquivoJSON("data/projetos.json")
		departmentsMainPage.Close()
		mainPage.Show()
	})

	departmentsMainPage.Show()
}
