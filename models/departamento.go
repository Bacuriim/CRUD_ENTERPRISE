package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const arquivoDepartamentos = "data/departamentos.json"

type Departamento struct {
	ID      int    `json:"id"`
	Nome    string `json:"nome"`
	ChefeID int    `json:"chefe_id"`
}

// Create
func (d *Departamento) Salvar() {
	departamentos, err := carregarDepartamentos()
	if err != nil {
		fmt.Printf("Erro ao carregar departamentos: %v\n", err)
		return
	}

	for _, departamento := range departamentos {
		if departamento.ID == d.ID {
			fmt.Printf("Erro: já existe um departamento com o ID %d. Ação não realizada.\n", d.ID)
			return
		}
	}

	departamentos = append(departamentos, *d)
	if err := salvarDepartamentos(departamentos); err != nil {
		fmt.Printf("Erro ao salvar departamento: %v\n", err)
		return
	}

	fmt.Printf("Departamento com ID %d salvo com sucesso.\n", d.ID)
}

// Update
func (d *Departamento) Atualizar() {
	departamentos, err := carregarDepartamentos()
	if err != nil {
		fmt.Printf("Erro ao carregar departamentos: %v\n", err)
		return
	}

	for i, departamento := range departamentos {
		if departamento.ID == d.ID {
			departamentos[i] = *d
			if err := salvarDepartamentos(departamentos); err != nil {
				fmt.Printf("Erro ao atualizar departamento: %v\n", err)
				return
			}
			fmt.Printf("Departamento com ID %d atualizado com sucesso.\n", d.ID)
			return
		}
	}

	fmt.Printf("Erro: departamento com ID %d não encontrado para atualização. Ação não realizada.\n", d.ID)
}

// Delete
func (d *Departamento) Deletar() {
	departamentos, err := carregarDepartamentos()
	if err != nil {
		fmt.Printf("Erro ao carregar departamentos: %v\n", err)
		return
	}

	for i, departamento := range departamentos {
		if departamento.ID == d.ID {
			departamentos = append(departamentos[:i], departamentos[i+1:]...)
			if err := salvarDepartamentos(departamentos); err != nil {
				fmt.Printf("Erro ao deletar departamento: %v\n", err)
				return
			}
			fmt.Printf("Departamento com ID %d deletado com sucesso.\n", d.ID)
			return
		}
	}

	fmt.Printf("Erro: departamento com ID %d não encontrado para exclusão. Ação não realizada.\n", d.ID)
}

// Read
func ListarDepartamentos() {
	departamentos, err := carregarDepartamentos()
	if err != nil {
		fmt.Printf("Erro ao carregar departamentos: %v\n", err)
		return
	}

	fmt.Println("Lista de departamentos:")
	for _, d := range departamentos {
		fmt.Printf("%+v\n", d)
	}
}

// Funções utilitárias

func carregarDepartamentos() ([]Departamento, error) {
	var departamentos []Departamento

	file, err := os.Open(arquivoDepartamentos)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return departamentos, nil // Arquivo não existe ainda
		}
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&departamentos); err != nil {
		return nil, err
	}

	return departamentos, nil
}

func salvarDepartamentos(departamentos []Departamento) error {
	file, err := os.Create(arquivoDepartamentos)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(departamentos)
}
