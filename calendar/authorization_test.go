package calendar

import "testing"

func TestRandomStringURLSafe(t *testing.T) {
	for i := 0; i < 10; i++ {
		res, err := randomStringURLSafe(i)
		if err != nil {
			t.Errorf("randomStringURLSafe(%d): returned an error %v", i, err)
		}

		actual := len(res)
		if actual != i {
			t.Errorf("randomStringURLSafe(%d): expected %d, actual %d", i, i, actual)
		}
	}
}
