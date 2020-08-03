package repository

import (
	"log"

	"github.com/heliosmc89/api-rest-gowww/api/models"
	"github.com/jmoiron/sqlx"
)

type CityInterface interface {
	Create(city *models.City) (*models.City, error)
	Update(ID string, city *models.City) (*models.City, error)
	Delete(ID string) error
	FindOne(ID string) (*models.City, error)
	FindAll() ([]models.City, error)
}

type CityRepo struct {
	Log *log.Logger
	DB  *sqlx.DB
}

func NewRepoCity(logging *log.Logger, db *sqlx.DB) CityInterface {
	return &CityRepo{
		Log: logging,
		DB:  db,
	}
}

func (c *CityRepo) Create(city *models.City) (*models.City, error) {

	// Create prepared statement.
	stmt, err := c.DB.Prepare(`INSERT INTO city("Name", "CountryCode", "District", "Population") VALUES($1, $2, $3, $4)`)
	if err != nil {
		return nil, err
	}

	// Execute the prepared statement and retrieve the results.
	res, err := stmt.Exec(city.Name, city.CountryCode, city.District, city.Population)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer stmt.Close()
	_, err = res.RowsAffected()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return city, err
}

func (c *CityRepo) Update(ID string, city *models.City) (*models.City, error) {
	// Create prepared statement.
	stmt, err := c.DB.Prepare(`UPDATE city
	SET "Name"=$1, "CountryCode"=$2, "Population"=$3, "District"=$4
	WHERE id=$5`)
	if err != nil {
		return nil, err
	}

	// Execute the prepared statement and retrieve the results.
	res, err := stmt.Exec(city.Name, city.CountryCode, city.Population, city.District, ID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer stmt.Close()
	_, err = res.RowsAffected()
	if err != nil {
		return nil, err
	}

	return city, nil
}

func (c *CityRepo) Delete(ID string) error {
	// Create prepared statement.
	stmt, err := c.DB.Prepare(`DELETE FROM city WHERE id=$1`)
	if err != nil {
		c.Log.Println(err)
		return err
	}

	// Execute the prepared statement and retrieve the results.
	res, err := stmt.Exec(ID)
	if err != nil {
		c.Log.Println(err)
		return err
	}
	defer stmt.Close()
	_, err = res.RowsAffected()
	if err != nil {
		c.Log.Println(err)
		return err
	}

	return nil
}

func (c *CityRepo) FindOne(ID string) (*models.City, error) {
	var city models.City

	err := c.DB.QueryRow(`SELECT id, "Name", "CountryCode", "District", "Population" FROM city WHERE id=$1 LIMIT 1`, ID).Scan(&city.ID, &city.Name, &city.CountryCode, &city.District, &city.Population)
	if err != nil {
		return nil, err
	}
	return &city, nil
}

func (c *CityRepo) FindAll() ([]models.City, error) {
	var cities []models.City
	var city models.City

	cityResults, err := c.DB.Query(`SELECT id, "Name", "CountryCode", "District", "Population" FROM city`)
	if err != nil {
		c.Log.Fatal(err)
		return nil, err
	}
	defer cityResults.Close()

	for cityResults.Next() {
		err = cityResults.Scan(&city.ID, &city.Name, &city.CountryCode, &city.District, &city.Population)
		if err != nil {
			return nil, err
		}

		cities = append(cities, city)
	}

	return cities, nil
}
