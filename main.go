package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/robfig/cron"
)

func backup(sourceDir, backupDir string) {
	now := time.Now()
	backupFileName := fmt.Sprintf("backup_%s.zip", now.Format("2006-01-02_15-04-05"))

	zipFile, err := os.Create(filepath.Join(backupDir, backupFileName))
	if err != nil {
		log.Printf("Erro ao criar o arquivo zip: %v\n", err)
		return
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	err = copyDirToZip(sourceDir, zipWriter)
	if err != nil {
		log.Printf("Erro ao realizar o backup: %v\n", err)
	}
}

func copyDirToZip(sourceDir string, zipWriter *zip.Writer) error {
	err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, _ := filepath.Rel(sourceDir, path)

		if info.IsDir() {
			return nil
		}

		fileToZip, err := os.Open(path)
		if err != nil {
			return err
		}
		defer fileToZip.Close()

		dest, err := zipWriter.Create(relPath)
		if err != nil {
			return err
		}

		_, err = io.Copy(dest, fileToZip)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

func scheduleBackup(sourceDir, backupDir string) {
	c := cron.New()

	//as  17:00  será realizado o backup
	c.AddFunc("0 17 * * *", func() {
		backup(sourceDir, backupDir)
	})
	c.Start()

	select {}
}

func main() {
	var sourceDir, backupDir string

	fmt.Print("Digite o caminho do diretório de origem: ")
	fmt.Scanln(&sourceDir)

	fmt.Print("Digite o caminho do diretório de destino: ")
	fmt.Scanln(&backupDir)

	fmt.Println("Iniciando o agendamento de backups...")
	scheduleBackup(sourceDir, backupDir)
}
