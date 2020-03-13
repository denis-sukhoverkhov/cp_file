package cpfile

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCopyFailedScenarios(t *testing.T) {
	tesFileName := "test1_1"
	toFile := "new" + tesFileName

	defer func(tesFileName string) {
		err := os.Remove(tesFileName)
		if err != nil {
			t.Errorf("Error of removing test_file=%v", tesFileName)
		}
	}(tesFileName)

	createTestFile(tesFileName, 100)

	t.Run("copy not existing file ", func(t *testing.T) {
		err := Copy(tesFileName+"wrong_name", toFile, 0, 0)
		expectedErrorString := "open test1_1wrong_name: no such file or directory"
		assert.EqualError(t, err, expectedErrorString)
	})
	t.Run("copy file with null size", func(t *testing.T) {
		err := Copy("/dev/urandom", toFile, 0, 0)
		expectedErrorString := "fileSizeInBytes == 0"
		assert.EqualError(t, err, expectedErrorString)
	})

	t.Run("toFile already exists", func(t *testing.T) {
		err := Copy(tesFileName, tesFileName, 0, 0)
		assert.NoError(t, err)
	})
}

func TestCopySuccess(t *testing.T) {
	tesFileName := "test1_2"
	toFile := "new" + tesFileName

	defer func(tesFileName string, toFile string) {
		err := os.Remove(tesFileName)
		if err != nil {
			t.Errorf("Error of removing test_file=%v", tesFileName)
		}
		err = os.Remove(toFile)
		if err != nil {
			t.Errorf("Error of removing test_file=%v", toFile)
		}
	}(tesFileName, toFile)

	createTestFile(tesFileName, 100)

	t.Run("copy file with offset which more than file", func(t *testing.T) {
		err := Copy(tesFileName, toFile, 0, 10000)
		assert.NoError(t, err)
		assert.FileExists(t, toFile)
		assert.Equal(t, int64(0), GetFileSize(toFile))
	})

	t.Run("copy file without offset and limit", func(t *testing.T) {
		err := Copy(tesFileName, toFile, -1, 0)
		assert.NoError(t, err)
		assert.FileExists(t, toFile)
		assert.Equal(t, int64(100), GetFileSize(toFile))
		assert.True(t, IsEqualFiles(tesFileName, toFile))
	})

	t.Run("copy file with offset and limit", func(t *testing.T) {
		err := Copy(tesFileName, toFile, 45, 20)
		assert.NoError(t, err)
	})
}

func TestCopyBigFile(t *testing.T) {
	tesFileName := "test1_3"
	toFile := "new" + tesFileName

	defer func(tesFileName string, toFile string) {
		err := os.Remove(tesFileName)
		if err != nil {
			t.Errorf("Error of removing test_file=%v", tesFileName)
		}
		err = os.Remove(toFile)
		if err != nil {
			t.Errorf("Error of removing test_file=%v", toFile)
		}
	}(tesFileName, toFile)

	createTestFile(tesFileName, 100000)

	t.Run("copy file with offset and limit", func(t *testing.T) {
		err := Copy(tesFileName, toFile, 913, 292)
		assert.NoError(t, err)
		assert.Equal(t, int64(913), GetFileSize(toFile))
	})

}
