package main

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func toTimestamptz(t *time.Time) pgtype.Timestamptz {
    if t != nil {
        return pgtype.Timestamptz{
            Time:  *t,
            Valid: true,
        }
    }
		
    return pgtype.Timestamptz{} 
}
