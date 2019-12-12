```
Creating a Block from scratch
Creating your own blocks is easy! Simply put your data in a file and run ipfs block put <yourfile> on it. Or, you can pipe your filedata into ipfs block put, like so:

$ echo "This is some data" | ipfs block put
QmfQ5QAjvg4GtA3wg3adpnDJug8ktA1BxurVqBD8rtgVjM
$ ipfs block get QmfQ5QAjvg4GtA3wg3adpnDJug8ktA1BxurVqBD8rtgVjM
This is some data
Note: When making your own block data, you won’t be able to read the data with ipfs cat. This is because you are inputting raw data without the unixfs data format. To read raw blocks, use ipfs block get as shown in the example.
```
