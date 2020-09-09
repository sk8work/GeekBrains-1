package packet

import "testing"

func TestSomeFunc(t *testing.T) {
	expected := "some string"
	got := SomeFunc()

	if expected != got {
		t.Errorf("TestSomeFunc: got %s, expected %s", got, expected)
	}
}

func TestTriangle(t *testing.T) {
	var testData = []struct {
		a, b, c  int
		expected int
		name     string
	}{
		{
			a:        2,
			b:        1,
			c:        1,
			expected: 4,
			name:     "first test",
		},
		{
			a:        2,
			b:        1,
			c:        2,
			expected: 5,
			name:     "second test",
		},
	}

	for i, v := range testData {

		t.Run(v.name, func(t *testing.T) {
			got := Triangle(v.a, v.b, v.c)
			if got != v.expected {
				t.Errorf("TestTriangle[%d]: got %d, expected %d", i, got, v.expected)
			}
		})
	}
}

func TestSomeBenchFunc(t *testing.T) {
	SomeBenchFunc(1)
}

func BenchmarkSomeBenchFunc(b *testing.B) {

	b.Run("some_testing_name", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SomeBenchFunc(10)
		}
	})
}
