package photoprism

import (
	"os"
	"testing"

	"github.com/photoprism/photoprism/internal/thumb"
	"github.com/photoprism/photoprism/pkg/fs"

	"github.com/photoprism/photoprism/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestMediaFile_DateCreated(t *testing.T) {
	conf := config.TestConfig()

	t.Run("iphone_7.heic", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.heic")
		assert.Nil(t, err)
		date := mediaFile.DateCreated().UTC()
		assert.Equal(t, "2018-09-10 03:16:13 +0000 UTC", date.String())
		assert.Empty(t, err)
	})
	t.Run("canon_eos_6d.dng", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/canon_eos_6d.dng")
		assert.Nil(t, err)
		date := mediaFile.DateCreated().UTC()
		assert.Equal(t, "2019-06-06 07:29:51 +0000 UTC", date.String())
		assert.Empty(t, err)
	})
	t.Run("elephants.jpg", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/elephants.jpg")
		assert.Nil(t, err)
		date := mediaFile.DateCreated().UTC()
		assert.Equal(t, "2013-11-26 13:53:55 +0000 UTC", date.String())
		assert.Empty(t, err)
	})
	t.Run("dog_created_1919.jpg", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/dog_created_1919.jpg")
		assert.Nil(t, err)
		date := mediaFile.DateCreated().UTC()
		assert.Equal(t, "1919-05-04 05:59:26 +0000 UTC", date.String())
		assert.Empty(t, err)
	})
}

func TestMediaFile_HasTimeAndPlace(t *testing.T) {
	t.Run("/beach_wood.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/beach_wood.jpg")
		assert.Nil(t, err)
		assert.Equal(t, true, mediaFile.HasTimeAndPlace())
	})
	t.Run("/peacock_blue.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/peacock_blue.jpg")
		assert.Nil(t, err)
		assert.Equal(t, false, mediaFile.HasTimeAndPlace())
	})
}
func TestMediaFile_CameraModel(t *testing.T) {
	t.Run("/beach_wood.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/beach_wood.jpg")
		assert.Nil(t, err)
		assert.Equal(t, "iPhone SE", mediaFile.CameraModel())
	})
	t.Run("/iphone_7.heic", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.heic")
		assert.Nil(t, err)
		assert.Equal(t, "iPhone 7", mediaFile.CameraModel())
	})
}

func TestMediaFile_CameraMake(t *testing.T) {
	t.Run("/beach_wood.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/beach_wood.jpg")
		assert.Nil(t, err)
		assert.Equal(t, "Apple", mediaFile.CameraMake())
	})
	t.Run("/peacock_blue.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/peacock_blue.jpg")
		assert.Nil(t, err)
		assert.Equal(t, "", mediaFile.CameraMake())
	})
}

func TestMediaFile_LensModel(t *testing.T) {
	t.Run("/beach_wood.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/beach_wood.jpg")
		assert.Nil(t, err)
		assert.Equal(t, "iPhone SE back camera 4.15mm f/2.2", mediaFile.LensModel())
	})
	t.Run("/canon_eos_6d.dng", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/canon_eos_6d.dng")
		assert.Nil(t, err)
		assert.Equal(t, "EF24-105mm f/4L IS USM", mediaFile.LensModel())
	})
}

func TestMediaFile_LensMake(t *testing.T) {
	t.Run("/cat_brown.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/cat_brown.jpg")
		assert.Nil(t, err)
		assert.Equal(t, "Apple", mediaFile.LensMake())
	})
	t.Run("/elephants.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/elephants.jpg")
		assert.Nil(t, err)
		assert.Equal(t, "", mediaFile.LensMake())
	})
}

func TestMediaFile_FocalLength(t *testing.T) {
	t.Run("/cat_brown.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/cat_brown.jpg")
		assert.Nil(t, err)
		assert.Equal(t, 29, mediaFile.FocalLength())
	})
	t.Run("/elephants.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/elephants.jpg")
		assert.Nil(t, err)
		assert.Equal(t, 111, mediaFile.FocalLength())
	})
}

