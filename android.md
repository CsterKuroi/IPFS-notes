# Android

* https://github.com/hazae41/sweet-ipfs
* https://github.com/ligi/IPFSDroid

```
        val env = arrayOf("IPFS_PATH=" + getRepoPath().absoluteFile)
        val command = getBinaryFile().absolutePath + " " + cmd

        return Runtime.getRuntime().exec(command, env)
```
