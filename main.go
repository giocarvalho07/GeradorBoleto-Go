package main

import (
	"fmt"
	"generatorPDF/htmlParser"
	"io/ioutil"
	"log"

	"github.com/jung-kurt/gofpdf"
)

type Data struct {
	Name           string
	LastName       string
	Vencimento     string
	ValorDocumento float64
	NumeroDoc      string
	LocalPagamento string
	AGCCDAC        string
	CPF            string
	Rua            string
	Cidade         string
}

func main() {
	h := htmlParser.New("tmp")
	//wk := pdfGenerator.NewWkHtmlToPdf("tmp")

	dataHTML := Data{
		Name:           "Giovanni",
		LastName:       "Carvalho",
		Vencimento:     "30/06/2023",
		ValorDocumento: 1290.86,
		NumeroDoc:      "34567890",
		LocalPagamento: "Agencia caixa",
		AGCCDAC:        "078/4353-0",
		CPF:            "678876678-87",
		Rua:            "Rua dos bobos, nº 0",
		Cidade:         "São Paulo - SP",
	}

	htmlGenerated, err := h.Create("templates/example.html", dataHTML)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("html gerado: ", htmlGenerated)

	fmt.Println(htmlGenerated)

	htmlLog, err := h.Create("templates/log.html", dataHTML)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(htmlLog)

	filePDFName := htmlLog

	content, err := ioutil.ReadFile(filePDFName)
	if err != nil {
		log.Fatalf("arquivo não encontrado")
	}

	pdf := gofpdf.New("p", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.MultiCell(190, 5, string(content), "0", "0", false)

	err = pdf.OutputFileAndClose("logBoleto.pdf")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Pdf de log compilado com sucesso")

}
