package populator

import (
	"github.com/golang/mock/gomock"
	"github.com/rostyslavborovyk/MQP/config"
	mock_providers "github.com/rostyslavborovyk/MQP/pkg/providers/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPopulator_pushMessage(t *testing.T) {
	t.Run("Test correct", func(t *testing.T) {
		c := gomock.NewController(t)
		defer c.Finish()

		provider := mock_providers.NewMockProvider(c)
		provider.EXPECT().PushMessage("random-queue", "text/plain", "var1").Return(true)

		populator := &Populator{
			provider:     provider,
			queuesConfig: nil,
		}

		variations := make([]interface{}, 1)
		variations[0] = "var1"

		populator.pushMessage(config.Queue{
			Name: "random-queue",
			Message: config.Message{
				BodyVariations: config.BodyVariations{
					Type:       "text/plain",
					Variations: variations,
				},
				IncludeTimestamp: false,
				IncludeRandom:    false,
				Frequency:        1,
			},
		})
	})

	t.Run("Test delay", func(t *testing.T) {
		c := gomock.NewController(t)
		defer c.Finish()

		provider := mock_providers.NewMockProvider(c)
		provider.EXPECT().PushMessage("random-queue", "text/plain", "var1").Return(true)

		populator := &Populator{
			provider:     provider,
			queuesConfig: nil,
		}

		variations := make([]interface{}, 1)
		variations[0] = "var1"

		start := time.Now()

		populator.pushMessage(config.Queue{
			Name: "random-queue",
			Message: config.Message{
				BodyVariations: config.BodyVariations{
					Type:       "text/plain",
					Variations: variations,
				},
				IncludeTimestamp: false,
				IncludeRandom:    false,
				Frequency:        1,
			},
		})

		elapsed := time.Since(start)
		assert.GreaterOrEqual(t, elapsed, time.Second)
		assert.Less(t, elapsed, time.Millisecond*1100)
	})
}
