package command

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

/*
* 从传入的*Controller获取附件, 支持多附件上传. 并且将附件存储在app.conf中uploadpath
* 配置指定的路径中
* @author 独照幔纱
* @param 携带请求信息(RequestBody)的请求
* @param 请求中文件的key
* @return 形如[文件名称: 文件路径]的map结构, 如果过程出错则返回nil
* @return 抛出的错误信息
 */
func SaveUploadFiles(c *beego.Controller,
	fileKey string) (map[string]string, error) {

	mapFilePath := make(map[string]string)
	fileSavePath := beego.AppConfig.String("uploadpath")

	// 从RequestBody获取上传文件
	fileHeaders, err := c.GetFiles(fileKey)
	if nil != err {
		beego.Error("GET FILES: ", err.Error())
		return nil, err
	}

	// 保存文件到全局配置的指定路径
	for count, _ := range fileHeaders {
		fileName, err := SaveFile(fileHeaders[count])
		if nil != err {
			beego.Error("SAVE FILE: ", err.Error())
			return nil, err
		}

		mapFilePath[fileName] = fileSavePath + fileName
	}

	return mapFilePath, nil
}

/*
* 生成一个MD5串, 如果出错返回空字符串
* @author 独照幔纱
* @return MD5串
 */
func MakeMD5() string {
	hash := md5.New()
	_, err := io.WriteString(hash, strconv.FormatInt(time.Now().UnixNano(), 36))
	if nil != err {
		beego.Error("HASH WRITTING: ", err)
		return ""
	}

	return fmt.Sprintf("%X", hash.Sum(nil))
}

/*
* Stores the file path where set "uploadpath" at app.conf
* @author 独照幔纱
* @param Describes uploaded file
* @return Randomly generated file name
* @return The error infomation
 */
func SaveFile(fileHeader *multipart.FileHeader) (string, error) {
	// Opens & returns fileHeader's associated File
	src, err := fileHeader.Open()
	if nil != err {
		beego.Error("OPEN MULTIPARTFORM FILE: ", err.Error())
		return "", err
	}
	defer src.Close()

	// Generate a random file name and storage path Get
	tempName := MakeMD5()
	if "" == tempName {
		err = errors.New("A EMPTY NAME.")
		beego.Error("GENERATE FILE NAME: ", err.Error())
		return "", err
	}
	fileName := tempName + filepath.Ext(fileHeader.Filename)
	fileSavePath := beego.AppConfig.String("uploadpath")

	// Create a copy of the file
	dst, err := os.Create(fileSavePath + fileName)
	if nil != err {
		beego.Error("CREATE FILE: ", err.Error())
		return "", err
	}
	defer dst.Close()

	src.Seek(0, 0)
	if _, err := io.Copy(dst, src); nil != err {
		beego.Error("COPY FILE: ", err.Error())
		return "", err
	}

	return fileName, nil
}
