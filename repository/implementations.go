package repository

import (
	"log"
)

// func (r *Repository) GetTestById(ctx context.Context, input GetTestByIdInput) (output GetTestByIdOutput, err error) {
// 	// err = r.Db.QueryRowContext(ctx, "SELECT name FROM test WHERE id = $1", input.Id).Scan(&output.Name)
// 	// if err != nil {
// 	// 	return
// 	// }
// 	return
// }

func (r Repository) SetEstate(e Estate) error {
	newEstate := Estate{
		Length: e.Length,
		Width:  e.Width,
	}

	result := r.Gorm.Create(&newEstate)
	if result.Error != nil {
		log.Println("Error : ", result.Error)
		return result.Error
	}
	log.Println("Success")
	return nil
}

// func (r *Repository) SetEstate(e Estate) error {
// 	fmt.Println("Interface Length : ", e.Length)
// 	fmt.Println("Interface Width : ", e.Width)
// 	newEstate := Estate{
// 		Length: e.Length,
// 		Width:  e.Width,
// 	}

// 	result := r.Gorm.Create(&newEstate)
// 	if result.Error != nil {
// 		log.Println("Error : ", result.Error)
// 		return result.Error
// 	}
// 	log.Println("Success")
// 	return nil
// }
