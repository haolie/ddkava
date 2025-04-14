package try

import (
	"context"
	"fmt"
	"time"

	"fileSharker/src/Server"
	"fileSharker/src/db/dFile"
	"fileSharker/src/db/fileInfo"
	"fileSharker/src/db/sFile"
	"fileSharker/src/tool"
)

func init() {
	Server.RegisterStartFun("try", start)
}

func start(ctx context.Context) (errList []error) {

	return tSFileModel()
}

func tFileInfo() (errList []error) {
	var item = fileInfo.NewFInfo(fmt.Sprintf("%d", tool.CreateId()), "", 444, time.Now())
	err := fileInfo.AddItem(item)
	if err != nil {
		errList = append(errList, err)
		return
	}

	fmt.Printf("Table:%s 保存成功\n", item.GetTableName())

	obj, exists, err := fileInfo.GetItem(item.FKey)
	if err != nil {
		errList = append(errList, err)
		return
	}

	if exists {
		fmt.Printf("%+v", obj)
	} else {
		fmt.Printf("Table:%s 查询失败\n", item.GetTableName())
	}

	return
}

func tDFileModel() (errList []error) {
	var item = &dFile.DFileModel{
		Id:           tool.CreateId(),
		FKey:         fmt.Sprintf("%d", tool.CreateId()),
		CompressName: "CompressName",
		CompressSize: 64,
		DownUrl:      "www.xunniu.com",
		VName:        "",
		SameLink:     0,
		Status:       0,
		UpdateTime:   time.Now(),
		CrtTime:      time.Now(),
	}

	// 添加
	err := dFile.AddItem(item)
	if err != nil {
		errList = append(errList, fmt.Errorf("DFileModel 添加失败"))
		return
	}

	// 修改
	item.CompressSize = 128
	err = dFile.UpdateItem(item)
	if err != nil {
		errList = append(errList, fmt.Errorf("DFileModel 修改失败"))
		return
	}

	// 查询id
	obj, exists, err := dFile.GetItem(item.Id)
	if err != nil || !exists {
		errList = append(errList, fmt.Errorf("DFileModel 查询失败"))
		return
	}

	if obj.CompressSize != 128 {
		errList = append(errList, fmt.Errorf("DFileModel 修改失败"))
		return
	}

	// 查询Fkey
	list, exists, err := dFile.GetListWithKey(item.FKey)
	if err != nil || !exists || len(list) == 0 {
		errList = append(errList, fmt.Errorf("DFileModel 查询Fkey失败"))
		return
	}

	// 查询Cname
	list, exists, err = dFile.GetListWithCompressName(item.CompressName)
	if err != nil || !exists || len(list) == 0 {
		errList = append(errList, fmt.Errorf("DFileModel 查询Cname失败"))
		return
	}

	// 查询url
	list, exists, err = dFile.GetListWithUrl(item.DownUrl)
	if err != nil || !exists || len(list) == 0 {
		errList = append(errList, fmt.Errorf("DFileModel 查询url失败"))
		return
	}

	// 查询vname
	list, exists, err = dFile.GetListWithVName(item.VName)
	if err != nil || !exists || len(list) == 0 {
		errList = append(errList, fmt.Errorf("DFileModel 查询vname失败"))
		return
	}

	// 删除
	err = dFile.DeleteItem(item)
	if err != nil {
		errList = append(errList, fmt.Errorf("DFileModel 删除失败"))
		return
	}

	// 查询id
	_, exists, err = dFile.GetItem(item.Id)
	if err != nil || exists {
		errList = append(errList, fmt.Errorf("DFileModel 删除失败"))
		return
	}

	return
}

func tSFileModel() (errList []error) {
	//var item=&dFile.DFileModel{
	//	Id:           tool.CreateId(),
	//	FKey:         fmt.Sprintf("%d",tool.CreateId()),
	//	CompressName: "CompressName",
	//	CompressSize: 64,
	//	DownUrl:      "www.xunniu.com",
	//	VName:        "",
	//	SameLink:     0,
	//	Status:       0,
	//	UpdateTime:   time.Now(),
	//	CrtTime:      time.Now(),
	//}

	item := &sFile.SFileModel{
		Id:         tool.CreateId(),
		FKey:       fmt.Sprintf("%d", tool.CreateId()),
		ParentID:   tool.CreateId(),
		FName:      "FName",
		DownId:     tool.CreateId(),
		DownUrl:    fmt.Sprintf("%d", tool.CreateId()),
		VName:      fmt.Sprintf("%d", tool.CreateId()),
		UpdateTime: time.Time{},
		CrtTime:    time.Time{},
	}

	// 添加
	err := sFile.AddItem(item)
	if err != nil {
		errList = append(errList, fmt.Errorf("SFileModel 添加失败"))
		return
	}

	// 修改
	item.FName = "NewFName"
	err = sFile.UpdateItem(item)
	if err != nil {
		errList = append(errList, fmt.Errorf("SFileModel 修改失败"))
		return
	}

	// 查询id
	obj, exists, err := sFile.GetItem(item.Id)
	if err != nil || !exists {
		errList = append(errList, fmt.Errorf("SFileModel 查询失败"))
		return
	}

	if obj.FName != "NewFName" {
		errList = append(errList, fmt.Errorf("SFileModel 修改失败"))
		return
	}

	// 查询Fkey
	list, exists, err := sFile.GetListWithFKey(item.FKey)
	if err != nil || !exists || len(list) == 0 {
		errList = append(errList, fmt.Errorf("SFileModel 查询Fkey失败"))
		return
	}

	// 查询parentId
	list, exists, err = sFile.GetListWithParentId(item.ParentID)
	if err != nil || !exists || len(list) == 0 {
		errList = append(errList, fmt.Errorf("SFileModel 查询parentId失败"))
		return
	}

	// 查询downId
	list, exists, err = sFile.GetListWithDownId(item.DownId)
	if err != nil || !exists || len(list) == 0 {
		errList = append(errList, fmt.Errorf("SFileModel 查询downId失败"))
		return
	}

	// 查询url
	list, exists, err = sFile.GetListWithDUrl(item.DownUrl)
	if err != nil || !exists || len(list) == 0 {
		errList = append(errList, fmt.Errorf("SFileModel 查询url失败"))
		return
	}

	// 查询vname
	list, exists, err = sFile.GetListWithVName(item.VName)
	if err != nil || !exists || len(list) == 0 {
		errList = append(errList, fmt.Errorf("SFileModel 查询vname失败"))
		return
	}

	// 查询FName
	list, exists, err = sFile.GetListWithFKey(item.FKey)
	if err != nil || !exists || len(list) == 0 {
		errList = append(errList, fmt.Errorf("SFileModel 查询FName失败"))
		return
	}

	//// 删除
	//err=sFile.DeleteItem(item)
	//if err!=nil{
	//	errList=append(errList,fmt.Errorf("DFileModel 删除失败"))
	//	return
	//}
	//
	//// 查询id
	//_,exists,err=sFile.GetItem(item.Id)
	//if err!=nil||exists{
	//	errList=append(errList,fmt.Errorf("DFileModel 删除失败"))
	//	return
	//}

	return
}