func TestMediaFile_FNumber(t *testing.T) {
	t.Run("/cat_brown.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/cat_brown.jpg")
		assert.Nil(t, err)
		assert.Equal(t, 2.2, mediaFile.FNumber())
	})
	t.Run("/elephants.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/elephants.jpg")
		assert.Nil(t, err)
		assert.Equal(t, 10.0, mediaFile.FNumber())
	})
}

func TestMediaFile_Iso(t *testing.T) {
	t.Run("/cat_brown.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/cat_brown.jpg")
		assert.Nil(t, err)
		assert.Equal(t, 32, mediaFile.Iso())
	})
	t.Run("/elephants.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/elephants.jpg")
		assert.Nil(t, err)
		assert.Equal(t, 200, mediaFile.Iso())
	})
}

func TestMediaFile_Exposure(t *testing.T) {
	t.Run("/cat_brown.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/cat_brown.jpg")
		assert.Nil(t, err)
		assert.Equal(t, "1/50", mediaFile.Exposure())
	})
	t.Run("/elephants.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/elephants.jpg")
		assert.Nil(t, err)
		assert.Equal(t, "1/640", mediaFile.Exposure())
	})
}

func TestMediaFileCanonicalName(t *testing.T) {
	conf := config.TestConfig()

	mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/beach_wood.jpg")
	assert.Nil(t, err)
	assert.Equal(t, "20180111_110938_B6B8AB4F", mediaFile.CanonicalName())
}

func TestMediaFileCanonicalNameFromFile(t *testing.T) {
	t.Run("/beach_wood.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/beach_wood.jpg")
		assert.Nil(t, err)
		assert.Equal(t, "beach_wood", mediaFile.CanonicalNameFromFile())
	})
	t.Run("/airport_grey", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/airport_grey")
		assert.Nil(t, err)
		assert.Equal(t, "airport_grey", mediaFile.CanonicalNameFromFile())
	})
}

func TestMediaFile_CanonicalNameFromFileWithDirectory(t *testing.T) {
	conf := config.TestConfig()

	mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/beach_wood.jpg")
	assert.Nil(t, err)
	assert.Equal(t, conf.ExamplesPath()+"/beach_wood", mediaFile.CanonicalNameFromFileWithDirectory())
}

func TestMediaFile_EditedFilename(t *testing.T) {
	conf := config.TestConfig()

	t.Run("IMG_4120.JPG", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/IMG_4120.JPG")
		assert.Nil(t, err)
		assert.Nil(t, err)
		assert.Equal(t, conf.ExamplesPath()+"/IMG_E4120.JPG", mediaFile.EditedName())
	})

	t.Run("fern_green.jpg", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/fern_green.jpg")
		assert.Nil(t, err)
		assert.Nil(t, err)
		assert.Equal(t, "", mediaFile.EditedName())
	})
}

func TestMediaFile_RelatedFiles(t *testing.T) {
	conf := config.TestConfig()

	t.Run("canon_eos_6d.dng", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/canon_eos_6d.dng")

		assert.Nil(t, err)

		expectedBaseFilename := conf.ExamplesPath() + "/canon_eos_6d"

		related, err := mediaFile.RelatedFiles()

		assert.Nil(t, err)

		assert.Len(t, related.files, 3)

		for _, result := range related.files {
			t.Logf("FileName: %s", result.FileName())

			filename := result.FileName()

			extension := result.Extension()

			baseFilename := filename[0 : len(filename)-len(extension)]

			assert.Equal(t, expectedBaseFilename, baseFilename)
		}
	})

	t.Run("iphone_7.heic", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.heic")

		assert.Nil(t, err)

		expectedBaseFilename := conf.ExamplesPath() + "/iphone_7"

		related, err := mediaFile.RelatedFiles()

		assert.Nil(t, err)

		assert.Len(t, related.files, 3)

		for _, result := range related.files {
			t.Logf("FileName: %s", result.FileName())

			filename := result.FileName()

			extension := result.Extension()

			baseFilename := filename[0 : len(filename)-len(extension)]

			assert.Equal(t, expectedBaseFilename, baseFilename)
		}
	})
}

