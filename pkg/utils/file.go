package utils

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

//@author 尹冲
//param 文件完整路径
//@description 判断文件是否存在
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

//@author 尹冲
//param 需要创建文件目录
//@description 可以递归创建目录
func CreateDir(filePath string) error {
	if !FileExists(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		return err
	}
	return nil
}

func RenameDir(oldName string, newName string) error {
	err := os.Rename(oldName, newName)
	return err
}

/**
读取所有文件内容
*/
func ReadAllFile(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

func RemoveDir(path string) (err error) {
	err = os.RemoveAll(path)
	return err
}

//@author 尹冲
//param  文件或者目录完整路径
//@description 删除文件及其下面目录
func RemoveFile(path string) error {
	err := os.Remove(path)
	return err
}

//@author 尹冲
//param 目录路径
//@description 第一个参数为文件目录本身
func ListDir(dirpath string) ([]string, error) {
	var dir_list []string
	dir_err := filepath.Walk(dirpath,
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() {
				dir_list = append(dir_list, path)
				return nil
			}

			return nil
		})
	return dir_list, dir_err
}

//获取文件下面的子目录非文件（非递归进行下去）
func ListChildDir(folder string) (filePaths []string) {
	files, _ := ioutil.ReadDir(folder) //specify the current dir
	for _, file := range files {
		if file.IsDir() {
			filePath := folder + "/" + file.Name()
			fmt.Println(filePath)
			filePaths = append(filePaths, filePath)
		} else {
		}
	}

	return filePaths

}

//非递归获取当前文件夹下面的文件路径
func ListFile(folder string) (filePaths []string) {
	files, _ := ioutil.ReadDir(folder) //specify the current dir
	for _, file := range files {
		if file.IsDir() {
			//listFile(folder  + file.Name())
		} else {
			filePath := folder + file.Name()
			filePaths = append(filePaths, filePath)
		}
	}
	return filePaths
}

//@author 尹冲
//param  文件完整路径
//@description 获取文件大小
func GetFileSize(filename string) int64 {
	var result int64
	filepath.Walk(filename, func(path string, f os.FileInfo, err error) error {
		result = f.Size()
		return nil
	})
	return result
}

//@author 尹冲
//param srcFilePath 压缩文件完整地址 destDirPath 需要解压到的路径
//@description 解压.tar.gz
func UncompressTarGz(srcFilePath string, destDirPath string) error {
	Info("begin remove uzip path:" + destDirPath)
	e := RemoveDir(destDirPath)
	if e != nil {
		Error(e)
		return e
	}
	os.Mkdir(destDirPath, os.ModePerm)
	fr, err := os.Open(srcFilePath)
	if err != nil {
		return err
	}
	defer fr.Close()
	gr, err := gzip.NewReader(fr)
	if err != nil {
		return err
	}
	defer gr.Close()
	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if hdr.Typeflag != tar.TypeDir {
			os.MkdirAll(destDirPath+"/"+path.Dir(hdr.Name), os.ModePerm)
			fw, err := os.OpenFile(destDirPath+"/"+hdr.Name, os.O_CREATE|os.O_WRONLY, os.FileMode(hdr.Mode))
			if err != nil {
				return err
			}
			defer fw.Close()
			_, err = io.Copy(fw, tr)
			if err != nil {
				return err
			}
		}
	}
	Info("uzip file:" + srcFilePath + " to dir:" + destDirPath)
	return nil
}

/**
 * @description: 按深度，列出目录下的子目录
 * @param  dir目录名
 * @param  depth:相对当前目录的开始层级查找,对应mindepth
 * @param  maxDepth:遍历的最大层级目录,对应maxdepth find -d --maxdepth --minpath
 * @return  []string：depth-maxdepth层级下的所有目录
 * @author: wzl
 * @time  : 2019/3/28 10:25
 */
func ListDirByDepth(dirpath string, depth, maxDepth int) []string {
	if maxDepth < depth { //大于设定的深度
		return nil
	}

	var allSubDir []string
	var dir string
	var i int
	for i = 0; i < maxDepth; i++ {
		files, err := ioutil.ReadDir(dirpath) //读取目录下文件
		if err != nil {
			return nil
		}
		for _, file := range files {
			if file.IsDir() {
				dir = filepath.Join(dirpath, file.Name())
				if i >= depth {
					allSubDir = append(allSubDir, dir)
				}
			}
		}
		dirpath = dir
	}

	return allSubDir
}
