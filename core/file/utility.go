/*
========================================================================================================================
Name        : core/task/utility.go
Author      : Andrea Graziani
Description : This file includes some utility function about file managing.
========================================================================================================================
*/
package file

import (
	"Word-Count/core"
	"Word-Count/core/utility"
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// This function is used to empty specified directory.
func EmptyDirectory(pDirectoryPath string) error {

	var mError error
	var mDirectory *os.File
	var mNames []string

	// Open specified directory and defer its closing...
	if mDirectory, mError = os.Open(pDirectoryPath); mError != nil {
		return mError
	}
	defer func() {
		utility.CheckPanicError(mDirectory.Close())
	}()

	// Reads and returns a slice of names from specified directory...
	if mNames, mError = mDirectory.Readdirnames(-1); mError != nil {
		return mError
	}

	// Removing...
	for _, name := range mNames {
		if mError = os.RemoveAll(filepath.Join(pDirectoryPath, name)); mError != nil {
			return mError
		}
	}
	return nil
}

// This function is used to split a specified file into "pNumber"-files.
// Files will be created into input file directory.
func SplitByWord(pInputFileDirectory string, pInputFileName string, pNumber int) error {

	var mInputFile *os.File
	var mError error

	// Open input file..
	// ====================================================================== //
	if mInputFile, mError = os.OpenFile(filepath.Join(pInputFileDirectory, pInputFileName), os.O_RDONLY, 0666); mError != nil {
		return mError
	}
	defer func() {
		utility.CheckPanicError(mInputFile.Close())
	}()

	// Creating output files and "*Writer" object...
	// ====================================================================== //
	mOutputFile := make([]*os.File, pNumber)                    // Allocation array used to store "*os.File" objects.
	mOutputWriters := make([]*bufio.Writer, pNumber)            // Allocation "*Writer" object to perform write operation.
	mReplacer := strings.NewReplacer(",", "", ".", "", ";", "") // Allocation "*Replacer" object to remove punctuation from string.

	for x := 0; x < pNumber; x++ {

		if mOutputFile[x], mError = os.OpenFile(filepath.Join(pInputFileDirectory, fmt.Sprintf("%s%d", core.DefaultFileName, x)), os.O_WRONLY|os.O_CREATE, 0666); mError != nil {
			return mError
		}
		mOutputWriters[x] = bufio.NewWriter(mOutputFile[x])
	}

	// Scanning input file...
	// ====================================================================== //
	mWordScanner := bufio.NewScanner(mInputFile)
	mWordScanner.Split(bufio.ScanWords)

	// Splitting...
	// ====================================================================== //
	for x := 0; mWordScanner.Scan(); x++ {

		// Restart cycle if necessary
		if x == pNumber {
			x = 0
		}

		// Get a word from file and change all his Unicode letters to their lower case removing punctuation...
		mCurrentWord := strings.ToLower(mWordScanner.Text())
		mCurrentWord = mReplacer.Replace(mCurrentWord)

		// Writing to file...
		if _, mError = mOutputWriters[x].WriteString(mCurrentWord + " "); mError != nil {
			return mError
		}

		// Flush...
		if mError = mOutputWriters[x].Flush(); mError != nil {
			return mError
		}
	}

	// Closing files...
	// ====================================================================== //
	for _, element := range mOutputFile {
		utility.CheckPanicError(element.Close())
	}

	return nil
}

// This function is used to send a file through network.
func Send(pInputFileName string, pFileHash string, pAddress string) error {

	var mError error
	var mInputFile *os.File
	var mInputFileInfo os.FileInfo
	var mBuffer []byte
	var mConnection net.Conn

	// Open input-file and connection; defer their close...
	// ====================================================================== //

	if mInputFile, mError = os.OpenFile(pInputFileName, os.O_RDONLY, 0666); mError != nil {
		return mError
	}
	defer func() {
		utility.CheckPanicError(mInputFile.Close())
	}()

	if mConnection, mError = net.Dial(core.DefaultNetwork, pAddress); mError != nil {
		return mError
	}
	defer func() {
		utility.CheckPanicError(mConnection.Close())
	}()

	// Getting "file-name" and "file-size" about file and send them to receiver...
	// ====================================================================== //
	if mInputFileInfo, mError = mInputFile.Stat(); mError != nil {
		return mError
	}

	// Getting file-name and file-size...
	mFileSize := fillString(strconv.FormatInt(mInputFileInfo.Size(), 10), 10)
	mFileName := fillString(mInputFileInfo.Name(), 64)
	mFileHash := fillString(pFileHash, 64)

	// Sending...
	if _, mError = mConnection.Write([]byte(mFileSize)); mError != nil {
		return mError
	}
	if _, mError = mConnection.Write([]byte(mFileName)); mError != nil {
		return mError
	}
	if _, mError = mConnection.Write([]byte(mFileHash)); mError != nil {
		return mError
	}

	// Sending File...
	// ====================================================================== //

	// Initialize a buffer for reading parts of the file in
	mBuffer = make([]byte, core.BufferSize)

	// Reading file...
	for {
		_, mError = mInputFile.Read(mBuffer)

		if mError == io.EOF {
			break
		}
		if mError != nil && mError != io.EOF {
			return mError
		}

		if _, mError = mConnection.Write(mBuffer); mError != nil {
			return mError
		}
	}
	return nil
}

// This function is used to receive a file from network.
func Receive(pListener net.Listener) (string, error) {

	var mOutputDirectory string
	var mOutputFile *os.File
	var mError error
	var mReceivedBytes int64
	var mFileSize int64
	var mFileName string
	var mFileHash string
	var mConnection net.Conn

	// Wait client...
	// ====================================================================== //
	mConnection, mError = pListener.Accept()
	utility.CheckPanicError(mError)

	// Getting "file-name" and "file-size" from sender...
	// ====================================================================== //
	mBufferHashFile := make([]byte, 64)
	mBufferFileName := make([]byte, 64)
	mBufferFileSize := make([]byte, 10)

	// Getting File-Name...
	if _, mError = mConnection.Read(mBufferFileSize); mError != nil {
		return "", mError
	}
	// Getting File-Size...
	if _, mError = mConnection.Read(mBufferFileName); mError != nil {
		return "", mError
	}
	// Getting Hash-Name...
	if _, mError = mConnection.Read(mBufferHashFile); mError != nil {
		return "", mError
	}

	// Strip the ':' from the received size, convert it to an int64...
	mFileSize, _ = strconv.ParseInt(strings.Trim(string(mBufferFileSize), ":"), 10, 64)
	// Strip the ':' from the received file name...
	mFileName = strings.Trim(string(mBufferFileName), ":")
	// Strip the ':' from the received file-hash
	mFileHash = strings.Trim(string(mBufferHashFile), ":")

	// Open/Create output directory...
	//
	// Use received "FileHash" to compute output directory.
	// If working directory already exist empty them, otherwise create it.
	// ====================================================================== //
	mOutputDirectory = filepath.Join(os.TempDir(), mFileHash)

	if _, mError = os.Stat(mOutputDirectory); mError == nil {
		utility.CheckPanicError(EmptyDirectory(mOutputDirectory))
	} else {
		utility.CheckPanicError(os.Mkdir(mOutputDirectory, os.ModePerm))
	}

	// Open/Create output-file and defer its closing...
	// ====================================================================== //
	if mOutputFile, mError = os.OpenFile(filepath.Join(mOutputDirectory, core.DefaultFileName), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666); mError != nil {
		return "", mError
	}
	defer func() {
		utility.CheckPanicError(mOutputFile.Close())
	}()

	// Start receiving...
	// ====================================================================== //
	for {
		if (mFileSize - mReceivedBytes) < core.BufferSize {

			if _, mError = io.CopyN(mOutputFile, mConnection, mFileSize-mReceivedBytes); mError != nil {
				panic(mError)
			}

			// Empty the remaining bytes that we don't need from the network buffer
			if _, mError = mConnection.Read(make([]byte, (mReceivedBytes+core.BufferSize)-mFileSize)); mError != nil {
				panic(mError)
			}
			break

		} else {
			if _, mError = io.CopyN(mOutputFile, mConnection, core.BufferSize); mError != nil {
				panic(mError)
			}
		}

		// Increment the counter...
		mReceivedBytes += core.BufferSize
	}

	return mFileName, nil
}

func fillString(retunString string, toLength int) string {
	for {
		lengtString := len(retunString)
		if lengtString < toLength {
			retunString = retunString + ":"
			continue
		}
		break
	}
	return retunString
}
