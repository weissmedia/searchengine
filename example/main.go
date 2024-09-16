package main

import (
	"fmt"
	"github.com/weissmedia/searchengine/internal/log"
	"github.com/weissmedia/searchengine/pkg/searchengine"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"time"
)

func main() {
	start := time.Now()
	// Retrieve the globally initialized logger
	logger := log.GetLogger()

	// Load the configuration
	cfg, err := searchengine.NewConfig()
	if err != nil {
		logger.Error("Error loading config", zap.Error(err))
		return
	}

	// Initialize the search engine with the configuration and logger
	engine := searchengine.NewEngine(cfg, logger)
	ctx := context.Background()

	//query := "((keyA = 'val1' OR keyB = 'val3') AND (keyC = 'val2' OR keyD = 'val2')) AND data_keyE=[0 100] AND keyC='val2'"
	query := `
	(owner_id = 'ZHsit105' OR case_leader_id = 'ZHsit105') AND
	case_state_text IN (
		'generated',
		'businesscasevalidation',
		'validated',
		'inprogress',
		'partnerexamination',
		'qualityassurance',
		'upstreamprocess'
	) AND
	inbox IN (
		'AR_IB_VB_IK_KORR',
		'AR_IB_VB_IK_BGS',
		'AR_IB_VB_IK_LD',
		'AR_IB_VB_IK_TEAM',
		'AR_BG_AK_SB_HE',
		'AR_BG_AK_SB_IK',
		'AR_BG_AK_SB_RE'
	) AND
	cases_stat_partner_protection_code <= 9`
	fmt.Println("Query:", query)
	// Execute a search query
	searchResult, err := engine.Search(ctx, query)
	if err != nil {
		logger.Error("Error executing query", zap.Error(err))
		return
	}

	marshal, err := searchResult.Marshal()
	if err != nil {
		return
	}
	fmt.Println("Marshal Results:", marshal)
	// Ausgabe der Gesamtzeit
	fmt.Printf("Total Time: %.3f ms\n", searchResult.TotalExecutionTime)

	// Ensure all log entries are written before exiting
	log.SyncLogger()
	duration := time.Since(start).Seconds() * 1000
	fmt.Println("Duration:", duration)
}
