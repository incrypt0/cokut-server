package brokers

import "github.com/incrypt0/cokut-server/models"

// CokutBroker is the interface whoch abstracts all the store functionalities
// It acts as a mediator between actual store and the handler object
type CokutBroker interface {
	InsertUser(u *models.User) (id string, err error)
	AddUserAddress(id string, address models.Address) (user *models.User, err error)
	RemoveUserAddress(uid string, address models.Address) (user *models.User, err error)
	GetUser(uid string) (l interface{}, err error)
	CheckUserPhoneExistence(phone string) (bool, error)
	CheckUserExistence(phone string, email string) (bool, error)
	CheckUserExistenceByGID(gid string) (bool, error)
	CheckUserExistenceByUID(UID string) (bool, error)
	InsertRestaurant(r *models.Restaurant) (id string, err error)
	GetAllRestaurants() (l []interface{}, err error)
	GetAllRegularRestaurants() (l []interface{}, err error)
	GetAllHomeMade() (l []interface{}, err error)
	InsertMeal(m *models.Meal) (id string, err error)
	InsertSpecial(id string) (string, error)
	GetMealsByRestaurant(rid string) (l []interface{}, err error)
	GetSpecialMeals() (l []interface{}, err error)
	GetSpiceyMeals() (l []interface{}, err error)
	CreateOrder(o *models.Order) (po *models.Order, err error)
	GetAllOrders() (l []interface{}, err error)
	GetOrdersByUser(uid string) (l []interface{}, err error)
}

// DbBroker is the interface whoch abstracts all the db functionalities
// It acts as a mediator between actual db workers and the store object
type DbBroker interface {
	DropTest() error
	Add(collectionName string, i interface{}) (id string, err error)
	DeleteOne(collectionName string, i interface{}) (n int64, err error)
	Get(collectionName string, i interface{}) (l []interface{}, err error)
	GetMultipleByID(collectionName string, model interface{}, ids []string) (l []interface{}, err error)
	FindOneAndUpdate(collectionName string, i interface{}, u interface{}) (l interface{}, err error)
	DeleteFromMap(collectionName string, filter interface{}, update interface{}) (l interface{}, err error)
	FindOneAndPush(collectionName string, i interface{}, u interface{}, field string) (l interface{}, err error)
	FindOneAndPull(collectionName string, filter interface{}, update interface{}, field string) (l interface{}, err error)
	FindOne(collectionName string, i interface{}) (l interface{}, err error)
	FindOneWithOr(collectionName string, i ...interface{}) (l interface{}, err error)
}
