package postgres

import (
	"database/sql"
	"fmt"
	"testing"

	"project/genproto/public"

	"github.com/stretchr/testify/assert"
	_ "github.com/lib/pq"
)

// Helper function to create a new test DB pool
func newTestDBpool(t *testing.T) *sql.DB {
	connStr := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		"localhost",
		"postgres",
		"votingsystem",
		"1234",
		5432)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		t.Fatalf("Failed to ping database: %v", err)
	}

	return db
}

// Helper function to create a test public vote
func createTestPublicVote() *public.CreatePublicVoteReq {
	return &public.CreatePublicVoteReq{
		ElectionId:  "d09eded5-0ecd-4290-af42-ba534c48060c",
		PublicId:    "d09eded5-0ecd-4290-af42-ba534c48060a",
		CandidateId: "d09eded5-0ecd-4290-af42-ba534c48060b",
	}
}

func TestCreatePublicVote(t *testing.T) {
	db := newTestDBpool(t)
	defer db.Close()

	publicVoteManager := NewPublicVoteManager(db)

	testPublicVote := createTestPublicVote()
	err := publicVoteManager.Create(testPublicVote)
	assert.NoError(t, err, "Error creating public vote")

}

func TestGetAllPublicVotes(t *testing.T) {
	db := newTestDBpool(t)
	defer db.Close()

	publicVoteManager := NewPublicVoteManager(db)

	filter := &public.Filter{
		Limit:  10,
		Offset: 0,
	}
	resp, err := publicVoteManager.GetAll(filter)
	assert.NoError(t, err, "Error getting all public votes")

	assert.NotNil(t, resp)
}