func TestMediaFile_RelatedFiles_Ordering(t *testing.T) {
	conf := config.TestConfig()

	mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/IMG_4120.JPG")

	assert.Nil(t, err)

	related, err := mediaFile.RelatedFiles()

	assert.Nil(t, err)

	assert.Len(t, related.files, 5)

	assert.Equal(t, conf.ExamplesPath()+"/IMG_4120.AAE", related.files[0].FileName())
	assert.Equal(t, conf.ExamplesPath()+"/IMG_4120.JPG", related.files[1].FileName())

	for _, result := range related.files {
		filename := result.FileName()
		t.Logf("FileName: %s", filename)
	}
}

func TestMediaFile_SetFilename(t *testing.T) {
	conf := config.TestConfig()

	mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/turtle_brown_blue.jpg")
	assert.Nil(t, err)
	mediaFile.SetFileName("newFilename")
	assert.Equal(t, "newFilename", mediaFile.fileName)
	mediaFile.SetFileName("turtle_brown_blue")
	assert.Equal(t, "turtle_brown_blue", mediaFile.fileName)
}

func TestMediaFile_RelativeFilename(t *testing.T) {
	conf := config.TestConfig()

	mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/tree_white.jpg")
	assert.Nil(t, err)

	t.Run("directory with end slash", func(t *testing.T) {
		filename := mediaFile.RelativeName("/go/src/github.com/photoprism/photoprism/assets/resources/")
		assert.Equal(t, "examples/tree_white.jpg", filename)
	})

	t.Run("directory without end slash", func(t *testing.T) {
		filename := mediaFile.RelativeName("/go/src/github.com/photoprism/photoprism/assets/resources")
		assert.Equal(t, "examples/tree_white.jpg", filename)
	})
	t.Run("directory not part of filename", func(t *testing.T) {
		filename := mediaFile.RelativeName("xxx/")
		assert.Equal(t, conf.ExamplesPath()+"/tree_white.jpg", filename)
	})
	t.Run("directory equals example path", func(t *testing.T) {
		filename := mediaFile.RelativeName("/go/src/github.com/photoprism/photoprism/assets/resources/examples")
		assert.Equal(t, "tree_white.jpg", filename)
	})
}

func TestMediaFile_RelativePath(t *testing.T) {

	conf := config.TestConfig()

	mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/tree_white.jpg")
	assert.Nil(t, err)

	t.Run("directory with end slash", func(t *testing.T) {
		path := mediaFile.RelativePath("/go/src/github.com/photoprism/photoprism/assets/resources/")
		assert.Equal(t, "examples", path)
	})
	t.Run("directory without end slash", func(t *testing.T) {
		path := mediaFile.RelativePath("/go/src/github.com/photoprism/photoprism/assets/resources")
		assert.Equal(t, "examples", path)
	})
	t.Run("directory equals filepath", func(t *testing.T) {
		path := mediaFile.RelativePath("/go/src/github.com/photoprism/photoprism/assets/resources/examples")
		assert.Equal(t, "", path)
	})
	t.Run("directory does not match filepath", func(t *testing.T) {
		path := mediaFile.RelativePath("xxx")
		assert.Equal(t, "/go/src/github.com/photoprism/photoprism/assets/resources/examples", path)
	})
}

