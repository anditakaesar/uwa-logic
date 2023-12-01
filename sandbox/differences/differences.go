package differences

import (
	"fmt"

	"github.com/google/uuid"
)

func Run() {
	uuid1 := uuid.New()
	uuid2 := uuid.New()
	uuid3 := uuid.New()
	uuid4 := uuid.New()

	request := []uuid.UUID{
		uuid1,
		uuid2,
		uuid3,
	}

	current := []uuid.UUID{
		uuid1,
		uuid2,
		uuid4,
	}

	toDelete, toCreate := separateDifferences(request, current)
	fmt.Printf("request: %v\ncurrent: %v\ntoDelete: %v\ntoCreate: %v\n", request, current, toDelete, toCreate)
}

func separateDifferences(request []uuid.UUID, current []uuid.UUID) (toDelete []uuid.UUID, toCreate []uuid.UUID) {
	for _, r := range request {
		found := false
		for i, c := range current {
			if r == c {
				current = append(current[:i], current[i+1:]...)
				found = true
				break
			}
		}
		if !found {
			toCreate = append(toCreate, r)
		}
	}

	toDelete = current

	return toDelete, toCreate
}
