package handlers

import (
	"fmt"
	"net/http"

	"api-ai/internal/logger"

	"go.uber.org/zap"
)

// UploadPDF godoc
// @Summary      Upload a PDF file with a prompt
// @Description  Accepts a PDF file and an optional prompt for processing
// @Tags         api
// @Accept       multipart/form-data
// @Produce      plain
// @Param        pdf     formData  file   true  "PDF file to upload"
// @Param        prompt  formData  string false "Optional processing prompt"
// @Success      200     {string}  string "PDF uploaded"
// @Failure      400     {object}  ErrorResponse
// @Router       /api/upload [post]
//
// TODO: This need to be reviwed when we get the documentation.
func UploadPDF(w http.ResponseWriter, r *http.Request) {
	// Parse form
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		logger.Default.Warn("Invalid form data", zap.Error(err))
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}
	prompt := r.FormValue("prompt")
	file, _, err := r.FormFile("pdf")
	if err != nil {
		logger.Default.Warn("PDF missing in form", zap.Error(err))
		http.Error(w, "PDF not found in request", http.StatusBadRequest)
		return
	}
	defer file.Close()

	logger.Default.Info("PDF uploaded",
		zap.String("prompt", prompt),
	)

	fmt.Fprintln(w, "PDF uploaded")
}
