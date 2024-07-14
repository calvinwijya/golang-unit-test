package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFizzbuzz(t *testing.T) {
	Buzz(100)
}

//sub table benchmark
func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "calvin",
			request: "calvin",
		},
		{
			name:    "phil",
			request: "phil",
		},
		{
			name:    "calvin wijaya wijaya wijaya",
			request: "calvin wijaya wijaya wijaya",
		},
	}
	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

//sub benchmark
func BenchmarkSub(b *testing.B) {
	b.Run("calvin", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("calvin")
		}
	})
	b.Run("phil", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("phil")
		}
	})
}

//bench mark
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("calvin")
	}
}

func BenchmarkHelloPhil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("phil")
	}
}

//table unit testing

func TestTableTes(t *testing.T) {
	Tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "calvin",
			request:  "calvin",
			expected: "hello calvin",
		},
		{
			name:     "phil",
			request:  "phil",
			expected: "hello phil",
		},
		{
			name:     "acin",
			request:  "acin",
			expected: "hello acin",
		},
	}
	for _, test := range Tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

//subtest function didalam function

func TestSubTes(t *testing.T) {
	t.Run("calvin", func(t *testing.T) {
		result := HelloWorld("calvin")
		require.Equal(t, "hello calvin", result)
	})
	t.Run("phil", func(t *testing.T) {
		result := HelloWorld("phil")
		require.Equal(t, "hello phil", result)
	})
}

//test main test yg lansung mengeksekusi semua unit tes di suatu package

func TestMain(m *testing.M) {
	//before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	//after
	fmt.Println("AFTER UNIT TEST")
}

//skipping unit test with t.Skip

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("cant run on mac os")
	}
	result := HelloWorld("calvin")
	require.Equal(t, "hello calvin", result)
}

//unit test with testify Framework

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("cavin")
	require.Equal(t, "hello calvin", result, "result must be hello calvin")
	fmt.Println("TesHelloWorldReuqire done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("calvin")
	assert.Equal(t, "hello calvin", result, "result must be hello calvin")
	fmt.Println("TesHelloWorldAssert done")
}

//unit test with if else

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("calvin")
	if result != "hello calvin" {
		//unit test failed'
		t.Error("it must be hello calvin")
	}
	fmt.Println("TestHelloWorld succces")
}

func TestHelloToPhil(t *testing.T) {
	result := HelloWorld("phil")
	if result != "hello phil" {
		//unit test failed'
		t.Fatal("result must be hello phil")
	}
	fmt.Println("TestHelloToPhil")
}
