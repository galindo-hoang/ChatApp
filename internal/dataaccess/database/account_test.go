package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type testAccountContext struct {
	some                string
	accountDataAccessor AccountDataAccessor
	cleanUp             func()
	context             int64
}

func (c *testAccountContext) beforeEachCreate() {
	access, cleanUp, err := getAccountAccessor("../../../configs/configs_test.yaml")
	if err != nil {
		cleanUp()
	}

	c.accountDataAccessor = access
	c.cleanUp = cleanUp
}

func (c *testAccountContext) afterEachCreate() {
	err := c.accountDataAccessor.DeleteAll(context.Background())
	if err != nil {
		fmt.Printf("fail to delete all accounts: %s \n", err)
	}
	c.cleanUp()
}

func testCaseCreate(test func(t *testing.T, c *testAccountContext)) func(*testing.T) {
	return func(t *testing.T) {
		ctx := &testAccountContext{}
		ctx.beforeEachCreate()
		defer ctx.afterEachCreate()
		test(t, ctx)
	}
}

func Test_CreateAccount(t *testing.T) {
	t.Run("create data", testCaseCreate(func(t *testing.T, c *testAccountContext) {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		acc, err := c.accountDataAccessor.CreateAccount(ctx, &Accounts{AccountName: "1", Email: "asd", Password: "123"})
		if err != nil {
			fmt.Println(err)
		} else {
			assert.Equal(t, acc.AccountName, "1")
		}
	}))

	t.Run("create with duplicate email", testCaseCreate(func(t *testing.T, c *testAccountContext) {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		_, err := c.accountDataAccessor.CreateAccount(ctx, &Accounts{AccountName: "1", Email: "asd1", Password: "123"})
		if err != nil {
			fmt.Println(err)
			assert.Error(t, err)
		} else {
			_, err := c.accountDataAccessor.CreateAccount(ctx, &Accounts{AccountName: "1", Email: "asd1", Password: "123"})
			if err != nil {
				assert.True(t, true)
			} else {
				assert.Error(t, errors.New("create with duplicate unique"))
			}
		}
	}))

	t.Run("create with empty name", testCaseCreate(func(t *testing.T, c *testAccountContext) {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		_, err := c.accountDataAccessor.CreateAccount(ctx, &Accounts{Email: "asd123", Password: "123"})
		if err != nil {
			assert.True(t, true)
		} else {
			assert.Fail(t, "create with empty name")
		}
	}))

	t.Run("create with empty password", testCaseCreate(func(t *testing.T, c *testAccountContext) {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		_, err := c.accountDataAccessor.CreateAccount(ctx, &Accounts{AccountName: "1", Email: "asd432"})
		if err != nil {
			assert.True(t, true)
		} else {
			assert.Fail(t, "create with empty password")
		}
	}))
}

func Test_GetAccountWithID(t *testing.T) {
	access, cleanUp, err := getAccountAccessor("../../../configs/configs_test.yaml")
	if err != nil {
		assert.Fail(t, err.Error())
		cleanUp()
	}

	defer func() {
		err := access.DeleteAll(context.Background())
		if err != nil {
			fmt.Printf("fail to delete all accounts: %s \n", err)
		}
		cleanUp()
	}()

	_, err = access.CreateAccount(context.Background(), &Accounts{AccountName: "1", Email: "asd", Password: "123"})
	if err != nil {
		assert.Fail(t, err.Error())
	}

	t.Run("get data with existing account", func(t *testing.T) {
		res, err := access.GetAccountByID(context.Background(), 1)
		if err != nil {
			assert.Fail(t, err.Error())
		}
		assert.Equal(t, "1", res.AccountName)
	})

	t.Run("get data with non-existing account", func(t *testing.T) {
		res, err := access.GetAccountByID(context.Background(), 2)
		if res != nil || err != nil {
			assert.Fail(t, "expected error")
		}
	})
}

func Test_GetAccountWithEmail(t *testing.T) {
	access, cleanUp, err := getAccountAccessor("../../../configs/configs_test.yaml")
	if err != nil {
		assert.Fail(t, err.Error())
		cleanUp()
	}

	defer func() {
		err := access.DeleteAll(context.Background())
		if err != nil {
			fmt.Printf("fail to delete all accounts: %s \n", err)
		}
		cleanUp()
	}()

	_, err = access.CreateAccount(context.Background(), &Accounts{AccountName: "1", Email: "asd", Password: "123"})
	if err != nil {
		assert.Fail(t, err.Error())
	}

	t.Run("get data with existing account", func(t *testing.T) {
		res, err := access.GetAccountByEmail(context.Background(), "asd")
		if err != nil {
			assert.Fail(t, err.Error())
		}
		assert.Equal(t, "1", res.AccountName)
	})

	t.Run("get data with non-existing account", func(t *testing.T) {
		res, err := access.GetAccountByEmail(context.Background(), "qwe")
		if res != nil || err != nil {
			assert.Fail(t, "expected error")
		}
	})
}
