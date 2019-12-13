# ipfs-senc

* https://github.com/jbenet/ipfs-senc

* ipfs-senc share <path-to-file-or-directory>

```
BundleEncryptAndPut
  Bundle -> TarAndZip -> filepath.Walk
  Encrypt -> NewCTR -> MultiReader
  Put -> Add
```

* ipfs-senc --key <secret-key> download <ipfs-link> [<local-destination-dir>]

```
GetDecryptAndUnbundle
  Get
  Decrypt -> NewCTR
  Unbundle -> UnzipAndUntar -> MkdirAll -> OpenFile
```
