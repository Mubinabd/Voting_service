package postgres

import (
	"database/sql"
	"fmt"
	pb "project/genproto/public"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

// Helper function to create a new test DB pool
func NewTestDBPool(t *testing.T) *ElectionManager {
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

	return NewElectionManager(db)
}

// Helper function to create a test election
func createTestElection() *pb.CreateElectionReq {
	return &pb.CreateElectionReq{
		Name: "Test Election",
		Date: "2024-06-01",
	}
}

func TestCreateElection(t *testing.T) {
	electionDB := NewTestDBPool(t)
	defer electionDB.conn.Close()

	testElection := createTestElection()
	err := electionDB.Create(testElection)
	assert.NoError(t, err, "Error creating election")

	getByIdReq := &pb.GetByIdReq{Id: testElection.Name}
	getByIdRes, err := electionDB.Get(getByIdReq)
	assert.NoError(t, err, "Error getting election by ID")
	assert.Equal(t, testElection.Name, getByIdRes.Name, "Election name does not match")
	assert.Equal(t, testElection.Date, getByIdRes.Date, "Election date does not match")
}

func TestUpdateElection(t *testing.T) {
	electionDB := NewTestDBPool(t)
	defer electionDB.conn.Close()

	testElection := createTestElection()
	err := electionDB.Create(testElection)
	assert.NoError(t, err, "Error creating election")

	updatedElection := &pb.ElectionUpdate{
		Name: "Updated Election Name",
		Date: "2024-07-01",
	}

	err = electionDB.Update(updatedElection)
	assert.NoError(t, err, "Error updating election")

	getByIdReq := &pb.GetByIdReq{Id: testElection.Name}
	getByIdRes, err := electionDB.Get(getByIdReq)
	assert.NoError(t, err, "Error getting election by ID")
	assert.Equal(t, updatedElection.Name, getByIdRes.Name, "Updated election name does not match")
	assert.Equal(t, updatedElection.Date, getByIdRes.Date, "Updated election date does not match")
}

func TestDeleteElection(t *testing.T) {
	electionDB := NewTestDBPool(t)
	defer electionDB.conn.Close()

	testElection := createTestElection()
	err := electionDB.Create(testElection)
	assert.NoError(t, err, "Error creating election")

	deleteReq := &pb.GetByIdReq{Id: testElection.Name}
	err = electionDB.Delete(deleteReq)
	assert.NoError(t, err, "Error deleting election")

	_, err = electionDB.Get(&pb.GetByIdReq{Id: deleteReq.Id})
	assert.Error(t, err, "Expected error when getting deleted election")
}

func TestGetAllElections(t *testing.T) {
	electionDB := NewTestDBPool(t)
	defer electionDB.conn.Close()

	// Creating multiple test elections
	for i := 0; i < 5; i++ {
		testElection := createTestElection()
		err := electionDB.Create(testElection)
		assert.NoError(t, err, "Error creating election")
	}

	filter := &pb.Filter{
		Limit:  5,
		Offset: 0,
	}

	electionList, err := electionDB.GetAll(filter)
	assert.NoError(t, err, "Error getting all elections")
	assert.GreaterOrEqual(t, len(electionList.Elections), 5, "Expected at least 5 elections in the list")
}
