// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
    "github.com/openui-backend-go/common/database"
    "strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	chatFieldNames          = builder.RawFieldNames(&Chat{})
	chatRows                = strings.Join(chatFieldNames, ",")
	chatRowsExpectAutoSet   = strings.Join(stringx.Remove(chatFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	chatRowsWithPlaceHolder = strings.Join(stringx.Remove(chatFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheChatIdPrefix = "cache:chat:id:"
)

type (
	chatModel interface {
		Insert(ctx context.Context, data *Chat) error
		FindOne(ctx context.Context, id int64) (*Chat, error)
		Update(ctx context.Context, data *Chat) error
		Delete(ctx context.Context, id int64) error
        List(ctx context.Context) (*[]Chat, error)
	}

	defaultChatModel struct {
        db *database.GormDao
        cache *database.DcRedisClient
		table string
	}

	Chat struct {
		Id         int64     `db:"id"`
		UserId     string     `db:"user_id"` // 用户ID
		Title      string    `db:"title"`   // 标题
		Chat       string    `db:"chat"`
		Archived   int64     `db:"archived"`
		ShareId    string     `db:"share_id"` // 分享用户ID
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
	}
)

func newChatModel(conn *database.GormDao, c *database.DcRedisClient) *defaultChatModel {
	return &defaultChatModel{
        db:   conn,
        cache: c,
		table:      "chat",
	}
}

func (m *defaultChatModel) Delete(ctx context.Context, id int64) error {
    chat := &Chat{Id: id}
    err := m.db.Delete(ctx, m.table, chat, false)
    if err != nil {
        return err
    }
	return nil
}

func (m *defaultChatModel) FindOne(ctx context.Context, id int64) (*Chat, error) {
    chat := Chat{
        Id: id,
    }
    err := m.db.First(ctx, m.table, &chat, chat)
    if err != nil {
        return nil, err
    }
    return &chat, nil
}

func (m *defaultChatModel) Insert(ctx context.Context, data *Chat) error {
    err := m.db.Create(ctx, m.table, data)
    if err != nil {
        return err
    }
    return nil
}

func (m *defaultChatModel) Update(ctx context.Context, data *Chat) error {
    err := m.db.Update(ctx, m.table, uint(data.Id), data)
    if err != nil {
        return err
    }
    return nil
}

func (m *defaultChatModel) List(ctx context.Context) (*[]Chat, error) {
    var chats []Chat
    err := m.db.Find(ctx, m.table, &chats, "")
    if err != nil {
        return nil, err
    }
    return &chats, nil
}

func (m *defaultChatModel) tableName() string {
	return m.table
}
