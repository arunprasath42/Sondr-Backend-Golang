package repository

import (
	"sondr-backend/src/models"
	"sondr-backend/utils/database"
)

func (r *MySqlRepositoryRepo) FindAllKycStatusRequest(obj interface{}, pageNo, pageSize int) (int, error) {

	offset := (pageNo - 1) * pageSize

	var count int
	err := database.DB.Debug().Table("kycs").Where("status = 'Requested'").Select("kycs.created_at,users.id,users.first_name,users.last_name,users.last_visited").
		Joins("INNER JOIN users ON kycs.user_id=users.id").Order("user_id DESC").Count(&count).Limit(pageSize).Offset(offset).Find(obj).Error

	if err != nil {
		return count, nil
	}

	return count, nil
}
func (r *MySqlRepositoryRepo) FindAllKycStatusVerify(obj interface{}, pageNo, pageSize int) (int, error) {

	offset := (pageNo - 1) * pageSize
	var count int
	err := database.DB.Debug().Table("kycs").Where("status = 'Approved'").Select("users.id,users.first_name,users.last_name,kycs.verified_by,kycs.updated_at").
		Joins("INNER JOIN users ON kycs.user_id=users.id").Order("updated_at DESC").Count(&count).Limit(pageSize).Offset(offset).Find(obj).Error
	if err != nil {
		return count, err
	}
	return count, nil
}
func (r *MySqlRepositoryRepo) KycApproveAndDisApprove(id uint, status string, email string) error {
	var update models.Kycs
	update.Status = status
	update.VerifiedBy = email
	err := database.DB.Debug().Table("kycs").Where("status = 'Requested' && user_id = ?", id).Find(&models.Kycs{}).Update(&update).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *MySqlRepositoryRepo) InsertOrUpdateKYC(obj *models.Kycs) error {
	err := database.DB.Debug().Table("kycs").Where("user_id = ?", obj.UserId).Find(&models.Kycs{}).Error
	if err != nil {
		if err.Error() == "record not found" {
			err := database.DB.Create(&obj).Error
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	err = database.DB.Debug().Table("kycs").Where("user_id = ?", obj.UserId).Update(obj).Error
	if err != nil {
		return err
	}

	return nil
}
