package aliyun_test

import (
	"testing"

	"gitee.com/infraboard/go-course/day8/cloudstation/store/provider/aliyun"
	"github.com/stretchr/testify/assert"
)

var (
	endpoint   = "http://oss-cn-chengdu.aliyuncs.com"
	ak         = "LTAI5tMvG5NA51eiH3ENZDaa"
	sk         = "vWOGbrPKQGmLVo4CKSgmAB62vdum10"
	bucketName = "cloud-station"
	filePath   = "store.go"
)

func TestUploadFile(t *testing.T) {
	should := assert.New(t)

	uploader, err := aliyun.NewUploader(endpoint, ak, sk)
	if should.NoError(err) {
		err = uploader.UploadFile(bucketName, filePath, filePath)
		should.NoError(err)
	}
}
