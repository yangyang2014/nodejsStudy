package picture

type picture struct {
	data              interface{}
	index             int
	dir               string
	total             int16
	meetingId         uint32
	images            []string
	fakeTeachingTools []string
	min               int
	max               int
	texts             []string
	bkImages          map[string]string
}

type picData struct {
	bk            []interface{}
	texts         map[string]interface{}
	imgs          map[string]interface{}
	lines         map[string]interface{}
	teachingTools map[string]interface{}
	//TODO 需要支持任何值，需要考虑使用空接口还是any
	// bk []interface{}
	//TODO  50min 如何定义一个数组支持任意数据类型
}

func NewPicture(data1 interface{}, dir1 string, index1 int, bkImages1 map[string]string, meetingId1 uint32) *picture {
	return &picture{
		data:              data1,
		index:             index1,
		dir:               dir1,
		total:             0,
		meetingId:         meetingId1,
		images:            []string{},
		fakeTeachingTools: []string{},
		min:               100000,
		max:               0,
		// texts:             data1.texts || []string{},
		bkImages: bkImages1,
	}
}

type ImageInfo struct {
	Name string
	Size int64
}

//TODO 绘制图片
func (this *picture) GetPicture(dir string) ImageInfo {
	return ImageInfo{"testfile", 1000}
}
