/*
*

	@author: junwang
	@since: 2023/8/25
	@desc: //TODO

*
*/
package generator
import (
"fmt"
"io"
"io/ioutil"
"log"
"os"
	"os/exec"
	"path/filepath"
)

func copyFile(srcPath, dstPath string) error {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}

func CopyDirectory(srcDir, dstDir string) error {
	files, err := ioutil.ReadDir(srcDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		srcPath := filepath.Join(srcDir, file.Name())
		dstPath := filepath.Join(dstDir, file.Name())

		if file.IsDir() {
			err = os.Mkdir(dstPath, file.Mode())
			if err != nil {
				return err
			}

			err = CopyDirectory(srcPath, dstPath)
			if err != nil {
				return err
			}
		} else {
			err = copyFile(srcPath, dstPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func CloneProject(repoURL, destinationDir string) error{
	// Git 仓库 URL
	//repoURL := "https://github.com/username/repository.git"
	// 目标目录 destinationDir := "path/to/destination"
	// 创建目标目录
	err := os.MkdirAll(destinationDir, 0755)
	if err != nil {
		log.Fatal(err)
	}

	// 执行 git clone 命令
	cmd := exec.Command("git", "clone", repoURL, destinationDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return err
	}
	fmt.Println("项目创建完成")
	return nil
}

func main() {
	srcDir := "../simbaproject"
	dstDir := "./directory"

	err := os.MkdirAll(dstDir, 0755)
	if err != nil {
		log.Fatal(err)
	}

	err = CopyDirectory(srcDir, dstDir)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("文件和子文件夹已复制到目标目录中")
}

