package chatgpt

import (
	"context"
	"time"

	"github.com/avast/retry-go"

	"github.com/eryajf/chatgpt-dingtalk/pkg/logger"
	"github.com/eryajf/chatgpt-dingtalk/public"
)

// SingleQa 单聊
func SingleQa(question, userId string) (answer string, err error) {
	chat := New(userId)
	defer chat.Close()
	// 定义一个重试策略
	retryStrategy := []retry.Option{
		retry.Delay(100 * time.Millisecond),
		retry.Attempts(3),
		retry.LastErrorOnly(true),
	}
	// 使用重试策略进行重试
	err = retry.Do(
		func() error {
			answer, err = chat.ChatWithContext(question)
			if err != nil {
				return err
			}
			return nil
		},
		retryStrategy...)
	return
}

// ContextQa 串聊
func ContextQa(question, userId string) (chat *ChatGPT, answer string, err error) {
	chat = New(userId)
	if public.UserService.GetUserSessionContext(userId) != "" {
		err := chat.ChatContext.LoadConversation(userId)
		if err != nil {
			logger.Warning("load station failed: %v\n", err)
		}
	}
	retryStrategy := []retry.Option{
		retry.Delay(100 * time.Millisecond),
		retry.Attempts(3),
		retry.LastErrorOnly(true)}
	// 使用重试策略进行重试
	err = retry.Do(
		func() error {
			answer, err = chat.ChatWithContext(question)
			if err != nil {
				return err
			}
			return nil
		},
		retryStrategy...)
	return
}

// ImageQa 生成图片
func ImageQa(ctx context.Context, question, userId string) (answer string, err error) {
	chat := New(userId)
	defer chat.Close()
	// 定义一个重试策略
	retryStrategy := []retry.Option{
		retry.Delay(100 * time.Millisecond),
		retry.Attempts(3),
		retry.LastErrorOnly(true),
	}
	// 使用重试策略进行重试
	err = retry.Do(
		func() error {
			answer, err = chat.GenerateImage(ctx, question)
			if err != nil {
				return err
			}
			return nil
		},
		retryStrategy...)
	return
}

// SingleQa 单聊
func SingleQaAi302(question, userId, model, apiKey string) (answer string, err error) {
	chat := NewAi302(userId, model, apiKey)
	defer chat.Close()
	// 定义一个重试策略
	retryStrategy := []retry.Option{
		retry.Delay(100 * time.Millisecond),
		retry.Attempts(3),
		retry.LastErrorOnly(true),
	}
	// 使用重试策略进行重试
	err = retry.Do(
		func() error {
			answer, err = chat.ChatWithAi30Context(question, model)
			if err != nil {
				return err
			}
			return nil
		},
		retryStrategy...)
	return
}

// ContextQa 串聊
func ContextQaAi302(question, userId, model, apiKey string) (chat *ChatGPT, answer string, err error) {
	chat = NewAi302(userId, model, apiKey)
	if public.UserService.GetUserSessionContext(userId) != "" {
		err := chat.ChatContext.LoadConversation(userId)
		if err != nil {
			logger.Warning("load station failed: %v\n", err)
		}
	}
	retryStrategy := []retry.Option{
		retry.Delay(100 * time.Millisecond),
		retry.Attempts(3),
		retry.LastErrorOnly(true)}
	// 使用重试策略进行重试
	err = retry.Do(
		func() error {
			answer, err = chat.ChatWithAi30Context(question, model)
			if err != nil {
				return err
			}
			return nil
		},
		retryStrategy...)
	return
}

// ImageQa 生成图片
func ImageQaAi302(ctx context.Context, question, userId, model, apiKey string) (answer string, err error) {
	chat := NewAi302(userId, model, apiKey)
	defer chat.Close()
	// 定义一个重试策略
	retryStrategy := []retry.Option{
		retry.Delay(100 * time.Millisecond),
		retry.Attempts(3),
		retry.LastErrorOnly(true),
	}
	// 使用重试策略进行重试
	err = retry.Do(
		func() error {
			answer, err = chat.GenerateImageAi302(ctx, question, model)
			if err != nil {
				return err
			}
			return nil
		},
		retryStrategy...)
	return
}
