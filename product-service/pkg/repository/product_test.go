package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/utils/request"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestSaveCategory(t *testing.T) {

	testCases := map[string]struct {
		input          request.AddCategory
		buildStub      func(mock sqlmock.Sqlmock)
		expectedOutput uint64
		expectedError  error
	}{
		"DBErrorShouldReturnError": {
			input: request.AddCategory{CategoryName: "CategoryName"},
			buildStub: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`INSERT INTO categories \(name\) VALUES\(\$1\) RETURNING id AS category_id`).
					WillReturnError(errors.New("database error"))
			},
			expectedOutput: 0,
			expectedError:  errors.New("database error"),
		},
		"WithMainCategoryIDShouldSaveMainCategory": {
			input: request.AddCategory{MainCategoryID: 4, CategoryName: "NewCategory"},
			buildStub: func(mock sqlmock.Sqlmock) {
				categoryID := 2

				mock.ExpectQuery(`INSERT INTO categories \(category_id, name\) VALUES\(\$1, \$2\) 
				RETURNING id AS category_id`).
					WithArgs(4, "NewCategory").
					WillReturnRows(sqlmock.NewRows([]string{"id"}).
						AddRow(categoryID))
			},
			expectedOutput: 2,
			expectedError:  nil,
		},
		"NoMainCategoryIDShouldSaveCategoryOnly": {
			input: request.AddCategory{CategoryName: "NoMainCategory"},
			buildStub: func(mock sqlmock.Sqlmock) {
				categoryID := 5
				mock.ExpectQuery(`INSERT INTO categories \(name\) VALUES\(\$1\) RETURNING id AS category_id`).
					WithArgs("NoMainCategory").
					WillReturnRows(sqlmock.NewRows([]string{"id"}).
						AddRow(categoryID))
			},
			expectedOutput: 5,
			expectedError:  nil,
		},
	}

	for testName, test := range testCases {
		test := test
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()

			gormDB, err := gorm.Open(postgres.New(postgres.Config{
				Conn: db,
			}), &gorm.Config{})
			assert.NoError(t, err)

			test.buildStub(mock)

			repo := NewProductRepository(gormDB)

			actualOutput, actualErr := repo.SaveCategory(context.Background(), test.input)
			assert.Equal(t, test.expectedError, actualErr)
			assert.Equal(t, test.expectedOutput, actualOutput)
		})
	}

}

func TestSaveVariation(t *testing.T) {

	testCases := map[string]struct {
		input          request.AddVariation
		buildStub      func(mock sqlmock.Sqlmock)
		expectedOutput uint64
		expectedError  error
	}{
		"DBErrorShouldReturnError": {
			input: request.AddVariation{CategoryID: 3, VariationName: "Variation"},
			buildStub: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`INSERT INTO variations \(category_id, name\) 
				VALUES\(\$1, \$2\) RETURNING id AS variation_id`).
					WithArgs(3, "Variation").
					WillReturnError(errors.New("database error"))
			},
			expectedOutput: 0,
			expectedError:  errors.New("database error"),
		},
	}

	for testName, test := range testCases {
		test := test
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()

			gormDB, err := gorm.Open(postgres.New(postgres.Config{
				Conn: db,
			}), &gorm.Config{})
			assert.NoError(t, err)

			test.buildStub(mock)

			repo := NewProductRepository(gormDB)

			actualOutput, actualErr := repo.SaveVariation(context.Background(), test.input)
			assert.Equal(t, test.expectedError, actualErr)
			assert.Equal(t, test.expectedOutput, actualOutput)
		})
	}
}
