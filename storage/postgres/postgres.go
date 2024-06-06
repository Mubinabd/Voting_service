package postgres

import (
	"database/sql"
	"fmt"
	"project/config"

	"project/storage"

	_ "github.com/lib/pq"
)

type Storage struct {
	Db          *sql.DB
	CandidateS  storage.CandidateI
	ElectionS   storage.ElectionI
	PublicVoteS storage.PublicVotesI
}

func ConnectDb(cfg config.Config) (*Storage, error) {
	con := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)
	db, err := sql.Open("postgres", con)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	var str = &Storage{Db: db}
	str.CandidateS = NewCandidateManager(db)
	str.ElectionS = NewElectionManager(db)
	str.PublicVoteS = NewPublicVoteManager(db)

	return str, nil
}

func (s *Storage) Candidate() storage.CandidateI {
	if s.CandidateS == nil {
		s.CandidateS = NewCandidateManager(s.Db)
	}

	return s.CandidateS
}

func (s *Storage) Election() storage.ElectionI {
	if s.ElectionS == nil {
		s.ElectionS = NewElectionManager(s.Db)
	}

	return s.ElectionS
}

func (s *Storage) PublicVotes() storage.PublicVotesI {
	if s.PublicVoteS == nil {
		s.PublicVoteS = NewPublicVoteManager(s.Db)
	}

	return s.PublicVoteS
}
