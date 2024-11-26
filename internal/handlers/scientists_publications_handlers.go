package handlers

import (
	"context"
	"io-project-api/internal/requests"
	"io-project-api/internal/responses"
	"io-project-api/internal/services"

	"github.com/danielgtaylor/huma/v2"
)

func GetScientistPublicationByID(ctx context.Context, input *requests.ScientistPublicationID) (*responses.ScientistPublicationResponse, error) {
	response := &responses.ScientistPublicationResponse{}
	resultingScientistsPublications, err := services.GetScientistPublicationByID(input.ID)
	if err != nil {
		return nil, huma.Error400BadRequest("Failed to get ScientistPublication by ID")
	}
	response.Body = resultingScientistsPublications
	return response, nil
}