func TestMediaFile_RelativeBasename(t *testing.T) {
	conf := config.TestConfig()

	mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/tree_white.jpg")
	assert.Nil(t, err)

	t.Run("directory with end slash", func(t *testing.T) {
		basename := mediaFile.RelativeBase("/go/src/github.com/photoprism/photoprism/assets/resources/")
		assert.Equal(t, "examples/tree_white", basename)
	})
	t.Run("directory without end slash", func(t *testing.T) {
		basename := mediaFile.RelativeBase("/go/src/github.com/photoprism/photoprism/assets/resources")
		assert.Equal(t, "examples/tree_white", basename)
	})
	t.Run("directory equals example path", func(t *testing.T) {
		basename := mediaFile.RelativeBase("/go/src/github.com/photoprism/photoprism/assets/resources/examples/")
		assert.Equal(t, "tree_white", basename)
	})

}

func TestMediaFile_Directory(t *testing.T) {
	t.Run("/limes.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/limes.jpg")
		assert.Nil(t, err)
		assert.Equal(t, conf.ExamplesPath(), mediaFile.Directory())
	})
}

func TestMediaFile_Basename(t *testing.T) {
	t.Run("/limes.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/limes.jpg")
		assert.Nil(t, err)
		assert.Equal(t, "limes", mediaFile.Base())
	})
	t.Run("/IMG_4120 copy.JPG", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/IMG_4120 copy.JPG")
		assert.Nil(t, err)
		assert.Equal(t, "IMG_4120", mediaFile.Base())
	})
	t.Run("/IMG_4120 (1).JPG", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/IMG_4120 (1).JPG")
		assert.Nil(t, err)
		assert.Equal(t, "IMG_4120", mediaFile.Base())
	})
}

func TestMediaFile_MimeType(t *testing.T) {
	conf := config.TestConfig()

	t.Run("elephants.jpg", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/elephants.jpg")
		assert.Nil(t, err)
		assert.Nil(t, err)
		assert.Equal(t, "image/jpeg", mediaFile.MimeType())
	})

	t.Run("canon_eos_6d.dng", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/canon_eos_6d.dng")
		assert.Nil(t, err)
		assert.Nil(t, err)
		assert.Equal(t, "application/octet-stream", mediaFile.MimeType())

	})

	t.Run("iphone_7.xmp", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.xmp")
		assert.Nil(t, err)
		assert.Nil(t, err)
		assert.Equal(t, "text/plain; charset=utf-8", mediaFile.MimeType())
	})

	t.Run("iphone_7.json", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.json")
		assert.Nil(t, err)
		assert.Nil(t, err)
		assert.Equal(t, "text/plain; charset=utf-8", mediaFile.MimeType())
	})

	t.Run("iphone_7.heic", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.heic")
		assert.Nil(t, err)
		assert.Nil(t, err)
		assert.Equal(t, "application/octet-stream", mediaFile.MimeType())
	})

	t.Run("IMG_4120.AAE", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/IMG_4120.AAE")
		assert.Nil(t, err)
		assert.Nil(t, err)
		assert.Equal(t, "text/xml; charset=utf-8", mediaFile.MimeType())
	})
}

func TestMediaFile_Exists(t *testing.T) {
	conf := config.TestConfig()

	mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/cat_black.jpg")
	assert.Nil(t, err)
	assert.NotNil(t, mediaFile)
	assert.True(t, mediaFile.Exists())

	mediaFile, err = NewMediaFile(conf.ExamplesPath() + "/xxz.jpg")
	assert.NotNil(t, err)
	assert.Nil(t, mediaFile)
}

func TestMediaFile_Move(t *testing.T) {
	conf := config.TestConfig()

	tmpPath := conf.CachePath() + "/_tmp/TestMediaFile_Move"
	origName := tmpPath + "/original.jpg"
	destName := tmpPath + "/destination.jpg"

	os.MkdirAll(tmpPath, os.ModePerm)

	defer os.RemoveAll(tmpPath)

	f, err := NewMediaFile(conf.ExamplesPath() + "/table_white.jpg")
	assert.Nil(t, err)
	f.Copy(origName)
	assert.True(t, fs.FileExists(origName))

	m, err := NewMediaFile(origName)
	assert.Nil(t, err)

	if err = m.Move(destName); err != nil {
		t.Errorf("failed to move: %s", err)
	}

	assert.True(t, fs.FileExists(destName))
	assert.Equal(t, destName, m.FileName())
}

