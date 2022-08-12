package dumpmdbx

import (
	"fmt"
	"os"

	_ "github.com/btcsuite/btcd/database/ffldb"
)

const (
	dump    = "dump"
	restore = "restore"
)

func Start() {

	// comfil := "/Users/andy/dev/btcd/dbtest/tmp/000000000.fdb.bin"
	// decompressFile(comfil, comfil+".dec")
	// return

	if len(os.Args) < 4 {
		printUsageInfo()
		return
	}

	if dump == os.Args[1] {
		sourceDBPath := os.Args[2]
		targeFileName := os.Args[3]
		StartDump(sourceDBPath, targeFileName)
	} else if restore == os.Args[1] {
		sourceFileName := os.Args[2]
		targeDBPath := os.Args[3]
		StartRestore(sourceFileName, targeDBPath)
	} else {
		printUsageInfo()
		return
	}
}

func printUsageInfo() {
	fmt.Println("Usage 1: dbtool dump [source DB directory] [target directory name]")
	fmt.Println("Usage 2: dbtool restore [source directory name] [target DB directory]")
}
