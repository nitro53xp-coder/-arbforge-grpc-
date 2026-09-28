package main

import (
	"context"
	"fmt"
	"math"

	pb "github.com/nitro53xp-coder/-arbforge-grpc-/gen"
)

type ArbForgeServer struct {
	pb.UnimplementedArbForgeServiceServer
}

func (s *ArbForgeServer) ValidateDeal(ctx context.Context, req *pb.ValidateDealRequest) (*pb.ValidateDealResponse, error) {
	if req == nil {
		return &pb.ValidateDealResponse{
			Valid:   false,
			Message: "missing request payload",
			Status:  "error",
		}, nil
	}

	buyPrice := req.GetBuyPrice()
	productName := req.GetProductName()

	roi := 0.0
	if buyPrice > 0 {
		roi = ((25.0 - buyPrice) / buyPrice) * 100.0
	}

	valid := buyPrice > 0 && buyPrice < 15 && len(productName) > 0

	resp := &pb.ValidateDealResponse{
		Valid:         valid,
		Message:      fmt.Sprintf("Deal validation for %s", productName),
		EstimatedRoi: roi,
		Status:       "ok",
	}

	if !valid {
		resp.Status = "needs_review"
		resp.Message = "Deal is outside the expected price and validation range"
	}

	return resp, nil
}

func (s *ArbForgeServer) AskMonica(ctx context.Context, req *pb.AskMonicaRequest) (*pb.AskMonicaResponse, error) {
	prompt := req.GetPrompt()
	if prompt == "" {
		prompt = "Tell me the best flip opportunities this week."
	}

	answer := fmt.Sprintf("Monica analysis: Based on the prompt '%s', prioritize products with strong margins, low cost, and clear resale demand.", prompt)

	return &pb.AskMonicaResponse{
		Answer: answer,
		Tone:   "confident",
	}, nil
}

func (s *ArbForgeServer) GetConfidenceScore(ctx context.Context, req *pb.GetConfidenceScoreRequest) (*pb.GetConfidenceScoreResponse, error) {
	buyPrice := req.GetBuyPrice()
	sellPrice := req.GetSellPrice()

	if buyPrice <= 0 || sellPrice <= 0 {
		return &pb.GetConfidenceScoreResponse{
			Score:   0,
			Label:   "invalid",
			Summary: "Buy and sell prices must be greater than zero.",
		}, nil
	}

	margin := sellPrice - buyPrice
	score := (margin / sellPrice) * 100.0
	if score < 0 {
		score = 0
	}
	if score > 100 {
		score = 100
	}

	label := "low"
	switch {
	case score >= 75:
		label = "high"
	case score >= 50:
		label = "medium"
	}

	return &pb.GetConfidenceScoreResponse{
		Score:   math.Round(score*100) / 100,
		Label:   label,
		Summary: fmt.Sprintf("Expected margin is %.2f%% on this deal.", score),
	}, nil
}

func (s *ArbForgeServer) GenerateListing(ctx context.Context, req *pb.GenerateListingRequest) (*pb.GenerateListingResponse, error) {
	title := req.GetProductName()
	if title == "" {
		title = "Hot Wheels Collector Item"
	}

	buyPrice := req.GetBuyPrice()
	sellPrice := req.GetSellPrice()

	if buyPrice <= 0 {
		buyPrice = 2.0
	}
	if sellPrice <= 0 {
		sellPrice = 35.0
	}

	listingTitle := fmt.Sprintf("%s - Buy for $%.2f, Sell for $%.2f", title, buyPrice, sellPrice)
	listingBody := fmt.Sprintf(
		"Looking for a strong flip opportunity? This %s is priced to move fast. Buy at $%.2f and target $%.2f with strong demand in collector channels.",
		title,
		buyPrice,
		sellPrice,
	)

	return &pb.GenerateListingResponse{
		ListingTitle: listingTitle,
		ListingBody:  listingBody,
		Tags:         "hotwheels,collector,flip,deal",
	}, nil
}
