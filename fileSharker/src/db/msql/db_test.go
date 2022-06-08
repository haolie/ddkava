package msql

import (
	"fmt"
	"strconv"
	"testing"
	"time"
	"unsafe"

	"fileSharker/config"
	"fileSharker/src/model"
)

func TestDownF(t *testing.T) {

	err := config.Load("../../../")
	if err != nil {
		t.Error(err)
		return
	}

	fm := &model.DFileModel{
		Id:           111111,
		FKey:         "1111111",
		CompressName: "111111",
		CompressSize: 111111,
		DownUrl:      "111111",
		VName:        "111111",
		SameLink:     0,
		UpdateTime:   time.Now(),
		CrtTime:      time.Now(),
	}

	dbLayer := new(mysqlDbLayer)

	//err=dbLayer.AddFile(model.SysCon_Table_Name_Down,fm)
	//if err != nil {
	//	t.Errorf("AddFile Err:%v",err)
	//	//return
	//}

	_, exists, err := dbLayer.GetDFListWithKey("Id", strconv.Itoa(int(fm.Id)))
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("GetFileWithId :%v", exists)
	}

	_, exists, err = dbLayer.GetDFListWithKey("FKey", fm.FKey)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("GetFileWithKey :%v", exists)
	}

	list, exists, err := dbLayer.GetDFListWithKey("DownUrl", fm.DownUrl)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("GetFilesWithDown :%d", len(list))
	}

	list, exists, err = dbLayer.GetDFListWithSameKey("DownUrl", "11")
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("GetFilesWithCompressName :%d", len(list))
	}

	fm.DownUrl += "323"
	err = dbLayer.UpdateFile(model.SysCon_Table_Name_Down, fm)
	if err != nil {
		t.Errorf("DownUrl Err:%v", err)
	}

	//err= dbLayer.Remove(model.SysCon_Table_Name_Down,fm)
	//if err!=nil{
	//	t.Error(err)
	//}else {
	//	_,exists,err:= dbLayer.GetDFListWithSameKey("Id",strconv.Itoa(int(fm.Id)))
	//	if err!=nil{
	//		t.Error(err)
	//	}else {
	//		if exists{
	//			t.Logf("Remove Fail")
	//		}else {
	//			t.Logf("Remove Success")
	//		}
	//
	//	}
	//}
}

func TestY(t *testing.T) {

	type e struct {
		d bool
	}
	hchanSize := unsafe.Sizeof(e{}) + uintptr(-int(unsafe.Sizeof(e{}))&(7-1))
	fmt.Println(-int(unsafe.Sizeof(e{})))
	fmt.Println(int(unsafe.Sizeof(e{})))
	fmt.Println(-int(unsafe.Sizeof(e{})) & (8 - 1))
	fmt.Println(hchanSize)

	fmt.Println(-1 & 7)
	fmt.Println(1 & 7)

	var i uint8 = 1
	i = i << 7

	fmt.Printf("%d", int8(i))
}
