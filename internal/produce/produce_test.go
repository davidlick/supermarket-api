package produce

import (
	"errors"
	"testing"

	"github.com/Rhymond/go-money"
	"github.com/davidlick/supermarket-api/internal/mocks"
	"github.com/davidlick/supermarket-api/pkg/ramdb"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestService_Add(t *testing.T) {
	tests := []struct {
		test          string
		expectFunc    func(t *testing.T, mockRamDB *mocks.MockRamDB) []Item
		expectedError error
	}{
		{
			test: "it should add all Items",
			expectFunc: func(t *testing.T, mockRamDB *mocks.MockRamDB) []Item {
				items := []Item{
					{Code: "code-1", Name: "name-1", Price: money.New(101, "USD")},
					{Code: "code-2", Name: "name-2", Price: money.New(202, "USD")},
					{Code: "code-3", Name: "name-3", Price: money.New(303, "USD")},
					{Code: "code-4", Name: "name-4", Price: money.New(404, "USD")},
					{Code: "code-5", Name: "name-5", Price: money.New(505, "USD")},
				}

				for _, item := range items {
					rec, err := ramdb.NewRecord(item.Code, KeyProduceCode, item)
					if err != nil {
						t.Error(err)
					}

					mockRamDB.EXPECT().Insert(rec).Return(nil)
				}

				return items
			},
		},
		{
			test: "it should return an error from ramdb",
			expectFunc: func(t *testing.T, mockRamDB *mocks.MockRamDB) []Item {
				mockRamDB.EXPECT().Insert(gomock.Any()).Return(errors.New("test error"))
				return []Item{
					{Code: "code-1", Name: "name-1", Price: money.New(101, "USD")},
				}
			},
			expectedError: errors.New("test error"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.test, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRamDB := mocks.NewMockRamDB(ctrl)

			items := tc.expectFunc(t, mockRamDB)

			svc := NewService(mockRamDB)
			err := svc.Add(items)

			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestService_Add_Lowercased(t *testing.T) {
	t.Run("it should store produce codes lowercased", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		item := Item{
			Code:  "TEST_CODE",
			Name:  "Name",
			Price: money.New(101, "USD"),
		}

		mockRamDB := mocks.NewMockRamDB(ctrl)
		rec, err := ramdb.NewRecord("test_code", KeyProduceCode, item)
		if err != nil {
			t.Error(err)
		}

		mockRamDB.EXPECT().Insert(rec).Return(nil)

		svc := NewService(mockRamDB)
		err = svc.Add([]Item{item})

		assert.Nil(t, err)
	})
}

func TestService_Remove(t *testing.T) {
	tests := []struct {
		test          string
		expectFunc    func(t *testing.T, mockRamDB *mocks.MockRamDB) Item
		expectedError error
	}{
		{
			test: "it should successfully remove an item",
			expectFunc: func(t *testing.T, mockRamDB *mocks.MockRamDB) Item {
				item := Item{Code: "code-1", Name: "name-1", Price: money.New(101, "USD")}
				rec, err := ramdb.NewRecord(item.Code, KeyProduceCode, item)
				if err != nil {
					t.Error(err)
				}

				mockRamDB.EXPECT().Delete(rec).Return(nil)
				return item
			},
		},
		{
			test: "it should return an error from ramdb",
			expectFunc: func(t *testing.T, mockRamDB *mocks.MockRamDB) Item {
				mockRamDB.EXPECT().Delete(gomock.Any()).Return(errors.New("test error"))
				return Item{}
			},
			expectedError: errors.New("test error"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.test, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRamDB := mocks.NewMockRamDB(ctrl)

			item := tc.expectFunc(t, mockRamDB)

			svc := NewService(mockRamDB)
			err := svc.Remove(item)

			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestService_Get(t *testing.T) {
	tests := []struct {
		test          string
		expectFunc    func(t *testing.T, mockRamDB *mocks.MockRamDB) Item
		expectedError error
	}{
		{
			test: "it should return the item successfully",
			expectFunc: func(t *testing.T, mockRamDB *mocks.MockRamDB) Item {
				item := Item{Code: "produce_code", Name: "name-1", Price: money.New(101, "USD")}
				rec, err := ramdb.NewRecord(item.Code, KeyProduceCode, item)
				if err != nil {
					t.Error(err)
				}

				mockRamDB.EXPECT().Get(KeyProduceCode, "test_code").Return(rec, nil)
				return item
			},
		},
		{
			test: "it should return an error from ramdb",
			expectFunc: func(t *testing.T, mockRamDB *mocks.MockRamDB) Item {
				mockRamDB.EXPECT().Get(KeyProduceCode, "test_code").Return(nil, errors.New("test error"))
				return Item{}
			},
			expectedError: errors.New("test error"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.test, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRamDB := mocks.NewMockRamDB(ctrl)

			expectedItem := tc.expectFunc(t, mockRamDB)

			svc := NewService(mockRamDB)
			item, err := svc.Get("test_code")

			assert.Equal(t, expectedItem, item)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestService_Get_CaseInsensitive(t *testing.T) {
	t.Run("it should lowercase produce codes before getting", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		searchItem := Item{
			Code: "TEST_CODE",
		}

		expectedItem := Item{
			Code: "test_code",
		}

		rec, err := ramdb.NewRecord("test_code", KeyProduceCode, expectedItem)
		if err != nil {
			t.Error(err)
		}

		mockRamDB := mocks.NewMockRamDB(ctrl)
		mockRamDB.EXPECT().Get(KeyProduceCode, "test_code").Return(rec, nil)

		svc := NewService(mockRamDB)

		item, err := svc.Get(searchItem.Code)

		assert.Equal(t, item, expectedItem)
		assert.Nil(t, err)
	})
}

func TestService_All(t *testing.T) {
	tests := []struct {
		test          string
		expectFunc    func(t *testing.T, mockRamDB *mocks.MockRamDB) []Item
		expectedError error
	}{
		{
			test: "it should return all items",
			expectFunc: func(t *testing.T, mockRamDB *mocks.MockRamDB) []Item {
				items := []Item{
					{Code: "code-1", Name: "name-1", Price: money.New(101, "USD")},
					{Code: "code-2", Name: "name-2", Price: money.New(202, "USD")},
					{Code: "code-3", Name: "name-3", Price: money.New(303, "USD")},
					{Code: "code-4", Name: "name-4", Price: money.New(404, "USD")},
					{Code: "code-5", Name: "name-5", Price: money.New(505, "USD")},
				}

				var recs []*ramdb.Record
				for _, item := range items {
					rec, err := ramdb.NewRecord(item.Code, KeyProduceCode, item)
					if err != nil {
						t.Error(err)
					}

					recs = append(recs, rec)
				}

				mockRamDB.EXPECT().Select(KeyProduceCode).Return(recs, nil)

				return items
			},
		},
		{
			test: "it should return a ramdb error",
			expectFunc: func(t *testing.T, mockRamDB *mocks.MockRamDB) []Item {
				mockRamDB.EXPECT().Select(gomock.Any()).Return(nil, errors.New("test error"))
				return nil
			},
			expectedError: errors.New("test error"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.test, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRamDB := mocks.NewMockRamDB(ctrl)

			expectedItems := tc.expectFunc(t, mockRamDB)

			svc := NewService(mockRamDB)
			items, err := svc.All()

			assert.Equal(t, expectedItems, items)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
