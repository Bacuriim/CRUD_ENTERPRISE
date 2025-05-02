package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const (
	arquivoFuncionariosJSON = "data/json/funcionarios.json"
	arquivoFuncionariosTXT  = "data/txt/funcionarios.txt"
)

type Funcionario struct {
	ID             string `json:"id"`
	Nome           string `json:"nome"`
	CPF            string `json:"cpf"`
	CEP            string `json:"cep"`
	Salario        string `json:"salario"`
	DataNascimento string `json:"data_nascimento"`
	Sexo           string `json:"sexo"`
	DepartamentoID int    `json:"departamento_id"`
}

// Create
func (f *Funcionario) Salvar() {
	funcionarios, err := CarregarFuncionarios()
	if err != nil {
		fmt.Printf("Erro ao carregar funcionários: %v\n", err)
		return
	}

	for _, funcionario := range funcionarios {
		if funcionario.ID == f.ID {
			fmt.Printf("Erro: já existe um funcionário com o ID %s. Ação não realizada.\n", f.ID)
			return
		}
	}

	funcionarios = append(funcionarios, *f)
	if err := salvarFuncionarios(funcionarios); err != nil {
		fmt.Printf("Erro ao salvar funcionário: %v\n", err)
		return
	}

	fmt.Printf("Funcionário com ID %s salvo com sucesso.\n", f.ID)
	sincronizarFuncionariosTxt(funcionarios)
}

// Update
func (f *Funcionario) Atualizar() {
	funcionarios, err := CarregarFuncionarios()
	if err != nil {
		fmt.Printf("Erro ao carregar funcionários: %v\n", err)
		return
	}

	for i, funcionario := range funcionarios {
		if funcionario.ID == f.ID {
			funcionarios[i] = *f
			if err := salvarFuncionarios(funcionarios); err != nil {
				fmt.Printf("Erro ao atualizar funcionário: %v\n", err)
				return
			}
			fmt.Printf("Funcionário com ID %s atualizado com sucesso.\n", f.ID)
			sincronizarFuncionariosTxt(funcionarios)
			return
		}
	}

	fmt.Printf("Erro: funcionário com ID %s não encontrado para atualização. Ação não realizada.\n", f.ID)
}

// Delete
func (f *Funcionario) Deletar() {
	funcionarios, err := CarregarFuncionarios()
	if err != nil {
		fmt.Printf("Erro ao carregar funcionários: %v\n", err)
		return
	}

	for i, funcionario := range funcionarios {
		if funcionario.ID == f.ID {
			funcionarios = append(funcionarios[:i], funcionarios[i+1:]...)
			if err := salvarFuncionarios(funcionarios); err != nil {
				fmt.Printf("Erro ao deletar funcionário: %v\n", err)
				return
			}
			fmt.Printf("Funcionário com ID %s deletado com sucesso.\n", f.ID)
			sincronizarFuncionariosTxt(funcionarios)
			return
		}
	}

	fmt.Printf("Erro: funcionário com ID %s não encontrado para exclusão. Ação não realizada.\n", f.ID)
}

// Read
func ListarFuncionarios() {
	funcionarios, err := CarregarFuncionarios()
	if err != nil {
		fmt.Printf("Erro ao carregar funcionários: %v\n", err)
		return
	}

	fmt.Println("Lista de funcionários:")
	for _, f := range funcionarios {
		fmt.Printf("ID: %s, Nome: %s, CPF: %s, CEP: %s, Salário: %s, Data de Nascimento: %s, Sexo: %s, DepartamentoID: %d\n",
			f.ID, f.Nome, f.CPF, f.CEP, f.Salario, f.DataNascimento, f.Sexo, f.DepartamentoID)
	}
}

// Funções utilitárias

func CarregarFuncionarios() ([]Funcionario, error) {
	var funcionarios []Funcionario

	file, err := os.Open(arquivoFuncionariosJSON)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return funcionarios, nil // Arquivo não existe ainda
		}
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&funcionarios); err != nil {
		return nil, err
	}

	return funcionarios, nil
}

func salvarFuncionarios(funcionarios []Funcionario) error {
	file, err := os.Create(arquivoFuncionariosJSON)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(funcionarios)
}

// Sincronizar o estado do JSON com o arquivo TXT
func sincronizarFuncionariosTxt(funcionarios []Funcionario) {
	file, err := os.Create(arquivoFuncionariosTXT)
	if err != nil {
		fmt.Printf("Erro ao criar arquivo TXT: %v\n", err)
		return
	}
	defer file.Close()

	for _, f := range funcionarios {
		linha := fmt.Sprintf("ID: %s, Nome: %s, CPF: %s, CEP: %s, Salário: %s, Data de Nascimento: %s, Sexo: %s, DepartamentoID: %d\n",
			f.ID, f.Nome, f.CPF, f.CEP, f.Salario, f.DataNascimento, f.Sexo, f.DepartamentoID)
		_, err := file.WriteString(linha)
		if err != nil {
			fmt.Printf("Erro ao escrever no arquivo TXT: %v\n", err)
			return
		}
	}

	fmt.Println("Arquivo TXT sincronizado com sucesso.")
}
