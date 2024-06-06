package postgres

import (
	"database/sql"
	pb "project/genproto/public"

	"github.com/google/uuid"
)

type PublicVoteManager struct {
	conn *sql.DB
}

func NewPublicVoteManager(conn *sql.DB) *PublicVoteManager {
	return &PublicVoteManager{conn: conn}
}

func (pvm *PublicVoteManager) Create(pubvote *pb.CreatePublicVoteReq) error {
	t, err := pvm.conn.Begin()
	if err != nil {
		return err
	}

	defer t.Commit()

	id := uuid.NewString()
	query := `INSERT INTO public_vote(id,election_id,public_id)VALUES($1,$2,$3)`

	_, err = t.Exec(query, id, pubvote.ElectionId, pubvote.PublicId)
	if err != nil {
		return err
	}

	id = uuid.NewString()
	query = `INSERT INTO votes(id,candidate_id)VALUES($1,$2)`

	_, err = t.Exec(query, id, pubvote.CandidateId)
	return err
}

func (pvm *PublicVoteManager) GetAll(filter *pb.Filter) (*pb.PublicVotesGetAllResp, error) {
	query := `SELECT 
		p.id,
		p.election_id,
		p.public_id,
		p.candidate_id,
		e.name AS election_name,
		e.date AS election_date,
		pb.first_name,
		pb.last_name,
		pb.birthday,
		pb.gender,
		pb.nation,
		c.id
	FROM 
		public_vote p
	JOIN 
		election e ON p.election_id = e.id
	JOIN 
		public pb ON p.public_id = pb.id
	JOIN 
		candidate c ON p.candidate_id = c.id
	WHERE 
		p.id = $1
`
	if filter.GetLimit() > 0 {
		query += " LIMIT " + string(filter.GetLimit())
	}
	if filter.GetOffset() > 0 {
		query += " OFFSET " + string(filter.GetOffset())
	}
	rows, err := pvm.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	publicVotesList := &pb.PublicVotesGetAllResp{}

	for rows.Next() {
		var publicVotes pb.PublicVoteRes
		err := rows.Scan(
			&publicVotes.Id,
			&publicVotes.Public.Id,
			&publicVotes.Public.FirstName,
			&publicVotes.Public.LastName,
			&publicVotes.Public.Birthday,
			&publicVotes.Public.Gender,
			&publicVotes.Public.Nation,
			&publicVotes.Election.Id,
			&publicVotes.Election.Name,
			&publicVotes.Election.Date,
		)
		if err != nil {
			return nil, err
		}
		publicVotesList.Publicvotes = append(publicVotesList.Publicvotes, &publicVotes)
	}
	return publicVotesList, nil
}
