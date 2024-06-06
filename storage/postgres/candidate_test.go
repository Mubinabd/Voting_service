package postgres

import (
	"database/sql"
	"fmt"
	"testing"

	pb "project/genproto/public"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

// Helper function to create a new test DB pool
func newTestDBPool(t *testing.T) *CandidateManager {
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

	return NewCandidateManager(db)
}

// Helper function to create a test candidate
func createTestCandidate(electionID, publicID string) *pb.CreateCandidateReq {
	return &pb.CreateCandidateReq{
		Id:         uuid.New().String(),
		ElectionId: electionID,
		PublicId:   publicID,
	}
}

func TestCreateCandidate(t *testing.T) {
	candidateDB := newTestDBPool(t)
	defer candidateDB.conn.Close()

	// Create test election and public entities
	electionID := uuid.New().String()
	publicID := uuid.New().String()

	_, err := candidateDB.conn.Exec("INSERT INTO election(id, name, date) VALUES ($1, 'Test Election', '2024-06-06')", electionID)
	assert.NoError(t, err, "Error creating test election")

	_, err = candidateDB.conn.Exec("INSERT INTO public(id, first_name, last_name, birthday, gender, nation) VALUES ($1, 'John', 'Doe', '1990-01-01', 'M', 'Testland')", publicID)
	assert.NoError(t, err, "Error creating test public entity")

	// Create test candidate
	testCandidate := createTestCandidate(electionID, publicID)
	err = candidateDB.Create(testCandidate)
	assert.NoError(t, err, "Error creating candidate")

	getByIdReq := &pb.GetByIdReq{Id: testCandidate.Id}
	getByIdRes, err := candidateDB.Get(&pb.GetCandidate{Id: getByIdReq.Id})
	assert.NoError(t, err, "Error getting candidate by ID")

	assert.Equal(t, testCandidate.Id, getByIdRes.Id)
	assert.Equal(t, testCandidate.ElectionId, getByIdRes.Election.Id)
	assert.Equal(t, testCandidate.PublicId, getByIdRes.Public.Id)
}
