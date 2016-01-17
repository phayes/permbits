# PermBits

[![Coverage Status](https://coveralls.io/repos/phayes/permbits/badge.svg?branch=master&service=github)](https://coveralls.io/github/phayes/permbits?branch=master) 

Easy file permissions for golang. Easily get and set file permission bits. 

This package makes it a breeze to check and modify file permission bits in Linux, Mac, and other Unix systems. 

##Example

```golang
permissions, err := permbits.Stat("/path/to/my/file");
if err != nil {
  return err
}

// Check to make sure the group can write to the file
// If they can't write, update the permissions so they can
if !permissions.GroupWrite() {
  permissions.SetGroupWrite(true)
  err := permbits.Chmod("/path/to/my/file", permissions)
  if err != nil {
    return errors.New("error setting permission on file", err)
  }
}

// Also works well with os.File
fileInfo, err := file.Stat()
if err != nil {
  return err
}
fileMode := fileInfo.Mode()
permissions := permbits.FileMode(fileMode)

// Disable write access to the file for everyone but the user
permissions.SetGroupWrite(false)
permissions.SetOtherWrite(false)
permbits.UpdateFileMode(&fileMode, permissions)


```
