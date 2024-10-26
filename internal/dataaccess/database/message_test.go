package database

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type testMessageContext struct {
	cleanup             func()
	messageDataAccessor MessageDataAccessor

	account1 Accounts
	account2 Accounts
}

var (
	//go:embed migrations/mock_messages.json
	rawMessage []byte

	sut = testMessageContext{}
)

func mockAccount(accessor AccountDataAccessor, account *Accounts) *Accounts {
	account, err := accessor.CreateAccount(context.Background(), account)
	if err != nil {
		fmt.Printf("error creating account %v: %v\n", account, err)
		return nil
	}
	return account
}

func (s *testMessageContext) setupSUT() {
	account, cleanup, err := getAccountAccessor("../../../configs/configs_test.yaml")
	res, err := account.GetAccountByEmail(context.Background(), "huy@gmail.com")
	if res == nil {
		sut.account1 = *mockAccount(account, &Accounts{AccountName: "huy", Email: "huy@gmail.com", Password: "12345"})
	} else {
		sut.account1 = *res
	}
	res, err = account.GetAccountByEmail(context.Background(), "huy1@gmail.com")
	if res == nil {
		sut.account2 = *mockAccount(account, &Accounts{AccountName: "huy1", Email: "huy1@gmail.com", Password: "12345"})
	} else {
		sut.account2 = *res
	}

	cleanup()

	access, cleanup, err := getMessageAccessor("../../../configs/configs_test.yaml")
	if err != nil {
		cleanup()
	}

	s.messageDataAccessor = access
	s.cleanup = cleanup

}

func (s *testMessageContext) teardownSUT() {
	err := s.messageDataAccessor.DeleteAll(context.Background())
	if err != nil {
		fmt.Printf("Error cleaning up test data: %s\n", err)
	}
	s.cleanup()
}

func Test_CreateMessage(t *testing.T) {
	sut.setupSUT()
	defer sut.teardownSUT()

	t.Run("create message", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		res, err := sut.messageDataAccessor.CreateMessage(ctx, &Messages{
			AccountTo:   sut.account1,
			AccountFrom: sut.account2,
			Content:     "hello world",
			CreatedAt:   time.Now(),
			UpdateAt:    time.Now(),
		})
		if err != nil {
			assert.Failf(t, "error creating message", err.Error())
		}
		assert.Equal(t, "hello world", res.Content)
	})

	t.Run("create message with empty content", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		_, err := sut.messageDataAccessor.CreateMessage(ctx, &Messages{
			AccountTo:   sut.account1,
			AccountFrom: sut.account2,
			CreatedAt:   time.Now(),
			UpdateAt:    time.Now(),
		})

		if err != nil {
			assert.True(t, true)
		} else {
			assert.Fail(t, "error creating message with empty content")
		}
	})

	t.Run("create message without account from", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		_, err := sut.messageDataAccessor.CreateMessage(ctx, &Messages{
			AccountTo: sut.account1,
			Content:   "hello world",
			CreatedAt: time.Now(),
			UpdateAt:  time.Now(),
		})

		if err != nil {
			assert.True(t, true)
		} else {
			assert.Fail(t, "error creating message without account from")
		}
	})

	t.Run("create message without account to", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		_, err := sut.messageDataAccessor.CreateMessage(ctx, &Messages{
			AccountFrom: sut.account1,
			Content:     "hello world",
			CreatedAt:   time.Now(),
			UpdateAt:    time.Now(),
		})

		if err != nil {
			assert.True(t, true)
		} else {
			assert.Fail(t, "error creating message without account from")
		}
	})

}

func Test_EditMessage(t *testing.T) {
	sut.setupSUT()
	message, err := sut.messageDataAccessor.CreateMessage(context.Background(), &Messages{
		AccountTo:   sut.account1,
		AccountFrom: sut.account2,
		Content:     "hello world",
		CreatedAt:   time.Now(),
		UpdateAt:    time.Now(),
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	defer sut.teardownSUT()

	t.Run("edit message", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		message.Content = "new content"
		res, err := sut.messageDataAccessor.EditMessage(ctx, message)
		if err != nil {
			assert.Fail(t, "error editing message", err.Error())
			return
		}
		assert.Equal(t, "new content", res.Content)
	})

	t.Run("edit message with empty", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		message.Content = ""
		_, err := sut.messageDataAccessor.EditMessage(ctx, message)
		if err != nil {
			assert.True(t, true)
			return
		}
		assert.Fail(t, "error editing message with empty content")
	})
}

func Test_DeleteMessage(t *testing.T) {
	sut.setupSUT()
	message, err := sut.messageDataAccessor.CreateMessage(context.Background(), &Messages{
		AccountTo:   sut.account1,
		AccountFrom: sut.account2,
		Content:     "hello world",
		CreatedAt:   time.Now(),
		UpdateAt:    time.Now(),
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sut.teardownSUT()

	t.Run("delete message", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		if err := sut.messageDataAccessor.DeleteMessage(ctx, message.ID); err == nil {
			tmp, err := sut.messageDataAccessor.getMessageByID(ctx, message.ID)
			if err != nil {
				assert.Fail(t, err.Error())
			} else {
				assert.True(t, tmp.MessageFrom == nil)
			}
		} else {
			assert.Fail(t, err.Error())
		}
	})
}

func addMockMessage() error {
	var res []*Messages
	err := json.Unmarshal(rawMessage, &res)
	if err != nil {
		return err
	}

	for _, v := range res {
		_, err := sut.messageDataAccessor.CreateMessage(context.Background(), v)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	return nil
}

func Test_GetMessage(t *testing.T) {
	sut.setupSUT()
	if err := addMockMessage(); err != nil {
		fmt.Println(err)
		return
	}
	defer sut.teardownSUT()

	t.Run("get message", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		res, err := sut.messageDataAccessor.GetMessages(ctx, 1, 2, 0, 20)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			assert.True(t, res[19].Content == "hello 20")
		}
	})
}
