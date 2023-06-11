package service

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"shareMe/db"
	"shareMe/util"
	"strings"
)

func AddFile(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["fileList"]
	var prevNames []string
	shareCode := db.GenerateShareCode()
	if shareCode == "" {
		c.JSON(500, gin.H{"status": "failed"})
		return
	}
	for _, file := range files {
		err := c.SaveUploadedFile(file, util.UploadFolder+"/"+shareCode+"/"+file.Filename)
		if err != nil {
			c.JSON(500, gin.H{"status": "failed"})
			for _, name := range prevNames {
				err := os.Remove(util.UploadFolder + "/" + shareCode + "/" + name)
				if err != nil {
					fmt.Println("AddFile Delete Error!")
				}
			}
			return
		} else {
			prevNames = append(prevNames, file.Filename)
		}
	}

	res := db.CreateFileRecord(prevNames, shareCode)
	if !res {
		c.JSON(500, gin.H{"status": "failed"})
		return
	}

	c.JSON(200, gin.H{"status": "success", "shareCode": shareCode})
}

func GetFile(c *gin.Context) {
	shareCode := c.Param("shareCode")
	if db.QueryShareCodeExist(shareCode) {
		files := db.QueryFileByShareCode(shareCode)
		c.JSON(200, gin.H{"status": "success", "files": files})
	} else {
		c.JSON(500, gin.H{"err": "Incorrect ShareCode"})
	}
}

func DownloadFile(c *gin.Context) {
	// shareCode := c.Param("shareCode")
	shareCode := c.Param("shareCode")
	fileName := c.Param("fileName")
	// Open
	_, errByOpenFile := os.Open(util.UploadFolder + "/" + shareCode + "/" + fileName)
	//非空处理
	if errByOpenFile != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": "Cannot Open the File",
		})
		return
	}

	c.Header("Content-Type", "application/json; application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(util.UploadFolder + "/" + shareCode + "/" + fileName)
	return
}

func DownloadAllFile(c *gin.Context) {
	// shareCode := c.Param("shareCode")
	shareCode := c.Param("shareCode")
	zipName := "target.tar.gz"
	// Open
	_, errByOpenFile := os.Open(util.UploadFolder + "/" + shareCode + "/" + zipName)
	//非空处理
	if errByOpenFile != nil {
		Zip(util.UploadFolder+"/"+shareCode, util.UploadFolder+"/"+shareCode+"/"+zipName)
	} else {
	}

	c.Header("Content-Type", "application/json; application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+zipName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(util.UploadFolder + "/" + shareCode + "/" + zipName)
	return
}
func Zip(src, dst string) (err error) {
	// 创建文件
	fw, err := os.Create(dst)
	if err != nil {
		return
	}
	defer fw.Close()

	// 将 tar 包使用 gzip 压缩，其实添加压缩功能很简单，
	// 只需要在 fw 和 tw 之前加上一层压缩就行了，和 Linux 的管道的感觉类似
	gw := gzip.NewWriter(fw)
	defer gw.Close()

	// 创建 Tar.Writer 结构
	tw := tar.NewWriter(gw)
	// 如果需要启用 gzip 将上面代码注释，换成下面的

	defer tw.Close()

	// 下面就该开始处理数据了，这里的思路就是递归处理目录及目录下的所有文件和目录
	// 这里可以自己写个递归来处理，不过 Golang 提供了 filepath.Walk 函数，可以很方便的做这个事情
	// 直接将这个函数的处理结果返回就行，需要传给它一个源文件或目录，它就可以自己去处理
	// 我们就只需要去实现我们自己的 打包逻辑即可，不需要再去路径相关的事情
	return filepath.Walk(src, func(fileName string, fi os.FileInfo, err error) error {
		// 因为这个闭包会返回个 error ，所以先要处理一下这个
		if err != nil {
			return err
		}

		// 这里就不需要我们自己再 os.Stat 了，它已经做好了，我们直接使用 fi 即可
		hdr, err := tar.FileInfoHeader(fi, "")
		if err != nil {
			return err
		}
		// 这里需要处理下 hdr 中的 Name，因为默认文件的名字是不带路径的，
		// 打包之后所有文件就会堆在一起，这样就破坏了原本的目录结果
		// 例如： 将原本 hdr.Name 的 syslog 替换程 log/syslog
		// 这个其实也很简单，回调函数的 fileName 字段给我们返回来的就是完整路径的 log/syslog
		// strings.TrimPrefix 将 fileName 的最左侧的 / 去掉，
		// 熟悉 Linux 的都知道为什么要去掉这个
		hdr.Name = strings.TrimPrefix(fileName, string(filepath.Separator))

		// 写入文件信息
		if err := tw.WriteHeader(hdr); err != nil {
			return err
		}

		// 判断下文件是否是标准文件，如果不是就不处理了，
		// 如： 目录，这里就只记录了文件信息，不会执行下面的 copy
		if !fi.Mode().IsRegular() {
			return nil
		}

		// 打开文件
		fr, err := os.Open(fileName)
		defer fr.Close()
		if err != nil {
			return err
		}

		// copy 文件数据到 tw
		n, err := io.Copy(tw, fr)
		if err != nil {
			return err
		}

		// 记录下过程，这个可以不记录，这个看需要，这样可以看到打包的过程
		log.Printf("Successfully Packed %s ，Wrote %d B Data\n", fileName, n)

		return nil
	})
}
