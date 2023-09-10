package repository

import (
	"sondr-backend/utils/database"
)

var Repo MysqlRepository

type MySqlRepositoryRepo struct{}

func MySqlInit() {
	Repo = &MySqlRepositoryRepo{}
}

/***Fetching data from database***/
func (r *MySqlRepositoryRepo) FindById(obj interface{}, id int) error {
	if err := database.DB.Debug().Where("id = ?", id).Find(obj).Error; err != nil {
		return err
	}
	return nil
}

/***For Updating tableName ***/
func (r *MySqlRepositoryRepo) Update(obj interface{}, id int, update interface{}) error {

	if err := database.DB.Debug().Where("Unique_id IN (?) ", id).First(obj).Updates(update).Error; err != nil {
		return err
	}
	return nil
}

/***********for Updating blocked user*****************/
func (r *MySqlRepositoryRepo) UpdateBlockUser(obj interface{}, id int, update interface{}) error {

	if err := database.DB.Debug().Where("id = ? ", id).First(obj).Updates(update).Error; err != nil {
		return err
	}
	return nil
}

func (r *MySqlRepositoryRepo) Delete(obj interface{}, id int) error {
	if err := database.DB.Debug().Where("Unique_id IN (?) ", id).First(obj).Delete(obj).Error; err != nil {
		return err
	}
	return nil
}

func (r *MySqlRepositoryRepo) Find(obj interface{}, tableName string, selectQuery string, whereQuery string, value ...interface{}) error {
	db := database.DB.Debug().Table(tableName)
	if selectQuery != "" {
		db.Select(selectQuery)
	}
	if err := db.Where(whereQuery, value...).Find(obj).Error; err != nil {
		return err
	}
	return nil
}

func (r *MySqlRepositoryRepo) InsertPromptQuestion(obj interface{}) error {
	if err := database.DB.Debug().Create(obj).Error; err != nil {
		return err
	}
	return nil
}

func (r *MySqlRepositoryRepo) DeleteLikeRequest(obj interface{}, id int) error {
	if err := database.DB.Debug().Where("id = ? ", id).First(obj).Unscoped().Delete(obj).Error; err != nil {
		return err
	}
	return nil
}

func (r *MySqlRepositoryRepo) DropTable(tableName string) error {
	if err := database.DB.Debug().DropTable(tableName).Error; err != nil {
		return err
	}
	return nil
}