func TestMediaFile_Copy(t *testing.T) {
	conf := config.TestConfig()

	tmpPath := conf.CachePath() + "/_tmp/TestMediaFile_Copy"

	os.MkdirAll(tmpPath, os.ModePerm)

	defer os.RemoveAll(tmpPath)

	mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/table_white.jpg")
	assert.Nil(t, err)
	mediaFile.Copy(tmpPath + "table_whitecopy.jpg")
	assert.True(t, fs.FileExists(tmpPath+"table_whitecopy.jpg"))
}

func TestMediaFile_Extension(t *testing.T) {
	t.Run("/iphone_7.json", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.json")
		assert.Nil(t, err)
		assert.Equal(t, ".json", mediaFile.Extension())
	})
	t.Run("/iphone_7.heic", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.heic")
		assert.Nil(t, err)
		assert.Equal(t, ".heic", mediaFile.Extension())
	})
	t.Run("/canon_eos_6d.dng", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/canon_eos_6d.dng")
		assert.Nil(t, err)
		assert.Equal(t, ".dng", mediaFile.Extension())
	})
	t.Run("/elephants.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/elephants.jpg")
		assert.Nil(t, err)
		assert.Equal(t, ".jpg", mediaFile.Extension())
	})
}

func TestMediaFile_IsJpeg(t *testing.T) {
	t.Run("/iphone_7.json", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.json")
		assert.Nil(t, err)
		assert.Equal(t, false, mediaFile.IsJpeg())
	})
	t.Run("/iphone_7.heic", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.heic")
		assert.Nil(t, err)
		assert.Equal(t, false, mediaFile.IsJpeg())
	})
	t.Run("/canon_eos_6d.dng", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/canon_eos_6d.dng")
		assert.Nil(t, err)
		assert.Equal(t, false, mediaFile.IsJpeg())
	})
	t.Run("/elephants.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/elephants.jpg")
		assert.Nil(t, err)
		assert.Equal(t, true, mediaFile.IsJpeg())
	})
}

func TestMediaFile_HasType(t *testing.T) {
	t.Run("/iphone_7.heic", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.heic")
		assert.Nil(t, err)
		assert.Equal(t, false, mediaFile.HasType("jpg"))
	})
	t.Run("/iphone_7.heic", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.heic")
		assert.Nil(t, err)
		assert.Equal(t, true, mediaFile.HasType("heif"))
	})
	t.Run("/iphone_7.xmp", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.xmp")
		assert.Nil(t, err)
		assert.Equal(t, true, mediaFile.HasType("xmp"))
	})
}

func TestMediaFile_IsHEIF(t *testing.T) {
	t.Run("/iphone_7.json", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.json")
		assert.Nil(t, err)
		assert.Equal(t, false, mediaFile.IsHEIF())
	})
	t.Run("/iphone_7.heic", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.heic")
		assert.Nil(t, err)
		assert.Equal(t, true, mediaFile.IsHEIF())
	})
	t.Run("/canon_eos_6d.dng", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/canon_eos_6d.dng")
		assert.Nil(t, err)
		assert.Equal(t, false, mediaFile.IsHEIF())
	})
	t.Run("/elephants.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/elephants.jpg")
		assert.Nil(t, err)
		assert.Equal(t, false, mediaFile.IsHEIF())
	})
}

