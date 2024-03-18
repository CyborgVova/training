package observer

import "testing"

func TestObserver(t *testing.T) {
	producer := NewProducer()
	consumer1 := NewConsumer()
	consumer2 := NewConsumer()
	consumer3 := NewConsumer()

	producer.AddConsumer(consumer1)
	producer.AddConsumer(consumer2, consumer3)

	tests := []struct {
		name  string
		state string
	}{
		{
			name:  "First",
			state: "First state",
		}, {
			name:  "Second",
			state: "Second state",
		}, {
			name:  "Third",
			state: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			producer.SetState(test.state)
			producer.Notify()
			want := test.state
			got := consumer1.State
			if want != got {
				t.Errorf("Consumer1: want: %s, got: %s\n", want, got)
			}

			got = consumer2.State
			if want != got {
				t.Errorf("Consumer2: want: %s, got: %s\n", want, got)
			}

			got = consumer3.State
			if want != got {
				t.Errorf("Consumer3: want: %s, got: %s\n", want, got)
			}
		})
	}
}
