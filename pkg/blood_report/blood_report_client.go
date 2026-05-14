package bloodreportclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/iamHemaAI/hema-ai-bot/internal/domain/dto"
)

type BloodReportClient struct {
	baseURL string
	http    *http.Client
}

func NewBloodReportClient(baseURL string) *BloodReportClient {
	return &BloodReportClient{
		baseURL: baseURL,
		http:    &http.Client{},
	}
}

type analyzeRequest struct {
	Hemoglobin   int     `json:"hemoglobin"`
	RBC          float64 `json:"rbc"`
	WBC          float64 `json:"wbc"`
	PLT          int     `json:"plt"`
	HCT          int     `json:"hct"`
	MCV          int     `json:"mcv"`
	MCH          int     `json:"mch"`
	ESR          int     `json:"esr"`
	Glucose      float64 `json:"glucose"`
	TotalProtein int     `json:"total_protein"`
}

func (c *BloodReportClient) Analyze(ctx context.Context, req *dto.BloodRequest) ([]byte, error) {
	body, err := json.Marshal(analyzeRequest{
		Hemoglobin:   req.Hb,
		RBC:          req.Rbc,
		WBC:          req.Wbc,
		PLT:          req.Plt,
		HCT:          req.Hct,
		MCV:          req.Mcv,
		MCH:          req.Mch,
		ESR:          req.Esr,
		Glucose:      req.Glucose,
		TotalProtein: req.Protein,
	})
	if err != nil {
		return nil, fmt.Errorf("marshal analyze request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/api/v1/analyze", bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("create analyze request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.http.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("do analyze request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("analyze request failed: status %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}