func TestMediaFile_IsRaw(t *testing.T) {
	t.Run("/iphone_7.json", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.json")
		assert.Nil(t, err)
		assert.Equal(t, false, mediaFile.IsRaw())
	})
	t.Run("/iphone_7.heic", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.heic")
		assert.Nil(t, err)
		assert.Equal(t, false, mediaFile.IsRaw())
	})
	t.Run("/canon_eos_6d.dng", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/canon_eos_6d.dng")
		assert.Nil(t, err)
		assert.Equal(t, true, mediaFile.IsRaw())
	})
	t.Run("/elephants.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/elephants.jpg")
		assert.Nil(t, err)
		assert.Equal(t, false, mediaFile.IsRaw())
	})
}

func TestMediaFile_IsPng(t *testing.T) {
	t.Run("/iphone_7.json", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.json")
		assert.Nil(t, err)
		assert.Equal(t, false, mediaFile.IsPng())
	})
	t.Run("/tweethog.png", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/tweethog.png")
		assert.Nil(t, err)
		assert.Equal(t, true, mediaFile.IsPng())
	})
}

func TestMediaFile_IsTiff(t *testing.T) {
	t.Run("/iphone_7.json", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.json")
		assert.Nil(t, err)
		assert.Equal(t, false, mediaFile.IsTiff())
	})
	t.Run("/purple.tiff", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/purple.tiff")
		assert.Nil(t, err)
		assert.Equal(t, true, mediaFile.IsTiff())
	})
}

func TestMediaFile_IsImageOther(t *testing.T) {
	t.Run("/iphone_7.json", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.json")
		assert.Nil(t, err)
		assert.Equal(t, false, mediaFile.IsImageOther())
	})
	t.Run("/purple.tiff", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/purple.tiff")
		assert.Nil(t, err)
		assert.Equal(t, true, mediaFile.IsImageOther())
	})
	t.Run("/tweethog.png", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/tweethog.png")
		assert.Nil(t, err)
		assert.Equal(t, true, mediaFile.IsImageOther())
	})
	t.Run("/yellow_rose-small.bmp", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/yellow_rose-small.bmp")
		assert.Nil(t, err)
		assert.Equal(t, true, mediaFile.IsImageOther())
	})
	t.Run("/preloader.gif", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/preloader.gif")
		assert.Nil(t, err)
		assert.Equal(t, true, mediaFile.IsImageOther())
	})
}

func TestMediaFile_IsSidecar(t *testing.T) {
	t.Run("/iphone_7.xmp", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.xmp")
		assert.Nil(t, err)
		assert.Equal(t, true, mediaFile.IsSidecar())
	})
	t.Run("/IMG_4120.AAE", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/IMG_4120.AAE")
		assert.Nil(t, err)
		assert.Equal(t, true, mediaFile.IsSidecar())
	})
	t.Run("/test.xml", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/test.xml")

		assert.Nil(t, err)
		assert.Equal(t, true, mediaFile.IsSidecar())
	})
	t.Run("/test.txt", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/test.txt")

		assert.Nil(t, err)
		assert.Equal(t, true, mediaFile.IsSidecar())
	})
	t.Run("/test.yml", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/test.yml")

		assert.Nil(t, err)
		assert.Equal(t, true, mediaFile.IsSidecar())
	})
	t.Run("/test.md", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/test.md")

		assert.Nil(t, err)
		assert.Equal(t, true, mediaFile.IsSidecar())
	})
	t.Run("/canon_eos_6d.dng", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/canon_eos_6d.dng")

		assert.Nil(t, err)
		assert.Equal(t, false, mediaFile.IsSidecar())
	})
}

func TestMediaFile_IsPhoto(t *testing.T) {
	t.Run("/iphone_7.json", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.json")
		assert.Nil(t, err)
		assert.Equal(t, false, mediaFile.IsPhoto())
	})
	t.Run("/iphone_7.xmp", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.xmp")
		assert.Nil(t, err)
		assert.Equal(t, false, mediaFile.IsPhoto())
	})
	t.Run("/iphone_7.heic", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.heic")
		assert.Nil(t, err)
		assert.Equal(t, true, mediaFile.IsPhoto())
	})
	t.Run("/canon_eos_6d.dng", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/canon_eos_6d.dng")
		assert.Nil(t, err)
		assert.Equal(t, true, mediaFile.IsPhoto())
	})
	t.Run("/elephants.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/elephants.jpg")
		assert.Nil(t, err)
		assert.Equal(t, true, mediaFile.IsPhoto())
	})
}

func TestMediaFile_IsVideo(t *testing.T) {
	t.Run("/christmas.mp4", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/christmas.mp4")
		assert.Nil(t, err)
		assert.Equal(t, false, mediaFile.IsPhoto())
	})
	t.Run("/canon_eos_6d.dng", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/canon_eos_6d.dng")
		assert.Nil(t, err)
		assert.Equal(t, true, mediaFile.IsPhoto())
	})
}

func TestMediaFile_Jpeg(t *testing.T) {
	t.Run("/Random.docx", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/Random.docx")
		assert.Nil(t, err)
		file, err := mediaFile.Jpeg()
		assert.Nil(t, file)
		assert.Equal(t, "jpeg file does not exist: "+conf.ExamplesPath()+"/Random.jpg", err.Error())
	})
	t.Run("/ferriswheel_colorful.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/ferriswheel_colorful.jpg")
		assert.Nil(t, err)
		file, err := mediaFile.Jpeg()
		assert.Nil(t, err)
		assert.FileExists(t, file.fileName)
	})
	t.Run("/iphone_7.json", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.json")
		assert.Nil(t, err)
		file, err := mediaFile.Jpeg()
		assert.Nil(t, file)
		assert.Equal(t, "jpeg file does not exist: "+conf.ExamplesPath()+"/iphone_7.jpg", err.Error())
	})
}

func TestMediaFile_decodeDimension(t *testing.T) {
	t.Run("/Random.docx", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/Random.docx")
		assert.Nil(t, err)
		decodeErr := mediaFile.decodeDimensions()
		assert.Equal(t, "not a photo: "+conf.ExamplesPath()+"/Random.docx", decodeErr.Error())
	})
	t.Run("/clock_purple.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/clock_purple.jpg")
		assert.Nil(t, err)
		decodeErr := mediaFile.decodeDimensions()
		assert.Nil(t, decodeErr)
	})
	t.Run("/iphone_7.heic", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.heic")
		assert.Nil(t, err)
		decodeErr := mediaFile.decodeDimensions()
		assert.Nil(t, decodeErr)
	})
}

func TestMediaFile_Width(t *testing.T) {
	t.Run("/Random.docx", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/Random.docx")
		assert.Nil(t, err)
		width := mediaFile.Width()
		assert.Equal(t, 0, width)
	})
	t.Run("/elephant_mono.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/elephant_mono.jpg")
		assert.Nil(t, err)
		width := mediaFile.Width()
		assert.Equal(t, 416, width)
	})
}

func TestMediaFile_Height(t *testing.T) {
	t.Run("/Random.docx", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/Random.docx")
		assert.Nil(t, err)
		height := mediaFile.Height()
		assert.Equal(t, 0, height)
	})
	t.Run("/elephants.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/elephants.jpg")
		assert.Nil(t, err)
		height := mediaFile.Height()
		assert.Equal(t, 331, height)
	})
}

func TestMediaFile_AspectRatio(t *testing.T) {
	t.Run("/iphone_7.heic", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.heic")
		assert.Nil(t, err)
		ratio := mediaFile.AspectRatio()
		assert.Equal(t, float64(0), ratio)
	})
	t.Run("/fern_green.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/fern_green.jpg")
		assert.Nil(t, err)
		ratio := mediaFile.AspectRatio()
		assert.Equal(t, float64(1), ratio)
	})
	t.Run("/elephants.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/elephants.jpg")
		assert.Nil(t, err)
		ratio := mediaFile.AspectRatio()
		assert.Equal(t, 1.501510574018127, ratio)
	})
}

func TestMediaFile_Orientation(t *testing.T) {
	t.Run("/iphone_7.heic", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.heic")
		assert.Nil(t, err)
		orientation := mediaFile.Orientation()
		assert.Equal(t, 6, orientation)
	})
	t.Run("/turtle_brown_blue.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/turtle_brown_blue.jpg")
		assert.Nil(t, err)
		orientation := mediaFile.Orientation()
		assert.Equal(t, 1, orientation)
	})
}

func TestMediaFile_Thumbnail(t *testing.T) {
	conf := config.TestConfig()

	if err := conf.CreateDirectories(); err != nil {
		t.Error(err)
	}

	thumbsPath := conf.CachePath() + "/_tmp"

	defer os.RemoveAll(thumbsPath)

	t.Run("/elephants.jpg", func(t *testing.T) {
		image, err := NewMediaFile(conf.ExamplesPath() + "/elephants.jpg")
		assert.Nil(t, err)

		thumbnail, err := image.Thumbnail(thumbsPath, "tile_500")

		assert.Empty(t, err)

		assert.FileExists(t, thumbnail)
	})
	t.Run("invalid image format", func(t *testing.T) {
		image, err := NewMediaFile(conf.ExamplesPath() + "/canon_eos_6d.xmp")
		assert.Nil(t, err)

		thumbnail, err := image.Thumbnail(thumbsPath, "tile_500")

		assert.Equal(t, "mediafile: could not create thumbnail (image: unknown format)", err.Error())
		t.Log(thumbnail)
	})
	t.Run("invalid thumbnail type", func(t *testing.T) {
		image, err := NewMediaFile(conf.ExamplesPath() + "/elephants.jpg")
		assert.Nil(t, err)

		thumbnail, err := image.Thumbnail(thumbsPath, "invalid_500")

		assert.Equal(t, "mediafile: invalid type invalid_500", err.Error())
		t.Log(thumbnail)
	})
}

func TestMediaFile_Resample(t *testing.T) {
	conf := config.TestConfig()

	if err := conf.CreateDirectories(); err != nil {
		t.Error(err)
	}

	thumbsPath := conf.CachePath() + "/_tmp"

	defer os.RemoveAll(thumbsPath)
	t.Run("/elephants.jpg", func(t *testing.T) {
		image, err := NewMediaFile(conf.ExamplesPath() + "/elephants.jpg")
		assert.Nil(t, err)

		thumbnail, err := image.Resample(thumbsPath, "tile_500")

		assert.Empty(t, err)
		assert.NotEmpty(t, thumbnail)

	})
	t.Run("invalid type", func(t *testing.T) {
		image, err := NewMediaFile(conf.ExamplesPath() + "/elephants.jpg")
		assert.Nil(t, err)

		thumbnail, err := image.Resample(thumbsPath, "xxx_500")

		assert.Equal(t, "mediafile: invalid type xxx_500", err.Error())
		assert.Empty(t, thumbnail)

	})

}

func TestMediaFile_RenderDefaultThumbs(t *testing.T) {
	conf := config.TestConfig()

	thumbsPath := conf.CachePath() + "/_tmp"

	defer os.RemoveAll(thumbsPath)

	if err := conf.CreateDirectories(); err != nil {
		t.Error(err)
	}

	m, err := NewMediaFile(conf.ExamplesPath() + "/elephants.jpg")
	assert.Nil(t, err)

	err = m.ResampleDefault(thumbsPath, true)

	assert.Empty(t, err)

	thumbFilename, err := thumb.Filename(m.Hash(), thumbsPath, thumb.Types["tile_50"].Width, thumb.Types["tile_50"].Height, thumb.Types["tile_50"].Options...)

	assert.Empty(t, err)

	assert.FileExists(t, thumbFilename)

	err = m.ResampleDefault(thumbsPath, false)

	assert.Empty(t, err)
}
